# serio

> 🌟 Modern, idiomatic Go serial port library.  
> Effortless, context-aware, stream-friendly serial communication, powered by [bugst/go-serial](https://github.com/bugst/go-serial).

[![Go](https://github.com/hawkli-1994/serio/actions/workflows/go.yml/badge.svg)](https://github.com/hawkli-1994/serio/actions/workflows/go.yml)
[![GitHub Pages](https://github.com/hawkli-1994/serio/actions/workflows/pages.yml/badge.svg)](https://github.com/hawkli-1994/serio/actions/workflows/pages.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hawkli-1994/serio)](https://goreportcard.com/report/github.com/hawkli-1994/serio)
[![GoDoc](https://godoc.org/github.com/hawkli-1994/serio?status.svg)](https://godoc.org/github.com/hawkli-1994/serio)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Repository:** [github.com/hawkli-1994/serio](https://github.com/hawkli-1994/serio)

## Features

- **io.ReadWriteCloser interface** — Use Go's standard IO tools (`io.Copy`, `bufio`, etc.)
- **Context support** — Open, read, write operations support context for timeout/cancel.
- **Modern configuration struct** — Easy, readable config.
- **Cross-platform** — Works on Linux, Windows, macOS, ARM.
- **Automatic resource management** — Safe, defer-friendly.
- **Timeout control** — Granular control over read/write timeouts with SetWriteTimeout and SetDeadline.
- **Event-driven (future)** — Ready for async data/error callbacks.

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

## List available ports

```go
ports, err := serio.ListPorts()
if err != nil {
    log.Fatal(err)
}
for _, name := range ports {
    fmt.Println("Available port:", name)
}
```

## Use with io.Copy (streaming data)

```go
// Copy stdin to serial port
go io.Copy(port, os.Stdin)
// Copy serial data to stdout
go io.Copy(os.Stdout, port)
```

## Timeout control

The library provides granular timeout control for serial operations:

```go
// Set a specific timeout for write operations
port.SetWriteTimeout(2 * time.Second)

// Set a deadline for both read and write operations
port.SetDeadline(time.Now().Add(5 * time.Second))

// Disable timeout for write operations
port.SetWriteTimeout(0)

// Remove deadline
port.SetDeadline(time.Time{})
```

## API

### Open

```go
func Open(ctx context.Context, cfg Config) (*Port, error)
```

### Port (implements io.ReadWriteCloser)

```go
type Port struct { ... }
func (p *Port) Read(b []byte) (int, error)
func (p *Port) Write(b []byte) (int, error)
func (p *Port) Close() error
func (p *Port) SetWriteTimeout(t time.Duration) error
func (p *Port) SetDeadline(t time.Time) error
```

### Config

```go
type Config struct {
    Name     string
    Baud     int
    DataBits int
    StopBits int
    Parity   Parity
    Timeout  time.Duration
}
```

## License

MIT

---

Inspired by [bugst/go-serial](https://github.com/bugst/go-serial).