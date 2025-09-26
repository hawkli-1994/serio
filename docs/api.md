# API Reference

## Types

### Port

```go
type Port struct {
    // contains filtered or unexported fields
}
```

Port is a wrapper around the underlying serial port implementation that provides additional functionality like timeout control.

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

Config represents the configuration for opening a serial port.

### Parity

```go
type Parity serial.Parity

const (
    None  Parity = Parity(serial.NoParity)
    Odd   Parity = Parity(serial.OddParity)
    Even  Parity = Parity(serial.EvenParity)
    Mark  Parity = Parity(serial.MarkParity)
    Space Parity = Parity(serial.SpaceParity)
)
```

Parity represents the parity setting for a serial port.

## Functions

### Open

```go
func Open(ctx context.Context, cfg Config) (*Port, error)
```

Open creates and configures a new serial port. The context can be used to cancel the operation.

### ListPorts

```go
func ListPorts() ([]string, error)
```

ListPorts returns a list of available serial port names.

## Port Methods

### Read

```go
func (p *Port) Read(b []byte) (int, error)
```

Read reads data from the serial port. It implements the io.Reader interface.

### Write

```go
func (p *Port) Write(b []byte) (int, error)
```

Write writes data to the serial port. It implements the io.Writer interface. The write operation is subject to the write timeout or deadline if set.

### Close

```go
func (p *Port) Close() error
```

Close closes the serial port. It implements the io.Closer interface.

### SetWriteTimeout

```go
func (p *Port) SetWriteTimeout(t time.Duration) error
```

SetWriteTimeout sets the timeout for Write operations. If t is 0, no timeout will be applied. If t is negative, the operation will block indefinitely.

### SetDeadline

```go
func (p *Port) SetDeadline(t time.Time) error
```

SetDeadline sets the read and write deadlines associated with the port. It is equivalent to calling both SetReadDeadline and SetWriteDeadline. A zero value for t means I/O operations will not time out.