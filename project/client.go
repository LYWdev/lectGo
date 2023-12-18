package main

import (
    "fmt"
    "net"
)

func main() {
    // 서버에 연결
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        return
    }
    defer conn.Close()

    // 서버로 메시지 전송
    message := "Hello, server!"
    conn.Write([]byte(message))

    // 서버로부터 응답 수신
    buffer := make([]byte, 1024)
    _, err = conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
        return
    }

    // 서버 응답 출력
    fmt.Println("Server response:", string(buffer))
}
