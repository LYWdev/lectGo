package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 서버 주소와 포트 설정
	serverAddress := "192.168.64.16"
	serverPort := "8080"

	// 서버에 연결
	conn, err := net.Dial("tcp", serverAddress+":"+serverPort)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	log.Println("Connected to server. You can start chatting.")

	go receiveMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}

// 서버로부터 메시지 수신 함수
func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error receiving message:", err)
			break
		}

		message = message[:len(message)-1]
		fmt.Println("Received message:", message)
	}

	log.Println("Disconnected from server.")
}

