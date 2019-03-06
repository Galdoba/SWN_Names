package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//makeTagsFunction()
	action := 0
	action, _ = TakeOptions("Select action", "Create new Planet", "Select Existing Planet")
	if action == 1 {
		CreaNewPlanetFile()
	}
	if action == 2 {
		SelectPlanetFile()
	}
}

func SelectPlanetFile() {
	readCurrentDir()
}

func readCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	var textList []string

	for _, name := range list {
		if strings.Contains(name, ".txt") {
			plntName := strings.Trim(name, ".txt")
			textList = append(textList, plntName)
		}
	}
	_, selected := TakeOptions("which?", textList...)

	b, err := ioutil.ReadFile(selected + ".txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	fmt.Println(str)
	lines := strings.Split(str, "\n")
	fmt.Println("lines[5]")
	fmt.Println(lines[5])
	fmt.Println(lines[5])

}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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

func CreaNewPlanetFile() {
	seed := randomSeed()
	save := 0

	fmt.Println(RandomColonyReport())
	w := &World{}
	for save != 1 {
		seed := randomSeed()
		fmt.Println("Seed =", seed)
		w = NewWorld()
		fmt.Println(w.toString())
		fmt.Println("w.TotalPopulation()", w.TotalPopulation())
		fmt.Println(TradeAntagonists(w.TradeTag))
		fmt.Println(TradeAuthorities(w.TradeTag))
		leader := NewLeader()
		army := NewArmy(w.TotalPopulation()/100, w.TechLevelInt())
		fmt.Println(leader.toString())
		fmt.Println(army)
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
	fmt.Println("Planet Sugested Name:", GiveName(randomSite(), "Place", -1))
}
