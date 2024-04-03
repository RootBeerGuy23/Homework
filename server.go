package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		ClientConnection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleServerConnection(ClientConnection)
	}
}

func HandleServerConnection(ClientConnection net.Conn) {

	deadline := time.Now().Add(10 * time.Second) // Timeout set to 10 seconds
	err := ClientConnection.SetReadDeadline(deadline)
	if err != nil {
		panic(err)
	}

	var size uint32
	err = binary.Read(ClientConnection, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	BytMessage := make([]byte, size)
	ClientConnection.Read(BytMessage)

	RealMessage := string(BytMessage)
	fmt.Printf("Message From Client: %s\n", RealMessage)

}
