/**
Code reference: http://jan.newmarch.name/go/rpc/chapter-rpc.html
*/

package main

import (
	"fmt"
	"net/rpc"
	"errors"
	"net/http"
	"github.com/goDB/single-server/data"
)

type Server struct {

}

var dataStore data.Store

func initialize(){
	dataStore.M = make(map[string]int)
}

func (t *Server) Put(keyValue *data.KeyValue, reply *int) error {
	dataStore.M[keyValue.K.String]=keyValue.V.Value
	return nil
}

func (t *Server) Get(key *data.Key, value *data.Val) error {
	val,ok := dataStore.M[key.String]
	if !ok {
		return errors.New("key not present")
	}
	value.Value = val
	return nil
}

func main() {
	dataStore.M = make(map[string]int)
	dataServer := new(Server)
	rpc.Register(dataServer)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
