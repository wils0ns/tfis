# tfis

[![Go Report Card](https://goreportcard.com/badge/github.com/wils0ns/tfis?style=flat-square)](https://goreportcard.com/report/github.com/wils0ns/tfis)

Terraform import syntax

## Install

Using go:

```bash
git clone https://github.com/wils0ns/tfis.git
cd tfis
go install
```

Or download the binary for a particular [release](https://github.com/wils0ns/tfis/releases).

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
