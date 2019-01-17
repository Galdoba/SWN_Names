package main

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
	world.Atmosphere = rollAtmosphere()
	world.Temperature = rollTemperature(world.Atmosphere)
	world.Biosphere = rollBiosphere(world.Atmosphere)
	world.Populate()
	return &world
}

func (w *World) toString() string {
	str := ""
	str = str + "Atmosphere: " + w.Atmosphere + "\n"
	str = str + "Temperature: " + w.Temperature + "\n"
	str = str + "Biosphere: " + w.Biosphere + "\n"
	str = str + "Population: " + w.Population + "\n"
	str = str + "Tech Level: " + w.TechLevel + "\n"
	str = str + "WorldTag1: " + w.WorldTag1 + "\n"
	str = str + "WorldTag2: " + w.WorldTag2 + "\n"
	str = str + "TradeTag: " + w.TradeTag + "\n"
	str = str + "LivingStandart: " + w.LivingStandart + "\n"
	return str
}

func (w *World) Populate() {
	r := rollXdY(2, 6)
	if w.Temperature == "Warm" || w.Temperature == "Cold" {
		r = r - 1
	}
	if w.Temperature == "Cold-Temperate" || w.Temperature == "Temperate-Warm" {
		r = r - 2
	}
	if w.Temperature == "Burning" || w.Temperature == "Frozen" {
		r = r - 3
	}
	if r < 2 {
		w.Population = "Outpost"
		return
	}
	if r > 12 {
		w.Population = "100 Million"
		if w.Atmosphere != "Breathable Mix" || w.Atmosphere != "Thick Atmosphere" || w.Atmosphere != "Thin Atmosphere" {
			w.Population = "10 Million"
			return
		}
	}
	if w.Atmosphere == "Breathable Mix" {
		switch r {
		case 2:
			w.Population = "Failed Colony"
		case 3:
			w.Population = "Outpost"
		case 4:
			w.Population = "10 Thousand"
		case 5:
			w.Population = "10 Thousand"
		case 6:
			w.Population = "100 Thousand"
		case 7:
			w.Population = "100 Thousand"
		case 8:
			w.Population = "Million"
		case 9:
			w.Population = "10 Million"
		case 10:
			w.Population = "100 Million"
		case 11:
			w.Population = "Billion"
		case 12:
			w.Population = "Alien"
		}
	}
	if w.Atmosphere == "Thick Atmosphere" {
		switch r {
		case 2:
			w.Population = "Failed Colony"
		case 3:
			w.Population = "Outpost"
		case 4:
			w.Population = "Outpost"
		case 5:
			w.Population = "10 Thousand"
		case 6:
			w.Population = "10 Thousand"
		case 7:
			w.Population = "100 Thousand"
		case 8:
			w.Population = "100 Thousand"
		case 9:
			w.Population = "100 Thousand"
		case 10:
			w.Population = "Million"
		case 11:
			w.Population = "10 Million"
		case 12:
			w.Population = "Alien"
		}
	}
	if w.Atmosphere == "Thin Atmosphere" {
		switch r {
		case 2:
			w.Population = "Failed Colony"
		case 3:
			w.Population = "Outpost"
		case 4:
			w.Population = "Outpost"
		case 5:
			w.Population = "10 Thousand"
		case 6:
			w.Population = "10 Thousand"
		case 7:
			w.Population = "100 Thousand"
		case 8:
			w.Population = "100 Thousand"
		case 9:
			w.Population = "100 Thousand"
		case 10:
			w.Population = "Million"
		case 11:
			w.Population = "10 Million"
		case 12:
			w.Population = "Alien"
		}
	}
	if w.Atmosphere == "Airless" || w.Atmosphere == "Inert Gas" {
		switch r {
		case 2:
			w.Population = "Failed Colony"
		case 3:
			w.Population = "Outpost"
		case 4:
			w.Population = "Outpost"
		case 5:
			w.Population = "Outpost"
		case 6:
			w.Population = "Outpost"
		case 7:
			w.Population = "10 Thousand"
		case 8:
			w.Population = "10 Thousand"
		case 9:
			w.Population = "10 Thousand"
		case 10:
			w.Population = "100 Thousand"
		case 11:
			w.Population = "Million"
		case 12:
			w.Population = "Alien"
		}
	}
	if w.Atmosphere == "Corrosive" || w.Atmosphere == "Invasive Atmosphere" || w.Atmosphere == "Corrosive and Invasive Atmosphere" {
		switch r {
		case 2:
			w.Population = "Failed Colony"
		case 3:
			w.Population = "Outpost"
		case 4:
			w.Population = "Outpost"
		case 5:
			w.Population = "Outpost"
		case 6:
			w.Population = "Outpost"
		case 7:
			w.Population = "Outpost"
		case 8:
			w.Population = "10 Thousand"
		case 9:
			w.Population = "10 Thousand"
		case 10:
			w.Population = "100 Thousand"
		case 11:
			w.Population = "Million"
		case 12:
			w.Population = "Alien"
		}
	}
}

func rollAtmosphere() string {
	r := rollXdY(2, 6)
	return pickAtmosphere(r)
}

func pickAtmosphere(index int) string {
	atmosphere := "Unknown"
	switch index {
	case 2:
		atmosphere = "Corrosive"
	case 3:
		atmosphere = "Inert Gas"
	case 4:
		atmosphere = "Airless"
	case 5:
		atmosphere = "Thin Atmosphere"
	case 6:
		atmosphere = "Breathable Mix"
	case 7:
		atmosphere = "Breathable Mix"
	case 8:
		atmosphere = "Breathable Mix"
	case 9:
		atmosphere = "Breathable Mix"
	case 10:
		atmosphere = "Thick Atmosphere"
	case 11:
		atmosphere = "Invasive Atmosphere"
	case 12:
		atmosphere = "Corrosive and Invasive Atmosphere"
	}
	return atmosphere
}

func pickTemperature(atmosphere string, index int) string {
	temperature := "Unknown"
	if atmosphere == "Breathable Mix" {
		switch index {
		case 2:
			temperature = "Cold-Temperate"
		case 3:
			temperature = "Cold-Temperate"
		case 4:
			temperature = "Cold"
		case 5:
			temperature = "Cold"
		case 6:
			temperature = "Temperate"
		case 7:
			temperature = "Temperate"
		case 8:
			temperature = "Temperate"
		case 9:
			temperature = "Warm"
		case 10:
			temperature = "Warm"
		case 11:
			temperature = "Temperate-Warm"
		case 12:
			temperature = "Temperate-Warm"
		}
	}
	if atmosphere == "Thick Atmosphere" {
		switch index {
		case 2:
			temperature = "Cold-Temperate"
		case 3:
			temperature = "Cold"
		case 4:
			temperature = "Cold"
		case 5:
			temperature = "Temperate"
		case 6:
			temperature = "Temperate"
		case 7:
			temperature = "Temperate"
		case 8:
			temperature = "Warm"
		case 9:
			temperature = "Warm"
		case 10:
			temperature = "Warm"
		case 11:
			temperature = "Warm"
		case 12:
			temperature = "Burning"
		}
	}
	if atmosphere == "Thin Atmosphere" {
		switch index {
		case 2:
			temperature = "Cold-Temperate"
		case 3:
			temperature = "Cold"
		case 4:
			temperature = "Cold"
		case 5:
			temperature = "Cold"
		case 6:
			temperature = "Cold"
		case 7:
			temperature = "Temperate"
		case 8:
			temperature = "Temperate"
		case 9:
			temperature = "Temperate"
		case 10:
			temperature = "Warm"
		case 11:
			temperature = "Warm"
		case 12:
			temperature = "Temperate-Warm"
		}
	}
	if atmosphere == "Airless" || atmosphere == "Inert Gas" {
		switch index {
		case 2:
			temperature = "Frozen"
		case 3:
			temperature = "Frozen"
		case 4:
			temperature = "Cold-Temperate"
		case 5:
			temperature = "Cold"
		case 6:
			temperature = "Cold"
		case 7:
			temperature = "Temperate"
		case 8:
			temperature = "Temperate"
		case 9:
			temperature = "Warm"
		case 10:
			temperature = "Warm"
		case 11:
			temperature = "Temperate-Warm"
		case 12:
			temperature = "Burning"
		}
	}
	if atmosphere == "Corrosive" || atmosphere == "Invasive Atmosphere" || atmosphere == "Corrosive and Invasive Atmosphere" {
		switch index {
		case 2:
			temperature = "Frozen"
		case 3:
			temperature = "Cold-Temperate"
		case 4:
			temperature = "Cold"
		case 5:
			temperature = "Cold"
		case 6:
			temperature = "Cold"
		case 7:
			temperature = "Temperate"
		case 8:
			temperature = "Warm"
		case 9:
			temperature = "Warm"
		case 10:
			temperature = "Warm"
		case 11:
			temperature = "Temperate-Warm"
		case 12:
			temperature = "Burning"
		}
	}
	return temperature
}

func rollTemperature(atmosphere string) string {
	r := rollXdY(2, 6)
	return pickTemperature(atmosphere, r)
}

func pickBiosphere(atmosphere string, index int) string {
	biosphere := "Unknown"
	if atmosphere == "Breathable Mix" {
		switch index {
		case 2:
			biosphere = "Remnant"
		case 3:
			biosphere = "None"
		case 4:
			biosphere = "Microbial"
		case 5:
			biosphere = "Microbial"
		case 6:
			biosphere = "Miscible"
		case 7:
			biosphere = "Miscible"
		case 8:
			biosphere = "Miscible"
		case 9:
			biosphere = "Immiscible"
		case 10:
			biosphere = "Hybrid"
		case 11:
			biosphere = "Hybrid"
		case 12:
			biosphere = "Engineered"
		}
	}
	if atmosphere == "Thick Atmosphere" {
		switch index {
		case 2:
			biosphere = "Remnant"
		case 3:
			biosphere = "None"
		case 4:
			biosphere = "Microbial"
		case 5:
			biosphere = "Microbial"
		case 6:
			biosphere = "Microbial"
		case 7:
			biosphere = "Miscible"
		case 8:
			biosphere = "Immiscible"
		case 9:
			biosphere = "Immiscible"
		case 10:
			biosphere = "Immiscible"
		case 11:
			biosphere = "Hybrid"
		case 12:
			biosphere = "Engineered"
		}
	}
	if atmosphere == "Thin Atmosphere" {
		switch index {
		case 2:
			biosphere = "Remnant"
		case 3:
			biosphere = "None"
		case 4:
			biosphere = "None"
		case 5:
			biosphere = "Microbial"
		case 6:
			biosphere = "Microbial"
		case 7:
			biosphere = "Miscible"
		case 8:
			biosphere = "Immiscible"
		case 9:
			biosphere = "Immiscible"
		case 10:
			biosphere = "Hybrid"
		case 11:
			biosphere = "Hybrid"
		case 12:
			biosphere = "Engineered"
		}
	}
	if atmosphere == "Airless" || atmosphere == "Inert Gas" {
		switch index {
		case 2:
			biosphere = "Remnant"
		case 3:
			biosphere = "None"
		case 4:
			biosphere = "None"
		case 5:
			biosphere = "None"
		case 6:
			biosphere = "None"
		case 7:
			biosphere = "None"
		case 8:
			biosphere = "None"
		case 9:
			biosphere = "None"
		case 10:
			biosphere = "Engineered"
		case 11:
			biosphere = "Engineered"
		case 12:
			biosphere = "Immiscible"
		}
	}
	if atmosphere == "Corrosive" || atmosphere == "Invasive Atmosphere" || atmosphere == "Corrosive and Invasive Atmosphere" {
		switch index {
		case 2:
			biosphere = "Remnant"
		case 3:
			biosphere = "None"
		case 4:
			biosphere = "None"
		case 5:
			biosphere = "None"
		case 6:
			biosphere = "Microbial"
		case 7:
			biosphere = "Microbial"
		case 8:
			biosphere = "Immiscible"
		case 9:
			biosphere = "Immiscible"
		case 10:
			biosphere = "Immiscible"
		case 11:
			biosphere = "Immiscible"
		case 12:
			biosphere = "Engineered"
		}
	}
	return biosphere
}

func rollBiosphere(atmosphere string) string {
	r := rollXdY(2, 6)
	return pickBiosphere(atmosphere, r)
}
