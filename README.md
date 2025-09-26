# serio

> 🌟 Modern, idiomatic Go serial port library.  
> Effortless, context-aware, stream-friendly serial communication, powered by [bugst/go-serial](https://github.com/bugst/go-serial).

[![Go](https://github.com/hawkli-1994/serio/actions/workflows/go.yml/badge.svg)](https://github.com/hawkli-1994/serio/actions/workflows/go.yml)
[![GitHub Pages](https://github.com/hawkli-1994/serio/actions/workflows/jekyll-gh-pages.yml/badge.svg)](https://github.com/hawkli-1994/serio/actions/workflows/jekyll-gh-pages.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hawkli-1994/serio)](https://goreportcard.com/report/github.com/hawkli-1994/serio)
[![GoDoc](https://godoc.org/github.com/hawkli-1994/serio?status.svg)](https://godoc.org/github.com/hawkli-1994/serio)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Repository:** [github.com/hawkli-1994/serio](https://github.com/hawkli-1994/serio)

## Features

- **io.ReadWriteCloser interface** — Use Go's standard IO tools (`io.Copy`, `bufio`, etc.`)
- **Context support** — Open, read, write operations support context for timeout/cancel
- **Modern configuration struct** — Easy, readable config
- **Cross-platform** — Supports Linux, Windows, and macOS on both x86 and ARM architectures
- **Automatic resource management** — Safe, defer-friendly
- **Timeout control** — Granular control over read/write timeouts with `SetWriteTimeout` and `SetDeadline`

## Documentation

- [API Reference](docs/api.md)
- [GitHub Pages](https://hawkli-1994.github.io/serio)

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/hawkli-1994/serio"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    port, err := serio.Open(ctx, serio.Config{
        Name:     "/dev/ttyUSB0",
        Baud:     115200,
        DataBits: 8,
        StopBits: 1,
        Parity:   serio.None,
        Timeout:  time.Second * 2,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer port.Close()

    _, err = port.Write([]byte("hello serial!"))
    if err != nil {
        log.Fatal(err)
    }

    buf := make([]byte, 128)
    n, err := port.Read(buf)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Received: %s\n", buf[:n])
}
```

## API Reference

### Opening a Port

To open a serial port, use the `Open` function with a context and configuration:

```go
port, err := serio.Open(ctx, serio.Config{
    Name:     "/dev/ttyUSB0",  // Port name
    Baud:     115200,          // Baud rate
    DataBits: 8,               // Data bits
    StopBits: 1,               // Stop bits
    Parity:   serio.None,      // Parity
    Timeout:  time.Second * 2, // Timeout
})
```

### Reading and Writing

The port implements `io.ReadWriteCloser`, so you can use standard Go IO operations:

```go
// Writing
n, err := port.Write([]byte("hello"))

// Reading
buf := make([]byte, 128)
n, err := port.Read(buf)
```

### Streaming Data

You can easily stream data using Go's `io.Copy`:

```go
// Copy stdin to serial port
go io.Copy(port, os.Stdin)
// Copy serial data to stdout
go io.Copy(os.Stdout, port)
```

### Timeout Control

serio provides granular timeout control:

```go
// Set a specific timeout for write operations
port.SetWriteTimeout(2 * time.Second)

// Set a deadline for both read and write operations
port.SetDeadline(time.Now().Add(5 * time.Second))
```

### Listing Ports

To get a list of available serial ports:

```go
ports, err := serio.ListPorts()
if err != nil {
    log.Fatal(err)
}
for _, name := range ports {
    fmt.Println("Available port:", name)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/hawkli-1994/serio/blob/main/LICENSE) file for details.

---

Powered by [bugst/go-serial](https://github.com/bugst/go-serial).