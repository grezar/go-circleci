# go-circleci

[![CircleCI](https://circleci.com/gh/grezar/go-circleci/tree/main.svg?style=svg)](https://circleci.com/gh/grezar/go-circleci/tree/main)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This client supports the [CircleCI V2 API](https://circleci.com/docs/api/v2/).

Note this client is in beta. While I am using this client in my personal
projects, many of the methods are not yet used in real projects and have not
been fully tested. Therefore, this client may involve a lot of breaking changes
until it reaches v1.0. If you find any missing features or bugs, please kindly
report it via an Issue or Pull Request.

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

contexts, err := client.Contexts.List(context.Background(), circleci.ContextListOptions{
	OwnerSlug: circleci.String("org"),
})
if err != nil {
	log.Fatal(err)
}
```

## Documentation
TODO: Write code comments for Go Doc.

## Contribution
If you find any issues with this package, please report an Issue.

## LICENSE
[The MIT License (MIT)](https://github.com/grezar/go-circleci/blob/main/LICENSE)
