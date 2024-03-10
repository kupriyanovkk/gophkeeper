package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/internal/client/storage"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PrivateService struct {
	storage storage.MemoryAbstractStorage
	ctx     *model.GlobalContext
	client  pb.PrivateClient
	crypt   crypt.CryptAbstract
	sync    storage.SyncAbstract
}

// GetPrivateData retrieves private data by ID.
//
// It takes an integer ID as a parameter and returns an interface{} and an error.
func (s *PrivateService) GetPrivateData(id int) (interface{}, error) {
	if data, ok := s.storage.FindPrivateData(id); ok {
		return data, nil
	}

	response, err := s.client.GetPrivateData(s.ctx.Ctx, &pb.GetPrivateDataRequest{Id: uint32(id)})
	if err != nil {
		return nil, err
	}

	decoded, err := s.crypt.Decode(string(response.Content))
	if err != nil {
		return nil, err
	}

	var privateData interface{}
	switch response.Type {
	case 1:
		privateData = &model.PrivateLoginPass{}
	case 2:
		privateData = &model.PrivateText{}
	case 4:
		privateData = &model.PrivateCard{}
	}

	if err = json.Unmarshal([]byte(decoded), privateData); err != nil {
		return nil, err
	}

	return privateData, nil
}

// GetPrivateDataList retrieves a list of private data for the given id.
//
// Parameter(s): id int
// Return type(s): []*pb.PrivateData, error
func (s *PrivateService) GetPrivateDataList(id int) ([]*pb.PrivateData, error) {
	list := s.storage.GetPrivateData(storage.PrivateType(id))
	if len(list) > 0 {
		return list, nil
	}

	res, err := s.client.GetPrivateDataByType(s.ctx.Ctx, &pb.GetPrivateDataByTypeRequest{
		TypeId: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

// GetPrivateBinary retrieves private binary data and writes it to the specified destination file.
//
// Parameters:
// - id int: the ID of the private binary data to retrieve.
// - dst string: the destination file path to write the retrieved binary data.
// Return type: error
func (s *PrivateService) GetPrivateBinary(id int, dst string) error {
	res, err := s.client.GetPrivateData(s.ctx.Ctx, &pb.GetPrivateDataRequest{Id: uint32(id)})
	if err != nil {
		return err
	}

	if res.Type != 3 {
		return errors.New("invalid type")
	}

	decoded, err := s.crypt.Decode(string(res.Content))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(decoded))

	return err
}

// CreatePrivate is a function to create a private record.
//
// It takes in title (string), recordTypeID (int), and contentStr (string) as parameters.
// It returns an error.
func (s *PrivateService) CreatePrivate(title string, recordTypeID int, contentStr string) error {
	content := []byte(s.crypt.Encode(contentStr))

	_, err := s.client.CreatePrivateData(s.ctx.Ctx, &pb.CreatePrivateDataRequest{
		Content: content,
		Type:    uint32(recordTypeID),
		Title:   title,
	})

	s.sync.SyncAll()

	return err
}

func (s *PrivateService) EditPrivate(id int, title string, recordType int, strContent string, isForce bool) error {
	var updated time.Time
	localPrivate, _ := s.GetPrivateData(id)

	switch v := localPrivate.(type) {
	case *model.PrivateLoginPass:
		updated = v.Updated
	case *model.PrivateText:
		updated = v.Updated
	case *model.PrivateCard:
		updated = v.Updated
	}

	content := []byte(s.crypt.Encode(strContent))
	_, err := s.client.EditPrivateData(s.ctx.Ctx, &pb.EditPrivateDataRequest{
		Id:      uint32(id),
		Content: content,
		Type:    uint32(recordType),
		Title:   title,
		IsForce: isForce,
		Updated: timestamppb.New(updated),
	})
	if err != nil {
		return err
	}

	fmt.Println("private data successfully edited")
	s.sync.SyncAll()

	return nil
}

// DeletePrivate deletes private data.
//
// Takes an integer id as a parameter and returns an error.
func (s *PrivateService) DeletePrivate(id int) error {
	_, err := s.client.DeletePrivateData(s.ctx.Ctx, &pb.DeletePrivateDataRequest{
		Id: uint32(id),
	})
	if err != nil {
		return err
	}

	s.storage.ResetStorage()
	s.sync.SyncAll()

	return nil
}

// NewPrivateService creates a new PrivateService.
//
// It takes in a MemoryAbstractStorage, GlobalContext pointer, PrivateClient, CryptAbstract, and SyncAbstract and returns a pointer to PrivateService.
func NewPrivateService(storage storage.MemoryAbstractStorage, ctx *model.GlobalContext, client pb.PrivateClient, crypt crypt.CryptAbstract, sync storage.SyncAbstract) *PrivateService {
	return &PrivateService{
		storage: storage,
		ctx:     ctx,
		client:  client,
		crypt:   crypt,
		sync:    sync,
	}
}
