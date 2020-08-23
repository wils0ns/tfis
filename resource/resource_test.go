package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wils0ns/tfis/provider"
)

var googleProvider = &provider.Provider{
	ID: "354",
	Attributes: &provider.Attributes{
		Name:     "google",
		FullName: "hashicorp/google",
		Source:   "https://github.com/hashicorp/terraform-provider-google",
	},
}

var digitalOceanProvider = &provider.Provider{
	ID: "605",
	Attributes: &provider.Attributes{
		Name:     "digitalocean",
		FullName: "digitalocean/digitalocean",
		Source:   "https://github.com/digitalocean/terraform-provider-digitalocean",
	},
}

var samples = []struct {
	name           string
	resource       *Resource
	docsURL        string
	importSyntaxes []string
}{
	{
		"google_project",
		&Resource{
			Name:     "project",
			Type:     "google_project",
			Provider: googleProvider,
		},
		"https://www.terraform.io/docs/providers/google/r/google_project.html",
		[]string{
			"terraform import google_project.my_project your-project-id",
		},
	},
	{
		"google_datastore_index",
		&Resource{
			Name:     "datastore_index",
			Type:     "google_datastore_index",
			Provider: googleProvider,
		},
		"https://www.terraform.io/docs/providers/google/r/datastore_index.html",
		[]string{
			"terraform import google_datastore_index.default projects/{{project}}/indexes/{{index_id}}",
			"terraform import google_datastore_index.default {{project}}/{{index_id}}",
			"terraform import google_datastore_index.default {{index_id}}",
		},
	},
	{
		"google_service_networking_connection",
		&Resource{
			Name:     "service_networking_connection",
			Type:     "google_service_networking_connection",
			Provider: googleProvider,
		},
		"https://www.terraform.io/docs/providers/google/r/service_networking_connection.html",
		nil,
	},
	{
		"digitalocean_ssh_key",
		&Resource{
			Name:     "ssh_key",
			Type:     "digitalocean_ssh_key",
			Provider: digitalOceanProvider,
		},
		"https://www.terraform.io/docs/providers/digitalocean/r/ssh_key.html",
		[]string{
			"terraform import digitalocean_ssh_key.mykey 263654",
		},
	},
}

func TestNew(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			res, err := New(s.name)
			if err != nil {
				t.Fatal(t, err)
			}

			assert.Equal(t, s.resource, res)
		})
	}
}

func TestDocsURL(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			res, err := New(s.name)
			if err != nil {
				t.Fatal(t, err)
			}

			url, err := res.DocsURL()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, s.docsURL, url)
		})
	}
}

func TestImportSyntax(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			res, err := New(s.name)
			if err != nil {
				t.Fatal(t, err)
			}

			syntaxes, err := res.ImportSyntax()
			if err != nil {
				if s.importSyntaxes != nil {
					t.Fatal(err)
				}
				assert.Error(t, err)
			}
			assert.Equal(t, s.importSyntaxes, syntaxes)
		})
	}
}
