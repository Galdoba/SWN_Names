package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/Galdoba/utils"
)

var expertize int

type planet struct {
	name      string
	friction  int
	localSpec map[string]int
}

func NewPlanet(name string) *planet {
	pl := &planet{}
	pl.name = name
	pl.friction = friction(pl.name)
	pl.localSpec = tradeSpec(pl.name)
	return pl
}

func main() {
	expertize = utils.InputInt("Set Trader's Experteese:")
	fmt.Println("Expertize =", expertize)
	var planets []string
	var saleLots []tradePosition
	planetNames := utils.FileNames("./", ".txt")
	for i := range planetNames {
		planets = append(planets, strings.TrimSuffix(planetNames[i], ".txt"))
	}
	_, currentPlanet := utils.TakeOptions("Choose Planet:", planets...)
	fmt.Println("Current planet:", currentPlanet)
	fmt.Println("Friction:", friction(currentPlanet))
	fmt.Println("Export:")
	planetLines := utils.LinesFromTXT(currentPlanet + ".txt")
	for j := range planetLines {
		if strings.Contains(planetLines[j], "Commoditie") {
			//fmt.Println(planetLines[j])
			saleLots = append(saleLots, *NewLot(planetLines[j], currentPlanet))

		}
	}
	var export []string
	for i := range saleLots {
		if saleLots[i].name != "UNKNOWN" {
			fmt.Println(saleLots[i].name, "Tags:", saleLots[i].tags, "Standard Price:", lotExportPrice(saleLots[i], currentPlanet), "(", saleLots[i].priceBase, ")")
			export = append(export, saleLots[i].name)
		}
	}
	checkLot := ""
	if len(export) > 0 {
		_, checkLot = utils.TakeOptions("Survey Commodity:", export...)
	} else {
		panic("End Program")
	}
	lot := pickLot(checkLot, saleLots)
	for i := range planets {
		fmt.Println(planets[i], "importing for ", lotImportPrice(lot, planets[i]))
	}

}

func CommoditieList(allLines []string) []string {
	var resultLines []string
	for i := range allLines {
		if strings.Contains(allLines[i], "Commoditie") {
			resultLines = append(resultLines, allLines[i])
		}
	}
	return resultLines
}

type tradePosition struct {
	origin    string
	name      string
	priceBase int
	priceMod  int
	tags      []string
}

func NewLot(commoditie string, planet string) *tradePosition {
	lot := tradePosition{}
	lot.name, lot.tags, lot.priceMod, lot.priceBase = parseLot(commoditie)
	lot.origin = planet
	//lot.costMod = friction(lot.origin)
	//lot.priceReal = int(float64(lot.priceBase) * salesChart(10+lot.costMod))
	return &lot
}

func lotExportPrice(lot tradePosition, planet string) int {
	planetData := NewPlanet(planet)
	dealMod := 0
	for i := range lot.tags {
		for k, v := range planetData.localSpec {
			if lot.tags[i] == k {
				dealMod = dealMod + v
			}
		}
	}
	tradeChart := salesChart(10 + planetData.friction + dealMod - expertize)
	price := int(float64(lot.priceBase) * tradeChart)
	return price
}

func lotImportPrice(lot tradePosition, planet string) int {
	planetData := NewPlanet(planet)
	dealMod := 0
	for i := range lot.tags {
		for k, v := range planetData.localSpec {
			if lot.tags[i] == k {
				dealMod = dealMod + v
			}
		}
	}
	tradeChart := salesChart(11 - planetData.friction + dealMod + expertize)
	price := int(float64(lot.priceBase) * tradeChart)
	return price
}

func parseLot(str string) (string, []string, int, int) {
	//"Commoditie 1:{UNKNOWN} {Livestock, Survival} {priceMod:0} {pricePerUnit:5000}"
	btStr := []byte(str)
	var args []string
	a := 0
	read := false
	var btWord []byte
	for i := range btStr {
		if string(btStr[i]) == "{" {
			read = true
			args = append(args, "")
			continue
		}
		if string(btStr[i]) == "}" {
			read = false
			args[a] = string(btWord[:])
			btWord = nil
			a++
			continue
		}
		if read {
			btWord = append(btWord, btStr[i])
			//args[a] += string(btStr[i])
		}
	}
	name := args[0]
	tags := strings.Split(args[1], ", ")
	pModS := strings.TrimPrefix(args[2], "priceMod:")
	priceMod, _ := strconv.Atoi(pModS)
	pUnitS := strings.TrimPrefix(args[3], "pricePerUnit:")
	priceUnit, _ := strconv.Atoi(pUnitS)
	return name, tags, priceMod, priceUnit
}

func parseSpec(str string) map[string]int {
	//World trade specific: {Religious -2} {Consumer -1} {Sapient 1} {Survival 2}
	specMap := make(map[string]int)
	btStr := []byte(str)
	var args []string
	a := 0
	read := false
	for i := range btStr {
		if string(btStr[i]) == "{" {
			read = true
			args = append(args, "")
			continue
		}
		if string(btStr[i]) == "}" {
			read = false
			a++
			continue
		}
		if read {
			args[a] += string(btStr[i])
		}
	}

	for i := range args {
		parts := strings.Split(args[i], ": ")
		tag := parts[0]
		mod, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		specMap[tag] = mod
	}

	return specMap
}

func salesChart(roll int) float64 {
	if roll < 2 {
		roll = 2
	}
	if roll > 19 {
		roll = 19
	}
	switch roll {
	case 2:
		return 0.1
	case 3:
		return 0.3
	case 4:
		return 0.4
	case 5:
		return 0.5
	case 6:
		return 0.6
	case 7:
		return 0.7
	case 8:
		return 0.8
	case 9:
		return 0.9
	case 10:
		return 1.0
	case 11:
		return 1.0
	case 12:
		return 1.1
	case 13:
		return 1.2
	case 14:
		return 1.4
	case 15:
		return 1.6
	case 16:
		return 1.8
	case 17:
		return 2.0
	case 18:
		return 2.5
	case 19:
		return 3.0
	}
	return 0.0
}

func friction(planet string) int {
	planetLines := utils.LinesFromTXT(planet + ".txt")
	friction := 0
	for i := range planetLines {
		if strings.Contains(planetLines[i], "Law Level: ") {
			lawStr := strings.TrimPrefix(planetLines[i], "Law Level: ")
			law, err := strconv.Atoi(lawStr)
			if err != nil {
				panic(err)
			}
			friction = lawToFriction(law)
		}
	}
	return friction
}

func tradeSpec(planet string) (specMap map[string]int) {
	planetLines := utils.LinesFromTXT(planet + ".txt")
	for i := range planetLines {
		if strings.Contains(planetLines[i], "World trade specific: ") {
			specMap = parseSpec(planetLines[i])
		}
	}
	return specMap
}

func lawToFriction(law int) int {
	if law < 7 {
		addon := 7 - law
		return 2 + addon
	}
	if law > 7 {
		return law - 5
	}
	return 2
}

func pickLot(name string, allLots []tradePosition) tradePosition {
	for i := range allLots {
		if allLots[i].name == name {
			return allLots[i]
		}
	}
	fmt.Println("Position was not founded (ERROR)")
	return tradePosition{}
}

// if strings.Contains(scanner.Text(), "Commoditie ") {
// 	fmt.Println(scanner.Text())
// }
