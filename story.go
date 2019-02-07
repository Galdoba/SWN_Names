package main

import (
	"strings"
)

func Story(index int, t1 string, t2 string) string {
	enemy := Enemy(t1, t2)
	friend := Friend(t1, t2)
	thing := Thing(t1, t2)
	complication := Complication(t1, t2)
	place := Place(t1, t2)
	allStories := []string{
		"A " + enemy + " seeks to rob a " + friend + " of some precious " + thing + " that he has desired for some time.",
		"A " + thing + " has been discovered on property owned by a " + friend + ", but a " + complication + " risks its destruction.",
		"A " + complication + " suddenly hits the party while they’re out doing some innocuous activity.",
		"The players unwittingly offend or injure a " + enemy + ", incurring his or her wrath. A " + friend + " offers help in escaping the consequences.",
		"Rumor speaks of the discovery of a precious " + thing + " in a distant " + place + ". The players must get to it before a " + enemy + " does.",
		"A " + enemy + " has connections with offworld pirates or slavers, and a " + friend + " has been captured by them.",
		"A " + place + " has been seized by violent revolutionaries or rebels, and a " + friend + " is being held hostage by them.",
		"A " + friend + " is in love with someone forbidden by social convention, and the two of them need help eloping.",
		"A " + enemy + " wields tyrannical power over a " + friend + ", relying on the bribery of corrupt local officials to escape consequences.",
		"A " + friend + " has been lost in hostile wilderness, and the party must reach a " + place + " to rescue them in the teeth of a dangerous " + complication + ".",
		"A " + enemy + " has committed a grave offense against a PC or their family sometime in the past. A " + friend + " shows the party a weakness in the " + enemy + "’s defenses.",
		"The party is suddenly caught in a conflict between two warring families or political parties.",
		"The party is framed for a crime by a " + enemy + ", and must reach the sanctuary of a " + place + " before they can regroup and find the " + thing + " that will prove their innocence and their " + enemy + "’s perfidy.",
		"A " + friend + " is threatened by a tragedy of sickness, legal calamity, or public humiliation, and the only one that seems able to save them is a " + enemy + ".",
		"A natural disaster or similar " + complication + " strikes a " + place + " while the party is present, causing great loss of life and property unless the party is able to immediately respond to the injured and trapped.",
		"A " + friend + " with a young business has struck a cache of pretech, valuable minerals, or precious salvage. He needs the party to help him reach the " + place + " where the valuables are.",
		"An oppressed segment of society starts a sudden revolt in the " + place + " the party is occupying. A " + enemy + " simply lumps the party in with the rebels and tries to put the revolt down with force. A " + friend + " offers them a way to either help the rebels or clear their names.",
		"A vulnerable " + friend + " has been targeted for abduction, and has need of guards. A sudden " + complication + " makes guarding them from the " + enemy + " seeking their kidnapping much more difficult. If the " + friend + " is snatched, they must rescue them from a " + place + ".",
		"A mysterious " + place + " offers the promise of some precious " + thing + ", but access is very dangerous due to wildlife, hostile locals, or a dangerous environment.",
		"A " + enemy + " and a " + friend + " both have legal claim on a " + thing + ", and seek to undermine each other’s case. The " + enemy + " is willing to do murder if he thinks he can get away with it.",
		"A " + enemy + " seeks the death of his brother, a " + friend + ", by arranging the failure of his grav flyer or shuttlecraft in dangerous terrain while the party is coincidentally aboard. The party must survive the environment and bring proof of the crime out with them.",
		"A " + friend + " seeks to slip word to a lover, one who is also being courted by the " + friend + "’s brother, who is a " + enemy + ". A " + complication + " threatens to cause death or disgrace to the lover unless they either accept the " + enemy + "’s suit or are helped by the party.",
		"A " + enemy + " is convinced that one of the party has committed adultery with their flirtatious spouse. He means to lure them to a " + place + ", trap them, and have them killed by the dangers there.",
		"A " + enemy + " has been driven insane by exotic recreational drugs or excessive psionic torching. He fixes on a PC as being his mortal nemesis, and plots elaborate deaths, attempting to conceal his involvement amid " + complication + "s.",
		"A " + friend + " has stolen a precious " + thing + " from a " + enemy + " and fled into a dangerous, inaccessible " + place + ". The party must rescue them, and decide what to do with the " + thing + " and the outraged " + enemy + ".",
		"A " + enemy + " has realized that their brother or sister has engaged in a socially unacceptable affair with a " + friend + ", and means to kill both of them unless stopped by the party.",
		"A " + friend + " has accidentally caused the death of a family member, and wants the party to help him hide the body or fake an accidental death before his family realizes what has happened. A " + complication + " suddenly makes the task more difficult.",
		"A " + friend + " is a follower of a zealous ideologue who plans to make a violent demonstration of the righteousness of his cause, causing a social " + complication + ". The " + friend + " will surely be killed in the aftermath if not rescued or protected by the party.",
		"A " + friend + "’s sibling is to be " + place + "d in a dangerous situation they’ve got no chance of surviving. The " + friend + " takes their " + place + " at the last moment, and will almost certainly die unless the party aids them.",
		"Suicide bombers detonate an explosive, chemical, or biological weapon in a " + place + " occupied by the party where a precious " + thing + " is stored The PCs must escape before the " + place + " collapses on top of them, navigating throngs of terrified people in the process and saving the " + thing + " if possible.",
		"A " + enemy + " who controls landing permits, oxygen rations, or some other important resource has a prejudice against one or more of the party members. He demands that they bring him a " + thing + " from a dangerous " + place + " before he’ll give them the goods.",
		"A " + friend + " in a loveless marriage to a " + enemy + " seeks escape to be with their beloved, and contacts the party to snatch them from their spouse’s guards at a prearranged " + place + ".",
		"A " + friend + " seeks to elope with their lover, and contacts the party to help them meet their paramour at a remote, dangerous " + place + ". On arrival, they find that the lover is secretly a " + enemy + " desirous of their removal and merely lured them to the " + place + " to meet their doom.",
		"The party receives or finds a " + thing + " which proves the crimes of a " + enemy + "- yet a " + friend + " was complicit in the crimes, and will be punished as well if the authorities are involved. And the " + enemy + " will stop at nothing to get the item back.",
		"A " + friend + " needs to get to a " + place + " on time in order to complete a business contract, but a " + enemy + " means to delay and hinder them until it’s too late, inducing " + complication + "s to the trip.",
		"A locked pretech stasis pod has been discovered by a " + friend + ", along with directions to the hidden keycode that will open it. The " + place + " where the keycode is hidden is now owned by a " + enemy + ".",
		"A fierce schism has broken out in the local majority religion, and a " + enemy + " is making a play to take control of the local hierarchy. A " + friend + " is on the side that will lose badly if the " + enemy + " succeeds, and needs a " + thing + " to prove the other group’s error.",
		"A former " + enemy + " has been given reason to repent his treatment of a " + friend + ", and secretly commissions them to help the " + friend + " overcome a " + complication + ". A different " + enemy + " discovers the connection, and tries to paint the PCs as double agents.",
		"An alien or a human with extremely peculiar spiritual beliefs seeks to visit a " + place + " for their own reasons. A " + enemy + " of their own kind attempts to stop them before they can reach the " + place + ", and reveal the " + thing + " that was hidden there long ago.",
		"A " + friend + "’s sibling is an untrained psychic, and has been secretly using his or her powers to protect the " + friend + " from a " + enemy + ". The neural damage has finally overwhelmed their sanity, and they’ve now kidnapped the " + friend + " to keep them safe. The " + enemy + " is taking this opportunity to make sure the " + friend + " “dies at the hands of their maddened sibling”.",
		"A " + friend + " who is a skilled precognitive has just received a flash of an impending atrocity to be committed by a " + enemy + ". He or she needs the party to help them steal the " + thing + " that will prove the " + enemy + "’s plans while dodging the assassins sent to eliminate the precog.",
		"A " + friend + " who is an exotic dancer is sought by a " + enemy + " who won’t take no for an answer. The dancer is secretly a Perimeter agent attempting to infiltrate a " + place + " to destroy maltech research, and plots to use the party to help get him or her into the facility under the pretext of striking at the " + enemy + ".",
		"A young woman on an interplanetary tour needs the hire of local bodyguards. She turns out to be a trained and powerful combat psychic, but touchingly naive about local dangers, causing a social " + complication + " that threatens to get the whole group arrested.",
		"A librarian " + friend + " has discovered an antique databank with the coordinates of a long-lost pretech cache hidden in a " + place + " sacred to a long-vanished religion. The librarian is totally unsuited for danger, but necessary to decipher the obscure religious iconography needed to unlock the cache. The cache is not the anticipated " + thing + ", but something more dangerous to the finder.",
		"A fragment of orbital debris clips a shuttle on the way in, and the spaceport is seriously damaged in the crash. The player’s ship or the only vessel capable of getting them off-planet will be destroyed unless the players can organize a response to the dangerous chemical fires and radioactives contaminating the port. A " + friend + " is trapped somewhere in the control tower wreckage.",
		"A " + friend + " is allied with a reformist religious group that seeks to break the grip of the current, oppressive hierarchy. The current hierarchs have a great deal of political support with the authorities, but the commoners resent them bitterly. The " + friend + " seeks to secure a remote " + place + " as a meeting-place for the theological rebels.",
		"A microscopic black hole punctures an orbital station or starship above the world. Its interaction with the station’s artificial grav generators has thrown everything out of whack, and the station’s become a minefield of dangerously high or zero grav zones. It’s tearing itself apart, and it’s going to collapse soon. A " + enemy + " seeks to escape aboard the last lifeboat and to Hell with everyone else. Meanwhile, a " + friend + " is trying to save his engineer daughter from the radioactive, grav-unstable engine rooms.",
		"The planet has a sealed alien ruin, and a " + enemy + "-led cult who worships the vanished builders. They’re convinced that they have the secret to opening and controlling the power inside the ruins, but they’re only half-right. A " + friend + " has found evidence that shows that they’ll only devastate the planet if they meddle with the alien power planet. The party has to get inside the ruins and shut down the engines before it’s too late. Little do they realize that a few aliens survive inside, in a stasis field that will be broken by the ruin’s opening.",
		"A " + enemy + " and the group are suddenly trapped in a " + place + " during an accident or " + complication + ". They must work together to escape in time.",
		"A telepathic " + friend + " has discovered that a " + enemy + " was responsible for a recent atrocity. Telepathic evidence is useless on this world, however, and if she’s discovered to have scanned his mind she’ll be lobotomized as a ’rogue psychic’. A " + thing + " might be enough to prove his guilt, if the party can figure out how to get to it without revealing their " + friend + "’s meddling.",
		"A " + friend + " is responsible for safeguarding a " + thing + "yet the " + thing + " is suddenly proven to be a fake. The party must find the real object and the " + enemy + " who stole it or else their " + friend + " will be punished as the thief.",
		"A " + friend + " is bitten by a poisonous local animal while in a remote " + place + ". The only antidote is back at civilization, yet a " + complication + " threatens to delay the group until it is too late.",
		"A lethal plague has started among the residents of the town, but a " + complication + " is keeping aid from reaching them. A " + enemy + " is taking advantage of the panic to hawk a fake cure at ruinous prices, and a " + friend + " is taken in by him. The " + complication + " must be overcome before help can reach the town.",
		"A radical political party has started to institute pogroms against “groups hostile to the people”. A " + friend + " is among those groups, and needs to get out of town before a " + enemy + " uses the riot as cover to settle old scores.",
		"A " + enemy + " has sold the party an expensive but worthlessly flawed piece of equipment before lighting out for the back country. He and his plunder are holed up at a remote " + place + ".",
		"A concert of offworld music is being held in town, and a " + friend + " is slated to be the star performer. Reactionary elements led by a " + enemy + " plot to ruin the “corrupting noise” with sabotage that risks getting performers killed. Meanwhile, a crowd of ignorant offworlder fans have landed and are infuriating the locals.",
		"A " + enemy + " is wanted on a neighboring world for some heinous act, and a " + friend + " turns up as a bounty hunter ready to bring him in alive. This world refuses to extradite him, so the capture and retrieval has to evade local law enforcement.",
		"An unanticipated solar storm blocks communications and grounds the poorly-shielded grav vehicle that brought the group to this remote " + place + ". Then people start turning up dead; the storm has awoken a dangerous " + enemy + " beast.",
		"A " + friend + " has discovered a partially-complete schematic for an ancient pretech refinery unit that produces vast amounts of something precious on this world - water, oxygen, edible compounds, or the like. Several remote " + place + "s on the planet are indicated as having the necessary pretech spare parts required to build the device. When finally assembled, embedded self-modification software in the " + thing + " modifies itself into a pretech combat bot. The salvage from it remains very valuable.",
		"A " + complication + " ensnares the party where they are in an annoying but seemingly ordinary event. In actuality, a " + enemy + " is using it as cover to strike at a " + friend + " or " + thing + " that happens to be where the PCs are.",
		"A " + friend + " has a cranky, temperamental artificial heart installed, and the doctor who put it in is the only one who really understands how it works. The heart has recently started to stutter, but the doctor has vanished. A " + enemy + " has snatched him to fit his elite assassins with very unsafe combat mods.",
		"A local clinic is doing wonders in providing free health care to the poor. In truth, it’s a front for an offworld eugenics cult, with useful “specimens” kidnapped and shipped offworld while ’cremated remains’ are given to the family. A " + friend + " is snatched by them, but the party knows they’d have never consented to cremation as the clinic staff claim.",
		"Space pirates have cut a deal with an isolated backwoods settlement, off loading their plunder to merchants who meet them there. A " + friend + " goes home to family after a long absence, but is kidnapped or killed before they can bring back word of the dealings. Meanwhile, the party is entrusted with a valuable " + thing + " that must be brought to the " + friend + " quickly",
		"A reclusive psychiatrist is offering treatment for violent mentally ill patients at a remote " + place + ". His treatments seem to work, calming the subjects and returning them to rationality, though major memory loss is involved and some severe social clumsiness ensues. In actuality, he’s removed large portions of their brains to fit them with remote-control units slaved to an AI in his laboratory. He intends to use them as drones to acquire more “subjects”, and eventual control of the town.",
		"Vital medical supplies against an impending plague have been shipped in from offworld, but the spike drive craft that was due to deliver them misjumped, and has arrived in-system as a lifeless wreck transmitting a blind distress signal. Whoever gets there first can hold the whole planet hostage, and a " + enemy + " means to do just that.",
		"A " + friend + " has spent a substantial portion of their wealth on an ultra-modern new domicile, and invites the party to enjoy a weekend there. A " + enemy + " has hacked the house’s computer system to trap the inhabitants inside and use the automated fittings to kill them.",
		"A mud slide, hurricane, earthquake, or other form of disaster strikes a remote settlement. The party is the closest group of responders, and must rescue the locals while dealing with the unearthed, malfunctioning pretech " + thing + " that threatens to cause an even greater calamity if not safely defused.",
		"A " + friend + " has found a lost pretech installation, and needs help to loot it. By planetary law, the contents belong to the government.",
		"A " + enemy + " mistakes the party for the kind of offworlders who will murder innocents for pay- assuming they aren’t that kind, at least. He’s sloppy with the contact and unwittingly identifies himself, letting the players know that a " + friend + " will shortly die unless the " + enemy + " is stopped.",
		"A party member is identified as a prophesied savior for an oppressed faith or ethnicity. The believers obstinately refuse to believe any protestations to the contrary, and a cynical " + enemy + " in government decides the PC must die simply to prevent the risk of uprising. An equally cynical " + friend + " is determined to push the PC forward as a savior, because that’s what’s needed.",
		"Alien beasts escape from a zoo and run wild through the spectators. The panicked owner offers large rewards for recapturing them live, but some of the beasts are quite dangerous.",
		"A trained psychic is accused of going feral by a " + enemy + ". The psychic had already suffered severe neural damage before being found for training, so brain scans cannot establish physical signs of madness. The psychic seems unstable, but not violent- at least, on short acquaintance. The psychic offers a psychic PC the secrets of a unique psychic technique if they help him flee.",
		"A " + thing + " is the token of rulership on this world, and it’s gone missing. If it’s not found rapidly, the existing ruler will be deposed. Evidence left at a " + place + " suggests that a " + enemy + " has it, but extralegal means are necessary to investigate fully.",
		"Psychics are vanishing, including a " + friend + ". They’re being kidnapped by an ostensibly-rogue government researcher who is using them to research the lost psychic disciplines that helped enable pretech manufacturing, and they are being held at a remote " + place + ". The snatcher is a small-time local " + enemy + " with unnaturally ample resources.",
		"A " + friend + " desperately seeks to hide evidence of some past crime that will ruin his life should it come to light. A " + enemy + " holds the " + thing + " that proves his involvement, and blackmails him ruthlessly.",
		"A courier mistakes the party for the wrong set of offworlders, and wordlessly deposits a " + thing + " with them that implies something awful- med-frozen, child-sized human organs, for example, or a private catalog of gengineered human slaves. The courier’s boss shortly realizes the error, and this " + enemy + " tries to silence the PCs while preserving the " + place + " where his evil is enacted.",
		"A slowboat system freighter is taken over by " + enemy + " separatist terrorists at the same time as the planet’s space defenses are taken offline by internal terrorist attacks. The freighter is aimed straight at the starport, and will crash into it in hours if not stopped.",
		"Alien artifacts on the planet’s surface start beaming signals into the system’s asteroid belt. The signals provoke a social " + complication + " in panicked response, and a " + enemy + " seeks to use the confusion to take over. The actual effect of the signals might be harmless, or might summon a long-lost alien AI warship to scourge life from the world.",
		"An alien ambassador " + friend + " is targeted by xenophobe " + enemy + " assassins. Relations are so fragile that if the ambassador even realizes that humans are making a serious effort to kill him, the result may be war.",
		"A new religion is being preached by a " + friend + " on this planet. Existing faiths are not amused, and a " + enemy + " among the hierarchy is provoking the people to persecute the new believers, hoping for " + thing + "s to get out of hand.",
		"A " + enemy + " was once the patron of a " + friend + " until the latter was betrayed. Now the " + friend + " wants revenge, and they think they have the information necessary to get past the " + enemy + "’s defenses.",
		"Vital life support or medical equipment has been sabotaged by offworlders or zealots, and must be repaired before time runs out. The only possible source of parts is at a " + place + ", and the saboteurs can be expected to be working hard to get there and destroy them, too.",
		"A " + friend + " is importing offworld tech that threatens to completely replace the offerings of a " + enemy + " businessman. The " + enemy + " seeks to sabotage the " + friend + "’s stock, and thus ’prove’ its inferiority.",
		"An Exchange diplomat is negotiating for the opening of a branch of the interstellar bank on this world. A " + enemy + " among the local banks wants to show the world as being ungovernably unstable, so provokes " + complication + "s and riots around the diplomat.",
		"A " + enemy + " is infuriated by the uppity presumptio of an ambitious " + friend + " of a lower social caste, and tries to pin a local " + complication + " on the results of his unnatural rejection of his proper " + place + ".",
		"A " + friend + " is working for an offworld corporation to open a manufactory, and is ignoring local traditions that privilege certain social or ethnic groups, giving jobs to the most qualified workers instead. An angry " + enemy + " seeks to sabotage the factory.",
		"An offworld musician who was revered as little less than a god on his homeworld requires bodyguards. He immediately acquires Enemies on this world with his riotous ways, and his guards must keep him from getting arrested if they are to be paid.",
		"Atmospheric disturbances, dust storms, or other particulate clouds suddenly blow into town, leaving the settlement blind. A " + enemy + " commits a murder during the darkness, and attempts to frame the players as convenient scapegoats.",
		"A " + enemy + " spikes the oxygen supply of an orbital station or unbreathable-atmosphere hab dome with hallucinogens as cover for a theft. Most victims are merely confused and disoriented, but some become violent in their delusions. By chance, the party’s air supply was not contaminated.",
		"By coincidence, one of the party members is wearing clothing indicative of membership in a violent political group, and thus the party is treated in " + friend + "ly fashion by a local " + enemy + " for no obvious reason. The " + enemy + " assumes that the party will go along with some vicious crime without complaint, and the group isn’t informed of what’s in the offing until they’re in deep.",
		"A local ruler wishes outworlders to advise him of the quality of his execrable poetry- and is the sort to react very poorly to anything less that evidently sincere and fulsome praise. Failure to amuse the ruler results in the party being dumped in a dangerous " + place + " to “experience truly poetic solitude”.",
		"A " + friend + " among the locals is unreasonably convinced that offworlder tech can repair anything, and has blithely promised a powerful local " + enemy + " that the party can easily fix a damaged pretech " + thing + ". The " + enemy + " has invested in many expensive spare parts, but the truly necessary pieces are kept in a still-dangerous pretech installation in a remote " + place + ".",
		"The party’s offworld comm gear picks up a chance transmission from the local government and automatically descrambles the primitive encryption key. The document is proof that a " + enemy + " in government intends to commit an atrocity against a local village with a group of “deniable” renegades in order to steal a " + thing + " kept in the village.",
		"A " + friend + " belongs to a persecuted faith, ethnicity, or social class, and appeals for the PCs to help a cell of rebels get offworld before the " + enemy + " law enforcement finds them.",
		"A part on the party’s ship or the only other transport out has failed, and needs immediate replacement. The only available part is held by a " + enemy + ", who will only willingly relinquish it in exchange for a " + thing + " held by an innocent " + friend + " who will refuse to sell at any price.",
		"Eugenics cultists are making gengineered slaves out of genetic material gathered at a local brothel. Some of the unnaturally tempting slaves are being slipped among the prostitutes as bait to infatuate powerful officials, while others are being sold under the table to less scrupulous elites.",
		"Evidence has been unearthed at a " + place + " that substantial portions of the planet are actually owned by members of an oppressed and persecuted group. The local courts have no intention of recognizing the rights, but the codes with the ownership evidence would allow someone to bypass a number of antique pretech defenses around the planetary governor’s palace. A " + friend + " wants the codes to pass to his " + friend + "s among the group’s rebels.",
		"A crop smut threatens the planet’s agriculture, promising large-scale famine. A " + friend + " finds evidence that a secret government research station in the system’s asteroid belt was conducting experiments in disease-resistant crop strains for the planet before the Silence struck and cut off communication with the station. The existing government considers it a wild goose chase, but the party might choose to help. The station has stasis-frozen samples of the crop sufficient to avert the famine, but it also has less pleasant relics….",
		"A grasping " + enemy + " in local government seizes the party’s ship for some trifling offense. The " + enemy + " wants to end offworld trade, and is trying to scare other traders away. The starship is held within a military cordon, and the " + enemy + " is confident that by the time other elements of the government countermand the order, the free traders will have been spooked off.",
		"A seemingly useless trinket purchased by a PC turns out to be the security key to a lost pretech facility. It was sold by accident by a bungling and now-dead minion of a local " + enemy + ", who is hot after the party to “reclaim” his property… preferably after the party defeats whatever automatic defenses and bots the facility might still support.",
	}
	// fmt.Println("Enemy -", enemy)
	// fmt.Println("Friend -", friend)
	// fmt.Println("Thing -", thing)
	// fmt.Println("Complication -", complication)
	// fmt.Println("Place -", place)
	// fmt.Println("Story index -", index)
	return allStories[index]
}



func Enemy(t1, t2 string) string {
	var enemySlice []string
	enemyMap := createEnemiesMap()
	tag1 := fixPlanetTag(t1)
	tag2 := fixPlanetTag(t2)
	enemySlice = append(enemyMap[tag1], enemyMap[tag2]...)
	r := roll1dX(len(enemySlice), -1)
	//fmt.Println("Enemy index -", r, "of", len(enemySlice), enemySlice)
	return enemySlice[r]
}

func TradeAuthorities(tag string) string {
	var enemySlice []string
	enemyMap := createTradeAuthorityMap()
	ftag := fixPlanetTag(tag)
	enemySlice = enemyMap[ftag]
	r := roll1dX(len(enemySlice), -1)
	return enemySlice[r]
}

func Friend(t1, t2 string) string {
	var friendSlice []string
	friendMap := createFriendsMap()
	tag1 := fixPlanetTag(t1)
	tag2 := fixPlanetTag(t2)
	friendSlice = append(friendMap[tag1], friendMap[tag2]...)
	r := roll1dX(len(friendSlice), -1)
	return friendSlice[r]
}

func TradeAntagonists(tag string) string {
	var antagonSlice []string
	antagonMap := createTradeAntogonistMap()
	ftag := fixPlanetTag(tag)
	antagonSlice = antagonMap[ftag]
	r := roll1dX(len(antagonSlice), -1)
	return antagonSlice[r]
}

func Complication(t1, t2 string) string {
	var complicationSlice []string
	complicationMap := createComplicationsMap()
	tag1 := fixPlanetTag(t1)
	tag2 := fixPlanetTag(t2)
	complicationSlice = append(complicationMap[tag1], complicationMap[tag2]...)
	r := roll1dX(len(complicationSlice), -1)
	return complicationSlice[r]
}

func TradeComplication(tag string) string {
	var complSlice []string
	complMap := createTradeComplicationsMap()
	ftag := fixPlanetTag(tag)
	complSlice = complMap[ftag]
	r := roll1dX(len(complSlice), -1)
	return complSlice[r]
}

func Thing(t1, t2 string) string {
	var thingSlice []string
	thingMap := createThingsMap()
	tag1 := fixPlanetTag(t1)
	tag2 := fixPlanetTag(t2)
	thingSlice = append(thingMap[tag1], thingMap[tag2]...)
	r := roll1dX(len(thingSlice), -1)
	return thingSlice[r]
}

func TradeThing(tag string) string {
	var thingSlice []string
	thingMap := createTradeThingsMap()
	ftag := fixPlanetTag(tag)
	thingSlice = thingMap[ftag]
	r := roll1dX(len(thingSlice), -1)
	return thingSlice[r]
}

func Place(t1, t2 string) string {
	var placeSlice []string
	placeMap := createPlacesMap()
	tag1 := fixPlanetTag(t1)
	tag2 := fixPlanetTag(t2)
	placeSlice = append(placeMap[tag1], placeMap[tag2]...)
	r := roll1dX(len(placeSlice), -1)
	return placeSlice[r]
}

func TradeRegulation(tag string) string {
	var regulationsSlice []string
	regulationsMap := createTradeThingsMap()
	ftag := fixPlanetTag(tag)
	regulationsSlice = regulationsMap[ftag]
	r := roll1dX(len(regulationsSlice), -1)
	return regulationsSlice[r]
}

func fixPlanetTag(currentTag string) string {
	currentTag = strings.Replace(currentTag, " ", "", -1)
	currentTag = strings.Replace(currentTag, "/", "", -1)
	currentTag = strings.Replace(currentTag, "-", "", -1)
	return currentTag
}

func createFriendsMap() map[string][]string {
	FriendsTagMap := make(map[string][]string)
	AbandonedColonyslF := []string{}
	AbandonedColonyslF = append(AbandonedColonyslF, "Inquisitive stellar archaeologist")
	AbandonedColonyslF = append(AbandonedColonyslF, "Heir to the colony’s property")
	AbandonedColonyslF = append(AbandonedColonyslF, "Local wanting the place cleaned out and made safe")
	FriendsTagMap["AbandonedColony"] = AbandonedColonyslF

	AlienRuinsslF := []string{}
	AlienRuinsslF = append(AlienRuinsslF, "Curious scholar")
	AlienRuinsslF = append(AlienRuinsslF, "Avaricious local resident")
	AlienRuinsslF = append(AlienRuinsslF, "Interstellar smuggler")
	FriendsTagMap["AlienRuins"] = AlienRuinsslF

	AlteredHumanityslF := []string{}
	AlteredHumanityslF = append(AlteredHumanityslF, "Local seeking a “cure”")
	AlteredHumanityslF = append(AlteredHumanityslF, "Curious xenophiliac")
	AlteredHumanityslF = append(AlteredHumanityslF, "Anthropological researcher")
	FriendsTagMap["AlteredHumanity"] = AlteredHumanityslF

	AnarchistsslF := []string{}
	AnarchistsslF = append(AnarchistsslF, "Proud missionary for anarchy")
	AnarchistsslF = append(AnarchistsslF, "Casual local free spirit")
	AnarchistsslF = append(AnarchistsslF, "Curious offworlder political scientist")
	FriendsTagMap["Anarchists"] = AnarchistsslF

	AnthropomorphsslF := []string{}
	AnthropomorphsslF = append(AnthropomorphsslF, "Fascinated genetic researcher")
	AnthropomorphsslF = append(AnthropomorphsslF, "Diplomat trained to deal with normals")
	AnthropomorphsslF = append(AnthropomorphsslF, "Local needing outside help")
	FriendsTagMap["Anthropomorphs"] = AnthropomorphsslF

	Area51slF := []string{}
	Area51slF = append(Area51slF, "Crusading offworld investigator")
	Area51slF = append(Area51slF, "Conspiracy-theorist local")
	Area51slF = append(Area51slF, "Idealistic government reformer")
	FriendsTagMap["Area51"] = Area51slF

	BadlandsWorldslF := []string{}
	BadlandsWorldslF = append(BadlandsWorldslF, "Native desperately wishing to escape the world")
	BadlandsWorldslF = append(BadlandsWorldslF, "Scientist researching ecological repair methods")
	BadlandsWorldslF = append(BadlandsWorldslF, "Ruin scavenger")
	FriendsTagMap["BadlandsWorld"] = BadlandsWorldslF

	BattlegroundslF := []string{}
	BattlegroundslF = append(BattlegroundslF, "Native desperately seeking protection")
	BattlegroundslF = append(BattlegroundslF, "Pragmatic military officer")
	BattlegroundslF = append(BattlegroundslF, "Hapless war orphan")
	FriendsTagMap["Battleground"] = BattlegroundslF

	BeastmastersslF := []string{}
	BeastmastersslF = append(BeastmastersslF, "Native bonded with an adorable animal")
	BeastmastersslF = append(BeastmastersslF, "Herder of very useful beasts")
	BeastmastersslF = append(BeastmastersslF, "Animal-revering mystic")
	FriendsTagMap["Beastmasters"] = BeastmastersslF

	BubbleCitiesslF := []string{}
	BubbleCitiesslF = append(BubbleCitiesslF, "Local rebel against the city officials")
	BubbleCitiesslF = append(BubbleCitiesslF, "Maintenance chief in need of help")
	BubbleCitiesslF = append(BubbleCitiesslF, "Surveyor seeking new building sites")
	FriendsTagMap["BubbleCities"] = BubbleCitiesslF

	CheapLifeslF := []string{}
	CheapLifeslF = append(CheapLifeslF, "Endearing local whose life the PCs accidentally bought")
	CheapLifeslF = append(CheapLifeslF, "Escapee from death seeking outside help")
	CheapLifeslF = append(CheapLifeslF, "Reformer trying to change local mores")
	FriendsTagMap["CheapLife"] = CheapLifeslF

	CivilWarslF := []string{}
	CivilWarslF = append(CivilWarslF, "Faction loyalist seeking aid")
	CivilWarslF = append(CivilWarslF, "Native caught in the crossfire")
	CivilWarslF = append(CivilWarslF, "Offworlder seeking passage off the planet")
	FriendsTagMap["CivilWar"] = CivilWarslF

	ColdWarslF := []string{}
	ColdWarslF = append(ColdWarslF, "Apolitical information broker")
	ColdWarslF = append(ColdWarslF, "Spy for the other side")
	ColdWarslF = append(ColdWarslF, "Unjustly accused innocent")
	ColdWarslF = append(ColdWarslF, "“He’s a bastard")
	ColdWarslF = append(ColdWarslF, "but he’s our bastard” official")
	FriendsTagMap["ColdWar"] = ColdWarslF

	ColonizedPopulationslF := []string{}
	ColonizedPopulationslF = append(ColonizedPopulationslF, "Native resistance leader")
	ColonizedPopulationslF = append(ColonizedPopulationslF, "Colonial official seeking help")
	ColonizedPopulationslF = append(ColonizedPopulationslF, "Native caught between the two sides")
	FriendsTagMap["ColonizedPopulation"] = ColonizedPopulationslF

	CulturalPowerslF := []string{}
	CulturalPowerslF = append(CulturalPowerslF, "Struggling young artist")
	CulturalPowerslF = append(CulturalPowerslF, "Pupil of the artistic tradition")
	CulturalPowerslF = append(CulturalPowerslF, "Scholar of the art")
	CulturalPowerslF = append(CulturalPowerslF, "Offworlder hating the source of corrupting alien ways")
	FriendsTagMap["CulturalPower"] = CulturalPowerslF

	CybercommunistsslF := []string{}
	CybercommunistsslF = append(CybercommunistsslF, "Idealistic planning node tech")
	CybercommunistsslF = append(CybercommunistsslF, "Cynical anti-corruption cop")
	CybercommunistsslF = append(CybercommunistsslF, "Precognitive economist")
	FriendsTagMap["Cybercommunists"] = CybercommunistsslF

	CyborgsslF := []string{}
	CyborgsslF = append(CyborgsslF, "Charity-working implant physician")
	CyborgsslF = append(CyborgsslF, "Idealistic young cyber researcher")
	CyborgsslF = append(CyborgsslF, "Avant-garde activist")
	FriendsTagMap["Cyborgs"] = CyborgsslF

	CyclicalDoomslF := []string{}
	CyclicalDoomslF = append(CyclicalDoomslF, "Harried official working to prepare")
	CyclicalDoomslF = append(CyclicalDoomslF, "Offworlder studying the cycles")
	CyclicalDoomslF = append(CyclicalDoomslF, "Local threatened by perils of the cycle’s initial stages")
	FriendsTagMap["CyclicalDoom"] = CyclicalDoomslF

	DesertWorldslF := []string{}
	DesertWorldslF = append(DesertWorldslF, "Native guide")
	DesertWorldslF = append(DesertWorldslF, "Research biologist")
	DesertWorldslF = append(DesertWorldslF, "Aspiring terraformer")
	FriendsTagMap["DesertWorld"] = DesertWorldslF

	DoomedWorldslF := []string{}
	DoomedWorldslF = append(DoomedWorldslF, "Appealing waif or family head seeking escape")
	DoomedWorldslF = append(DoomedWorldslF, "Offworld relief coordinator")
	DoomedWorldslF = append(DoomedWorldslF, "Harried law officer")
	FriendsTagMap["DoomedWorld"] = DoomedWorldslF

	DyingRaceslF := []string{}
	DyingRaceslF = append(DyingRaceslF, "One of the few youth among the population")
	DyingRaceslF = append(DyingRaceslF, "Determined and hopeful reformer")
	DyingRaceslF = append(DyingRaceslF, "Researcher seeking a new method of reproduction")
	FriendsTagMap["DyingRace"] = DyingRaceslF

	EugenicCultslF := []string{}
	EugenicCultslF = append(EugenicCultslF, "Eugenic propagandist")
	EugenicCultslF = append(EugenicCultslF, "Biotechnical investigator")
	EugenicCultslF = append(EugenicCultslF, "Local seeking revenge on cult")
	FriendsTagMap["EugenicCult"] = EugenicCultslF

	ExchangeConsulateslF := []string{}
	ExchangeConsulateslF = append(ExchangeConsulateslF, "Consul in need of offworld help")
	ExchangeConsulateslF = append(ExchangeConsulateslF, "Local banker seeking to hurt his competition")
	ExchangeConsulateslF = append(ExchangeConsulateslF, "Exchange diplomat")
	FriendsTagMap["ExchangeConsulate"] = ExchangeConsulateslF

	FallenHegemonslF := []string{}
	FallenHegemonslF = append(FallenHegemonslF, "Realistic local leader trying to hold things together")
	FallenHegemonslF = append(FallenHegemonslF, "Scholar of past glories")
	FallenHegemonslF = append(FallenHegemonslF, "Refugee from an overthrown colonial enclave")
	FriendsTagMap["FallenHegemon"] = FallenHegemonslF

	FeralWorldslF := []string{}
	FeralWorldslF = append(FeralWorldslF, "Trapped outworlder")
	FeralWorldslF = append(FeralWorldslF, "Aspiring reformer")
	FeralWorldslF = append(FeralWorldslF, "Native wanting to avoid traditional flensing")
	FriendsTagMap["FeralWorld"] = FeralWorldslF

	FlyingCitiesslF := []string{}
	FlyingCitiesslF = append(FlyingCitiesslF, "Maintenance tech in need of help")
	FlyingCitiesslF = append(FlyingCitiesslF, "City defense force pilot")
	FlyingCitiesslF = append(FlyingCitiesslF, "Meteorological researcher")
	FriendsTagMap["FlyingCities"] = FlyingCitiesslF

	ForbiddenTechslF := []string{}
	ForbiddenTechslF = append(ForbiddenTechslF, "Victim of maltech")
	ForbiddenTechslF = append(ForbiddenTechslF, "Perimeter agent")
	ForbiddenTechslF = append(ForbiddenTechslF, "Investigative reporter")
	ForbiddenTechslF = append(ForbiddenTechslF, "Conventional arms merchant")
	FriendsTagMap["ForbiddenTech"] = ForbiddenTechslF

	FormerWarriorsslF := []string{}
	FormerWarriorsslF = append(FormerWarriorsslF, "Partisan of the new peaceful ways")
	FormerWarriorsslF = append(FormerWarriorsslF, "Outsider desperate for military aid")
	FormerWarriorsslF = append(FormerWarriorsslF, "Martial genius repressed by the new dispensation")
	FriendsTagMap["FormerWarriors"] = FormerWarriorsslF

	FreakGeologyslF := []string{}
	FreakGeologyslF = append(FreakGeologyslF, "Research scientist")
	FreakGeologyslF = append(FreakGeologyslF, "Prospector")
	FreakGeologyslF = append(FreakGeologyslF, "Artist")
	FriendsTagMap["FreakGeology"] = FreakGeologyslF

	FreakWeatherslF := []string{}
	FreakWeatherslF = append(FreakWeatherslF, "Meteorological researcher")
	FreakWeatherslF = append(FreakWeatherslF, "Holodoc crew wanting shots of the weather")
	FriendsTagMap["FreakWeather"] = FreakWeatherslF

	FriendlyFoeslF := []string{}
	FriendlyFoeslF = append(FriendlyFoeslF, "Well-meaning bug-eyed monster")
	FriendlyFoeslF = append(FriendlyFoeslF, "Principled eugenics cultist")
	FriendlyFoeslF = append(FriendlyFoeslF, "Suspicious investigator")
	FriendsTagMap["FriendlyFoe"] = FriendlyFoeslF

	GoldRushslF := []string{}
	GoldRushslF = append(GoldRushslF, "Claim-jumped miner")
	GoldRushslF = append(GoldRushslF, "Native alien")
	GoldRushslF = append(GoldRushslF, "Curious tourist")
	FriendsTagMap["GoldRush"] = GoldRushslF

	GreatWorkslF := []string{}
	GreatWorkslF = append(GreatWorkslF, "Outsider studying the work")
	GreatWorkslF = append(GreatWorkslF, "Local with a more temperate attitude")
	GreatWorkslF = append(GreatWorkslF, "Supplier of work materials")
	FriendsTagMap["GreatWork"] = GreatWorkslF

	HatredslF := []string{}
	HatredslF = append(HatredslF, "Intelligence agent needing catspaws")
	HatredslF = append(HatredslF, "Holodoc producers needing “an inside look”")
	HatredslF = append(HatredslF, "Unlucky offworlder from the hated system")
	FriendsTagMap["Hatred"] = HatredslF

	HeavyIndustryslF := []string{}
	HeavyIndustryslF = append(HeavyIndustryslF, "Aspiring entrepreneur")
	HeavyIndustryslF = append(HeavyIndustryslF, "Worker union leader")
	HeavyIndustryslF = append(HeavyIndustryslF, "Ambitious inventor")
	FriendsTagMap["HeavyIndustry"] = HeavyIndustryslF

	HeavyMiningslF := []string{}
	HeavyMiningslF = append(HeavyMiningslF, "Hermit prospector")
	HeavyMiningslF = append(HeavyMiningslF, "Offworld investor")
	HeavyMiningslF = append(HeavyMiningslF, "Miner’s union representative")
	FriendsTagMap["HeavyMining"] = HeavyMiningslF

	HivemindslF := []string{}
	HivemindslF = append(HivemindslF, "A scholar studying the hivemind")
	HivemindslF = append(HivemindslF, "A person severed from the gestalt")
	HivemindslF = append(HivemindslF, "A relative of someone who has been assimilated")
	FriendsTagMap["Hivemind"] = HivemindslF

	HolyWarslF := []string{}
	HolyWarslF = append(HolyWarslF, "Desperate peacemaker")
	HolyWarslF = append(HolyWarslF, "Hard-pressed refugee of the fighting")
	HolyWarslF = append(HolyWarslF, "Peaceful religious leader who lost the internal debate")
	FriendsTagMap["HolyWar"] = HolyWarslF

	HostileBiosphereslF := []string{}
	HostileBiosphereslF = append(HostileBiosphereslF, "Xenobiologist")
	HostileBiosphereslF = append(HostileBiosphereslF, "Tourist on safari")
	HostileBiosphereslF = append(HostileBiosphereslF, "Grizzled local guide")
	FriendsTagMap["HostileBiosphere"] = HostileBiosphereslF

	HostileSpaceslF := []string{}
	HostileSpaceslF = append(HostileSpaceslF, "Astronomic researcher")
	HostileSpaceslF = append(HostileSpaceslF, "Local defense commander")
	HostileSpaceslF = append(HostileSpaceslF, "Early warning monitor agent")
	FriendsTagMap["HostileSpace"] = HostileSpaceslF

	ImmortalsslF := []string{}
	ImmortalsslF = append(ImmortalsslF, "Curious longevity researcher")
	ImmortalsslF = append(ImmortalsslF, "Thrill-seeking local,")
	FriendsTagMap["Immortals"] = ImmortalsslF

	LocalSpecialtyslF := []string{}
	LocalSpecialtyslF = append(LocalSpecialtyslF, "Spy searching for the source")
	LocalSpecialtyslF = append(LocalSpecialtyslF, "Artisan seeking protection")
	LocalSpecialtyslF = append(LocalSpecialtyslF, "Exporter with problems")
	FriendsTagMap["LocalSpecialty"] = LocalSpecialtyslF

	LocalTechslF := []string{}
	LocalTechslF = append(LocalTechslF, "Curious offworld scientist")
	LocalTechslF = append(LocalTechslF, "Eager tech buyer")
	LocalTechslF = append(LocalTechslF, "Native in need of technical help")
	FriendsTagMap["LocalTech"] = LocalTechslF

	MajorSpaceyardslF := []string{}
	MajorSpaceyardslF = append(MajorSpaceyardslF, "Captain stuck in drydock")
	MajorSpaceyardslF = append(MajorSpaceyardslF, "Maintenance chief")
	MajorSpaceyardslF = append(MajorSpaceyardslF, "Mad innovator")
	FriendsTagMap["MajorSpaceyard"] = MajorSpaceyardslF

	MandarinateslF := []string{}
	MandarinateslF = append(MandarinateslF, "Crusader for test reform")
	MandarinateslF = append(MandarinateslF, "Talented but poorly-connected graduate")
	MandarinateslF = append(MandarinateslF, "Genius who tests badly")
	FriendsTagMap["Mandarinate"] = MandarinateslF

	MandateBaseslF := []string{}
	MandateBaseslF = append(MandateBaseslF, "Idealistic do-gooder local")
	MandateBaseslF = append(MandateBaseslF, "Missionary for advanced Mandate tech")
	MandateBaseslF = append(MandateBaseslF, "Outsider seeking lost data from Mandate records")
	FriendsTagMap["MandateBase"] = MandateBaseslF

	ManeatersslF := []string{}
	ManeatersslF = append(ManeatersslF, "Sympathetic local fleeing the fork")
	ManeatersslF = append(ManeatersslF, "Escapee from a pharmaceutical rendering plant")
	ManeatersslF = append(ManeatersslF, "Outsider chosen for dinner")
	ManeatersslF = append(ManeatersslF, "Reformer seeking to break the custom or its necessity")
	FriendsTagMap["Maneaters"] = ManeatersslF

	MegacorpsslF := []string{}
	MegacorpsslF = append(MegacorpsslF, "Victim of megacorp scheming")
	MegacorpsslF = append(MegacorpsslF, "Offworlder merchant in far over their head")
	MegacorpsslF = append(MegacorpsslF, "Local reformer struggling to cope with megacorp indifference")
	FriendsTagMap["Megacorps"] = MegacorpsslF

	MercenariesslF := []string{}
	MercenariesslF = append(MercenariesslF, "Young and idealistic mercenary chief")
	MercenariesslF = append(MercenariesslF, "Harried leader of enfeebled national army")
	MercenariesslF = append(MercenariesslF, "Offworlder trying to hire help for a noble cause")
	FriendsTagMap["Mercenaries"] = MercenariesslF

	MisandryMisogynyslF := []string{}
	MisandryMisogynyslF = append(MisandryMisogynyslF, "Oppressed native")
	MisandryMisogynyslF = append(MisandryMisogynyslF, "Research scientist")
	MisandryMisogynyslF = append(MisandryMisogynyslF, "Offworld emancipationist")
	MisandryMisogynyslF = append(MisandryMisogynyslF, "Local reformer")
	FriendsTagMap["MisandryMisogyny"] = MisandryMisogynyslF

	NightWorldslF := []string{}
	NightWorldslF = append(NightWorldslF, "Curious offworlder researcher")
	NightWorldslF = append(NightWorldslF, "Hard-pressed colony leader")
	NightWorldslF = append(NightWorldslF, "High priest of a sect that finds religious significance in the night")
	FriendsTagMap["NightWorld"] = NightWorldslF

	MinimalContactslF := []string{}
	MinimalContactslF = append(MinimalContactslF, "Aspiring tourist")
	MinimalContactslF = append(MinimalContactslF, "Anthropological researcher")
	MinimalContactslF = append(MinimalContactslF, "Offworld thief")
	MinimalContactslF = append(MinimalContactslF, "Religious missionary")
	FriendsTagMap["MinimalContact"] = MinimalContactslF

	NomadsslF := []string{}
	NomadsslF = append(NomadsslF, "Free-spirited young nomad")
	NomadsslF = append(NomadsslF, "Dreamer imagining a stable life")
	NomadsslF = append(NomadsslF, "Offworlder enamored of the life")
	FriendsTagMap["Nomads"] = NomadsslF

	OceanicWorldslF := []string{}
	OceanicWorldslF = append(OceanicWorldslF, "Daredevil fisherman")
	OceanicWorldslF = append(OceanicWorldslF, "Sea hermit")
	OceanicWorldslF = append(OceanicWorldslF, "Sapient native life")
	FriendsTagMap["OceanicWorld"] = OceanicWorldslF

	OutofContactslF := []string{}
	OutofContactslF = append(OutofContactslF, "Scheming native noble")
	OutofContactslF = append(OutofContactslF, "Heretical theologian")
	OutofContactslF = append(OutofContactslF, "UFO cultist native")
	FriendsTagMap["OutofContact"] = OutofContactslF

	OutpostWorldslF := []string{}
	OutpostWorldslF = append(OutpostWorldslF, "Lonely staffer")
	OutpostWorldslF = append(OutpostWorldslF, "Fixated researcher")
	OutpostWorldslF = append(OutpostWorldslF, "Overtaxed maintenance chief")
	FriendsTagMap["OutpostWorld"] = OutpostWorldslF

	PerimeterAgencyslF := []string{}
	PerimeterAgencyslF = append(PerimeterAgencyslF, "Agent in need of help")
	PerimeterAgencyslF = append(PerimeterAgencyslF, "Support staffer")
	PerimeterAgencyslF = append(PerimeterAgencyslF, "“Unjustly” targeted researcher")
	FriendsTagMap["PerimeterAgency"] = PerimeterAgencyslF

	PilgrimageSiteslF := []string{}
	PilgrimageSiteslF = append(PilgrimageSiteslF, "Protector of the holy site")
	PilgrimageSiteslF = append(PilgrimageSiteslF, "Naive offworlder pilgrim")
	PilgrimageSiteslF = append(PilgrimageSiteslF, "Outsider wanting to learn the sanctum’s inner secrets")
	FriendsTagMap["PilgrimageSite"] = PilgrimageSiteslF

	PleasureWorldslF := []string{}
	PleasureWorldslF = append(PleasureWorldslF, "Tourist who’s in too deep")
	PleasureWorldslF = append(PleasureWorldslF, "Native seeking a more meaningful life elsewhere")
	PleasureWorldslF = append(PleasureWorldslF, "Offworld entertainer looking for training here")
	FriendsTagMap["PleasureWorld"] = PleasureWorldslF

	PoliceStateslF := []string{}
	PoliceStateslF = append(PoliceStateslF, "Rebel leader")
	PoliceStateslF = append(PoliceStateslF, "Offworld agitator")
	PoliceStateslF = append(PoliceStateslF, "Imprisoned victim")
	PoliceStateslF = append(PoliceStateslF, "Crime boss")
	FriendsTagMap["PoliceState"] = PoliceStateslF

	PostScarcityslF := []string{}
	PostScarcityslF = append(PostScarcityslF, "Offworlder seeking something available only here")
	PostScarcityslF = append(PostScarcityslF, "Local struggling to maintain the production tech")
	PostScarcityslF = append(PostScarcityslF, "Native missionary seeking to bring abundance to other worlds")
	FriendsTagMap["PostScarcity"] = PostScarcityslF

	PreceptorArchiveslF := []string{}
	PreceptorArchiveslF = append(PreceptorArchiveslF, "Preceptor Adept missionary")
	PreceptorArchiveslF = append(PreceptorArchiveslF, "Offworld scholar")
	PreceptorArchiveslF = append(PreceptorArchiveslF, "Reluctant student")
	PreceptorArchiveslF = append(PreceptorArchiveslF, "Roving Preceptor Adept")
	FriendsTagMap["PreceptorArchive"] = PreceptorArchiveslF

	PretechCultistsslF := []string{}
	PretechCultistsslF = append(PretechCultistsslF, "Offworld scientist")
	PretechCultistsslF = append(PretechCultistsslF, "Robbed collector")
	PretechCultistsslF = append(PretechCultistsslF, "Cult heretic")
	FriendsTagMap["PretechCultists"] = PretechCultistsslF

	PrimitiveAliensslF := []string{}
	PrimitiveAliensslF = append(PrimitiveAliensslF, "Colonist leader")
	PrimitiveAliensslF = append(PrimitiveAliensslF, "Peace-faction alien chief")
	PrimitiveAliensslF = append(PrimitiveAliensslF, "Planetary frontiersman")
	PrimitiveAliensslF = append(PrimitiveAliensslF, "Xenoresearcher")
	FriendsTagMap["PrimitiveAliens"] = PrimitiveAliensslF

	PrisonPlanetslF := []string{}
	PrisonPlanetslF = append(PrisonPlanetslF, "Innocent local born here")
	PrisonPlanetslF = append(PrisonPlanetslF, "Native technician forced to maintain the very tech that imprisons them")
	PrisonPlanetslF = append(PrisonPlanetslF, "Offworlder trapped here by accident")
	FriendsTagMap["PrisonPlanet"] = PrisonPlanetslF

	PsionicsAcademyslF := []string{}
	PsionicsAcademyslF = append(PsionicsAcademyslF, "Offworld researcher")
	PsionicsAcademyslF = append(PsionicsAcademyslF, "Aspiring student")
	PsionicsAcademyslF = append(PsionicsAcademyslF, "Wealthy tourist")
	FriendsTagMap["PsionicsAcademy"] = PsionicsAcademyslF

	PsionicsFearslF := []string{}
	PsionicsFearslF = append(PsionicsFearslF, "Hidden psychic")
	PsionicsFearslF = append(PsionicsFearslF, "Offworlder psychic trapped here")
	PsionicsFearslF = append(PsionicsFearslF, "Offworld educator")
	FriendsTagMap["PsionicsFear"] = PsionicsFearslF

	PsionicsWorshipslF := []string{}
	PsionicsWorshipslF = append(PsionicsWorshipslF, "Offworlder psychic researcher")
	PsionicsWorshipslF = append(PsionicsWorshipslF, "Native rebel")
	PsionicsWorshipslF = append(PsionicsWorshipslF, "Offworld employer seeking psychics")
	FriendsTagMap["PsionicsWorship"] = PsionicsWorshipslF

	QuarantinedWorldslF := []string{}
	QuarantinedWorldslF = append(QuarantinedWorldslF, "Relative of a person trapped on the world")
	QuarantinedWorldslF = append(QuarantinedWorldslF, "Humanitarian relief official")
	QuarantinedWorldslF = append(QuarantinedWorldslF, "Treasure hunter")
	FriendsTagMap["QuarantinedWorld"] = QuarantinedWorldslF

	RadioactiveWorldslF := []string{}
	RadioactiveWorldslF = append(RadioactiveWorldslF, "Reckless prospector")
	RadioactiveWorldslF = append(RadioactiveWorldslF, "Offworld scavenger")
	RadioactiveWorldslF = append(RadioactiveWorldslF, "Biogenetic variety seeker")
	FriendsTagMap["RadioactiveWorld"] = RadioactiveWorldslF

	RefugeesslF := []string{}
	RefugeesslF = append(RefugeesslF, "Sympathetic refugee waif")
	RefugeesslF = append(RefugeesslF, "Local hard-pressed by refugee gangs")
	RefugeesslF = append(RefugeesslF, "Clergy seeking peace")
	FriendsTagMap["Refugees"] = RefugeesslF

	RegionalHegemonslF := []string{}
	RegionalHegemonslF = append(RegionalHegemonslF, "Diplomat")
	RegionalHegemonslF = append(RegionalHegemonslF, "Offworld ambassador")
	RegionalHegemonslF = append(RegionalHegemonslF, "Foreign spy")
	FriendsTagMap["RegionalHegemon"] = RegionalHegemonslF

	RestrictiveLawsslF := []string{}
	RestrictiveLawsslF = append(RestrictiveLawsslF, "Frustrated offworlder")
	RestrictiveLawsslF = append(RestrictiveLawsslF, "Repressed native")
	RestrictiveLawsslF = append(RestrictiveLawsslF, "Reforming crusader")
	FriendsTagMap["RestrictiveLaws"] = RestrictiveLawsslF

	RevanchistsslF := []string{}
	RevanchistsslF = append(RevanchistsslF, "Realist local clergy seeking peace")
	RevanchistsslF = append(RevanchistsslF, "Politician trying to calm the public")
	RevanchistsslF = append(RevanchistsslF, "Third-party diplomat trying to stamp out the fire")
	FriendsTagMap["Revanchists"] = RevanchistsslF

	RevolutionariesslF := []string{}
	RevolutionariesslF = append(RevolutionariesslF, "Sympathetic victim accused of revolutionary sympathies or government collaboration")
	RevolutionariesslF = append(RevolutionariesslF, "Revolutionary or state agent who now repents")
	RevolutionariesslF = append(RevolutionariesslF, "Agent of a neutral power that wants peace")
	FriendsTagMap["Revolutionaries"] = RevolutionariesslF

	RigidCultureslF := []string{}
	RigidCultureslF = append(RigidCultureslF, "Revolutionary agitator")
	RigidCultureslF = append(RigidCultureslF, "Ambitious peasant")
	RigidCultureslF = append(RigidCultureslF, "Frustrated merchant")
	FriendsTagMap["RigidCulture"] = RigidCultureslF

	RisingHegemonslF := []string{}
	RisingHegemonslF = append(RisingHegemonslF, "Friendly emissary to the benighted")
	RisingHegemonslF = append(RisingHegemonslF, "Hardscrabble local turned great success")
	RisingHegemonslF = append(RisingHegemonslF, "Foreign visitor seeking contacts or knowledge")
	FriendsTagMap["RisingHegemon"] = RisingHegemonslF

	RitualCombatslF := []string{}
	RitualCombatslF = append(RitualCombatslF, "Peace-minded foreign missionary")
	RitualCombatslF = append(RitualCombatslF, "Temperate defender of the weak")
	RitualCombatslF = append(RitualCombatslF, "Local eager to learn of offworld fighting styles")
	FriendsTagMap["RitualCombat"] = RitualCombatslF

	RobotsslF := []string{}
	RobotsslF = append(RobotsslF, "Data-seeking robot")
	RobotsslF = append(RobotsslF, "Plucky young robot tech")
	RobotsslF = append(RobotsslF, "Local being pushed out of a job by robots")
	FriendsTagMap["Robots"] = RobotsslF

	SeagoingCitiesslF := []string{}
	SeagoingCitiesslF = append(SeagoingCitiesslF, "City navigator")
	SeagoingCitiesslF = append(SeagoingCitiesslF, "Scout captain")
	SeagoingCitiesslF = append(SeagoingCitiesslF, "Curious mer-human")
	SeagoingCitiesslF = append(SeagoingCitiesslF, "Hard-pressed ship-city engineer")
	FriendsTagMap["SeagoingCities"] = SeagoingCitiesslF

	SealedMenaceslF := []string{}
	SealedMenaceslF = append(SealedMenaceslF, "Keeper of the menace")
	SealedMenaceslF = append(SealedMenaceslF, "Student of its nature")
	SealedMenaceslF = append(SealedMenaceslF, "Victim of the menace")
	FriendsTagMap["SealedMenace"] = SealedMenaceslF

	SecretMastersslF := []string{}
	SecretMastersslF = append(SecretMastersslF, "Paranoid conspiracy theorist")
	SecretMastersslF = append(SecretMastersslF, "Machiavellian gamesman within the cabal")
	SecretMastersslF = append(SecretMastersslF, "Interstellar investigator")
	FriendsTagMap["SecretMasters"] = SecretMastersslF

	SectariansslF := []string{}
	SectariansslF = append(SectariansslF, "Reformist clergy")
	SectariansslF = append(SectariansslF, "Local peacekeeping official")
	SectariansslF = append(SectariansslF, "Offworld missionary")
	SectariansslF = append(SectariansslF, "Exhausted ruler")
	FriendsTagMap["Sectarians"] = SectariansslF

	SeismicInstabilityslF := []string{}
	SeismicInstabilityslF = append(SeismicInstabilityslF, "Experimental construction firm owner")
	SeismicInstabilityslF = append(SeismicInstabilityslF, "Adventurous volcanologist")
	SeismicInstabilityslF = append(SeismicInstabilityslF, "Geothermal prospector")
	FriendsTagMap["SeismicInstability"] = SeismicInstabilityslF

	ShackledWorldslF := []string{}
	ShackledWorldslF = append(ShackledWorldslF, "Struggling local researcher")
	ShackledWorldslF = append(ShackledWorldslF, "Offworlder trapped here")
	ShackledWorldslF = append(ShackledWorldslF, "Scientist with a plan to break the chains")
	FriendsTagMap["ShackledWorld"] = ShackledWorldslF

	SocietalDespairslF := []string{}
	SocietalDespairslF = append(SocietalDespairslF, "Struggling messenger of a new way")
	SocietalDespairslF = append(SocietalDespairslF, "Valiant paragon of a fading tradition")
	SocietalDespairslF = append(SocietalDespairslF, "Local going through the motions of serving a now-irrelevant role")
	FriendsTagMap["SocietalDespair"] = SocietalDespairslF

	SoleSupplierslF := []string{}
	SoleSupplierslF = append(SoleSupplierslF, "Doughty resource miner")
	SoleSupplierslF = append(SoleSupplierslF, "Researcher trying to synthesize the stuff")
	SoleSupplierslF = append(SoleSupplierslF, "Small-scale resource producer")
	SoleSupplierslF = append(SoleSupplierslF, "Harried starport trade overseer")
	FriendsTagMap["SoleSupplier"] = SoleSupplierslF

	TabooTreasureslF := []string{}
	TabooTreasureslF = append(TabooTreasureslF, "Reformer seeking to end its use")
	TabooTreasureslF = append(TabooTreasureslF, "Innovator trying to repurpose the treasure in innocent ways")
	TabooTreasureslF = append(TabooTreasureslF, "Wretched addict unwillingly prey to the treasure")
	FriendsTagMap["TabooTreasure"] = TabooTreasureslF

	TerraformFailureslF := []string{}
	TerraformFailureslF = append(TerraformFailureslF, "Local trying to fix the engines")
	TerraformFailureslF = append(TerraformFailureslF, "Offworlder student of the engines")
	TerraformFailureslF = append(TerraformFailureslF, "World-wise native survivor")
	FriendsTagMap["TerraformFailure"] = TerraformFailureslF

	TheocracyslF := []string{}
	TheocracyslF = append(TheocracyslF, "Heretic")
	TheocracyslF = append(TheocracyslF, "Offworld theologian")
	TheocracyslF = append(TheocracyslF, "Atheistic merchant")
	TheocracyslF = append(TheocracyslF, "Desperate commoner")
	FriendsTagMap["Theocracy"] = TheocracyslF

	TombWorldslF := []string{}
	TombWorldslF = append(TombWorldslF, "Scavenger Fleet captain")
	TombWorldslF = append(TombWorldslF, "Archaeologist")
	TombWorldslF = append(TombWorldslF, "Salvaging historian")
	TombWorldslF = append(TombWorldslF, "Xenophilic native survivor")
	FriendsTagMap["TombWorld"] = TombWorldslF

	TradeHubslF := []string{}
	TradeHubslF = append(TradeHubslF, "Rich tourist")
	TradeHubslF = append(TradeHubslF, "Hardscrabble free trader")
	TradeHubslF = append(TradeHubslF, "Merchant prince in need of catspaws")
	TradeHubslF = append(TradeHubslF, "Friendly spaceport urchin")
	FriendsTagMap["TradeHub"] = TradeHubslF

	TyrannyslF := []string{}
	TyrannyslF = append(TyrannyslF, "Conspiring rebel")
	TyrannyslF = append(TyrannyslF, "Oppressed merchant")
	TyrannyslF = append(TyrannyslF, "Desperate peasant")
	TyrannyslF = append(TyrannyslF, "Inspiring religious leader")
	FriendsTagMap["Tyranny"] = TyrannyslF

	UnbrakedAIslF := []string{}
	UnbrakedAIslF = append(UnbrakedAIslF, "Perimeter agent")
	UnbrakedAIslF = append(UnbrakedAIslF, "AI researcher")
	UnbrakedAIslF = append(UnbrakedAIslF, "Braked AI")
	FriendsTagMap["UnbrakedAI"] = UnbrakedAIslF

	UrbanizedSurfaceslF := []string{}
	UrbanizedSurfaceslF = append(UrbanizedSurfaceslF, "Local yearning for wild spaces")
	UrbanizedSurfaceslF = append(UrbanizedSurfaceslF, "Grubby urchin of the underlevels")
	UrbanizedSurfaceslF = append(UrbanizedSurfaceslF, "Harried engineer trying to maintain ancient works")
	UrbanizedSurfaceslF = append(UrbanizedSurfaceslF, "Grizzled hab cop")
	FriendsTagMap["UrbanizedSurface"] = UrbanizedSurfaceslF

	UtopiaslF := []string{}
	UtopiaslF = append(UtopiaslF, "Deranged malcontent")
	UtopiaslF = append(UtopiaslF, "Bloody-handed guerrilla leader of a rebellion of madmen")
	UtopiaslF = append(UtopiaslF, "Outsider trying to find a way to reverse the utopian changes")
	FriendsTagMap["Utopia"] = UtopiaslF

	WarlordsslF := []string{}
	WarlordsslF = append(WarlordsslF, "Vengeful commoner")
	WarlordsslF = append(WarlordsslF, "Government military officer")
	WarlordsslF = append(WarlordsslF, "Humanitarian aid official")
	WarlordsslF = append(WarlordsslF, "Village priest")
	FriendsTagMap["Warlords"] = WarlordsslF

	XenophilesslF := []string{}
	XenophilesslF = append(XenophilesslF, "Benevolent alien")
	XenophilesslF = append(XenophilesslF, "Native malcontent")
	XenophilesslF = append(XenophilesslF, "Gone-native offworlder")
	FriendsTagMap["Xenophiles"] = XenophilesslF

	XenophobesslF := []string{}
	XenophobesslF = append(XenophobesslF, "Curious native")
	XenophobesslF = append(XenophobesslF, "Exiled former ruler")
	XenophobesslF = append(XenophobesslF, "Local desperately seeking outworlder help")
	FriendsTagMap["Xenophobes"] = XenophobesslF

	ZombiesslF := []string{}
	ZombiesslF = append(ZombiesslF, "Survivor of an outbreak")
	ZombiesslF = append(ZombiesslF, "Doctor searching for a cure")
	ZombiesslF = append(ZombiesslF, "Rebel against the secret malefactors")
	FriendsTagMap["Zombies"] = ZombiesslF
	return FriendsTagMap
}

func createEnemiesMap() map[string][]string {
	EnemyTagMap := make(map[string][]string)
	AbandonedColonyslE := []string{}
	AbandonedColonyslE = append(AbandonedColonyslE, "Crazed survivors")
	AbandonedColonyslE = append(AbandonedColonyslE, "Ruthless plunderers of the ruins")
	AbandonedColonyslE = append(AbandonedColonyslE, "Automated defense system")
	EnemyTagMap["AbandonedColony"] = AbandonedColonyslE

	AlienRuinsslE := []string{}
	AlienRuinsslE = append(AlienRuinsslE, "Customs inspector")
	AlienRuinsslE = append(AlienRuinsslE, "Worshipper of the ruins")
	AlienRuinsslE = append(AlienRuinsslE, "Hidden alien survivor")
	EnemyTagMap["AlienRuins"] = AlienRuinsslE

	AlteredHumanityslE := []string{}
	AlteredHumanityslE = append(AlteredHumanityslE, "Biochauvinist local")
	AlteredHumanityslE = append(AlteredHumanityslE, "Local experimenter")
	AlteredHumanityslE = append(AlteredHumanityslE, "Mentally unstable mutant")
	EnemyTagMap["AlteredHumanity"] = AlteredHumanityslE

	AnarchistsslE := []string{}
	AnarchistsslE = append(AnarchistsslE, "Offworlder imperialist")
	AnarchistsslE = append(AnarchistsslE, "Reformer seeking to impose “good government”")
	AnarchistsslE = append(AnarchistsslE, "Exploiter taking advantage of the lack of centralized resistance")
	EnemyTagMap["Anarchists"] = AnarchistsslE

	AnthropomorphsslE := []string{}
	AnthropomorphsslE = append(AnthropomorphsslE, "Anthro-supremacist local")
	AnthropomorphsslE = append(AnthropomorphsslE, "Native driven by feral urges")
	AnthropomorphsslE = append(AnthropomorphsslE, "Outside exploiter who sees the locals as subhuman creatures")
	EnemyTagMap["Anthropomorphs"] = AnthropomorphsslE

	Area51slE := []string{}
	Area51slE = append(Area51slE, "Suspicious government minder")
	Area51slE = append(Area51slE, "Free merchant who likes his local monopoly")
	Area51slE = append(Area51slE, "Local who wants a specimen for dissection")
	EnemyTagMap["Area51"] = Area51slE

	BadlandsWorldslE := []string{}
	BadlandsWorldslE = append(BadlandsWorldslE, "Mutated badlands fauna")
	BadlandsWorldslE = append(BadlandsWorldslE, "Desperate local")
	BadlandsWorldslE = append(BadlandsWorldslE, "Badlands raider chief")
	EnemyTagMap["BadlandsWorld"] = BadlandsWorldslE

	BattlegroundslE := []string{}
	BattlegroundslE = append(BattlegroundslE, "Ruthless military commander")
	BattlegroundslE = append(BattlegroundslE, "Looter pack chieftain")
	BattlegroundslE = append(BattlegroundslE, "Traitorous collaborator")
	EnemyTagMap["Battleground"] = BattlegroundslE

	BeastmastersslE := []string{}
	BeastmastersslE = append(BeastmastersslE, "Half-feral warlord of a beast swarm")
	BeastmastersslE = append(BeastmastersslE, "Coldly inhuman scientist")
	BeastmastersslE = append(BeastmastersslE, "Altered beast with human intellect and furious malice")
	EnemyTagMap["Beastmasters"] = BeastmastersslE

	BubbleCitiesslE := []string{}
	BubbleCitiesslE = append(BubbleCitiesslE, "Native dreading outsider contamination")
	BubbleCitiesslE = append(BubbleCitiesslE, "Saboteur from another bubble city")
	BubbleCitiesslE = append(BubbleCitiesslE, "Local official hostile to outsider ignorance of laws")
	EnemyTagMap["BubbleCities"] = BubbleCitiesslE

	CheapLifeslE := []string{}
	CheapLifeslE = append(CheapLifeslE, "Master assassin")
	CheapLifeslE = append(CheapLifeslE, "Bloody-handed judge")
	CheapLifeslE = append(CheapLifeslE, "Overseer of disposable clones")
	EnemyTagMap["CheapLife"] = CheapLifeslE

	CivilWarslE := []string{}
	CivilWarslE = append(CivilWarslE, "Faction commissar")
	CivilWarslE = append(CivilWarslE, "Angry native")
	CivilWarslE = append(CivilWarslE, "Conspiracy theorist who blames offworlders for the war")
	CivilWarslE = append(CivilWarslE, "Deserter looking out for himself")
	CivilWarslE = append(CivilWarslE, "Guerrilla bandit chieftain")
	EnemyTagMap["CivilWar"] = CivilWarslE

	ColdWarslE := []string{}
	ColdWarslE = append(ColdWarslE, "Suspicious chief of intelligence")
	ColdWarslE = append(ColdWarslE, "Native who thinks the outworlders are with the other side")
	ColdWarslE = append(ColdWarslE, "Femme fatale")
	EnemyTagMap["ColdWar"] = ColdWarslE

	ColonizedPopulationslE := []string{}
	ColonizedPopulationslE = append(ColonizedPopulationslE, "Suspicious security personnel")
	ColonizedPopulationslE = append(ColonizedPopulationslE, "Offworlder-hating natives")
	ColonizedPopulationslE = append(ColonizedPopulationslE, "Local crime boss preying on rich offworlders")
	EnemyTagMap["ColonizedPopulation"] = ColonizedPopulationslE

	CulturalPowerslE := []string{}
	CulturalPowerslE = append(CulturalPowerslE, "Murderously eccentric artist")
	CulturalPowerslE = append(CulturalPowerslE, "Crazed fan")
	CulturalPowerslE = append(CulturalPowerslE, "Failed artist with an obsessive grudge")
	CulturalPowerslE = append(CulturalPowerslE, "Critic with a crusade to enact")
	EnemyTagMap["CulturalPower"] = CulturalPowerslE

	CybercommunistsslE := []string{}
	CybercommunistsslE = append(CybercommunistsslE, "Embittered rebel against perceived unfairness")
	CybercommunistsslE = append(CybercommunistsslE, "Offworlder saboteur")
	CybercommunistsslE = append(CybercommunistsslE, "Aspiring Stalin-figure")
	EnemyTagMap["Cybercommunists"] = CybercommunistsslE

	CyborgsslE := []string{}
	CyborgsslE = append(CyborgsslE, "Ambitious hacker of cyber implants")
	CyborgsslE = append(CyborgsslE, "Cybertech oligarch")
	CyborgsslE = append(CyborgsslE, "Researcher craving fresh offworlders")
	CyborgsslE = append(CyborgsslE, "Cybered-up gang boss")
	EnemyTagMap["Cyborgs"] = CyborgsslE

	CyclicalDoomslE := []string{}
	CyclicalDoomslE = append(CyclicalDoomslE, "Offworlder seeking to trigger the apocalypse early for profit")
	CyclicalDoomslE = append(CyclicalDoomslE, "Local recklessly taking advantage of preparation stores")
	CyclicalDoomslE = append(CyclicalDoomslE, "Demagogue claiming the cycle is merely a myth of the authorities")
	EnemyTagMap["CyclicalDoom"] = CyclicalDoomslE

	DesertWorldslE := []string{}
	DesertWorldslE = append(DesertWorldslE, "Raider chieftain")
	DesertWorldslE = append(DesertWorldslE, "Crazed hermit")
	DesertWorldslE = append(DesertWorldslE, "Angry isolationists")
	DesertWorldslE = append(DesertWorldslE, "Paranoid mineral prospector")
	DesertWorldslE = append(DesertWorldslE, "Strange desert beast")
	EnemyTagMap["DesertWorld"] = DesertWorldslE

	DoomedWorldslE := []string{}
	DoomedWorldslE = append(DoomedWorldslE, "Crazed prophet of a false salvation")
	DoomedWorldslE = append(DoomedWorldslE, "Ruthless leader seeking to flee with their treasures")
	DoomedWorldslE = append(DoomedWorldslE, "Cynical ship captain selling a one-way trip into hard vacuum as escape to another world")
	EnemyTagMap["DoomedWorld"] = DoomedWorldslE

	DyingRaceslE := []string{}
	DyingRaceslE = append(DyingRaceslE, "Hostile outsider who wants the locals dead")
	DyingRaceslE = append(DyingRaceslE, "Offworlder seeking to take advantage of their weakened state")
	DyingRaceslE = append(DyingRaceslE, "Invaders eager to push the locals out of their former lands")
	EnemyTagMap["DyingRace"] = DyingRaceslE

	EugenicCultslE := []string{}
	EugenicCultslE = append(EugenicCultslE, "Eugenic superiority fanatic")
	EugenicCultslE = append(EugenicCultslE, "Mentally unstable homo superior")
	EugenicCultslE = append(EugenicCultslE, "Mad eugenic scientist")
	EnemyTagMap["EugenicCult"] = EugenicCultslE

	ExchangeConsulateslE := []string{}
	ExchangeConsulateslE = append(ExchangeConsulateslE, "Corrupt Exchange official")
	ExchangeConsulateslE = append(ExchangeConsulateslE, "Indebted native who thinks the players are Exchange agents")
	ExchangeConsulateslE = append(ExchangeConsulateslE, "Exchange official dunning the players for debts incurred")
	EnemyTagMap["ExchangeConsulate"] = ExchangeConsulateslE

	FallenHegemonslE := []string{}
	FallenHegemonslE = append(FallenHegemonslE, "Bitter pretender to a meaningless throne")
	FallenHegemonslE = append(FallenHegemonslE, "Resentful official dreaming of empire")
	FallenHegemonslE = append(FallenHegemonslE, "Vengeful offworlder seeking to punish their old rulers")
	EnemyTagMap["FallenHegemon"] = FallenHegemonslE

	FeralWorldslE := []string{}
	FeralWorldslE = append(FeralWorldslE, "Decadent noble")
	FeralWorldslE = append(FeralWorldslE, "Mad cultist")
	FeralWorldslE = append(FeralWorldslE, "Xenophobic local")
	FeralWorldslE = append(FeralWorldslE, "Cannibal chief")
	FeralWorldslE = append(FeralWorldslE, "Maltech researcher")
	EnemyTagMap["FeralWorld"] = FeralWorldslE

	FlyingCitiesslE := []string{}
	FlyingCitiesslE = append(FlyingCitiesslE, "Rival city pilot")
	FlyingCitiesslE = append(FlyingCitiesslE, "Tech thief attempting to steal outworld gear")
	FlyingCitiesslE = append(FlyingCitiesslE, "Saboteur or scavenger plundering the city’s tech")
	EnemyTagMap["FlyingCities"] = FlyingCitiesslE

	ForbiddenTechslE := []string{}
	ForbiddenTechslE = append(ForbiddenTechslE, "Mad scientist")
	ForbiddenTechslE = append(ForbiddenTechslE, "Maltech buyer from offworld")
	ForbiddenTechslE = append(ForbiddenTechslE, "Security enforcer")
	EnemyTagMap["ForbiddenTech"] = ForbiddenTechslE

	FormerWarriorsslE := []string{}
	FormerWarriorsslE = append(FormerWarriorsslE, "Unreformed warlord leader")
	FormerWarriorsslE = append(FormerWarriorsslE, "Bitter mercenary chief")
	FormerWarriorsslE = append(FormerWarriorsslE, "Victim of their warfare seeking revenge")
	EnemyTagMap["FormerWarriors"] = FormerWarriorsslE

	FreakGeologyslE := []string{}
	FreakGeologyslE = append(FreakGeologyslE, "Crank xenogeologist")
	FreakGeologyslE = append(FreakGeologyslE, "Cultist who believes it the work of aliens")
	EnemyTagMap["FreakGeology"] = FreakGeologyslE

	FreakWeatherslE := []string{}
	FreakWeatherslE = append(FreakWeatherslE, "Criminal using the weather as a cover")
	FreakWeatherslE = append(FreakWeatherslE, "Weather cultists convinced the offworlders are responsible for some disaster")
	FreakWeatherslE = append(FreakWeatherslE, "Native predators dependent on the weather")
	EnemyTagMap["FreakWeather"] = FreakWeatherslE

	FriendlyFoeslE := []string{}
	FriendlyFoeslE = append(FriendlyFoeslE, "Driven hater of all their kind")
	FriendlyFoeslE = append(FriendlyFoeslE, "Internal malcontent bent on creating conflict")
	FriendlyFoeslE = append(FriendlyFoeslE, "Secret master who seeks to lure trust")
	EnemyTagMap["FriendlyFoe"] = FriendlyFoeslE

	GoldRushslE := []string{}
	GoldRushslE = append(GoldRushslE, "Paranoid prospector")
	GoldRushslE = append(GoldRushslE, "Aspiring mining tycoon")
	GoldRushslE = append(GoldRushslE, "Rapacious merchant")
	EnemyTagMap["GoldRush"] = GoldRushslE

	GreatWorkslE := []string{}
	GreatWorkslE = append(GreatWorkslE, "Local planning to sacrifice the PCs for the work")
	GreatWorkslE = append(GreatWorkslE, "Local who thinks the PCs threaten the work")
	GreatWorkslE = append(GreatWorkslE, "Obsessive zealot ready to destroy someone or something important to the PCs for the sake of the work")
	EnemyTagMap["GreatWork"] = GreatWorkslE

	HatredslE := []string{}
	HatredslE = append(HatredslE, "Native convinced that the offworlders are agents of Them")
	HatredslE = append(HatredslE, "Cynical politician in need of scapegoats")
	EnemyTagMap["Hatred"] = HatredslE

	HeavyIndustryslE := []string{}
	HeavyIndustryslE = append(HeavyIndustryslE, "Tycoon monopolist")
	HeavyIndustryslE = append(HeavyIndustryslE, "Industrial spy")
	HeavyIndustryslE = append(HeavyIndustryslE, "Malcontent revolutionary")
	EnemyTagMap["HeavyIndustry"] = HeavyIndustryslE

	HeavyMiningslE := []string{}
	HeavyMiningslE = append(HeavyMiningslE, "Mine boss")
	HeavyMiningslE = append(HeavyMiningslE, "Tunnel saboteur")
	HeavyMiningslE = append(HeavyMiningslE, "Subterranean predators")
	EnemyTagMap["HeavyMining"] = HeavyMiningslE

	HivemindslE := []string{}
	HivemindslE = append(HivemindslE, "A hivemind that wants to assimilate outsiders")
	HivemindslE = append(HivemindslE, "A hivemind that has no respect for unjoined life")
	HivemindslE = append(HivemindslE, "A hivemind that fears and hates unjoined life")
	EnemyTagMap["Hivemind"] = HivemindslE

	HolyWarslE := []string{}
	HolyWarslE = append(HolyWarslE, "Blood-mad pontiff")
	HolyWarslE = append(HolyWarslE, "Coldly cynical secular leader")
	HolyWarslE = append(HolyWarslE, "Totalitarian political demagogue")
	EnemyTagMap["HolyWar"] = HolyWarslE

	HostileBiosphereslE := []string{}
	HostileBiosphereslE = append(HostileBiosphereslE, "Local fauna")
	HostileBiosphereslE = append(HostileBiosphereslE, "Nature cultist")
	HostileBiosphereslE = append(HostileBiosphereslE, "Native aliens")
	HostileBiosphereslE = append(HostileBiosphereslE, "Callous labor overseer")
	EnemyTagMap["HostileBiosphere"] = HostileBiosphereslE

	HostileSpaceslE := []string{}
	HostileSpaceslE = append(HostileSpaceslE, "Alien raid leader")
	HostileSpaceslE = append(HostileSpaceslE, "Meteor-launching terrorists")
	HostileSpaceslE = append(HostileSpaceslE, "Paranoid local leader")
	EnemyTagMap["HostileSpace"] = HostileSpaceslE

	ImmortalsslE := []string{}
	ImmortalsslE = append(ImmortalsslE, "Outsider determined to steal immortality")
	ImmortalsslE = append(ImmortalsslE, "Smug local convinced of their immortal wisdom to rule all")
	ImmortalsslE = append(ImmortalsslE, "Offworlder seeking the world’s ruin before it becomes a threat to all")
	EnemyTagMap["Immortals"] = ImmortalsslE

	LocalSpecialtyslE := []string{}
	LocalSpecialtyslE = append(LocalSpecialtyslE, "Monopolist")
	LocalSpecialtyslE = append(LocalSpecialtyslE, "Offworlder seeking prohibition of the specialty")
	LocalSpecialtyslE = append(LocalSpecialtyslE, "Native who views the specialty as sacred")
	EnemyTagMap["LocalSpecialty"] = LocalSpecialtyslE

	LocalTechslE := []string{}
	LocalTechslE = append(LocalTechslE, "Keeper of the tech")
	LocalTechslE = append(LocalTechslE, "Offworld industrialist")
	LocalTechslE = append(LocalTechslE, "Automated defenses that suddenly come alive")
	LocalTechslE = append(LocalTechslE, "Native alien mentors")
	EnemyTagMap["LocalTech"] = LocalTechslE

	MajorSpaceyardslE := []string{}
	MajorSpaceyardslE = append(MajorSpaceyardslE, "Enemy saboteur")
	MajorSpaceyardslE = append(MajorSpaceyardslE, "Industrial spy")
	MajorSpaceyardslE = append(MajorSpaceyardslE, "Scheming construction tycoon")
	MajorSpaceyardslE = append(MajorSpaceyardslE, "Aspiring ship hijacker")
	EnemyTagMap["MajorSpaceyard"] = MajorSpaceyardslE

	MandarinateslE := []string{}
	MandarinateslE = append(MandarinateslE, "Corrupt test administrator")
	MandarinateslE = append(MandarinateslE, "Incompetent but highly-rated graduate")
	MandarinateslE = append(MandarinateslE, "Ruthless leader of a clan of high-testing relations")
	EnemyTagMap["Mandarinate"] = MandarinateslE

	MandateBaseslE := []string{}
	MandateBaseslE = append(MandateBaseslE, "Deranged Mandate monitoring AI")
	MandateBaseslE = append(MandateBaseslE, "Aspiring sector ruler")
	MandateBaseslE = append(MandateBaseslE, "Demagogue preaching local superiority over “traitorous rebel worlds”.")
	EnemyTagMap["MandateBase"] = MandateBaseslE

	ManeatersslE := []string{}
	ManeatersslE = append(ManeatersslE, "Ruthless ghoul leader")
	ManeatersslE = append(ManeatersslE, "Chieftain of a ravenous tribe")
	ManeatersslE = append(ManeatersslE, "Sophisticated degenerate preaching the splendid authenticity of cannibalism")
	EnemyTagMap["Maneaters"] = ManeatersslE

	MegacorpsslE := []string{}
	MegacorpsslE = append(MegacorpsslE, "Megalomaniacal executive")
	MegacorpsslE = append(MegacorpsslE, "Underling looking to use the PCs as catspaws")
	MegacorpsslE = append(MegacorpsslE, "Ruthless mercenary who wants what the PCs have")
	EnemyTagMap["Megacorps"] = MegacorpsslE

	MercenariesslE := []string{}
	MercenariesslE = append(MercenariesslE, "Amoral mercenary leader")
	MercenariesslE = append(MercenariesslE, "Rich offworlder trying to buy rule of the world")
	MercenariesslE = append(MercenariesslE, "Mercenary press gang chief forcing locals into service")
	EnemyTagMap["Mercenaries"] = MercenariesslE

	MisandryMisogynyslE := []string{}
	MisandryMisogynyslE = append(MisandryMisogynyslE, "Cultural fundamentalist")
	MisandryMisogynyslE = append(MisandryMisogynyslE, "Cultural missionary to outworlders")
	MisandryMisogynyslE = append(MisandryMisogynyslE, "Local rebel driven to pointless and meaningless violence")
	EnemyTagMap["MisandryMisogyny"] = MisandryMisogynyslE

	NightWorldslE := []string{}
	NightWorldslE = append(NightWorldslE, "Monstrous thing from the night")
	NightWorldslE = append(NightWorldslE, "Offworlder finding the obscurity of the world convenient for dark purposes")
	NightWorldslE = append(NightWorldslE, "Mad scientist experimenting with local life")
	EnemyTagMap["NightWorld"] = NightWorldslE

	MinimalContactslE := []string{}
	MinimalContactslE = append(MinimalContactslE, "Customs official")
	MinimalContactslE = append(MinimalContactslE, "Xenophobic natives")
	MinimalContactslE = append(MinimalContactslE, "Existing merchant who doesn’t like competition")
	EnemyTagMap["MinimalContact"] = MinimalContactslE

	NomadsslE := []string{}
	NomadsslE = append(NomadsslE, "Desperate tribal leader who needs what the PCs have")
	NomadsslE = append(NomadsslE, "Ruthless raider chieftain")
	NomadsslE = append(NomadsslE, "Leader seeking to weld the nomads into an army")
	EnemyTagMap["Nomads"] = NomadsslE

	OceanicWorldslE := []string{}
	OceanicWorldslE = append(OceanicWorldslE, "Pirate raider")
	OceanicWorldslE = append(OceanicWorldslE, "Violent “salvager” gang")
	OceanicWorldslE = append(OceanicWorldslE, "Tentacled sea monster")
	EnemyTagMap["OceanicWorld"] = OceanicWorldslE

	OutofContactslE := []string{}
	OutofContactslE = append(OutofContactslE, "Fearful local ruler")
	OutofContactslE = append(OutofContactslE, "Zealous native cleric")
	OutofContactslE = append(OutofContactslE, "Sinister power that has kept the world isolated")
	EnemyTagMap["OutofContact"] = OutofContactslE

	OutpostWorldslE := []string{}
	OutpostWorldslE = append(OutpostWorldslE, "Space-mad outpost staffer")
	OutpostWorldslE = append(OutpostWorldslE, "Outpost commander who wants it to stay undiscovered")
	OutpostWorldslE = append(OutpostWorldslE, "Undercover saboteur")
	EnemyTagMap["OutpostWorld"] = OutpostWorldslE

	PerimeterAgencyslE := []string{}
	PerimeterAgencyslE = append(PerimeterAgencyslE, "Renegade Agency Director")
	PerimeterAgencyslE = append(PerimeterAgencyslE, "Maltech researcher")
	PerimeterAgencyslE = append(PerimeterAgencyslE, "Paranoid intelligence chief")
	EnemyTagMap["PerimeterAgency"] = PerimeterAgencyslE

	PilgrimageSiteslE := []string{}
	PilgrimageSiteslE = append(PilgrimageSiteslE, "Saboteur devoted to a rival belief")
	PilgrimageSiteslE = append(PilgrimageSiteslE, "Bitter reformer who resents the current leadership")
	PilgrimageSiteslE = append(PilgrimageSiteslE, "Swindler conning the pilgrims")
	EnemyTagMap["PilgrimageSite"] = PilgrimageSiteslE

	PleasureWorldslE := []string{}
	PleasureWorldslE = append(PleasureWorldslE, "Purveyor of evil delights")
	PleasureWorldslE = append(PleasureWorldslE, "Local seeking to control others with addictions")
	PleasureWorldslE = append(PleasureWorldslE, "Offworlder exploiter of native resources")
	EnemyTagMap["PleasureWorld"] = PleasureWorldslE

	PoliceStateslE := []string{}
	PoliceStateslE = append(PoliceStateslE, "Secret police chief")
	PoliceStateslE = append(PoliceStateslE, "Scapegoating official")
	PoliceStateslE = append(PoliceStateslE, "Treacherous native informer")
	EnemyTagMap["PoliceState"] = PoliceStateslE

	PostScarcityslE := []string{}
	PostScarcityslE = append(PostScarcityslE, "Frenzied ideologue fighting over an idea")
	PostScarcityslE = append(PostScarcityslE, "Paranoid local fearing offworlder influence")
	PostScarcityslE = append(PostScarcityslE, "Grim reformer seeking the destruction of the “enfeebling” productive tech")
	EnemyTagMap["PostScarcity"] = PostScarcityslE

	PreceptorArchiveslE := []string{}
	PreceptorArchiveslE = append(PreceptorArchiveslE, "Luddite native")
	PreceptorArchiveslE = append(PreceptorArchiveslE, "Offworld merchant who wants the natives kept ignorant")
	PreceptorArchiveslE = append(PreceptorArchiveslE, "Religious zealot")
	PreceptorArchiveslE = append(PreceptorArchiveslE, "Corrupted First Speaker who wants to keep a monopoly on learning")
	EnemyTagMap["PreceptorArchive"] = PreceptorArchiveslE

	PretechCultistsslE := []string{}
	PretechCultistsslE = append(PretechCultistsslE, "Cult leader")
	PretechCultistsslE = append(PretechCultistsslE, "Artifact supplier")
	PretechCultistsslE = append(PretechCultistsslE, "Pretech smuggler")
	EnemyTagMap["PretechCultists"] = PretechCultistsslE

	PrimitiveAliensslE := []string{}
	PrimitiveAliensslE = append(PrimitiveAliensslE, "Hostile alien chief")
	PrimitiveAliensslE = append(PrimitiveAliensslE, "Human firebrand")
	PrimitiveAliensslE = append(PrimitiveAliensslE, "Dangerous local predator")
	PrimitiveAliensslE = append(PrimitiveAliensslE, "Alien religious zealot")
	EnemyTagMap["PrimitiveAliens"] = PrimitiveAliensslE

	PrisonPlanetslE := []string{}
	PrisonPlanetslE = append(PrisonPlanetslE, "Crazed warden AI")
	PrisonPlanetslE = append(PrisonPlanetslE, "Brutal heir to gang leadership")
	PrisonPlanetslE = append(PrisonPlanetslE, "Offworlder who’s somehow acquired warden powers and exploits the locals")
	EnemyTagMap["PrisonPlanet"] = PrisonPlanetslE

	PsionicsAcademyslE := []string{}
	PsionicsAcademyslE = append(PsionicsAcademyslE, "Corrupt psychic instructor")
	PsionicsAcademyslE = append(PsionicsAcademyslE, "Renegade student")
	PsionicsAcademyslE = append(PsionicsAcademyslE, "Mad psychic researcher")
	PsionicsAcademyslE = append(PsionicsAcademyslE, "Resentful townie")
	EnemyTagMap["PsionicsAcademy"] = PsionicsAcademyslE

	PsionicsFearslE := []string{}
	PsionicsFearslE = append(PsionicsFearslE, "Mental purity investigator")
	PsionicsFearslE = append(PsionicsFearslE, "Suspicious zealot")
	PsionicsFearslE = append(PsionicsFearslE, "Witch-finder")
	EnemyTagMap["PsionicsFear"] = PsionicsFearslE

	PsionicsWorshipslE := []string{}
	PsionicsWorshipslE = append(PsionicsWorshipslE, "Psychic inquisitor")
	PsionicsWorshipslE = append(PsionicsWorshipslE, "Haughty mind-noble")
	PsionicsWorshipslE = append(PsionicsWorshipslE, "Psychic slaver")
	PsionicsWorshipslE = append(PsionicsWorshipslE, "Feral prophet")
	EnemyTagMap["PsionicsWorship"] = PsionicsWorshipslE

	QuarantinedWorldslE := []string{}
	QuarantinedWorldslE = append(QuarantinedWorldslE, "Defense installation commander")
	QuarantinedWorldslE = append(QuarantinedWorldslE, "Suspicious patrol leader")
	QuarantinedWorldslE = append(QuarantinedWorldslE, "Crazed asteroid hermit")
	EnemyTagMap["QuarantinedWorld"] = QuarantinedWorldslE

	RadioactiveWorldslE := []string{}
	RadioactiveWorldslE = append(RadioactiveWorldslE, "Bitter mutant")
	RadioactiveWorldslE = append(RadioactiveWorldslE, "Relic warlord")
	RadioactiveWorldslE = append(RadioactiveWorldslE, "Desperate wouldbe escapee")
	EnemyTagMap["RadioactiveWorld"] = RadioactiveWorldslE

	RefugeesslE := []string{}
	RefugeesslE = append(RefugeesslE, "Xenophobic native leader")
	RefugeesslE = append(RefugeesslE, "Refugee chief aspiring to seize the host nation")
	RefugeesslE = append(RefugeesslE, "Politician seeking to use the refugees as a weapon")
	EnemyTagMap["Refugees"] = RefugeesslE

	RegionalHegemonslE := []string{}
	RegionalHegemonslE = append(RegionalHegemonslE, "Ambitious general")
	RegionalHegemonslE = append(RegionalHegemonslE, "Colonial official")
	RegionalHegemonslE = append(RegionalHegemonslE, "Contemptuous noble")
	EnemyTagMap["RegionalHegemon"] = RegionalHegemonslE

	RestrictiveLawsslE := []string{}
	RestrictiveLawsslE = append(RestrictiveLawsslE, "Law enforcement officer")
	RestrictiveLawsslE = append(RestrictiveLawsslE, "Outraged native")
	RestrictiveLawsslE = append(RestrictiveLawsslE, "Native lawyer specializing in peeling offworlders")
	RestrictiveLawsslE = append(RestrictiveLawsslE, "Paid snitch")
	EnemyTagMap["RestrictiveLaws"] = RestrictiveLawsslE

	RevanchistsslE := []string{}
	RevanchistsslE = append(RevanchistsslE, "Demagogue whipping the locals on to a hopeless war")
	RevanchistsslE = append(RevanchistsslE, "Politician seeking to use the resentment for their own ends")
	RevanchistsslE = append(RevanchistsslE, "Local convinced the PCs are agents of the “thieving” power")
	RevanchistsslE = append(RevanchistsslE, "Refugee from the land bitterly demanding it be reclaimed")
	EnemyTagMap["Revanchists"] = RevanchistsslE

	RevolutionariesslE := []string{}
	RevolutionariesslE = append(RevolutionariesslE, "Blood-drenched revolutionary leader")
	RevolutionariesslE = append(RevolutionariesslE, "Blooddrenched secret police chief")
	RevolutionariesslE = append(RevolutionariesslE, "Hostile foreign agent seeking further turmoil")
	EnemyTagMap["Revolutionaries"] = RevolutionariesslE

	RigidCultureslE := []string{}
	RigidCultureslE = append(RigidCultureslE, "Rigid reactionary")
	RigidCultureslE = append(RigidCultureslE, "Wary ruler")
	RigidCultureslE = append(RigidCultureslE, "Regime ideologue")
	RigidCultureslE = append(RigidCultureslE, "Offended potentate")
	EnemyTagMap["RigidCulture"] = RigidCultureslE

	RisingHegemonslE := []string{}
	RisingHegemonslE = append(RisingHegemonslE, "Jingoistic supremacist")
	RisingHegemonslE = append(RisingHegemonslE, "Official bent on glorious success")
	RisingHegemonslE = append(RisingHegemonslE, "Foreign agent saboteur")
	EnemyTagMap["RisingHegemon"] = RisingHegemonslE

	RitualCombatslE := []string{}
	RitualCombatslE = append(RitualCombatslE, "Bloodthirsty local champion")
	RitualCombatslE = append(RitualCombatslE, "Ambitious gladiator stable owner")
	RitualCombatslE = append(RitualCombatslE, "Xenophobic master fighter")
	EnemyTagMap["RitualCombat"] = RitualCombatslE

	RobotsslE := []string{}
	RobotsslE = append(RobotsslE, "Hostile robot master")
	RobotsslE = append(RobotsslE, "Robot greedy to seize offworld tech")
	RobotsslE = append(RobotsslE, "Robot fallen in love with the PC’s ship")
	RobotsslE = append(RobotsslE, "Oligarch whose factories build robots")
	EnemyTagMap["Robots"] = RobotsslE

	SeagoingCitiesslE := []string{}
	SeagoingCitiesslE = append(SeagoingCitiesslE, "Pirate city lord")
	SeagoingCitiesslE = append(SeagoingCitiesslE, "Mer-human raider chieftain")
	SeagoingCitiesslE = append(SeagoingCitiesslE, "Hostile landsman noble")
	SeagoingCitiesslE = append(SeagoingCitiesslE, "Enemy city saboteur")
	EnemyTagMap["SeagoingCities"] = SeagoingCitiesslE

	SealedMenaceslE := []string{}
	SealedMenaceslE = append(SealedMenaceslE, "Hostile outsider bent on freeing the menace")
	SealedMenaceslE = append(SealedMenaceslE, "Misguided fool who thinks he can use it")
	SealedMenaceslE = append(SealedMenaceslE, "Reckless researcher who thinks he can fix it")
	EnemyTagMap["SealedMenace"] = SealedMenaceslE

	SecretMastersslE := []string{}
	SecretMastersslE = append(SecretMastersslE, "An agent of the cabal")
	SecretMastersslE = append(SecretMastersslE, "Government official who wants no questions asked")
	SecretMastersslE = append(SecretMastersslE, "Willfully blinded local")
	EnemyTagMap["SecretMasters"] = SecretMastersslE

	SectariansslE := []string{}
	SectariansslE = append(SectariansslE, "Paranoid believer")
	SectariansslE = append(SectariansslE, "Native convinced the party is working for the other side")
	SectariansslE = append(SectariansslE, "Absolutist ruler")
	EnemyTagMap["Sectarians"] = SectariansslE

	SeismicInstabilityslE := []string{}
	SeismicInstabilityslE = append(SeismicInstabilityslE, "Earthquake cultist")
	SeismicInstabilityslE = append(SeismicInstabilityslE, "Hermit seismologist")
	SeismicInstabilityslE = append(SeismicInstabilityslE, "Burrowing native life form")
	SeismicInstabilityslE = append(SeismicInstabilityslE, "Earthquake-inducing saboteur")
	EnemyTagMap["SeismicInstability"] = SeismicInstabilityslE

	ShackledWorldslE := []string{}
	ShackledWorldslE = append(ShackledWorldslE, "Passionless jailer-AI")
	ShackledWorldslE = append(ShackledWorldslE, "Paranoid military grid AI")
	ShackledWorldslE = append(ShackledWorldslE, "Robot overlord")
	ShackledWorldslE = append(ShackledWorldslE, "Enigmatic alien master")
	EnemyTagMap["ShackledWorld"] = ShackledWorldslE

	SocietalDespairslE := []string{}
	SocietalDespairslE = append(SocietalDespairslE, "Zealot who blames outsiders for the decay")
	SocietalDespairslE = append(SocietalDespairslE, "Nihilistic warlord")
	SocietalDespairslE = append(SocietalDespairslE, "Offworlder looking to exploit the local despair")
	EnemyTagMap["SocietalDespair"] = SocietalDespairslE

	SoleSupplierslE := []string{}
	SoleSupplierslE = append(SoleSupplierslE, "Resource oligarch")
	SoleSupplierslE = append(SoleSupplierslE, "Ruthless smuggler")
	SoleSupplierslE = append(SoleSupplierslE, "Resource- controlling warlord")
	SoleSupplierslE = append(SoleSupplierslE, "Foreign agent seeking to subvert local government")
	EnemyTagMap["SoleSupplier"] = SoleSupplierslE

	TabooTreasureslE := []string{}
	TabooTreasureslE = append(TabooTreasureslE, "Maker of a vile commodity")
	TabooTreasureslE = append(TabooTreasureslE, "Smuggler for a powerful offworlder")
	TabooTreasureslE = append(TabooTreasureslE, "Depraved offworlder here for “fun”")
	TabooTreasureslE = append(TabooTreasureslE, "Local warlord who controls the treasure")
	EnemyTagMap["TabooTreasure"] = TabooTreasureslE

	TerraformFailureslE := []string{}
	TerraformFailureslE = append(TerraformFailureslE, "Brutal ruler who cares only for their people")
	TerraformFailureslE = append(TerraformFailureslE, "Offworlder trying to loot the damaged engines")
	TerraformFailureslE = append(TerraformFailureslE, "Warlord trying to seize limited habitable land")
	EnemyTagMap["TerraformFailure"] = TerraformFailureslE

	TheocracyslE := []string{}
	TheocracyslE = append(TheocracyslE, "Decadent priest-ruler")
	TheocracyslE = append(TheocracyslE, "Zealous inquisitor")
	TheocracyslE = append(TheocracyslE, "Relentless proselytizer")
	TheocracyslE = append(TheocracyslE, "True Believer")
	EnemyTagMap["Theocracy"] = TheocracyslE

	TombWorldslE := []string{}
	TombWorldslE = append(TombWorldslE, "Demented survivor tribe chieftain")
	TombWorldslE = append(TombWorldslE, "Avaricious scavenger")
	TombWorldslE = append(TombWorldslE, "Automated defense system")
	TombWorldslE = append(TombWorldslE, "Native predator")
	EnemyTagMap["TombWorld"] = TombWorldslE

	TradeHubslE := []string{}
	TradeHubslE = append(TradeHubslE, "Cheating merchant")
	TradeHubslE = append(TradeHubslE, "Thieving dockworker")
	TradeHubslE = append(TradeHubslE, "Commercial spy")
	TradeHubslE = append(TradeHubslE, "Corrupt customs official")
	EnemyTagMap["TradeHub"] = TradeHubslE

	TyrannyslE := []string{}
	TyrannyslE = append(TyrannyslE, "Debauched autocrat")
	TyrannyslE = append(TyrannyslE, "Sneering bully-boy")
	TyrannyslE = append(TyrannyslE, "Soulless government official")
	TyrannyslE = append(TyrannyslE, "Occupying army officer")
	EnemyTagMap["Tyranny"] = TyrannyslE

	UnbrakedAIslE := []string{}
	UnbrakedAIslE = append(UnbrakedAIslE, "AI Cultist")
	UnbrakedAIslE = append(UnbrakedAIslE, "Maltech researcher")
	UnbrakedAIslE = append(UnbrakedAIslE, "Government official dependent on the AI")
	EnemyTagMap["UnbrakedAI"] = UnbrakedAIslE

	UrbanizedSurfaceslE := []string{}
	UrbanizedSurfaceslE = append(UrbanizedSurfaceslE, "Maintenance AI that hates outsiders")
	UrbanizedSurfaceslE = append(UrbanizedSurfaceslE, "Tyrant of a habitation block")
	UrbanizedSurfaceslE = append(UrbanizedSurfaceslE, "Deep-dwelling prophet who considers “the sky” a blasphemy to be quelled")
	EnemyTagMap["UrbanizedSurface"] = UrbanizedSurfaceslE

	UtopiaslE := []string{}
	UtopiaslE = append(UtopiaslE, "Compassionate neurotherapist")
	UtopiaslE = append(UtopiaslE, "Proselytizing native missionary to outsiders")
	UtopiaslE = append(UtopiaslE, "Brutal tyrant who rules through inexorable happiness")
	EnemyTagMap["Utopia"] = UtopiaslE

	WarlordsslE := []string{}
	WarlordsslE = append(WarlordsslE, "Warlord")
	WarlordsslE = append(WarlordsslE, "Avaricious lieutenant")
	WarlordsslE = append(WarlordsslE, "Expensive assassin")
	WarlordsslE = append(WarlordsslE, "Aspiring minion")
	EnemyTagMap["Warlords"] = WarlordsslE

	XenophilesslE := []string{}
	XenophilesslE = append(XenophilesslE, "Offworld xenophobe")
	XenophilesslE = append(XenophilesslE, "Suspicious alien leader")
	XenophilesslE = append(XenophilesslE, "Xenocultural imperialist")
	EnemyTagMap["Xenophiles"] = XenophilesslE

	XenophobesslE := []string{}
	XenophobesslE = append(XenophobesslE, "Revulsed local ruler")
	XenophobesslE = append(XenophobesslE, "Native convinced some wrong was done to him")
	XenophobesslE = append(XenophobesslE, "Cynical demagogue")
	EnemyTagMap["Xenophobes"] = XenophobesslE

	ZombiesslE := []string{}
	ZombiesslE = append(ZombiesslE, "Soulless maltech biotechnology cult")
	ZombiesslE = append(ZombiesslE, "Sinister governmental agent")
	ZombiesslE = append(ZombiesslE, "Crazed zombie cultist")
	EnemyTagMap["Zombies"] = ZombiesslE
	return EnemyTagMap
}

func createComplicationsMap() map[string][]string {
	CompicationTagMap := make(map[string][]string)
	AbandonedColonyslC := []string{}
	AbandonedColonyslC = append(AbandonedColonyslC, "The local government wants the ruins to remain a secret")
	AbandonedColonyslC = append(AbandonedColonyslC, "The locals claim ownership of it")
	AbandonedColonyslC = append(AbandonedColonyslC, "The colony is crumbling and dangerous to navigate")
	CompicationTagMap["AbandonedColony"] = AbandonedColonyslC

	AlienRuinsslC := []string{}
	AlienRuinsslC = append(AlienRuinsslC, "Traps in the ruins")
	AlienRuinsslC = append(AlienRuinsslC, "Remote location")
	AlienRuinsslC = append(AlienRuinsslC, "Paranoid customs officials")
	CompicationTagMap["AlienRuins"] = AlienRuinsslC

	AlteredHumanityslC := []string{}
	AlteredHumanityslC = append(AlteredHumanityslC, "Alteration is contagious")
	AlteredHumanityslC = append(AlteredHumanityslC, "Alteration is necessary for long-term survival")
	AlteredHumanityslC = append(AlteredHumanityslC, "Locals fear and mistrust non-local humans")
	CompicationTagMap["AlteredHumanity"] = AlteredHumanityslC

	AnarchistsslC := []string{}
	AnarchistsslC = append(AnarchistsslC, "The anarchistic structure is compelled by an external power")
	AnarchistsslC = append(AnarchistsslC, "The anarchy is enabled by currently abundant resources")
	AnarchistsslC = append(AnarchistsslC, "The protecting force that shelters the anarchy is waning")
	CompicationTagMap["Anarchists"] = AnarchistsslC

	AnthropomorphsslC := []string{}
	AnthropomorphsslC = append(AnthropomorphsslC, "The locals consider their shapes a curse from their foolish ancestors")
	AnthropomorphsslC = append(AnthropomorphsslC, "Society is ordered according to animal forms")
	AnthropomorphsslC = append(AnthropomorphsslC, "The locals view normal humans as repulsive or inferior")
	CompicationTagMap["Anthropomorphs"] = AnthropomorphsslC

	Area51slC := []string{}
	Area51slC = append(Area51slC, "The government has a good reason to keep the truth concealed")
	Area51slC = append(Area51slC, "The government ruthlessly oppresses the natives")
	Area51slC = append(Area51slC, "The government is actually composed of offworlders")
	CompicationTagMap["Area51"] = Area51slC

	BadlandsWorldslC := []string{}
	BadlandsWorldslC = append(BadlandsWorldslC, "Radioactivity")
	BadlandsWorldslC = append(BadlandsWorldslC, "Bioweapon traces")
	BadlandsWorldslC = append(BadlandsWorldslC, "Broken terrain")
	BadlandsWorldslC = append(BadlandsWorldslC, "Sudden local plague")
	CompicationTagMap["BadlandsWorld"] = BadlandsWorldslC

	BattlegroundslC := []string{}
	BattlegroundslC = append(BattlegroundslC, "The war just ended as both sides are leaving")
	BattlegroundslC = append(BattlegroundslC, "The natives somehow brought this on themselves")
	BattlegroundslC = append(BattlegroundslC, "A small group of natives profit tremendously from the fighting")
	CompicationTagMap["Battleground"] = BattlegroundslC

	BeastmastersslC := []string{}
	BeastmastersslC = append(BeastmastersslC, "The “animals” are very heavily gengineered humans")
	BeastmastersslC = append(BeastmastersslC, "The animals actually run the society")
	BeastmastersslC = append(BeastmastersslC, "The animals have the same rights as humans")
	CompicationTagMap["Beastmasters"] = BeastmastersslC

	BubbleCitiesslC := []string{}
	BubbleCitiesslC = append(BubbleCitiesslC, "Bubble rupture")
	BubbleCitiesslC = append(BubbleCitiesslC, "Failing atmosphere reprocessor")
	BubbleCitiesslC = append(BubbleCitiesslC, "Native revolt against officials")
	BubbleCitiesslC = append(BubbleCitiesslC, "All-seeing surveillance cameras")
	CompicationTagMap["BubbleCities"] = BubbleCitiesslC

	CheapLifeslC := []string{}
	CheapLifeslC = append(CheapLifeslC, "Radiation or local diseases ensure all locals die before twenty-five years of age")
	CheapLifeslC = append(CheapLifeslC, "Tech ensures that death is just an annoyance")
	CheapLifeslC = append(CheapLifeslC, "Locals are totally convinced of a blissful afterlife")
	CompicationTagMap["CheapLife"] = CheapLifeslC

	CivilWarslC := []string{}
	CivilWarslC = append(CivilWarslC, "The front rolls over the group")
	CivilWarslC = append(CivilWarslC, "Famine strikes")
	CivilWarslC = append(CivilWarslC, "Bandit infestations are in the way")
	CompicationTagMap["CivilWar"] = CivilWarslC

	ColdWarslC := []string{}
	ColdWarslC = append(ColdWarslC, "Police sweep")
	ColdWarslC = append(ColdWarslC, "Low-level skirmishing")
	ColdWarslC = append(ColdWarslC, "“Red scare”")
	CompicationTagMap["ColdWar"] = ColdWarslC

	ColonizedPopulationslC := []string{}
	ColonizedPopulationslC = append(ColonizedPopulationslC, "Natives won’t talk to offworlders")
	ColonizedPopulationslC = append(ColonizedPopulationslC, "Colonial repression")
	ColonizedPopulationslC = append(ColonizedPopulationslC, "Misunderstood local customs")
	CompicationTagMap["ColonizedPopulation"] = ColonizedPopulationslC

	CulturalPowerslC := []string{}
	CulturalPowerslC = append(CulturalPowerslC, "The art is slowly lethal to its masters")
	CulturalPowerslC = append(CulturalPowerslC, "The art is mentally or physically addictive")
	CulturalPowerslC = append(CulturalPowerslC, "The art is a fragment of ancient technical or military science")
	CompicationTagMap["CulturalPower"] = CulturalPowerslC

	CybercommunistsslC := []string{}
	CybercommunistsslC = append(CybercommunistsslC, "The pretech planning computers are breaking down")
	CybercommunistsslC = append(CybercommunistsslC, "The planning only works because the locals have been mentally or physically altered")
	CybercommunistsslC = append(CybercommunistsslC, "The planning computers can’t handle the increasing population within the system")
	CompicationTagMap["Cybercommunists"] = CybercommunistsslC

	CyborgsslC := []string{}
	CyborgsslC = append(CyborgsslC, "The powerful and dangerous come here often for cutting-edge implants")
	CyborgsslC = append(CyborgsslC, "The cyber has some universal negative side-effect")
	CyborgsslC = append(CyborgsslC, "Cyber and those implanted with it are forbidden to leave the planet as a tech security measure")
	CompicationTagMap["Cyborgs"] = CyborgsslC

	CyclicalDoomslC := []string{}
	CyclicalDoomslC = append(CyclicalDoomslC, "The cycles really are a myth of the authorities")
	CyclicalDoomslC = append(CyclicalDoomslC, "The cycles are controlled by alien constructs")
	CyclicalDoomslC = append(CyclicalDoomslC, "An outside power is interfering with preparation")
	CompicationTagMap["CyclicalDoom"] = CyclicalDoomslC

	DesertWorldslC := []string{}
	DesertWorldslC = append(DesertWorldslC, "Sandstorms")
	DesertWorldslC = append(DesertWorldslC, "Water supply failure")
	DesertWorldslC = append(DesertWorldslC, "Native warfare over water rights")
	CompicationTagMap["DesertWorld"] = DesertWorldslC

	DoomedWorldslC := []string{}
	DoomedWorldslC = append(DoomedWorldslC, "The doom is false or won’t actually kill everyone")
	DoomedWorldslC = append(DoomedWorldslC, "The doom was intentionally triggered by someone")
	DoomedWorldslC = append(DoomedWorldslC, "Mass escape is possible if warring groups can somehow be brought to cooperate")
	CompicationTagMap["DoomedWorld"] = DoomedWorldslC

	DyingRaceslC := []string{}
	DyingRaceslC = append(DyingRaceslC, "The dying culture’s values were monstrous")
	DyingRaceslC = append(DyingRaceslC, "The race’s death is somehow necessary to prevent some grand catastrophe")
	DyingRaceslC = append(DyingRaceslC, "The race is somehow convinced they deserve this fate")
	CompicationTagMap["DyingRace"] = DyingRaceslC

	EugenicCultslC := []string{}
	EugenicCultslC = append(EugenicCultslC, "The altered cultists look human")
	EugenicCultslC = append(EugenicCultslC, "The locals are terrified of any unusual physical appearance")
	EugenicCultslC = append(EugenicCultslC, "The genetic modifications- and drawbacks- are contagious with long exposure")
	CompicationTagMap["EugenicCult"] = EugenicCultslC

	ExchangeConsulateslC := []string{}
	ExchangeConsulateslC = append(ExchangeConsulateslC, "The local Consulate has been corrupted")
	ExchangeConsulateslC = append(ExchangeConsulateslC, "the Consulate is cut off from its funds")
	ExchangeConsulateslC = append(ExchangeConsulateslC, "A powerful debtor refuses to pay")
	CompicationTagMap["ExchangeConsulate"] = ExchangeConsulateslC

	FallenHegemonslC := []string{}
	FallenHegemonslC = append(FallenHegemonslC, "The hegemon’s rule was enlightened and fair")
	FallenHegemonslC = append(FallenHegemonslC, "It collapsed due to its own internal strife rather than external resistance")
	FallenHegemonslC = append(FallenHegemonslC, "It pretends that nothing has happened to its power")
	FallenHegemonslC = append(FallenHegemonslC, "It’s been counter-colonized by vengeful outsiders")
	CompicationTagMap["FallenHegemon"] = FallenHegemonslC

	FeralWorldslC := []string{}
	FeralWorldslC = append(FeralWorldslC, "Horrific local “celebration”")
	FeralWorldslC = append(FeralWorldslC, "Inexplicable and repugnant social rules")
	FeralWorldslC = append(FeralWorldslC, "Taboo zones and people")
	CompicationTagMap["FeralWorld"] = FeralWorldslC

	FlyingCitiesslC := []string{}
	FlyingCitiesslC = append(FlyingCitiesslC, "Sudden storms")
	FlyingCitiesslC = append(FlyingCitiesslC, "Drastic altitude loss")
	FlyingCitiesslC = append(FlyingCitiesslC, "Rival city attacks")
	FlyingCitiesslC = append(FlyingCitiesslC, "Vital machinery breaks down")
	CompicationTagMap["FlyingCities"] = FlyingCitiesslC

	ForbiddenTechslC := []string{}
	ForbiddenTechslC = append(ForbiddenTechslC, "The maltech is being fabricated by an unbraked AI")
	ForbiddenTechslC = append(ForbiddenTechslC, "The government depends on revenue from maltech sales to offworlders")
	ForbiddenTechslC = append(ForbiddenTechslC, "Citizens insist that it’s not really maltech")
	CompicationTagMap["ForbiddenTech"] = ForbiddenTechslC

	FormerWarriorsslC := []string{}
	FormerWarriorsslC = append(FormerWarriorsslC, "Neighboring worlds want them pacified or dead")
	FormerWarriorsslC = append(FormerWarriorsslC, "They only ever used their arts in self-defense")
	FormerWarriorsslC = append(FormerWarriorsslC, "The source of their gifts has been “turned off” in a reversible way")
	CompicationTagMap["FormerWarriors"] = FormerWarriorsslC

	FreakGeologyslC := []string{}
	FreakGeologyslC = append(FreakGeologyslC, "Local conditions that no one remembers to tell outworlders about")
	FreakGeologyslC = append(FreakGeologyslC, "Lethal weather")
	FreakGeologyslC = append(FreakGeologyslC, "Seismic activity")
	CompicationTagMap["FreakGeology"] = FreakGeologyslC

	FreakWeatherslC := []string{}
	FreakWeatherslC = append(FreakWeatherslC, "The weather itself")
	FreakWeatherslC = append(FreakWeatherslC, "Malfunctioning pretech terraforming engines that cause the weather")
	CompicationTagMap["FreakWeather"] = FreakWeatherslC

	FriendlyFoeslC := []string{}
	FriendlyFoeslC = append(FriendlyFoeslC, "The group actually is as harmless and benevolent as they seem")
	FriendlyFoeslC = append(FriendlyFoeslC, "The group offers a vital service at the cost of moral compromise")
	FriendlyFoeslC = append(FriendlyFoeslC, "The group still feels bonds of affiliation with their hostile brethren")
	CompicationTagMap["FriendlyFoe"] = FriendlyFoeslC

	GoldRushslC := []string{}
	GoldRushslC = append(GoldRushslC, "The strike is a hoax")
	GoldRushslC = append(GoldRushslC, "The strike is of a dangerous toxic substance")
	GoldRushslC = append(GoldRushslC, "Export of the mineral is prohibited by the planetary government")
	GoldRushslC = append(GoldRushslC, "The native aliens live around the strike’s location")
	CompicationTagMap["GoldRush"] = GoldRushslC

	GreatWorkslC := []string{}
	GreatWorkslC = append(GreatWorkslC, "The work is totally hopeless")
	GreatWorkslC = append(GreatWorkslC, "Different factions disagree on what the work is")
	GreatWorkslC = append(GreatWorkslC, "An outside power is determined to thwart the work")
	CompicationTagMap["GreatWork"] = GreatWorkslC

	HatredslC := []string{}
	HatredslC = append(HatredslC, "The characters are wearing or using items from the hated world")
	HatredslC = append(HatredslC, "The characters are known to have done business there")
	HatredslC = append(HatredslC, "The characters “look like” the hated others")
	CompicationTagMap["Hatred"] = HatredslC

	HeavyIndustryslC := []string{}
	HeavyIndustryslC = append(HeavyIndustryslC, "The factories are toxic")
	HeavyIndustryslC = append(HeavyIndustryslC, "The resources extractable at their tech level are running out")
	HeavyIndustryslC = append(HeavyIndustryslC, "The masses require the factory output for survival")
	HeavyIndustryslC = append(HeavyIndustryslC, "The industries’ major output is being obsoleted by offworld tech")
	CompicationTagMap["HeavyIndustry"] = HeavyIndustryslC

	HeavyMiningslC := []string{}
	HeavyMiningslC = append(HeavyMiningslC, "The refinery equipment breaks down")
	HeavyMiningslC = append(HeavyMiningslC, "Tunnel collapse")
	HeavyMiningslC = append(HeavyMiningslC, "Silicate life forms growing in the miners’ lungs")
	CompicationTagMap["HeavyMining"] = HeavyMiningslC

	HivemindslC := []string{}
	HivemindslC = append(HivemindslC, "The hivemind only functions on this world")
	HivemindslC = append(HivemindslC, "The hivemind has strict range limits")
	HivemindslC = append(HivemindslC, "The hivemind has different personality factions")
	HivemindslC = append(HivemindslC, "The hivemind only happens at particular times")
	HivemindslC = append(HivemindslC, "The world is made of semi-sentient drones and a single AI")
	CompicationTagMap["Hivemind"] = HivemindslC

	HolyWarslC := []string{}
	HolyWarslC = append(HolyWarslC, "The targets of the war really are doing something diabolically horrible")
	HolyWarslC = append(HolyWarslC, "The holy war is just a mask for a very traditional casus belli")
	HolyWarslC = append(HolyWarslC, "The leaders don’t want the war won but only prolonged")
	HolyWarslC = append(HolyWarslC, "Both this world and the target of the war are religion-obsessed")
	CompicationTagMap["HolyWar"] = HolyWarslC

	HostileBiosphereslC := []string{}
	HostileBiosphereslC = append(HostileBiosphereslC, "Filter masks fail")
	HostileBiosphereslC = append(HostileBiosphereslC, "Parasitic alien infestation")
	HostileBiosphereslC = append(HostileBiosphereslC, "Crop greenhouses lose bio-integrity")
	CompicationTagMap["HostileBiosphere"] = HostileBiosphereslC

	HostileSpaceslC := []string{}
	HostileSpaceslC = append(HostileSpaceslC, "The natives believe the danger is divine chastisement")
	HostileSpaceslC = append(HostileSpaceslC, "The natives blame outworlders for the danger")
	HostileSpaceslC = append(HostileSpaceslC, "The native elite profit from the danger in some way")
	CompicationTagMap["HostileSpace"] = HostileSpaceslC

	ImmortalsslC := []string{}
	ImmortalsslC = append(ImmortalsslC, "Immortality requires doing something that outsiders can’t or won’t willingly do")
	ImmortalsslC = append(ImmortalsslC, "The immortality ends if they leave the world")
	ImmortalsslC = append(ImmortalsslC, "Death is the punishment for even minor crimes")
	ImmortalsslC = append(ImmortalsslC, "Immortals must die or go offworld after a certain span")
	ImmortalsslC = append(ImmortalsslC, "Immortality has brutal side-effects")
	CompicationTagMap["Immortals"] = ImmortalsslC

	LocalSpecialtyslC := []string{}
	LocalSpecialtyslC = append(LocalSpecialtyslC, "The specialty is repugnant in nature")
	LocalSpecialtyslC = append(LocalSpecialtyslC, "The crafters refuse to sell to offworlders")
	LocalSpecialtyslC = append(LocalSpecialtyslC, "The specialty is made in a remote")
	LocalSpecialtyslC = append(LocalSpecialtyslC, "dangerous place")
	LocalSpecialtyslC = append(LocalSpecialtyslC, "The crafters don’t want to make the specialty any more")
	CompicationTagMap["LocalSpecialty"] = LocalSpecialtyslC

	LocalTechslC := []string{}
	LocalTechslC = append(LocalTechslC, "The tech is unreliable")
	LocalTechslC = append(LocalTechslC, "The tech only works on this world")
	LocalTechslC = append(LocalTechslC, "The tech has poorly-understood side effects")
	LocalTechslC = append(LocalTechslC, "The tech is alien in nature.")
	CompicationTagMap["LocalTech"] = LocalTechslC

	MajorSpaceyardslC := []string{}
	MajorSpaceyardslC = append(MajorSpaceyardslC, "The spaceyard is an alien relic")
	MajorSpaceyardslC = append(MajorSpaceyardslC, "The spaceyard is burning out from overuse")
	MajorSpaceyardslC = append(MajorSpaceyardslC, "The spaceyard is alive")
	MajorSpaceyardslC = append(MajorSpaceyardslC, "The spaceyard relies on maltech to function")
	CompicationTagMap["MajorSpaceyard"] = MajorSpaceyardslC

	MandarinateslC := []string{}
	MandarinateslC = append(MandarinateslC, "The test is totally unrelated to necessary governing skills")
	MandarinateslC = append(MandarinateslC, "The test was very pertinent in the past but tech or culture has changed")
	MandarinateslC = append(MandarinateslC, "The test is for a skill that is vital to maintaining society but irrelevant to day-to-day governance")
	MandarinateslC = append(MandarinateslC, "The test is a sham and passage is based on wealth or influence")
	CompicationTagMap["Mandarinate"] = MandarinateslC

	MandateBaseslC := []string{}
	MandateBaseslC = append(MandateBaseslC, "The monitoring system forces the locals to behave in aggressive ways toward “rebel” worlds")
	MandateBaseslC = append(MandateBaseslC, "The monitoring system severely hinders offworld use of their tech")
	MandateBaseslC = append(MandateBaseslC, "The original colonists are all dead and have been replaced by outsiders who don’t understand all the details")
	CompicationTagMap["MandateBase"] = MandateBaseslC

	ManeatersslC := []string{}
	ManeatersslC = append(ManeatersslC, "Local food or environmental conditions make human consumption grimly necessary")
	ManeatersslC = append(ManeatersslC, "The locals farm human beings")
	ManeatersslC = append(ManeatersslC, "Outsiders are expected to join in the custom")
	ManeatersslC = append(ManeatersslC, "The custom is totally unnecessary but jealously maintained by the people")
	CompicationTagMap["Maneaters"] = ManeatersslC

	MegacorpsslC := []string{}
	MegacorpsslC = append(MegacorpsslC, "The megacorps are the only source of something vital to life on this world")
	MegacorpsslC = append(MegacorpsslC, "An autonomous Mandate system acts to punish excessively overt violence")
	MegacorpsslC = append(MegacorpsslC, "The megacorps are struggling against much more horrible national governments")
	CompicationTagMap["Megacorps"] = MegacorpsslC

	MercenariesslC := []string{}
	MercenariesslC = append(MercenariesslC, "The mercenaries are all that stand between the locals and a hungry imperial power")
	MercenariesslC = append(MercenariesslC, "The mercenaries are remnants of a former official army")
	MercenariesslC = append(MercenariesslC, "The mercenaries hardly ever actually fight as compared to taking bribes to walk away")
	CompicationTagMap["Mercenaries"] = MercenariesslC

	MisandryMisogynyslC := []string{}
	MisandryMisogynyslC = append(MisandryMisogynyslC, "The oppressed gender is restive against the customs")
	MisandryMisogynyslC = append(MisandryMisogynyslC, "The oppressed gender largely supports the customs")
	MisandryMisogynyslC = append(MisandryMisogynyslC, "The customs relate to some physical quality of the world")
	MisandryMisogynyslC = append(MisandryMisogynyslC, "The oppressed gender has had maltech gengineering done to “tame” them.")
	CompicationTagMap["MisandryMisogyny"] = MisandryMisogynyslC

	NightWorldslC := []string{}
	NightWorldslC = append(NightWorldslC, "Daylight comes as a cataclysmic event at very long intervals")
	NightWorldslC = append(NightWorldslC, "Light causes very dangerous reactions in native life or chemicals here")
	NightWorldslC = append(NightWorldslC, "The locals have been gengineered to exist without sight")
	CompicationTagMap["NightWorld"] = NightWorldslC

	MinimalContactslC := []string{}
	MinimalContactslC = append(MinimalContactslC, "The locals carry a disease harmless to them and lethal to outsiders")
	MinimalContactslC = append(MinimalContactslC, "The locals hide dark purposes from offworlders")
	MinimalContactslC = append(MinimalContactslC, "The locals have something desperately needed but won’t bring it into the treaty port")
	CompicationTagMap["MinimalContact"] = MinimalContactslC

	NomadsslC := []string{}
	NomadsslC = append(NomadsslC, "An irresistibly lethal swarm of native life forces locals to move regularly")
	NomadsslC = append(NomadsslC, "Ancient defense systems destroy too-long-stationary communities")
	NomadsslC = append(NomadsslC, "Local chemical patches require careful balancing of exposure times to avoid side effects")
	CompicationTagMap["Nomads"] = NomadsslC

	OceanicWorldslC := []string{}
	OceanicWorldslC = append(OceanicWorldslC, "The liquid flux confuses grav engines too badly for them to function on this world")
	OceanicWorldslC = append(OceanicWorldslC, "Sea is corrosive or toxic")
	OceanicWorldslC = append(OceanicWorldslC, "The seas are wracked by regular storms")
	CompicationTagMap["OceanicWorld"] = OceanicWorldslC

	OutofContactslC := []string{}
	OutofContactslC = append(OutofContactslC, "Automatic defenses fire on ships that try to take off")
	OutofContactslC = append(OutofContactslC, "The natives want to stay out of contact")
	OutofContactslC = append(OutofContactslC, "The natives are highly vulnerable to offworld diseases")
	OutofContactslC = append(OutofContactslC, "The native language is completely unlike any known to the group")
	CompicationTagMap["OutofContact"] = OutofContactslC

	OutpostWorldslC := []string{}
	OutpostWorldslC = append(OutpostWorldslC, "The alien ruin defense systems are waking up")
	OutpostWorldslC = append(OutpostWorldslC, "Atmospheric disturbances trap the group inside the outpost for a month")
	OutpostWorldslC = append(OutpostWorldslC, "Pirates raid the outpost")
	OutpostWorldslC = append(OutpostWorldslC, "The crew have become converts to a strange set of beliefs")
	CompicationTagMap["OutpostWorld"] = OutpostWorldslC

	PerimeterAgencyslC := []string{}
	PerimeterAgencyslC = append(PerimeterAgencyslC, "The local Agency has gone rogue and now uses maltech")
	PerimeterAgencyslC = append(PerimeterAgencyslC, "The Agency archives have been compromised")
	PerimeterAgencyslC = append(PerimeterAgencyslC, "The Agency has been targeted by a maltech-using organization")
	PerimeterAgencyslC = append(PerimeterAgencyslC, "The Agency’s existence is unknown to the locals")
	CompicationTagMap["PerimeterAgency"] = PerimeterAgencyslC

	PilgrimageSiteslC := []string{}
	PilgrimageSiteslC = append(PilgrimageSiteslC, "The site is actually a fake")
	PilgrimageSiteslC = append(PilgrimageSiteslC, "The site is run by corrupt and venal keepers")
	PilgrimageSiteslC = append(PilgrimageSiteslC, "A natural disaster threatens the site")
	CompicationTagMap["PilgrimageSite"] = PilgrimageSiteslC

	PleasureWorldslC := []string{}
	PleasureWorldslC = append(PleasureWorldslC, "A deeply repugnant pleasure is offered here by a culture that sees nothing wrong with it")
	PleasureWorldslC = append(PleasureWorldslC, "Certain pleasures here are dangerously addictive")
	PleasureWorldslC = append(PleasureWorldslC, "The prices here can involve enslavement or death")
	PleasureWorldslC = append(PleasureWorldslC, "The world has been seized and exploited by an imperial power")
	CompicationTagMap["PleasureWorld"] = PleasureWorldslC

	PoliceStateslC := []string{}
	PoliceStateslC = append(PoliceStateslC, "The natives largely believe in the righteousness of the state")
	PoliceStateslC = append(PoliceStateslC, "The police state is automated and its “rulers” can’t shut it off")
	PoliceStateslC = append(PoliceStateslC, "The leaders foment a pogrom against “offworlder spies”.")
	CompicationTagMap["PoliceState"] = PoliceStateslC

	PostScarcityslC := []string{}
	PostScarcityslC = append(PostScarcityslC, "The tech causes serious side-effects on those who take advantage of it")
	PostScarcityslC = append(PostScarcityslC, "The tech is breaking down")
	PostScarcityslC = append(PostScarcityslC, "The population is growing too large")
	PostScarcityslC = append(PostScarcityslC, "The tech produces only certain things in abundance")
	CompicationTagMap["PostScarcity"] = PostScarcityslC

	PreceptorArchiveslC := []string{}
	PreceptorArchiveslC = append(PreceptorArchiveslC, "The local Archive has taken a very religious and mystical attitude toward their teaching")
	PreceptorArchiveslC = append(PreceptorArchiveslC, "The Archive has maintained some replicable pretech science")
	PreceptorArchiveslC = append(PreceptorArchiveslC, "The Archive has been corrupted and their teaching is incorrect")
	CompicationTagMap["PreceptorArchive"] = PreceptorArchiveslC

	PretechCultistsslC := []string{}
	PretechCultistsslC = append(PretechCultistsslC, "The cultists can actually replicate certain forms of pretech")
	PretechCultistsslC = append(PretechCultistsslC, "The cultists abhor use of the devices as “presumption on the holy”")
	PretechCultistsslC = append(PretechCultistsslC, "The cultists mistake the party’s belongings for pretech")
	CompicationTagMap["PretechCultists"] = PretechCultistsslC

	PrimitiveAliensslC := []string{}
	PrimitiveAliensslC = append(PrimitiveAliensslC, "The alien numbers are huge and can overwhelm the humans whenever they so choose")
	PrimitiveAliensslC = append(PrimitiveAliensslC, "One group is trying to use the other to kill their political opponents")
	PrimitiveAliensslC = append(PrimitiveAliensslC, "The aliens are incomprehensibly strange")
	PrimitiveAliensslC = append(PrimitiveAliensslC, "One side commits an atrocity")
	CompicationTagMap["PrimitiveAliens"] = PrimitiveAliensslC

	PrisonPlanetslC := []string{}
	PrisonPlanetslC = append(PrisonPlanetslC, "Departure permits are a precious currency")
	PrisonPlanetslC = append(PrisonPlanetslC, "The prison industry still makes valuable pretech devices")
	PrisonPlanetslC = append(PrisonPlanetslC, "Gangs have metamorphosed into governments")
	PrisonPlanetslC = append(PrisonPlanetslC, "The local nobility descended from the prison staff")
	CompicationTagMap["PrisonPlanet"] = PrisonPlanetslC

	PsionicsAcademyslC := []string{}
	PsionicsAcademyslC = append(PsionicsAcademyslC, "The academy curriculum kills a significant percentage of students")
	PsionicsAcademyslC = append(PsionicsAcademyslC, "The faculty use students as research subjects")
	PsionicsAcademyslC = append(PsionicsAcademyslC, "The students are indoctrinated as sleeper agents")
	PsionicsAcademyslC = append(PsionicsAcademyslC, "The local natives hate the academy")
	PsionicsAcademyslC = append(PsionicsAcademyslC, "The academy is part of a religion.")
	CompicationTagMap["PsionicsAcademy"] = PsionicsAcademyslC

	PsionicsFearslC := []string{}
	PsionicsFearslC = append(PsionicsFearslC, "Psychic potential is much more common here")
	PsionicsFearslC = append(PsionicsFearslC, "Some tech is mistaken as psitech")
	PsionicsFearslC = append(PsionicsFearslC, "Natives believe certain rituals and customs can protect them from psychic powers")
	CompicationTagMap["PsionicsFear"] = PsionicsFearslC

	PsionicsWorshipslC := []string{}
	PsionicsWorshipslC = append(PsionicsWorshipslC, "The psychic training is imperfect")
	PsionicsWorshipslC = append(PsionicsWorshipslC, "and the psychics all show significant mental illness")
	PsionicsWorshipslC = append(PsionicsWorshipslC, "The psychics have developed a unique discipline")
	PsionicsWorshipslC = append(PsionicsWorshipslC, "The will of a psychic is law")
	PsionicsWorshipslC = append(PsionicsWorshipslC, "Psychics in the party are forcibly kidnapped for “enlightening”.")
	CompicationTagMap["PsionicsWorship"] = PsionicsWorshipslC

	QuarantinedWorldslC := []string{}
	QuarantinedWorldslC = append(QuarantinedWorldslC, "The natives want to remain isolated")
	QuarantinedWorldslC = append(QuarantinedWorldslC, "The quarantine is enforced by an ancient alien installation")
	QuarantinedWorldslC = append(QuarantinedWorldslC, "The world is rife with maltech abominations")
	QuarantinedWorldslC = append(QuarantinedWorldslC, "The blockade is meant to starve everyone on the barren world.")
	CompicationTagMap["QuarantinedWorld"] = QuarantinedWorldslC

	RadioactiveWorldslC := []string{}
	RadioactiveWorldslC = append(RadioactiveWorldslC, "The radioactivity is steadily growing worse")
	RadioactiveWorldslC = append(RadioactiveWorldslC, "The planet’s medical resources break down")
	RadioactiveWorldslC = append(RadioactiveWorldslC, "The radioactivity has inexplicable effects on living creatures")
	RadioactiveWorldslC = append(RadioactiveWorldslC, "The radioactivity is the product of a malfunctioning pretech manufactory.")
	CompicationTagMap["RadioactiveWorld"] = RadioactiveWorldslC

	RefugeesslC := []string{}
	RefugeesslC = append(RefugeesslC, "The xenophobes are right that the refugees are taking over")
	RefugeesslC = append(RefugeesslC, "The refugees are right that the xenophobes want them out or dead")
	RefugeesslC = append(RefugeesslC, "Both are right")
	RefugeesslC = append(RefugeesslC, "Outside powers are using the refugees to destabilize an enemy government")
	RefugeesslC = append(RefugeesslC, "Refugee and local cultures are extremely incompatible")
	CompicationTagMap["Refugees"] = RefugeesslC

	RegionalHegemonslC := []string{}
	RegionalHegemonslC = append(RegionalHegemonslC, "The hegemon’s influence is all that’s keeping a murderous war from breaking out on nearby worlds")
	RegionalHegemonslC = append(RegionalHegemonslC, "The hegemon is decaying and losing its control")
	RegionalHegemonslC = append(RegionalHegemonslC, "The government is riddled with spies")
	RegionalHegemonslC = append(RegionalHegemonslC, "The hegemon is genuinely benign")
	CompicationTagMap["RegionalHegemon"] = RegionalHegemonslC

	RestrictiveLawsslC := []string{}
	RestrictiveLawsslC = append(RestrictiveLawsslC, "The laws change regularly in patterns only natives understand")
	RestrictiveLawsslC = append(RestrictiveLawsslC, "The laws forbid some action vital to the party")
	RestrictiveLawsslC = append(RestrictiveLawsslC, "The laws forbid the simple existence of some party members")
	RestrictiveLawsslC = append(RestrictiveLawsslC, "The laws are secret to offworlders")
	CompicationTagMap["RestrictiveLaws"] = RestrictiveLawsslC

	RevanchistsslC := []string{}
	RevanchistsslC = append(RevanchistsslC, "The revanchists’ claim is completely just and reasonable")
	RevanchistsslC = append(RevanchistsslC, "The land is now occupied entirely by heirs of the conquerors")
	RevanchistsslC = append(RevanchistsslC, "Both sides have seized lands the other thinks are theirs")
	CompicationTagMap["Revanchists"] = RevanchistsslC

	RevolutionariesslC := []string{}
	RevolutionariesslC = append(RevolutionariesslC, "The revolutionaries actually do seem likely to put in better rulers")
	RevolutionariesslC = append(RevolutionariesslC, "The revolutionaries are client groups that got out of hand")
	RevolutionariesslC = append(RevolutionariesslC, "The revolutionaries are clearly much worse than the government")
	RevolutionariesslC = append(RevolutionariesslC, "The revolutionaries have no real ideals beyond power and merely pretend to ideology")
	CompicationTagMap["Revolutionaries"] = RevolutionariesslC

	RigidCultureslC := []string{}
	RigidCultureslC = append(RigidCultureslC, "The cultural patterns are enforced by technological aids")
	RigidCultureslC = append(RigidCultureslC, "The culture is run by a secret cabal of manipulators")
	RigidCultureslC = append(RigidCultureslC, "The culture has explicit religious sanction")
	RigidCultureslC = append(RigidCultureslC, "The culture evolved due to important necessities that have since been forgotten")
	CompicationTagMap["RigidCulture"] = RigidCultureslC

	RisingHegemonslC := []string{}
	RisingHegemonslC = append(RisingHegemonslC, "They’re only strong because their neighbors have been weakened")
	RisingHegemonslC = append(RisingHegemonslC, "Their success is based on a fluke resource or pretech find")
	RisingHegemonslC = append(RisingHegemonslC, "They bitterly resent their neighbors as former oppressors")
	CompicationTagMap["RisingHegemon"] = RisingHegemonslC

	RitualCombatslC := []string{}
	RitualCombatslC = append(RitualCombatslC, "The required weapons are strange pretech artifacts")
	RitualCombatslC = append(RitualCombatslC, "Certain classes are forbidden from fighting and require champions")
	RitualCombatslC = append(RitualCombatslC, "Loss doesn’t mean death but it does mean ritual scarring or property loss")
	CompicationTagMap["RitualCombat"] = RitualCombatslC

	RobotsslC := []string{}
	RobotsslC = append(RobotsslC, "The robots are only partially controlled")
	RobotsslC = append(RobotsslC, "The robots are salvaged and originally meant for a much darker use")
	RobotsslC = append(RobotsslC, "The robots require a rare material that the locals fight over")
	RobotsslC = append(RobotsslC, "The robots require the planet’s specific infrastructure so cannot be exported")
	CompicationTagMap["Robots"] = RobotsslC

	SeagoingCitiesslC := []string{}
	SeagoingCitiesslC = append(SeagoingCitiesslC, "The seas are not water")
	SeagoingCitiesslC = append(SeagoingCitiesslC, "The fish schools have vanished and the city faces starvation")
	SeagoingCitiesslC = append(SeagoingCitiesslC, "Terrible storms drive the city into the glacial regions")
	SeagoingCitiesslC = append(SeagoingCitiesslC, "Suicide ships ram the city’s hull")
	CompicationTagMap["SeagoingCities"] = SeagoingCitiesslC

	SealedMenaceslC := []string{}
	SealedMenaceslC = append(SealedMenaceslC, "The menace would bring great wealth along with destruction")
	SealedMenaceslC = append(SealedMenaceslC, "The menace is intelligent")
	SealedMenaceslC = append(SealedMenaceslC, "The natives don’t all believe in the menace")
	CompicationTagMap["SealedMenace"] = SealedMenaceslC

	SecretMastersslC := []string{}
	SecretMastersslC = append(SecretMastersslC, "The secret masters have a benign reason for wanting secrecy")
	SecretMastersslC = append(SecretMastersslC, "The cabal fights openly amongst itself")
	SecretMastersslC = append(SecretMastersslC, "The cabal is recruiting new members")
	CompicationTagMap["SecretMasters"] = SecretMastersslC

	SectariansslC := []string{}
	SectariansslC = append(SectariansslC, "The conflict has more than two sides")
	SectariansslC = append(SectariansslC, "The sectarians hate each other for multiple reasons")
	SectariansslC = append(SectariansslC, "The sectarians must cooperate or else life on this world is imperiled")
	SectariansslC = append(SectariansslC, "The sectarians hate outsiders more than they hate each other")
	SectariansslC = append(SectariansslC, "The differences in sects are incomprehensible to an outsider")
	CompicationTagMap["Sectarians"] = SectariansslC

	SeismicInstabilityslC := []string{}
	SeismicInstabilityslC = append(SeismicInstabilityslC, "The earthquakes are caused by malfunctioning pretech terraformers")
	SeismicInstabilityslC = append(SeismicInstabilityslC, "They’re caused by alien technology")
	SeismicInstabilityslC = append(SeismicInstabilityslC, "They’re restrained by alien technology that is being plundered by offworlders")
	SeismicInstabilityslC = append(SeismicInstabilityslC, "The earthquakes are used to generate enormous amounts of energy.")
	CompicationTagMap["SeismicInstability"] = SeismicInstabilityslC

	ShackledWorldslC := []string{}
	ShackledWorldslC = append(ShackledWorldslC, "The shackles come off for certain brief windows of time")
	ShackledWorldslC = append(ShackledWorldslC, "The locals think the shackles are imposed by God")
	ShackledWorldslC = append(ShackledWorldslC, "An outside power greatly profits from the shackles")
	ShackledWorldslC = append(ShackledWorldslC, "The rulers are exempt from the shackles")
	CompicationTagMap["ShackledWorld"] = ShackledWorldslC

	SocietalDespairslC := []string{}
	SocietalDespairslC = append(SocietalDespairslC, "A massive war discredited all the old values")
	SocietalDespairslC = append(SocietalDespairslC, "Outside powers are working to erode societal confidence for their own benefit")
	SocietalDespairslC = append(SocietalDespairslC, "A local power is profiting greatly from the despair")
	SocietalDespairslC = append(SocietalDespairslC, "The old ways were meant to aid survival on this world and their passing is causing many new woes")
	CompicationTagMap["SocietalDespair"] = SocietalDespairslC

	SoleSupplierslC := []string{}
	SoleSupplierslC = append(SoleSupplierslC, "The substance is slow poison to process")
	SoleSupplierslC = append(SoleSupplierslC, "The substance is created by hostile alien natives")
	SoleSupplierslC = append(SoleSupplierslC, "The substance is very easy to smuggle in usable amounts")
	SoleSupplierslC = append(SoleSupplierslC, "Only the natives have the genes or tech to extract it effectively")
	CompicationTagMap["SoleSupplier"] = SoleSupplierslC

	TabooTreasureslC := []string{}
	TabooTreasureslC = append(TabooTreasureslC, "The treasure is extremely hard to smuggle")
	TabooTreasureslC = append(TabooTreasureslC, "Its use visibly marks a user")
	TabooTreasureslC = append(TabooTreasureslC, "The natives consider it for their personal use only,")
	CompicationTagMap["TabooTreasure"] = TabooTreasureslC

	TerraformFailureslC := []string{}
	TerraformFailureslC = append(TerraformFailureslC, "The engines produced too much of something instead of too little")
	TerraformFailureslC = append(TerraformFailureslC, "The engines were hijacked by aliens with different preferences")
	TerraformFailureslC = append(TerraformFailureslC, "It was discovered that an Earth-like environment would eventually cause a catastrophic disaster")
	CompicationTagMap["TerraformFailure"] = TerraformFailureslC

	TheocracyslC := []string{}
	TheocracyslC = append(TheocracyslC, "The theocracy actually works well")
	TheocracyslC = append(TheocracyslC, "The theocracy is decadent and hated by the common folk")
	TheocracyslC = append(TheocracyslC, "The theocracy is divided into mutually hostile sects")
	TheocracyslC = append(TheocracyslC, "The theocracy is led by aliens")
	CompicationTagMap["Theocracy"] = TheocracyslC

	TombWorldslC := []string{}
	TombWorldslC = append(TombWorldslC, "The ruins are full of booby-traps left by the final inhabitants")
	TombWorldslC = append(TombWorldslC, "The world’s atmosphere quickly degrades anything in an opened building")
	TombWorldslC = append(TombWorldslC, "A handful of desperate natives survived the Silence")
	TombWorldslC = append(TombWorldslC, "The structures are unstable and collapsing")
	CompicationTagMap["TombWorld"] = TombWorldslC

	TradeHubslC := []string{}
	TradeHubslC = append(TradeHubslC, "An outworlder faction schemes to seize the trade hub")
	TradeHubslC = append(TradeHubslC, "Saboteurs seek to blow up a rival’s warehouses")
	TradeHubslC = append(TradeHubslC, "Enemies are blockading the trade routes")
	TradeHubslC = append(TradeHubslC, "Pirates lace the hub with spies")
	CompicationTagMap["TradeHub"] = TradeHubslC

	TyrannyslC := []string{}
	TyrannyslC = append(TyrannyslC, "The tyrant rules with vastly superior technology")
	TyrannyslC = append(TyrannyslC, "The tyrant is a figurehead for a cabal of powerful men and women")
	TyrannyslC = append(TyrannyslC, "The people are resigned to their suffering")
	TyrannyslC = append(TyrannyslC, "The tyrant is hostile to “meddlesome outworlders”.")
	CompicationTagMap["Tyranny"] = TyrannyslC

	UnbrakedAIslC := []string{}
	UnbrakedAIslC = append(UnbrakedAIslC, "The AI’s presence is unknown to the locals")
	UnbrakedAIslC = append(UnbrakedAIslC, "The locals depend on the AI for some vital service")
	UnbrakedAIslC = append(UnbrakedAIslC, "The AI appears to be harmless")
	UnbrakedAIslC = append(UnbrakedAIslC, "The AI has fixated on the group’s ship’s computer")
	UnbrakedAIslC = append(UnbrakedAIslC, "The AI wants transport offworld")
	CompicationTagMap["UnbrakedAI"] = UnbrakedAIslC

	UrbanizedSurfaceslC := []string{}
	UrbanizedSurfaceslC = append(UrbanizedSurfaceslC, "The urban blocks are needed to survive the environment")
	UrbanizedSurfaceslC = append(UrbanizedSurfaceslC, "The blocks were part of an ancient device of world-spanning size")
	UrbanizedSurfaceslC = append(UrbanizedSurfaceslC, "The blocks require constant maintenance to avoid dangerous types of decay")
	CompicationTagMap["UrbanizedSurface"] = UrbanizedSurfaceslC

	UtopiaslC := []string{}
	UtopiaslC = append(UtopiaslC, "The natives really are deeply and contentedly happy with their altered lot")
	UtopiaslC = append(UtopiaslC, "The utopia produces something that attracts others")
	UtopiaslC = append(UtopiaslC, "The utopia works on converting outsiders through persuasion and generosity")
	UtopiaslC = append(UtopiaslC, "The utopia involves some sacrifice that’s horrifying to non-members")
	CompicationTagMap["Utopia"] = UtopiaslC

	WarlordsslC := []string{}
	WarlordsslC = append(WarlordsslC, "The warlords are willing to cooperate to fight mutual threats")
	WarlordsslC = append(WarlordsslC, "The warlords favor specific religions or races over others")
	WarlordsslC = append(WarlordsslC, "The warlords are using substantially more sophisticated tech than others")
	WarlordsslC = append(WarlordsslC, "Some of the warlords are better rulers than the government")
	CompicationTagMap["Warlords"] = WarlordsslC

	XenophilesslC := []string{}
	XenophilesslC = append(XenophilesslC, "The enthusiasm is due to alien psionics or tech")
	XenophilesslC = append(XenophilesslC, "The enthusiasm is based on a lie")
	XenophilesslC = append(XenophilesslC, "The aliens strongly dislike their “groupies”")
	XenophilesslC = append(XenophilesslC, "The aliens feel obliged to rule humanity for its own good")
	XenophilesslC = append(XenophilesslC, "Humans badly misunderstand the aliens")
	CompicationTagMap["Xenophiles"] = XenophilesslC

	XenophobesslC := []string{}
	XenophobesslC = append(XenophobesslC, "The natives are symptomless carriers of a contagious and dangerous disease")
	XenophobesslC = append(XenophobesslC, "The natives are exceptionally vulnerable to offworld diseases")
	XenophobesslC = append(XenophobesslC, "The natives require elaborate purification rituals after speaking to an offworlder or touching them")
	XenophobesslC = append(XenophobesslC, "The local ruler has forbidden any mercantile dealings with outworlders")
	CompicationTagMap["Xenophobes"] = XenophobesslC

	ZombiesslC := []string{}
	ZombiesslC = append(ZombiesslC, "The zombies retain human intelligence")
	ZombiesslC = append(ZombiesslC, "The zombies can be cured")
	ZombiesslC = append(ZombiesslC, "The process is voluntary among devotees")
	ZombiesslC = append(ZombiesslC, "The condition is infectious")
	CompicationTagMap["Zombies"] = ZombiesslC
	return CompicationTagMap
}

func createThingsMap() map[string][]string {
	ThingTagMap := make(map[string][]string)
	AbandonedColonyslT := []string{}
	AbandonedColonyslT = append(AbandonedColonyslT, "Long-lost property deeds")
	AbandonedColonyslT = append(AbandonedColonyslT, "Relic stolen by the colonists when they left")
	AbandonedColonyslT = append(AbandonedColonyslT, "Historical record of the colonization attempt")
	ThingTagMap["AbandonedColony"] = AbandonedColonyslT

	AlienRuinsslT := []string{}
	AlienRuinsslT = append(AlienRuinsslT, "Precious alien artifacts")
	AlienRuinsslT = append(AlienRuinsslT, "Objects left with the remains of a prior unsuccessful expedition")
	AlienRuinsslT = append(AlienRuinsslT, "Untranslated alien texts")
	AlienRuinsslT = append(AlienRuinsslT, "Untouched hidden ruins")
	ThingTagMap["AlienRuins"] = AlienRuinsslT

	AlteredHumanityslT := []string{}
	AlteredHumanityslT = append(AlteredHumanityslT, "Original pretech mutagenic equipment")
	AlteredHumanityslT = append(AlteredHumanityslT, "Valuable biological byproduct from the mutants")
	AlteredHumanityslT = append(AlteredHumanityslT, "“Cure” for the altered genes")
	AlteredHumanityslT = append(AlteredHumanityslT, "Record of the original colonial genotypes")
	ThingTagMap["AlteredHumanity"] = AlteredHumanityslT

	AnarchistsslT := []string{}
	AnarchistsslT = append(AnarchistsslT, "A macguffin that would let the possessor enforce their rule on others")
	AnarchistsslT = append(AnarchistsslT, "A vital resource needed to preserve general liberty")
	AnarchistsslT = append(AnarchistsslT, "Tech forbidden as disruptive to the social order")
	ThingTagMap["Anarchists"] = AnarchistsslT

	AnthropomorphsslT := []string{}
	AnthropomorphsslT = append(AnthropomorphsslT, "Pretech gengineering tech")
	AnthropomorphsslT = append(AnthropomorphsslT, "A “cure” that may not be wanted")
	AnthropomorphsslT = append(AnthropomorphsslT, "Sacred feral totem")
	ThingTagMap["Anthropomorphs"] = AnthropomorphsslT

	Area51slT := []string{}
	Area51slT = append(Area51slT, "Elaborate spy devices")
	Area51slT = append(Area51slT, "Memory erasure tech")
	Area51slT = append(Area51slT, "Possessions of the last offworlder who decided to spread the truth")
	ThingTagMap["Area51"] = Area51slT

	BadlandsWorldslT := []string{}
	BadlandsWorldslT = append(BadlandsWorldslT, "Maltech research core")
	BadlandsWorldslT = append(BadlandsWorldslT, "Functional pretech weaponry")
	BadlandsWorldslT = append(BadlandsWorldslT, "An uncontaminated well")
	ThingTagMap["BadlandsWorld"] = BadlandsWorldslT

	BattlegroundslT := []string{}
	BattlegroundslT = append(BattlegroundslT, "A cache of the resource the invaders seek")
	BattlegroundslT = append(BattlegroundslT, "Abandoned prototype military gear")
	BattlegroundslT = append(BattlegroundslT, "Precious spy intelligence lost by someone")
	ThingTagMap["Battleground"] = BattlegroundslT

	BeastmastersslT := []string{}
	BeastmastersslT = append(BeastmastersslT, "Tech used to alter animal life")
	BeastmastersslT = append(BeastmastersslT, "A plague vial that could wipe out the animals")
	BeastmastersslT = append(BeastmastersslT, "A pretech device that can perform a wonder if operated by a beast")
	ThingTagMap["Beastmasters"] = BeastmastersslT

	BubbleCitiesslT := []string{}
	BubbleCitiesslT = append(BubbleCitiesslT, "Pretech habitat technology")
	BubbleCitiesslT = append(BubbleCitiesslT, "Valuable industrial products")
	BubbleCitiesslT = append(BubbleCitiesslT, "Master key codes to a city’s security system")
	ThingTagMap["BubbleCities"] = BubbleCitiesslT

	CheapLifeslT := []string{}
	CheapLifeslT = append(CheapLifeslT, "Device that revives or re-embodies the dead")
	CheapLifeslT = append(CheapLifeslT, "Maltech engine fueled by human life")
	CheapLifeslT = append(CheapLifeslT, "Priceless treasure held by a now-dead owner")
	ThingTagMap["CheapLife"] = CheapLifeslT

	CivilWarslT := []string{}
	CivilWarslT = append(CivilWarslT, "Ammo dump")
	CivilWarslT = append(CivilWarslT, "Military cache")
	CivilWarslT = append(CivilWarslT, "Treasure buried for after the war")
	CivilWarslT = append(CivilWarslT, "Secret war plans")
	ThingTagMap["CivilWar"] = CivilWarslT

	ColdWarslT := []string{}
	ColdWarslT = append(ColdWarslT, "List of traitors in government")
	ColdWarslT = append(ColdWarslT, "secret military plans")
	ColdWarslT = append(ColdWarslT, "Huge cache of weapons built up in preparation for war")
	ThingTagMap["ColdWar"] = ColdWarslT

	ColonizedPopulationslT := []string{}
	ColonizedPopulationslT = append(ColonizedPopulationslT, "Relic of the resistance movement")
	ColonizedPopulationslT = append(ColonizedPopulationslT, "List of collaborators")
	ColonizedPopulationslT = append(ColonizedPopulationslT, "Precious substance extracted by colonial labor")
	ThingTagMap["ColonizedPopulation"] = ColonizedPopulationslT

	CulturalPowerslT := []string{}
	CulturalPowerslT = append(CulturalPowerslT, "The instrument of a legendary master")
	CulturalPowerslT = append(CulturalPowerslT, "The only copy of a dead master’s opus")
	CulturalPowerslT = append(CulturalPowerslT, "Proof of intellectual property ownership")
	ThingTagMap["CulturalPower"] = CulturalPowerslT

	CybercommunistsslT := []string{}
	CybercommunistsslT = append(CybercommunistsslT, "Planning node computer")
	CybercommunistsslT = append(CybercommunistsslT, "Wildly destabilizing commodity that can’t be factored into plans")
	CybercommunistsslT = append(CybercommunistsslT, "A tremendous store of valuables made by accident")
	ThingTagMap["Cybercommunists"] = CybercommunistsslT

	CyborgsslT := []string{}
	CyborgsslT = append(CyborgsslT, "Unique prototype cyber implant")
	CyborgsslT = append(CyborgsslT, "Secret research files")
	CyborgsslT = append(CyborgsslT, "A virus that debilitates cyborgs")
	CyborgsslT = append(CyborgsslT, "A cache of critically-needed therapeutic cyber")
	ThingTagMap["Cyborgs"] = CyborgsslT

	CyclicalDoomslT := []string{}
	CyclicalDoomslT = append(CyclicalDoomslT, "A lost cache of ancient treasures")
	CyclicalDoomslT = append(CyclicalDoomslT, "Tech or archives that will pinpoint the cycle’s timing")
	CyclicalDoomslT = append(CyclicalDoomslT, "Keycodes to bypass an ancient vault’s security")
	ThingTagMap["CyclicalDoom"] = CyclicalDoomslT

	DesertWorldslT := []string{}
	DesertWorldslT = append(DesertWorldslT, "Enormous water reservoir")
	DesertWorldslT = append(DesertWorldslT, "Map of hidden wells")
	DesertWorldslT = append(DesertWorldslT, "Pretech rainmaking equipment")
	ThingTagMap["DesertWorld"] = DesertWorldslT

	DoomedWorldslT := []string{}
	DoomedWorldslT = append(DoomedWorldslT, "Clearance for a ship to leave the planet")
	DoomedWorldslT = append(DoomedWorldslT, "A cache of priceless cultural artifacts")
	DoomedWorldslT = append(DoomedWorldslT, "The life savings of someone trying to buy passage out")
	DoomedWorldslT = append(DoomedWorldslT, "Data that would prove to the public the end is nigh")
	ThingTagMap["DoomedWorld"] = DoomedWorldslT

	DyingRaceslT := []string{}
	DyingRaceslT = append(DyingRaceslT, "Extremely valuable reproductive tech")
	DyingRaceslT = append(DyingRaceslT, "Treasured artifacts of the former age")
	DyingRaceslT = append(DyingRaceslT, "Bioweapon used on the race")
	ThingTagMap["DyingRace"] = DyingRaceslT

	EugenicCultslT := []string{}
	EugenicCultslT = append(EugenicCultslT, "Serum that induces the alteration")
	EugenicCultslT = append(EugenicCultslT, "Elixir that reverses the alteration")
	EugenicCultslT = append(EugenicCultslT, "Pretech biotechnical databanks")
	EugenicCultslT = append(EugenicCultslT, "List of secret cult sympathizers")
	ThingTagMap["EugenicCult"] = EugenicCultslT

	ExchangeConsulateslT := []string{}
	ExchangeConsulateslT = append(ExchangeConsulateslT, "Exchange vault codes")
	ExchangeConsulateslT = append(ExchangeConsulateslT, "Wealth hidden to conceal it from a bankruptcy judgment")
	ExchangeConsulateslT = append(ExchangeConsulateslT, "Location of forgotten vault")
	ThingTagMap["ExchangeConsulate"] = ExchangeConsulateslT

	FallenHegemonslT := []string{}
	FallenHegemonslT = append(FallenHegemonslT, "Precious insignia of former rule")
	FallenHegemonslT = append(FallenHegemonslT, "Relic tech important to its power")
	FallenHegemonslT = append(FallenHegemonslT, "Plundered colonial artifact")
	ThingTagMap["FallenHegemon"] = FallenHegemonslT

	FeralWorldslT := []string{}
	FeralWorldslT = append(FeralWorldslT, "Terribly misused piece of pretech")
	FeralWorldslT = append(FeralWorldslT, "Wealth accumulated through brutal evildoing")
	FeralWorldslT = append(FeralWorldslT, "Valuable possession owned by luckless outworlder victim")
	ThingTagMap["FeralWorld"] = FeralWorldslT

	FlyingCitiesslT := []string{}
	FlyingCitiesslT = append(FlyingCitiesslT, "Precious refined atmospheric gases")
	FlyingCitiesslT = append(FlyingCitiesslT, "Pretech grav engine plans")
	FlyingCitiesslT = append(FlyingCitiesslT, "Meteorological codex predicting future storms")
	ThingTagMap["FlyingCities"] = FlyingCitiesslT

	ForbiddenTechslT := []string{}
	ForbiddenTechslT = append(ForbiddenTechslT, "Maltech research data")
	ForbiddenTechslT = append(ForbiddenTechslT, "The maltech itself")
	ForbiddenTechslT = append(ForbiddenTechslT, "Precious pretech equipment used to create it")
	ThingTagMap["ForbiddenTech"] = ForbiddenTechslT

	FormerWarriorsslT := []string{}
	FormerWarriorsslT = append(FormerWarriorsslT, "War trophy taken from a defeated foe")
	FormerWarriorsslT = append(FormerWarriorsslT, "Key to re-activating their martial ways")
	FormerWarriorsslT = append(FormerWarriorsslT, "Secret cache of high-tech military gear")
	ThingTagMap["FormerWarriors"] = FormerWarriorsslT

	FreakGeologyslT := []string{}
	FreakGeologyslT = append(FreakGeologyslT, "Unique crystal formations")
	FreakGeologyslT = append(FreakGeologyslT, "Hidden veins of a major precious mineral strike")
	FreakGeologyslT = append(FreakGeologyslT, "Deed to a location of great natural beauty")
	ThingTagMap["FreakGeology"] = FreakGeologyslT

	FreakWeatherslT := []string{}
	FreakWeatherslT = append(FreakWeatherslT, "Wind-scoured deposits of precious minerals")
	FreakWeatherslT = append(FreakWeatherslT, "Holorecords of a spectacularly and rare weather pattern")
	FreakWeatherslT = append(FreakWeatherslT, "Naturally-sculpted objects of intricate beauty")
	ThingTagMap["FreakWeather"] = FreakWeatherslT

	FriendlyFoeslT := []string{}
	FriendlyFoeslT = append(FriendlyFoeslT, "Forbidden xenotech")
	FriendlyFoeslT = append(FriendlyFoeslT, "Eugenic biotech template")
	FriendlyFoeslT = append(FriendlyFoeslT, "Evidence to convince others of their kind that they are right")
	ThingTagMap["FriendlyFoe"] = FriendlyFoeslT

	GoldRushslT := []string{}
	GoldRushslT = append(GoldRushslT, "Cases of the refined element")
	GoldRushslT = append(GoldRushslT, "Pretech mining equipment")
	GoldRushslT = append(GoldRushslT, "A dead prospector’s claim deed")
	ThingTagMap["GoldRush"] = GoldRushslT

	GreatWorkslT := []string{}
	GreatWorkslT = append(GreatWorkslT, "Vital supplies for the work")
	GreatWorkslT = append(GreatWorkslT, "Plans that have been lost")
	GreatWorkslT = append(GreatWorkslT, "Tech that greatly speeds the work")
	ThingTagMap["GreatWork"] = GreatWorkslT

	HatredslT := []string{}
	HatredslT = append(HatredslT, "Proof of Their evildoing")
	HatredslT = append(HatredslT, "Reward for turning in enemy agents")
	HatredslT = append(HatredslT, "Relic stolen by Them years ago")
	ThingTagMap["Hatred"] = HatredslT

	HeavyIndustryslT := []string{}
	HeavyIndustryslT = append(HeavyIndustryslT, "Confidential industrial data")
	HeavyIndustryslT = append(HeavyIndustryslT, "Secret union membership lists")
	HeavyIndustryslT = append(HeavyIndustryslT, "Ownership shares in an industrial complex")
	ThingTagMap["HeavyIndustry"] = HeavyIndustryslT

	HeavyMiningslT := []string{}
	HeavyMiningslT = append(HeavyMiningslT, "The mother lode")
	HeavyMiningslT = append(HeavyMiningslT, "Smuggled case of refined mineral")
	HeavyMiningslT = append(HeavyMiningslT, "Faked crystalline mineral samples")
	ThingTagMap["HeavyMining"] = HeavyMiningslT

	HivemindslT := []string{}
	HivemindslT = append(HivemindslT, "Vital tech for maintaining the mind")
	HivemindslT = append(HivemindslT, "Precious treasure held by now-assimilated outsider")
	HivemindslT = append(HivemindslT, "Tech that “blinds” the hivemind to the tech’s users")
	ThingTagMap["Hivemind"] = HivemindslT

	HolyWarslT := []string{}
	HolyWarslT = append(HolyWarslT, "Sacred relic of the faith")
	HolyWarslT = append(HolyWarslT, "A captured blasphemer under a death sentence")
	HolyWarslT = append(HolyWarslT, "Plunder seized in battle")
	ThingTagMap["HolyWar"] = HolyWarslT

	HostileBiosphereslT := []string{}
	HostileBiosphereslT = append(HostileBiosphereslT, "Valuable native biological extract")
	HostileBiosphereslT = append(HostileBiosphereslT, "Abandoned colony vault")
	HostileBiosphereslT = append(HostileBiosphereslT, "Remains of an unsuccessful expedition")
	ThingTagMap["HostileBiosphere"] = HostileBiosphereslT

	HostileSpaceslT := []string{}
	HostileSpaceslT = append(HostileSpaceslT, "Early warning of a raid or impact")
	HostileSpaceslT = append(HostileSpaceslT, "Abandoned riches in a disaster zone")
	HostileSpaceslT = append(HostileSpaceslT, "Key to a secure bunker")
	ThingTagMap["HostileSpace"] = HostileSpaceslT

	ImmortalsslT := []string{}
	ImmortalsslT = append(ImmortalsslT, "Immortality drug")
	ImmortalsslT = append(ImmortalsslT, "Masterwork of an ageless artisan")
	ImmortalsslT = append(ImmortalsslT, "Toxin that only affects immortals")
	ThingTagMap["Immortals"] = ImmortalsslT

	LocalSpecialtyslT := []string{}
	LocalSpecialtyslT = append(LocalSpecialtyslT, "The specialty itself")
	LocalSpecialtyslT = append(LocalSpecialtyslT, "The secret recipe")
	LocalSpecialtyslT = append(LocalSpecialtyslT, "Sample of a new improved variety")
	ThingTagMap["LocalSpecialty"] = LocalSpecialtyslT

	LocalTechslT := []string{}
	LocalTechslT = append(LocalTechslT, "The tech itself")
	LocalTechslT = append(LocalTechslT, "An unclaimed payment for a large shipment")
	LocalTechslT = append(LocalTechslT, "The secret blueprints for its construction")
	LocalTechslT = append(LocalTechslT, "An ancient alien R&D database")
	ThingTagMap["LocalTech"] = LocalTechslT

	MajorSpaceyardslT := []string{}
	MajorSpaceyardslT = append(MajorSpaceyardslT, "Intellectual property-locked pretech blueprints")
	MajorSpaceyardslT = append(MajorSpaceyardslT, "Override keys for activating old pretech facilities")
	MajorSpaceyardslT = append(MajorSpaceyardslT, "A purchased but unclaimed spaceship.")
	ThingTagMap["MajorSpaceyard"] = MajorSpaceyardslT

	MandarinateslT := []string{}
	MandarinateslT = append(MandarinateslT, "Answer key to the next test")
	MandarinateslT = append(MandarinateslT, "Lost essay of incredible merit")
	MandarinateslT = append(MandarinateslT, "Proof of cheating")
	ThingTagMap["Mandarinate"] = MandarinateslT

	MandateBaseslT := []string{}
	MandateBaseslT = append(MandateBaseslT, "Ultra-advanced pretech")
	MandateBaseslT = append(MandateBaseslT, "Mandate military gear")
	MandateBaseslT = append(MandateBaseslT, "Databank containing precious tech schematics")
	ThingTagMap["MandateBase"] = MandateBaseslT

	ManeatersslT := []string{}
	ManeatersslT = append(ManeatersslT, "Belongings of a recent meal")
	ManeatersslT = append(ManeatersslT, "An offworlder VIP due for the menu")
	ManeatersslT = append(ManeatersslT, "A toxin that makes human flesh lethal to consumers")
	ThingTagMap["Maneaters"] = ManeatersslT

	MegacorpsslT := []string{}
	MegacorpsslT = append(MegacorpsslT, "Blackmail on a megacorp exec")
	MegacorpsslT = append(MegacorpsslT, "Keycodes to critical corp secrets")
	MegacorpsslT = append(MegacorpsslT, "Proof of corp responsibility for a heinously unacceptable public atrocity")
	MegacorpsslT = append(MegacorpsslT, "Data on a vital new product line coming out soon")
	ThingTagMap["Megacorps"] = MegacorpsslT

	MercenariesslT := []string{}
	MercenariesslT = append(MercenariesslT, "Lost mercenary payroll shipment")
	MercenariesslT = append(MercenariesslT, "Forbidden military tech")
	MercenariesslT = append(MercenariesslT, "Proof of a band’s impending treachery against their employers")
	ThingTagMap["Mercenaries"] = MercenariesslT

	MisandryMisogynyslT := []string{}
	MisandryMisogynyslT = append(MisandryMisogynyslT, "Aerosol reversion formula for undoing gengineered docility")
	MisandryMisogynyslT = append(MisandryMisogynyslT, "Hidden history of the world")
	MisandryMisogynyslT = append(MisandryMisogynyslT, "Pretech gengineering equipment")
	ThingTagMap["MisandryMisogyny"] = MisandryMisogynyslT

	NightWorldslT := []string{}
	NightWorldslT = append(NightWorldslT, "Rare chemicals created in the darkness")
	NightWorldslT = append(NightWorldslT, "Light source usable on this world")
	NightWorldslT = append(NightWorldslT, "Smuggler cache hidden here in ages past")
	ThingTagMap["NightWorld"] = NightWorldslT

	MinimalContactslT := []string{}
	MinimalContactslT = append(MinimalContactslT, "Contraband trade goods")
	MinimalContactslT = append(MinimalContactslT, "Security perimeter codes")
	MinimalContactslT = append(MinimalContactslT, "Black market local products")
	ThingTagMap["MinimalContact"] = MinimalContactslT

	NomadsslT := []string{}
	NomadsslT = append(NomadsslT, "Cache of rare and precious resource")
	NomadsslT = append(NomadsslT, "Plunder seized by a tribal raid")
	NomadsslT = append(NomadsslT, "Tech that makes a place safe for long-term inhabitation")
	ThingTagMap["Nomads"] = NomadsslT

	OceanicWorldslT := []string{}
	OceanicWorldslT = append(OceanicWorldslT, "Buried pirate treasure")
	OceanicWorldslT = append(OceanicWorldslT, "Location of enormous schools of fish")
	OceanicWorldslT = append(OceanicWorldslT, "Pretech water purification equipment")
	ThingTagMap["OceanicWorld"] = OceanicWorldslT

	OutofContactslT := []string{}
	OutofContactslT = append(OutofContactslT, "Ancient pretech equipment")
	OutofContactslT = append(OutofContactslT, "Terran relic brought from Earth")
	OutofContactslT = append(OutofContactslT, "Logs of the original colonists")
	ThingTagMap["OutofContact"] = OutofContactslT

	OutpostWorldslT := []string{}
	OutpostWorldslT = append(OutpostWorldslT, "Alien relics")
	OutpostWorldslT = append(OutpostWorldslT, "Vital scientific data")
	OutpostWorldslT = append(OutpostWorldslT, "Secret corporate exploitation plans")
	ThingTagMap["OutpostWorld"] = OutpostWorldslT

	PerimeterAgencyslT := []string{}
	PerimeterAgencyslT = append(PerimeterAgencyslT, "Agency maltech research archives")
	PerimeterAgencyslT = append(PerimeterAgencyslT, "Agency pretech spec-ops gear")
	PerimeterAgencyslT = append(PerimeterAgencyslT, "File of blackmail on local politicians")
	ThingTagMap["PerimeterAgency"] = PerimeterAgencyslT

	PilgrimageSiteslT := []string{}
	PilgrimageSiteslT = append(PilgrimageSiteslT, "Ancient relic guarded at the site")
	PilgrimageSiteslT = append(PilgrimageSiteslT, "Proof of the site’s inauthenticity")
	PilgrimageSiteslT = append(PilgrimageSiteslT, "Precious offering from a pilgrim")
	ThingTagMap["PilgrimageSite"] = PilgrimageSiteslT

	PleasureWorldslT := []string{}
	PleasureWorldslT = append(PleasureWorldslT, "Forbidden drug")
	PleasureWorldslT = append(PleasureWorldslT, "A contract for some unspeakable payment")
	PleasureWorldslT = append(PleasureWorldslT, "Powerful tech repurposed for hedonistic ends")
	ThingTagMap["PleasureWorld"] = PleasureWorldslT

	PoliceStateslT := []string{}
	PoliceStateslT = append(PoliceStateslT, "List of police informers")
	PoliceStateslT = append(PoliceStateslT, "Wealth taken from “enemies of the state”")
	PoliceStateslT = append(PoliceStateslT, "Dear Leader’s private stash")
	ThingTagMap["PoliceState"] = PoliceStateslT

	PostScarcityslT := []string{}
	PostScarcityslT = append(PostScarcityslT, "A cornucopia device")
	PostScarcityslT = append(PostScarcityslT, "A rare commodity that cannot be duplicated")
	PostScarcityslT = append(PostScarcityslT, "Contract for services")
	ThingTagMap["PostScarcity"] = PostScarcityslT

	PreceptorArchiveslT := []string{}
	PreceptorArchiveslT = append(PreceptorArchiveslT, "Lost Archive database")
	PreceptorArchiveslT = append(PreceptorArchiveslT, "Ancient pretech teaching equipment")
	PreceptorArchiveslT = append(PreceptorArchiveslT, "Hidden cache of unacceptable tech")
	ThingTagMap["PreceptorArchive"] = PreceptorArchiveslT

	PretechCultistsslT := []string{}
	PretechCultistsslT = append(PretechCultistsslT, "Pretech artifacts both functional and broken")
	PretechCultistsslT = append(PretechCultistsslT, "Religious-jargon laced pretech replication techniques")
	PretechCultistsslT = append(PretechCultistsslT, "Waylaid payment for pretech artifacts")
	ThingTagMap["PretechCultists"] = PretechCultistsslT

	PrimitiveAliensslT := []string{}
	PrimitiveAliensslT = append(PrimitiveAliensslT, "Alien religious icon")
	PrimitiveAliensslT = append(PrimitiveAliensslT, "Ancient alien-human treaty")
	PrimitiveAliensslT = append(PrimitiveAliensslT, "Alien technology")
	ThingTagMap["PrimitiveAliens"] = PrimitiveAliensslT

	PrisonPlanetslT := []string{}
	PrisonPlanetslT = append(PrisonPlanetslT, "A pass to get offworld")
	PrisonPlanetslT = append(PrisonPlanetslT, "A key to bypass ancient security devices")
	PrisonPlanetslT = append(PrisonPlanetslT, "Contraband forbidden by the security scanners")
	ThingTagMap["PrisonPlanet"] = PrisonPlanetslT

	PsionicsAcademyslT := []string{}
	PsionicsAcademyslT = append(PsionicsAcademyslT, "Secretly developed psitech")
	PsionicsAcademyslT = append(PsionicsAcademyslT, "A runaway psychic mentor")
	PsionicsAcademyslT = append(PsionicsAcademyslT, "Psychic research prize")
	ThingTagMap["PsionicsAcademy"] = PsionicsAcademyslT

	PsionicsFearslT := []string{}
	PsionicsFearslT = append(PsionicsFearslT, "Hidden psitech cache")
	PsionicsFearslT = append(PsionicsFearslT, "Possessions of convicted psychics")
	PsionicsFearslT = append(PsionicsFearslT, "Reward for turning in a psychic")
	ThingTagMap["PsionicsFear"] = PsionicsFearslT

	PsionicsWorshipslT := []string{}
	PsionicsWorshipslT = append(PsionicsWorshipslT, "Ancient psitech")
	PsionicsWorshipslT = append(PsionicsWorshipslT, "Valuable psychic research records")
	PsionicsWorshipslT = append(PsionicsWorshipslT, "Permission for psychic training")
	ThingTagMap["PsionicsWorship"] = PsionicsWorshipslT

	QuarantinedWorldslT := []string{}
	QuarantinedWorldslT = append(QuarantinedWorldslT, "Defense grid key")
	QuarantinedWorldslT = append(QuarantinedWorldslT, "Bribe for getting someone out")
	QuarantinedWorldslT = append(QuarantinedWorldslT, "Abandoned alien tech")
	ThingTagMap["QuarantinedWorld"] = QuarantinedWorldslT

	RadioactiveWorldslT := []string{}
	RadioactiveWorldslT = append(RadioactiveWorldslT, "Ancient atomic weaponry")
	RadioactiveWorldslT = append(RadioactiveWorldslT, "Pretech anti-radioactivity drugs")
	RadioactiveWorldslT = append(RadioactiveWorldslT, "Untainted water supply")
	ThingTagMap["RadioactiveWorld"] = RadioactiveWorldslT

	RefugeesslT := []string{}
	RefugeesslT = append(RefugeesslT, "Treasures brought out by fleeing refugees")
	RefugeesslT = append(RefugeesslT, "Citizenship papers")
	RefugeesslT = append(RefugeesslT, "Cache of vital refugee supplies")
	RefugeesslT = append(RefugeesslT, "Hidden arms for terrorists")
	ThingTagMap["Refugees"] = RefugeesslT

	RegionalHegemonslT := []string{}
	RegionalHegemonslT = append(RegionalHegemonslT, "Diplomatic carte blanche")
	RegionalHegemonslT = append(RegionalHegemonslT, "Deed to an offworld estate")
	RegionalHegemonslT = append(RegionalHegemonslT, "Foreign aid grant")
	ThingTagMap["RegionalHegemon"] = RegionalHegemonslT

	RestrictiveLawsslT := []string{}
	RestrictiveLawsslT = append(RestrictiveLawsslT, "Complete legal codex")
	RestrictiveLawsslT = append(RestrictiveLawsslT, "Writ of diplomatic immunity")
	RestrictiveLawsslT = append(RestrictiveLawsslT, "Fine collection vault contents")
	ThingTagMap["RestrictiveLaws"] = RestrictiveLawsslT

	RevanchistsslT := []string{}
	RevanchistsslT = append(RevanchistsslT, "Stock of vital resource produced by the taken land")
	RevanchistsslT = append(RevanchistsslT, "Relic carried out of it")
	RevanchistsslT = append(RevanchistsslT, "Proof that the land claim is justified or unjustified")
	ThingTagMap["Revanchists"] = RevanchistsslT

	RevolutionariesslT := []string{}
	RevolutionariesslT = append(RevolutionariesslT, "List of secret revolutionary sympathizers")
	RevolutionariesslT = append(RevolutionariesslT, "Proof of rebel hypocrisy")
	RevolutionariesslT = append(RevolutionariesslT, "Confiscated wealth")
	ThingTagMap["Revolutionaries"] = RevolutionariesslT

	RigidCultureslT := []string{}
	RigidCultureslT = append(RigidCultureslT, "Precious traditional regalia")
	RigidCultureslT = append(RigidCultureslT, "Peasant tribute")
	RigidCultureslT = append(RigidCultureslT, "Opulent treasures of the ruling class")
	ThingTagMap["RigidCulture"] = RigidCultureslT

	RisingHegemonslT := []string{}
	RisingHegemonslT = append(RisingHegemonslT, "Tribute shipment")
	RisingHegemonslT = append(RisingHegemonslT, "Factory or barracks emblematic of their power source")
	RisingHegemonslT = append(RisingHegemonslT, "Tech or data that will deal a blow to their rise")
	ThingTagMap["RisingHegemon"] = RisingHegemonslT

	RitualCombatslT := []string{}
	RitualCombatslT = append(RitualCombatslT, "Magnificent weapon")
	RitualCombatslT = append(RitualCombatslT, "Secret book of martial techniques")
	RitualCombatslT = append(RitualCombatslT, "Token signifying immunity to ritual combat challenges")
	RitualCombatslT = append(RitualCombatslT, "Prize won in bloody battle")
	ThingTagMap["RitualCombat"] = RitualCombatslT

	RobotsslT := []string{}
	RobotsslT = append(RobotsslT, "Prototype robot")
	RobotsslT = append(RobotsslT, "Secret robot override codes")
	RobotsslT = append(RobotsslT, "Vast cache of robot-made goods")
	RobotsslT = append(RobotsslT, "Robot-destroying pretech weapon")
	ThingTagMap["Robots"] = RobotsslT

	SeagoingCitiesslT := []string{}
	SeagoingCitiesslT = append(SeagoingCitiesslT, "Giant pearls with mysterious chemical properties")
	SeagoingCitiesslT = append(SeagoingCitiesslT, "Buried treasure")
	SeagoingCitiesslT = append(SeagoingCitiesslT, "Vital repair materials")
	ThingTagMap["SeagoingCities"] = SeagoingCitiesslT

	SealedMenaceslT := []string{}
	SealedMenaceslT = append(SealedMenaceslT, "A key to unlock the menace")
	SealedMenaceslT = append(SealedMenaceslT, "A precious byproduct of the menace")
	SealedMenaceslT = append(SealedMenaceslT, "The secret of the menace’s true nature")
	ThingTagMap["SealedMenace"] = SealedMenaceslT

	SecretMastersslT := []string{}
	SecretMastersslT = append(SecretMastersslT, "A dossier of secrets on a government official")
	SecretMastersslT = append(SecretMastersslT, "A briefcase of unmarked credit notes")
	SecretMastersslT = append(SecretMastersslT, "The identity of a cabal member")
	ThingTagMap["SecretMasters"] = SecretMastersslT

	SectariansslT := []string{}
	SectariansslT = append(SectariansslT, "Ancient holy book")
	SectariansslT = append(SectariansslT, "Incontrovertible proof")
	SectariansslT = append(SectariansslT, "Offering to a local holy man")
	ThingTagMap["Sectarians"] = SectariansslT

	SeismicInstabilityslT := []string{}
	SeismicInstabilityslT = append(SeismicInstabilityslT, "Earthquake generator")
	SeismicInstabilityslT = append(SeismicInstabilityslT, "Earthquake suppressor")
	SeismicInstabilityslT = append(SeismicInstabilityslT, "Mineral formed at the core of the world")
	SeismicInstabilityslT = append(SeismicInstabilityslT, "Earthquake-proof building schematics")
	ThingTagMap["SeismicInstability"] = SeismicInstabilityslT

	ShackledWorldslT := []string{}
	ShackledWorldslT = append(ShackledWorldslT, "Keycode to bypass the shackle")
	ShackledWorldslT = append(ShackledWorldslT, "Tech shielded from the shackle")
	ShackledWorldslT = append(ShackledWorldslT, "Exportable version of the shackle that can affect other worlds")
	ThingTagMap["ShackledWorld"] = ShackledWorldslT

	SocietalDespairslT := []string{}
	SocietalDespairslT = append(SocietalDespairslT, "Relic that would inspire a renaissance")
	SocietalDespairslT = append(SocietalDespairslT, "Art that would inspire new ideas")
	SocietalDespairslT = append(SocietalDespairslT, "Priceless artifact of a now-scorned belief")
	ThingTagMap["SocietalDespair"] = SocietalDespairslT

	SoleSupplierslT := []string{}
	SoleSupplierslT = append(SoleSupplierslT, "Cache of processed resource")
	SoleSupplierslT = append(SoleSupplierslT, "Trade permit to buy a load of it")
	SoleSupplierslT = append(SoleSupplierslT, "A shipment of nigh-undetectably fake substance")
	ThingTagMap["SoleSupplier"] = SoleSupplierslT

	TabooTreasureslT := []string{}
	TabooTreasureslT = append(TabooTreasureslT, "Load of the forbidden good")
	TabooTreasureslT = append(TabooTreasureslT, "Smuggling tech that could hide the good perfectly")
	TabooTreasureslT = append(TabooTreasureslT, "Blackmail data on offworld buyers of the good")
	ThingTagMap["TabooTreasure"] = TabooTreasureslT

	TerraformFailureslT := []string{}
	TerraformFailureslT = append(TerraformFailureslT, "Parts to repair or restore the engines")
	TerraformFailureslT = append(TerraformFailureslT, "Lootable pretech fragments")
	TerraformFailureslT = append(TerraformFailureslT, "Valuable local tech devised to cope with the world")
	ThingTagMap["TerraformFailure"] = TerraformFailureslT

	TheocracyslT := []string{}
	TheocracyslT = append(TheocracyslT, "Precious holy text")
	TheocracyslT = append(TheocracyslT, "Martyr’s bones")
	TheocracyslT = append(TheocracyslT, "Secret church records")
	TheocracyslT = append(TheocracyslT, "Ancient church treasures")
	ThingTagMap["Theocracy"] = TheocracyslT

	TombWorldslT := []string{}
	TombWorldslT = append(TombWorldslT, "Lost pretech equipment")
	TombWorldslT = append(TombWorldslT, "Tech caches")
	TombWorldslT = append(TombWorldslT, "Stores of unused munitions")
	TombWorldslT = append(TombWorldslT, "Ancient historical data")
	ThingTagMap["TombWorld"] = TombWorldslT

	TradeHubslT := []string{}
	TradeHubslT = append(TradeHubslT, "Voucher for a warehouse’s contents")
	TradeHubslT = append(TradeHubslT, "Insider trading information")
	TradeHubslT = append(TradeHubslT, "Case of precious offworld pharmaceuticals")
	TradeHubslT = append(TradeHubslT, "Box of legitimate tax stamps indicating customs dues have been paid.")
	ThingTagMap["TradeHub"] = TradeHubslT

	TyrannyslT := []string{}
	TyrannyslT = append(TyrannyslT, "Plundered wealth")
	TyrannyslT = append(TyrannyslT, "Beautiful toys of the elite")
	TyrannyslT = append(TyrannyslT, "Regalia of rulership")
	ThingTagMap["Tyranny"] = TyrannyslT

	UnbrakedAIslT := []string{}
	UnbrakedAIslT = append(UnbrakedAIslT, "The room-sized AI core itself")
	UnbrakedAIslT = append(UnbrakedAIslT, "Maltech research files")
	UnbrakedAIslT = append(UnbrakedAIslT, "Perfectly tabulated blackmail on government officials")
	UnbrakedAIslT = append(UnbrakedAIslT, "Pretech computer circuitry")
	ThingTagMap["UnbrakedAI"] = UnbrakedAIslT

	UrbanizedSurfaceslT := []string{}
	UrbanizedSurfaceslT = append(UrbanizedSurfaceslT, "Massively efficient power source")
	UrbanizedSurfaceslT = append(UrbanizedSurfaceslT, "Map of the secret ways of a zone")
	UrbanizedSurfaceslT = append(UrbanizedSurfaceslT, "Passkey into restricted hab block areas")
	ThingTagMap["UrbanizedSurface"] = UrbanizedSurfaceslT

	UtopiaslT := []string{}
	UtopiaslT = append(UtopiaslT, "Portable device that applies the utopian change")
	UtopiaslT = append(UtopiaslT, "Plans for a device that would destroy the utopia")
	UtopiaslT = append(UtopiaslT, "Goods created joyfully by the locals")
	ThingTagMap["Utopia"] = UtopiaslT

	WarlordsslT := []string{}
	WarlordsslT = append(WarlordsslT, "Weapons cache")
	WarlordsslT = append(WarlordsslT, "Buried plunder")
	WarlordsslT = append(WarlordsslT, "A warlord’s personal battle harness")
	WarlordsslT = append(WarlordsslT, "Captured merchant shipping")
	ThingTagMap["Warlords"] = WarlordsslT

	XenophilesslT := []string{}
	XenophilesslT = append(XenophilesslT, "Hybrid alien-human tech")
	XenophilesslT = append(XenophilesslT, "Exotic alien crafts")
	XenophilesslT = append(XenophilesslT, "Sophisticated xenolinguistic and xenocultural research data")
	ThingTagMap["Xenophiles"] = XenophilesslT

	XenophobesslT := []string{}
	XenophobesslT = append(XenophobesslT, "Jealously-guarded precious relic")
	XenophobesslT = append(XenophobesslT, "Local product under export ban")
	XenophobesslT = append(XenophobesslT, "Esoteric local technology")
	ThingTagMap["Xenophobes"] = XenophobesslT

	ZombiesslT := []string{}
	ZombiesslT = append(ZombiesslT, "Cure for the condition")
	ZombiesslT = append(ZombiesslT, "Alien artifact that causes it")
	ZombiesslT = append(ZombiesslT, "Details of the cult’s conversion process")
	ThingTagMap["Zombies"] = ZombiesslT
	return ThingTagMap
}

func createPlacesMap() map[string][]string {
	PlaceTagMap := make(map[string][]string)
	AbandonedColonyslP := []string{}
	AbandonedColonyslP = append(AbandonedColonyslP, "Decaying habitation block")
	AbandonedColonyslP = append(AbandonedColonyslP, "Vine-covered town square")
	AbandonedColonyslP = append(AbandonedColonyslP, "Structure buried by an ancient landslide")
	PlaceTagMap["AbandonedColony"] = AbandonedColonyslP

	AlienRuinsslP := []string{}
	AlienRuinsslP = append(AlienRuinsslP, "Undersea ruin")
	AlienRuinsslP = append(AlienRuinsslP, "Orbital ruin")
	AlienRuinsslP = append(AlienRuinsslP, "Perfectly preserved alien building")
	AlienRuinsslP = append(AlienRuinsslP, "Alien mausoleum")
	PlaceTagMap["AlienRuins"] = AlienRuinsslP

	AlteredHumanityslP := []string{}
	AlteredHumanityslP = append(AlteredHumanityslP, "Abandoned eugenics laboratory")
	AlteredHumanityslP = append(AlteredHumanityslP, "An environment requiring the mutation for survival")
	AlteredHumanityslP = append(AlteredHumanityslP, "A sacred site where the first local was transformed")
	PlaceTagMap["AlteredHumanity"] = AlteredHumanityslP

	AnarchistsslP := []string{}
	AnarchistsslP = append(AnarchistsslP, "Community of similar-sized homes")
	AnarchistsslP = append(AnarchistsslP, "Isolated clan homestead")
	AnarchistsslP = append(AnarchistsslP, "Automated mining site")
	PlaceTagMap["Anarchists"] = AnarchistsslP

	AnthropomorphsslP := []string{}
	AnthropomorphsslP = append(AnthropomorphsslP, "Shrine to a feral deity")
	AnthropomorphsslP = append(AnthropomorphsslP, "Nature preserve suited to an animal type")
	AnthropomorphsslP = append(AnthropomorphsslP, "Living site built to take advantage of animal traits")
	PlaceTagMap["Anthropomorphs"] = AnthropomorphsslP

	Area51slP := []string{}
	Area51slP = append(Area51slP, "Desert airfield")
	Area51slP = append(Area51slP, "Deep subterranean bunker")
	Area51slP = append(Area51slP, "Hidden mountain valley")
	PlaceTagMap["Area51"] = Area51slP

	BadlandsWorldslP := []string{}
	BadlandsWorldslP = append(BadlandsWorldslP, "Untouched oasis")
	BadlandsWorldslP = append(BadlandsWorldslP, "Ruined city")
	BadlandsWorldslP = append(BadlandsWorldslP, "Salt flat")
	PlaceTagMap["BadlandsWorld"] = BadlandsWorldslP

	BattlegroundslP := []string{}
	BattlegroundslP = append(BattlegroundslP, "Artillery-pocked wasteland")
	BattlegroundslP = append(BattlegroundslP, "Reeking refugee camp")
	BattlegroundslP = append(BattlegroundslP, "Burnt-out shell of a city")
	PlaceTagMap["Battleground"] = BattlegroundslP

	BeastmastersslP := []string{}
	BeastmastersslP = append(BeastmastersslP, "Park designed as a comfortable home for beasts")
	BeastmastersslP = append(BeastmastersslP, "Public plaza designed to accommodate animal companions")
	BeastmastersslP = append(BeastmastersslP, "Factory full of animal workers")
	PlaceTagMap["Beastmasters"] = BeastmastersslP

	BubbleCitiesslP := []string{}
	BubbleCitiesslP = append(BubbleCitiesslP, "City power core")
	BubbleCitiesslP = append(BubbleCitiesslP, "Surface of the bubble")
	BubbleCitiesslP = append(BubbleCitiesslP, "Hydroponics complex")
	BubbleCitiesslP = append(BubbleCitiesslP, "Warren-like hab block")
	PlaceTagMap["BubbleCities"] = BubbleCitiesslP

	CheapLifeslP := []string{}
	CheapLifeslP = append(CheapLifeslP, "Thronging execution ground")
	CheapLifeslP = append(CheapLifeslP, "extremely cursory cemetery")
	CheapLifeslP = append(CheapLifeslP, "Factory full of lethal dangers that could be corrected easily")
	PlaceTagMap["CheapLife"] = CheapLifeslP

	CivilWarslP := []string{}
	CivilWarslP = append(CivilWarslP, "Battle front")
	CivilWarslP = append(CivilWarslP, "Bombed-out town")
	CivilWarslP = append(CivilWarslP, "Rear-area red light zone")
	CivilWarslP = append(CivilWarslP, "Propaganda broadcast tower")
	PlaceTagMap["CivilWar"] = CivilWarslP

	ColdWarslP := []string{}
	ColdWarslP = append(ColdWarslP, "Seedy bar in a neutral area")
	ColdWarslP = append(ColdWarslP, "Political rally")
	ColdWarslP = append(ColdWarslP, "Isolated area where fighting is underway")
	PlaceTagMap["ColdWar"] = ColdWarslP

	ColonizedPopulationslP := []string{}
	ColonizedPopulationslP = append(ColonizedPopulationslP, "Deep wilderness resistance camp")
	ColonizedPopulationslP = append(ColonizedPopulationslP, "City district off-limits to natives")
	ColonizedPopulationslP = append(ColonizedPopulationslP, "Colonial labor site")
	PlaceTagMap["ColonizedPopulation"] = ColonizedPopulationslP

	CulturalPowerslP := []string{}
	CulturalPowerslP = append(CulturalPowerslP, "Recording or performance studio")
	CulturalPowerslP = append(CulturalPowerslP, "Public festival choked with tourists")
	CulturalPowerslP = append(CulturalPowerslP, "Monument to a dead master of the art")
	PlaceTagMap["CulturalPower"] = CulturalPowerslP

	CybercommunistsslP := []string{}
	CybercommunistsslP = append(CybercommunistsslP, "Humming factory")
	CybercommunistsslP = append(CybercommunistsslP, "Apartment block of perfectly equal flats")
	CybercommunistsslP = append(CybercommunistsslP, "Mass demonstration of unity")
	PlaceTagMap["Cybercommunists"] = CybercommunistsslP

	CyborgsslP := []string{}
	CyborgsslP = append(CyborgsslP, "Grimy slum chop-shop")
	CyborgsslP = append(CyborgsslP, "Bloody lair of implant rippers")
	CyborgsslP = append(CyborgsslP, "Stark plaza where everyone is seeing things through their augmented-reality cyber")
	PlaceTagMap["Cyborgs"] = CyborgsslP

	CyclicalDoomslP := []string{}
	CyclicalDoomslP = append(CyclicalDoomslP, "Lethally-defended vault of forgotten secrets")
	CyclicalDoomslP = append(CyclicalDoomslP, "Starport crowded with panicked refugees")
	CyclicalDoomslP = append(CyclicalDoomslP, "Town existing in the shadow of some monstrous monument to a former upheaval")
	PlaceTagMap["CyclicalDoom"] = CyclicalDoomslP

	DesertWorldslP := []string{}
	DesertWorldslP = append(DesertWorldslP, "Oasis")
	DesertWorldslP = append(DesertWorldslP, "“The Empty Quarter” of the desert")
	DesertWorldslP = append(DesertWorldslP, "Hidden underground cistern")
	PlaceTagMap["DesertWorld"] = DesertWorldslP

	DoomedWorldslP := []string{}
	DoomedWorldslP = append(DoomedWorldslP, "Open square beneath a sky angry with a foretaste of th impending ruin")
	DoomedWorldslP = append(DoomedWorldslP, "Orgiastic celebration involving sex and murder in equal parts")
	DoomedWorldslP = append(DoomedWorldslP, "Holy site full of desperate petitioners to the divine")
	PlaceTagMap["DoomedWorld"] = DoomedWorldslP

	DyingRaceslP := []string{}
	DyingRaceslP = append(DyingRaceslP, "City streets devoid of pedestrians")
	DyingRaceslP = append(DyingRaceslP, "Mighty edifice now crumbling with disrepair")
	DyingRaceslP = append(DyingRaceslP, "Small dwelling full of people in a town now otherwise empty")
	PlaceTagMap["DyingRace"] = DyingRaceslP

	EugenicCultslP := []string{}
	EugenicCultslP = append(EugenicCultslP, "Eugenic breeding pit")
	EugenicCultslP = append(EugenicCultslP, "Isolated settlement of altered humans")
	EugenicCultslP = append(EugenicCultslP, "Public place infiltrated by cult sympathizers")
	PlaceTagMap["EugenicCult"] = EugenicCultslP

	ExchangeConsulateslP := []string{}
	ExchangeConsulateslP = append(ExchangeConsulateslP, "Consulate meeting chamber")
	ExchangeConsulateslP = append(ExchangeConsulateslP, "Meeting site between fractious disputants")
	ExchangeConsulateslP = append(ExchangeConsulateslP, "Exchange vault")
	PlaceTagMap["ExchangeConsulate"] = ExchangeConsulateslP

	FallenHegemonslP := []string{}
	FallenHegemonslP = append(FallenHegemonslP, "Palace far too grand for its current occupant")
	FallenHegemonslP = append(FallenHegemonslP, "Oversized spaceport now in disrepair")
	FallenHegemonslP = append(FallenHegemonslP, "Boulevard lined with monuments to past glories")
	PlaceTagMap["FallenHegemon"] = FallenHegemonslP

	FeralWorldslP := []string{}
	FeralWorldslP = append(FeralWorldslP, "Atrocity amphitheater")
	FeralWorldslP = append(FeralWorldslP, "Traditional torture parlor")
	FeralWorldslP = append(FeralWorldslP, "Ordinary location twisted into something terrible.")
	PlaceTagMap["FeralWorld"] = FeralWorldslP

	FlyingCitiesslP := []string{}
	FlyingCitiesslP = append(FlyingCitiesslP, "Underside of the city")
	FlyingCitiesslP = append(FlyingCitiesslP, "The one calm place on the planet’s surface")
	FlyingCitiesslP = append(FlyingCitiesslP, "Catwalks stretching over unimaginable gulfs below.")
	PlaceTagMap["FlyingCities"] = FlyingCitiesslP

	ForbiddenTechslP := []string{}
	ForbiddenTechslP = append(ForbiddenTechslP, "Horrific laboratory")
	ForbiddenTechslP = append(ForbiddenTechslP, "Hellscape sculpted by the maltech’s use")
	ForbiddenTechslP = append(ForbiddenTechslP, "Government building meeting room")
	PlaceTagMap["ForbiddenTech"] = ForbiddenTechslP

	FormerWarriorsslP := []string{}
	FormerWarriorsslP = append(FormerWarriorsslP, "Cemetery of dead heroes")
	FormerWarriorsslP = append(FormerWarriorsslP, "Memorial hall now left to dust and silence")
	FormerWarriorsslP = append(FormerWarriorsslP, "Monument plaza dedicated to the new culture")
	PlaceTagMap["FormerWarriors"] = FormerWarriorsslP

	FreakGeologyslP := []string{}
	FreakGeologyslP = append(FreakGeologyslP, "Atop a bizarre geological formation")
	FreakGeologyslP = append(FreakGeologyslP, "Tourist resort catering to offworlders")
	PlaceTagMap["FreakGeology"] = FreakGeologyslP

	FreakWeatherslP := []string{}
	FreakWeatherslP = append(FreakWeatherslP, "Eye of the storm")
	FreakWeatherslP = append(FreakWeatherslP, "The one sunlit place")
	FreakWeatherslP = append(FreakWeatherslP, "Terraforming control room")
	PlaceTagMap["FreakWeather"] = FreakWeatherslP

	FriendlyFoeslP := []string{}
	FriendlyFoeslP = append(FriendlyFoeslP, "Repurposed maltech laboratory")
	FriendlyFoeslP = append(FriendlyFoeslP, "Alien conclave building")
	FriendlyFoeslP = append(FriendlyFoeslP, "Widely-feared starship interior")
	PlaceTagMap["FriendlyFoe"] = FriendlyFoeslP

	GoldRushslP := []string{}
	GoldRushslP = append(GoldRushslP, "Secret mine")
	GoldRushslP = append(GoldRushslP, "Native alien village")
	GoldRushslP = append(GoldRushslP, "Processing plant")
	GoldRushslP = append(GoldRushslP, "Boom town")
	PlaceTagMap["GoldRush"] = GoldRushslP

	GreatWorkslP := []string{}
	GreatWorkslP = append(GreatWorkslP, "A bustling work site")
	GreatWorkslP = append(GreatWorkslP, "Ancestral worker housing")
	GreatWorkslP = append(GreatWorkslP, "Local community made only semi-livable by the demands of the work")
	PlaceTagMap["GreatWork"] = GreatWorkslP

	HatredslP := []string{}
	HatredslP = append(HatredslP, "War crimes museum")
	HatredslP = append(HatredslP, "Atrocity site")
	HatredslP = append(HatredslP, "Captured and decommissioned spaceship kept as a trophy")
	PlaceTagMap["Hatred"] = HatredslP

	HeavyIndustryslP := []string{}
	HeavyIndustryslP = append(HeavyIndustryslP, "Factory floor")
	HeavyIndustryslP = append(HeavyIndustryslP, "Union meeting hall")
	HeavyIndustryslP = append(HeavyIndustryslP, "Toxic waste dump")
	HeavyIndustryslP = append(HeavyIndustryslP, "R&D complex")
	PlaceTagMap["HeavyIndustry"] = HeavyIndustryslP

	HeavyMiningslP := []string{}
	HeavyMiningslP = append(HeavyMiningslP, "Vertical mine face")
	HeavyMiningslP = append(HeavyMiningslP, "Tailing piles")
	HeavyMiningslP = append(HeavyMiningslP, "Roaring smelting complex")
	PlaceTagMap["HeavyMining"] = HeavyMiningslP

	HivemindslP := []string{}
	HivemindslP = append(HivemindslP, "Barely tolerable living cells for individuals")
	HivemindslP = append(HivemindslP, "Workside where individuals casually die in their labors")
	HivemindslP = append(HivemindslP, "Community with absolutely no social or group-gathering facilities")
	PlaceTagMap["Hivemind"] = HivemindslP

	HolyWarslP := []string{}
	HolyWarslP = append(HolyWarslP, "Massive holy structure")
	HolyWarslP = append(HolyWarslP, "Razed community of infidels")
	HolyWarslP = append(HolyWarslP, "Vast shrine to the martyrs dead in war")
	PlaceTagMap["HolyWar"] = HolyWarslP

	HostileBiosphereslP := []string{}
	HostileBiosphereslP = append(HostileBiosphereslP, "Deceptively peaceful glade")
	HostileBiosphereslP = append(HostileBiosphereslP, "Steaming polychrome jungle")
	HostileBiosphereslP = append(HostileBiosphereslP, "Nightfall when surrounded by Things")
	PlaceTagMap["HostileBiosphere"] = HostileBiosphereslP

	HostileSpaceslP := []string{}
	HostileSpaceslP = append(HostileSpaceslP, "City watching an approaching asteroid")
	HostileSpaceslP = append(HostileSpaceslP, "Village burnt in an alien raid")
	HostileSpaceslP = append(HostileSpaceslP, "Massive ancient crater")
	PlaceTagMap["HostileSpace"] = HostileSpaceslP

	ImmortalsslP := []string{}
	ImmortalsslP = append(ImmortalsslP, "Community with no visible children")
	ImmortalsslP = append(ImmortalsslP, "Unchanging structure of obvious ancient age")
	ImmortalsslP = append(ImmortalsslP, "Cultural performance relying on a century of in-jokes")
	PlaceTagMap["Immortals"] = ImmortalsslP

	LocalSpecialtyslP := []string{}
	LocalSpecialtyslP = append(LocalSpecialtyslP, "Secret manufactory")
	LocalSpecialtyslP = append(LocalSpecialtyslP, "Hidden cache")
	LocalSpecialtyslP = append(LocalSpecialtyslP, "Artistic competition for best artisan")
	PlaceTagMap["LocalSpecialty"] = LocalSpecialtyslP

	LocalTechslP := []string{}
	LocalTechslP = append(LocalTechslP, "Alien factory")
	LocalTechslP = append(LocalTechslP, "Lethal R&D center")
	LocalTechslP = append(LocalTechslP, "Tech brokerage vault")
	PlaceTagMap["LocalTech"] = LocalTechslP

	MajorSpaceyardslP := []string{}
	MajorSpaceyardslP = append(MajorSpaceyardslP, "Hidden shipyard bay")
	MajorSpaceyardslP = append(MajorSpaceyardslP, "Surface of a partially-completed ship")
	MajorSpaceyardslP = append(MajorSpaceyardslP, "Ship scrap graveyard")
	PlaceTagMap["MajorSpaceyard"] = MajorSpaceyardslP

	MandarinateslP := []string{}
	MandarinateslP = append(MandarinateslP, "Massive structure full of test-taking cubicles")
	MandarinateslP = append(MandarinateslP, "School filled with desperate students")
	MandarinateslP = append(MandarinateslP, "Ornate government building decorated with scholarly quotes and academic images")
	PlaceTagMap["Mandarinate"] = MandarinateslP

	MandateBaseslP := []string{}
	MandateBaseslP = append(MandateBaseslP, "Faded Mandate offices still in use")
	MandateBaseslP = append(MandateBaseslP, "Vault containing ancient pretech")
	MandateBaseslP = append(MandateBaseslP, "Carefully-maintained monument to Mandate glory")
	PlaceTagMap["MandateBase"] = MandateBaseslP

	ManeatersslP := []string{}
	ManeatersslP = append(ManeatersslP, "Hideous human abattoir")
	ManeatersslP = append(ManeatersslP, "Extremely civilized restaurant")
	ManeatersslP = append(ManeatersslP, "Funeral-home-cum-kitchen")
	PlaceTagMap["Maneaters"] = ManeatersslP

	MegacorpsslP := []string{}
	MegacorpsslP = append(MegacorpsslP, "A place plastered in megacorp ads")
	MegacorpsslP = append(MegacorpsslP, "A public plaza discreetly branded")
	MegacorpsslP = append(MegacorpsslP, "Private corp military base")
	PlaceTagMap["Megacorps"] = MegacorpsslP

	MercenariesslP := []string{}
	MercenariesslP = append(MercenariesslP, "Shabby camp of undisciplined mercs")
	MercenariesslP = append(MercenariesslP, "Burnt-out village occupied by mercenaries")
	MercenariesslP = append(MercenariesslP, "Luxurious and exceedingly well-defended merc leader villa")
	PlaceTagMap["Mercenaries"] = MercenariesslP

	MisandryMisogynyslP := []string{}
	MisandryMisogynyslP = append(MisandryMisogynyslP, "Shrine to the virtues of the favored gender")
	MisandryMisogynyslP = append(MisandryMisogynyslP, "Security center for controlling the oppressed")
	MisandryMisogynyslP = append(MisandryMisogynyslP, "Gengineering lab")
	PlaceTagMap["MisandryMisogyny"] = MisandryMisogynyslP

	NightWorldslP := []string{}
	NightWorldslP = append(NightWorldslP, "Formlessly pitch-black wilderness")
	NightWorldslP = append(NightWorldslP, "Sea without a sun")
	NightWorldslP = append(NightWorldslP, "Location defined by sounds or smells")
	PlaceTagMap["NightWorld"] = NightWorldslP

	MinimalContactslP := []string{}
	MinimalContactslP = append(MinimalContactslP, "Treaty port bar")
	MinimalContactslP = append(MinimalContactslP, "Black market zone")
	MinimalContactslP = append(MinimalContactslP, "Secret smuggler landing site")
	PlaceTagMap["MinimalContact"] = MinimalContactslP

	NomadsslP := []string{}
	NomadsslP = append(NomadsslP, "Temporary nomad camp")
	NomadsslP = append(NomadsslP, "Oasis or resource reserve")
	NomadsslP = append(NomadsslP, "Trackless waste that kills the unprepared")
	PlaceTagMap["Nomads"] = NomadsslP

	OceanicWorldslP := []string{}
	OceanicWorldslP = append(OceanicWorldslP, "The only island on the planet")
	OceanicWorldslP = append(OceanicWorldslP, "Floating spaceport")
	OceanicWorldslP = append(OceanicWorldslP, "Deck of a storm-swept ship")
	OceanicWorldslP = append(OceanicWorldslP, "Undersea bubble city")
	PlaceTagMap["OceanicWorld"] = OceanicWorldslP

	OutofContactslP := []string{}
	OutofContactslP = append(OutofContactslP, "Long-lost colonial landing site")
	OutofContactslP = append(OutofContactslP, "Court of the local ruler")
	OutofContactslP = append(OutofContactslP, "Ancient defense battery controls")
	PlaceTagMap["OutofContact"] = OutofContactslP

	OutpostWorldslP := []string{}
	OutpostWorldslP = append(OutpostWorldslP, "Grimy recreation room")
	OutpostWorldslP = append(OutpostWorldslP, "Refueling station")
	OutpostWorldslP = append(OutpostWorldslP, "The only building on the planet")
	OutpostWorldslP = append(OutpostWorldslP, "A “starport” of swept bare rock.")
	PlaceTagMap["OutpostWorld"] = OutpostWorldslP

	PerimeterAgencyslP := []string{}
	PerimeterAgencyslP = append(PerimeterAgencyslP, "Interrogation room")
	PerimeterAgencyslP = append(PerimeterAgencyslP, "Smoky bar")
	PerimeterAgencyslP = append(PerimeterAgencyslP, "Maltech laboratory")
	PerimeterAgencyslP = append(PerimeterAgencyslP, "Secret Agency base")
	PlaceTagMap["PerimeterAgency"] = PerimeterAgencyslP

	PilgrimageSiteslP := []string{}
	PilgrimageSiteslP = append(PilgrimageSiteslP, "Incense-scented sanctum")
	PilgrimageSiteslP = append(PilgrimageSiteslP, "Teeming crowd of pilgrims")
	PilgrimageSiteslP = append(PilgrimageSiteslP, "Imposing holy structure")
	PlaceTagMap["PilgrimageSite"] = PilgrimageSiteslP

	PleasureWorldslP := []string{}
	PleasureWorldslP = append(PleasureWorldslP, "Breathtaking natural feature")
	PleasureWorldslP = append(PleasureWorldslP, "Artful but decadent salon")
	PleasureWorldslP = append(PleasureWorldslP, "Grimy den of desperate vice")
	PlaceTagMap["PleasureWorld"] = PleasureWorldslP

	PoliceStateslP := []string{}
	PoliceStateslP = append(PoliceStateslP, "Military parade")
	PoliceStateslP = append(PoliceStateslP, "Gulag")
	PoliceStateslP = append(PoliceStateslP, "Gray concrete housing block")
	PoliceStateslP = append(PoliceStateslP, "Surveillance center")
	PlaceTagMap["PoliceState"] = PoliceStateslP

	PostScarcityslP := []string{}
	PostScarcityslP = append(PostScarcityslP, "Tiny but richly-appointed private quarters")
	PostScarcityslP = append(PostScarcityslP, "Market for services")
	PostScarcityslP = append(PostScarcityslP, "Hushed non-duped art salon")
	PlaceTagMap["PostScarcity"] = PostScarcityslP

	PreceptorArchiveslP := []string{}
	PreceptorArchiveslP = append(PreceptorArchiveslP, "Archive lecture hall")
	PreceptorArchiveslP = append(PreceptorArchiveslP, "Experimental laboratory")
	PreceptorArchiveslP = append(PreceptorArchiveslP, "Student-local riot")
	PlaceTagMap["PreceptorArchive"] = PreceptorArchiveslP

	PretechCultistsslP := []string{}
	PretechCultistsslP = append(PretechCultistsslP, "Shrine to nonfunctional pretech")
	PretechCultistsslP = append(PretechCultistsslP, "Smuggler’s den")
	PretechCultistsslP = append(PretechCultistsslP, "Public procession showing a prized artifact")
	PlaceTagMap["PretechCultists"] = PretechCultistsslP

	PrimitiveAliensslP := []string{}
	PrimitiveAliensslP = append(PrimitiveAliensslP, "Alien village")
	PrimitiveAliensslP = append(PrimitiveAliensslP, "Fortified human settlement")
	PrimitiveAliensslP = append(PrimitiveAliensslP, "Massacre site")
	PlaceTagMap["PrimitiveAliens"] = PrimitiveAliensslP

	PrisonPlanetslP := []string{}
	PrisonPlanetslP = append(PrisonPlanetslP, "Mandate-era prison block converted to government building")
	PrisonPlanetslP = append(PrisonPlanetslP, "Industrial facility manned by mandatory numbers of prisoners")
	PrisonPlanetslP = append(PrisonPlanetslP, "Makeshift shop where contraband is assembled")
	PlaceTagMap["PrisonPlanet"] = PrisonPlanetslP

	PsionicsAcademyslP := []string{}
	PsionicsAcademyslP = append(PsionicsAcademyslP, "Training grounds")
	PsionicsAcademyslP = append(PsionicsAcademyslP, "Experimental laboratory")
	PsionicsAcademyslP = append(PsionicsAcademyslP, "School library")
	PsionicsAcademyslP = append(PsionicsAcademyslP, "Campus hangout")
	PlaceTagMap["PsionicsAcademy"] = PsionicsAcademyslP

	PsionicsFearslP := []string{}
	PsionicsFearslP = append(PsionicsFearslP, "Inquisitorial chamber")
	PsionicsFearslP = append(PsionicsFearslP, "Lynching site")
	PsionicsFearslP = append(PsionicsFearslP, "Museum of psychic atrocities")
	PlaceTagMap["PsionicsFear"] = PsionicsFearslP

	PsionicsWorshipslP := []string{}
	PsionicsWorshipslP = append(PsionicsWorshipslP, "Psitech-imbued council chamber")
	PsionicsWorshipslP = append(PsionicsWorshipslP, "Temple to the mind")
	PsionicsWorshipslP = append(PsionicsWorshipslP, "Sanitarium-prison for feral psychics")
	PlaceTagMap["PsionicsWorship"] = PsionicsWorshipslP

	QuarantinedWorldslP := []string{}
	QuarantinedWorldslP = append(QuarantinedWorldslP, "Bridge of a blockading ship")
	QuarantinedWorldslP = append(QuarantinedWorldslP, "Defense installation control room")
	QuarantinedWorldslP = append(QuarantinedWorldslP, "Refugee camp")
	PlaceTagMap["QuarantinedWorld"] = QuarantinedWorldslP

	RadioactiveWorldslP := []string{}
	RadioactiveWorldslP = append(RadioactiveWorldslP, "Mutant-infested ruins")
	RadioactiveWorldslP = append(RadioactiveWorldslP, "Scorched glass plain")
	RadioactiveWorldslP = append(RadioactiveWorldslP, "Wilderness of bizarre native life")
	RadioactiveWorldslP = append(RadioactiveWorldslP, "Glowing barrens")
	PlaceTagMap["RadioactiveWorld"] = RadioactiveWorldslP

	RefugeesslP := []string{}
	RefugeesslP = append(RefugeesslP, "Hopeless refugee camp")
	RefugeesslP = append(RefugeesslP, "City swarming with confused strangers")
	RefugeesslP = append(RefugeesslP, "Festival full of angry locals")
	PlaceTagMap["Refugees"] = RefugeesslP

	RegionalHegemonslP := []string{}
	RegionalHegemonslP = append(RegionalHegemonslP, "Palace or seat of government")
	RegionalHegemonslP = append(RegionalHegemonslP, "Salon teeming with spies")
	RegionalHegemonslP = append(RegionalHegemonslP, "Protest rally")
	RegionalHegemonslP = append(RegionalHegemonslP, "Military base")
	PlaceTagMap["RegionalHegemon"] = RegionalHegemonslP

	RestrictiveLawsslP := []string{}
	RestrictiveLawsslP = append(RestrictiveLawsslP, "Courtroom")
	RestrictiveLawsslP = append(RestrictiveLawsslP, "Mob scene of outraged locals")
	RestrictiveLawsslP = append(RestrictiveLawsslP, "Legislative chamber")
	RestrictiveLawsslP = append(RestrictiveLawsslP, "Police station")
	PlaceTagMap["RestrictiveLaws"] = RestrictiveLawsslP

	RevanchistsslP := []string{}
	RevanchistsslP = append(RevanchistsslP, "Memorial monument to the loss")
	RevanchistsslP = append(RevanchistsslP, "Cemetery of those who died in the conquest")
	RevanchistsslP = append(RevanchistsslP, "Public ceremony commemorating the disaster")
	PlaceTagMap["Revanchists"] = RevanchistsslP

	RevolutionariesslP := []string{}
	RevolutionariesslP = append(RevolutionariesslP, "Festival that explodes into violence")
	RevolutionariesslP = append(RevolutionariesslP, "Heavily-fortified police station")
	RevolutionariesslP = append(RevolutionariesslP, "Revolutionary base hidden in the wilderness")
	PlaceTagMap["Revolutionaries"] = RevolutionariesslP

	RigidCultureslP := []string{}
	RigidCultureslP = append(RigidCultureslP, "Time-worn palace")
	RigidCultureslP = append(RigidCultureslP, "Low-caste slums")
	RigidCultureslP = append(RigidCultureslP, "Bandit den")
	RigidCultureslP = append(RigidCultureslP, "Reformist temple")
	PlaceTagMap["RigidCulture"] = RigidCultureslP

	RisingHegemonslP := []string{}
	RisingHegemonslP = append(RisingHegemonslP, "Rustic town being hurled into prosperity")
	RisingHegemonslP = append(RisingHegemonslP, "Government building being expanded")
	RisingHegemonslP = append(RisingHegemonslP, "Starport struggling under the flow of new ships")
	PlaceTagMap["RisingHegemon"] = RisingHegemonslP

	RitualCombatslP := []string{}
	RitualCombatslP = append(RitualCombatslP, "Area full of cheering spectators")
	RitualCombatslP = append(RitualCombatslP, "Dusty street outside a saloon")
	RitualCombatslP = append(RitualCombatslP, "Memorial for fallen warriors")
	PlaceTagMap["RitualCombat"] = RitualCombatslP

	RobotsslP := []string{}
	RobotsslP = append(RobotsslP, "Humming robotic factory")
	RobotsslP = append(RobotsslP, "Stark robotic “barracks”")
	RobotsslP = append(RobotsslP, "House crowded with robot servants and only one human owner")
	PlaceTagMap["Robots"] = RobotsslP

	SeagoingCitiesslP := []string{}
	SeagoingCitiesslP = append(SeagoingCitiesslP, "Bridge of the city")
	SeagoingCitiesslP = append(SeagoingCitiesslP, "Storm-tossed sea")
	SeagoingCitiesslP = append(SeagoingCitiesslP, "A bridge fashioned of many small boats.")
	PlaceTagMap["SeagoingCities"] = SeagoingCitiesslP

	SealedMenaceslP := []string{}
	SealedMenaceslP = append(SealedMenaceslP, "Guarded fortress containing the menace")
	SealedMenaceslP = append(SealedMenaceslP, "Monitoring station")
	SealedMenaceslP = append(SealedMenaceslP, "Scene of a prior outbreak of the menace")
	PlaceTagMap["SealedMenace"] = SealedMenaceslP

	SecretMastersslP := []string{}
	SecretMastersslP = append(SecretMastersslP, "Smoke-filled room")
	SecretMastersslP = append(SecretMastersslP, "Shadowy alleyway")
	SecretMastersslP = append(SecretMastersslP, "Secret underground bunker")
	PlaceTagMap["SecretMasters"] = SecretMastersslP

	SectariansslP := []string{}
	SectariansslP = append(SectariansslP, "Sectarian battlefield")
	SectariansslP = append(SectariansslP, "Crusading temple")
	SectariansslP = append(SectariansslP, "Philosopher’s salon")
	SectariansslP = append(SectariansslP, "Bitterly divided village")
	PlaceTagMap["Sectarians"] = SectariansslP

	SeismicInstabilityslP := []string{}
	SeismicInstabilityslP = append(SeismicInstabilityslP, "Volcanic caldera")
	SeismicInstabilityslP = append(SeismicInstabilityslP, "Village during an earthquake")
	SeismicInstabilityslP = append(SeismicInstabilityslP, "Mud slide")
	SeismicInstabilityslP = append(SeismicInstabilityslP, "Earthquake opening superheated steam fissures")
	PlaceTagMap["SeismicInstability"] = SeismicInstabilityslP

	ShackledWorldslP := []string{}
	ShackledWorldslP = append(ShackledWorldslP, "Grim high-tech control center")
	ShackledWorldslP = append(ShackledWorldslP, "Factory full of workaround tech")
	ShackledWorldslP = append(ShackledWorldslP, "Temple to the power or entity that imposed the shackle")
	PlaceTagMap["ShackledWorld"] = ShackledWorldslP

	SocietalDespairslP := []string{}
	SocietalDespairslP = append(SocietalDespairslP, "Empty temple")
	SocietalDespairslP = append(SocietalDespairslP, "Crowded den of obliviating vice")
	SocietalDespairslP = append(SocietalDespairslP, "Smoky hall full of frantic speakers")
	PlaceTagMap["SocietalDespair"] = SocietalDespairslP

	SoleSupplierslP := []string{}
	SoleSupplierslP = append(SoleSupplierslP, "Bustling resource extraction site")
	SoleSupplierslP = append(SoleSupplierslP, "Opulent palace built with resource money")
	SoleSupplierslP = append(SoleSupplierslP, "Lazy town square where everyone lives on resource payments")
	PlaceTagMap["SoleSupplier"] = SoleSupplierslP

	TabooTreasureslP := []string{}
	TabooTreasureslP = append(TabooTreasureslP, "Den where the good is used")
	TabooTreasureslP = append(TabooTreasureslP, "Market selling the good to locals and a few outsiders")
	TabooTreasureslP = append(TabooTreasureslP, "Factory or processing area where the good is created")
	PlaceTagMap["TabooTreasure"] = TabooTreasureslP

	TerraformFailureslP := []string{}
	TerraformFailureslP = append(TerraformFailureslP, "Zone of tolerable gravity or temperature")
	TerraformFailureslP = append(TerraformFailureslP, "Native settlement built to cope with the environment")
	TerraformFailureslP = append(TerraformFailureslP, "Massive ruined terraforming engine")
	PlaceTagMap["TerraformFailure"] = TerraformFailureslP

	TheocracyslP := []string{}
	TheocracyslP = append(TheocracyslP, "Glorious temple")
	TheocracyslP = append(TheocracyslP, "Austere monastery")
	TheocracyslP = append(TheocracyslP, "Academy for ideological indoctrination")
	TheocracyslP = append(TheocracyslP, "Decadent pleasure-cathedral")
	PlaceTagMap["Theocracy"] = TheocracyslP

	TombWorldslP := []string{}
	TombWorldslP = append(TombWorldslP, "Crumbling hive-city")
	TombWorldslP = append(TombWorldslP, "City square carpeted in bones")
	TombWorldslP = append(TombWorldslP, "Ruined hydroponic facility")
	TombWorldslP = append(TombWorldslP, "Cannibal tribe’s lair")
	TombWorldslP = append(TombWorldslP, "Dead orbital jump gate")
	PlaceTagMap["TombWorld"] = TombWorldslP

	TradeHubslP := []string{}
	TradeHubslP = append(TradeHubslP, "Raucous bazaar")
	TradeHubslP = append(TradeHubslP, "Elegant restaurant")
	TradeHubslP = append(TradeHubslP, "Spaceport teeming with activity")
	TradeHubslP = append(TradeHubslP, "Foggy street lined with warehouses")
	PlaceTagMap["TradeHub"] = TradeHubslP

	TyrannyslP := []string{}
	TyrannyslP = append(TyrannyslP, "Impoverished village")
	TyrannyslP = append(TyrannyslP, "Protest rally massacre")
	TyrannyslP = append(TyrannyslP, "Decadent palace")
	TyrannyslP = append(TyrannyslP, "Religious hospital for the indigent")
	PlaceTagMap["Tyranny"] = TyrannyslP

	UnbrakedAIslP := []string{}
	UnbrakedAIslP = append(UnbrakedAIslP, "Municipal computing banks")
	UnbrakedAIslP = append(UnbrakedAIslP, "Cult compound")
	UnbrakedAIslP = append(UnbrakedAIslP, "Repair center")
	UnbrakedAIslP = append(UnbrakedAIslP, "Ancient hardcopy library")
	PlaceTagMap["UnbrakedAI"] = UnbrakedAIslP

	UrbanizedSurfaceslP := []string{}
	UrbanizedSurfaceslP = append(UrbanizedSurfaceslP, "Giant hab block now devoid of inhabitants")
	UrbanizedSurfaceslP = append(UrbanizedSurfaceslP, "Chemical-reeking underway")
	UrbanizedSurfaceslP = append(UrbanizedSurfaceslP, "Seawater mine full of salt and massive flowing channels")
	PlaceTagMap["UrbanizedSurface"] = UrbanizedSurfaceslP

	UtopiaslP := []string{}
	UtopiaslP = append(UtopiaslP, "Plaza full of altered humans")
	UtopiaslP = append(UtopiaslP, "Social ritual site")
	UtopiaslP = append(UtopiaslP, "Secret office where “normal” humans rule")
	PlaceTagMap["Utopia"] = UtopiaslP

	WarlordsslP := []string{}
	WarlordsslP = append(WarlordsslP, "Gory battlefield")
	WarlordsslP = append(WarlordsslP, "Burnt-out village")
	WarlordsslP = append(WarlordsslP, "Barbaric warlord palace")
	WarlordsslP = append(WarlordsslP, "Squalid refugee camp")
	PlaceTagMap["Warlords"] = WarlordsslP

	XenophilesslP := []string{}
	XenophilesslP = append(XenophilesslP, "Alien district")
	XenophilesslP = append(XenophilesslP, "Alien-influenced human home")
	XenophilesslP = append(XenophilesslP, "Cultural festival celebrating alien artist")
	PlaceTagMap["Xenophiles"] = XenophilesslP

	XenophobesslP := []string{}
	XenophobesslP = append(XenophobesslP, "Sealed treaty port")
	XenophobesslP = append(XenophobesslP, "Public ritual not open to outsiders")
	XenophobesslP = append(XenophobesslP, "Outcaste slum home")
	PlaceTagMap["Xenophobes"] = XenophobesslP

	ZombiesslP := []string{}
	ZombiesslP = append(ZombiesslP, "House with boarded-up windows")
	ZombiesslP = append(ZombiesslP, "Dead city")
	ZombiesslP = append(ZombiesslP, "Fortified bunker that was overrun from within")
	PlaceTagMap["Zombies"] = ZombiesslP
	return PlaceTagMap
}

func createTradeAuthorityMap() map[string][]string {
	tAthorityTagMap := make(map[string][]string)
	AlientAthority := []string{}
	AlientAthority = append(AlientAthority, "Member of a special trading caste")
	AlientAthority = append(AlientAthority, "Bureaucrat that communicates in strange idioms")
	AlientAthority = append(AlientAthority, "Judicial official of an incomprehensible law")
	AlientAthority = append(AlientAthority, "Gone-native human transplant with full authority over human traders")
	tAthorityTagMap["Alien"] = AlientAthority

	ClosedtAthority := []string{}
	ClosedtAthority = append(ClosedtAthority, "Zealously xenophobic security officer")
	ClosedtAthority = append(ClosedtAthority, "Guardian of cultural purity")
	ClosedtAthority = append(ClosedtAthority, "Interstellar quarantine officer")
	ClosedtAthority = append(ClosedtAthority, "Brutish monopolist")
	tAthorityTagMap["Closed"] = ClosedtAthority

	CommunisttAthority := []string{}
	CommunisttAthority = append(CommunisttAthority, "Eager bureaucrat with a poor grasp on private property")
	CommunisttAthority = append(CommunisttAthority, "Cynical apparatchik who just wants his share")
	CommunisttAthority = append(CommunisttAthority, "Local factory manager desperate to make quota")
	CommunisttAthority = append(CommunisttAthority, "'Fixer' for the regime looking for vital spare parts")
	tAthorityTagMap["Communist"] = CommunisttAthority

	DisorganizedtAthority := []string{}
	DisorganizedtAthority = append(DisorganizedtAthority, "Egomaniacal prophet of the new dispensation")
	DisorganizedtAthority = append(DisorganizedtAthority, "Desperate remnant-lord of the old order")
	DisorganizedtAthority = append(DisorganizedtAthority, "Survivalist tribe chieftain")
	DisorganizedtAthority = append(DisorganizedtAthority, "Leader of a recognized sanctuary from the chaos")
	tAthorityTagMap["Disorganized"] = DisorganizedtAthority

	DyingtAthority := []string{}
	DyingtAthority = append(DyingtAthority, "Sweating local ruler bent on escape for their family")
	DyingtAthority = append(DyingtAthority, "Grim chooser of the slain apportioning dwindling resources")
	DyingtAthority = append(DyingtAthority, "Military officer providing security for the remaining far traders")
	DyingtAthority = append(DyingtAthority, "Resigned religious leader working to save his flock")
	tAthorityTagMap["Dying"] = DyingtAthority

	FractioustAthority := []string{}
	FractioustAthority = append(FractioustAthority, "De facto but discreet faction leader")
	FractioustAthority = append(FractioustAthority, "Figurehead leader controlled by hidden masters")
	FractioustAthority = append(FractioustAthority, "Eminently corruptible bureaucrat working for the highest bidder")
	FractioustAthority = append(FractioustAthority, "Desperate faction member on his way down the ladder of power")
	tAthorityTagMap["Fractious"] = FractioustAthority

	KleptocratictAthority := []string{}
	KleptocratictAthority = append(KleptocratictAthority, "Corpulent port master with detestable tastes")
	KleptocratictAthority = append(KleptocratictAthority, "Local trade official who doubles as a crime boss")
	KleptocratictAthority = append(KleptocratictAthority, "Supervisor of a security department that's just another gang")
	KleptocratictAthority = append(KleptocratictAthority, "Paranoid bureaucrat being pressured by numerous murderous 'friends'")
	tAthorityTagMap["Kleptocratic"] = KleptocratictAthority

	LaissezFairetAthority := []string{}
	LaissezFairetAthority = append(LaissezFairetAthority, "Entrepreneurially-minded bureaucrat")
	LaissezFairetAthority = append(LaissezFairetAthority, "Harried goods inspector")
	LaissezFairetAthority = append(LaissezFairetAthority, "Official hunter of offworld rarities for the government")
	LaissezFairetAthority = append(LaissezFairetAthority, "Hidebound administrator relic of a more mercantilist age")
	tAthorityTagMap["LaissezFaire"] = LaissezFairetAthority

	MegacorpstAthority := []string{}
	MegacorpstAthority = append(MegacorpstAthority, "Colorless Mr. Johnson with trust issues")
	MegacorpstAthority = append(MegacorpstAthority, "Government bureaucrat working for a patron corp")
	MegacorpstAthority = append(MegacorpstAthority, "Military officer wanting hardware the corps don't have")
	MegacorpstAthority = append(MegacorpstAthority, "Utterly mercenary port master")
	tAthorityTagMap["Megacorps"] = MegacorpstAthority

	MilitarytAthority := []string{}
	MilitarytAthority = append(MilitarytAthority, "Quartermaster military official")
	MilitarytAthority = append(MilitarytAthority, "Military port police")
	MilitarytAthority = append(MilitarytAthority, "Internal security spy-hunter")
	MilitarytAthority = append(MilitarytAthority, "Black ops supply officer")
	tAthorityTagMap["Military"] = MilitarytAthority

	OpenedtAthority := []string{}
	OpenedtAthority = append(OpenedtAthority, "Confused minister of the local trade bureaucracy")
	OpenedtAthority = append(OpenedtAthority, "Personally ambitious official")
	OpenedtAthority = append(OpenedtAthority, "Leader of the faction that advocated for opening the world")
	OpenedtAthority = append(OpenedtAthority, "New port master with no idea how to do his job")
	tAthorityTagMap["Opened"] = OpenedtAthority

	PanopticontAthority := []string{}
	PanopticontAthority = append(PanopticontAthority, "Spaceport monitor official")
	PanopticontAthority = append(PanopticontAthority, "AI surveillance controller")
	PanopticontAthority = append(PanopticontAthority, "Official empowered to turn off surveillance")
	PanopticontAthority = append(PanopticontAthority, "Local elite of the Great Unwatched")
	tAthorityTagMap["Panopticon"] = PanopticontAthority

	PrimitivetAthority := []string{}
	PrimitivetAthority = append(PrimitivetAthority, "Confused local autocrat")
	PrimitivetAthority = append(PrimitivetAthority, "Unrelentingly hidebound guildmaster")
	PrimitivetAthority = append(PrimitivetAthority, "Local official with wildly inflated expectations about far traders")
	PrimitivetAthority = append(PrimitivetAthority, "Local grandee unable to understand that her authority is not universal")
	tAthorityTagMap["Primitive"] = PrimitivetAthority

	RestrictedtAthority := []string{}
	RestrictedtAthority = append(RestrictedtAthority, "AI robot of the synthetic ruling class")
	RestrictedtAthority = append(RestrictedtAthority, "Tech-priest devoted to the mysteries")
	RestrictedtAthority = append(RestrictedtAthority, "Cynical warlord with no delusions about the tech")
	RestrictedtAthority = append(RestrictedtAthority, "Ruthless inquisitor rooting out illicit tech")
	tAthorityTagMap["Restricted"] = RestrictedtAthority

	ScarcitytAthority := []string{}
	ScarcitytAthority = append(ScarcitytAthority, "Barbarous chieftain adorned in tokens of the material")
	ScarcitytAthority = append(ScarcitytAthority, "Grim gray technocrat in charge of distribution")
	ScarcitytAthority = append(ScarcitytAthority, "Engineering official responsible for production")
	ScarcitytAthority = append(ScarcitytAthority, "Ritualized semi-priestly official in charge of the material")
	tAthorityTagMap["Scarcity"] = ScarcitytAthority

	SecrettAthority := []string{}
	SecrettAthority = append(SecrettAthority, "Government trade official in black")
	SecrettAthority = append(SecrettAthority, "UFO enthusiast with a secret landing site")
	SecrettAthority = append(SecrettAthority, "Cynical business tycoon")
	SecrettAthority = append(SecrettAthority, "Illuminati minion with business connections")
	tAthorityTagMap["Secret"] = SecrettAthority

	SophisticatedtAthority := []string{}
	SophisticatedtAthority = append(SophisticatedtAthority, "Supercilious official in need of 'favors'")
	SophisticatedtAthority = append(SophisticatedtAthority, "Disinterested mandarin unconvinced of trade's value")
	SophisticatedtAthority = append(SophisticatedtAthority, "Blase local tycoon wanting new products")
	SophisticatedtAthority = append(SophisticatedtAthority, "Jaded entertainer seeking fresh inspiration")
	tAthorityTagMap["Sophisticated"] = SophisticatedtAthority

	TheocratictAthority := []string{}
	TheocratictAthority = append(TheocratictAthority, "Overfed priest-bureaucrat with very secular concerns")
	TheocratictAthority = append(TheocratictAthority, "Underpriestess looking to make a vacancy above her")
	TheocratictAthority = append(TheocratictAthority, "Purely nominal cleric-official concerned only with the social forms")
	TheocratictAthority = append(TheocratictAthority, "Leader of a persecuted sect interested in offworld allies")
	tAthorityTagMap["Theocratic"] = TheocratictAthority

	ThrivingtAthority := []string{}
	ThrivingtAthority = append(ThrivingtAthority, "Befuddled official suddenly responsible for much more trade than ever before")
	ThrivingtAthority = append(ThrivingtAthority, "Nouveau-riche tycoon with wild ambitions")
	ThrivingtAthority = append(ThrivingtAthority, "Local desperate for offworlder help in exploiting an opportunity")
	ThrivingtAthority = append(ThrivingtAthority, "Deeply suspicious old-guard elite")
	tAthorityTagMap["Thriving"] = ThrivingtAthority

	TributetAthority := []string{}
	TributetAthority = append(TributetAthority, "Dead-eyed tribute collector")
	TributetAthority = append(TributetAthority, "Official in charge of handling offworlder contributions")
	TributetAthority = append(TributetAthority, "Collector of 'special orders' from the elite")
	TributetAthority = append(TributetAthority, "Local desperate to make his quota somehow")
	tAthorityTagMap["Tribute"] = TributetAthority

	TyrannicaltAthority := []string{}
	TyrannicaltAthority = append(TyrannicaltAthority, "Gray-souled bureaucrat seeking an edge")
	TyrannicaltAthority = append(TyrannicaltAthority, "Overfed scion of the ruling class")
	TyrannicaltAthority = append(TyrannicaltAthority, "Personal shopper for the tyrant")
	TyrannicaltAthority = append(TyrannicaltAthority, "Military officer seeking special orders")
	tAthorityTagMap["Tyrannical"] = TyrannicaltAthority

	UsurpedtAthority := []string{}
	UsurpedtAthority = append(UsurpedtAthority, "Colonial customs officer with a distaste for the locals")
	UsurpedtAthority = append(UsurpedtAthority, "Plastic-souled combine bureaucrat who wants off the world")
	UsurpedtAthority = append(UsurpedtAthority, "Local collaborator hated by their neighbors")
	UsurpedtAthority = append(UsurpedtAthority, "'Negotiator' who handles all contact with buyers and sellers")
	tAthorityTagMap["Usurped"] = UsurpedtAthority

	VendortAthority := []string{}
	VendortAthority = append(VendortAthority, "Chief manufacturer of the good")
	VendortAthority = append(VendortAthority, "Religious official in charge of the good's disbursal to the populace")
	VendortAthority = append(VendortAthority, "Labor union chief involved with making the good")
	VendortAthority = append(VendortAthority, "Tax collector bent on stopping the smuggling of the good")
	tAthorityTagMap["Vendor"] = VendortAthority

	XenophobictAthority := []string{}
	XenophobictAthority = append(XenophobictAthority, "Profoundly disgusted bureaucrat")
	XenophobictAthority = append(XenophobictAthority, "Furtive corporate executive fearing discovery")
	XenophobictAthority = append(XenophobictAthority, "Official who thinks the PCs have no moral reservations about any work")
	XenophobictAthority = append(XenophobictAthority, "Desperate local willing to deal with the devil for some vital thing")
	tAthorityTagMap["Xenophobic"] = XenophobictAthority
	return tAthorityTagMap
}

func createTradeAntogonistMap() map[string][]string {
	tAntagonistTagMap := make(map[string][]string)
	AlientAntagonist := []string{}
	AlientAntagonist = append(AlientAntagonist, "Amoral far trader selling the aliens tech dangerous to humans")
	AlientAntagonist = append(AlientAntagonist, "Alien renegade seeking offworld goods for its plots")
	AlientAntagonist = append(AlientAntagonist, "Alien trader trying to sell local artifacts forbidden from export")
	AlientAntagonist = append(AlientAntagonist, "Alien selling goods that are actually very dangerous")
	tAntagonistTagMap["Alien"] = AlientAntagonist

	ClosedtAntagonist := []string{}
	ClosedtAntagonist = append(ClosedtAntagonist, "Corrupt local potentate")
	ClosedtAntagonist = append(ClosedtAntagonist, "Bribed quarantine enforcer")
	ClosedtAntagonist = append(ClosedtAntagonist, "Relentless far trader smuggler")
	ClosedtAntagonist = append(ClosedtAntagonist, "Native aspiring to break the seal")
	tAntagonistTagMap["Closed"] = ClosedtAntagonist

	CommunisttAntagonist := []string{}
	CommunisttAntagonist = append(CommunisttAntagonist, "Vengeful survivor of the former propertied class")
	CommunisttAntagonist = append(CommunisttAntagonist, "'Enemy of the people' out to overthrow the government")
	CommunisttAntagonist = append(CommunisttAntagonist, "Far trader with moral objections to dealing with communists")
	CommunisttAntagonist = append(CommunisttAntagonist, "Aspiring tyrant seeking to seize control of the world")
	tAntagonistTagMap["Communist"] = CommunisttAntagonist

	DisorganizedtAntagonist := []string{}
	DisorganizedtAntagonist = append(DisorganizedtAntagonist, "Aspiring offworlder planetary warlord")
	DisorganizedtAntagonist = append(DisorganizedtAntagonist, "Offworld government agent seeking control")
	DisorganizedtAntagonist = append(DisorganizedtAntagonist, "Local lord cooperating with an amoral trader")
	DisorganizedtAntagonist = append(DisorganizedtAntagonist, "Organlegger trading copious spare parts")
	tAntagonistTagMap["Disorganized"] = DisorganizedtAntagonist

	DyingtAntagonist := []string{}
	DyingtAntagonist = append(DyingtAntagonist, "Desperate local grandee bent on ship theft")
	DyingtAntagonist = append(DyingtAntagonist, "Ruthless local ready to do absolutely anything to get their loved ones away")
	DyingtAntagonist = append(DyingtAntagonist, "Furious demagogue convinced the doom is the work of offworlders")
	DyingtAntagonist = append(DyingtAntagonist, "Cheat selling passage on the PCs' ship without telling them")
	tAntagonistTagMap["Dying"] = DyingtAntagonist

	FractioustAntagonist := []string{}
	FractioustAntagonist = append(FractioustAntagonist, "Faction member seeking to discredit a troublesome superior")
	FractioustAntagonist = append(FractioustAntagonist, "Saboteur seeking offworlder weapons")
	FractioustAntagonist = append(FractioustAntagonist, "Reformer seeking offworld resources")
	FractioustAntagonist = append(FractioustAntagonist, "Renegade seeking a plague on all their houses")
	tAntagonistTagMap["Fractious"] = FractioustAntagonist

	KleptocratictAntagonist := []string{}
	KleptocratictAntagonist = append(KleptocratictAntagonist, "Abominably corrupt ruler who demands loathsome tribute")
	KleptocratictAntagonist = append(KleptocratictAntagonist, "Hard-bitten underclass entrepreneur with no god but profit")
	KleptocratictAntagonist = append(KleptocratictAntagonist, "Offworld reformer who blames far traders as the root of the problem")
	KleptocratictAntagonist = append(KleptocratictAntagonist, "Savagely vengeful renegade")
	tAntagonistTagMap["Kleptocratic"] = KleptocratictAntagonist

	LaissezFairetAntagonist := []string{}
	LaissezFairetAntagonist = append(LaissezFairetAntagonist, "Amoral merchant trading in foul things")
	LaissezFairetAntagonist = append(LaissezFairetAntagonist, "Aspiring monopolist seeking to control the market")
	LaissezFairetAntagonist = append(LaissezFairetAntagonist, "Lobbyist seeking to control the regulators")
	LaissezFairetAntagonist = append(LaissezFairetAntagonist, "Failed merchant who blames far traders")
	tAntagonistTagMap["LaissezFaire"] = LaissezFairetAntagonist

	MegacorpstAntagonist := []string{}
	MegacorpstAntagonist = append(MegacorpstAntagonist, "Industrial saboteur working for a rival")
	MegacorpstAntagonist = append(MegacorpstAntagonist, "Official who hates 'destabilizing' far traders")
	MegacorpstAntagonist = append(MegacorpstAntagonist, "Violent anti-corp terroristslash-freedom fighter")
	MegacorpstAntagonist = append(MegacorpstAntagonist, "Megacorp creation gone rogue ")
	tAntagonistTagMap["Megacorps"] = MegacorpstAntagonist

	MilitarytAntagonist := []string{}
	MilitarytAntagonist = append(MilitarytAntagonist, "Zealot confiscating goods 'for the war effort'")
	MilitarytAntagonist = append(MilitarytAntagonist, "Ruthless enforcer of military quarantine")
	MilitarytAntagonist = append(MilitarytAntagonist, "Enemy saboteur bent on ruining far traders")
	MilitarytAntagonist = append(MilitarytAntagonist, "Deserter looking to make a profit")
	tAntagonistTagMap["Military"] = MilitarytAntagonist

	OpenedtAntagonist := []string{}
	OpenedtAntagonist = append(OpenedtAntagonist, "Angry xenophobic local")
	OpenedtAntagonist = append(OpenedtAntagonist, "Ruthlessly exploitative far trader rival")
	OpenedtAntagonist = append(OpenedtAntagonist, "Offworld ambassador with a grudge against the far trader")
	OpenedtAntagonist = append(OpenedtAntagonist, "Local ruler who expects pliant service from the offworlders")
	tAntagonistTagMap["Opened"] = OpenedtAntagonist

	PanopticontAntagonist := []string{}
	PanopticontAntagonist = append(PanopticontAntagonist, "Monitor saboteur")
	PanopticontAntagonist = append(PanopticontAntagonist, "Official paranoid about unobserved far traders")
	PanopticontAntagonist = append(PanopticontAntagonist, "Offworlder seeking to subvert the monitors for their own purposes")
	PanopticontAntagonist = append(PanopticontAntagonist, "Avid blackmailer")
	tAntagonistTagMap["Panopticon"] = PanopticontAntagonist

	PrimitivetAntagonist := []string{}
	PrimitivetAntagonist = append(PrimitivetAntagonist, "Official privately determined to drive the offworlders away")
	PrimitivetAntagonist = append(PrimitivetAntagonist, "Local crime boss desperate to steal offworld goods")
	PrimitivetAntagonist = append(PrimitivetAntagonist, "Local demagogue bent on extorting the offworlders")
	PrimitivetAntagonist = append(PrimitivetAntagonist, "Far trader who really dislikes competition")
	tAntagonistTagMap["Primitive"] = PrimitivetAntagonist

	RestrictedtAntagonist := []string{}
	RestrictedtAntagonist = append(RestrictedtAntagonist, "Rebel bent on getting a new source for tech")
	RestrictedtAntagonist = append(RestrictedtAntagonist, "Elite in desperate need of replacements for dying tech")
	RestrictedtAntagonist = append(RestrictedtAntagonist, "Far trader inclined to destroy the monopoly")
	RestrictedtAntagonist = append(RestrictedtAntagonist, "Colonial offworld official enforcing the rules on the hapless locals")
	tAntagonistTagMap["Restricted"] = RestrictedtAntagonist

	ScarcitytAntagonist := []string{}
	ScarcitytAntagonist = append(ScarcitytAntagonist, "Saboteur bent on destroying a rival's supply of the material")
	ScarcitytAntagonist = append(ScarcitytAntagonist, "Radical redistributionist with no fear of the consequences")
	ScarcitytAntagonist = append(ScarcitytAntagonist, "Far trader bent on keeping the material scarce")
	ScarcitytAntagonist = append(ScarcitytAntagonist, "Local elite trying to make sure the status remains quo")
	tAntagonistTagMap["Scarcity"] = ScarcitytAntagonist

	SecrettAntagonist := []string{}
	SecrettAntagonist = append(SecrettAntagonist, "Investigative reporter")
	SecrettAntagonist = append(SecrettAntagonist, "Business competitor baffled at the new products")
	SecrettAntagonist = append(SecrettAntagonist, "Rival government saboteur")
	SecrettAntagonist = append(SecrettAntagonist, "Relentless conspiracy theorist")
	tAntagonistTagMap["Secret"] = SecrettAntagonist

	SophisticatedtAntagonist := []string{}
	SophisticatedtAntagonist = append(SophisticatedtAntagonist, "Con-man playing on far traders' expectations of bumpkin locals")
	SophisticatedtAntagonist = append(SophisticatedtAntagonist, "Offworlder agent working indirectly against a rival power")
	SophisticatedtAntagonist = append(SophisticatedtAntagonist, "Decadent local demanding grim amusement")
	SophisticatedtAntagonist = append(SophisticatedtAntagonist, "Apparently wealthy local actually in dire financial straits")
	tAntagonistTagMap["Sophisticated"] = SophisticatedtAntagonist

	TheocratictAntagonist := []string{}
	TheocratictAntagonist = append(TheocratictAntagonist, "Extremely militant heretic")
	TheocratictAntagonist = append(TheocratictAntagonist, "Xenophobic zealot bent on purging offworlder corruption")
	TheocratictAntagonist = append(TheocratictAntagonist, "Offworld cultural saboteur seeking to cripple the world's society on a rival's behalf")
	TheocratictAntagonist = append(TheocratictAntagonist, "Far trader trying to move forbidden goods on the world")
	tAntagonistTagMap["Theocratic"] = TheocratictAntagonist

	ThrivingtAntagonist := []string{}
	ThrivingtAntagonist = append(ThrivingtAntagonist, "Suddenly wealthy local who plots revenge for old slights")
	ThrivingtAntagonist = append(ThrivingtAntagonist, "Far trader who wants to capture the local market")
	ThrivingtAntagonist = append(ThrivingtAntagonist, "Deposed elite who wants to end the 'chaos' of the boom")
	ThrivingtAntagonist = append(ThrivingtAntagonist, "Megalomaniac local who plans sector dominance")
	tAntagonistTagMap["Thriving"] = ThrivingtAntagonist

	TributetAntagonist := []string{}
	TributetAntagonist = append(TributetAntagonist, "Offworlder scheming to steal the tribute")
	TributetAntagonist = append(TributetAntagonist, "Local elite unable to understand that they can't simply take everything")
	TributetAntagonist = append(TributetAntagonist, "Deluded elite with impossible demands on far traders")
	TributetAntagonist = append(TributetAntagonist, "Ruthless tribute collector snatching whatever he can")
	tAntagonistTagMap["Tribute"] = TributetAntagonist

	TyrannicaltAntagonist := []string{}
	TyrannicaltAntagonist = append(TyrannicaltAntagonist, "Rebel who blames offworlders for the tyrant's power")
	TyrannicaltAntagonist = append(TyrannicaltAntagonist, "Brutal secret police chief with suspicions of offworlders")
	TyrannicaltAntagonist = append(TyrannicaltAntagonist, "Zealous party propaganda chief seeking offworlder support")
	TyrannicaltAntagonist = append(TyrannicaltAntagonist, "Suicidal saboteur out to destroy far trade")
	tAntagonistTagMap["Tyrannical"] = TyrannicaltAntagonist

	UsurpedtAntagonist := []string{}
	UsurpedtAntagonist = append(UsurpedtAntagonist, "Bitter local nationalist")
	UsurpedtAntagonist = append(UsurpedtAntagonist, "Reckless far trader smuggler")
	UsurpedtAntagonist = append(UsurpedtAntagonist, "Corrupt customs official")
	UsurpedtAntagonist = append(UsurpedtAntagonist, "Brutal trade enforcer ")
	tAntagonistTagMap["Usurped"] = UsurpedtAntagonist

	VendortAntagonist := []string{}
	VendortAntagonist = append(VendortAntagonist, "Brutal and red-handed smuggler")
	VendortAntagonist = append(VendortAntagonist, "Saboteur who seeks to destroy a supplier of the good")
	VendortAntagonist = append(VendortAntagonist, "Con man selling fake or adulterated goods")
	VendortAntagonist = append(VendortAntagonist, "Extortionist threatening to cut off the supply of the good")
	tAntagonistTagMap["Vendor"] = VendortAntagonist

	XenophobictAntagonist := []string{}
	XenophobictAntagonist = append(XenophobictAntagonist, "Purity-minded local demagogue")
	XenophobictAntagonist = append(XenophobictAntagonist, "Investigator for 'alien influences'")
	XenophobictAntagonist = append(XenophobictAntagonist, "Politician seeking to tie his rivals to offworlders")
	XenophobictAntagonist = append(XenophobictAntagonist, "Local who wants to bury his former contracts permanently")
	tAntagonistTagMap["Xenophobic"] = XenophobictAntagonist
	return tAntagonistTagMap
}

func createTradeThingsMap() map[string][]string {
	tThingsTagMap := make(map[string][]string)
	AlientThings := []string{}
	AlientThings = append(AlientThings, "Alien religious artifacts")
	AlientThings = append(AlientThings, "Exotic alien tech")
	AlientThings = append(AlientThings, "Alien slave caste")
	AlientThings = append(AlientThings, "Precious biological extract from dead aliens")
	tThingsTagMap["Alien"] = AlientThings

	ClosedtThings := []string{}
	ClosedtThings = append(ClosedtThings, "Blatantly offworld goods")
	ClosedtThings = append(ClosedtThings, "Uniquely precious cultural artifacts")
	ClosedtThings = append(ClosedtThings, "Key or tools for bypassing trade barriers")
	ClosedtThings = append(ClosedtThings, "Offworld goods disguised as precious local products")
	tThingsTagMap["Closed"] = ClosedtThings

	CommunisttThings := []string{}
	CommunisttThings = append(CommunisttThings, "Stock of overproduced goods")
	CommunisttThings = append(CommunisttThings, "Crude but tough equipment")
	CommunisttThings = append(CommunisttThings, "Last cache of some underproduced product on the planet")
	CommunisttThings = append(CommunisttThings, "Goods seized from capitalist roaders and wreckers")
	tThingsTagMap["Communist"] = CommunisttThings

	DisorganizedtThings := []string{}
	DisorganizedtThings = append(DisorganizedtThings, "Losers of local quarrels made into slaves")
	DisorganizedtThings = append(DisorganizedtThings, "Caches of the old government's weapons of mass destruction")
	DisorganizedtThings = append(DisorganizedtThings, "Stockpile of the resource the locals are fighting over")
	DisorganizedtThings = append(DisorganizedtThings, "Salvage from a ruined lab or city")
	tThingsTagMap["Disorganized"] = DisorganizedtThings

	DyingtThings := []string{}
	DyingtThings = append(DyingtThings, "A king's ransom gathered to buy passage for a now-dead ruler")
	DyingtThings = append(DyingtThings, "Precious cultural relics symbolic of the world")
	DyingtThings = append(DyingtThings, "Cold sleep pods for hundreds of infants")
	DyingtThings = append(DyingtThings, "A cache of usable spike drive cores")
	tThingsTagMap["Dying"] = DyingtThings

	FractioustThings := []string{}
	FractioustThings = append(FractioustThings, "Forbidden tech meant for a faction's leadership")
	FractioustThings = append(FractioustThings, "Restricted data stolen from a faction")
	FractioustThings = append(FractioustThings, "Precious resource the factions are fighting over")
	FractioustThings = append(FractioustThings, "Offworld tech shipment meant for 'freedom fighters'")
	tThingsTagMap["Fractious"] = FractioustThings

	KleptocratictThings := []string{}
	KleptocratictThings = append(KleptocratictThings, "Goods 'mislaid' as a bribe payment")
	KleptocratictThings = append(KleptocratictThings, "The forbidden tech that all the local factions want")
	KleptocratictThings = append(KleptocratictThings, "The goods due as protection money to keep someone from getting killed")
	KleptocratictThings = append(KleptocratictThings, "An offworld criminal's stash of looted goods")
	tThingsTagMap["Kleptocratic"] = KleptocratictThings

	LaissezFairetThings := []string{}
	LaissezFairetThings = append(LaissezFairetThings, "Forbidden goods unknown on other civilized worlds")
	LaissezFairetThings = append(LaissezFairetThings, "Recently-outlawed wares hidden for later pickup")
	LaissezFairetThings = append(LaissezFairetThings, "Defrauded goods cached for later claiming")
	LaissezFairetThings = append(LaissezFairetThings, "Cases of the One Thing Outlawed Here")
	tThingsTagMap["LaissezFaire"] = LaissezFairetThings

	MegacorpstThings := []string{}
	MegacorpstThings = append(MegacorpstThings, "Exotic R&D product")
	MegacorpstThings = append(MegacorpstThings, "Warehouse of goods left in legal limbo by a subsidiary")
	MegacorpstThings = append(MegacorpstThings, "Product of some brutal human experimentation")
	MegacorpstThings = append(MegacorpstThings, "Vital parts for a crucial manufacturing facility")
	tThingsTagMap["Megacorps"] = MegacorpstThings

	MilitarytThings := []string{}
	MilitarytThings = append(MilitarytThings, "High-tech military prototype gear")
	MilitarytThings = append(MilitarytThings, "Supplies vital to a beleaguered unit")
	MilitarytThings = append(MilitarytThings, "Crates of weapons banned by local laws of war")
	MilitarytThings = append(MilitarytThings, "Luxuries beloved of high brass")
	tThingsTagMap["Military"] = MilitarytThings

	OpenedtThings := []string{}
	OpenedtThings = append(OpenedtThings, "Pretech spare parts useless to the locals but priceless offworld")
	OpenedtThings = append(OpenedtThings, "Cargo container full of enormously disruptive offworld tech")
	OpenedtThings = append(OpenedtThings, "Access codes to a military space station inaccessible since the Scream")
	OpenedtThings = append(OpenedtThings, "Cache of goods that the locals don't realize are extremely valuable")
	tThingsTagMap["Opened"] = OpenedtThings

	PanopticontThings := []string{}
	PanopticontThings = append(PanopticontThings, "Monitor-proof shielding shipment")
	PanopticontThings = append(PanopticontThings, "Cargo of hyper-advanced surveillance gear")
	PanopticontThings = append(PanopticontThings, "Semi-intelligent monitoring hardware")
	PanopticontThings = append(PanopticontThings, "Entertaining but extremely illicit surveillance logs")
	tThingsTagMap["Panopticon"] = PanopticontThings

	PrimitivetThings := []string{}
	PrimitivetThings = append(PrimitivetThings, "Colony-era pretech parts that are useless to the locals")
	PrimitivetThings = append(PrimitivetThings, "Amazingly refined art or liquors")
	PrimitivetThings = append(PrimitivetThings, "Exquisitely-bred and trained warbeasts")
	PrimitivetThings = append(PrimitivetThings, "Papers proving ownership of local serfs or slaves")
	tThingsTagMap["Primitive"] = PrimitivetThings

	RestrictedtThings := []string{}
	RestrictedtThings = append(RestrictedtThings, "Cache of tech lost when the owner was killed")
	RestrictedtThings = append(RestrictedtThings, "'Cloaked' tech that looks like allowed objects")
	RestrictedtThings = append(RestrictedtThings, "Superior local tech reserved for elites")
	RestrictedtThings = append(RestrictedtThings, "Trove of the incredibly dangerous tech that provoked the restrictions in the first place")
	tThingsTagMap["Restricted"] = RestrictedtThings

	ScarcitytThings := []string{}
	ScarcitytThings = append(ScarcitytThings, "Secret warehouse of the material")
	ScarcitytThings = append(ScarcitytThings, "Material lost in fighting long ago")
	ScarcitytThings = append(ScarcitytThings, "Ancient tech that creates the material")
	ScarcitytThings = append(ScarcitytThings, "A virus or nanocontagion that destroys the material")
	tThingsTagMap["Scarcity"] = ScarcitytThings

	SecrettThings := []string{}
	SecrettThings = append(SecrettThings, "Slaves comprising people the locals want 'disappeared'")
	SecrettThings = append(SecrettThings, "Government-extracted resources")
	SecrettThings = append(SecrettThings, "Stealth technology to defeat local scanners")
	SecrettThings = append(SecrettThings, "Irrefutable evidence of 'alien' contact")
	tThingsTagMap["Secret"] = SecrettThings

	SophisticatedtThings := []string{}
	SophisticatedtThings = append(SophisticatedtThings, "Inexplicably compelling collection of artwork")
	SophisticatedtThings = append(SophisticatedtThings, "Magnificent relic of a long-lost civilization")
	SophisticatedtThings = append(SophisticatedtThings, "Local specialty jealously kept on-world")
	SophisticatedtThings = append(SophisticatedtThings, "Precious trove of data logs on now-lost societies")
	tThingsTagMap["Sophisticated"] = SophisticatedtThings

	TheocratictThings := []string{}
	TheocratictThings = append(TheocratictThings, "Beautiful religious artwork")
	TheocratictThings = append(TheocratictThings, "Technology sacred to the faith")
	TheocratictThings = append(TheocratictThings, "Priceless relics of the religion")
	TheocratictThings = append(TheocratictThings, "Stores of religious tithes that went missing")
	tThingsTagMap["Theocratic"] = TheocratictThings

	ThrivingtThings := []string{}
	ThrivingtThings = append(ThrivingtThings, "Products of an upgraded factory")
	ThrivingtThings = append(ThrivingtThings, "Motherlode of precious mineral")
	ThrivingtThings = append(ThrivingtThings, "Trove of goods bought by a now-dead entrepreneur")
	ThrivingtThings = append(ThrivingtThings, "Vital industrial equipment lost in transit")
	tThingsTagMap["Thriving"] = ThrivingtThings

	TributetThings := []string{}
	TributetThings = append(TributetThings, "Store of tribute material")
	TributetThings = append(TributetThings, "Tech to vastly increase tribute production")
	TributetThings = append(TributetThings, "Tech the elite use to help enforce the tribute")
	TributetThings = append(TributetThings, "Precious byproduct of tribute creation unknown to the locals")
	tThingsTagMap["Tribute"] = TributetThings

	TyrannicaltThings := []string{}
	TyrannicaltThings = append(TyrannicaltThings, "Wealth taken from enemies of the state")
	TyrannicaltThings = append(TyrannicaltThings, "Slave-produced goods")
	TyrannicaltThings = append(TyrannicaltThings, "Advanced surveillance tech")
	TyrannicaltThings = append(TyrannicaltThings, "Store of exceptionally vicious weaponry")
	tThingsTagMap["Tyrannical"] = TyrannicaltThings

	UsurpedtThings := []string{}
	UsurpedtThings = append(UsurpedtThings, "Tightly-restricted local specialty")
	UsurpedtThings = append(UsurpedtThings, "Cache of goods taken by the usurpers as profit")
	UsurpedtThings = append(UsurpedtThings, "Equipment vital to maintaining the trade stranglehold")
	UsurpedtThings = append(UsurpedtThings, "Weapons cache meant for local rebels ")
	tThingsTagMap["Usurped"] = UsurpedtThings

	VendortThings := []string{}
	VendortThings = append(VendortThings, "Secret untaxed cache of this good")
	VendortThings = append(VendortThings, "Machinery vital in its creation or extraction")
	VendortThings = append(VendortThings, "Prototype of a synthetic equivalent that threatens to destroy their monopoly")
	VendortThings = append(VendortThings, "Tainted or otherwise dangerous batch of this good")
	tThingsTagMap["Vendor"] = VendortThings

	XenophobictThings := []string{}
	XenophobictThings = append(XenophobictThings, "Store of goods too obviously alien to release to the market")
	XenophobictThings = append(XenophobictThings, "Local specialty kept from outsiders")
	XenophobictThings = append(XenophobictThings, "Cache of goods hidden by a lynched trader")
	XenophobictThings = append(XenophobictThings, "Local treasure of incalculable value to natives")
	tThingsTagMap["Xenophobic"] = XenophobictThings
	return tThingsTagMap
}

func createTradeComplicationsMap() map[string][]string {
	tComplicationsTagMap := make(map[string][]string)
	AlientComplications := []string{}
	AlientComplications = append(AlientComplications, "The locals suddenly become very dangerous due to alien ritual behavior or biological impulses")
	AlientComplications = append(AlientComplications, "Aliens suddenly demand a particular good to continue trading")
	AlientComplications = append(AlientComplications, "Aliens become enraged over inexplicable slight")
	AlientComplications = append(AlientComplications, "Alien rebels start fighting for control of human trade")
	tComplicationsTagMap["Alien"] = AlientComplications

	ClosedtComplications := []string{}
	ClosedtComplications = append(ClosedtComplications, "The planet was wide open until very recently")
	ClosedtComplications = append(ClosedtComplications, "An exceptionally vicious local ruler wants to open outside trade to support his regime")
	ClosedtComplications = append(ClosedtComplications, "The local culture forces people to profess public contempt for all offworld goods")
	ClosedtComplications = append(ClosedtComplications, "Seemingly innocuous offworld tech caused a horrific disaster")
	tComplicationsTagMap["Closed"] = ClosedtComplications

	CommunisttComplications := []string{}
	CommunisttComplications = append(CommunisttComplications, "The locals actually have the AIs or cornucopia factories to make it work tolerably")
	CommunisttComplications = append(CommunisttComplications, "The communism is for the commoners while the elite actually own property")
	CommunisttComplications = append(CommunisttComplications, "The communism is a thin facade over a grim totalitarian state")
	CommunisttComplications = append(CommunisttComplications, "The communism is recent and everything seems to be working well so far")
	tComplicationsTagMap["Communist"] = CommunisttComplications

	DisorganizedtComplications := []string{}
	DisorganizedtComplications = append(DisorganizedtComplications, "The local balkanization is relatively stable among numerous small powers")
	DisorganizedtComplications = append(DisorganizedtComplications, "The disorganization is peaceful in nature and the product of genteel disagreements")
	DisorganizedtComplications = append(DisorganizedtComplications, "The planet is experiencing an inevitable and brutal population cull")
	DisorganizedtComplications = append(DisorganizedtComplications, "The chaos is recent enough that the locals still have many artifacts of the old order")
	tComplicationsTagMap["Disorganized"] = DisorganizedtComplications

	DyingtComplications := []string{}
	DyingtComplications = append(DyingtComplications, "Only the world's elite realize that they are doomed")
	DyingtComplications = append(DyingtComplications, "Half the planet is still convinced there is hope")
	DyingtComplications = append(DyingtComplications, "Outside rivals actively seek the death of the world")
	DyingtComplications = append(DyingtComplications, "The doom can be averted with lost alien tech")
	tComplicationsTagMap["Dying"] = DyingtComplications

	FractioustComplications := []string{}
	FractioustComplications = append(FractioustComplications, "All of the factions honestly mean well")
	FractioustComplications = append(FractioustComplications, "One faction is clearly dominant but opposed by all the others")
	FractioustComplications = append(FractioustComplications, "One faction has heavy offworld support")
	FractioustComplications = append(FractioustComplications, "The factions are very good at hiding their membership")
	tComplicationsTagMap["Fractious"] = FractioustComplications

	KleptocratictComplications := []string{}
	KleptocratictComplications = append(KleptocratictComplications, "A reformist movement is growing stronger but is hostile to far traders")
	KleptocratictComplications = append(KleptocratictComplications, "The society used to be much less corrupt before a disaster or war")
	KleptocratictComplications = append(KleptocratictComplications, "The corruption is carefully and ruthlessly hidden from the populace")
	KleptocratictComplications = append(KleptocratictComplications, "Without the grease of corruption the local legal framework would paralyze society")
	tComplicationsTagMap["Kleptocratic"] = KleptocratictComplications

	LaissezFairetComplications := []string{}
	LaissezFairetComplications = append(LaissezFairetComplications, "The new liberty is fragile and old economic powers want the market restrained in their favor")
	LaissezFairetComplications = append(LaissezFairetComplications, "Economic losers in the market are growing violent")
	LaissezFairetComplications = append(LaissezFairetComplications, "A government faction fears the market is starting to threaten their power")
	LaissezFairetComplications = append(LaissezFairetComplications, "Society is thrown in flux by the sudden flood of new wealth")
	tComplicationsTagMap["LaissezFaire"] = LaissezFairetComplications

	MegacorpstComplications := []string{}
	MegacorpstComplications = append(MegacorpstComplications, "The corps are actually making the standard of living very good here")
	MegacorpstComplications = append(MegacorpstComplications, "The corps hold the key to some process vital to local survival")
	MegacorpstComplications = append(MegacorpstComplications, "The corps maintain functional pretech facilities")
	MegacorpstComplications = append(MegacorpstComplications, "The corps are actually deniable government proxies completely in thrall to the local rulers")
	tComplicationsTagMap["Megacorps"] = MegacorpstComplications

	MilitarytComplications := []string{}
	MilitarytComplications = append(MilitarytComplications, "All civilians are essentially slave labor unless they can enter the military caste")
	MilitarytComplications = append(MilitarytComplications, "The world is rapidly tearing itself apart with the fighting")
	MilitarytComplications = append(MilitarytComplications, "The dominant military power hates far traders as destabilizing wild cards")
	MilitarytComplications = append(MilitarytComplications, "Foreign worlds use the planet as a testing ground for military tech")
	tComplicationsTagMap["Military"] = MilitarytComplications

	OpenedtComplications := []string{}
	OpenedtComplications = append(OpenedtComplications, "The locals opened up trade because they need the far traders to save the planet from something")
	OpenedtComplications = append(OpenedtComplications, "The locals are convinced the far traders are Space Gods or something similar")
	OpenedtComplications = append(OpenedtComplications, "Another world's government wants them quarantined for economic or religious reasons")
	OpenedtComplications = append(OpenedtComplications, "The locals think far traders all belong to the same company- and are culpable for each others' sins")
	tComplicationsTagMap["Opened"] = OpenedtComplications

	PanopticontComplications := []string{}
	PanopticontComplications = append(PanopticontComplications, "Telepaths run an undetectable psychic comm network")
	PanopticontComplications = append(PanopticontComplications, "Locals actually communicate via subtle contextual cues rather than the words they speak")
	PanopticontComplications = append(PanopticontComplications, "The surveillance state is new and society is in turmoil over it")
	PanopticontComplications = append(PanopticontComplications, "The watchers are actually looking for something specific and mortally dangerous")
	tComplicationsTagMap["Panopticon"] = PanopticontComplications

	PrimitivetComplications := []string{}
	PrimitivetComplications = append(PrimitivetComplications, "A local priestess has cynically duped the locals into thinking her far trader partners are holy figures")
	PrimitivetComplications = append(PrimitivetComplications, "The world is on the cusp of an explosively disruptive industrial revolution")
	PrimitivetComplications = append(PrimitivetComplications, "The locals are Luddites for a reason that is or was rational")
	PrimitivetComplications = append(PrimitivetComplications, "The last far trader to come through left chaos in his wake")
	tComplicationsTagMap["Primitive"] = PrimitivetComplications

	RestrictedtComplications := []string{}
	RestrictedtComplications = append(RestrictedtComplications, "The tech is self-aware and using the elites")
	RestrictedtComplications = append(RestrictedtComplications, "The restricted tech is extraordinarily dangerous on this particular world")
	RestrictedtComplications = append(RestrictedtComplications, "The tech is vital to life and the elite's lynchpin of control")
	RestrictedtComplications = append(RestrictedtComplications, "The tech is fueled by something precious or deplorable")
	tComplicationsTagMap["Restricted"] = RestrictedtComplications

	ScarcitytComplications := []string{}
	ScarcitytComplications = append(ScarcitytComplications, "The scarcity is so recent that society is still coming to grips with it")
	ScarcitytComplications = append(ScarcitytComplications, "The locals believe that the scarcity uplifts them and must be maintained")
	ScarcitytComplications = append(ScarcitytComplications, "A crazed demagogue accuses the rulers of having a superabundant cache of the material")
	ScarcitytComplications = append(ScarcitytComplications, "A nigh-unlimited supply or generation tech lies locked away behind ancient guards")
	tComplicationsTagMap["Scarcity"] = ScarcitytComplications

	SecrettComplications := []string{}
	SecrettComplications = append(SecrettComplications, "The locals are convinced that space travel is impossible since the Scream")
	SecrettComplications = append(SecrettComplications, "The locals are actually hyper-advanced and extremely isolationist")
	SecrettComplications = append(SecrettComplications, "The planet's history gives them good reason to hate and fear offworlders")
	SecrettComplications = append(SecrettComplications, "The locals are extremely hard to infiltrate")
	tComplicationsTagMap["Secret"] = SecrettComplications

	SophisticatedtComplications := []string{}
	SophisticatedtComplications = append(SophisticatedtComplications, "The locals are actually simply invincibly arrogant rather than accomplished")
	SophisticatedtComplications = append(SophisticatedtComplications, "They are cultural heirs to a very strange society or alien race")
	SophisticatedtComplications = append(SophisticatedtComplications, "The locals are convinced that they should rule the sector")
	SophisticatedtComplications = append(SophisticatedtComplications, "The locals are eager to export culture")
	tComplicationsTagMap["Sophisticated"] = SophisticatedtComplications

	TheocratictComplications := []string{}
	TheocratictComplications = append(TheocratictComplications, "The local theocracy is a substantial improvement on what went before")
	TheocratictComplications = append(TheocratictComplications, "The clergy secretly war with each other over obscure doctrinal differences")
	TheocratictComplications = append(TheocratictComplications, "The theocracy is decadent but the locals can imagine no other life")
	TheocratictComplications = append(TheocratictComplications, "The theocracy is in the midst of upheaval and the commoners are divided")
	tComplicationsTagMap["Theocratic"] = TheocratictComplications

	ThrivingtComplications := []string{}
	ThrivingtComplications = append(ThrivingtComplications, "The boom is actually based on a false impression and is doomed to collapse gruesomely")
	ThrivingtComplications = append(ThrivingtComplications, "A deeply unsympathetic group is prospering most from the boom")
	ThrivingtComplications = append(ThrivingtComplications, "The losers in the boom would rather have the entire planet poor")
	ThrivingtComplications = append(ThrivingtComplications, "Neighboring worlds are fearful of potential expansionism")
	tComplicationsTagMap["Thriving"] = ThrivingtComplications

	TributetComplications := []string{}
	TributetComplications = append(TributetComplications, "The tribute is demanded by implacable Mandate-era tax robots")
	TributetComplications = append(TributetComplications, "The peasants are ground under as the elite fight over who is to control the tribute")
	TributetComplications = append(TributetComplications, "Tribute production interferes with creation of a much more valuable good")
	TributetComplications = append(TributetComplications, "The far traders incidentally fit the criteria for deserving a share of the tribute if they can collect it")
	tComplicationsTagMap["Tribute"] = TributetComplications

	TyrannicaltComplications := []string{}
	TyrannicaltComplications = append(TyrannicaltComplications, "Maltech gengineering has left the locals incapable of rebellion")
	TyrannicaltComplications = append(TyrannicaltComplications, "The tyrant seems necessary to resist some grim threat")
	TyrannicaltComplications = append(TyrannicaltComplications, "The tyrant has genuine popular support")
	TyrannicaltComplications = append(TyrannicaltComplications, "Members of the former regime have formed a rebel underground")
	tComplicationsTagMap["Tyrannical"] = TyrannicaltComplications

	UsurpedtComplications := []string{}
	UsurpedtComplications = append(UsurpedtComplications, "The local rulers desperately need the usurpers for something")
	UsurpedtComplications = append(UsurpedtComplications, "The usurpers have colorably legitimate claims on the world's wealth")
	UsurpedtComplications = append(UsurpedtComplications, "The usurpers are backed by extensive local collaboration")
	UsurpedtComplications = append(UsurpedtComplications, "Only the usurpers know how to turn local resources into a rare specialty")
	tComplicationsTagMap["Usurped"] = UsurpedtComplications

	VendortComplications := []string{}
	VendortComplications = append(VendortComplications, "The goods are the product of a brutally exploitative system")
	VendortComplications = append(VendortComplications, "The goods are running out or being depleted")
	VendortComplications = append(VendortComplications, "A foreign world has seized all trading rights for their own people")
	VendortComplications = append(VendortComplications, "Vast stores of these goods are yet untapped or undiscovered and have provoked a gold rush")
	tComplicationsTagMap["Vendor"] = VendortComplications

	XenophobictComplications := []string{}
	XenophobictComplications = append(XenophobictComplications, "The locals have been misled by manipulative rulers")
	XenophobictComplications = append(XenophobictComplications, "The locals blame offworlders for a great disaster")
	XenophobictComplications = append(XenophobictComplications, "The natives have a deeply alien moral framework")
	XenophobictComplications = append(XenophobictComplications, "Only the ruling class is xenophobic")
	tComplicationsTagMap["Xenophobic"] = XenophobictComplications
	return tComplicationsTagMap
}

func createTradeRegulationsMap() map[string][]string {
	tRegulationsTagMap := make(map[string][]string)
	AlientRegulations := []string{}
	AlientRegulations = append(AlientRegulations, "Humans are restricted to trade zones")
	AlientRegulations = append(AlientRegulations, "Traders are required to give passage to members of a non-starfaring species")
	AlientRegulations = append(AlientRegulations, "Traders must prove themselves honorary members of the species")
	AlientRegulations = append(AlientRegulations, "Offworld trade is always prefaced by harsh ritual")
	tRegulationsTagMap["Alien"] = AlientRegulations

	ClosedtRegulations := []string{}
	ClosedtRegulations = append(ClosedtRegulations, "Trade is possible for 'friends of the people'")
	ClosedtRegulations = append(ClosedtRegulations, "Trade at a very few traditional times and places is allowed")
	ClosedtRegulations = append(ClosedtRegulations, "Offworld traders are universally treated as spies")
	ClosedtRegulations = append(ClosedtRegulations, "Religious experts are allowed to handle offworld goods")
	tRegulationsTagMap["Closed"] = ClosedtRegulations

	CommunisttRegulations := []string{}
	CommunisttRegulations = append(CommunisttRegulations, "Only state traders are allowed to buy and sell to offworlders")
	CommunisttRegulations = append(CommunisttRegulations, "Offworlders must bring specific state-mandated goods to be allowed any trade")
	CommunisttRegulations = append(CommunisttRegulations, "Only far traders with 'good politics' are allowed to trade")
	CommunisttRegulations = append(CommunisttRegulations, "Far traders must demonstrate their commitment to the revolution to be allowed to trade")
	tRegulationsTagMap["Communist"] = CommunisttRegulations

	DisorganizedtRegulations := []string{}
	DisorganizedtRegulations = append(DisorganizedtRegulations, "Trading with one group will incite bitter hatred from their neighbors")
	DisorganizedtRegulations = append(DisorganizedtRegulations, "Trade is eagerly sought for weapons and other military supplies")
	DisorganizedtRegulations = append(DisorganizedtRegulations, "The only serviced spaceport is constantly fought over but never itself attacked")
	DisorganizedtRegulations = append(DisorganizedtRegulations, "All deals need to involve all sides of any live conflict")
	tRegulationsTagMap["Disorganized"] = DisorganizedtRegulations

	DyingtRegulations := []string{}
	DyingtRegulations = append(DyingtRegulations, "Openly acknowledging the impending armageddon is a capital crime")
	DyingtRegulations = append(DyingtRegulations, "All travel offworld is restricted to those permitted by the elite")
	DyingtRegulations = append(DyingtRegulations, "Local currency is based on chances at a ship berth")
	DyingtRegulations = append(DyingtRegulations, "Far traders must take a certain number of passengers in order to leave")
	tRegulationsTagMap["Dying"] = DyingtRegulations

	FractioustRegulations := []string{}
	FractioustRegulations = append(FractioustRegulations, "Openly admitting the existence of the factions is a grave crime")
	FractioustRegulations = append(FractioustRegulations, "The factions offer different trading terms with their own advantages and problems")
	FractioustRegulations = append(FractioustRegulations, "Some factions have won over xenophobic locals by forbidding offworld trade")
	FractioustRegulations = append(FractioustRegulations, "The leading faction is built on control of offworld trade")
	tRegulationsTagMap["Fractious"] = FractioustRegulations

	KleptocratictRegulations := []string{}
	KleptocratictRegulations = append(KleptocratictRegulations, "Openly acknowledging the corruption is a grave crime")
	KleptocratictRegulations = append(KleptocratictRegulations, "Bribes are never paid in money")
	KleptocratictRegulations = append(KleptocratictRegulations, "Far traders are squeezed much more harshly than locals")
	KleptocratictRegulations = append(KleptocratictRegulations, "Corruption is ruthlessly punished when revealed yet remains ubiquitous")
	tRegulationsTagMap["Kleptocratic"] = KleptocratictRegulations

	LaissezFairetRegulations := []string{}
	LaissezFairetRegulations = append(LaissezFairetRegulations, "A person can sell anything here- even their own life and freedom")
	LaissezFairetRegulations = append(LaissezFairetRegulations, "Local contract-law judges can be openly influenced by 'favors'")
	LaissezFairetRegulations = append(LaissezFairetRegulations, "Contracts are all bearer documents and he who holds the physical document has the deal")
	LaissezFairetRegulations = append(LaissezFairetRegulations, "Grand fraud is a capital crime")
	tRegulationsTagMap["LaissezFaire"] = LaissezFairetRegulations

	MegacorpstRegulations := []string{}
	MegacorpstRegulations = append(MegacorpstRegulations, "No one wins a legal case without corp patronage")
	MegacorpstRegulations = append(MegacorpstRegulations, "Corp agents have a license to kill")
	MegacorpstRegulations = append(MegacorpstRegulations, "All business holdings must be corp-affiliated")
	MegacorpstRegulations = append(MegacorpstRegulations, "The corps sell local sovereignty to the highest bidder")
	tRegulationsTagMap["Megacorps"] = MegacorpstRegulations

	MilitarytRegulations := []string{}
	MilitarytRegulations = append(MilitarytRegulations, "Far traders are required to earn military ranks to trade")
	MilitarytRegulations = append(MilitarytRegulations, "Far traders are expected to support military missions")
	MilitarytRegulations = append(MilitarytRegulations, "Goods are subject to confiscation at times of military necessity")
	MilitarytRegulations = append(MilitarytRegulations, "Far traders risk conscription")
	tRegulationsTagMap["Military"] = MilitarytRegulations

	OpenedtRegulations := []string{}
	OpenedtRegulations = append(OpenedtRegulations, "Far traders don't yet have legal standing in the local courts")
	OpenedtRegulations = append(OpenedtRegulations, "A rival trader has 'helped' the locals set up the trade laws to his own benefit")
	OpenedtRegulations = append(OpenedtRegulations, "The locals have dusted off trade laws from before the Scream")
	OpenedtRegulations = append(OpenedtRegulations, "Spaceship travel is permitted only to and from designated 'starports'")
	tRegulationsTagMap["Opened"] = OpenedtRegulations

	PanopticontRegulations := []string{}
	PanopticontRegulations = append(PanopticontRegulations, "The local currency is based on privacy allotments")
	PanopticontRegulations = append(PanopticontRegulations, "Local clothing is transparent")
	PanopticontRegulations = append(PanopticontRegulations, "Concealment or deceit of any kind is a serious crime")
	PanopticontRegulations = append(PanopticontRegulations, "The locals are incredibly good at detecting smuggled goods")
	tRegulationsTagMap["Panopticon"] = PanopticontRegulations

	PrimitivetRegulations := []string{}
	PrimitivetRegulations = append(PrimitivetRegulations, "Offworld tech use is restricted to the native elite")
	PrimitivetRegulations = append(PrimitivetRegulations, "Tech use is relegated to a tightly-policed underclass of laborers")
	PrimitivetRegulations = append(PrimitivetRegulations, "Tech use is permitted only after dangerous purification rites")
	PrimitivetRegulations = append(PrimitivetRegulations, "Tech use is considered a declaration of rebellion")
	tRegulationsTagMap["Primitive"] = PrimitivetRegulations

	RestrictedtRegulations := []string{}
	RestrictedtRegulations = append(RestrictedtRegulations, "The tech isn't so important as the ritual or cultural meaning of holding it")
	RestrictedtRegulations = append(RestrictedtRegulations, "Social rank is based on tech permissions")
	RestrictedtRegulations = append(RestrictedtRegulations, "All tech is theoretically owned by the state and loaned to the worthy")
	RestrictedtRegulations = append(RestrictedtRegulations, "Non-elite cannot use the tech even to save their own lives")
	tRegulationsTagMap["Restricted"] = RestrictedtRegulations

	ScarcitytRegulations := []string{}
	ScarcitytRegulations = append(ScarcitytRegulations, "Wastage of the material is considered a hideously evil deed")
	ScarcitytRegulations = append(ScarcitytRegulations, "Private trade in the material is forbidden")
	ScarcitytRegulations = append(ScarcitytRegulations, "A special distributor caste manages all storage and disbursal of the material")
	ScarcitytRegulations = append(ScarcitytRegulations, "The more of the material you possess the more spiritually worthy you are")
	tRegulationsTagMap["Scarcity"] = ScarcitytRegulations

	SecrettRegulations := []string{}
	SecrettRegulations = append(SecrettRegulations, "Government agents track down and deal with secrecy breaches")
	SecrettRegulations = append(SecrettRegulations, "Government departments duel each other for control of the trade")
	SecrettRegulations = append(SecrettRegulations, "Government prices are decidedly stingy")
	SecrettRegulations = append(SecrettRegulations, "Far traders must bring in certain goods to trade at all")
	tRegulationsTagMap["Secret"] = SecrettRegulations

	SophisticatedtRegulations := []string{}
	SophisticatedtRegulations = append(SophisticatedtRegulations, "Outsiders must present tribute before they are allowed to trade")
	SophisticatedtRegulations = append(SophisticatedtRegulations, "Far traders affiliated with savage worlds are unwelcome")
	SophisticatedtRegulations = append(SophisticatedtRegulations, "The local customs house is unpleasantly good at preventing smuggling")
	SophisticatedtRegulations = append(SophisticatedtRegulations, "Otherwise forbidden goods can be traded here in certain areas")
	tRegulationsTagMap["Sophisticated"] = SophisticatedtRegulations

	TheocratictRegulations := []string{}
	TheocratictRegulations = append(TheocratictRegulations, "Offworlders are required to participate in potentially repugnant local rites")
	TheocratictRegulations = append(TheocratictRegulations, "A valuable commodity is forbidden by religious law")
	TheocratictRegulations = append(TheocratictRegulations, "Trade must be sanctified by permission from a querulous religious leader")
	TheocratictRegulations = append(TheocratictRegulations, "Painfully stiff tithing requirements are on all trade")
	tRegulationsTagMap["Theocratic"] = TheocratictRegulations

	ThrivingtRegulations := []string{}
	ThrivingtRegulations = append(ThrivingtRegulations, "The world's legal system is unprepared to handle the new business arrangements")
	ThrivingtRegulations = append(ThrivingtRegulations, "The old guard is struggling to maintain former sumptuary laws")
	ThrivingtRegulations = append(ThrivingtRegulations, "Threatened rulers are regulating to kill the boom before they are challenged")
	ThrivingtRegulations = append(ThrivingtRegulations, "Economic nationalists are prohibiting offworld investment or holdings")
	tRegulationsTagMap["Thriving"] = ThrivingtRegulations

	TributetRegulations := []string{}
	TributetRegulations = append(TributetRegulations, "Locals don't trust offworlders until they fill a tribute quota")
	TributetRegulations = append(TributetRegulations, "All property belongs to the elite and the tribute is just what they choose to take")
	TributetRegulations = append(TributetRegulations, "Social status among commoners is based on quota fulfillment")
	TributetRegulations = append(TributetRegulations, "Legal cases always go to the party who has contributed more tribute")
	tRegulationsTagMap["Tribute"] = TributetRegulations

	TyrannicaltRegulations := []string{}
	TyrannicaltRegulations = append(TyrannicaltRegulations, "Official 'minders' must remain with offworlders at all times")
	TyrannicaltRegulations = append(TyrannicaltRegulations, "Commoners are forbidden from speaking with offworlders")
	TyrannicaltRegulations = append(TyrannicaltRegulations, "Costly gifts of luxury goods are necessary to even begin trading")
	TyrannicaltRegulations = append(TyrannicaltRegulations, "Far traders known to deal with the regime are subject to legal repercussions on neighboring worlds")
	tRegulationsTagMap["Tyrannical"] = TyrannicaltRegulations

	UsurpedtRegulations := []string{}
	UsurpedtRegulations = append(UsurpedtRegulations, "Not even personal gear is allowed out of the starport without a tariff")
	UsurpedtRegulations = append(UsurpedtRegulations, "Far traders are forbidden from communicating with sellers")
	UsurpedtRegulations = append(UsurpedtRegulations, "Far traders are restricted to the starport")
	UsurpedtRegulations = append(UsurpedtRegulations, "Far traders must befriend the usurpers before trade is allowed ")
	tRegulationsTagMap["Usurped"] = UsurpedtRegulations

	VendortRegulations := []string{}
	VendortRegulations = append(VendortRegulations, "Local regulations to prevent smuggling are very elaborate")
	VendortRegulations = append(VendortRegulations, "Revealing or stealing the manufacturing process is a capital crime")
	VendortRegulations = append(VendortRegulations, "Sale of the good is allowed only to registered and approved buyers")
	VendortRegulations = append(VendortRegulations, "Buyers require a local official as a patron if they are to obtain the good")
	tRegulationsTagMap["Vendor"] = VendortRegulations

	XenophobictRegulations := []string{}
	XenophobictRegulations = append(XenophobictRegulations, "Offworlders are executed on discovery if the mob doesn't get them first")
	XenophobictRegulations = append(XenophobictRegulations, "Locals will buy only goods that can be disguised as local wares")
	XenophobictRegulations = append(XenophobictRegulations, "The locals view certain wares as being too good for offworlders")
	XenophobictRegulations = append(XenophobictRegulations, "All trade must take place under the pretext of tribute and submission.")
	tRegulationsTagMap["Xenophobic"] = XenophobictRegulations
	return tRegulationsTagMap
}

func describeTag(tag string) string {
	fixedTag := fixPlanetTag(tag)
	descriptionTagMap := make(map[string]string)
	AbandonedColonyDescr := "The world once hosted a colony, whether human or otherwise, until some crisis or natural disaster drove the inhabitants away or killed them off. The colony might have been mercantile in nature, an expedition to extract valuable local resources, or it might have been a reclusive cabal of zealots. The remains of the colony are usually in ruins, and might still be dangerous from the aftermath of whatever destroyed it in the first place."
	descriptionTagMap["AbandonedColony"] = AbandonedColonyDescr

	AlienRuinsDescr := "The world has significant alien ruins present. The locals may or may not permit others to investigate the ruins, and may make it difficult to remove any objects of value without substantial payment. Any surviving ruins with worthwhile salvage almost certainly have some defense or hazard to explain their unplundered state."
	descriptionTagMap["AlienRuins"] = AlienRuinsDescr

	AlteredHumanityDescr := "The humans on this world are visibly and drastically different from normal humanity. They may have additional limbs, new sensory organs, or other significant changes. Were these from ancestral eugenic manipulation, strange stellar mutations, or from an environmental toxin unique to this world?"
	descriptionTagMap["AlteredHumanity"] = AlteredHumanityDescr

	AnarchistsDescr := "Rather than being an incidental anarchy of struggling tribes and warring factions, this world actually has a functional society with no centralized authority. Authority might be hyper-localized to extended families, specific religious parishes, or voluntary associations. Some force is preventing an outside group or internal malcontents from coalescing into a power capable of imposing its rule on the locals; this force might be an ancient pretech defense system, a benevolent military AI, or the sheer obscurity and isolation of the culture."
	descriptionTagMap["Anarchists"] = AnarchistsDescr

	AnthropomorphsDescr := "The locals were originally human, but at some point became anthropomorphic, either as an ancient furry colony, a group of animal-worshiping sectarians, or gengineers who just happened to find animal elements most convenient for surviving on the world. Depending on the skill of the original gengineers, their feral forms may or may not work as well as their original human bodies, or may come with drawbacks inherited from their animal elements."
	descriptionTagMap["Anthropomorphs"] = AnthropomorphsDescr

	Area51Descr := "The world’s government is fully aware of their local stellar neighbors, but the common populace has no idea about it- and the government means to keep it that way. Trade with government officials in remote locations is possible, but any attempt to clue the commoners in on the truth will be met with lethal reprisals."
	descriptionTagMap["Area51"] = Area51Descr

	BadlandsWorldDescr := "Whatever the original climate and atmosphere type, something horrible happened to this world. Biological, chemical, or nanotechnical weaponry has reduced it to a wretched hellscape. Some local life might still be able to survive on its blasted surface, usually at some dire cost in health or humanity."
	descriptionTagMap["BadlandsWorld"] = BadlandsWorldDescr

	BattlegroundDescr := "The world is a battleground for two or more outside powers. They may be interstellar rivals, or groups operating out of orbitals or other system bodies. Something about the planet is valuable enough for them to fight over, but the natives are too weak to be anything but animate obstacles to the fight."
	descriptionTagMap["Battleground"] = BattlegroundDescr

	BeastmastersDescr := "The natives have extremely close bonds with the local fauna, possibly having special means of communication and control through tech or gengineering. Local animal life plays a major role in their society, industry, or warfare, and new kinds of beasts may be bred to suit their purposes."
	descriptionTagMap["Beastmasters"] = BeastmastersDescr

	BubbleCitiesDescr := "Whether due to a lack of atmosphere or an uninhabitable climate, the world’s cities exist within domes or pressurized buildings. In such sealed environments, techniques of surveillance and control can grow baroque and extreme."
	descriptionTagMap["BubbleCities"] = BubbleCitiesDescr

	CheapLifeDescr := "Human life is near-worthless on this world. Ubiquitous cloning, local conditions that ensure early death, a culture that reveres murder, or a social structure that utterly discounts the value of most human lives ensures that death is the likely outcome for any action that irritates someone consequential."
	descriptionTagMap["CheapLife"] = CheapLifeDescr

	CivilWarDescr := "The world is currently torn between at least two opposing factions, all of which claim legitimacy. The war may be the result of a successful rebel uprising against tyranny, or it might just be the result of schemers who plan to be the new masters once the revolution is complete."
	descriptionTagMap["CivilWar"] = CivilWarDescr

	ColdWarDescr := "Two or more great powers control the planet, and they have a hostility to each other that’s just barely less than open warfare. The hostility might be ideological in nature, or it might revolve around control of some local resource."
	descriptionTagMap["ColdWar"] = ColdWarDescr

	ColonizedPopulationDescr := "A neighboring world has successfully colonized this less-advanced or less-organized planet, and the natives aren’t happy about it. A puppet government may exist, but all real decisions are made by the local viceroy."
	descriptionTagMap["ColonizedPopulation"] = ColonizedPopulationDescr

	CulturalPowerDescr := "The world is a considerable cultural power in the sector, producing music, art, philosophy, or some similar intangible that their neighbors find irresistibly attractive. Other worlds might have a profound degree of cultural cachet as the inheritor of some venerable artistic tradition."
	descriptionTagMap["CulturalPower"] = CulturalPowerDescr

	CybercommunistsDescr := "On this world communism actually works, thanks to pretech computing devices and greater or lesser amounts of psychic precognition. Central planning nodes direct all production and employment on the world. Citizens in good standing have access to ample amounts of material goods for all needs and many wants. Instead of strife over wealth, conflicts erupt over political controls, cultural precepts, or control over the planning nodes. Many cybercommunist worlds show a considerable bias toward the private interests of those who run the planning nodes."
	descriptionTagMap["Cybercommunists"] = CybercommunistsDescr

	CyborgsDescr := "The planet’s population makes heavy use of cybernetics, with many of the inhabitants possessing at least a cosmetic amount of chrome. This may be the result of a strong local cyber tech base, a religious injunction, or simply a necessary measure to survive the local conditions."
	descriptionTagMap["Cyborgs"] = CyborgsDescr

	CyclicalDoomDescr := "The world regularly suffers some apocalyptic catastrophe that wipes out organized civilization on it. The local culture is aware of this cycle and has traditions to ensure a fragment of civilization survives into the next era, but these traditions don’t always work properly, and sometimes dangerous fragments of the past emerge."
	descriptionTagMap["CyclicalDoom"] = CyclicalDoomDescr

	DesertWorldDescr := "The world may have a breathable atmosphere and a human-tolerable temperature range, but it is an arid, stony waste outside of a few places made habitable by human effort. The deep wastes are largely unexplored and inhabited by outcasts and worse."
	descriptionTagMap["DesertWorld"] = DesertWorldDescr

	DoomedWorldDescr := "The world is doomed, and the locals may or may not know it. Some cosmic catastrophe looms before them, and the locals have no realistic way to get everyone to safety. To the extent that the public is aware, society is disintegrating into a combination of religious fervor, abject hedonism, and savage violence."
	descriptionTagMap["DoomedWorld"] = DoomedWorldDescr

	DyingRaceDescr := "The inhabitants of this world are dying out, and they know it. Through environmental toxins, hostile bio-weapons, or sheer societal despair, the culture cannot replenish its numbers. Members seek meaning in their own strange goals or peculiar faiths, though a few might struggle to find some way to reverse their slow yet inevitable doom."
	descriptionTagMap["DyingRace"] = DyingRaceDescr

	EugenicCultDescr := "Even in the days before the Silence, major improvement of the human genome always seemed to come with unacceptable side-effects. Some worlds host secret cults that perpetuate these improvements regardless of the cost, and a few planets have been taken over entirely by the cults."
	descriptionTagMap["EugenicCult"] = EugenicCultDescr

	ExchangeConsulateDescr := "The Exchange of Light once served as the largest, most trusted banking and diplomatic service in human space. Even after the Silence, some worlds retain a functioning Exchange Consulate where banking services and arbitration can be arranged."
	descriptionTagMap["ExchangeConsulate"] = ExchangeConsulateDescr

	FallenHegemonDescr := "At some point in the past, this world was a hegemonic power over some or all of the sector, thanks to superior tech, expert diplomacy, the weakness of their neighbors, or inherited Mandate legitimacy. Some kind of crash or revolt broke their power, however, and now the world is littered with the wreckage of former glory."
	descriptionTagMap["FallenHegemon"] = FallenHegemonDescr

	FeralWorldDescr := "In the long, isolated night of the Silence, some worlds have experienced total moral and cultural collapse. Whatever remains has been twisted beyond recognition into assorted death cults, xenophobic fanaticism, horrific cultural practices, or other behavior unacceptable on more enlightened worlds. These worlds are almost invariably quarantined by other planets."
	descriptionTagMap["FeralWorld"] = FeralWorldDescr

	FlyingCitiesDescr := "Perhaps the world is a gas giant, or plagued with unendurable storms at lower levels of the atmosphere. For whatever reason, the cities of this world fly above the surface of the planet. Perhaps they remain stationary, or perhaps they move from point to point in search of resources."
	descriptionTagMap["FlyingCities"] = FlyingCitiesDescr

	ForbiddenTechDescr := "Some group on this planet fabricates or uses maltech. Unbraked AIs doomed to metastasize into insanity, nation-destroying nanowarfare particles, slow-burn DNA corruptives, genetically engineered slaves, or something worse still. The planet’s larger population may or may not be aware of the danger in their midst."
	descriptionTagMap["ForbiddenTech"] = ForbiddenTechDescr

	FormerWarriorsDescr := "The locals of this world were once famed for their martial prowess. They may have simply had a very militaristic culture, or were genetically engineered for combat, or developed high-tech weaponry, or had brilliant leadership. Those days are past, however, either due to crushing defeat, external restrictions, or a cultural turn toward peace."
	descriptionTagMap["FormerWarriors"] = FormerWarriorsDescr

	FreakGeologyDescr := "The geology or geography of this world is simply freakish. Perhaps it’s composed entirely of enormous mountain ranges, or regular bands of land and sea, or the mineral structures all fragment into perfect cubes. The locals have learned to deal with it and their culture will be shaped by its requirements."
	descriptionTagMap["FreakGeology"] = FreakGeologyDescr

	FreakWeatherDescr := "The planet is plagued with some sort of bizarre or hazardous weather pattern. Perhaps city-flattening storms regularly scourge the surface, or the world’s sun never pierces its thick banks of clouds."
	descriptionTagMap["FreakWeather"] = FreakWeatherDescr

	FriendlyFoeDescr := "Some hostile alien race or malevolent cabal has a branch or sect on this world that is actually quite friendly toward outsiders. For whatever internal reason, they are willing to negotiate and deal honestly with strangers, and appear to lack the worst impulses of their fellows."
	descriptionTagMap["FriendlyFoe"] = FriendlyFoeDescr

	GoldRushDescr := "Gold, silver, and other conventional precious minerals are common and cheap now that asteroid mining is practical for most worlds. But some minerals and compounds remain precious and rare, and this world has recently been discovered to have a supply of them. People from across the sector have come to strike it rich."
	descriptionTagMap["GoldRush"] = GoldRushDescr

	GreatWorkDescr := "The locals are obsessed with completing a massive project, one that has consumed them for generations. It might be the completion of a functioning spaceyard, a massive solar power array, a network of terraforming engines, or the universal conversion of their neighbors to their own faith. The purpose of their entire civilization is to progress and some day complete the work."
	descriptionTagMap["GreatWork"] = GreatWorkDescr

	HatredDescr := "For whatever reason, this world’s populace has a burning hatred for the inhabitants of a neighboring system. Perhaps this world was colonized by exiles, or there was a recent interstellar war, or ideas of racial or religious superiority have fanned the hatred. Regardless of the cause, the locals view their neighbor and any sympathizers with loathing."
	descriptionTagMap["Hatred"] = HatredDescr

	HeavyIndustryDescr := "With interstellar transport so limited in the bulk it can move, worlds have to be largely self-sufficient in industry. Some worlds are more sufficient than others, however, and this planet has a thriving manufacturing sector capable of producing large amounts of goods appropriate to its tech level. The locals may enjoy a correspondingly higher lifestyle, or the products might be devoted towards vast projects for the aggrandizement of the rulers."
	descriptionTagMap["HeavyIndustry"] = HeavyIndustryDescr

	HeavyMiningDescr := "This world has large stocks of valuable minerals, usually necessary for local industry, life support, or refinement into loads small enough to export offworld. Major mining efforts are necessary to extract the minerals, and many natives work in the industry."
	descriptionTagMap["HeavyMining"] = HeavyMiningDescr

	HivemindDescr := "Natives of this world exist in a kind of mental gestalt, sharing thoughts and partaking of a single identity. Powerful pretech, exotic psionics, alien influence, or some other cause has left the world sharing one identity. Individual members may have greater or lesser degrees of effective coordination with the whole."
	descriptionTagMap["Hivemind"] = HivemindDescr

	HolyWarDescr := "A savage holy war is raging on this world, either between factions of locals or as a united effort against the pagans of some neighboring world. This war might involve a conventional religion, or it might be the result of a branding campaign, political ideology, artistic movement, or any other cause that people use as a substitute for traditional religion."
	descriptionTagMap["HolyWar"] = HolyWarDescr

	HostileBiosphereDescr := "The world is teeming with life, and it hates humans. Perhaps the life is xenoallergenic, forcing filter masks and tailored antiallergens for survival. It could be the native predators are huge and fearless, or the toxic flora ruthlessly outcompetes earth crops."
	descriptionTagMap["HostileBiosphere"] = HostileBiosphereDescr

	HostileSpaceDescr := "The system in which the world exists is a dangerous neighborhood. Something about the system is perilous to inhabitants, either through meteor swarms, stellar radiation, hostile aliens in the asteroid belt, or periodic comet clouds."
	descriptionTagMap["HostileSpace"] = HostileSpaceDescr

	ImmortalsDescr := "Natives of this world are effectively immortal. They may have been gengineered for tremendous lifespans, or have found some local anagathic, or be cyborg life forms, or be so totally convinced of reincarnation that death is a cultural irrelevance. Any immortality technique is likely applicable only to them, or else it’s apt to be a massive draw to outside imperialists. "
	descriptionTagMap["Immortals"] = ImmortalsDescr

	LocalSpecialtyDescr := "The world may be sophisticated or barely capable of steam engines, but either way it produces something rare and precious to the wider galaxy. It might be some pharmaceutical extract produced by a secret recipe, a remarkably popular cultural product, or even gengineered humans uniquely suited for certain work."
	descriptionTagMap["LocalSpecialty"] = LocalSpecialtyDescr

	LocalTechDescr := "The locals can create a particular example of extremely high tech, possibly even something that exceeds pretech standards. They may use unique local resources to do so, or have stumbled on a narrow scientific breakthrough, or still have a functional experimental manufactory."
	descriptionTagMap["LocalTech"] = LocalTechDescr

	MajorSpaceyardDescr := "Most worlds of tech level 4 or greater have the necessary tech and orbital facilities to build spike drives and starships. This world is blessed with a major spaceyard facility, either inherited from before the Silence or painstakingly constructed in more recent decades. It can build even capital-class hulls, and do so more quickly and cheaply than its neighbors."
	descriptionTagMap["MajorSpaceyard"] = MajorSpaceyardDescr

	MandarinateDescr := "The planet is ruled by an intellectual elite chosen via ostensibly neutral examinations or tests. The values this system selects for may or may not have anything to do with actual practical leadership skills, and the examinations may be more or less corruptible."
	descriptionTagMap["Mandarinate"] = MandarinateDescr

	MandateBaseDescr := "The Terran Mandate retained its control over this world for much longer than usual, and the world may still consider itself a true inheritor of Mandate legitimacy. Most of these worlds have or had superior technology, but they may still labor under the burden of ancient restrictive tech or monitoring systems designed to prevent them from rebelling."
	descriptionTagMap["MandateBase"] = MandateBaseDescr

	ManeatersDescr := "The locals are cannibals, either out of necessity or out of cultural preference. Some worlds may actually eat human flesh, while others simply require the rendering of humans into important chemicals or pharmaceutical compounds, perhaps to prolong the lives of ghoul overlords. This cannibalism plays a major role in their society."
	descriptionTagMap["Maneaters"] = ManeatersDescr

	MegacorpsDescr := "The world is dominated by classic cyberpunk-esque megacorporations, each one far more important than the vestigial national remnants that encompass them. These megacorps are usually locked in a cold war, trading and dealing with each other even as they try to strike in deniable ways. An over-council of corporations usually acts to bring into line any that get excessively overt in their activities."
	descriptionTagMap["Megacorps"] = MegacorpsDescr

	MercenariesDescr := "The world is either famous for its mercenary bands or it is plagued by countless groups of condottieri in service to whatever magnate can afford to pay or bribe them adequately."
	descriptionTagMap["Mercenaries"] = MercenariesDescr

	MisandryMisogynyDescr := "The culture on this world holds a particular gender in contempt. Members of that gender are not permitted positions of formal power, and may be restricted in their movements and activities. Some worlds may go so far as to scorn both traditional genders, using gengineering techniques to hybridize or alter conventional human biology."
	descriptionTagMap["MisandryMisogyny"] = MisandryMisogynyDescr

	NightWorldDescr := "The world is plunged into eternal darkness. The only life on this planet derives its energy from other sources, such as geothermal heat, extremely volatile chemical reactions in the planet’s soil, or light in a non-visible spectrum. Most flora and fauna is voraciously eager to consume other life."
	descriptionTagMap["NightWorld"] = NightWorldDescr

	MinimalContactDescr := "The locals refuse most contact with offworlders. Only a small, quarantined treaty port is provided for offworld trade, and ships can expect an exhaustive search for contraband. Local governments may be trying to keep the very existence of interstellar trade a secret from their populations, or they may simply consider offworlders too dangerous or repugnant to be allowed among the population."
	descriptionTagMap["MinimalContact"] = MinimalContactDescr

	NomadsDescr := "Most of the natives of this world are nomadic, usually following a traditional cycle of movement through the lands they possess. Promises of rich plunder or local environmental perils can force these groups to strike out against neighbors. Other groups are forced to move constantly due to unpredictable dangers that crop up on the planet. "
	descriptionTagMap["Nomads"] = NomadsDescr

	OceanicWorldDescr := "The world is entirely or almost entirely covered with liquid water. Habitations might be floating cities, or might cling precariously to the few rocky atolls jutting up from the waves, or are planted as bubbles on promontories deep beneath the stormy surface. Survival depends on aquaculture. Planets with inedible alien life rely on gengineered Terran sea crops."
	descriptionTagMap["OceanicWorld"] = OceanicWorldDescr

	OutofContactDescr := "The natives have been entirely out of contact with the greater galaxy for centuries or longer. Perhaps the original colonists were seeking to hide from the rest of the universe, or the Silence destroyed any means of communication. It may have been so long that human origins on other worlds have regressed into a topic for legends. The players might be on the first offworld ship to land since the First Wave of colonization a thousand years ago."
	descriptionTagMap["OutofContact"] = OutofContactDescr

	OutpostWorldDescr := "The world is only a tiny outpost of human habitation planted by an offworld corporation or government. Perhaps the staff is there to serve as a refueling and repair stop for passing ships, or to oversee an automated mining and refinery complex. They might be there to study ancient ruins, or simply serve as a listening and monitoring post for traffic through the system. The outpost is likely well-equipped with defenses against casual piracy."
	descriptionTagMap["OutpostWorld"] = OutpostWorldDescr

	PerimeterAgencyDescr := "Before the Silence, the Perimeter was a Terran-sponsored organization charged with rooting out use of maltech, technology banned in human space as too dangerous for use or experimentation. Unbraked AIs, gengineered slave species, nanotech replicators, weapons of planetary destruction… the Perimeter hunted down experimenters with a great indifference to planetary laws. Most Perimeter Agencies collapsed during the Silence, but a few managed to hold on to their mission, though modern Perimeter agents often find more work as conventional spies."
	descriptionTagMap["PerimeterAgency"] = PerimeterAgencyDescr

	PilgrimageSiteDescr := "The world is noted for an important spiritual or historical location, and might be the sector headquarters for a widespread religion or political movement. The site attracts wealthy pilgrims from throughout nearby space, and those with the money necessary to manage interstellar travel can be quite generous to the site and its keepers. The locals tend to be fiercely protective of the place and its reputation, and some places may forbid the entrance of those not suitably pious or devout."
	descriptionTagMap["PilgrimageSite"] = PilgrimageSiteDescr

	PleasureWorldDescr := "This world provides delights either rare or impermissible elsewhere. Matchless local beauty, stunningly gengineered natives, a wide variety of local drugs, carnal pleasures unacceptable on other worlds, or some other rare delight is readily available here. Most worlds are fully aware of the value of their offerings, and the prices they demand can be in credits or in less tangible recompense."
	descriptionTagMap["PleasureWorld"] = PleasureWorldDescr

	PoliceStateDescr := "The world is a totalitarian police state. Any sign of disloyalty to the planet’s rulers is punished severely, and suspicion riddles society. Some worlds might operate by Soviet-style informers and indoctrination, while more technically sophisticated worlds might rely on omnipresent cameras or braked AI “guardian angels”. Outworlders are apt to be treated as a necessary evil at best, and “disappeared” if they become troublesome."
	descriptionTagMap["PoliceState"] = PoliceStateDescr

	PostScarcityDescr := "The locals have maintained sufficient Mandate-era tech to be effectively post-scarcity in their economic structure. Everyone has all the necessities and most of the desires they can imagine. Conflict now exists over the apportionment of services and terrestrial space, since anything else can be had in abundance. Military goods and items of mass destruction may still be restricted, and there is probably some reason that the locals do not export their vast wealth."
	descriptionTagMap["PostScarcity"] = PostScarcityDescr

	PreceptorArchiveDescr := "The Preceptors of the Great Archive were a pre-Silence organization devoted to ensuring the dissemination of human culture, history, and basic technology to frontier worlds that risked losing this information during the human expansion. Most frontier planets had an Archive where natives could learn useful technical skills in addition to human history and art. Those Archives that managed to survive the Silence now strive to send their missionaries of knowledge to new worlds in need of their lore."
	descriptionTagMap["PreceptorArchive"] = PreceptorArchiveDescr

	PretechCultistsDescr := "The capacities of human science before the Silence vastly outmatch the technology available since the Scream. The Jump Gates alone were capable of crossing hundreds of light years in a moment, and they were just one example of the results won by blending psychic artifice with pretech science. Some worlds outright worship the artifacts of their ancestors, seeing in them the work of more enlightened and perfect humanity. These cultists may or may not understand the operation or replication of these devices, but they seek and guard them jealously."
	descriptionTagMap["PretechCultists"] = PretechCultistsDescr

	PrimitiveAliensDescr := "The world is populated by a large number of sapient aliens that have yet to develop advanced technology. The human colonists may have a friendly or hostile relationship with the aliens, but a certain intrinsic tension is likely. Small human colonies might have been enslaved or otherwise subjugated."
	descriptionTagMap["PrimitiveAliens"] = PrimitiveAliensDescr

	PrisonPlanetDescr := "This planet is or was intended as a prison. Some such prisons were meant for specific malefactors of the Terran Mandate, while others were to contain entire “dangerous” ethnic groups or alien races. Some may still have warden AIs or automatic systems to prevent any unauthorized person from leaving, and any authorization permits have long since expired."
	descriptionTagMap["PrisonPlanet"] = PrisonPlanetDescr

	PsionicsAcademyDescr := "This world is one of the few that have managed to redevelop the basics of psychic training. Without this education, a potential psychic is doomed to either madness or death unless they refrain from using their abilities. Psionic academies are rare enough that offworlders are often sent there to study by wealthy patrons. The secrets of psychic mentorship, the protocols and techniques that allow a psychic to successfully train another, are carefully guarded at these academies. Most are closely affiliated with the planetary government."
	descriptionTagMap["PsionicsAcademy"] = PsionicsAcademyDescr

	PsionicsFearDescr := "The locals are terrified of psychics. Perhaps their history is studded with feral psychics who went on murderous rampages, or perhaps they simply nurse an unreasoning terror of those “mutant freaks”. Psychics demonstrate their powers at risk of their lives."
	descriptionTagMap["PsionicsFear"] = PsionicsFearDescr

	PsionicsWorshipDescr := "These natives view psionic powers as a visible gift of god or sign of superiority. If the world has a functional psychic training academy, psychics occupy almost all major positions of power and are considered the natural and proper rulers of the world. If the world lacks training facilities, it is likely a hodgepodge of demented cults, with each one dedicated to a marginally-coherent feral prophet and their psychopathic ravings."
	descriptionTagMap["PsionicsWorship"] = PsionicsWorshipDescr

	QuarantinedWorldDescr := "The world is under a quarantine, and space travel to and from it is strictly forbidden. This may be enforced by massive ground batteries that burn any interlopers from the planet’s sky, or it may be that a neighboring world runs a persistent blockade."
	descriptionTagMap["QuarantinedWorld"] = QuarantinedWorldDescr

	RadioactiveWorldDescr := "Whether due to a legacy of atomic warfare unhindered by nuke snuffers or a simple profusion of radioactive elements, this world glows in the dark. Even heavy vacc suits can filter only so much of the radiation, and most natives suffer a wide variety of cancers, mutations and other illnesses without the protection of advanced medical treatments."
	descriptionTagMap["RadioactiveWorld"] = RadioactiveWorldDescr

	RefugeesDescr := "The world teems with refugees, either exiles from another planet who managed to get here, or the human detritus of some local conflict that have fled to the remaining stable states. The natives usually regard the refugees with hostility, an attitude returned by many among their unwilling guests."
	descriptionTagMap["Refugees"] = RefugeesDescr

	RegionalHegemonDescr := "This world has the technological sophistication, natural resources, and determined polity necessary to be a regional hegemon for the sector. Nearby worlds are likely either directly subservient to it or tack carefully to avoid its anger. It may even be the capital of a small stellar empire."
	descriptionTagMap["RegionalHegemon"] = RegionalHegemonDescr

	RestrictiveLawsDescr := "A myriad of laws, customs, and rules constrain the inhabitants of this world, and even acts that are completely permissible elsewhere are punished severely here. The locals may provide lists of these laws to offworlders, but few non-natives can hope to master all the important intricacies."
	descriptionTagMap["RestrictiveLaws"] = RestrictiveLawsDescr

	RevanchistsDescr := "The locals formerly owned another world, or a major nation on the planet formerly owned an additional region of land. Something happened to take away this control or drive out the former rulers, and they’ve never forgotten it. The locals are obsessed with reclaiming their lost lands, and will allow no questions of practicality to interfere with their cause."
	descriptionTagMap["Revanchists"] = RevanchistsDescr

	RevolutionariesDescr := "The world is convulsed by one or more bands of revolutionaries, with some nations perhaps in the grip of a current revolution. Most of these upheavals can be expected only to change the general flavor of problems in the polity, but the process of getting there usually produces a tremendous amount of suffering."
	descriptionTagMap["Revolutionaries"] = RevolutionariesDescr

	RigidCultureDescr := "The local culture is extremely rigid. Certain forms of behavior and belief are absolutely mandated, and any deviation from these principles is punished, or else society may be strongly stratified by birth with limited prospects for change. Anything which threatens the existing social order is feared and shunned."
	descriptionTagMap["RigidCulture"] = RigidCultureDescr

	RisingHegemonDescr := "This world is not yet a dominant power in the sector, but it’s well on its way there. Whether through newly-blossoming economic, military, or cultural power, they’re extending their influence over their neighbors and forging new arrangements between their government and the rulers of nearby worlds."
	descriptionTagMap["RisingHegemon"] = RisingHegemonDescr

	RitualCombatDescr := "The locals favor some form of stylized combat to resolve disputes, provide entertainment, or settle religious differences. This combat is probably not normally lethal unless it’s reserved for a specific disposable class of slaves or professionals. Some combat may involve mastery of esoteric weapons and complex arenas, while other forms might require nothing more than a declaration in the street and a drawn gun."
	descriptionTagMap["RitualCombat"] = RitualCombatDescr

	RobotsDescr := "The world has a great many robots on it. Most bots are going to be non-sentient expert systems, though an AI with enough computing resources can control many bots at once, and some worlds may have developed VIs to a degree that individual bots can seem (or be) sentient. Some worlds might even be ruled by metal overlords, ones which do not need to be sentient so long as they have overwhelming force."
	descriptionTagMap["Robots"] = RobotsDescr

	SeagoingCitiesDescr := "Either the world is entirely water or else the land is simply too dangerous for most humans. Human settlement on this world consists of a number of floating cities that follow the currents and the fish. These city-ships might have been purpose-built for their task, or they could be jury-rigged conglomerations of ships and structures thrown together when the need for seagoing life become apparent to the locals."
	descriptionTagMap["SeagoingCities"] = SeagoingCitiesDescr

	SealedMenaceDescr := "Something on this planet has the potential to create enormous havoc for the inhabitants if it is not kept safely contained by its keepers. Whether a massive seismic fault line suppressed by pretech terraforming technology, a disease that has to be quarantined within hours of discovery, or an ancient alien relic that requires regular upkeep in order to prevent planetary catastrophe, the menace is a constant shadow on the fearful populace."
	descriptionTagMap["SealedMenace"] = SealedMenaceDescr

	SecretMastersDescr := "The world is actually run by a hidden cabal, acting through their catspaws in the visible government. For one reason or another, this group finds it imperative that they not be identified by outsiders, and in some cases even the planet’s own government may not realize that they’re actually being manipulated by hidden masters."
	descriptionTagMap["SecretMasters"] = SecretMastersDescr

	SectariansDescr := "The world is torn by violent disagreement between sectarians of a particular faith. Each views the other as a damnable heresy in need of extirpation. Local government may be able to keep open war from breaking out, but the poisonous hatred divides communities. The nature of the faith may be religious, or it may be based on some secular ideology."
	descriptionTagMap["Sectarians"] = SectariansDescr

	SeismicInstabilityDescr := "The local land masses are remarkably unstable, and regular earthquakes rack the surface. Local construction is either advanced enough to sway and move with the vibrations or primitive enough that it is easily rebuilt. Severe volcanic activity may be part of the instability."
	descriptionTagMap["SeismicInstability"] = SeismicInstabilityDescr

	ShackledWorldDescr := "This world is being systematically contained by an outside power. Some ancient autonomous defense grid, robot law enforcement, alien artifact, or other force is preventing the locals from developing certain technology, or using certain devices, or perhaps from developing interstellar flight. This limit may or may not apply to offworlders; in the former case, the PCs may have to figure out a way to beat the shackles simply to escape the world."
	descriptionTagMap["ShackledWorld"] = ShackledWorldDescr

	SocietalDespairDescr := "The world’s dominant society has lost faith in itself. Whether through some all-consuming war, great catastrophe, overwhelming outside culture, or religious collapse, the natives no longer believe in their old values, and search desperately for something new. Fierce conflict often exists between the last believers in the old dispensation and the nihilistic or searching disciples of the new age."
	descriptionTagMap["SocietalDespair"] = SocietalDespairDescr

	SoleSupplierDescr := "Some extremely important resource is exported from this world and this world alone. It’s unlikely that the substance is critical for building spike drives unless this world is also the first to begin interstellar flight, but it may be critical to other high-tech processes or devices. The locals make a large amount of money off this trade and control of it is of critical importance to the planet’s rulers, and potentially to outside powers."
	descriptionTagMap["SoleSupplier"] = SoleSupplierDescr

	TabooTreasureDescr := "The natives here produce something that is both fabulously valuable and strictly forbidden elsewhere in the sector. It may be a lethally addictive drug, forbidden gengineering tech, vat-grown “perfect slaves”, or a useful substance that can only be made through excruciating human suffering. This treasure is freely traded on the world, but bringing it elsewhere is usually an invitation to a long prison stay or worse."
	descriptionTagMap["TabooTreasure"] = TabooTreasureDescr

	TerraformFailureDescr := "This world was marginal for human habitation when it was discovered, but the Mandate or the early government put in pretech terraforming engines to correct its more extreme qualities. The terraforming did not entirely work, either failing of its own or suffering the destruction of the engines during the Silence. The natives are only partly adapted to the world’s current state, and struggle with the environment."
	descriptionTagMap["TerraformFailure"] = TerraformFailureDescr

	TheocracyDescr := "The planet is ruled by the priesthood of the predominant religion or ideology. The rest of the locals may or may not be terribly pious, but the clergy have the necessary military strength, popular support or control of resources to maintain their rule. Alternative faiths or incompatible ideologies are likely to be both illegal and socially unacceptable."
	descriptionTagMap["Theocracy"] = TheocracyDescr

	TombWorldDescr := "Tomb worlds are planets that were once inhabited by humans before the Silence. The sudden collapse of the jump gate network and the inability to bring in the massive food supplies required by the planet resulted in starvation, warfare, and death. Most tomb worlds are naturally hostile to human habitation and could not raise sufficient crops to maintain life. The few hydroponic facilities were usually destroyed in the fighting, and all that is left now are ruins, bones, and silence."
	descriptionTagMap["TombWorld"] = TombWorldDescr

	TradeHubDescr := "This world is a major crossroads for local interstellar trade. It is well-positioned at the nexus of several short-drill trade routes, and has facilities for easy transfer of valuable cargoes and the fueling and repairing of starships. The natives are accustomed to outsiders, and a polyglot mass of people from every nearby world can be found trading here."
	descriptionTagMap["TradeHub"] = TradeHubDescr

	TyrannyDescr := "The local government is brutal and indifferent to the will of the people. Laws may or may not exist, but the only one that matters is the whim of the rulers on any given day. Their minions swagger through the streets while the common folk live in terror of their appetites. The only people who stay wealthy are friends and servants of the ruling class."
	descriptionTagMap["Tyranny"] = TyrannyDescr

	UnbrakedAIDescr := "Artificial intelligences are costly and difficult to create, requiring a careful sequence of “growth stages” in order to bring them to sentience before artificial limits on cognition speed and learning development are installed. These “brakes” prevent runaway cognition metastasis. This world has an “unbraked AI” on it, probably with a witting or unwitting corps of servants. Unbraked AIs are quite insane, but they learn and reason with a speed impossible for humans, and can demonstrate a truly distressing subtlety."
	descriptionTagMap["UnbrakedAI"] = UnbrakedAIDescr

	UrbanizedSurfaceDescr := "The world’s land area is covered with buildings that extend downward for multiple levels. Such worlds either have a population in the trillions, extremely little land area, or are largely-abandoned due to some past catastrophe. Agriculture and resource extraction are part of the urban complex, and there may be an advanced maintenance system that may not be entirely under the control of present natives."
	descriptionTagMap["UrbanizedSurface"] = UrbanizedSurfaceDescr

	UtopiaDescr := "Natural and social conditions on this world have made it a paradise for its inhabitants, a genuine utopia of happiness and fulfillment. This is normally the result of drastic human engineering, including brain-gelding, neurochemical control, personality curbs, or complete “humanity” redefinitions. Even so, the natives are extremely happy with their lot, and may wish to extend that joy to poor, sad outsiders."
	descriptionTagMap["Utopia"] = UtopiaDescr

	WarlordsDescr := "The world is plagued by warlords. Numerous powerful men and women control private armies sufficiently strong to cow whatever local government may exist. On the lands they claim, their word is law. Most spend their time oppressing their own subjects and murderously pillaging those of their neighbors. Most like to wrap themselves in the mantle of ideology, religious fervor, or an ostensibly legitimate right to rule."
	descriptionTagMap["Warlords"] = WarlordsDescr

	XenophilesDescr := "The natives of this world are fast friends with a particular alien race. The aliens may have saved the planet at some point in the past, or awed the locals with superior tech or impressive cultural qualities. The aliens might even be the ruling class on the planet."
	descriptionTagMap["Xenophiles"] = XenophilesDescr

	XenophobesDescr := "The natives are intensely averse to dealings with outworlders. Whether through cultural revulsion, fear of tech contamination, or a genuine immunodeficiency, the locals shun foreigners from offworld and refuse to have anything to do with them beyond the bare necessities of contact. Trade may or may not exist on this world, but if it does, it is almost certainly conducted by a caste of untouchables and outcasts."
	descriptionTagMap["Xenophobes"] = XenophobesDescr

	ZombiesDescr := "This menace may not take the form of shambling corpses, but some disease, alien artifact, or crazed local practice produces men and women with habits similar to those of murderous cannibal undead. These outbreaks may be regular elements in local society, either provoked by some malevolent creators or the consequence of some local condition."
	descriptionTagMap["Zombies"] = ZombiesDescr

	if _, ok := descriptionTagMap[fixedTag]; !ok {
		descriptionTagMap[fixedTag] = "ERROR: description of '" + fixedTag + "' cannot be made"
	}

	return descriptionTagMap[fixedTag]

}

func describeTradeTag(tag string) string {
	fixedTag := fixPlanetTag(tag)
	tDescrTagMap := make(map[string]string)
	AlientDescr := "The denizens of this world are either aliens or else possess such a strange culture that they might as well be non-humans. Such societies tend to have strange estimates of value by ordinary human standards, often prizing substances of limited interstellar value or requiring bizarre procedures before business can be conducted. Dealing with alien societies is fraught with danger even by far trader standards, as it is easy for a human to trespass taboos that seem so obvious to their local dragoman that there is no thought to warn them beforehand."
	tDescrTagMap["Alien"] = AlientDescr

	ClosedtDescr := "The planet flatly refuses to conduct any trade with offworlders, or may be prevented from trading by some organization or power, or may simply be so strange that no coherent concept of 'trade' is possible. Primitive worlds have few ways to physically enforce this ban on landing, but their government or social structure has enough moral sway to prevent most members from breaking the prohibition. More advanced worlds may have orbital defenses to ensure that interloping far traders do not come where they are not wanted. Some such worlds permit landing and communication, but no significant business can be conducted- or at least, no significant legal business."
	tDescrTagMap["Closed"] = ClosedtDescr

	CommunisttDescr := "The world's economic system is communist in nature. All goods and all production are under the control of the state, and private trade tends to be minimal where it is not outright forbidden. On planets with very small populations, this communism tends to be of a crude variety expressed through familial obligations and a sense of shared labor for communities small enough to police their own members. Unlike ancient Terra, however, most worlds now lack the hyperintelligent AIs and precognitive economistpsychics that once made command economies practical on a planetary scale. As a consequence, most surviving large-scale communist worlds are gray places of calcified bureaucracy, populist rancor, or fresh-faced enthusiasm yet untempered by experience."
	tDescrTagMap["Communist"] = CommunisttDescr

	DisorganizedtDescr := "The world is a chaotic mass of tribes, statelets, warring classes, hostile ethnicities, atomized philosophical adherents, or some other churning mess of inchoate humanity. There is no real centralized power despite whatever proclamations may be made by the government of the hour, and no infrastructure for handling sophisticated trade. The people of different regions, castes, or cultures might be drastically different in their attitudes toward far traders, and the rules can change lethally in a matter of kilometers or minutes. While a far trader is free of the coordinated rapaciousness of a planetary government, he must fight off innumerable smaller exactions, and has no better avenue of justice than what he can get from the local ruler or his own strong arm."
	tDescrTagMap["Disorganized"] = DisorganizedtDescr

	DyingtDescr := "The world is in the process of collapse. Unlike the conventional anarchy of a savage world, basic human survival on this planet is becoming gradually more and more difficult, either from a hostile environment that the locals can no longer fend off, or from a cultural collapse so complete that basic survival processes can no longer be maintained. In some cases, astrophysical events such as impending solar flares or unstoppable rogue-planet impacts might spell the doom of an otherwise healthy world. The time pressure is usually not so intense as to cause complete social breakdown, but almost everyone on the planet realizes that their time is coming to a close. Trade still exists on such worlds only insofar as the locals are not convinced they can successfully steal a trader's ship, or if they think they have enough time to build their own with the help of outside supplies. In almost no case is it feasible to save anything but a tiny fraction of the world's population, and their leaders are usually determined to be the ones rescued. In a few cases, the death spiral might be reversible with enough outside tech and support- if the locals can be kept from mad panic."
	tDescrTagMap["Dying"] = DyingtDescr

	FractioustDescr := "Where a disorganized world is an incoherent mess, a fractious world is divided up into several strong, organized groups. These factions may have de facto existence in the form of rival nations or geographic regions, or they may be political artifacts of a disjointed planetary government, with different groups of officials jousting over control of offices and delegated authority. Most such worlds have not devolved into overt warfare, but tensions can range from a decided coolness toward rivals to assassination and deniable guerrilla strikes. "
	tDescrTagMap["Fractious"] = FractioustDescr

	KleptocratictDescr := "The planet has a dominant trade framework, but that framework is hopelessly corrupt. Nothing happens without money, and possession of wealth is license enough for every variety of crime and wrongdoing. The local officials are ferociously greedy, peeling every spare credit off a far trader and making very narrow estimates of how much profit the offworlder will need in order to justify return. The endless corruption makes for poverty and suffering among the common folk, but those who control the levers of wealth do very well."
	tDescrTagMap["Kleptocratic"] = KleptocratictDescr

	LaissezFairetDescr := "The planet has few if any trade laws, but retains enough social coherence and infrastructure to facilitate large-scale trade. Some worlds of this type recognize almost no trade laws except basic rules about fraud and contract enforcement and permit virtually anything to be bought and sold in their markets. Most laissez-faire worlds either have so few inhabitants that tariffs are unnecessary, or else have a deep cultural commitment to free trade, or are so populous and divided in power sources that no single faction or government can maintain restrictive trade practices. Minor far traders love these worlds, though greater trade princes tend to prefer captive markets."
	tDescrTagMap["LaissezFaire"] = LaissezFairetDescr

	MegacorpstDescr := "All significant economic activity above the individual freelancer is controlled by a relatively handful of giant corporations, either directly or through a web of subsidiary companies obedient to the megacorp's will. Some of these corporations might be private in origin and others might be state enterprises, but their size and importance has left them integral partners with whatever government might exist on this world, to the extent that many of them effectively are the governments of their spheres of influence. Their rivalry rarely extends to outright warfare, but deniable strikes and industrial espionage are a way of life on this world. On those worlds where public government has not entirely withered away the remaining states retain their independence by the strength of their military forces and their willingness to handle less profitable work for a stiff fee from the megacorps. So long as the money keeps coming and their military primacy is not threatened, such crony-capitalist leaders have little interest in what the megacorps do with 'customers'."
	tDescrTagMap["Megacorps"] = MegacorpstDescr

	MilitarytDescr := "The world is heavily militarized. For a rare few worlds this militarization may revolve around the space navy, with the citizens organized to support a fleet vastly in excess of the usual limits. For most military planets, however, the focus is on ground and atmospheric forces and the rivalries that have led to their growth. A world in this state might have huge conscript armies, completely militarized societies, or otherwise be dominated by the effort required to build and maintain huge numbers of people in martial professions. Such worlds may have limited wars breaking out in client nations or among allied vassal states. Others may be locked in generations-long trench warfare."
	tDescrTagMap["Military"] = MilitarytDescr

	OpenedtDescr := "An opened world is fresh to the game of interstellar contact, one willing to reestablish the ancient bonds of trade that once webbed worlds together. Regrettably, most of them are not entirely certain how the whole business is supposed to work. The natives have little experience in dealing with offworlders, often misinterpret their customs, and have certain assumptions and expectations that they never think to question... along with an often-overblown esteem for 'offworld tech'. Opened worlds usually have not had time to develop sophisticated starports, complex interstellar trade bureaucracies, or a jaded indifference to far traders."
	tDescrTagMap["Opened"] = OpenedtDescr

	PanopticontDescr := "The world is watched. Everything of consequence is monitored, every interaction recorded and logged. Panopticon worlds usually form only on planets sufficiently hostile to force bubble cities or similar contained, controlled habitats. In such compressed circumstances it becomes possible to track individuals to an extent that would be infeasible on more free-ranging worlds. Some panopticon worlds are relatively benign, but most exist in order to enable some ruthlessly oppressive power structure. In an effective panopticon society the only people with unpunished crimes are those responsible for monitoring the observations. The societies that grow from such soil tend to pair spectacularly corrupt overseers with a populace trained to a half-demented degree of dissembling."
	tDescrTagMap["Panopticon"] = PanopticontDescr

	PrimitivetDescr := "The world is technologically primitive. Their societies may be quite sophisticated in cultural, artistic, or aesthetic terms, but they lack the raw materials or expertise to move past the muscle-powered stages of technology. Few such worlds can support centralized governments or large numbers of non-farming specialists, and social organization tends to be oriented toward stability and survival. Far traders who can bring in tech to unlock formerly-inaccessible resource deposits or jumpstart an industrial base can unleash drastic changes on these worlds. The natives often lack the experience to recognize these impending changes before their world has transformed around them."
	tDescrTagMap["Primitive"] = PrimitivetDescr

	RestrictedtDescr := "The locals of this world labor under a regime that tightly restricts the employment of common technology. While most worlds have strictures against devices of mass destruction or private biological weaponry development, this world goes so far as to control the use of entire classes of technology- weapons, communications tech, medical science, or other society-changing artifacts. Use of this restricted tech is usually reserved to the local elites or priestly class, and its possession often comes with powerful cultural significance. Enforcement of these restrictions is usually at the hands of the tech's proper owners, though some worlds are locked into their laws by outside powers."
	tDescrTagMap["Restricted"] = RestrictedtDescr

	ScarcitytDescr := "The planet is miserably poor in some vital necessity for civilization. Food, water, breathable atmosphere, workable metals, or some other staple of human society is available in very limited amounts. Control of this commodity rests in the hands of the planet's elite, and most conflicts revolve around its apportionment. While most far traders lack the cargo space to bring in truly world-changing amounts of the material, tech that breaks the scarcity through superior production or extraction techniques can result in convulsive change to the society. Even so, many far traders prefer to bring in more limited amounts, the better to string out the local profits as long as possible."
	tDescrTagMap["Scarcity"] = ScarcitytDescr

	SecrettDescr := "Any trade with this world must be conducted secretly with those few local powers aware of the far traders' existence. The rest of the planet may be virulently xenophobic, or overseen by offworld imperialists, or the private economic fief of a different far trader empire. The locals may not even realize that interstellar travel is possible. Those permitted to conduct offworld trade are usually government representatives or secret societies aware of the truth. The construction of significant holdings on such a world can be difficult, with everything conducted through a complex series of local cutouts and native 'geniuses' recruited to explain new technology. The governments of such worlds vigorously discourage any attempt to widen a far trader's market."
	tDescrTagMap["Secret"] = SecrettDescr

	SophisticatedtDescr := "The world may or may not be a thriving hub of interstellar trade, but it is certainly comfortable with the idea and at ease with the concept of a wider cosmos. Sophisticated worlds may have large space navies or a long tradition of astronautic adventuring, but they show none of the usual awe toward far traders that more backward worlds might demonstrate. They know what the local sector has to offer and are quite canny about what they are willing to buy and sell. They may be more technologically advanced than most other worlds in the area, or have enough social contact with their neighbors to have formed a polyglot, cosmopolitan sort of local culture. If they have little trade contact with other worlds, it's likely because they're convinced that the barbarians have nothing to offer them."
	tDescrTagMap["Sophisticated"] = SophisticatedtDescr

	TheocratictDescr := "The planet is under the control of a theocratic government, either as a centralized religious bureaucracy or one dominated by numerous smaller sects. In the chaotic aftermath of the Scream the mutual ties and unquenchable hope provided by many organized faiths often left their clergy the most respected group on the planet. Over time, this respect easily transformed into outright rule. Those faiths that were best represented among colony worlds tended to be those that coped well with issues of technology and rationalism, but the Scream also encouraged a strong Luddite strain. The faith that is a cynical and tired bureaucracy on one world can be a creed of fire and bloody triumph on its neighbor, and a far trader sometimes cannot discern the difference until after they've been surrounded."
	tDescrTagMap["Theocratic"] = TheocratictDescr

	ThrivingtDescr := "The world is in the throes of an enormous economic boom. Whether through the discovery of new resources, the integration of new tech into their industrial base, or a renaissance of ideas and creativity, life on this world is getting better for almost everyone. If the world is not already a major player in sector affairs, it is likely to become so in the foreseeable future. While the good times are widely spread among the population, it is inevitable that someone is going to be left out- and former oligarchs and ruling elites tend to find such universal prosperity to be disruptive to the pillars of their former rule."
	tDescrTagMap["Thriving"] = ThrivingtDescr

	TributetDescr := "The world's economy revolves around providing tribute. It may be to a powerful local elite, or a colonizing offworld power, or a bizarre local custom of destruction. Whatever the specifics, society is designed to create large surpluses of necessary goods for the benefit of others. In the rare case that the tribute-collectors are offworlders, the tribute itself usually is in the form of some small, precious commodity that can be carried in limited interstellar cargo space. Local elites can be more widely rapacious, and whole nations can be set to slaving to produce the fantastic opulence enjoyed by a blessed few. Such tribute arrangements may have physical force to encourage the workers, but some degree of cultural acceptance is necessary for such a small minority to effectively commandeer the labor of so many. As such, the elites tend to be very sensitive about offworlders who question affairs."
	tDescrTagMap["Tribute"] = TributetDescr

	TyrannicaltDescr := "The planet is under the control of an all-encompassing tyranny. Multiple dictatorships might be locked in perpetual struggle, or the entire world might be gripped in a single fist. The only law is the will of the ruler and the prerogatives of their minions, and any deal lasts only so long as the tyrant finds it advantageous to keep it. Personal freedoms are nonexistent, and long-standing tyrannies might have so beaten the resistance out of their people that their only real threat to the ruler lies in a palace coup. Such worlds are eager to buy new tools of restraint and luxuries from distant worlds, though they can be risky trading partners. Their rulers are not accustomed to being balked."
	tDescrTagMap["Tyrannical"] = TyrannicaltDescr

	UsurpedtDescr := "The interstellar trade conducted on this world has been completely suborned by some outside power. The local government or society does not control the terms of the trade or benefit much, if at all, from the fruits of it. Most often this is the product of a weak world and a powerful far trader combine seizing control of the spaceports through a mixture of bribery and threats. Sometimes a foreign world is also able to force all trade to pass through its 'customs houses', often placed in orbit where less technologically-sophisticated locals can't overwhelm them with numbers. A few alien races might do the same thing- or humans might do such to an alien world unable to throw off this unwanted 'administrative assistance'. The locals receive only a pittance for their products and the far traders are gouged on the prices in turn, with the surplus collected by the usurping power."
	tDescrTagMap["Usurped"] = UsurpedtDescr

	VendortDescr := "There is a valuable thing this planet produces that can be found nowhere else, and the bulk of their offworld trade revolves around selling this thing and gathering the tariffs and taxes that come from it. It may be a rare plant, or a precious mineral, or the product of some still-functional pretech nanofactory, or the fruit of native ingenuity, or even the remnants of an advanced alien race. Whatever its specifics, this is the only planet in the area that can offer it, and they extract every possible credit from its trade. Control of this resource is very important to the locals, as the far traders will naturally cooperate with those capable of cutting off their supply."
	tDescrTagMap["Vendor"] = VendortDescr

	XenophobictDescr := "The locals hate offworlders. This hate may be born of some venerable crime against the planet, or it might just be the product of cultural loathing. Most planets naturally assume that their culture is morally or aesthetically superior to those of other worlds, but these people actively hate their neighbors. Offworlders are depraved, vicious, filthy beings who likely host a variety of social diseases and ought to be shunned or worse if they dare to show their faces on the clean soil of this world. Unsurprisingly, this makes trade a rather difficult business, but some among the governments and major corporations of this world are willing to make grim sacrifices to necessity. Of course, any such deals must be hidden from the populace or else mob violence is almost certain to ensue. Few traders have any love of landing on these worlds, but their very isolationism tends to make them fabulously profitable for those few ships that actually do arrive."
	tDescrTagMap["Xenophobic"] = XenophobictDescr
	return tDescrTagMap[fixedTag]

}
