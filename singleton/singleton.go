package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ready")

	for i := 0; i < 30; i++ {
		getInstance()
	}
	fmt.Scanln()
}
