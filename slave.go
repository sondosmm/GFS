package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	//Listen for incoming TCP connection from client
	listener, err := net.Listen("tcp", ":9060")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// Accept incoming connection from client
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Receive the string from the remote device
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert bytes to string and print the received message
	receivedMessage := string(buffer[:n])
	fmt.Println("Received message:", receivedMessage)

	if receivedMessage != "" {
		// Connect to the remote device
		conn, err := net.Dial("tcp", "192.168.1.184:9060")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		// Read the file contents into a byte slice
		fileContents, err := ioutil.ReadFile("D:/Manar.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Convert the byte slice to a string
		fileText := string(fileContents)

		// Create a map and put the values of the text on it
		letterCount := make(map[string]int)
		for _, char := range fileText {
			if char != ' ' && char != ',' && char != '!' {
				letterCount[string(char)]++
			}
		}
		fmt.Println(letterCount)

		// Convert the map to JSON
		jsonData, err := json.Marshal(letterCount)
		if err != nil {
			panic(err)
		}

		// Send the JSON data over the TCP connection
		_, err = conn.Write(jsonData)
		if err != nil {
			panic(err)
		}

		fmt.Println("Map sent successfully")
	} else {
		fmt.Println("Wrong handshake")
	}
}