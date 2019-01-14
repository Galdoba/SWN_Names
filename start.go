package main

import "fmt"

func main() {
	function := map[string]func(int, int) int{
		"someFunction1": someFunction1,
		"someFunction2": someFunction2,
	}
	fmt.Println(someOtherFunction(3, 2, function["someFunction1"]))
	fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))
	fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))

	fmt.Println(placesRegistry("spanish", 2))
}

func someFunction1(a, b int) int {

	return a + b
}

func someFunction2(a, b int) int {
	return a - b
}

func someOtherFunction(a, b int, f func(int, int) int) int {
	return f(a, b)
}
