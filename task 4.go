package main

import "fmt"

func Swap()func(){
	return func() {
		defer fmt.Println("World")
		fmt.Println("Hello")
	}
}
func main() {
a :=Swap()
a()
}
