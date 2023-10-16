package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	///////////////////////////client///////////////////////////////////
	//Listen for incoming TCP connection from client
	listener, err := net.Listen("tcp", ":9050")
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

	//////////////////////////chuncks/////////////////////////////////////
	//send to the chunks
	if receivedMessage != "" {
		///////////////////////////////database/////////////////////////////
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/project")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		// Check the connection
		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}

		// Create a users table
		_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS data (
            letter char PRIMARY KEY,
            count int NOT NULL
        );
    `)
		if err != nil {
			panic(err)
			fmt.Println("db error")
		}

		var chuncks []string

		chuncks = append(chuncks, "192.168.1.188:9060", "192.168.1.206:9070")

		for i := 0; i < len(chuncks); i++ {
			fmt.Println("ip : " + chuncks[i])
			// Connect to the first chunck device
			conn1, err := net.Dial("tcp", chuncks[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn1.Close()

			// Send the string to the remote device
			chunk1 := "send me the file"
			_, err = conn1.Write([]byte(chunk1))
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("String sent successfully to chunk 1")

			parts := strings.Split(chuncks[i], ":")

			listener, err := net.Listen("tcp", ":"+parts[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			defer listener.Close()

			// Accept incoming connection from client
			conn2, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn2.Close()

			// Read the JSON data from the TCP connection
			buf := make([]byte, 4096)
			n, err = conn2.Read(buf)
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

			// Insert the data into the database
			for letter, count := range data {
				_, err = db.Exec(`INSERT INTO data (letter, count) VALUES (?, ?);`, letter, count)
				if err != nil {
					panic(err)
				}
			}

		}

		// Select all the data from the data table
		rows, err := db.Query(`SELECT letter, SUM(count) FROM data GROUP BY letter`)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// Build a map from the query results
		dataMap := make(map[string]int)
		for rows.Next() {
			var letter string
			var count int
			err := rows.Scan(&letter, &count)
			if err != nil {
				panic(err)
			}
			dataMap[letter] = count
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}
		// Print the map
		fmt.Println(dataMap)

		// Convert the map to JSON
		jsonData, err := json.Marshal(dataMap)
		if err != nil {
			panic(err)
		}

		// Send the JSON data over the TCP connection
		_, err = conn.Write(jsonData)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("wrong handshake")
	}

}
