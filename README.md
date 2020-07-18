# tfis

Terraform import syntax

## Install

```bash
git clone https://github.com/wilson-codeminus/tfis.git
cd tfis
go install
```

## Use

```bash
$ tfis google_datastore_index
==> google_datastore_index
Documentation URL: https://www.terraform.io/docs/providers/google/r/datastore_index.html
Import formats:
terraform import google_datastore_index.default projects/{{project}}/indexes/{{index_id}}
terraform import google_datastore_index.default {{project}}/{{index_id}}
terraform import google_datastore_index.default {{index_id}}
```
