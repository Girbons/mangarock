# Mangarock Go API Wrapper

[![Build Status](https://travis-ci.org/Girbons/mangarock.svg?branch=master)](https://travis-ci.org/Girbons/mangarock)
[![Coverage Status](https://coveralls.io/repos/github/Girbons/mangarock/badge.svg?branch=dev)](https://coveralls.io/github/Girbons/mangarock?branch=dev)
[![GoDoc](https://godoc.org/github.com/Girbons/mangarock?status.svg)](https://godoc.org/github.com/Girbons/mangarock)
[![Go Report Card](https://goreportcard.com/badge/github.com/Girbons/mangarock)](https://goreportcard.com/report/github.com/Girbons/mangarock)

## Getting started

### Installing

```
go get github.com/Girbons/mangarock
```

### Example

```go
package main

import (
	"fmt"

	"github.com/Girbons/mangarock"
)

func main() {
    // there are 2 ways to initialize a client
    // the first one is using `NewClient`
    client := mangarock.NewClient()

    //the other one is using `NewClientWithOptions`
    options := make(map[string]string{"country":"your country"})
    client := mangarock.NewClientWithOptions(options)

    info, err := client.Info("mrs-serie-35593")
    // SetOptions is still available
    // client.SetOptions(options)
    if err != nil {
        // do something
    }

    fmt.Println(info.Data.Author)
}
```

### Available Methods

- Info("mrs-series-id")
- Pages("mrs-chapter-id")

## Contribuiting

Feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
