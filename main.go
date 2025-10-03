package main

import (
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to estableshed connection: %v", err)
	}
	defer conn.Close()
	readAndWrite(conn)

}

func readAndWrite(conn net.Conn) {
	_, err := conn.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatalf("failed to write messege: %v", err)
	}
	conn.(*net.TCPConn).CloseWrite()

	buf, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalf("failed to read conn: %v", err)
	}

	log.Println(string(buf))
}
