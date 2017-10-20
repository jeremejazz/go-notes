package main

import "fmt"

func g() {
	/*
		Even when not calling g() you cant
		use invalid statements or even make unused variables without getting
		an error on build
	*/
	b := 1
	fmt.Println(b)
}
func main() {
	g()
	fmt.Println("vim-go")
}
