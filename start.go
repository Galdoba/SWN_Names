package main

import (
	"fmt"
)

func main() {
	randomSeed()
	for i := 0; i < 4; i++ {
		w := NewWorld()

		fmt.Println(w.toString())

	}

	// file, err := os.Open("tag.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// file2, err := os.Create("result.txt")
	// if err != nil {
	// 	log.Fatal("Cannot create file", err)
	// }
	// defer file2.Close()
	// scanner := bufio.NewScanner(file)
	// fmt.Fprintf(file2, "Hello Readers of golangcode.com")
	// allTags := tagList()
	// currentTag := ""
	// startFpart := false
	// friendText := ""
	// for scanner.Scan() { // internally, it advances token based on sperator

	// 	//fmt.Println(scanner.Text())  // token in unicode-char
	// 	//fmt.Println(scanner.Bytes()) // token in bytes
	// 	//currentTag = ""

	// 	fmt.Println("check tag:", currentTag)
	// 	//fmt.Println("check LINE:", scanner.Text())
	// 	if strings.Contains(scanner.Text(), "---F ") {
	// 		startFpart = true
	// 	}
	// 	if strings.Contains(scanner.Text(), "---E ") {
	// 		startFpart = false
	// 	}
	// 	if startFpart {
	// 		friendText = friendText + scanner.Text()
	// 	}
	// 	for i := range allTags {

	// 		if scanner.Text() == "---"+allTags[i] {
	// 			fmt.Println("found tag", allTags[i], "i =", i)
	// 			currentTag = allTags[i]
	// 		}

	// 	}

	// }
	// currentTag = ""
}
