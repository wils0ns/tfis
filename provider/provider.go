package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

const terraformBaseURL = "https://www.terraform.io"
const providersURL = "https://registry.terraform.io/v2/providers"
const providerVersionsURL = "https://registry.terraform.io/v2/provider-versions"

// Provider of Resources
type Provider struct {
	ID         string `json:"id"`
	Attributes *Attributes
	versions   []*Version
}

// Attributes of Provider
type Attributes struct {
	Name     string
	FullName string `json:"full-name"`
	Source   string
}

// Version of provider
type Version struct {
	ID   string `json:"id"`
	docs []*Docs
}

// Docs of provider
type Docs struct {
	ID         string `json:"id"`
	Attributes *DocsAttributes
}

// DocsAttributes of provider
type DocsAttributes struct {
	Title, Category, Content string
}

// New Provider
func New(name string) (*Provider, error) {
	var data = &struct{ Data []Provider }{}
	client := resty.New()
	_, err := client.R().
		SetResult(data).
		SetQueryParams(map[string]string{
			"filter[name]":             name,
			"filter[without-versions]": "true",
		}).
		Get(providersURL)
	if err != nil {
		return nil, err
	}

	for _, d := range data.Data {
		return &d, nil
	}

	return nil, fmt.Errorf("Provider not found: %v", name)
}

// SelfLink of the Provider
func (p *Provider) SelfLink() string {
	return fmt.Sprintf("%v/%v", providersURL, p.ID)
}

// Versions of the Provider (will cache)
func (p *Provider) Versions() []*Version {
	if p.versions == nil {
		var data = &struct{ Included []*Version }{}
		client := resty.New()
		_, err := client.R().
			SetResult(data).
			SetQueryParam("include", "provider-versions").
			Get(p.SelfLink())
		if err != nil {
			return nil
		}
		p.versions = data.Included
	}
	return p.versions
}

// LatestVersion of the Provider
func (p *Provider) LatestVersion() *Version {
	var latest = &Version{}
	for _, v := range p.Versions() {
		id, _ := strconv.Atoi(v.ID)
		latestID, _ := strconv.Atoi(latest.ID)
		if id > latestID {
			latest = v
		}
	}
	return latest
}

// SelfLink of the Version
func (v *Version) SelfLink() string {
	return fmt.Sprintf("%v/%v", providerVersionsURL, v.ID)
}

// Docs of the Version (will cache)
func (v *Version) Docs() []*Docs {
	if v.docs == nil {
		var data = &struct{ Included []*Docs }{}
		client := resty.New()
		_, err := client.R().
			SetResult(data).
			SetQueryParam("include", "provider-docs").
			Get(v.SelfLink())
		if err != nil {
			return nil
		}
		v.docs = data.Included
	}
	return v.docs
}

// ResourceDocs of the Version by title
func (v *Version) ResourceDocs(name string) (*Docs, error) {
	for _, i := range v.Docs() {
		if i.Attributes.Category == "resources" && i.Attributes.Title == name {
			return i, nil
		}
	}
	return nil, fmt.Errorf("Resource not found: %v", name)
}

// ImportDocs of a Docs
func ImportDocs(d *Docs) ([]string, error) {
	md := parser.New()
	node := md.Parse([]byte(d.Attributes.Content))
	items := node.GetChildren()
	for _, i := range items {
		switch v := i.(type) {
		case *ast.CodeBlock:
			code := string(v.Literal)
			if strings.Contains(code, "terraform import") {
				is := []string{}
				for _, i := range strings.Split(code, "\n") {
					if i != "" {
						is = append(is, strings.Replace(i, "$ ", "", 1))
					}
				}
				return is, nil
			}
		}
	}
	return nil, fmt.Errorf("Import syntax for '%v' not found within docs", d.Attributes.Title)
}
