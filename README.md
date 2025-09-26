# serio

> 🌟 Modern, idiomatic Go serial port library.  
> Effortless, context-aware, stream-friendly serial communication, powered by [bugst/go-serial](https://github.com/bugst/go-serial).

**Repository:** [github.com/hawkli-1994/serio](https://github.com/hawkli-1994/serio)

## Features

- **io.ReadWriteCloser interface** — Use Go's standard IO tools (`io.Copy`, `bufio`, etc.)
- **Context support** — Open, read, write operations support context for timeout/cancel.
- **Modern configuration struct** — Easy, readable config.
- **Cross-platform** — Works on Linux, Windows, macOS, ARM.
- **Automatic resource management** — Safe, defer-friendly.
- **Event-driven (future)** — Ready for async data/error callbacks.

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