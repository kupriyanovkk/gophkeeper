package storage

import (
	"testing"

	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
)

func TestGetPrivateData(t *testing.T) {
	storage := NewMemoryStorage()

	storage.SetPrivateLoginPass([]model.PrivateLoginPass{
		{Id: 1, Login: "login1", Password: "password1", Type: 0},
		{Id: 2, Login: "login2", Password: "password2", Type: 0},
	})

	storage.SetPrivateCard([]model.PrivateCard{
		{Id: 1, Title: "card1", CardNumber: "1111 2222 3333 4444", CVV: "123", Due: "2022-01-01", Type: 1},
		{Id: 2, Title: "card2", CardNumber: "5555 6666 7777 8888", CVV: "456", Due: "2022-02-01", Type: 1},
	})

	storage.SetPrivateText([]model.PrivateText{
		{Id: 1, Text: "text1", Type: 2},
		{Id: 2, Text: "text2", Type: 2},
	})

	t.Run("Retrieve private login pass data", func(t *testing.T) {
		result := storage.GetPrivateData(0)
		if len(result) != 2 {
			t.Errorf("Expected 2 private login pass data, but got %d", len(result))
		}
	})

	t.Run("Retrieve private card data", func(t *testing.T) {
		result := storage.GetPrivateData(1)
		if len(result) != 2 {
			t.Errorf("Expected 2 private card data, but got %d", len(result))
		}
	})

	t.Run("Retrieve private text data", func(t *testing.T) {
		result := storage.GetPrivateData(2)
		if len(result) != 2 {
			t.Errorf("Expected 2 private text data, but got %d", len(result))
		}
	})
}
