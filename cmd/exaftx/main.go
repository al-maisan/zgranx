package main

import "fmt"

var bts, rev string

func main() {
	version := fmt.Sprintf("%s::%s", bts, rev)
	fmt.Printf("exaftx: rev: %s\n", version)
}
