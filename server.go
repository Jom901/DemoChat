package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Listen on port 8080
	port, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer port.Close()
	connectionMap := make(map[net.Conn]string)
	for {
		// Accept a connection
		conn, err := port.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn, &connectionMap)
	}
}

func handleConnection(conn net.Conn, connectionMap *map[net.Conn]string) {
	defer conn.Close()

	// Read data from the connection
	buf := make([]byte, 1024)
	conn.Write([]byte("Enter username: \n"))
	var userName string
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if userName == "" {
				sendMessageToAllOtherConnections(conn, connectionMap, fmt.Sprintf("%s has disconnected\n", userName))
				delete((*connectionMap), conn)
			}
			return
		}
		// Print the received data
		message := string(buf[:n])
		message = strings.ReplaceAll(message, "\r\n", "")
		// If connection has no username, latest message IS the username
		if userName == "" {
			userName = message
			_, err = conn.Write([]byte(fmt.Sprintf("Welcome to the chat, %s !\n", userName)))
			(*connectionMap)[conn] = userName
		} else {
			fmt.Printf("%s: %s\n", userName, message)
			sendMessageToAllOtherConnections(conn, connectionMap, fmt.Sprintf("%s: %s\n", userName, message))
		}
	}
}

func sendMessageToAllOtherConnections(conn net.Conn, connMap *map[net.Conn]string, message string) {
	for otherConn, _ := range *connMap {
		if otherConn != conn {
			// Forward message to other clients
			_, err := otherConn.Write([]byte(message))
			if err != nil {
				fmt.Println("Ignoring exception to inform disconnect")
			}
		}
	}
}
