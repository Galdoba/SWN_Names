package main

import (
	"fmt"
)

func main() {
	randomSeed()

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

	fmt.Println("----------------------------")
	createVitalPointMilitary()

	name := RandomName(false)
	fmt.Println(name, Friend(w.WorldTag1, w.WorldTag2) )

	fmt.Println("BP CALC TEST----------------------------")
	baseBP(w)

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

func createVitalPointMilitary() {
	point := []string{
		"Elaborate fortification ",
		"Site with fine natural defenses ",
		"Ancient or new-made ruins ",
		"Newly-erected strongpoint ",
		"Recently-seized outpost ",
		"Transit chokepoint ",
	}
	vital := []string{
		"Commands a strategic transit route.",
		"Would savage any advance past it.",
		"Has supplies to support an offensive.",
		"Is a major recruit mustering point.",
		"Is an operational base for a major command.",
		"Is threatening an enemy position.",
	}
	p := roll1dX(len(point), -1)
	v := roll1dX(len(vital), -1)
	fmt.Println(point[p] + "is vital because it " + vital[v])
}
