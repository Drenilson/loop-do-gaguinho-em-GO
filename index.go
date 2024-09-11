// loop do Ceolinha e do Guaguinho em GO

package main

import (
	"fmt"
)

// Gaguinho
func Gaguinho() {
	fala := []string{"É isso ai", "Pe...", "Pe...", "Pe...", "Pessoal!"}

	for i := 0; i < len(fala); i++ {
		fmt.Println(fala[i])
	}
}

// Cebolinha (Turma da Mônica)
func Cebolinha() {
	fala := "Chola mais."

	for _, letra := range fala {
		fmt.Printf("%c", letra)
	}
	fmt.Println("------")
}

func main() {
	Cebolinha()
	Gaguinho()
}