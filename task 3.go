package main

import "fmt"

type Infer interface {
	inf()string
}

type Transport struct {
	Weight  float64
	Density float64
	Class string
	Model string
}

type Car struct {
	Transport
	Color   string
	Name    string
	Mileage float64
	Age     int
	Length  float64
}

func (t Transport) string() string {
	return fmt.Sprintf("%v %v %v %v", t.Weight, t.Density, t.Class, t.Model)
}
func volume(t Infer) string {
return t.inf()
}
func (t Transport) inf() string{
	return t.Class+" "+ t.Model
}
func main() {
	var t Infer
	t = Transport{1234.0, 42.0,  "Business", "rx 8" }
	mass := []Transport{{2543, 12, "Business", "rx 7"},
		{2000, 12, "Business", "rx 8"},
		{3000, 14, "Business", "3"},
		{2145, 11, "Business", "6"},}
	fmt.Println(volume(t))
	for i := range mass{
		fmt.Printf ("%d | %s \n", i, mass[i].string())
	}
}

