package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/serialx/serialx"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // 枚举端口
    ports, _ := serialx.ListPorts()
    fmt.Println("Available ports:", ports)

    // 打开串口
    port, err := serialx.Open(ctx, serialx.Config{
        Name:     "/dev/ttyUSB0",
        Baud:     115200,
        DataBits: 8,
        StopBits: 1,
        Parity:   serialx.None,
        Timeout:  time.Second * 3,
    })
    if err != nil {
        log.Fatal("Open error:", err)
    }
    defer port.Close()

    // 写数据
    _, err = port.Write([]byte("hello serial"))
    if err != nil {
        log.Fatal("Write error:", err)
    }

    // 读数据
    buf := make([]byte, 128)
    n, err := port.ReadWithContext(ctx, buf)
    if err != nil {
        log.Fatal("Read error:", err)
    }
    fmt.Printf("Received: %s\n", buf[:n])

    // 异步事件监听
    port.OnData(func(data []byte) {
        fmt.Printf("Async Received: %x\n", data)
    })
    port.OnError(func(err error) {
        log.Printf("Serial error: %v", err)
    })

    // 保持主线程
    select {}
}