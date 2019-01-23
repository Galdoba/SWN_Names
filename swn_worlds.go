package main

import (
	"strconv"
)

type World struct {
	Atmosphere     string
	Temperature    string
	Biosphere      string
	Population     string
	PopulationNum  int
	TechLevel      string
	WorldTag1      string
	WorldTag2      string
	TradeTag       string
	LivingStandard string
	BP             int
}

func NewWorld() *World {
	world := World{}
	world.rollAtmosphere()
	world.rollTemperature()
	world.rollBiosphere()
	world.Populate()
	world.rollTechLevel()
	world.rollWorldTags()
	world.rollLivingStandard()
	world.calculateBP()
	//world.PopulateNum()
	return &world
}

func (w *World) rollWorldTags() {
	w.WorldTag1 = randomWorldTag()
	w.WorldTag2 = randomWorldTag()
	for w.WorldTag1 == w.WorldTag2 {
		w.WorldTag2 = randomWorldTag()
	}
}

func (w *World) toString() string {
	str := ""
	str = str + "Atmosphere: " + w.Atmosphere + "\n"
	str = str + "Temperature: " + w.Temperature + "\n"
	str = str + "Biosphere: " + w.Biosphere + "\n"
	if w.PopulationNum != 0 {
		str = str + "Population: " + strconv.Itoa(w.PopulationNum) + w.Population + "\n"
	} else {
		str = str + "Population: " + w.Population + "\n"
	}
	str = str + "Tech Level: " + w.TechLevel + "\n"
	str = str + "WorldTag1: " + w.WorldTag1 + "\n"
	str = str + "WorldTag2: " + w.WorldTag2 + "\n"
	str = str + "TradeTag: " + w.TradeTag + "\n"
	str = str + "LivingStandard: " + w.LivingStandard + " (" + strconv.Itoa(w.BP) + " BP or " + strconv.Itoa(w.BP*200000) + ")\n"
	return str
}

func (w *World) Populate() {

	r := rollXdY(2, 6)
	switch r {
	case 2:
		w.Population = "Failed Colony"

	case 3:
		w.Population = "Outpost"

	case 4:
		w.Population = "K"
		w.PopulationNum = roll1dX(100, 0)
	case 5:
		w.Population = "K"
		w.PopulationNum = roll1dX(1000, 0)
	case 6:
		w.Population = "M"
		w.PopulationNum = roll1dX(10, 0)
	case 7:
		w.Population = "M"
		w.PopulationNum = roll1dX(10, 0)
	case 8:
		w.Population = "M"
		w.PopulationNum = roll1dX(100, 0)
	case 9:
		w.Population = "M"
		w.PopulationNum = roll1dX(100, 0)
	case 10:
		w.Population = "M"
		w.PopulationNum = roll1dX(1000, 0)
	case 11:
		w.Population = "B"
		w.PopulationNum = roll1dX(10, 0)
	case 12:
		w.Population = "Alien"
	}
}

func (w *World) rollAtmosphere() {
	r := rollXdY(2, 6)
	w.pickAtmosphere(r)
}

func (w *World) pickAtmosphere(index int) {
	switch index {
	case 2:
		w.Atmosphere = "Corrosive"
	case 3:
		w.Atmosphere = "Inert Gas"
	case 4:
		w.Atmosphere = "Thin Atmosphere"
	case 5:
		w.Atmosphere = "Breathable Mix"
	case 6:
		w.Atmosphere = "Breathable Mix"
	case 7:
		w.Atmosphere = "Breathable Mix"
	case 8:
		w.Atmosphere = "Breathable Mix"
	case 9:
		w.Atmosphere = "Breathable Mix"
	case 10:
		w.Atmosphere = "Thick Atmosphere"
	case 11:
		w.Atmosphere = "Invasive Atmosphere"
	case 12:
		w.Atmosphere = "Corrosive and Invasive Atmosphere"
	}

}

func (w *World) pickTemperature(index int) {

	switch index {
	case 2:
		w.Temperature = "Frozen"
	case 3:
		w.Temperature = "Cold"
	case 4:
		w.Temperature = "Variable Cold"
	case 5:
		w.Temperature = "Variable Cold"
	case 6:
		w.Temperature = "Temperate"
	case 7:
		w.Temperature = "Temperate"
	case 8:
		w.Temperature = "Temperate"
	case 9:
		w.Temperature = "Variable Warm"
	case 10:
		w.Temperature = "Variable Warm"
	case 11:
		w.Temperature = "Warm"
	case 12:
		w.Temperature = "Burning"
	}
}

func (w *World) rollTemperature() {
	r := rollXdY(2, 6)
	w.pickTemperature(r)
}

func (w *World) pickBiosphere(index int) {

	switch index {
	case 2:
		w.Biosphere = "Remnant"
	case 3:
		w.Biosphere = "Microbial"
	case 4:
		w.Biosphere = "None"
	case 5:
		w.Biosphere = "None"
	case 6:
		w.Biosphere = "Human-Miscible"
	case 7:
		w.Biosphere = "Human-Miscible"
	case 8:
		w.Biosphere = "Human-Miscible"
	case 9:
		w.Biosphere = "Immiscible"
	case 10:
		w.Biosphere = "Immiscible"
	case 11:
		w.Biosphere = "Hybrid"
	case 12:
		w.Biosphere = "Engineered"
	default:
		w.Biosphere = "None"
	}

}

func (w *World) rollBiosphere() {
	r := rollXdY(2, 6)
	w.pickBiosphere(r)
}

func (w *World) pickTechlevel(index int) {
	switch index {
	case 2:
		w.TechLevel = "TL0"
	case 3:
		w.TechLevel = "TL1"
	case 4:
		w.TechLevel = "TL2"
	case 5:
		w.TechLevel = "TL3+"
	case 6:
		w.TechLevel = "TL3"
	case 7:
		w.TechLevel = "TL4"
	case 8:
		w.TechLevel = "TL4"
	case 9:
		w.TechLevel = "TL4"
	case 10:
		w.TechLevel = "TL4-"
	case 11:
		w.TechLevel = "TL4+"
	case 12:
		w.TechLevel = "TL5"
	default:
		w.TechLevel = "None"
	}
}

func (w *World) rollTechLevel() {
	r := rollXdY(2, 6)
	w.pickTechlevel(r)
}

func worldTags() []string {
	tagList := []string{
		"Abandoned Colony",
		"Alien Ruins",
		"Altered Humanity",
		"Anarchists",
		"Anthropomorphs",
		"Area 51",
		"Badlands World",
		"Battleground",
		"Beastmasters",
		"Bubble Cities",
		"Cheap Life",
		"Civil War",
		"Cold War",
		"Colonized Population",
		"Cultural Power",
		"Cybercommunists",
		"Cyborgs",
		"Cyclical Doom",
		"Desert World",
		"Doomed World",
		"Dying Race",
		"Eugenic Cult",
		"Exchange Consulate",
		"Fallen Hegemon",
		"Feral World",
		"Flying Cities",
		"Forbidden Tech",
		"Former Warriors",
		"Freak Geology",
		"Freak Weather",
		"Friendly Foe",
		"Gold Rush",
		"Great Work",
		"Hatred",
		"Heavy Industry",
		"Heavy Mining",
		"Hivemind",
		"Holy War",
		"Hostile Biosphere",
		"Hostile Space",
		"Immortals",
		"Local Specialty",
		"Local Tech",
		"Major Spaceyard",
		"Mandarinate",
		"Mandate Base",
		"Maneaters",
		"Megacorps",
		"Mercenaries",
		"Minimal Contact",
		"Misandry/Misogyny",
		"Night World",
		"Nomads",
		"Oceanic World",
		"Out of Contact",
		"Outpost World",
		"Perimeter Agency",
		"Pilgrimage Site",
		"Pleasure World",
		"Police State",
		"Post-Scarcity",
		"Preceptor Archive",
		"Pretech Cultists",
		"Primitive Aliens",
		"Prison Planet",
		"Psionics Academy",
		"Psionics Fear",
		"Psionics Worship",
		"Quarantined World",
		"Radioactive World",
		"Refugees",
		"Regional Hegemon",
		"Restrictive Laws",
		"Revanchists",
		"Revolutionaries",
		"Rigid Culture",
		"Rising Hegemon",
		"Ritual Combat",
		"Robots",
		"Seagoing Cities",
		"Sealed Menace",
		"Secret Masters",
		"Sectarians",
		"Seismic Instability",
		"Shackled World",
		"Societal Despair",
		"Sole Supplier",
		"Taboo Treasure",
		"Terraform Failure",
		"Theocracy",
		"Tomb World",
		"Trade Hub",
		"Tyranny",
		"Unbraked AI",
		"Urbanized Surface",
		"Utopia",
		"Warlords",
		"Xenophiles",
		"Xenophobes",
		"Zombies",
	}
	return tagList
}

func pickWorldTag(index int) string {
	return worldTags()[index]
}

func randomWorldTag() string {
	r := roll1dX(len(worldTags()), -1)
	return pickWorldTag(r)
}

func (w *World) pickLivingStandard(index int) {
	switch index {
	case 2:
		w.LivingStandard = "Slum"
	case 3:
		w.LivingStandard = "Slum"
	case 4:
		w.LivingStandard = "Poor"
	case 5:
		w.LivingStandard = "Poor"
	case 6:
		w.LivingStandard = "Poor"
	case 7:
		w.LivingStandard = "Common"
	case 8:
		w.LivingStandard = "Common"
	case 9:
		w.LivingStandard = "Common"
	case 10:
		w.LivingStandard = "Common"
	case 11:
		w.LivingStandard = "Good"
	case 12:
		w.LivingStandard = "Good"
	default:
		w.LivingStandard = "Elite"
	}
}

func (w *World) rollLivingStandard() {
	r := rollXdY(2, 6)
	w.pickLivingStandard(r)
}

func (w *World) calculateBP() {
	switch w.Population {
	case "K":
		if w.PopulationNum > 100 {
			switch w.LivingStandard {
			case "Slum":
				w.BP = 20
			case "Poor":
				w.BP = 47
			case "Common":
				w.BP = 75
			case "Good":
				w.BP = 130
			case "Elite":
				w.BP = 260
			}
		}
	case "B":
		switch w.LivingStandard {
		case "Slum":
			w.BP = 338
		case "Poor":
			w.BP = 562
		case "Common":
			w.BP = 778
		case "Good":
			w.BP = 1250
		case "Elite":
			w.BP = 2500
		}
	default:
		if w.PopulationNum < 11 {
			switch w.LivingStandard {
			case "Slum":
				w.BP = 62
			case "Poor":
				w.BP = 125
			case "Common":
				w.BP = 187
			case "Good":
				w.BP = 312
			case "Elite":
				w.BP = 625
			}
			return
		}
		if w.PopulationNum < 101 {
			switch w.LivingStandard {
			case "Slum":
				w.BP = 125
			case "Poor":
				w.BP = 237
			case "Common":
				w.BP = 350
			case "Good":
				w.BP = 550
			case "Elite":
				w.BP = 1100
			}
			return
		}
		switch w.LivingStandard {
		case "Slum":
			w.BP = 225
		case "Poor":
			w.BP = 387
		case "Common":
			w.BP = 550
		case "Good":
			w.BP = 775
		case "Elite":
			w.BP = 1650
		}
		return
	}

}
