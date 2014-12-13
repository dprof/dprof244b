package main

import (
	"log"
	"os"
	"math/rand"
)




func main(){
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(log.Flags() + log.Lmicroseconds)
	log.SetOutput(f)
	

	for i := 0; i<20; i++ {
		
		log.Printf(" traceID:%d spanID:%d parentID:%d callerID:{dc:dc1;rack:22} calleeID:{dc:dc2:rack:rack} operationID:%d\n", rand.Intn(1<<32),rand.Intn(1<<32),rand.Intn(1<<32),rand.Intn(1<<32))
			

	}

	//Mar 12 12:27:00:000  traceID:<TID>  spanID:<TID>  parentID:<ID> callerID:<metadata> <dc, rack> calleeID:<timstamp metata, dc,rack>  OperationID

}

