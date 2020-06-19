package resource

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetProperties(t *testing.T) {
	actual := GetProperties("google_datastore_index")
	expected := []string{"google", "datastore_index"}

	lenActual := len(actual)
	if lenActual != 2 {
		t.Fatalf("GetProperties is expected to return a slice of length 2. Got %v", lenActual)
	}

	for i, item := range actual {
		if expected[i] != item {
			t.Fatalf("Expected %v, got: %v", expected, actual)
		}
	}
}

func TestNew(t *testing.T) {
	actual := New("google_project")
	expected := TerraformResource{
		Type:     "google_project",
		Provider: "google",
		Name:     "project",
	}

	if *actual != expected {
		t.Errorf("Expected %v, got %v", expected, *actual)
	}
}

func TestGetDocUrl(t *testing.T) {
	samples := []struct {
		resType string
		url     string
	}{
		{"google_project", "https://www.terraform.io/docs/providers/google/r/google_project.html"},
		{"google_datastore_index", "https://www.terraform.io/docs/providers/google/r/datastore_index.html"},
	}

	for i := range samples {
		sample := samples[i]
		t.Run(sample.resType, func(t *testing.T) {
			t.Parallel()
			res := New(sample.resType)
			url, err := res.GetDocUrl()

			if err != nil {
				t.Error(err)
			}

			if sample.url != url {
				t.Errorf("Expected %v, got %v", sample.url, url)
			}
		})
	}
}

func TestGetDocUrlReturnsError(t *testing.T) {
	res := New("res_not_found")
	_, err := res.GetDocUrl()

	if err == nil {
		t.Fatal("Expected an error to be returned, got nil")
	}

	errType := fmt.Sprintf("%T", err)
	expectedErrType := "*resource.docNotFoundError"

	if errType != expectedErrType {
		t.Errorf("Expected error to be %v, got %T", expectedErrType, errType)
	}
}

func TestGetImportSyntaxes(t *testing.T) {
	t.Parallel()

	samples := []struct {
		resType                string
		fakeHTML               string
		expectedImportSyntaxes []string
	}{
		{
			"google_project",
			`<pre><code>$ terraform import google_project.my_project your-project-id
			</code></pre>`,
			[]string{"terraform import google_project.my_project your-project-id"},
		},
		{
			"google_datastore_index",
			`<pre class="highlight plaintext"><code>$ terraform import google_datastore_index.default projects/{{project}}/indexes/{{index_id}}
			$ terraform import google_datastore_index.default {{project}}/{{index_id}}
			$ terraform import google_datastore_index.default {{index_id}}
			</code></pre>`,
			[]string{
				"terraform import google_datastore_index.default projects/{{project}}/indexes/{{index_id}}",
				"terraform import google_datastore_index.default {{project}}/{{index_id}}",
				"terraform import google_datastore_index.default {{index_id}}",
			},
		},
	}

	for i := range samples {
		// NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		sample := samples[i]
		t.Run(sample.resType, func(t *testing.T) {
			t.Parallel()
			res := New(sample.resType)

			w := httptest.NewRecorder()
			w.WriteString(sample.fakeHTML)

			isyntaxes, err := res.GetImportSyntaxes(w.Body)
			if err != nil {
				t.Fatal(err)
			}

			for ii, syntax := range isyntaxes {
				if sample.expectedImportSyntaxes[ii] != syntax {
					t.Errorf("Expected %v, got %v", sample.expectedImportSyntaxes[ii], syntax)
				}
			}

		})
	}
}

func TestGetImportSyntaxesReturnsError(t *testing.T) {
	res := New("google_service_networking_connection")

	w := httptest.NewRecorder()
	w.WriteString("")

	_, err := res.GetImportSyntaxes(w.Body)

	if err == nil {
		t.Fatal("Expected an error to be returned, got nil")
	}

	errType := fmt.Sprintf("%T", err)
	expectedErrType := "*resource.importSyntaxNotFoundError"

	if errType != expectedErrType {
		t.Log(err)
		t.Errorf("Expected error to be %v, got %T", expectedErrType, errType)
	}
}
