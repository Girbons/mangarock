# Mangarock Go API Wrapper

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
	client := mangarock.NewClient()

	info, err := client.Info("mrs-serie-35593")

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
