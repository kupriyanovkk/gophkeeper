package completer

import "github.com/c-bata/go-prompt"

type Completer struct{}

func NewCompleter() *Completer {
	return &Completer{}
}

// Complete returns a list of suggestions for the given prompt document.
//
// It takes a prompt.Document parameter and returns a slice of prompt.Suggest.
func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "login", Description: "Login to the service"},
		{Text: "register", Description: "Register to the service"},
		{Text: "create-login-pass", Description: "Create login and password"},
		{Text: "create-card", Description: "Create card"},
		{Text: "create-text", Description: "Create text"},
		{Text: "create-file", Description: "Create file"},
		{Text: "get-private", Description: "Get private data"},
		{Text: "get-private-by-type", Description: "Get private data by type"},
		{Text: "get-private-binary", Description: "Get private data binary"},
		{Text: "delete-private", Description: "Delete private data"},
		{Text: "update-private", Description: "Update private data"},
		{Text: "exit", Description: "Exit the program"},
	}

	word := d.GetWordBeforeCursor()

	return prompt.FilterHasPrefix(suggestions, word, true)
}
