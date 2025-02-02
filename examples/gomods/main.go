package main

import "fmt"

//go:generate go run github.com/weareyolo/enumer -type=Pill -json
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	fmt.Println(PillStrings())
	fmt.Println(Placebo.IsAPill())
	fmt.Println(Placebo)
}
