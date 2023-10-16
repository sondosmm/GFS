package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	// Connect to the master
	conn, err := net.Dial("tcp", "192.168.1.184:9050")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send string to the master
	message := "Hi Master"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected Successfully to Master!")

	//---------------------------

	// Read the JSON data from the TCP connection
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	// Decode the JSON data into a map
	var data map[string]int
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("data received successfully")
	fmt.Println(data) // Prints the received map data

}