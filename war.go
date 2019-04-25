package main

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/utils"
)

type vitalPoint struct {
	vpType  string
	point   string
	meaning string
}

type Army struct {
	leader    Leader
	manpower  int
	tl        int
	unitScore int
	units     []Unit
}

type Unit struct {
	name      string
	strength  int
	condition int
}

func NewArmy(manpower int, tl int) *Army {
	army := &Army{}
	army.leader = *NewLeader()
	army.tl = tl
	army.manpower = manpower
	army.unitScore = army.leader.getCommandScore()
	army.formCorps()

	return army
}

func (army *Army) report() string {
	repo := ""
	for i := range army.units {
		repo += "Unit: " + army.units[i].name
	}
	return repo
}

func NewUnit(unitStr int) *Unit {
	unit := &Unit{}

	unit.condition = unitStr
	unit.strength = unit.baseScore()
	unit.name = ""
	return unit
}

func (unit *Unit) nameit() {

}

func (unit *Unit) SetModifiedStrengh(newScore int) {
	unit.strength = newScore
}

func (unit *Unit) baseScore() int {
	cond := unit.condition
	switch cond / 4 {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 8
	case 5:
		return 12
	case 6:
		return 16
	}
	return -1
}

func (unit *Unit) toString() string {
	return unit.unitStatusStr() + " " + unit.unitTypeStr()
}

func (unit *Unit) unitTypeStr() string {
	cond := unit.condition
	switch cond / 4 {
	case 0:
		return "Battalion"
	case 1:
		return "Brigade"
	case 2:
		return "Division"
	case 3:
		return "Corps"
	case 4:
		return "Army"
	case 5:
		return "Army Front"
	case 6:
		return "Theater"
	}
	return "Unknown"
}

func (unit *Unit) unitStatusStr() string {
	cond := unit.condition
	switch cond % 4 {
	case 0:
		return "Crippled"
	case 1:
		return "Bloodied"
	case 2:
		return "Ready"
	case 3:
		return "Fresh"
	}
	return "Unknown"
}

func (leader *Leader) personalityScore() int {
	return atrMod(leader.intAtr) + atrMod(leader.wisAtr) + atrMod(leader.chaAtr)
}

func (army *Army) formCorps() {
	recruits := army.manpower
	maxUnits := army.tl + utils.Min(army.leader.leadSkl, army.leader.adminSkl) + army.leader.personalityScore()
	fmt.Println("maxUnits", maxUnits)

	fmt.Println("Total recruits:", recruits)
	for recruits/10000000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 10000000
		enlisted := 10000000
		if recruits < 0 {
			enlisted = 10000000 + recruits

		}
		army.units = append(army.units, *NewUnit(28))
		fmt.Println(enlisted, "men enlisted to Theater (power 16)")
	}
	for recruits/2000000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 2000000
		enlisted := 2000000
		if recruits < 0 {
			enlisted = 2000000 + recruits
		}
		army.units = append(army.units, *NewUnit(24))
		fmt.Println(enlisted, "men enlisted to Army Front (power 12)")
	}
	for recruits/200000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 200000
		enlisted := 200000
		if recruits < 0 {
			enlisted = 200000 + recruits
		}
		army.units = append(army.units, *NewUnit(20))
		fmt.Println(enlisted, "men enlisted to Army (power 8)")
	}
	for recruits/40000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 40000
		enlisted := 40000
		if recruits < 0 {
			enlisted = 40000 + recruits
		}
		army.units = append(army.units, *NewUnit(16))
		fmt.Println(enlisted, "men enlisted to Corps (power 4)")
	}
	for recruits/18000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 18000
		enlisted := 18000
		if recruits < 0 {
			enlisted = 18000 + recruits
		}
		army.units = append(army.units, *NewUnit(12))
		fmt.Println(enlisted, "men enlisted to Division (power 2)")
	}
	for recruits/5000 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 5000
		enlisted := 5000
		if recruits < 0 {
			enlisted = 5000 + recruits
		}
		army.units = append(army.units, *NewUnit(8))
		fmt.Println(enlisted, "men enlisted to Brigade (power 1)")
	}
	for recruits/600 > 0 {
		if len(army.units) > maxUnits {
			return
		}
		recruits = recruits - 600
		enlisted := 600
		if recruits < 0 {
			enlisted = 600 + recruits
		}
		army.units = append(army.units, *NewUnit(4))
		fmt.Println(enlisted, "men enlisted to Batalion (power 0)")
	}
	freeScore := maxUnits - len(army.units)
	for i := 0; i < freeScore; i++ {
		army.units[0].strength++
	}
}

type Leader struct {
	name         string
	leadSkl      int
	adminSkl     int
	intAtr       int
	wisAtr       int
	chaAtr       int
	commandScore int
	/*
		actions:
		Attack
		Defend
		Rest

	*/
}

func NewLeader() *Leader {
	leader := &Leader{}
	npc := CreateNPC()
	leader.name = npc.name
	leader.intAtr = npc.attribute["Int"]
	leader.wisAtr = npc.attribute["Wis"]
	leader.chaAtr = npc.attribute["Cha"]
	leader.leadSkl = npc.skill["Lead"].skillLevel
	leader.adminSkl = npc.skill["Administer"].skillLevel
	return leader
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
	str = combineStrings(str, "Command Score: "+strconv.Itoa(l.commandScore)+"\n")
	return str
}

func (l *Leader) getCommandScore() int {
	un := atrMod(l.intAtr) + atrMod(l.wisAtr) + atrMod(l.chaAtr)
	skillBuff := l.leadSkl
	if skillBuff > l.adminSkl {
		skillBuff = l.adminSkl
	}
	if skillBuff == 0 {
		un--
	} else {
		un = un + skillBuff
	}
	if un < 1 {
		un = 1
	}
	return un
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
