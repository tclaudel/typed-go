# typed

`typed` is a Go package providing a type-safe wrapper around the `sync.Map` data structure.

## Installation

To use `typed`, you need to have Go installed and set up on your machine. Then you can install the package using `go get`:

```sh
go get "github.com/tclaudel/typed"
```

## Usage

Import the package into your Go code:

```go
import "github.com/tclaudel/typed"
```

Here's a basic example demonstrating how to use `typed.SyncMap`:

```go
package main

import (
	"fmt"
	"github.com/tclaudel/typed"
)

func main() {
	// Create a new SyncMap
	m := typed.NewSyncMap[string]()

	// Store a key-value pair
	m.Store("key", "value")

	// Load the value for a key
	val, ok := m.Load("key")
	if ok {
		fmt.Println("Value found:", val)
	} else {
		fmt.Println("Value not found")
	}

	// Delete a key
	m.Delete("key")
}
```

## Features

- Type-safe wrapper around `sync.Map`
- Methods for safe concurrent map operations such as `Store`, `Load`, `Delete`, `Range`, etc.
- Support for comparing and swapping values atomically

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests on [GitHub](https://github.com/tclaudel/go-typed-map).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
