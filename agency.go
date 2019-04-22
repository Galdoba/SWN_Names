package main

import (
	"strconv"
)

type Agency struct {
	Attributes map[string]int
	Elements   map[string]element
}

func newAttributeMap() map[string]int {
	aMap := make(map[string]int)
	aMap["Connections"] = 0
	aMap["Infiltration"] = 0
	aMap["Mobility"] = 0
	aMap["Muscle"] = 0
	aMap["Resources"] = 0
	aMap["Security"] = 0
	aMap["Tech"] = 0
	return aMap
}

type element struct {
	//name string - не нужно будет в МАР
	level      int
	benefitsTo string
}

func newElementMap() map[string]element {
	eMap := make(map[string]element)
	eMap["Armory"] = element{0, "Resources"}
	eMap["Assassins"] = element{0, "Muscle"}
	eMap["Beamgates"] = element{0, "Mobility"}
	eMap["Black Codex"] = element{0, "Tech"}
	eMap["Criminal Ties"] = element{0, "Connections"}
	eMap["Early Warning"] = element{0, "Security"}
	eMap["Front Business"] = element{0, "Resources"}
	eMap["Hidden Strings"] = element{0, "Connections"}
	eMap["Identity Shop"] = element{0, "Infiltration"}
	eMap["Internal Security"] = element{0, "Security"}
	eMap["Legitimacy"] = element{0, "Connections"}
	eMap["Medical Lab"] = element{0, "Tech"}
	eMap["Military Backing"] = element{0, "Muscle"}
	eMap["Money"] = element{0, "Resources"}
	eMap["Pretech Lab"] = element{0, "Tech"}
	eMap["Psychics"] = element{0, "Infiltration"}
	eMap["Starships"] = element{0, "Muscle"}
	eMap["Stations"] = element{0, "Mobility"}
	eMap["Tradition"] = element{0, "Security"}
	eMap["Training"] = element{0, "NULL"}
	eMap["Transport"] = element{0, "Mobility"}
	return eMap
}

func SetElementLevel(name string, lvl int) element {
	if lvl > 3 {
		lvl = 3
	}
	if lvl < 0 {
		lvl = 0
	}
	elem := newElementMap()[name]
	elem.level = lvl
	return elem
}

func NewAgency(name string, lvl1, lvl2 int) *Agency {
	agency := &Agency{}
	agency.Attributes = newAttributeMap()
	agency.Elements = newElementMap()
	for key, val := range agency.Elements {
		if lvl1 > 0 {
			lvl1--
			agency.Elements[key] = SetElementLevel(key, 1)
			continue
		}
		if lvl2 > 0 {
			lvl2--
			agency.Elements[key] = SetElementLevel(key, 2)
			continue
		}

	}
	return agency
}

func (ag *Agency) Update() {
	ag.Attributes = newAttributeMap()
	for key, val := range ag.Elements {
		bonusInt := val.level
		ag[key] = val
	}
}

func (ag *Agency) Report() string {
	str := ""
	str += "Connections = " + strconv.Itoa(ag.Attributes["Connections"]) + "\n"
	str += "Infiltration = " + strconv.Itoa(ag.Attributes["Infiltration"]) + "\n"
	str += "Mobility = " + strconv.Itoa(ag.Attributes["Mobility"]) + "\n"
	str += "Muscle = " + strconv.Itoa(ag.Attributes["Muscle"]) + "\n"
	str += "Resources = " + strconv.Itoa(ag.Attributes["Resources"]) + "\n"
	str += "Security = " + strconv.Itoa(ag.Attributes["Security"]) + "\n"
	str += "Tech = " + strconv.Itoa(ag.Attributes["Tech"]) + "\n"
	return str
}
