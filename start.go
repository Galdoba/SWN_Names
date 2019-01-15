package main

import "fmt"

func main() {
	randomSeed()
	fmt.Println(GiveName("spanish", "Sur", 1))
	fmt.Println(NewFullName(false, "", "", ""))
}
