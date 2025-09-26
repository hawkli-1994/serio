package serio

import (
    "context"
    "errors"
    "go.bug.st/serial"
    "time"
)

type Port struct {
    p serial.Port
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
        if cfg.Timeout > 0 {
            deadline := time.Now().Add(cfg.Timeout)
            _ = p.SetReadTimeout(cfg.Timeout)
            _ = p.SetWriteTimeout(cfg.Timeout)
            _ = p.SetDeadline(deadline)
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