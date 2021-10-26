# go-circleci

[![CircleCI](https://circleci.com/gh/grezar/go-circleci/tree/main.svg?style=svg)](https://circleci.com/gh/grezar/go-circleci/tree/main)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This client supports the [CircleCI V2 API](https://circleci.com/docs/api/v2/).

## Installation

Installation can be done with a normal `go get`:

```sh
go get -u github.com/grezar/go-circleci
```

## Usage

```go
import "github.com/grezar/go-ciecleci"
```

Construct a new CircleCI client, then use the various services on the client to
access different parts of the CircleCI API. For example, to list all contexts:

```go
config := circleci.DefaultConfig()
config.Token = "put-your-circleci-token-here"

client, err := circleci.NewClient(config)
if err != nil {
	log.Fatal(err)
}

contexts, err := client.Contexts.List(context.Background(), ContextListOptions{
	OwnerSlug: circleci.String("org"),
})
if err != nil {
	log.Fatal(err)
}
```
