# tfis

Terraform import syntax

Usage example:

```bash
$ tfis google_datastore_index
==> google_datastore_index
Documentation URL: https://www.terraform.io/docs/providers/google/r/datastore_index.html
Import formats:
terraform import google_datastore_index.default projects/{{project}}/indexes/{{index_id}}
terraform import google_datastore_index.default {{project}}/{{index_id}}
terraform import google_datastore_index.default {{index_id}}
```
