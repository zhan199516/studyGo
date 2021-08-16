package main

import (
	//"fmt"
	"github.com/davecgh/go-spew/spew"
	//"log"
)

func main() {
	i := 123456
	s := "wo is String"
	//fmt.Println(i, s)
	//log.Println(i, s)
	spew.Dump(i, s)
}
