package main

import (
	"fmt"

	"github.com/Galdoba/utils"
)

type npc struct {
	name       string
	lvl        int
	class      string
	attribute  map[string]int
	atr_Str    int
	atr_Dex    int
	atr_Con    int
	atr_Int    int
	atr_Wis    int
	atr_Cha    int
	background string
	skill      map[string]int

	//Strength, Dexterity, Constitution, Intelligence, Wisdom, and Charisma
}

func validSkills() []string {
	vs := []string{
		"Administer",
		"Connect",
		"Exert",
		"Fix",
		"Heal",
		"Know",
		"Lead",
		"Notice",
		"Perform",
		"Pilot",
		"Program",
		"Punch",
		"Shoot",
		"Sneak",
		"Stab",
		"Survive",
		"Talk",
		"Trade",
		"Work",
		"Biopsionics",
		"Metapsionics",
		"Precognition",
		"Telekinesis",
		"Telepathy",
		"Teleportation",
	}
	return vs
}

func blankAtrSet() map[string]int {
	atrSet := make(map[string]int)
	atrSet["Str"] = 0
	atrSet["Dex"] = 0
	atrSet["Con"] = 0
	atrSet["Int"] = 0
	atrSet["Wis"] = 0
	atrSet["Cha"] = 0
	atrSet["lvl"] = 0
	return atrSet
}

func blankSkillSet() map[string]int {
	skills := make(map[string]int)
	for i := range validSkills() {
		skills[validSkills()[i]] = -1
	}
	return skills
}

func CreateNPC() *npc {
	npc := npc{}
	inputName := utils.InputString("Set Name:('-m' for Male Random, '-f' for Female Random )")
	if inputName == "-m" {
		inputName = RandomName(false)
	}
	if inputName == "-f" {
		inputName = RandomName(true)
	}
	npc.name = inputName
	npc.attribute = blankAtrSet()
	npc.atr_Str = utils.RollDice("3d6")
	npc.atr_Dex = utils.RollDice("3d6")
	npc.atr_Con = utils.RollDice("3d6")
	npc.atr_Int = utils.RollDice("3d6")
	npc.atr_Wis = utils.RollDice("3d6")
	npc.atr_Cha = utils.RollDice("3d6")
	npc.skill = blankSkillSet()

	fmt.Println(npc)
	return &npc
}
