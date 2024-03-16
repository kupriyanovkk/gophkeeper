package storage

import (
	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/proto"
)

type MemoryAbstractStorage interface {
	GetPrivateLoginPass(id int) (model.PrivateLoginPass, bool, error)
	SetPrivateLoginPass([]model.PrivateLoginPass)
	GetPrivateCard(id int) (model.PrivateCard, bool, error)
	SetPrivateCard([]model.PrivateCard)
	GetPrivateText(id int) (model.PrivateText, bool, error)
	SetPrivateText([]model.PrivateText)
	FindPrivateData(id int) (interface{}, bool)
	GetPrivateData(privateType PrivateType) []*proto.PrivateData
	ResetStorage()
}

type SyncAbstract interface {
	SyncAll()
	SyncPassLoginData() error
	SyncCardData() error
	SyncTextData() error
}

type PrivateType int

const (
	PrivateLoginPass PrivateType = iota
	PrivateCard
	PrivateText
	PrivateFile
)
