package resource

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wils0ns/tfis/provider"
)

const terraformBaseURL = "https://www.terraform.io"

// Resource of Terraform
type Resource struct {
	Type     string
	Provider *provider.Provider
	Name     string
}

// New terraform resource
func New(name string) (*Resource, error) {
	r := &Resource{}
	r.Type = name
	props := strings.SplitN(name, "_", 2)
	p, err := provider.New(props[0])
	if err != nil {
		return nil, err
	}
	r.Name = props[1]
	r.Provider = p
	return r, nil
}

// DocsURL of the Resource
func (r *Resource) DocsURL() (string, error) {
	possibleUrls := []string{
		fmt.Sprintf("%v/docs/providers/%v/r/%v.html", terraformBaseURL, r.Provider.Attributes.Name, r.Name),
		fmt.Sprintf("%v/docs/providers/%v/r/%v.html", terraformBaseURL, r.Provider.Attributes.Name, r.Type),
	}

	for _, url := range possibleUrls {
		resp, _ := http.Get(url)
		if resp.StatusCode == 200 {
			return url, nil
		}
	}
	return "", fmt.Errorf("Unable to find documentions for '%v'", r.Name)
}

// ImportSyntax of the Resource
func (r *Resource) ImportSyntax() ([]string, error) {
	resDocs, err := r.Provider.LatestVersion().ResourceDocs(r.Name)
	if err != nil {
		resDocs, err = r.Provider.LatestVersion().ResourceDocs(r.Type)
		if err != nil {
			return nil, err
		}
	}
	importDocs, err := provider.ImportDocs(resDocs)
	if err != nil {
		return nil, fmt.Errorf("Import syntax not found for '%v' resource", r.Type)
	}
	return importDocs, nil
}
