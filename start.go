package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

var seed int64

func main2() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	var txts []string
	for _, f := range files {
		if strings.Contains(f.Name(), ".txt") {
			txts = append(txts, f.Name())
		}
		fmt.Println(f.Name())
	}
	fmt.Println("Txt only:")
	for i := range txts {
		fmt.Println(txts[i])
		file, err := os.Open(txts[i])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "Commoditie ") {
				fmt.Println(scanner.Text())
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}

func CreatePlanet() {

	save := 0
	fmt.Println("Planet Sugested Name:", GiveName(randomSite(), "Place", -1))
	fmt.Println("")
	w := &World{}
	for save != 1 {
		//seed := randomSeed()
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
		fmt.Println("d4=", roll1dX(4, 0))
		fmt.Println("d6=", roll1dX(6, 0))
		fmt.Println("d8=", roll1dX(8, 0))
		fmt.Println("d10=", roll1dX(10, 0))
		fmt.Println("d12=", roll1dX(12, 0))
		fmt.Println("d20=", roll1dX(20, 0))
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

	fmt.Println(roll1dX(4, 0))
	fmt.Println(roll1dX(6, 0))
	fmt.Println(roll1dX(8, 0))
	fmt.Println(roll1dX(10, 0))
	fmt.Println(roll1dX(12, 0))
	fmt.Println(roll1dX(20, 0))
}

func main() {
	seed = randomSeed()

	optInt, optStr := utils.TakeOptions("select action:", "Create New Planet", "Load Planet", "Create NPC", "Create War", "Society")
	fmt.Println("Selected:", optInt, optStr)
	switch optInt {
	case 1:
		CreatePlanet()
	case 2:
		ReadPlanet()
	case 3:
		npc := CreateNPC()
		fmt.Println(npc.report())
	case 4:
		CreateWar()
	case 5:
		CreateSociety()
	default:
		fmt.Println("Program Exit")
	}
	fmt.Println("Program Ended")
}

func CreateWar() {
	manpower := InputInt("Manpower: ")
	army1 := NewArmy(manpower, 4)
	fmt.Println(army1.leader.toString())
}

func ReadPlanet() {
	var planets []string
	planetNames := utils.FileNames("./", ".txt")
	for i := range planetNames {
		planets = append(planets, strings.TrimSuffix(planetNames[i], ".txt"))
	}
	_, currentPlanet := utils.TakeOptions("Choose Planet:", planets...)
	fmt.Println("Current planet:", currentPlanet)
	planetLines := utils.LinesFromTXT(currentPlanet + ".txt")

	tag1, tag2, _ := getPlanetTags(planetLines)
	fmt.Println("---------------------------")
	fmt.Println(Story(utils.RollDice("d100"), tag1, tag2))
	fmt.Println("---------------------------")
	agency := NewAgency("Name", 4, 2)
	agency.Update()
	fmt.Println(agency.Report())
	agency.SetElementLevel("Assassins", 3)
	agency.Update()
	fmt.Println(agency.Report())

}

func getPlanetTags(data []string) (string, string, string) {
	tag1 := "tag1"
	tag2 := "tag2"
	tagTrade := "trade tag"
	for i := range data {
		if strings.Contains(data[i], "WorldTag1: ") {
			tag1 = strings.TrimPrefix(data[i], "WorldTag1: ")
		}
		if strings.Contains(data[i], "WorldTag2: ") {
			tag2 = strings.TrimPrefix(data[i], "WorldTag2: ")
		}
		if strings.Contains(data[i], "TradeTag: ") {
			tagTrade = strings.TrimPrefix(data[i], "TradeTag: ")
		}
	}

	return tag1, tag2, tagTrade
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

//import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	utils "github.com/Galdoba/utils"
// )

// func main() {
// 	//var allGoods []string
// 	var allLots []TradeLot
// 	planetNames := utils.FileNames("./", ".txt")
// 	for i := range planetNames {
// 		planet := strings.TrimSuffix(planetNames[i], ".txt")
// 		lines := utils.LinesFromTXT(planetNames[i])
// 		for j := range lines {
// 			if strings.Contains(lines[j], "Commoditie") {
// 				allLots = append(allLots, *NewTradeLot(lines[j], planet))
// 			}
// 		}
// 		for i := range allLots {
// 			fmt.Println(allLots[i])
// 		}
// 		// 	allGoods = append(allGoods, CommoditieList(lines)...)
// 		// }
// 		// fmt.Println(utils.RollDice("2d6"))
// 		// for i := range allGoods {
// 		// 	fmt.Println(allGoods[i])
// 		// 	fmt.Println(parseLot(allGoods[i]))
// 		// 	allLots = append(allLots, *NewTradeLot(allGoods[i], "Planet"))
// 		// 	fmt.Println(*NewTradeLot(allGoods[i], "Planet"))
// 		// }
// 	}
// }

// func CommoditieList(allLines []string) []string {
// 	var resultLines []string
// 	for i := range allLines {
// 		if strings.Contains(allLines[i], "Commoditie") {
// 			resultLines = append(resultLines, allLines[i])
// 		}
// 	}
// 	return resultLines
// }

// type TradeLot struct {
// 	origin    string
// 	name      string
// 	priceBase int
// 	priceMod  int
// 	priceReal int
// 	tags      []string
// 	friction  int
// }

// func NewTradeLot(commoditie string, planet string) *TradeLot {
// 	lot := TradeLot{}
// 	lot.name, lot.tags, lot.priceMod, lot.priceBase = parseLot(commoditie)
// 	lot.origin = planet
// 	return &lot
// }

// func parseLot(str string) (string, []string, int, int) {
// 	//"Commoditie 1:{UNKNOWN} {Livestock, Survival} {priceMod:0} {pricePerUnit:5000}"
// 	btStr := []byte(str)
// 	var args []string
// 	a := 0
// 	read := false
// 	for i := range btStr {
// 		if string(btStr[i]) == "{" {
// 			read = true
// 			args = append(args, "")
// 			continue
// 		}
// 		if string(btStr[i]) == "}" {
// 			read = false
// 			a++
// 			continue
// 		}
// 		if read {
// 			args[a] += string(btStr[i])
// 		}
// 	}
// 	name := args[0]
// 	tags := strings.Split(args[1], ", ")
// 	pModS := strings.TrimPrefix(args[2], "priceMod:")
// 	priceMod, _ := strconv.Atoi(pModS)
// 	pUnitS := strings.TrimPrefix(args[3], "pricePerUnit:")
// 	priceUnit, _ := strconv.Atoi(pUnitS)
// 	return name, tags, priceMod, priceUnit
// }

// func salesChart(roll int) float64 {
// 	if roll < 2 {
// 		roll = 2
// 	}
// 	if roll > 19 {
// 		roll = 19
// 	}
// 	switch roll {
// 	case 2:
// 		return 0.1
// 	case 3:
// 		return 0.3
// 	case 4:
// 		return 0.4
// 	case 5:
// 		return 0.5
// 	case 6:
// 		return 0.6
// 	case 7:
// 		return 0.7
// 	case 8:
// 		return 0.8
// 	case 9:
// 		return 0.9
// 	case 10:
// 		return 1.0
// 	case 11:
// 		return 1.0
// 	case 12:
// 		return 1.1
// 	case 13:
// 		return 1.2
// 	case 14:
// 		return 1.4
// 	case 15:
// 		return 1.6
// 	case 16:
// 		return 1.8
// 	case 17:
// 		return 2.0
// 	case 18:
// 		return 2.5
// 	case 19:
// 		return 3.0
// 	}
// 	return 0.0
// }

// // if strings.Contains(scanner.Text(), "Commoditie ") {
// // 	fmt.Println(scanner.Text())
// // }
