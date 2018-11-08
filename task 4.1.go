package main

import (
	"fmt"
)

func Swap()func(){
	return func() {
			defer fmt.Println("World")
			fmt.Println("Hello")
		}
	}
func main() {
	a := Swap()
	a()
	for i := 1; i < 10; i++ {
			go a()
		}
	fmt.Scanln()
}

