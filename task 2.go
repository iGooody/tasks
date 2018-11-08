package main

import "fmt"

type Transport struct {
	Weight  float64
	Density float64
	Class string
	Model string
}

type Car struct{
	Transport
	Color   string
	Name    string
	Mileage float64
	Age     int
	Length  float64
}
func main() {
	mass := []Transport{{2543, 12, "Business", "rx 7"},
		{2000, 12, "Business", "rx 8"},
		{3000, 14, "Business", "3"},
		{2145, 11, "Business", "6"},
	}
	for i := range mass {
		fmt.Printf("%d | ", i)
		fmt.Print(mass[i])
		fmt.Printf("\n")
	}
}