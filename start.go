package main

import "fmt"

func main() {
	randomSeed()
	for i := 0; i < 4; i++ {
		w := NewWorld()
		fmt.Println("")
		fmt.Println(w.toString())

	}
}
