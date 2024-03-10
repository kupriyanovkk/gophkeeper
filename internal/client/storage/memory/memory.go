package storage

import (
	"errors"
	"sync"

	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/internal/client/storage"
	"github.com/kupriyanovkk/gophkeeper/proto"
)

type MemoryStorage struct {
	mutex            sync.RWMutex
	PrivateLoginPass map[int]model.PrivateLoginPass
	PrivateCard      map[int]model.PrivateCard
	PrivateText      map[int]model.PrivateText
}

// SetPrivateLoginPass sets the private login pass in the MemoryStorage.
//
// data []model.PrivateLoginPass
func (s *MemoryStorage) SetPrivateLoginPass(data []model.PrivateLoginPass) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, d := range data {
		s.PrivateLoginPass[d.Id] = d
	}
}

// GetPrivateLoginPass returns the private login pass for the given ID, along with a boolean indicating if the data was found, and an error if applicable.
//
// Parameter(s): id int
// Return type(s): model.PrivateLoginPass, bool, error
func (s *MemoryStorage) GetPrivateLoginPass(id int) (model.PrivateLoginPass, bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, ok := s.PrivateLoginPass[id]
	if !ok {
		return model.PrivateLoginPass{}, false, errors.New("private login pass not found")
	}

	return data, ok, nil
}

// SetPrivateCard sets private card data in MemoryStorage.
//
// data []model.PrivateCard
func (s *MemoryStorage) SetPrivateCard(data []model.PrivateCard) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, d := range data {
		s.PrivateCard[d.Id] = d
	}
}

// GetPrivateCard retrieves the private card from memory storage.
//
// It takes an integer id as a parameter and returns a model.PrivateCard, a boolean, and an error.
func (s *MemoryStorage) GetPrivateCard(id int) (model.PrivateCard, bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, ok := s.PrivateCard[id]
	if !ok {
		return model.PrivateCard{}, false, errors.New("private card not found")
	}

	return data, ok, nil
}

// SetPrivateText sets the private text in the MemoryStorage.
//
// It takes a slice of model.PrivateText as parameter.
func (s *MemoryStorage) SetPrivateText(data []model.PrivateText) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, d := range data {
		s.PrivateText[d.Id] = d
	}
}

// GetPrivateText retrieves private text from MemoryStorage.
//
// It takes an integer id as a parameter and returns a model.PrivateText, a boolean, and an error.
func (s *MemoryStorage) GetPrivateText(id int) (model.PrivateText, bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, ok := s.PrivateText[id]
	if !ok {
		return model.PrivateText{}, false, errors.New("private text not found")
	}

	return data, ok, nil
}

// FindPrivateData searches for private data in the MemoryStorage.
//
// It takes an integer ID as a parameter and returns an interface and a boolean.
func (s *MemoryStorage) FindPrivateData(id int) (interface{}, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if _, ok := s.PrivateLoginPass[id]; ok {
		return s.PrivateLoginPass[id], ok
	}
	if _, ok := s.PrivateCard[id]; ok {
		return s.PrivateCard[id], ok
	}
	if _, ok := s.PrivateText[id]; ok {
		return s.PrivateText[id], ok
	}

	return nil, false
}

// ResetStorage resets the MemoryStorage.
//
// No parameters.
func (s *MemoryStorage) ResetStorage() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.PrivateLoginPass = make(map[int]model.PrivateLoginPass)
	s.PrivateCard = make(map[int]model.PrivateCard)
	s.PrivateText = make(map[int]model.PrivateText)
}

func (s *MemoryStorage) GetPrivateData(privateType storage.PrivateType) []*proto.PrivateData {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var privateData []*proto.PrivateData

	switch privateType {
	case storage.PrivateLoginPass:
		for _, pd := range s.PrivateLoginPass {
			privateData = append(privateData, &proto.PrivateData{
				Id:    uint32(pd.Id),
				Title: pd.Title,
			})
		}
	case storage.PrivateCard:
		for _, pd := range s.PrivateCard {
			privateData = append(privateData, &proto.PrivateData{
				Id:    uint32(pd.Id),
				Title: pd.Title,
			})
		}
	case storage.PrivateText:
		for _, pd := range s.PrivateText {
			privateData = append(privateData, &proto.PrivateData{
				Id:    uint32(pd.Id),
				Title: pd.Title,
			})
		}
	}

	return privateData
}

// NewMemoryStorage creates a new MemoryStorage instance.
//
// No parameters.
// Returns a pointer to MemoryStorage.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		PrivateLoginPass: make(map[int]model.PrivateLoginPass),
		PrivateCard:      make(map[int]model.PrivateCard),
		PrivateText:      make(map[int]model.PrivateText),
	}
}
