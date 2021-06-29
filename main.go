package main

import (
	"fmt"
)

var i float32 = 42
var (
	actorName string = "Tim Bergling"
)

func main() {
	var s string = "this is a string"
	var z []byte = []byte(s)
	fmt.Println(z)
}