package main

import (
	"fmt"
)

func main() {
	fooclose(420)
	fooclose(420)
	fooclose(420)

	fooclose(420)

	fooclose(420)
	fooclose(420)
	fooclose(420)
	fooclose(4203442)

}

func fooclose(n int) int {
	t := n * 34
	fmt.Println(t)
	fmt.Println("The var t only exists inside fooclose")
	return t
}
