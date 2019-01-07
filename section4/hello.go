package main

import "fmt"

func main() {
	fmt.Println("Hello everybody, this is the most awesome class in the entire world for having and learning the Go programming language!")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")

}

func bar() {
	fmt.Println("and then we exited")
}

// control flow
// (1) sequence
// (2) loo; iterative
// (3) conditrional
