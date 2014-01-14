/**
Code reference: http://jan.newmarch.name/go/rpc/chapter-rpc.html
*/

package main

import (
	"net/rpc"
	"fmt"
	"log"
	"os"
	"github.com/goDB/single-server/data"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":4321")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	// Synchronous call
	
	// Put value on server
	fmt.Println("Putting key-value pair, ('key',1) on server")
	keyValue := data.KeyValue{data.Key{"key"},data.Val{1}}
	var reply int
	err = client.Call("Server.Put", keyValue, &reply)
	if err != nil {
		log.Fatal("server error:", err)
	}
	
	// Get value from server
	var value data.Val
	var key = data.Key{"key"}
	err = client.Call("Server.Get", key, &value)
	if err != nil {
		log.Fatal("server error:", err)
	}
	fmt.Println("value for 'key' is",value.Value)
}
