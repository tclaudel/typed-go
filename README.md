# typed

`typed` is a Go package providing a type-safe wrapper around the `sync.Map` data structure.

## Installation

To use `typed`, you need to have Go installed and set up on your machine. Then you can install the package using `go get`:

```sh
go get "github.com/tclaudel/typed-go"
```

## Usage

Import the package into your Go code:

```go
import "github.com/tclaudel/typed-go"
```
### Using typed.SyncMap

Here's a basic example demonstrating how to use typed.SyncMap:

```go
package main

import (
	"fmt"
	"github.com/tclaudel/typed-go"
)

func main() {
	m := typed.NewSyncMap[string]()
	m.Store("key", "value")

	val, ok := m.Load("key")
	if ok {
		fmt.Println("Value found:", val)
	} else {
		fmt.Println("Value not found")
	}

	m.Delete("key")
}

```

### Using typed.List

Here's how you can use the doubly linked list implementation:

```go
package main

import (
	"fmt"
	"github.com/tclaudel/typed-go"
)

func main() {
	l := typed.NewList[int]()
	
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value())
	}
}
```

## Features

- Type-safe wrappers around sync.Map and doubly linked lists
- Methods for safe concurrent map operations such as Store, Load, Delete, Range, etc.
- Support for comparing and swapping values atomically in sync.Map
- Operations for manipulation of doubly linked lists including insertion, removal, and iteration

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests on [GitHub](https://github.com/tclaudel/go-typed-map).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
