package main

import "fmt"

func main() {

	seed := 0
	current := 1

	for i := 0; i < 10; i++ {
		fmt.Println(current)
		seed, current = current, seed+current

	}
}
