package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Galdoba/utils"
)

var atributeNames = [...]string{"Str", "Dex", "Con", "Int", "Wis", "Cha"}
var physicalAtr = [...]string{"Str", "Dex", "Con"}
var mentalAtr = [...]string{"Int", "Wis", "Cha"}

type npc struct {
	name       string
	lvl        int
	class      string
	attribute  map[string]int
	background string
	skill      map[string]int
	//Strength, Dexterity, Constitution, Intelligence, Wisdom, and Charisma
}

func (npc *npc) Atr(s string) int {
	if val, ok := npc.attribute[s]; ok {
		return val
	}
	return -999
}

func (npc *npc) IncreaseSkill(skill string) {
	npc.skill[skill] = npc.skill[skill] + 1
}

func (npc *npc) IncreaseStat(stat string) {
	npc.attribute[stat] = npc.attribute[stat] + 1
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

func nonPsyhicSkills() []string {
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
	}
	return vs
}

func psyhicSkills() []string {
	vs := []string{
		"Biopsionics",
		"Metapsionics",
		"Precognition",
		"Telekinesis",
		"Telepathy",
		"Teleportation",
	}
	return vs
}

func combatSkills() []string {
	vs := []string{
		"Punch",
		"Shoot",
		"Stab",
	}
	return vs
}

func nonCombatSkills() []string {
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
		"Sneak",
		"Survive",
		"Talk",
		"Trade",
		"Work",
	}
	return vs
}

func backgrounds() []string {
	bcgr := []string{
		"Barbarian",
		"Clergy",
		"Courtesan",
		"Criminal",
		"Dilettante",
		"Entertainer",
		"Merchant",
		"Noble",
		"Official",
		"Peasant",
		"Physician",
		"Pilot",
		"Politician",
		"Scholar",
		"Soldier",
		"Spacer",
		"Technician",
		"Thug",
		"Vagabond",
		"Worker",
	}
	return bcgr
}

func randomClass() string {
	classArray := []string{
		"Warrior",
		"Warrior",
		"Warrior",
		"Warrior",
		"Expert",
		"Expert",
		"Expert",
		"Expert",
		"Psyhic",
		"Psyhic",
		"Adventurer (W-E)",
		"Adventurer (W-E)",
		"Adventurer (W-P)",
		"Adventurer (E-P)",
	}
	return classArray[rand.Intn(len(classArray))]
}

func rollAtrSet() map[string]int {
	atrSet := make(map[string]int)
	var changeAtr []string
	for i := range atributeNames {
		atrSet[atributeNames[i]] = utils.RollDice("3d6")
		share := (14 - atrSet[atributeNames[i]])
		for sh := 0; sh < share; sh++ {
			changeAtr = append(changeAtr, atributeNames[i])
		}
	}
	if len(changeAtr) > 0 {
		rollRes := utils.RollDice("d" + strconv.Itoa(len(changeAtr)))
		pick := changeAtr[rollRes-1]
		atrSet[pick] = 14
	}
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
	inputName := utils.InputString("Set Name: \n('-m' for Male Random, '-f' for Female Random )\n")
	if inputName == "-m" {
		inputName = RandomName(false)
	}
	if inputName == "-f" {
		inputName = RandomName(true)
	}
	npc.name = inputName
	npc.attribute = rollAtrSet()
	npc.skill = blankSkillSet()
	npc.background = backgrounds()[rand.Intn(len(backgrounds()))]
	npc.backgroundGrowth()
	fmt.Println(npc.report())
	return &npc
}

func growthLearnPatern() []int {
	var patern []int
	for i := 0; i < 3; i++ {
		patern = append(patern, utils.RollDice("1d2")-1)
	}
	fmt.Println(patern)
	return patern
}

func (npc *npc) backgroundGrowth() {
	gPatern := growthLearnPatern()
	back := npc.background
	npc.develop(freeSkill(back))
	for i := range gPatern {
		action := "Growth"
		if gPatern[i] == 1 {
			action = "Learn"
		}
		key := back + action
		options := glMap()[key]
		res := options[rand.Intn(len(options))]
		npc.develop(res)

	}
}

func glMap() map[string][]string {
	glMap := make(map[string][]string)
	glMap["BarbarianGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["BarbarianLearn"] = []string{"AnyCombat", "Connect", "Exert", "Lead", "Notice", "Punch", "Sneak", "Survive"}
	glMap["ClergyGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["ClergyLearn"] = []string{"Administer", "Connect", "Know", "Lead", "Notice", "Perform", "Talk", "Talk"}
	glMap["CourtesanGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Physical2", "Connect", "AnySkill"}
	glMap["CourtesanLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Perform", "Survive", "Talk", "Trade"}
	glMap["CriminalGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["CriminalLearn"] = []string{"Administer", "AnyCombat", "Connect", "Notice", "Program", "Sneak", "Talk", "Trade"}
	glMap["DilettanteGrowth"] = []string{"AnyStat", "AnyStat", "AnyStat", "AnyStat", "Connect", "AnySkill"}
	glMap["DilettanteLearn"] = []string{"AnySkill", "AnySkill", "Connect", "Know", "Perform", "Pilot", "Talk", "Trade"}
	glMap["EntertainerGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Physical2", "Connect", "AnySkill"}
	glMap["EntertainerLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Perform", "Perform", "Sneak", "Talk"}
	glMap["MerchantGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["MerchantLearn"] = []string{"Administer", "AnyCombat", "Connect", "Fix", "Know", "Notice", "Trade", "Talk"}
	glMap["NobleGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["NobleLearn"] = []string{"Administer", "AnyCombat", "Connect", "Know", "Lead", "Notice", "Pilot", "Talk"}
	glMap["OfficialGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["OfficialLearn"] = []string{"Administer", "AnyCombat", "Connect", "Know", "Lead", "Notice", "Talk", "Trade"}
	glMap["PeasantGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Physical2", "Exert", "AnySkill"}
	glMap["PeasantLearn"] = []string{"Connect", "Exert", "Fix", "Notice", "Sneak", "Survive", "Trade", "Work"}
	glMap["PhysicianGrowth"] = []string{"AnyStat", "Physical2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["PhysicianLearn"] = []string{"Administer", "Connect", "Fix", "Heal", "Know", "Notice", "Talk", "Trade"}
	glMap["PilotGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["PilotLearn"] = []string{"Connect", "Exert", "Fix", "Notice", "Pilot", "Pilot", "Shoot", "Trade"}
	glMap["PoliticianGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["PoliticianLearn"] = []string{"Administer", "Connect", "Connect", "Lead", "Notice", "Perform", "Talk", "Talk"}
	glMap["ScholarGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["ScholarLearn"] = []string{"Administer", "Connect", "Fix", "Know", "Notice", "Perform", "Program", "Talk"}
	glMap["SoldierGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Physical2", "Exert", "AnySkill"}
	glMap["SoldierLearn"] = []string{"Administer", "AnyCombat", "Exert", "Fix", "Lead", "Notice", "Sneak", "Survive"}
	glMap["SpacerGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["SpacerLearn"] = []string{"Administer", "Connect", "Exert", "Fix", "Know", "Pilot", "Program", "Talk"}
	glMap["TechnicianGrowth"] = []string{"AnyStat", "Physical2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["TechnicianLearn"] = []string{"Administer", "Connect", "Exert", "Fix", "Fix", "Know", "Notice", "Pilot"}
	glMap["ThugGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Physical2", "Connect", "AnySkill"}
	glMap["ThugLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Sneak", "Stab", "Survive", "Talk", "Shoot"} // thug имеет "Stab or Shoot" на шестой позиции
	glMap["VagabondGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["VagabondLearn"] = []string{"AnyCombat", "Connect", "Notice", "Perform", "Pilot", "Sneak", "Survive", "Work"}
	glMap["WorkerGrowth"] = []string{"AnyStat", "AnyStat", "AnyStat", "AnyStat", "Exert", "AnySkill"}
	glMap["WorkerLearn"] = []string{"Administer", "AnySkill", "Connect", "Exert", "Fix", "Pilot", "Program", "Work"}

	return glMap
}

func freeSkill(background string) string {
	switch background {
	case "Barbarian":
		return "Survive"
	case "Clergy":
		return "Talk"
	case "Courtesan":
		return "Perform"
	case "Criminal":
		return "Sneak"
	case "Dilettante":
		return "Connect"
	case "Entertainer":
		return "Perform"
	case "Merchant":
		return "Trade"
	case "Noble":
		return "Lead"
	case "Official":
		return "Administer"
	case "Peasant":
		return "Exert"
	case "Physician":
		return "Heal"
	case "Pilot":
		return "Pilot"
	case "Politician":
		return "Talk"
	case "Scholar":
		return "Know"
	case "Soldier":
		return "AnyCombat"
	case "Spacer":
		return "Fix"
	case "Technician":
		return "Fix"
	case "Thug":
		return "AnyCombat"
	case "Vagabond":
		return "Survive"
	case "Worker":
		return "Work"
	}
	return "Free Skill Error"
}

func (npc *npc) develop(str string) {
	switch str {
	case "AnyStat":
		stat := atributeNames[rand.Intn(len(atributeNames))]
		npc.IncreaseStat(stat)
	case "Physical2":
		for i := 0; i < 2; i++ {
			stat := physicalAtr[rand.Intn(len(physicalAtr))]
			npc.IncreaseStat(stat)
		}
	case "Mental2":
		for i := 0; i < 2; i++ {
			stat := mentalAtr[rand.Intn(len(mentalAtr))]
			npc.IncreaseStat(stat)
		}
	case "AnySkill":
		skill := nonPsyhicSkills()[rand.Intn(len(nonPsyhicSkills()))]
		npc.IncreaseSkill(skill)
	case "AnyCombat":
		skill := combatSkills()[rand.Intn(len(combatSkills()))]
		npc.IncreaseSkill(skill)
	default:
		npc.IncreaseSkill(str)

	}
}

func (npc *npc) report() string {
	rep := ""
	rep += "Name: " + npc.name
	rep += "\nAttributes:"
	for i := range atributeNames {
		rep += "\n" + atributeNames[i] + ": " + strconv.Itoa(npc.attribute[atributeNames[i]]) + "(" + atrModS(npc.attribute[atributeNames[i]]) + ")"
	}
	rep += "\nBackground: " + npc.background
	rep += "\nSkills:"
	for key, val := range npc.skill {
		if val > -1 {
			rep += "\n" + key + "-" + strconv.Itoa(npc.skill[key])
		}
	}

	return rep
}

//increaseSkill
//increaseStat

//AnyStat
//Physical2
//Mental2
//AnySkill
//AnyCombat
