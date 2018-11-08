package main

import (
	"fmt"
)

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
func volume (v Transport) float64 {
	return v.Weight/v.Density
}
func (v Transport) info() string {
	return v.Class+" "+v.Model
}
func main() {
	v := Transport{2567  , 38, "Business", "Rx 8"}
	fmt.Printf("volume = %f\n", volume(v))
	fmt.Printf("info - %s\n", v.info())
}