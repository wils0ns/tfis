package main

import (
	"fmt"
	"os"

	resource "github.com/wils0ns/tfis/resource"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tfis <RESOURCE_ADDRESS>")
		fmt.Println("Example: tfis google_project")
		os.Exit(0)
	}

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
