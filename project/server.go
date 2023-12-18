package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var clients []net.Conn

func main() {
	// 서버 주소와 포트 설정
	address := "192.168.64.16"
	port := "80"

	// 웹 서버 시작
	go startWebServer(address, port)

	// 클라이언트와의 채팅을 위한 TCP 소켓 서버 시작
	startChatServer(address, "8080")
}

// 웹 서버 시작 함수
func startWebServer(address, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Go server")
	})
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Gogo server")
	})

	err := http.ListenAndServe(address+":"+port, nil)
	if err != nil {
		log.Fatal("Web server error: ", err)
	}
}

// 채팅 서버 시작 함수
func startChatServer(address, port string) {
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatal("Chat server error: ", err)
	}
	defer listener.Close()

	log.Println("Chat server started. Waiting for clients...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		clients = append(clients, conn)
		go handleChatConnection(conn)
	}
}

// 클라이언트와의 채팅 처리 함수
func handleChatConnection(conn net.Conn) {
	log.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from client:", err)
			break
		}

		message = strings.TrimSpace(message)

		if message == "/quit" {
			break
		}

		log.Println("Received message:", message)

		sendMessageToClients(message, conn)
	}

	conn.Close()
	log.Println("Client disconnected:", conn.RemoteAddr().String())

	removeClient(conn)
}

// 클라이언트에게 메시지 전송 함수
func sendMessageToClients(message string, sender net.Conn) {
	for _, client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message + "\n"))
			if err != nil {
				log.Println("Error writing to client:", err)
				continue
			}
		}
	}
}

// 클라이언트 제거 함수
func removeClient(client net.Conn) {
	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}