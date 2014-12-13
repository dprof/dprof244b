package main

import (
	"os"
	"io"
	"fmt"
	"time"
	"net"
)


//Send the log to the destination
//logTyp = 0 for Dynomite, 1 for Redis
func LogSender(logFileName string, logType int) {
	f, err := os.Open(logFileName)
	 if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	 }
	defer f.Close()
	if options.debug {
		fmt.Printf("Processing filename %s\n", logFileName)
	}
	buffer := make([]byte, 1)
	line := make([]byte, 256)
	var fileoffset int64 = 0
	n := 0
	for {
		if n, err = f.ReadAt(buffer,fileoffset); err != nil {
			if err == io.EOF {
				//Sleep some and retry reading - possisbly new log contents will come
				time.Sleep(3000 * time.Millisecond)
				fileoffset,_ = f.Seek(fileoffset,0)
				if options.debug {
					fmt.Printf("Waiting on new input at %s\n", logFileName)
				}
			} else {
				break
			}
		}
		if err != io.EOF{
			fileoffset += int64(n)
			//fmt.Printf("%s",string(buffer))
			if (buffer[0]) != '\n' {
				line = append(line,buffer[0]);	
			} else {
				//We have collected a line
				//if options.debug {
				//	fmt.Printf("Collected a line(%d):'%s'\n",logType,string(line))
				//}
				//Send over line to server
				conn, err := net.Dial("tcp", options.ipPort)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				n, e:= conn.Write(line)
				if e != nil {
					fmt.Println(e)
					os.Exit(1)
				}
				if options.debug {
					fmt.Printf("Collected %s and sent line (%d bytes) to %s\n\n", string(line), n, "127.0.0.1:4000")
				}
				conn.Close()
				line = make([]byte, 1024)
			}
		}

	}
	if err == io.EOF {
		err = nil
	}
	
}

