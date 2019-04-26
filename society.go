package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

type Society struct {
	reasonForColonization    string
	initialCulture           string
	initialGoverment         string
	traits                   string
	mainConflictBeforeScream string
	govermentAfterScream     string
	secondConflict           string
	currentGoverment         string
}

func CreateSociety() {
	soc := Society{}
	soc.reasonForColonization = reasonForColonization()
	fmt.Println(soc.reasonForColonization)
	soc.initialGoverment = initialGoverment()
	fmt.Println(soc.initialGoverment)
	desidionPool, _ := generateDesidionPool(7, "Military Ships", "Stations", "Civilian Ships")
	fmt.Println(desidionPoolStr("Fleet Policy", desidionPool))

}

func generateDesidionPool(dPoints int, options ...string) ([]string, error) {
	var desidionPool []string
	if dPoints < len(options) {
		return desidionPool, errors.New("Critical: Not enough Desidion Points")
	}
	for i := range options {

		desidionPool = append(desidionPool, options[i])
		dPoints--
	}
	for dPoints > 0 {

		desidionPool = append(desidionPool, utils.RandomFromList(options))
		dPoints--
	}
	return desidionPool, nil
}

func desidionPoolStr(policyName string, dPool []string) string {
	dMap := make(map[string]int)
	for i := range dPool {
		dMap[dPool[i]] = dMap[dPool[i]] + 1
	}
	var output []string
	for key, val := range dMap {
		line := "{" + key + ": " + strconv.Itoa(val) + "}"
		output = append(output, line)
	}
	sort.Sort(Alphabetic(output))
	return policyName + "  " + strings.Join(output, " ")
}

func randomPolicy() (int, int, int) {
	p1 := 1
	p2 := 1
	p3 := 1
	for i := 0; i < 4; i++ {
		pick := utils.RollDice("d3")
		switch pick {
		case 1:
			if p1 > 4 {
				i--
				p1--
			}
			p1++
		case 2:
			if p2 > 4 {
				i--
				p2--
			}
			p2++
		case 3:
			if p3 > 4 {
				i--
				p3--
			}
			p3++
		}
	}
	return p1, p2, p3
}

func reasonForColonization() string {
	r := utils.RollDice("d20")
	reazons := []string{
		"Castaways",
		"Corporate factory world",
		"Ethnic or national purity",
		"Excavation Site",
		"Exiles",
		"Exotic genotypes",
		"Homeworld overpopulation",
		"Invasion forces",
		"Liason outpost",
		"Mandate malcontents",
		"Military outpost",
		"Political liberty",
		"Precious exports",
		"Prison planet",
		"Refueling outpost",
		"Religious liberty",
		"Research outpost",
		"Rich natural resourses",
		"Social liberty",
		"Trade hub",
	}
	return reazons[r-1]
}

func initialGoverment() string {
	r := utils.RollDice("d12")
	goverment := []string{
		"Autocracy",
		"Corporatism",
		"Democracy",
		"Feudalism",
		"Hydraulic Despotism",
		"Military Dictatorship",
		"Monarchy",
		"Oligarchy",
		"Republic",
		"Technocraticy",
		"Theocracy",
		"Tribalism",
	}
	return goverment[r-1]
}

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}
