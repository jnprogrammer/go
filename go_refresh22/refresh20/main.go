package main

import (
	"fmt"
	"time"
)

func main() {

	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello Go for the 3220 time\nI KEEP FIGHTING TO CHANGE AND I KEEP FIGHTING")
	}()

	func() {
		fmt.Println("Cats are not food")
	}()
}
