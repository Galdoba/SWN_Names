package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Galdoba/utils"
)

var atributeNames = [...]string{"Str", "Dex", "Con", "Int", "Wis", "Cha"}
var physicalAtr = [...]string{"Str", "Dex", "Con"}
var mentalAtr = [...]string{"Int", "Wis", "Cha"}

type npc struct {
	name       string
	lvl        int
	hp         int
	class      string
	attribute  map[string]int
	background string
	skill      map[string]skill
	focus      map[string]int
	sp         int
	esp        int
	//Strength, Dexterity, Constitution, Intelligence, Wisdom, and Charisma
}

type skill struct {
	name       string
	skillType  string
	skillLevel int
	spSpent    int
}

func (npc *npc) spendSP(skillName string, i int) {
	if !npc.canLearn(skillName) {
		npc.sp = npc.sp + i
		return
	}
	sk := npc.skill[skillName]
	sk.spSpent = sk.spSpent + i

	if sk.spSpent+1 > sk.skillLevel { // -1 , 36
		// if sk.skillLevel < -1 {
		// 	break
		// }

		sk.skillLevel = sk.skillLevel + 1
		sk.spSpent = sk.spSpent - sk.skillLevel // 4,4
	}

	npc.skill[skillName] = sk
}

func (npc *npc) Attribute(s string) int {
	if val, ok := npc.attribute[s]; ok {
		return val
	}
	return -999
}

func (npc *npc) hpLevelBonus() int {
	bonus := 0
	switch npc.class {
	case "Warrior":
		bonus = 2
	case "Adventurer (W-E)":
		bonus = 2
	case "Adventurer (W-P)":
		bonus = 2
	}
	if npc.focus["Die Hard"] > 0 {
		bonus += 2
	}
	bonus += atrMod(npc.Attribute("Con"))
	return bonus
}

func (npc *npc) IncreaseSkill(skill string) {
	curSkill := npc.skill[skill]
	curSkill.skillLevel = curSkill.skillLevel + 1
	npc.skill[skill] = curSkill
}

func (npc *npc) IncreaseStat(stat string) {
	npc.attribute[stat] = npc.attribute[stat] + 1
}

func (npc *npc) increaseFocus(focus string) {
	if npc.focus[focus] > 1 || npc.focus[focus] < 0 {
		fmt.Println("Error increseFocus", focus, npc.focus[focus])
	}
	npc.focus[focus] = npc.focus[focus] + 1
	fmt.Println(npc.focus[focus], "---------")
	npc.applyFocusEffects(focus)
}

func (npc *npc) applyFocusEffects(focusName string) {
	switch focusName {
	case "Alert":
		if npc.focus[focusName] == 1 {
			npc.develop("Notice")
		}
	case "Armsman":
		if npc.focus[focusName] == 1 {
			npc.develop("Stab")
		}
	case "Assassin":
		if npc.focus[focusName] == 1 {
			npc.develop("Sneak")
		}
	case "Authority":
		if npc.focus[focusName] == 1 {
			npc.develop("Lead")
		}
	case "Close Combatant":
		if npc.focus[focusName] == 1 {
			npc.develop("AnyCombat")
		}
	case "Connected":
		if npc.focus[focusName] == 1 {
			npc.develop("Connect")
		}
	case "Diplomat":
		if npc.focus[focusName] == 1 {
			npc.develop("Talk")
		}
	case "Gunslinger":
		if npc.focus[focusName] == 1 {
			npc.develop("Shoot")
		}
	case "Hacker":
		if npc.focus[focusName] == 1 {
			npc.develop("Program")
		}
	case "Healer":
		if npc.focus[focusName] == 1 {
			npc.develop("Heal")
		}
	case "Henchkeeper":
		if npc.focus[focusName] == 1 {
			npc.develop("Lead")
		}
	case "Psychic Training":
		if npc.focus[focusName] == 1 {
			npc.develop("AnyPsyhic")
		}
	case "Savage Fray":
		if npc.focus[focusName] == 1 {
			npc.develop("Stab")
		}
	case "Shocking Assault":
		if npc.focus[focusName] == 1 {
			d := utils.RollDice("d2")
			if d == 1 {
				npc.develop("Punch")
			} else {
				npc.develop("Stab")
			}
		}
	case "Sniper":
		if npc.focus[focusName] == 1 {
			npc.develop("Shoot")
		}
	case "Star Captain":
		if npc.focus[focusName] == 1 {
			npc.develop("Lead")
		}
	case "Starfarer":
		if npc.focus[focusName] == 1 {
			npc.develop("Pilot")
		}
	case "Tinker":
		if npc.focus[focusName] == 1 {
			npc.develop("Fix")
		}
	case "Unarmed Combatant":
		if npc.focus[focusName] == 1 {
			npc.develop("Punch")
		}
	case "Wanderer":
		if npc.focus[focusName] == 1 {
			npc.develop("Survive")
		}
	case "Wild Psychic Talent":
		if npc.focus[focusName] == 1 {
			npc.develop("AnyPsyhic")
		}

	}
}

func validSkills() []string {
	vs := []string{
		"Administer",
		"Connect",
		"Exert",
		"Fix",
		"Heal",
		"Know",
		"Lead",
		"Notice",
		"Perform",
		"Pilot",
		"Program",
		"Punch",
		"Shoot",
		"Sneak",
		"Stab",
		"Survive",
		"Talk",
		"Trade",
		"Work",
		"Biopsionics",
		"Metapsionics",
		"Precognition",
		"Telekinesis",
		"Telepathy",
		"Teleportation",
	}
	return vs
}

func validSkillList() []skill {
	var skillList []skill
	skillList = append(skillList, skill{"Administer", "Expert", -1, 0})
	skillList = append(skillList, skill{"Connect", "Expert", -1, 0})
	skillList = append(skillList, skill{"Exert", "Expert", -1, 0})
	skillList = append(skillList, skill{"Fix", "Expert", -1, 0})
	skillList = append(skillList, skill{"Heal", "Expert", -1, 0})
	skillList = append(skillList, skill{"Know", "Expert", -1, 0})
	skillList = append(skillList, skill{"Lead", "Expert", -1, 0})
	skillList = append(skillList, skill{"Notice", "Expert", -1, 0})
	skillList = append(skillList, skill{"Perform", "Expert", -1, 0})
	skillList = append(skillList, skill{"Pilot", "Expert", -1, 0})
	skillList = append(skillList, skill{"Program", "Expert", -1, 0})
	skillList = append(skillList, skill{"Punch", "Warrior", -1, 0})
	skillList = append(skillList, skill{"Shoot", "Warrior", -1, 0})
	skillList = append(skillList, skill{"Sneak", "Expert", -1, 0})
	skillList = append(skillList, skill{"Stab", "Warrior", -1, 0})
	skillList = append(skillList, skill{"Survive", "Expert", -1, 0})
	skillList = append(skillList, skill{"Talk", "Expert", -1, 0})
	skillList = append(skillList, skill{"Trade", "Expert", -1, 0})
	skillList = append(skillList, skill{"Work", "Expert", -1, 0})
	skillList = append(skillList, skill{"Biopsionics", "Psionic", -10, 0})
	skillList = append(skillList, skill{"Metapsionics", "Psionic", -10, 0})
	skillList = append(skillList, skill{"Precognition", "Psionic", -10, 0})
	skillList = append(skillList, skill{"Telekinesis", "Psionic", -10, 0})
	skillList = append(skillList, skill{"Telepathy", "Psionic", -10, 0})
	skillList = append(skillList, skill{"Teleportation", "Psionic", -10, 0})
	return skillList
}

func nonPsyhicSkills() []string {
	vs := []string{
		"Administer",
		"Connect",
		"Exert",
		"Fix",
		"Heal",
		"Know",
		"Lead",
		"Notice",
		"Perform",
		"Pilot",
		"Program",
		"Punch",
		"Shoot",
		"Sneak",
		"Stab",
		"Survive",
		"Talk",
		"Trade",
		"Work",
	}
	return vs
}

func psyhicSkills() []string {
	vs := []string{
		"Biopsionics",
		"Metapsionics",
		"Precognition",
		"Telekinesis",
		"Telepathy",
		"Teleportation",
	}
	return vs
}

func combatSkills() []string {
	vs := []string{
		"Punch",
		"Shoot",
		"Stab",
	}
	return vs
}

func nonCombatSkills() []string {
	vs := []string{
		"Administer",
		"Connect",
		"Exert",
		"Fix",
		"Heal",
		"Know",
		"Lead",
		"Notice",
		"Perform",
		"Pilot",
		"Program",
		"Sneak",
		"Survive",
		"Talk",
		"Trade",
		"Work",
	}
	return vs
}

func backgrounds() []string {
	bcgr := []string{
		"Barbarian",
		"Clergy",
		"Courtesan",
		"Criminal",
		"Dilettante",
		"Entertainer",
		"Merchant",
		"Noble",
		"Official",
		"Peasant",
		"Physician",
		"Pilot",
		"Politician",
		"Scholar",
		"Soldier",
		"Spacer",
		"Technician",
		"Thug",
		"Vagabond",
		"Worker",
	}
	return bcgr
}

func randomClass() string {
	classArray := []string{
		"Warrior",
		"Warrior",
		"Warrior",
		"Warrior",
		"Expert",
		"Expert",
		"Expert",
		"Expert",
		"Psyhic",
		"Adventurer (W-E)",
		"Adventurer (W-E)",
		"Adventurer (W-E)",
		"Adventurer (W-P)",
		"Adventurer (E-P)",
	}
	return classArray[rand.Intn(len(classArray))]
}

func (npc *npc) addClassAbilities() {
	class := npc.class
	npc.increaseFocus(utils.RandomFromList(nonPsyhicFocusList()))
	switch class {
	case "Warrior":
		npc.addFullWarriorBonus()
	case "Expert":
		npc.addFullExpertBonus()
	case "Psyhic":
		npc.addFullPsyhicBonus()
	case "Adventurer (W-E)":
		d := utils.RollDice("d2")
		addfocus := true
		if d == 1 {
			addfocus = false
		}
		npc.addPartialWarriorBonus(addfocus)
		npc.addPartialExpertBonus(!addfocus)

	case "Adventurer (W-P)":
		npc.addPartialWarriorBonus(true)
		npc.addPartialPsyhicBonus()
	case "Adventurer (E-P)":
		npc.addPartialExpertBonus(true)
		npc.addPartialPsyhicBonus()
	}

}

func (npc *npc) addFullWarriorBonus() {
	foc := utils.RandomFromList(warriorFocusList())
	fmt.Println("Debug: add:", foc)
	npc.increaseFocus(foc)
	// 1 gain a free level in a combat-related focus associated with your background. The GM decides if a focus qualifies if it’s an ambiguous case
	// You gain two extra maximum hit points at each character level. //DONE
}

func (npc *npc) addPartialWarriorBonus(addFocus bool) {
	if addFocus {
		foc := utils.RandomFromList(warriorFocusList())
		fmt.Println("Debug: add:", foc)
		npc.increaseFocus(foc)
	}
	// 1 gain a free level in a combat-related focus associated with your background. The GM decides if a focus qualifies if it’s an ambiguous case
	// You gain two extra maximum hit points at each character level. //DONE
}

func (npc *npc) addFullExpertBonus() {
	//npc.increaseFocus(utils.RandomFromList(expertFocusList()))
	foc := utils.RandomFromList(expertFocusList())
	fmt.Println("Debug: add:", foc)
	npc.increaseFocus(foc)
	// 1 gain a free level in a combat-related focus associated with your background. The GM decides if a focus qualifies if it’s an ambiguous case
	// You gain two extra maximum hit points at each character level. //DONE
}

func (npc *npc) addPartialExpertBonus(addFocus bool) {
	//npc.increaseFocus(utils.RandomFromList(expertFocusList()))
	if addFocus {
		foc := utils.RandomFromList(expertFocusList())
		fmt.Println("Debug: add:", foc)
		npc.increaseFocus(foc)
	}
	// 1 gain a free level in a combat-related focus associated with your background. The GM decides if a focus qualifies if it’s an ambiguous case
	// You gain two extra maximum hit points at each character level. //DONE
}

func (npc *npc) addFullPsyhicBonus() {
	psySkList := psyhicSkills()
	for i := range psySkList {
		npc.unlockPsiSkill(psySkList[i])
	}
	skill := utils.RandomFromList(psyhicSkills())
	npc.develop(skill)
	d := utils.RollDice("d2")
	if d == 1 {
		npc.develop(skill)
	} else {
		npc.develop(utils.RandomFromList(psyhicSkills()))
	}
	// TODO: Effort
}

func (npc *npc) learningOptions() []string {
	/*
		1 скилы из листа Learn
		2 пси скилы
		3 уже изученные
	*/
	learnOptions := glMap()[npc.background+"Learn"]
	learnOptions = append(learnOptions, npc.openPsySkills()...)
	learnOptions = append(learnOptions, npc.knownSkills()...)
	return learnOptions
}

func (npc *npc) openPsySkills() []string {
	var validSkills []string
	for key, val := range npc.skill {
		if val.skillLevel > -2 && val.skillType == "Psionic" {
			validSkills = append(validSkills, key)
		}
	}
	return validSkills
}

func (npc *npc) knownSkills() []string {
	var validSkills []string
	for key, val := range npc.skill {
		if val.skillLevel > -1 {
			validSkills = append(validSkills, key)
		}
	}
	return validSkills
}

func (npc *npc) unlockPsiSkill(skillName string) {
	sk := skill{skillName, "Psionic", -1, 0}
	npc.skill[skillName] = sk
}

func (npc *npc) addPartialPsyhicBonus() {
	skill := utils.RandomFromList(psyhicSkills())
	npc.unlockPsiSkill(skill)
	npc.develop(skill)
	// TODO: Effort
}

func rollAtrSet() map[string]int {
	atrSet := make(map[string]int)
	var changeAtr []string
	for i := range atributeNames {
		atrSet[atributeNames[i]] = utils.RollDice("3d6")
		share := (14 - atrSet[atributeNames[i]])
		for sh := 0; sh < share; sh++ {
			changeAtr = append(changeAtr, atributeNames[i])
		}
	}
	if len(changeAtr) > 0 {
		rollRes := utils.RollDice("d" + strconv.Itoa(len(changeAtr)))
		pick := changeAtr[rollRes-1]
		atrSet[pick] = 14
	}
	return atrSet
}

func blankSkillSet() map[string]skill {
	skills := make(map[string]skill)
	for i := range validSkillList() {
		skillName := validSkillList()[i].name
		skills[skillName] = validSkillList()[i]
	}
	return skills
}

func CreateNPC() *npc {
	npc := npc{}
	npc.name = RandomName(utils.RandomBool())
	npc.attribute = rollAtrSet()
	npc.skill = blankSkillSet()
	npc.focus = make(map[string]int)
	npc.background = utils.RandomFromList(backgrounds())
	npc.backgroundGrowth()
	npc.class = randomClass()
	npc.addClassAbilities()
	npc.develop("AnySkill")
	targLvl := utils.RollDice("2d6")
	for npc.lvl < targLvl {
		npc.levelUP()
	}
	npc.spendSP("Lead", npc.sp)
	fmt.Println(npc.report())
	return &npc
}

func (npc *npc) levelUP() {
	// получаем значение нового уровня
	npc.lvl = npc.lvl + 1
	// ролим хп для нового уровня
	npc.rollHP()
	// получаем сп за уровень
	npc.sp = npc.sp + 3
	if npc.class == "Expert" || npc.class == "Adventurer (W-E)" || npc.class == "Adventurer (E-P)" {
		npc.sp = npc.sp + 1
	}
	// вычисляем максимальные значения навыков - TODO: не надо - для каждого уровня значение будет фиксированно
	// проверяем есть ли новый фокус  - 2, 5, 7, and 10
	npc.addAddtionalFocus()
	// пересчитываем спасброски - значение считается от уровня TODO: уточнить есть ли модификаторы на спасброски
	// пересчитываем модификаторы атаки
	// пересчитываем Effort
	learnOptions := npc.learningOptions()
	npc.learn(learnOptions)

}

func (npc *npc) learn(learnOptions []string) {
	fmt.Println("npc.sp", npc.sp)
	for len(learnOptions) > 0 {
		pick := utils.RollDice("d"+strconv.Itoa(len(learnOptions)), -1)
		skillToLearn := learnOptions[pick]
		if npc.canLearn(skillToLearn) {
			npc.spendSP(skillToLearn, npc.skill[skillToLearn].skillLevel+1)
			npc.sp = npc.sp - npc.skill[skillToLearn].skillLevel
		} else {
			//a = append(a[:i], a[i+1:]...)
			learnOptions = append(learnOptions[:pick], learnOptions[pick+1:]...)
		}
	}
}

func (npc *npc) canLearn(skillName string) bool {
	canLearn := true
	if npc.skill[skillName].skillLevel+1 > maxSkillLevel(npc) {
		canLearn = canLearn && false
	}
	if npc.skill[skillName].skillLevel < -1 {
		canLearn = canLearn && false
	}
	if npc.sp < npc.skill[skillName].skillLevel+1 {
		fmt.Println("CANLEARN?", skillName, "npc.sp", npc.sp)
		canLearn = canLearn && false
	}
	return canLearn
}

func maxSkillLevel(npc *npc) int {
	if npc.lvl < 3 {
		return 1
	}
	if npc.lvl < 6 {
		return 2
	}
	if npc.lvl < 9 {
		return 3
	}
	return 4
}

func (npc *npc) addAddtionalFocus() {
	switch npc.lvl {
	case 2, 5, 7, 10:
		npc.increaseFocus(utils.RandomFromList(nonPsyhicFocusList()))
	}
}

func (npc *npc) rollHP() {
	diceQ := strconv.Itoa(npc.lvl)
	hpMod := npc.hpLevelBonus() * (npc.lvl)
	newHP := utils.RollDice(diceQ+"d6", hpMod)
	if npc.hp+1 >= newHP {
		npc.hp = npc.hp + 1
	} else {
		npc.hp = newHP
	}
}

func (npc *npc) setStartingHP() {
	npc.lvl = 1
	npc.hp = utils.RollDice("d6") + npc.hpLevelBonus()
	if npc.hp < 1 {
		npc.hp = 1
	}
}

func growthLearnPatern() []int {
	var patern []int
	for i := 0; i < 3; i++ {
		patern = append(patern, utils.RollDice("1d2")-1)
	}
	fmt.Println(patern)
	return patern
}

func (npc *npc) backgroundGrowth() {
	gPatern := growthLearnPatern()
	back := npc.background
	npc.develop(freeSkill(back))
	for i := range gPatern {
		action := "Growth"
		if gPatern[i] == 1 {
			action = "Learn"
		}
		key := back + action
		options := glMap()[key]
		res := options[rand.Intn(len(options))]
		npc.develop(res)

	}
}

func glMap() map[string][]string {
	glMap := make(map[string][]string)
	glMap["BarbarianGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["BarbarianLearn"] = []string{"AnyCombat", "Connect", "Exert", "Lead", "Notice", "Punch", "Sneak", "Survive"}
	glMap["ClergyGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["ClergyLearn"] = []string{"Administer", "Connect", "Know", "Lead", "Notice", "Perform", "Talk", "Talk"}
	glMap["CourtesanGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Physical2", "Connect", "AnySkill"}
	glMap["CourtesanLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Perform", "Survive", "Talk", "Trade"}
	glMap["CriminalGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["CriminalLearn"] = []string{"Administer", "AnyCombat", "Connect", "Notice", "Program", "Sneak", "Talk", "Trade"}
	glMap["DilettanteGrowth"] = []string{"AnyStat", "AnyStat", "AnyStat", "AnyStat", "Connect", "AnySkill"}
	glMap["DilettanteLearn"] = []string{"AnySkill", "AnySkill", "Connect", "Know", "Perform", "Pilot", "Talk", "Trade"}
	glMap["EntertainerGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Physical2", "Connect", "AnySkill"}
	glMap["EntertainerLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Perform", "Perform", "Sneak", "Talk"}
	glMap["MerchantGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["MerchantLearn"] = []string{"Administer", "AnyCombat", "Connect", "Fix", "Know", "Notice", "Trade", "Talk"}
	glMap["NobleGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["NobleLearn"] = []string{"Administer", "AnyCombat", "Connect", "Know", "Lead", "Notice", "Pilot", "Talk"}
	glMap["OfficialGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["OfficialLearn"] = []string{"Administer", "AnyCombat", "Connect", "Know", "Lead", "Notice", "Talk", "Trade"}
	glMap["PeasantGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Physical2", "Exert", "AnySkill"}
	glMap["PeasantLearn"] = []string{"Connect", "Exert", "Fix", "Notice", "Sneak", "Survive", "Trade", "Work"}
	glMap["PhysicianGrowth"] = []string{"AnyStat", "Physical2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["PhysicianLearn"] = []string{"Administer", "Connect", "Fix", "Heal", "Know", "Notice", "Talk", "Trade"}
	glMap["PilotGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Connect", "AnySkill"}
	glMap["PilotLearn"] = []string{"Connect", "Exert", "Fix", "Notice", "Pilot", "Pilot", "Shoot", "Trade"}
	glMap["PoliticianGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["PoliticianLearn"] = []string{"Administer", "Connect", "Connect", "Lead", "Notice", "Perform", "Talk", "Talk"}
	glMap["ScholarGrowth"] = []string{"AnyStat", "Mental2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["ScholarLearn"] = []string{"Administer", "Connect", "Fix", "Know", "Notice", "Perform", "Program", "Talk"}
	glMap["SoldierGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Physical2", "Exert", "AnySkill"}
	glMap["SoldierLearn"] = []string{"Administer", "AnyCombat", "Exert", "Fix", "Lead", "Notice", "Sneak", "Survive"}
	glMap["SpacerGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["SpacerLearn"] = []string{"Administer", "Connect", "Exert", "Fix", "Know", "Pilot", "Program", "Talk"}
	glMap["TechnicianGrowth"] = []string{"AnyStat", "Physical2", "Mental2", "Mental2", "Connect", "AnySkill"}
	glMap["TechnicianLearn"] = []string{"Administer", "Connect", "Exert", "Fix", "Fix", "Know", "Notice", "Pilot"}
	glMap["ThugGrowth"] = []string{"AnyStat", "Mental2", "Physical2", "Physical2", "Connect", "AnySkill"}
	glMap["ThugLearn"] = []string{"AnyCombat", "Connect", "Exert", "Notice", "Sneak", "Stab", "Survive", "Talk", "Shoot"} // thug имеет "Stab or Shoot" на шестой позиции
	glMap["VagabondGrowth"] = []string{"AnyStat", "Physical2", "Physical2", "Mental2", "Exert", "AnySkill"}
	glMap["VagabondLearn"] = []string{"AnyCombat", "Connect", "Notice", "Perform", "Pilot", "Sneak", "Survive", "Work"}
	glMap["WorkerGrowth"] = []string{"AnyStat", "AnyStat", "AnyStat", "AnyStat", "Exert", "AnySkill"}
	glMap["WorkerLearn"] = []string{"Administer", "AnySkill", "Connect", "Exert", "Fix", "Pilot", "Program", "Work"}

	return glMap
}

func freeSkill(background string) string {
	switch background {
	case "Barbarian":
		return "Survive"
	case "Clergy":
		return "Talk"
	case "Courtesan":
		return "Perform"
	case "Criminal":
		return "Sneak"
	case "Dilettante":
		return "Connect"
	case "Entertainer":
		return "Perform"
	case "Merchant":
		return "Trade"
	case "Noble":
		return "Lead"
	case "Official":
		return "Administer"
	case "Peasant":
		return "Exert"
	case "Physician":
		return "Heal"
	case "Pilot":
		return "Pilot"
	case "Politician":
		return "Talk"
	case "Scholar":
		return "Know"
	case "Soldier":
		return "AnyCombat"
	case "Spacer":
		return "Fix"
	case "Technician":
		return "Fix"
	case "Thug":
		return "AnyCombat"
	case "Vagabond":
		return "Survive"
	case "Worker":
		return "Work"
	}
	return "Free Skill Error"
}

func (npc *npc) develop(str string) {
	switch str {
	case "AnyStat":
		stat := atributeNames[rand.Intn(len(atributeNames))]
		npc.IncreaseStat(stat)
	case "Physical2":
		for i := 0; i < 2; i++ {
			stat := physicalAtr[rand.Intn(len(physicalAtr))]
			npc.IncreaseStat(stat)
		}
	case "Mental2":
		for i := 0; i < 2; i++ {
			stat := mentalAtr[rand.Intn(len(mentalAtr))]
			npc.IncreaseStat(stat)
		}
	case "AnySkill":
		skill := nonPsyhicSkills()[rand.Intn(len(nonPsyhicSkills()))]
		if npc.lvl > 0 {
			npc.learn(nonPsyhicSkills())
			return
		}
		npc.IncreaseSkill(skill)
	case "AnyPsyhic":
		skill := psyhicSkills()[rand.Intn(len(psyhicSkills()))]
		if npc.lvl > 0 {
			npc.learn(psyhicSkills())
			return
		}
		npc.IncreaseSkill(skill)
	case "AnyCombat":
		skill := combatSkills()[rand.Intn(len(combatSkills()))]
		if npc.lvl > 0 {
			npc.learn(combatSkills())

			return
		}
		npc.IncreaseSkill(skill)
	case "Str":
		npc.IncreaseStat(str)
	case "Dex":
		npc.IncreaseStat(str)
	case "Con":
		npc.IncreaseStat(str)
	case "Int":
		npc.IncreaseStat(str)
	case "Wis":
		npc.IncreaseStat(str)
	case "Cha":
		npc.IncreaseStat(str)
	default:
		npc.IncreaseSkill(str)
		fmt.Println(str, "==================================")
	}
}

func (npc *npc) developNEW(str string) {
	switch str {
	case "AnyStat":
		stat := atributeNames[rand.Intn(len(atributeNames))]
		npc.IncreaseStat(stat)
	case "Physical2":
		for i := 0; i < 2; i++ {
			stat := physicalAtr[rand.Intn(len(physicalAtr))]
			npc.IncreaseStat(stat)
		}
	case "Mental2":
		for i := 0; i < 2; i++ {
			stat := mentalAtr[rand.Intn(len(mentalAtr))]
			npc.IncreaseStat(stat)
		}
	case "AnySkill":
		skill := nonPsyhicSkills()[rand.Intn(len(nonPsyhicSkills()))]
		npc.IncreaseSkill(skill)
	case "AnyPsyhic":
		skill := psyhicSkills()[rand.Intn(len(psyhicSkills()))]
		npc.IncreaseSkill(skill)
	case "AnyCombat":
		//skill := combatSkills()[rand.Intn(len(combatSkills()))]
		npc.learn(combatSkills())
	case "Str":
		npc.IncreaseStat(str)
	case "Dex":
		npc.IncreaseStat(str)
	case "Con":
		npc.IncreaseStat(str)
	case "Int":
		npc.IncreaseStat(str)
	case "Wis":
		npc.IncreaseStat(str)
	case "Cha":
		npc.IncreaseStat(str)
	default:
		npc.IncreaseSkill(str)

	}
}

func (npc *npc) report() string {
	rep := ""
	rep += "Name: " + npc.name
	rep += "\nlvl: " + strconv.Itoa(npc.lvl)
	rep += "\nAttributes:"
	for i := range atributeNames {
		rep += "\n" + atributeNames[i] + ": " + strconv.Itoa(npc.attribute[atributeNames[i]]) + "(" + atrModS(npc.attribute[atributeNames[i]]) + ")"
	}
	rep += "\nBackground: " + npc.background
	rep += "\nClass: " + npc.class
	rep += "\nSkills:"
	for key, val := range npc.skill {
		if val.skillLevel > -1 {
			rep += "\n" + key + "-" + strconv.Itoa(npc.skill[key].skillLevel)
		}
	}
	rep += "\nFocus:"
	for key, val := range npc.focus {
		if val > 0 {
			rep += "\n" + key + ": " + strconv.Itoa(npc.focus[key])
		}
	}
	rep += "\nHP:" + strconv.Itoa(npc.hp)
	rep += "\nFree SP:" + strconv.Itoa(npc.sp)

	return rep
}

func expertFocusList() []string {
	ef := []string{
		"Alert",
		"Authority",
		"Connected",
		"Diplomat",
		"Hacker",
		"Healer",
		"Henchkeeper",
		"Specialist",
		"Star Captain",
		"Starfarer",
		"Tinker",
		"Wanderer",
		"Unique Gift",
		"Wild Psychic Talent",
	}
	return ef
}

func warriorFocusList() []string {
	wf := []string{
		"Alert",
		"Armsman",
		"Assassin",
		"Close Combatant",
		"Die Hard",
		"Gunslinger",
		"Ironhide",
		"Savage Fray",
		"Shocking Assault",
		"Sniper",
		"Unarmed Combatant",
		"Unique Gift",
		"Wanderer",
		"Wild Psychic Talent",
	}
	return wf
}

func nonPsyhicFocusList() []string {
	npf := []string{
		"Armsman",
		"Assassin",
		"Close Combatant",
		"Die Hard",
		"Gunslinger",
		"Ironhide",
		"Savage Fray",
		"Shocking Assault",
		"Sniper",
		"Unarmed Combatant",
		"Alert",
		"Authority",
		"Connected",
		"Diplomat",
		"Hacker",
		"Healer",
		"Henchkeeper",
		"Specialist",
		"Star Captain",
		"Starfarer",
		"Tinker",
		"Wanderer",
		"Unique Gift",
		"Wild Psychic Talent",
	}
	return npf
}

func psyhicFocusList() []string {
	npf := []string{
		"Psychic Training",
	}
	return npf
}

//RandFromList(slice)

//ExpertFocus
/*
Alert
Authority
Connected
Diplomat
Hacker
Healer
Henchkeeper
Specialist
Star Captain
Starfarer
Tinker
Wanderer
Unique Gift
Wild Psychic Talent
*/

//CombatFocus
/*
Armsman
Assassin
Close Combatant
Die Hard
Gunslinger
Ironhide
Savage Fray
Shocking Assault
Sniper
Unarmed Combatant
Unique Gift
Wild Psychic Talent
*/
