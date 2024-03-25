package main

import (
	"encoding/binary"
	"fmt"
	"net"
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
	var size uint32

	err := binary.Read(ClientConnection, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	BytMessage := make([]byte, size)
	ClientConnection.Read(BytMessage)

	RealMessage := string(BytMessage)
	fmt.Printf("Message From Client: %s\n", RealMessage)

}
