package main

import (
	"github.com/goDB/single-server/data"
	"fmt"
)

func main(){
	s := new(data.Store)
	s.M = make(map[string]int)
	s.M["it"]=1
	fmt.Printf(s.M)
}
