package main

import "fmt"

type foo struct{}

type bar interface {
	myfunc() interface{}
}

func (f foo) myfunc() string {
	return "foo"
}

func main() {
	var hoge interface{} = "string"
	fmt.Println(hoge)

	var fuga bar = foo{}
	fmt.Println(fuga)
}
