package main

import (
	"fmt"
	"strconv"
)

type vitalPoint struct {
	vpType  string
	point   string
	meaning string
}

type Leader struct {
	name     string
	leadSkl  int
	adminSkl int
	intAtr   int
	wisAtr   int
	chaAtr   int
	tl       int
	Units    int
	/*
		actions:
		Attack
		Defend
		Rest

	*/
}

func NewLeader() *Leader {
	leader := Leader{}
	leader.name = RandomName(false)
	leader.intAtr = rollXdY(3, 6)
	leader.wisAtr = rollXdY(3, 6)
	leader.chaAtr = rollXdY(3, 6)
	leader.leadSkl = setSkill()
	leader.adminSkl = setSkill()
	leader.tl = setSkill() + 1
	return &leader
}
func (l *Leader) toString() string {
	str := ""
	str = combineStrings(str, "Name: "+l.name+"\n")
	str = combineStrings(str, "Attributes\n")
	str = combineStrings(str, "Intelligence: "+strconv.Itoa(l.intAtr)+" ("+atrModS(l.intAtr)+")\n")
	str = combineStrings(str, "Wisdom      : "+strconv.Itoa(l.wisAtr)+" ("+atrModS(l.wisAtr)+")\n")
	str = combineStrings(str, "Charisma    : "+strconv.Itoa(l.chaAtr)+" ("+atrModS(l.chaAtr)+")\n")
	str = combineStrings(str, "Relevant Skills\n")
	str = combineStrings(str, "Lead        : "+strconv.Itoa(l.leadSkl)+"\n")
	str = combineStrings(str, "Admin       : "+strconv.Itoa(l.adminSkl)+"\n")
	str = combineStrings(str, "TL          : "+strconv.Itoa(l.tl)+"\n")
	return str
}

func (l *Leader) totalUnits() {
	un := atrMod(l.intAtr) + atrMod(l.wisAtr) + atrMod(l.chaAtr) + atrMod(l.adminSkl) + atrMod(l.leadSkl)
	if l.leadSkl == 0 {
		un--
	}
	if l.adminSkl == 0 {
		un--
	}
}

func atrMod(i int) int {
	if i > 17 {
		return 2
	}
	if i > 13 {
		return 1
	}
	if i < 4 {
		return -2
	}
	if i < 8 {
		return -1
	}
	return 0
}

func atrModS(i int) string {
	return strconv.Itoa(atrMod(i))
}

func setSkill() int {
	r := rollXdY(2, 6)
	skill := (r / 2) - 2
	if skill < 0 {
		skill = 0
	}
	return skill
}

func diverse() {
	a := 1
	b := 1
	c := 1
	for i := 0; i < 4; i++ {
		set := rollXdY(1, 3)
		switch set {
		case 1:
			a++
		case 2:
			b++
		case 3:
			c++
		}

	}
	fmt.Println("a =", a, "b =", b, "c =", c)
}

func main0() {
	randomSeed()
	for i := 0; i < 8; i++ {
		fmt.Println("Vital point", i+1, ":")
		vpType := vpType(rollXdY(1, 4))
		vpPoint := rollXdY(1, 6) - 1
		vpVital := rollXdY(1, 6) - 1
		place := GiveName("Any", "Place", -1)
		fmt.Println("   "+place+":", thisPoint(vpType, vpPoint), "is vital because it", vitalBecause(vpType, vpVital), " ("+vpType+")")
		diverse()
	}
}

func vpType(i int) string {
	switch i {
	case 1:
		return "Military"
	case 2:
		return "Industrial"
	case 3:
		return "Social"
	case 4:
		return "Political"
	}
	return "Unknown"
}

func thisPoint(vpType string, i int) string {
	switch vpType {
	case "Military":
		point := []string{
			"Elaborate fortification",
			"Site with fine natural defenses",
			"Ancient or new-made ruins",
			"Newly-erected strongpoint",
			"Recently-seized outpost",
			"Transit chokepoint",
		}
		return point[i]
	case "Industrial":
		point := []string{
			"Factory complex",
			"Heavily industrialized town",
			"Research or training center",
			"Supply depot transit nexus",
			"Community of skilled workers",
			"Mine, cropland, or resource site",
		}
		return point[i]
	case "Social":
		point := []string{
			"Major population center",
			"Sacred site",
			"Traditional seat of rulership",
			"Site of famous former victory",
			"Group of revered personages",
			"Place of safety and sanctuary",
		}
		return point[i]
	case "Political":
		point := []string{
			"Home of major nobility",
			"Town of regime supporters",
			"City of dubious loyalties",
			"Capital of region or polity",
			"Loyalty of regional gentry",
			"Oligarch’s business holding",
		}
		return point[i]
	}
	return "Unknown"
}

func vitalBecause(vpType string, i int) string {

	switch vpType {
	case "Military":
		point := []string{
			"Commands a strategic transit route",
			"Would savage any advance past it",
			"Has supplies to support an offensive",
			"Is a major recruit mustering point",
			"Is an operational base for a major command",
			"Is threatening an enemy position",
		}
		return point[i]
	case "Industrial":
		point := []string{
			"Produces a critical part or supply",
			"Is developing dangerous new weaponry",
			"Provides repair services to the army",
			"Has a crucial store of military supplies",
			"Is only source of a vital component",
			"Is producing morale-critical consumer goods",
		}
		return point[i]
	case "Social":
		point := []string{
			"It has historical value to morale",
			"Is a symbol of regime legitimacy",
			"Is a major source of recruits",
			"Steadies and calms the populace",
			"Is a symbol of normal life going on",
			"Is important to neutral or allied neighbors",
		}
		return point[i]
	case "Political":
		point := []string{
			"Its owners will turn on the regime if it falls",
			"Its owners are providing vital moral support",
			"It commands large sums of money",
			"Vital governmental functions happen there",
			"Votes or regime officials come from there",
			"Its security was promised to a critical ally",
		}
		return point[i]
	}
	return "Unknown"
}
