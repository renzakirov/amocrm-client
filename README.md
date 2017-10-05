<p align="center">
  <h3 align="center">A Golang client libary for AmoCRM API</h3>
  <p align="center">
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/esporo-org/amocrm-client"><img alt="Travis" src="https://img.shields.io/travis/esporo-org/amocrm-client.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/esporo-org/amocrm-client"><img alt="Codecov branch" src="https://img.shields.io/codecov/c/github/esporo-org/amocrm-client/master.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/esporo-org/amocrm-client"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/esporo-org/amocrm-client?style=flat-square"></a>
  <a href="http://godoc.org/github.com/esporo-org/amocrm-client"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
</p>

---

# amocrm-client


## AmoCRM API description

You can find the AmoCRM API for next [link](https://developers.amocrm.com/rest_api/)

## Installation

```shell
go get -u github.com/esporo-org/amocrm-client
```
### For glide
```shell
glide get github.com/esporo-org/amocrm-client
```
### For dep
```shell
dep ensure -add github.com/esporo-org/amocrm-client
```

## Getting Started

You can run this code in examples directory

```
package main

import (
	"os"
    "log"
    "github.com/esporo-org/amocrm-client"
)


func main() {
  // TODO write examples code
}
```

## How run examples 

## Tests

The test files are in the directory `tests`. 
To run tests code you need:

1. Rename file `env.example.json` to `env.json`

2. Write down your values domain name and authorization keys

3. Run:
```go test github.com/esporo-org/amocrm-client/tests``` or ```go test -v github.com/esporo-org/amocrm-client/tests```
