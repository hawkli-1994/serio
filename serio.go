package serio

import (
	"context"
	// "errors"
	"go.bug.st/serial"
	"time"
)

type Parity serial.Parity

const (
	None  Parity = Parity(serial.NoParity)
	Odd   Parity = Parity(serial.OddParity)
	Even  Parity = Parity(serial.EvenParity)
	Mark  Parity = Parity(serial.MarkParity)
	Space Parity = Parity(serial.SpaceParity)
)

type Config struct {
	Name     string
	Baud     int
	DataBits int
	StopBits int
	Parity   Parity
	Timeout  time.Duration
}

type Port struct {
	p            serial.Port
	writeTimeout time.Duration
	deadline     time.Time
}

func Open(ctx context.Context, cfg Config) (*Port, error) {
	mode := &serial.Mode{
		BaudRate: cfg.Baud,
		DataBits: cfg.DataBits,
		StopBits: serial.StopBits(cfg.StopBits),
		Parity:   serial.Parity(cfg.Parity),
	}

	// 使用通道配合 context 控制 open
	result := make(chan struct {
		port serial.Port
		err  error
	}, 1)

	go func() {
		p, err := serial.Open(cfg.Name, mode)
		if err != nil {
			result <- struct {
				port serial.Port
				err  error
			}{nil, err}
			return
		}
		// 串口打开成功，设置超时
		port := &Port{p: p}
		if cfg.Timeout > 0 {
			port.writeTimeout = cfg.Timeout
			port.deadline = time.Now().Add(cfg.Timeout)
			_ = p.SetReadTimeout(cfg.Timeout)
		}
		result <- struct {
			port serial.Port
			err  error
		}{p, nil}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err() // 超时或取消
	case r := <-result:
		if r.err != nil {
			return nil, r.err
		}
		return &Port{p: r.port}, nil
	}
}

// Implements io.Reader
func (p *Port) Read(b []byte) (int, error) {
	// Check deadline
	if !p.deadline.IsZero() && time.Now().After(p.deadline) {
		return 0, &serial.PortError{}
	}

	// Set read timeout based on remaining deadline
	if !p.deadline.IsZero() {
		remaining := time.Until(p.deadline)
		if remaining > 0 {
			_ = p.p.SetReadTimeout(remaining)
		} else {
			return 0, &serial.PortError{}
		}
	}

	return p.p.Read(b)
}

// Implements io.Writer
func (p *Port) Write(b []byte) (int, error) {
	// Check deadline
	if !p.deadline.IsZero() && time.Now().After(p.deadline) {
		return 0, &serial.PortError{}
	}

	// Create channel for result
	result := make(chan struct {
		n   int
		err error
	}, 1)

	// Create context with timeout or deadline
	var ctx context.Context
	var cancel context.CancelFunc

	if !p.deadline.IsZero() {
		ctx, cancel = context.WithDeadline(context.Background(), p.deadline)
	} else if p.writeTimeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), p.writeTimeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	// Perform write in goroutine
	go func() {
		n, err := p.p.Write(b)
		result <- struct {
			n   int
			err error
		}{n, err}
	}()

	// Wait for result or context cancellation
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case r := <-result:
		return r.n, r.err
	}
}

// Implements io.Closer
func (p *Port) Close() error {
	return p.p.Close()
}

// SetWriteTimeout sets the timeout for Write operations
func (p *Port) SetWriteTimeout(t time.Duration) error {
	p.writeTimeout = t
	return nil
}

// SetDeadline sets the read and write deadlines for the port
func (p *Port) SetDeadline(t time.Time) error {
	p.deadline = t
	return nil
}
