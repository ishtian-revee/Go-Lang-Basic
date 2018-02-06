package main

import(
	"fmt"
	"io"
	"net"
	"time"
)

func main(){

	lis, err := net.Listen("tcp", "0.0.0.0:9000")

	if err != nil {

		fmt.Println("panic here!!", err)
		panic(err)
	}

	// closing the listener when all the functions are done
	defer lis.Close()

	for {

		// Accept waits for and returns the next connection to the listener
		conn, err := lis.Accept()

		if err != nil {

			fmt.Println("panic here!!", err)
			panic(err)
		}

		// WriteString takes a writer which is the connection (con) and a string
		io.WriteString(conn, fmt.Sprint("Yaaoo Rev\n", time.Now(), "\n"))
		// closing the connection
		conn.Close()
	}
}