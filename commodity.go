package main

type Commodite struct {
	costPerUnit int
	minTech     int
	tags        []string
	pricemod    int
}

func NewCommoditie() *Commodite {
	commodity := Commodite{}

	commodity.getTags()
	size := roll1dX(20, 0)
	if size == 1 {
		commodity.tags = append(commodity.tags, "Compact")
	}
	if inRange(size, 19, 20) {
		commodity.tags = append(commodity.tags, "Bulky")
	}
	rarity := roll1dX(20, 0)
	if inRange(rarity, 1, 2) {
		commodity.tags = append(commodity.tags, "Common")
	}
	if inRange(rarity, 20, 20) {
		commodity.tags = append(commodity.tags, "Rare")
	}
	commodity.pricemod = priceMod(commodity.tags)
	commodity.costPerUnit = price(commodity.pricemod)
	return &commodity
}

func (c *Commodite) getTags() {
	var tagArray []int
	tagQty := roll1dX(2, 0)
	tagArray = append(tagArray, 0)
	tagArray = append(tagArray, 0)

	for !tagsAreGood(tagArray[0], tagArray[1]) {
		tagArray[0] = roll1dX(20, 0)
		tagArray[1] = roll1dX(20, 0)
	}
	for i := 0; i < tagQty; i++ {
		c.tags = append(c.tags, pickCargoType(tagArray[i]-1))
	}
}

func tagsAreGood(tag1, tag2 int) bool {
	if tag1 == tag2 {
		return false
	}
	if tag1 == 8 && tag2 == 14 {
		return false
	}
	if tag1 == 8 && tag2 == 15 {
		return false
	}
	if tag1 == 14 && tag2 == 8 {
		return false
	}
	if tag1 == 14 && tag2 == 15 {
		return false
	}
	if tag1 == 15 && tag2 == 8 {
		return false
	}
	if tag1 == 15 && tag2 == 14 {
		return false
	}
	return true
}

func priceMod(sl []string) int {
	priceMod := 0
	for i := range sl {
		if sl[i] == "Agricultural" {
			priceMod = priceMod - 2
		}
		if sl[i] == "Alien" {
			priceMod = priceMod + 2
		}
		if sl[i] == "Astronautic" {
			priceMod = priceMod + 1
		}
		if sl[i] == "Biotech" {
			priceMod = priceMod + 1
		}
		if sl[i] == "Low Tech" {
			priceMod = priceMod - 1
		}
		if sl[i] == "Luxury" {
			priceMod = priceMod + 2
		}
		if sl[i] == "Maltech" {
			priceMod = priceMod + 4
		}
		if sl[i] == "Medical" {
			priceMod = priceMod + 2
		}
		if sl[i] == "Military" {
			priceMod = priceMod + 1
		}
		if sl[i] == "Mineral" {
			priceMod = priceMod - 1
		}
		if sl[i] == "Pretech" {
			priceMod = priceMod + 3
		}
		if sl[i] == "Sapient" {
			priceMod = priceMod + 2
		}
		if sl[i] == "Tool" {
			priceMod = priceMod + 1
		}
		if sl[i] == "Vehicle" {
			priceMod = priceMod + 1
		}
		if sl[i] == "Common" {
			priceMod = priceMod - 1
		}
		if sl[i] == "Rare" {
			priceMod = priceMod + 1
		}
	}
	return priceMod
}

func price(priceMod int) int {
	if priceMod == -5 {
		return 100
	}
	if priceMod == -4 {
		return 250
	}
	if priceMod == -3 {
		return 500
	}
	if priceMod == -2 {
		return 1000
	}
	if priceMod == -1 {
		return 2000
	}
	if priceMod == 0 {
		return 5000
	}
	if priceMod == 1 {
		return 10000
	}
	if priceMod == 2 {
		return 25000
	}
	if priceMod == 3 {
		return 50000
	}
	if priceMod == 4 {
		return 100000
	}
	if priceMod == 5 {
		return 200000
	}
	if priceMod == 6 {
		return 400000
	}
	if priceMod == 7 {
		return 800000
	}
	if priceMod == 8 {
		return 1600000
	}
	if priceMod == 9 {
		return 3200000
	}
	if priceMod == 10 {
		return 6400000
	}
	return 0
}

func pickCargoType(index int) string {
	typeList := []string{
		"Agricultural",
		"Alien",
		"Astronautic",
		"Biotech",
		"Consumer",
		"Cultural",
		"Livestock",
		"Low Tech",
		"Luxury",
		"Maltech",
		"Medical",
		"Military",
		"Mineral",
		"Postech",
		"Pretech",
		"Religious",
		"Sapient",
		"Survival",
		"Tool",
		"Vehicle",
	}
	return typeList[index]
}
