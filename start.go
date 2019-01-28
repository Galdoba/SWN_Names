package main

import (
	"fmt"
)

func main() {
	randomSeed()

	for i := 0; i < 1; i++ {
		w := NewWorld()

		fmt.Println(w.toString())

		fmt.Println("Cultural report:")
		fmt.Println(oneRollSociety())
		fmt.Println(describeTag(w.WorldTag1))
		fmt.Println(describeTag(w.WorldTag2))
		fmt.Println("")
		fmt.Println("Political Report:")
		fmt.Println(oneRollRulers())
		fmt.Println(oneRollRuled())
		//r := roll1dX(len(worldTags()), -1)
		//fmt.Println(Story(r, w.WorldTag1, w.WorldTag2))

	}
	fmt.Println("----------------------------")
	var comSl []Commodite
	for i := 0; i < 10; i++ {
		com := NewCommoditie()
		fmt.Println(com)
		comSl = append(comSl, *com)
	}
	marketCheck()

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

// file, err := os.Open("tag.txt")
// if err != nil {
// 	log.Fatal(err)
// }
// defer file.Close()
// file2, err := os.Create("result.txt")
// if err != nil {
// 	log.Fatal("Cannot create file", err)
// }
// defer file2.Close()
// scanner := bufio.NewScanner(file)
// allTags := tagList()
// currentTag := ""
// fileLine := 0
// fmt.Fprintf(file2, "PlaceTagMap := make(map[string][]string)")
// for scanner.Scan() { // internally, it advances token based on sperator
// 	for i := range allTags {
// 		if scanner.Text() == allTags[i] {
// 			//fmt.Println("found tag", allTags[i], "i =", i)
// 			currentTag = allTags[i]
// 		}
// 	}

// 	//fmt.Fprintf(file2, strconv.Itoa(fileLine)+" Tag:"+currentTag+" ### "+scanner.Text()+"\n")
// 	//fmt.Println("check LINE:", scanner.Text())
// 	// if fileLine == 2 {
// 	// 	currentTag = strings.Replace(currentTag, " ", "", -1)
// 	// 	currentTag = strings.Replace(currentTag, "/", "", -1)
// 	// 	currentTag = strings.Replace(currentTag, "-", "", -1)
// 	// 	str := "\n"
// 	// 	tags := strings.Split(scanner.Text(), ", ")
// 	// 	str = str + currentTag + "slF := []string{}\n"
// 	// 	for i := range tags {
// 	// 		str = str + currentTag + "slF = append(" + currentTag + "slF, '" + tags[i] + "')\n"
// 	// 	}

// 	// 	str = str + "FriendTagMap['" + currentTag + "'] = " + currentTag + "slF\n"

// 	// 	fmt.Fprintf(file2, str)
// 	// }

// 	if fileLine == 6 {
// 		currentTag = strings.Replace(currentTag, " ", "", -1)
// 		currentTag = strings.Replace(currentTag, "/", "", -1)
// 		currentTag = strings.Replace(currentTag, "-", "", -1)
// 		str := "\n"
// 		tags := strings.Split(scanner.Text(), ", ")
// 		str = str + currentTag + "slP := []string{}\n"
// 		for i := range tags {
// 			str = str + currentTag + "slP = append(" + currentTag + "slP, '" + tags[i] + "')\n"
// 		}

// 		str = str + "PlaceTagMap['" + currentTag + "'] = " + currentTag + "slP\n"

// 		fmt.Fprintf(file2, str)
// 	}

// 	fileLine++
// 	if fileLine > 7 {
// 		fileLine = 0
// 	}
// }
// fmt.Fprintf(file2, "return PlaceTagMap")
// // currentTag = ""
// //fmt.Println(oneRollContact())

/*





 */
