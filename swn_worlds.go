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
	LivingStandart string
}

func NewWorld() *World {
	world := World{}
	world.rollAtmosphere()
	world.rollTemperature()
	world.rollBiosphere()
	world.Populate()
	//world.PopulateNum()
	return &world
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
	str = str + "LivingStandart: " + w.LivingStandart + "\n"
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
		w.PopulationNum = roll1dX(10, 0)
	case 5:
		w.Population = "K"
		w.PopulationNum = roll1dX(100, 0)
	case 6:
		w.Population = "M"
		w.PopulationNum = roll1dX(5, 0)
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

func (w *World) PopulateNum() {
	switch w.Population {
	case "Failed Colony":
		r := roll1dX(6, 0)
		if r < 4 {
			w.PopulationNum = roll1dX(100, 0)
		} else {
			w.Population = "Outpost"
		}
	case "Alien":
		r := roll1dX(6, 0)
		switch r {
		case 1:
			w.Population = "Remnant Alien"
		case 2:
			w.Population = "100 Thousand Alien"
		case 3:
			w.Population = "Million Alien"
		case 4:
			w.Population = "10 Million Alien"
		case 5:
			w.Population = "100 Million Alien"
		case 6:
			w.Population = "Billion Alien"
		}
	case "Outpost":
		r := roll1dX(6, 0)
		if inRange(r, 1, 3) {
			w.PopulationNum = roll1dX(10, 0) * 100
		}
		if inRange(r, 4, 5) {
			w.PopulationNum = roll1dX(10, 0) * 500
		}
		if r == 6 {
			w.PopulationNum = roll1dX(10, 0) * 1000
		}
	case "10 Thousand":
		w.PopulationNum = roll1dX(10, 0) * 10000
	case "100 Thousand":
		w.PopulationNum = roll1dX(10, 0) * 100000
	case "Million":
		w.PopulationNum = roll1dX(10, 0) * 1000000
	case "10 Million":
		w.PopulationNum = roll1dX(10, 0) * 10000000
	case "100 Million":
		w.PopulationNum = roll1dX(10, 0) * 100000000
	case "Billion":
		r := roll1dX(6, 0)
		if inRange(r, 1, 3) {
			w.PopulationNum = 1000000000
		}
		if r == 4 {
			w.PopulationNum = roll1dX(3, 0) * 1000000000
		}
		if r == 5 {
			w.PopulationNum = roll1dX(6, 0) * 1000000000
		}
		if r == 6 {
			w.PopulationNum = roll1dX(10, 0) * 1000000000
		}

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
