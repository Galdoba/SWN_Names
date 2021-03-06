package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

type World struct {
	Atmosphere        string
	Temperature       string
	Biosphere         string
	Population        string
	PopulationNum     int
	PopCode           int
	GovCode           int
	GovType           string
	LawCode           int
	RestrictionPolicy []string
	TechLevel         string
	WorldTag1         string
	WorldTag2         string
	TradeTag          string
	BudgetPolicy      []string
	LivingStandard    string
	BP                int
	Export            []Commodite
	TradeSpecifics    []int
	FleetPolicy       []string
	Fleet             []ship

	ColonyReport string
}

func NewWorld() *World {
	world := &World{}
	world.rollAtmosphere()
	world.rollTemperature()
	world.rollBiosphere()
	world.Populate()
	if world.Population == "Failed Colony" {
		world.ColonyReport = RandomColonyReport()
	}
	world.rollTechLevel()
	world.rollWorldTags()
	world.rollTradeTag()
	world.rollLivingStandard()
	world.generateBudgetPolicy()
	world.calculateBP()
	world.adjustBP()
	world.law()

	world.TradeSpecifics = marketSpecific()
	again := true
	for again {
		newCommoditie := *NewCommodityWithTag(pickCargoType(world.TradeSpecifics[0]))
		if !newCommoditie.isContainedBy(world.Export) {
			world.Export = append(world.Export, newCommoditie)
			again = false
		}
	}
	//world.Export = append(world.Export, *NewCommodityWithTag(pickCargoType(world.TradeSpecifics[0])))
	//	world.Export = append(world.Export, *NewCommodityWithTag(pickCargoType(world.TradeSpecifics[0])))
	//	world.Export = append(world.Export, *NewCommodityWithTag(pickCargoType(world.TradeSpecifics[1])))
	normCom := CreateCommodities(9)
	for i := range normCom {
		world.Export = append(world.Export, normCom[i])
	}
	for again {
		newCommoditie := *NewCommodityWithTag(pickCargoType(world.TradeSpecifics[0]))
		if !newCommoditie.isContainedBy(world.Export) {
			world.Export = append(world.Export, newCommoditie)
			again = false
		}
	}
	world.assembleFleet()
	return world
}

func NewUniqueCommoditie(commList []Commodite, tag string) *Commodite {

	return NewCommoditie()
}

func (w *World) rollWorldTags() {
	w.WorldTag1 = randomWorldTag()
	w.WorldTag2 = randomWorldTag()
	for w.WorldTag1 == w.WorldTag2 {
		w.WorldTag2 = randomWorldTag()
	}
}

func (w *World) showExport() string {
	str := ""
	for i := range w.Export {
		str += "\n Commoditie " + strconv.Itoa(i) + ":"
		str += commoditieStr(w.Export[i])
	}
	return str
}

func commoditieStr(c Commodite) string {
	str := "{" + c.name + "} {"
	for i := range c.tags {
		str += c.tags[i] + ", "
	}
	str = strings.TrimSuffix(str, ", ")
	str += "} {priceMod:" + strconv.Itoa(c.pricemod)
	str += "} {pricePerUnit:" + strconv.Itoa(c.costPerUnit) + "}"

	return str
}

func (comm *Commodite) CommoditeStr() string {
	str := ""
	str += comm.tags[0]
	return str
}

func (w *World) rollTradeTag() {
	w.TradeTag = randomTradeTag()
}

func (w *World) generateBudgetPolicy() {
	desidionPool, _ := generateDesidionPool(10, "Import", "Export", "Military")
	w.BudgetPolicy = desidionPool
}

func (w *World) toString() string {
	str := ""
	str = str + "Atmosphere: " + w.Atmosphere + "\n"
	str = str + "Temperature: " + w.Temperature + "\n"
	str = str + "Biosphere: " + w.Biosphere + "\n"
	if w.PopulationNum != 0 {
		str = str + "Population: " + strconv.Itoa(w.PopulationNum) + w.Population + "\n"
	} else {
		str = str + "Population: " + w.populationStr() + "\n"
	}
	str = str + "Population: " + w.populationStr() + "\n"
	str = str + "Tech Level: " + w.TechLevel + "\n"
	str = str + "WorldTag1: " + w.WorldTag1 + "\n"
	str = str + "WorldTag2: " + w.WorldTag2 + "\n"
	str = str + "TradeTag: " + w.TradeTag + "\n"
	str = str + "World trade specific: {" + pickCargoType(w.TradeSpecifics[0]) + ": -2} {" + pickCargoType(w.TradeSpecifics[1]) + ": -1} {" + pickCargoType(w.TradeSpecifics[2]) + ": 1} {" + pickCargoType(w.TradeSpecifics[3]) + ": 2}\n"
	str = str + "TradeCode: " + strconv.Itoa(w.TradeSpecifics[0]) + "-" + strconv.Itoa(w.TradeSpecifics[1]) + "-" + strconv.Itoa(w.TradeSpecifics[2]) + "-" + strconv.Itoa(w.TradeSpecifics[3]) + "\n"
	str = str + "World Export:" + w.showExport() + "\n"
	str += "Colony can produce " + strconv.Itoa(w.planetaryTax()*policyShare(w.BudgetPolicy, "Export")/10*2) + " credits worth commodities\n"
	str += "Colony can consume " + strconv.Itoa(w.planetaryTax()*policyShare(w.BudgetPolicy, "Import")/10) + " credits worth commodities\n"
	str = str + "LivingStandard: " + w.LivingStandard + " (" + strconv.Itoa(w.BP) + " BP or " + strconv.Itoa(w.BP*200000) + ")\n"
	str = str + "Law Level: " + strconv.Itoa(w.LawCode) + "\n"
	str += "Goverment Type: " + w.GovType + "\n"
	str += w.fleetReport() + "\n"
	str = str + "\nWorldTag1 (descr): " + describeTag(w.WorldTag1) + "\n"
	str = str + "\nWorldTag2 (descr): " + describeTag(w.WorldTag2) + "\n"
	str = str + "\nTrade Tag (descr): " + describeTradeTag(w.TradeTag) + "\n"
	str = str + "\nSociety: " + oneRollSociety() + "\n"
	str = str + "\nRulers: " + oneRollRulers() + "\n"
	str = str + "\nCommon Folk: " + oneRollRuled() + "\n"
	return str
}

func (w *World) Populate() {

	r := rollXdY(2, 6)
	switch r {
	case 2:
		w.Population = "Failed Colony"

	case 3:
		w.Population = "Outpost"
		w.PopCode = 4
	// case 4:
	// 	w.Population = "K"
	// 	w.PopulationNum = roll1dX(400, 100)
	// 	w.PopCode = 5
	// case 5:
	// 	w.Population = "K"
	// 	w.PopulationNum = roll1dX(900, 100)
	// 	w.PopCode = 5
	// case 6:
	// 	w.Population = "M"
	// 	w.PopulationNum = roll1dX(10, 0)
	// 	w.PopCode = 6
	// case 7:
	// 	w.Population = "M"
	// 	w.PopulationNum = roll1dX(40, 10)
	// 	w.PopCode = 7
	// case 8:
	// 	w.Population = "M"
	// 	w.PopulationNum = roll1dX(90, 10)
	// 	w.PopCode = 7
	// case 9:
	// 	w.Population = "M"
	// 	w.PopulationNum = roll1dX(400, 100)
	// 	w.PopCode = 8
	// case 10:
	// 	w.Population = "M"
	// 	w.PopulationNum = roll1dX(900, 100)
	// 	w.PopCode = 8
	// case 11:
	// 	w.Population = "B"
	// 	w.PopulationNum = roll1dX(10, 0)
	// 	w.PopCode = 9
	case 12:
		w.Population = "Alien"
	default:
		w.PopCode = utils.RollDice("2d6", -2)
		if w.PopCode < 3 {
			w.PopCode = 3
		}
		w.PopulationNum = 10
		for i := 1; i < w.PopCode; i++ {
			w.PopulationNum = w.PopulationNum * 10
		}
		w.PopulationNum = w.PopulationNum * utils.RollDice("2d6") / 10
		pops := w.PopulationNum
		w.PopCode = 0
		for pops > 0 {
			w.PopCode++
			pops = pops / 10
		}
	}

}

func (w *World) populationStr() string {
	pops := w.PopulationNum
	suffix := ""
	if pops/1000 > 0 {
		suffix = "K"
	}
	if pops/1000000 > 0 {
		suffix = "M"
	}
	if pops/1000000000 > 0 {
		suffix = "B"
	}
	if pops/1000000000000 > 0 {
		suffix = "T"
	}
	showNum := 0
	if suffix == "K" {
		showNum = pops / 1000
	}
	if suffix == "M" {
		showNum = pops / 1000000
	}
	if suffix == "B" {
		showNum = pops / 1000000000
	}
	return strconv.Itoa(showNum) + suffix
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

func (w *World) TechLevelInt() int {
	switch w.TechLevel {
	case "TL0":
		return 0
	case "TL1":
		return 1
	case "TL2":
		return 2
	case "TL3":
		return 3
	case "TL3+":
		return 3
	case "TL4-":
		return 4
	case "TL4":
		return 4
	case "TL4+":
		return 4
	case "TL5":
		return 5
	}
	return -1
}

func (w *World) rollTechLevel() {
	r := rollXdY(2, 6)
	w.pickTechlevel(r)
}

func (w *World) friction() int {
	if w.LawCode < 7 {
		addon := 7 - w.LawCode
		return 2 + addon

	}
	if w.LawCode > 7 {
		return w.LawCode - 5
	}
	return 2
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

func tradeTags() []string {
	tagList := []string{
		"Alien",
		"Closed",
		"Communist",
		"Disorganized",
		"Dying",
		"Fractious",
		"Kleptocratic",
		"Laissez Faire",
		"Megacorps",
		"Military",
		"Opened",
		"Panopticon",
		"Primitive",
		"Restricted",
		"Scarcity",
		"Secret",
		"Sophisticated",
		"Theocratic",
		"Thriving",
		"Tribute",
		"Tyrannical",
		"Usurped",
		"Vendor",
		"Xenophobic",
	}
	return tagList
}

func pickTradeTag(index int) string {
	return tradeTags()[index]
}

func randomTradeTag() string {
	r := roll1dX(len(tradeTags()), -1)
	return pickTradeTag(r)
}

func (w *World) planetaryTax() int {
	oneManTax := 1
	switch w.LivingStandard {
	case "Slum":
		oneManTax = 1
	case "Poor":
		oneManTax = 2
	case "Common":
		oneManTax = 3
	case "Good":
		oneManTax = 5
	case "Elite":
		oneManTax = 40
	case "Peerless":
		oneManTax = 200
	}
	totalTax := oneManTax * w.PopulationNum
	return totalTax
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

func (w *World) adjustBP() {

	if w.TechLevel == "TL3" || w.TechLevel == "TL3+" {
		w.BP = w.BP / 4
	}
	if w.TechLevel == "TL2" {
		w.BP = w.BP / 8
	}
	if w.TechLevel == "TL1" || w.TechLevel == "TL0" {
		w.BP = w.BP / 16
	}
}

func lifestyleCode(lstyle string) int {
	switch lstyle {
	case "Slum":
		return 1
	case "Poor":
		return 2
	case "Common":
		return 3
	case "Good":
		return 4
	case "Elite":
		return 5
	case "Peerless":
		return 6
	}
	return 0
}

func (w *World) calculateBP() {
	policyMod := policyShare(w.BudgetPolicy, "Military")
	militaryTax := w.planetaryTax() * 10 / policyMod
	fmt.Println("DEBUG: TAX", militaryTax, w.planetaryTax())
	w.BP = militaryTax / 200000
}

func oneRollContact() string {
	result := ""
	interest := []string{
		"They have other possibilities; it was never crucial to them",
		"They’d suffer a significant loss if this deal goes badly",
		"This would ruin their quarterly numbers if it fails",
		"It was vital to their interests, either in volume or timing.",
	}
	r := roll1dX(len(interest), -1)
	//result = result + "How interested were they in this deal?\n"
	result = result + "   " + interest[r] + "\n"

	deal := []string{
		"They need immediate cash or the good they tried to buy",
		"They know the market will soon be flooded/starved for it",
		"A government contact told them to do it or else",
		"They know a great deal but need the goods/credits for it",
		"Their supervisor/creditors made it clear that they had to",
		"They have a deal that needs the goods/credits to complete",
	}
	r = roll1dX(len(deal), -1)
	//result = result + "Why did they want to make the deal?\n"
	result = result + "   " + deal[r] + "\n"

	money := []string{
		"Old-money family with large resources",
		"Head of a profitable native business",
		"Government official in charge of stocking this resource",
		"Retired far trader who now deals only on this world",
		"Ambitious young entrepreneur with much borrowed cash",
		"Crime boss who runs legit businesses as cover",
		"Rebel or revolutionary leader with a facade of legality",
		"Local fixer who handles “problems” for important clients",
	}
	r = roll1dX(len(money), -1)
	//result = result + "Where does their money or power come from?\n"
	result = result + "   " + money[r] + "\n"

	government := []string{
		"They’re actually government officials",
		"The local government is highly suspicious of their trade",
		"They’re from a government-disliked minority or religion",
		"They’re favored by one faction and hated by another",
		"They used to be in the ruling faction, but no more",
		"They backed the faction that recently won control",
		"They operate only so long as the officials tolerate them",
		"They hate the government and want it overthrown",
		"They’re secret agents of the government law enforcement",
		"They’re operating illegally, unknown to the local lawmen",
	}
	r = roll1dX(len(government), -1)
	//result = result + "What’s their relationship with the local government?\n"
	result = result + "   " + government[r] + "\n"

	enemy := []string{
		"A jilted lover or abandoned spouse",
		"A cheated business partner",
		"A former friend they left for dead as part of their rise",
		"A former colleague they ruined to advance",
		"A far trader they cheated in a prior deal",
		"A government official they got thrown out for corruption",
		"A crime boss who they failed to pay off for protection",
		"An official who wants to show far trade as unprofitable",
		"A local industrialist who wants to sabotage far trade",
		"A local rival in the same trade who wants them ruined",
		"An employee they’ve abused and mistreated",
		"A family member who wants to control the business",
	}
	r = roll1dX(len(enemy), -1)
	//result = result + "Who would want them to suffer a loss?\n"
	result = result + "   " + enemy[r] + "\n"

	rumor := []string{
		"They’re hopelessly in debt to a local crime boss",
		"They’re actually with the hated secret police",
		"They’re secretly of a sect outlawed and persecuted here",
		"They trade cheap and shoddy goods to far traders",
		"They love to cheat far traders with native cultural dodges",
		"They trade in goods that are toxic or dangerous to buyers",
		"They’re just a puppet for a secret mastermind",
		"They’re addicted to some offworlder drug",
		"They have truly detestable sexual appetites",
		"They cheat their employees of their rightful wages",
		"The port longshoremen hate them and spoil their goods",
		"They’re secretly in league with rebels or revolutionaries",
		"They’re violently abusive toward a local minority",
		"The local government wants to ruin them for some slight",
		"They carry a socially detested disease",
		"They’re actually a very coherent feral psychic",
		"They’re somehow involved in forbidden maltech use",
		"They actually stole their wealth and have enemies for it",
		"They abuse their servants and minions in awful ways",
		"They’re building up credibility for a huge betrayal",
	}
	r = roll1dX(len(rumor), -1)
	//result = result + "What’s the most problematic rumor about them?\n"
	result = result + "   " + rumor[r] + "\n"
	return result
}

func oneRollSociety() string {
	result := ""

	howOld := []string{
		"An ancient First Wave colony",
		"Founded during the Second Wave",
		"Founded sometime around the Scream",
		"Founded within the past century",
	}
	r := roll1dX(len(howOld), -1)
	result = result + "This colony is " + howOld[r] + ". "

	reason := []string{
		"Castaways",
		"Corporate factory world",
		"Ethnic or national purity",
		"Excavation site",
		"Exiles from Old Terra or a losing regime",
		"Exotic genotype designed for here",
		"Homeworld overpopulation",
		"Invasion force",
		"Liaison outpost",
		"Mandate malcontents",
		"Military outpost",
		"Political liberty",
		"Precious export",
		"Prison planet",
		"Refueling outpost",
		"Religious liberty",
		"Research outpost",
		"Rich natural resources",
		"Social liberty",
		"Trade hub",
	}
	r = roll1dX(len(reason), -1)
	result = result + "Initial reason for colonization of this planet was: '" + reason[r] + "'. "

	prior := []string{
		"This culture has persisted since founding",
		"The culture’s changed, but has continuity",
		"Founding group splintered, this is one heir",
		"Founding group collapsed and became this",
		"Founding group wiped out by new colonists",
		"there was a several founding groups before this one",
	}
	r = roll1dX(len(prior), -1)
	result = result + "Right now it is safe to say that " + prior[r] + ". "

	modern := []string{
		"There is no other meaningful group here",
		"The society has significant sub-groups",
		"The society is a unified world government",
		"The only rivals have been conquered here",
		"There are several minor rival nations here",
		"Theres at least one major planetary rival",
		"Alliances exist of semi-equal rival nations",
		"There are dozens of significant societies",
	}
	r = roll1dX(len(modern), -1)
	result = result + "Modern political situation looks like " + modern[r] + ". "

	resource := []string{
		"Critical resources for making spike drives",
		"Nexus of interstellar spike drill routes",
		"Valuable alien tech relics or remnants",
		"Abundant food resources",
		"Important medical compound or extract",
		"Industry salvaged from a prior colony try",
		"Friendly alien population",
		"Valuable raw resources for luxury goods",
		"Local environment augments humans here",
		"Only semi-safe habitable place in system",
		"Raw materials to maintain TL5 industry",
		"Important Mandate naval base site",
	}
	r = roll1dX(len(resource), -1)
	result = result + "It is important to note that this planet benefits from " + resource[r] + ". "

	remnants := []string{
		"Vital advanced tech was left behind",
		"Ritually-important religious centers",
		"Dangerous ruins of high-security areas",
		"Abandoned cities, now dangerous to enter",
		"Bunker-caches of stored valuables and tech",
		"Massive megastructures of strange purpose",
		"Terraforming tech that needs maintenance",
		"Ethnic group with a grudge of some kind",
		"Bloodline of former rulers, which are resentful now",
		"Ancient resource extraction facilities",
	}
	r = roll1dX(len(remnants), -1)
	result = result + remnants[r] + " connects history of this planet with Terran Mandate."
	return result
}

func oneRollRulers() string {
	result := ""

	formOfRule := []string{
		"Autocracy of a single popular ruler",
		"Corporatism among guilds/classes/corps",
		"Democracy, one sentient, one vote",
		"Feudalism, many near-free sub-rulers",
		"Hydraulic Despotism over a vital resource",
		"Military Dictatorship via martial force",
		"Monarchy, single ruler via bloodline",
		"Oligarchy of the society’s powerful elite",
		"Republic of representative delegates",
		"Technocracy of intellectual elites",
		"Theocracy by the religious leadership",
		"Tribalism without structure beyond blood",
	}
	r := roll1dX(len(formOfRule), -1)
	result = result + "Basic form of rule can be described as " + formOfRule[r] + " "

	legitimacy := []string{
		"A glorious bloodline or honored family",
		"Control of overwhelming martial force",
		"Popular support among a wide class",
		"Loyalty of a major ethnic/religious group",
		"Social compact among the ruled groups",
		"Possession of pretech artifacts",
		"Religiously-legitimated sacredness",
		"Personal merit among the ruling class",
	}
	r = roll1dX(len(legitimacy), -1)
	result = result + "which draws it's legitimacy from " + legitimacy[r] + ". "

	control := []string{
		"All aspects of life are touched by the rulers",
		"The rulers firmly control the populace",
		"The rulers control only critical elements",
		"The rulers have little control of the ruled",
	}
	r = roll1dX(len(control), -1)
	result = result + "" + control[r] + ". "

	conflRulers := []string{
		"struggle between Peripheral elites against the central power",
		"Old leadership group deposed by new one",
		"Dividing the profits of taxation or tribute",
		"Starting or stopping a current war",
		"“Reformists” with new ideas for control",
		"Sectarian religious groups struggle",
		"Expanding membership in the ruling class",
		"Enacting a major public building project",
		"Support for favored corporations/groups",
		"Dividing power and offices among them",
	}
	r = roll1dX(len(conflRulers), -1)
	result = result + "Rulling class experiencing inner conflict caused by " + conflRulers[r] + ". "

	conflRuled := []string{
		"Their taxation is intolerably high",
		"Crimes against the ruled are ignored",
		"They trample on cherished customs",
		"They hold the ruled in obvious contempt",
		"The law is designed to favor them greatly",
		"They have immunity to onerous taxes",
		"Disrespect for common religious belief",
		"They failed or are failing in a recent war",
		"They waste taxes and labor on vain things",
		"Ways to enter the class have been removed",
		"The rulers are all of a different ethnicity",
		"They have different basic moral values",
		"The rulers ignore rights when it’s useful",
		"They deposed former popular rulers",
		"State connections are vital to all success",
		"They have fine ideas that ignore public will",
		"They have removed a prized ancient right",
		"They’re seen as puppets of a hated group",
		"Their source of legitimacy is crumbling",
		"The leadership is deeply incompetent",
	}
	r = roll1dX(len(conflRuled), -1)
	result = result + "On the other hand major population think that " + conflRuled[r] + ". "

	security := []string{
		"The rulers teeter on the brink of collapse",
		"They seem likely to fall soon",
		"They’ve recently overcome a real threat",
		"They have no serious threats to their rule",
		"No alternative is currently imaginable",
		"They’ve ruled undisputed for ages",
	}
	r = roll1dX(len(security), -1)
	result = result + "Still " + security[r] + ". "

	return result
}

func oneRollRuled() string {
	result := ""

	howUniform := []string{
		"consider themselves a single group",
		"consists of two general groups or factions",
		"consists of weakly-bounded sub-groups",
		"consists of many strong factions/ethnicities/classes",
	}
	r := roll1dX(len(howUniform), -1)
	result = result + "Ruled Class " + howUniform[r] + " which "

	power := []string{
		"Strong guilds of common workers",
		"Powerful influence in the military",
		"Ability to get support from a rival power",
		"Local religion is largely on their side",
		"Their magnates have poor state relations",
		"Strong tradition of self-organized rule",
		"Keen unity in pursuit of their own interest",
		"A faction of the ruling class is their ally",
		"They control the state’s income stream",
		"The rulers dread the threat of revolt",
	}
	r = roll1dX(len(power), -1)
	result = result + "builds it's power on " + power[r] + ". "

	stability := []string{
		"Is in active insurgency or revolt",
		"Are going to revolt at any time",
		"Feels serious restiveness and regular troubles",
		"Is generally content, with patches of trouble",
		"Fells widespread contentment or submission",
		"commited, Only individual can be met resistance, if even that",
	}
	r = roll1dX(len(stability), -1)
	result = result + "General population " + stability[r] + ". "

	threar := []string{
		"a peasant uprising of discontented proles",
		"new technology embraced by bourgeois",
		"religious schism threatening the state",
		"a dangerous popular demagogue arose",
		"foreign-backed regional insurgency",
		"civil war backing a deposed ruler or exile",
		"mass reluctance to support a vital war",
		"a new political philosophy spread widely",
	}
	r = roll1dX(len(threar), -1)
	result = result + "Their last Major Threat to the Rulers was " + threar[r] + ". "

	innerConflict := []string{
		"Harsh conflict between economic strata",
		"Substantial ethnic conflict and disunity",
		"Drastic changes in the local economy",
		"Competition to enter the ruling class",
		"Secessionist traditions or urges in a group",
		"Religious differences provoke trouble",
		"Regional identities are in conflict",
		"Formerly prosperous group is embittered",
		"New economic opportunity is fought over",
		"Foreign influence is causing conflict",
		"A forceful social reform movement spreads",
		"A mass appealing delusion is growing",
	}
	r = roll1dX(len(innerConflict), -1)
	result = result + "Inside the Ruled Class there is " + innerConflict[r] + ". "

	trend := []string{
		"The rich are mimicking the ruling style",
		"Coordinated tax evasion or smuggling",
		"Methodical bribery or suborning officials",
		"Popular agitation for war with a rival",
		"Radical pro-traditional social movement",
		"A communistic insurgency rising",
		"Fascist groups are gaining support",
		"Grows of large, restless youth population",
		"Demagogues promoting group tensions",
		"Embrace of “self-improving” tech or ways",
		"Now rising artificial group identity",
		"Serious and widespread drug addiction",
		"De-facto chattel debt slavery is spreading",
		"The rich, who oppose threatening tech advances",
		"A resentation of the sub-group as the ruler’s pet",
		"A powerful colonial urge is in the populace",
		"A spirit of decadent ennui is pervasive",
		"Spread of self-protection fraternities",
		"Internal disputes are becoming bloody",
		"Progressive loss of faith in their culture",
	}
	r = roll1dX(len(trend), -1)
	result = result + "Overall society's main trand is " + trend[r] + ". "

	return result
}

func (w *World) law() {
	w.GovCode = rollXdY(2, 6) - 7 + w.PopCode
	if w.GovCode < 0 {
		w.GovCode = 0
	}

	w.LawCode = rollXdY(2, 6) - 7 + w.GovCode
	if w.LawCode < 0 {
		w.LawCode = 0
	}
	w.setGovermentType(w.GovCode)
}

func restrictionsList() []string {
	strSl := []string{
		"Weapons",
		"Drugs",
		"Information",
		"Technology",
		"Travellers",
		"Psionics",
	}
	return strSl
}

func (w *World) setGovermentType(i int) {
	switch i {
	case 0:
		w.GovType = "None"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "None")
	case 1:
		w.GovType = "Company/corporation"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Weapons", "Drugs", "Travellers")
	case 2:
		w.GovType = "Participating democracy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Drugs")
	case 3:
		w.GovType = "Self-perpetuating oligarchy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Technology", "Weapons", "Travellers")
	case 4:
		w.GovType = "Representative democracy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Drugs", "Weapons", "Psionics")
	case 5:
		w.GovType = "Feudal technocracy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Technology", "Weapons", "Information")
	case 6:
		w.GovType = "Captive government"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Weapons", "Technology", "Travellers")
	case 7:
		w.GovType = "Balkanisation/Tribalism"
		restrictionOptions := utils.MaybeSlice(restrictionsList()...)
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, restrictionOptions...)
	case 8:
		w.GovType = "Civil service bureaucracy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Drugs", "Weapons")
	case 9:
		w.GovType = "Impersonal Bureaucracy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Technology", "Weapons", "Drugs", "Travellers", "Psionics")
	case 10:
		w.GovType = "Charismatic dictator"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "None")
	case 11:
		w.GovType = "Non-charismatic leader"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Weapons", "Technology", "Information")
	case 12:
		w.GovType = "Charismatic oligarchy"
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, "Weapons")
	case 13:
		w.GovType = "Religious dictatorship"
		restrictionOptions := utils.MaybeSlice(restrictionsList()...)
		w.RestrictionPolicy = generateDesidionPoolSafe(w.LawCode*2, restrictionOptions...)
	}
	fmt.Println(desidionPoolStr("Restrictions:", w.RestrictionPolicy))
	lawsList := restrictionsList()
	fmt.Println("Local Laws:")
	for i := range lawsList {
		fmt.Println(lawsList[i]+" restrictions:", w.describeRestriction(lawsList[i]))
	}
}

func (w *World) TotalPopulation() int {
	return w.PopulationNum
}

func RandomColonyReport() string {
	d1 := roll1dX(20, 0)
	d2 := roll1dX(12, 0)
	d3 := roll1dX(20, 0)
	return ColonyReport(d1, d2, d3)
}

func ColonyReport(d1, d2, d3 int) string {
	str := "Historical data: "
	str += developmentQ(d1) + " "
	str += colonizationReasonQ(d2) + " "
	str += colonyFailedQ(d3) + "\n"
	return str
}

func developmentQ(index int) string {
	if index < 5 {
		return "Landing Party. Nothing but the barest survival structures and raw beginnings of a colony.\n"
	}
	if index < 11 {
		return "Outpost. At least one town-sized population center and spaceport."
	}
	if index < 16 {
		return "Young colony. Multiple population centers and farms or other facilities for self-sustaining existence."
	}
	if index < 18 {
		return "Mature colony. Population centers, agricultural	zones, and some industrial facilities."
	}
	if index == 18 {
		return "Frontier world. Hundreds of thousands of residents with supporting industry."
	}
	if index == 19 {
		return "Developed world. Millions of residents and astronautic industries, if the planet had the resources."
	}
	if index == 20 {
		return "Advanced world. Hundreds of millions of citizens and multiple major cities and provinces."
	}
	return "Error"
}

func colonizationReasonQ(index int) string {
	reazonStr := []string{
		"It had a vital position in the spike drive trade lanes.",
		"It was exceptionally rich in rare minerals and ores.",
		"The climate and biosphere was unusually friendly to human inhabitants.",
		"It was barely habitable and thus perfect for convicts and exiles.",
		"Its location was of military importance in holding back aliens or dangerous humans.",
		"Local flora or fauna were abundant and commercially valuable.",
		"It was the site of some mysterious, research-worthy natural phenomenon.",
		"It was very distant from then-inhabited worlds and the locals wanted privacy.",
		"Aliens natives that may still be around were very friendly toward humans.",
		"It had the raw resources to produce a type of good that other worlds lacked.",
		"It was ordained as a new homeworld by some prophet or ideological leader.",
		"It was accidental, colonized out of desperate need or mistake rather than any choice.",
	}
	if index > len(reazonStr)-1 {
		index = roll1dX(len(reazonStr), -1)
	}
	return reazonStr[index]
}

func colonyFailedQ(index int) string {
	failReazStr := []string{
		"A virulent plague wiped out the young colony, but there is reason to believe someone has a vaccine.",
		"The locals killed each other off in an orgy of political, religious, or ethnic violence.",
		"The mineral strike or plunder opportunity that brought the original colonists played out and everyone left.",
		"Pirate raiders hit them repeatedly and drove them beneath a viable population for survival.",
		"A piece of infrastructure vital for surviving the planet's atmosphere or biosphere broke down.",
		"Some maltech weapon of mass destruction was unleashed on the planet.",
		"Hostile alien naval forces wiped them out on their way through the system.",
		"The population changed voluntarily or otherwise into something rather less than human.",
		"A psychic rampage during the Scream caused unsurvivable chaos in the colony.",
		"Solar flares or astronomic conditions rendered the world uninhabitable, but all that's over now. Honest.",
		"The colonists were slaughtered by indigenous natives, but took their assailants with them.",
		"The colony was too small for genetic viability without newcomers, and gradually dwindled away to nothing.",
		"The locals effectively committed suicide in service of some crazed ideology or religious demands",
		"A bloody anarchy destroyed the colony after they were cut off from the power that was forcing them to cooperate",
		"The planet was laid waste in the process of unseating a tyrannical ruler, whether local or an offworld warlord",
		"An unexpected phase of planetary weather or biospheric activity destroyed the unprepared colonists.",
		"The Mandate or a sector hegemon smashed the world as a rebel or hostile colony.",
		"Despair at their harsh and hopeless existence after the Scream gradually wore them to nothing",
		"They gambled on a failed space program or great public work that ended up collapsing their remaining industry.",
		"They were killed or left to die by their founding world due to politics or a change in the colony's value.",
	}
	if index > len(failReazStr)-1 {
		index = roll1dX(len(failReazStr), -1)
	}
	return failReazStr[index]
}

func policyShare(policy []string, policyType string) int {
	rang := 0
	for i := range policy {
		if policy[i] == policyType {
			rang++
		}
	}
	return rang
}

func (w *World) describeRestriction(restr string) string {
	share := policyShare(w.RestrictionPolicy, restr)
	switch restr {
	case "Weapons":
		return weaponRestriction(share)
	case "Drugs":
		return drugsRestriction(share)
	case "Information":
		return informationRestriction(share)
	case "Technology":
		return technologyRestriction(share)
	case "Travellers":
		return travallersRestriction(share)
	case "Psionics":
		return psionicRestriction(share)
	}
	return "No meaningful restrictions found"
}

func weaponRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Poison gas, explosives, undetectable weapons, WMD"
	case 2:
		return "Poison gas, explosives, undetectable weapons, WMD, Portable energy weapons (except ship-mounted weapons)"
	case 3:
		return "All Heavy and Mounted Weapons"
	case 4:
		return "All Heavy and Mounted Weapons, Light assault weapons and submachine guns"
	case 5:
		return "All Firearms except Pistol, Shotguns and Stunners. Consealable waepons Restricted. All weapons must be registered"
	case 6:
		return "All Firearms except Shotguns and Stunners. Carring weapons discouraged"
	case 7:
		return "All leathal Firearms. Carring weapons is highly discouraged"
	case 8:
		return "All Firearms, Bladed weapons and Stunners"
	default:
		return "All Weapons"
	}
	return "None"
}

func drugsRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Highly addictive and dangerous narcotics"
	case 2:
		return "Highly addictive narcotics"
	case 3:
		return "Highly addictive narcotics and Combat Drugs"
	case 4:
		return "Addictive narcotics and Combat Drugs"
	case 5:
		return "Addictive narcotics, Combat Drugs and Anagathics (longevity drugs)"
	case 6:
		return "Addictive narcotics, Combat Drugs, Anagathics, Fast and Slow drugs"
	case 7:
		return "All narcotics and Combat Drugs"
	case 8:
		return "All narcotics, Combat and Medical Drugs"
	default:
		return "All Drugs"
	}
	return "None"
}

func informationRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Unbracked AI"
	case 2:
		return "All AI"
	case 3:
		return "AI and Intrusion Programs"
	case 4:
		return "AI, Intrusion and Security Programs (combat capable Expert Programs)"
	case 5:
		return "All Expert Programs"
	case 6:
		return "All Expert Programs, Resent news from outworld"
	case 7:
		return "All Expert Programs, Library programs and unfiltered data about other worlds. Free speech curtailed."
	case 8:
		return "All Information Technology, any non-critical data from offworld, personal media"
	default:
		return "All Information Technology, any data from offworld. No free press."
	}
	return "None"
}

func technologyRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Maltech technology"
	case 2:
		return "Maltech and Alien technology"
	case 3:
		return "Maltech, Alien and Sapient technology. All TL5 combat related items"
	case 4:
		return "Maltech, Alien, Sapient and Pretech technology. All Items TL5"
	case 5:
		return "Maltech, Alien, Sapient and Pretech technology. All items TL5 and combat related TL4 items"
	case 6:
		return "Maltech, Alien, Sapient and Pretech technology. All TL4+ spaceship and combat related items"
	case 7:
		return "Maltech, Alien, Sapient, Post-tech and Pretech technology. All items TL4+"
	case 8:
		return "Maltech, Alien, Sapient, Astronautic, Post-tech and Pretech technology. All items TL3+"
	default:
		return "All non-Lowtech, All items TL3 and above"
	}
	return "None"
}

func travallersRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Visitors must contact planetary authorities by radio, landing is permitted anywhere"
	case 2:
		return "Visitors must report passenger and cargo manifest by radio, landing is permitted anywhere"
	case 3:
		return "Landing is permitted only at starport or other authorised sites"
	case 4:
		return "Landing is permitted only at starport"
	case 5:
		return "Citizen must register offworld travel. Visitors must register all buisiness"
	case 6:
		return "Visits discoureged. Excessive contact with cotizens forbidden"
	case 7:
		return "Citizen may not leave planet. Visitors may not leave starport"
	case 8:
		return "Landing permitted only to govermen agents"
	default:
		return "No offworlders permitted"
	}
	return "None"
}

func psionicRestriction(i int) string {
	switch i {
	case 0:
		return "None"
	case 1:
		return "Dangerous talents must be registered."
	case 2:
		return "Psionic powers must be registered. Use of dangerous powers forbidden."
	case 3:
		return "Psionic powers must be registered. Use of telepathy restricted to government approved psionics."
	case 4:
		return "Psionic powers must be registered. Use of telepathy, teleportation and clairvoyance restricted to government approved psionics."
	case 5:
		return "Psionic powers must be registered. Use of all psionic powers restricted to government approved psionics."
	case 6:
		return "Psionic powers must be registered. Use of all psionic powers restricted to government approved psionics. Possesion of psionic drugs banned."
	case 7:
		return "Use of psionic powers is forbidden."
	case 8:
		return "Psionic related technology is forbidden."
	default:
		return "All Psionics"
	}
	return "None"
}
