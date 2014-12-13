package main

import (
	"fmt"
	"net"
	"os"
	"github.com/gocql"
)




const (
	COLLECTOR_CONN_PORT = "4000"
	COLLECTOR_CONN_TYPE = "tcp"
	LOG_BUFFER_SIZE = 16384
	DEBUG = true
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(COLLECTOR_CONN_TYPE, ":"+COLLECTOR_CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	

	//Initialize Connection
	session := DataStoreInit()

	fmt.Println("Listening on " + "Port" + COLLECTOR_CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go func(){
			handleInComingLogBatch(conn, session)
		}()
	}
	
	fmt.Println("Exiting Aggregator")
}


// Handles incoming requests - batch and write to Cassandra
func handleInComingLogBatch(conn net.Conn, session *gocql.Session) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, LOG_BUFFER_SIZE)

	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	} else {
		if false {
			fmt.Printf("Read (%d bytes) string(%s)\n",len, string(buf))
		}
		
		//Cleanup junk from transfer
		tmp := make([]byte,0)
		for _,v := range buf {
			if v>=32 && v<127 {
				tmp = append(tmp,v)
			}
		}

		DataStoreRunCommand(string(tmp),session)
	}
	conn.Close()
}


