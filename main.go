package main

import (
	"fmt"
	"strconv"
)

var i float32 = 42
var (
	actorName string = "Tim Bergling"
)

func main() {
	i := 45
	fmt.Printf("%v, %T\n", i, i)
	k := strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

}