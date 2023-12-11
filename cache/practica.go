package main

import "fmt"

// type struct
type EmailSender struct {
}

// type func
type EmailChannel func(port, serial, encode string) (EmailSender, int, error)

func calcularSuma(n1, n2 int) int {
	fmt.Println(n1)
	fmt.Println(n2)
	return n1 + n2
}

func mainx() {

	//channel
	//	c := make(chan string)
	// map
	//	ma := make(map[string]bool)

	go calcularSuma(2, 2)

}
