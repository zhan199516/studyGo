package main

import (
	//"fmt"
	"errors"
	"github.com/davecgh/go-spew/spew"
	//"log"
)

func main() {
	i := 18
	s := "wo is zhanshuaiqiang"
	m := map[int]string{1: "value", 2: "2"}
	e := errors.New("Error, system is not fund")
	p := person{Id: 26, Name: "zhan"}
	spew.Dump(i, s, m, e, p)
}

type person struct {
	Id   int
	Name string
}
