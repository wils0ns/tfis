package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var samples = []struct {
	name     string
	provider *Provider
	selfLink string
}{
	{
		"google",
		&Provider{
			ID: "354",
			Attributes: &Attributes{
				Name:     "google",
				FullName: "hashicorp/google",
				Source:   "https://github.com/hashicorp/terraform-provider-google",
			},
		},
		"https://registry.terraform.io/v2/providers/354",
	},
	{
		"digitalocean",
		&Provider{
			ID: "605",
			Attributes: &Attributes{
				Name:     "digitalocean",
				FullName: "digitalocean/digitalocean",
				Source:   "https://github.com/digitalocean/terraform-provider-digitalocean",
			},
		},
		"https://registry.terraform.io/v2/providers/605",
	},
}

func TestNew(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			p, err := New(s.name)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, s.provider, p)
		})
	}
}

func TestSelfLink(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			p, err := New(s.name)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, s.selfLink, p.SelfLink())
		})
	}
}

func TestLatestVersion(t *testing.T) {
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			p, err := New(s.name)
			if err != nil {
				t.Fatal(err)
			}
			assert.NotEmpty(t, p.LatestVersion().ID)
		})
	}
}

func TestLatestVersionDocs(t *testing.T) {

	p, err := New("digitalocean")
	if err != nil {
		t.Fatal(err)
	}
	docs := p.LatestVersion().Docs()
	assert.Less(t, 1, len(docs))

}

func TestResourceDocs(t *testing.T) {
	p, err := New("digitalocean")
	if err != nil {
		t.Fatal(err)
	}
	docs := p.LatestVersion().Docs()
	assert.Greater(t, len(docs), 1)

	_, err = p.LatestVersion().ResourceDocs("ssh_ke1y")
	assert.Error(t, err)

	res, err := p.LatestVersion().ResourceDocs("ssh_key")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, res.Attributes.Content)
}

func TestImportDocs(t *testing.T) {
	p, err := New("digitalocean")
	if err != nil {
		t.Fatal(err)
	}

	res, err := p.LatestVersion().ResourceDocs("ssh_key")
	if err != nil {
		t.Fatal(err)
	}

	is, err := ImportDocs(res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Greater(t, len(is), 0)
	assert.Contains(t, is[0], "terraform import digitalocean_ssh_key")

	p, err = New("google")
	if err != nil {
		t.Fatal(err)
	}

	res, err = p.LatestVersion().ResourceDocs("datastore_index")
	if err != nil {
		t.Fatal(err)
	}

	is, err = ImportDocs(res)
	if err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, len(is), 0)
	assert.NotContains(t, is[0], "$ terraform")

	res, err = p.LatestVersion().ResourceDocs("service_networking_connection")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ImportDocs(res)
	assert.Error(t, err)

}
