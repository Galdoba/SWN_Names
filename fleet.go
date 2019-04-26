package main

import (
	"fmt"

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
	shipType         string
	shipHull         string
	maintainanceCost int
}

func (w *World) assembleFleet() {
	bp := w.BP
	w.generateFleetPolicy()
	policy := w.FleetPolicy
	var Ships []ship
	site := randomSite()
	theme := RandomFromSliceStr([]string{"Place", "Sur"})
	for bp > w.BP*10/100 {
		shpType := utils.RandomFromList(policy)
		funcMap := make(map[string][]string)
		funcMap["Civilian Ships"] = civilanShipsList()
		funcMap["Military Ships"] = militaryShipsList()
		funcMap["Stations"] = stationsList()
		hull := utils.RandomFromList(funcMap[shpType])
		if bp > shipYerlyBP(hull) {
			Ships = append(Ships, ship{GiveName(site, theme, -1), hull, shipYerlyBP(hull)})
			bp = bp - shipYerlyBP(hull)
		}
		//fmt.Println(bp, shpType, hull)
	}
	totalMaint := w.BP - bp
	fmt.Println("Total BP spent on maintainance", totalMaint)
	for i := range Ships {
		fmt.Println(Ships[i])
	}
	fmt.Println(desidionPoolStr("Fleet policy", w.FleetPolicy))

}
