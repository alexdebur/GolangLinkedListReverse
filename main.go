package main

import (
	"./classes"
	"fmt"
)

func main() {
	fmt.Println("test")
	list := classes.NewList(5)
	fmt.Println(list.ToString())
	list.Reverse()
	fmt.Println(list.ToString())
	list.Reverse()
	fmt.Println(list.ToString())
}
