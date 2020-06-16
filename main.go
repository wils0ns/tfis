package main

import (
	"fmt"
	"os"

	resource "github.com/wilson-codeminus/terraform-import-syntax/resource"
)

func main() {
	resource := resource.New(os.Args[1])
	url, err := resource.GetDocUrl()
	if err != nil {
		PrintErrorAndExist(err)
	}
	fmt.Println("==>", resource.Type)
	fmt.Println("Documentation URL:", url)

	syntaxes, err := resource.GetImportSyntaxes()
	if err != nil {
		PrintErrorAndExist(err)
	}

	fmt.Println("Import formats:")
	for _, s := range syntaxes {
		fmt.Println(s)
	}
}

func PrintErrorAndExist(e error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", e)
	os.Exit(1)
}
