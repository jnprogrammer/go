package main

import (
	"fmt"
)

const (
	_ = iota // _ tells Go I don't care about what returns here.

	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}
