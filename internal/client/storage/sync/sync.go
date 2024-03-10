package storage

import (
	"encoding/json"

	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/internal/client/storage"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
)

type Sync struct {
	storage storage.MemoryAbstractStorage
	client  pb.PrivateClient
	ctx     *model.GlobalContext
	crypt   crypt.CryptAbstract
}

// SyncPassLoginData syncs the private login data.
//
// No parameters.
// Returns an error.
func (s *Sync) SyncPassLoginData() error {
	response, err := s.client.GetPrivateDataByType(s.ctx.Ctx, &pb.GetPrivateDataByTypeRequest{
		TypeId: uint32(storage.PrivateLoginPass),
	})
	if err != nil {
		return err
	}

	var list []model.PrivateLoginPass
	for _, data := range response.Data {
		decoded, err := s.crypt.Decode(string(data.Content))
		if err != nil {
			return err
		}

		var entry model.PrivateLoginPass
		if err := json.Unmarshal([]byte(decoded), &entry); err != nil {
			return err
		}

		entry.Id = int(data.Id)
		entry.Updated = data.Updated.AsTime()

		list = append(list, entry)
	}

	s.storage.SetPrivateLoginPass(list)
	return nil
}

// SyncCardData synchronizes card data.
//
// No parameters.
// Returns an error.
func (s *Sync) SyncCardData() error {
	cards, err := s.client.GetPrivateDataByType(s.ctx.Ctx, &pb.GetPrivateDataByTypeRequest{TypeId: uint32(storage.PrivateCard)})
	if err != nil {
		panic(err)
	}

	var list []model.PrivateCard
	for _, card := range cards.Data {
		decoded, errDecode := s.crypt.Decode(string(card.Content))
		if errDecode != nil {
			return errDecode
		}

		entry := model.PrivateCard{}
		errUnmarshal := json.Unmarshal([]byte(decoded), &entry)
		if errUnmarshal != nil {
			return errUnmarshal
		}
		entry.Id = int(card.Id)
		entry.Updated = card.Updated.AsTime()

		list = append(list, entry)
	}

	s.storage.SetPrivateCard(list)

	return nil
}

// SyncTextData syncs text data from the client and stores it in the storage.
//
// No parameters.
// Returns an error.
func (s *Sync) SyncTextData() error {
	texts, err := s.client.GetPrivateDataByType(s.ctx.Ctx, &pb.GetPrivateDataByTypeRequest{TypeId: uint32(storage.PrivateText)})
	if err != nil {
		panic(err)
	}

	var list []model.PrivateText
	for _, text := range texts.Data {
		entry := model.PrivateText{}

		decoded, errDecode := s.crypt.Decode(string(text.Content))
		if errDecode != nil {
			return errDecode
		}

		errUnmarshal := json.Unmarshal([]byte(decoded), &entry)
		if errUnmarshal != nil {
			return errUnmarshal
		}

		entry.Id = int(text.Id)
		entry.Updated = text.Updated.AsTime()

		list = append(list, entry)
	}

	s.storage.SetPrivateText(list)

	return nil
}

// SyncAll synchronizes all data.
func (s *Sync) SyncAll() {
	s.SyncPassLoginData()
	s.SyncCardData()
	s.SyncTextData()
}

// NewSync initializes and returns a new Sync struct.
//
// storage: memory abstract storage
// client: private client
// ctx: global context pointer
// crypt: crypt abstract
// *Sync
func NewSync(storage storage.MemoryAbstractStorage, client pb.PrivateClient, ctx *model.GlobalContext, crypt crypt.CryptAbstract) *Sync {
	return &Sync{
		storage: storage,
		client:  client,
		ctx:     ctx,
		crypt:   crypt,
	}
}
