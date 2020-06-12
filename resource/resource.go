package resource

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const TerraformBaseUrl = "https://www.terraform.io"

type TerraformResource struct {
	Type     string
	Provider string
	Name     string
}

func GetProperties(resourceType string) []string {
	return strings.SplitN(resourceType, "_", 2)
}

func InitResource(resourceType string) *TerraformResource {
	tr := &TerraformResource{}
	tr.Type = resourceType

	props := GetProperties(tr.Type)
	tr.Provider = props[0]
	tr.Name = props[1]

	return tr
}

func (r *TerraformResource) GetDocUrl() (string, error) {

	possibleUrls := []string{
		TerraformBaseUrl + "/docs/providers/" + r.Provider + "/r/" + r.Name + ".html",
		TerraformBaseUrl + "/docs/providers/" + r.Provider + "/r/" + r.Type + ".html",
	}

	for _, url := range possibleUrls {
		resp, _ := http.Get(url)
		if resp.StatusCode == 200 {
			return url, nil
		}
	}

	return "", errors.New("Unable to find documentation for " + r.Type)

}

func (r *TerraformResource) GetImportSyntaxes() ([]string, error) {
	url, err := r.GetDocUrl()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	syntaxes := []string{}
	doc.Find("pre").Each(func(i int, item *goquery.Selection) {
		if strings.Contains(item.Text(), "terraform import "+r.Type) {
			syntaxes = strings.Split(item.Text(), "\n")
		}
	})

	if len(syntaxes) == 0 {
		return nil, errors.New("Unable to find import syntax in documentation")
	}
	return syntaxes, nil
}
