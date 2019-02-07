package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	seed := randomSeed()
	save := 0
	//makeTagsFunction()
	fmt.Println("Planet Sugested Name:", GiveName(randomSite(), "Place", -1))
	fmt.Println("")
	w := World{}
	for save != 1 {
		seed := randomSeed()
		fmt.Println("Seed =", seed)
		w = *NewWorld()
		fmt.Println(w.toString())
		fmt.Println(TradeAntagonists(w.TradeTag))
		fmt.Println(TradeAuthorities(w.TradeTag))
		leader := NewLeader()
		fmt.Println(leader.toString())
		save, _ = TakeOptions("Save Planet?", "Yes", "No")
	}
	fmt.Println("Name Planet:")
	planetName := InputString()
	fileOut, err := os.Create(planetName + ".txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer fileOut.Close()
	fmt.Fprintf(fileOut, "Seed: "+strconv.Itoa(int(seed))+"\n")
	fmt.Fprintf(fileOut, "Planet: "+planetName+"\n")
	fmt.Fprintf(fileOut, w.toString())
	fmt.Fprintf(fileOut, "Sample Story:\n")
	r1 := roll1dX(100, 0)
	fmt.Fprintf(fileOut, Story(r1, w.WorldTag1, w.WorldTag2))
	fmt.Fprintf(fileOut, " \n")
	fmt.Fprintf(fileOut, " \n")
	r2 := roll1dX(100, 0)
	fmt.Fprintf(fileOut, Story(r2, w.WorldTag1, w.WorldTag2))
	fmt.Fprintf(fileOut, " \n")
	fmt.Fprintf(fileOut, " \n")
	r3 := roll1dX(100, 0)
	fmt.Fprintf(fileOut, Story(r3, w.WorldTag1, w.WorldTag2))
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
// allTags := tagList()
// currentTag := ""
// fileLine := 0
// fmt.Fprintf(file2, "PlaceTagMap := make(map[string][]string)")
// for scanner.Scan() { // internally, it advances token based on sperator
// 	for i := range allTags {
// 		if scanner.Text() == allTags[i] {
// 			//fmt.Println("found tag", allTags[i], "i =", i)
// 			currentTag = allTags[i]
// 		}
// 	}

// 	//fmt.Fprintf(file2, strconv.Itoa(fileLine)+" Tag:"+currentTag+" ### "+scanner.Text()+"\n")
// 	//fmt.Println("check LINE:", scanner.Text())
// 	// if fileLine == 2 {
// 	// 	currentTag = strings.Replace(currentTag, " ", "", -1)
// 	// 	currentTag = strings.Replace(currentTag, "/", "", -1)
// 	// 	currentTag = strings.Replace(currentTag, "-", "", -1)
// 	// 	str := "\n"
// 	// 	tags := strings.Split(scanner.Text(), ", ")
// 	// 	str = str + currentTag + "slF := []string{}\n"
// 	// 	for i := range tags {
// 	// 		str = str + currentTag + "slF = append(" + currentTag + "slF, '" + tags[i] + "')\n"
// 	// 	}

// 	// 	str = str + "FriendTagMap['" + currentTag + "'] = " + currentTag + "slF\n"

// 	// 	fmt.Fprintf(file2, str)
// 	// }

// 	if fileLine == 6 {
// 		currentTag = strings.Replace(currentTag, " ", "", -1)
// 		currentTag = strings.Replace(currentTag, "/", "", -1)
// 		currentTag = strings.Replace(currentTag, "-", "", -1)
// 		str := "\n"
// 		tags := strings.Split(scanner.Text(), ", ")
// 		str = str + currentTag + "slP := []string{}\n"
// 		for i := range tags {
// 			str = str + currentTag + "slP = append(" + currentTag + "slP, '" + tags[i] + "')\n"
// 		}

// 		str = str + "PlaceTagMap['" + currentTag + "'] = " + currentTag + "slP\n"

// 		fmt.Fprintf(file2, str)
// 	}

// 	fileLine++
// 	if fileLine > 7 {
// 		fileLine = 0
// 	}
// }
// fmt.Fprintf(file2, "return PlaceTagMap")
// // currentTag = ""
// //fmt.Println(oneRollContact())

/*





 */

func makeTagsFunction() {
	file, err := os.Open("tag.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file2, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file)
	allTags := tradeTags()
	currentTag := ""
	fileLine := 0
	fmt.Fprintf(file2, "tRegulationsTagMap := make(map[string][]string)")
	for scanner.Scan() { // internally, it advances token based on sperator
		for i := range allTags {
			if scanner.Text() == allTags[i] {
				//fmt.Println("found tag", allTags[i], "i =", i)
				currentTag = allTags[i]
			}
		}

		//fmt.Fprintf(file2, strconv.Itoa(fileLine)+" Tag:"+currentTag+" ### "+scanner.Text()+"\n")
		//fmt.Println("check LINE:", scanner.Text())
		// if fileLine == 2 {
		// 	currentTag = strings.Replace(currentTag, " ", "", -1)
		// 	currentTag = strings.Replace(currentTag, "/", "", -1)
		// 	currentTag = strings.Replace(currentTag, "-", "", -1)
		// 	str := "\n"
		// 	tags := strings.Split(scanner.Text(), ", ")
		// 	str = str + currentTag + "slF := []string{}\n"
		// 	for i := range tags {
		// 		str = str + currentTag + "slF = append(" + currentTag + "slF, '" + tags[i] + "')\n"
		// 	}

		// 	str = str + "FriendTagMap['" + currentTag + "'] = " + currentTag + "slF\n"

		// 	fmt.Fprintf(file2, str)
		// }

		if fileLine == 6 {
			currentTag = strings.Replace(currentTag, " ", "", -1)
			currentTag = strings.Replace(currentTag, "/", "", -1)
			currentTag = strings.Replace(currentTag, "-", "", -1)
			str := "\n"
			tags := strings.Split(scanner.Text(), ", ")
			str = str + currentTag + "tRegulations := []string{}\n"
			for i := range tags {
				str = str + currentTag + "tRegulations = append(" + currentTag + "tRegulations, '" + tags[i] + "')\n"
			}

			str = str + "tRegulationsTagMap['" + currentTag + "'] = " + currentTag + "tRegulations\n"

			fmt.Fprintf(file2, str)
		}

		fileLine++
		if fileLine > 7 {
			fileLine = 0
		}
	}
	fmt.Fprintf(file2, "return tRegulationsTagMap")
}
