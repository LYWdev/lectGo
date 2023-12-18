package main

import (
    "fmt"
    "net"
)

func main() {
    // TCP 서버 소켓 생성
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()

    fmt.Println("Server started. Waiting for clients...")

    for {
        // 클라이언트 연결 대기
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err.Error())
            return
        }

        // 클라이언트와의 통신을 위한 고루틴 실행
        go handleRequest(conn)
    }
}

// 클라이언트 요청 처리 함수
func handleRequest(conn net.Conn) {
    // 클라이언트 주소 출력
    fmt.Println("Client connected:", conn.RemoteAddr().String())

    // 클라이언트로부터 데이터 수신
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    // 수신한 데이터 출력
    fmt.Println("Received message:", string(buffer))

    // 클라이언트에게 응답 전송
    conn.Write([]byte("Message received."))
    conn.Close()
}
