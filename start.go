package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	randomSeed()

	for i := 0; i < 1; i++ {
		w := NewWorld()

		fmt.Println(w.toString())
		r := roll1dX(len(tagList()), -1)
		fmt.Println(Story(r, w.WorldTag1, w.WorldTag2))

	}

	file, err := os.Open("tag.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file2, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file)
	allTags := tagList()
	currentTag := ""
	fileLine := 0
	fmt.Fprintf(file2, "FriendTagMap := make(map[string][]string)")
	for scanner.Scan() { // internally, it advances token based on sperator
		for i := range allTags {
			if scanner.Text() == allTags[i] {
				//fmt.Println("found tag", allTags[i], "i =", i)
				currentTag = allTags[i]
			}
		}

		//fmt.Fprintf(file2, strconv.Itoa(fileLine)+" Tag:"+currentTag+" ### "+scanner.Text()+"\n")
		//fmt.Println("check LINE:", scanner.Text())
		if fileLine == 2 {
			currentTag = strings.Replace(currentTag, " ", "", -1)
			currentTag = strings.Replace(currentTag, "/", "", -1)
			currentTag = strings.Replace(currentTag, "-", "", -1)
			str := "\n"
			tags := strings.Split(scanner.Text(), ", ")
			str = str + currentTag + "slF := []string{}\n"
			for i := range tags {
				str = str + currentTag + "slF = append(" + currentTag + "slF, '" + tags[i] + "')\n"
			}

			str = str + "FriendTagMap['" + currentTag + "'] = " + currentTag + "slF\n"

			fmt.Fprintf(file2, str)
		}

		fileLine++
		if fileLine > 7 {
			fileLine = 0
		}
	}
	fmt.Fprintf(file2, "return &FriendTagMap")
	// currentTag = ""
	//fmt.Println(oneRollContact())
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
	result = result + "How interested were they in this deal?\n"
	result = result + interest[r] + "\n"

	deal := []string{
		"They need immediate cash or the good they tried to buy",
		"They know the market will soon be flooded/starved for it",
		"A government contact told them to do it or else",
		"They know a great deal but need the goods/credits for it",
		"Their supervisor/creditors made it clear that they had to",
		"They have a deal that needs the goods/credits to complete",
	}
	r = roll1dX(len(deal), -1)
	result = result + "Why did they want to make the deal?\n"
	result = result + deal[r] + "\n"

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
	result = result + "Where does their money or power come from?\n"
	result = result + money[r] + "\n"

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
	result = result + "What’s their relationship with the local government?\n"
	result = result + government[r] + "\n"

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
	result = result + "Who would want them to suffer a loss?\n"
	result = result + enemy[r] + "\n"

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
	result = result + "What’s the most problematic rumor about them?\n"
	result = result + rumor[r] + "\n"
	return result
}
