package main

import (
	"fmt"

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
