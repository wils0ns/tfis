package main

import (
	"fmt"
	"os"

	"github.com/wils0ns/tfis/resource"
)

func printErrorAndExit(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tfis <RESOURCE_ADDRESS>")
		fmt.Println("Example: tfis google_project")
		os.Exit(0)
	}

	resource, err := resource.New(os.Args[1])
	printErrorAndExit(err)

	url, err := resource.DocsURL()
	printErrorAndExit(err)

	fmt.Println("==>", resource.Type)
	fmt.Println("Documentation URL:", url)

	syntaxes, err := resource.ImportSyntax()
	printErrorAndExit(err)

	fmt.Println("Import formats:")
	for _, s := range syntaxes {
		fmt.Println(s)
	}
}
