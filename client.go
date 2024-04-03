package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send Message")
		fmt.Println("2. Exit")
		fmt.Print(">>")
		scanner.Scan()
		choose := scanner.Text()
		if choose == "1" {
			SendMessageMenu()
		} else if choose == "2" {
			fmt.Println("GoodBye!")
			break
		}

	}
}

func SendMessageMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var Message string

	for {
		fmt.Print("Input Message To Send: ")
		scanner.Scan()
		Message = scanner.Text()
		break
	}
	SendMessageToServer(Message)

}

func SendMessageToServer(Message string) {
	ServerConn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer ServerConn.Close()

	deadline := time.Now().Add(5 * time.Second) // Timeout set to 5 seconds
	err = ServerConn.SetDeadline(deadline)
	if err != nil {
		panic(err)
	}
	err = binary.Write(ServerConn, binary.LittleEndian, uint32(len(Message)))
	if err != nil {
		panic(err)
	}

	_, err = ServerConn.Write([]byte(Message))
	if err != nil {
		panic(err)
	}

}

func main() {
	menu()
}
