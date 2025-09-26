package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/hawkli-1994/serio"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // 枚举端口
    ports, _ := serio.ListPorts()
    fmt.Println("Available ports:", ports)

    // 打开串口
    port, err := serio.Open(ctx, serio.Config{
        Name:     "/dev/ttyUSB0",
        Baud:     115200,
        DataBits: 8,
        StopBits: 1,
        Parity:   serio.None,
        Timeout:  time.Second * 3,
    })
    if err != nil {
        log.Fatal("Open error:", err)
    }
    defer port.Close()

    // 设置写超时
    port.SetWriteTimeout(2 * time.Second)
    
    // 设置读写截止时间
    port.SetDeadline(time.Now().Add(5 * time.Second))

    // 写数据
    _, err = port.Write([]byte("hello serial"))
    if err != nil {
        log.Fatal("Write error:", err)
    }

    // 读数据
    buf := make([]byte, 128)
    n, err := port.Read(buf)
    if err != nil {
        log.Fatal("Read error:", err)
    }
    fmt.Printf("Received: %s\n", buf[:n])

    // 保持主线程
    select {}
}