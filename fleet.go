package main

import (
	"strconv"

	"github.com/Galdoba/utils"
)

func (w *World) generateFleetPolicy() {
	desidionPool, _ := generateDesidionPool(7, "Civilian Ships", "Military Ships", "Stations")
	w.FleetPolicy = desidionPool
}

func shipYerlyBP(key string) int {
	sshipCost := make(map[string]int)
	sshipCost["Fighter"] = 1
	sshipCost["Shuttle"] = 2
	sshipCost["Free Merchant"] = 1
	sshipCost["Naval Courier"] = 3
	sshipCost["Patrol Boat"] = 4
	sshipCost["Corvette"] = 7
	sshipCost["Bulk Freighter"] = 7
	sshipCost["Frigate"] = 8
	sshipCost["Heavy Frigate"] = 22
	sshipCost["Cruiser"] = 29
	sshipCost["Siege Cruiser"] = 28
	sshipCost["Troop Transport"] = 13
	sshipCost["Logistics Ship"] = 23
	sshipCost["Battleship"] = 239
	sshipCost["Carrier"] = 363
	sshipCost["Bannerjee 12"] = 9
	sshipCost["Peerless"] = 12
	sshipCost["Shantadurga"] = 41
	sshipCost["Scutum"] = 78
	sshipCost["Arx"] = 96
	return sshipCost[key]
}

func civilanShipsList() []string {
	sList := []string{
		"Shuttle",
		"Free Merchant",
		"Bulk Freighter",
	}
	return sList
}

func militaryShipsList() []string {
	sList := []string{
		"Fighter",
		"Naval Courier",
		"Patrol Boat",
		"Corvette",
		"Frigate",
		"Heavy Frigate",
		"Cruiser",
		"Siege Cruiser",
		"Troop Transport",
		"Logistics Ship",
		"Battleship",
		"Carrier",
	}
	return sList
}

func stationsList() []string {
	sList := []string{
		"Bannerjee 12",
		"Peerless",
		"Shantadurga",
		"Scutum",
		"Arx",
	}
	return sList
}

type ship struct {
	name             string
	shipHull         string
	maintainanceCost int
}

func (w *World) assembleFleet() {
	bp := w.BP
	w.generateFleetPolicy()
	policy := w.FleetPolicy
	theme := RandomFromSliceStr([]string{"Place", "Sur"})
	for bp > w.BP*10/100 {
		shpType := utils.RandomFromList(policy)
		funcMap := make(map[string][]string)
		funcMap["Civilian Ships"] = civilanShipsList()
		funcMap["Military Ships"] = militaryShipsList()
		funcMap["Stations"] = stationsList()
		hull := utils.RandomFromList(funcMap[shpType])
		if bp >= shipYerlyBP(hull) {
			w.Fleet = append(w.Fleet, *NewShip(GiveName(randomSite(), theme, -1), hull))
			bp = bp - shipYerlyBP(hull)
		}
	}
}

func (w *World) addToFleet(sh ship) {
	w.Fleet = append(w.Fleet, sh)

}

func NewShip(name string, hull string) *ship {
	newShip := &ship{}
	newShip.name = name
	newShip.shipHull = hull
	newShip.maintainanceCost = shipYerlyBP(hull)
	return newShip
}

func (shp *ship) isStation() bool {
	list := stationsList()
	for i := range list {
		if shp.shipHull == list[i] {
			return true
		}
	}
	return false
}

func (shp *ship) isMilitary() bool {
	list := militaryShipsList()
	for i := range list {
		if shp.shipHull == list[i] {
			return true
		}
	}
	return false
}

func (shp *ship) isCivilian() bool {
	list := civilanShipsList()
	for i := range list {
		if shp.shipHull == list[i] {
			return true
		}
	}
	return false
}

func (w *World) fleetReport() string {
	report := "Fleet Report: \n"
	report += "Stations: \n"
	stationMap := make(map[string]int)
	for i := range w.Fleet {
		curShip := w.Fleet[i]
		if curShip.isStation() {
			stationMap[curShip.shipHull] = stationMap[curShip.shipHull] + 1
		}
	}
	for key, val := range stationMap {
		report += key + " " + strconv.Itoa(val) + "\n"
	}
	report += "Civilian Ships: \n"
	civilMap := make(map[string]int)
	for i := range w.Fleet {
		curShip := w.Fleet[i]
		if curShip.isCivilian() {
			civilMap[curShip.shipHull] = civilMap[curShip.shipHull] + 1
		}
	}
	for key, val := range civilMap {
		report += key + " " + strconv.Itoa(val) + "\n"
	}
	report += "Military Ships: \n"
	milMap := make(map[string]int)
	for i := range w.Fleet {
		curShip := w.Fleet[i]
		if curShip.isMilitary() {
			milMap[curShip.shipHull] = milMap[curShip.shipHull] + 1
		}
	}
	for key, val := range milMap {
		report += key + " " + strconv.Itoa(val) + "\n"
	}
	return report
}
