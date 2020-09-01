# tfis

[![GoDoc](https://img.shields.io/badge/go-docs-blue.svg)](http://godoc.org/github.com/wils0ns/tfis)
[![Go Report Card](https://goreportcard.com/badge/github.com/wils0ns/tfis)](https://goreportcard.com/report/github.com/wils0ns/tfis)
[![Codecov](https://img.shields.io/codecov/c/github/wils0ns/tfis.svg)](https://codecov.io/gh/wils0ns/tfis)

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
