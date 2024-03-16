package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kupriyanovkk/gophkeeper/internal/client/completer"
	"github.com/kupriyanovkk/gophkeeper/internal/client/executor"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"

func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
}

func main() {
	printBuildInfo()

	p := prompt.New(
		executor.NewExecutor().Execute,
		completer.NewCompleter().Complete,
		prompt.OptionTitle("GophKeeper"),
		prompt.OptionPrefix(">>>"),
	)

	p.Run()
}
