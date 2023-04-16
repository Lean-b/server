package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)


func files(){
	
}


func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)

	for {
		//Read msg
		mensseger, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read: %s", err)
			break
		}
		fmt.Println("Hello", mensseger)

		//Wiriten
		response := "Hello from server!\n"
		_, err = c.Write([]byte(response))
		if err != nil {
			log.Printf("Failed to send response to client: %s", err)
			break
		}
	}
}

func main() {

	l, err := net.Listen("tcp4", ":5480")

	if err != nil {
		log.Fatalf("%\n", err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatalf("%\n", err)
		}
		log.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
