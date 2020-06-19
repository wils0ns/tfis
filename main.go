package main

import (
	"fmt"
	"os"

	resource "github.com/wilson-codeminus/tfis/resource"
)

func main() {
	resource := resource.New(os.Args[1])
	url, err := resource.GetDocUrl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("==>", resource.Type)
	fmt.Println("Documentation URL:", url)

	syntaxes, err := resource.GetImportSyntaxes(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Import formats:")
	for _, s := range syntaxes {
		fmt.Println(s)
	}
}
