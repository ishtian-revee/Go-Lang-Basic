package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main(){

	conn, err := net.Dial("tcp", "localhost:9000")

	if err != nil {

		fmt.Println("panic here!!", err)
		panic(err)
	}

	// closes the connection
	defer conn.Close()

	// returns a slice of bytes
	bs, _ := ioutil.ReadAll(conn)
	// converting that slice of bytes to string and printing it
	fmt.Println(string(bs))
}