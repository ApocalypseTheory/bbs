package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := "23"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on port " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to the Apocalypse!\r\n"))
	conn.Write([]byte("Type :h for help or :q to quit\r\n"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == ":q" {
			break
		}
		conn.Write([]byte("You typed: " + text + "\r\n"))
	}
}
