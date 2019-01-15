package main

func Story(index int, t1 string, t2 string) string {
	allStories := []string{
		"An " + Enemy(t1, t2) + " seeks to rob a " + Friend(t1, t2) + " of some precious " + Thing(t1, t2) + " that he has desired for some time.",
		"A " + Thing(t1, t2) + " has been discovered on property owned by a " + Friend(t1, t2) + ", but a " + Complication(t1, t2) + " risks its destruction.",
		"A " + Complication(t1, t2) + " suddenly hits the party while they’re out doing some innocuous activity.",
		"The players unwittingly offend or injure an " + Enemy(t1, t2) + ", incurring his or her wrath. A " + Friend(t1, t2) + " offers help in escaping the consequences.",
		"Rumor speaks of the discovery of a precious " + Thing(t1, t2) + " in a distant " + Place(t1, t2) + ". The players must get to it before an " + Enemy(t1, t2) + " does.",
		"An " + Enemy(t1, t2) + " has connections with offworld pirates or slavers, and a " + Friend(t1, t2) + " has been captured by them.",
		"A " + Place(t1, t2) + " has been seized by violent revolutionaries or rebels, and a " + Friend(t1, t2) + " is being held hostage by them.",
		"A " + Friend(t1, t2) + " is in love with someone forbidden by social convention, and the two of them need help eloping.",
		"An " + Enemy(t1, t2) + " wields tyrannical power over a " + Friend(t1, t2) + ", relying on the bribery of corrupt local officials to escape consequences.",
		"A " + Friend(t1, t2) + " has been lost in hostile wilderness, and the party must reach a " + Place(t1, t2) + " to rescue them in the teeth of a dangerous " + Complication(t1, t2) + ".",
		"An " + Enemy(t1, t2) + " has committed a grave offense against a PC or their family sometime in the past. A " + Friend(t1, t2) + " shows the party a weakness in the " + Enemy(t1, t2) + "’s defenses.",
		"The party is suddenly caught in a conflict between two warring families or political parties.",
		"The party is framed for a crime by an " + Enemy(t1, t2) + ", and must reach the sanctuary of a " + Place(t1, t2) + " before they can regroup and find the " + Thing(t1, t2) + " that will prove their innocence and their " + Enemy(t1, t2) + "’s perfidy.",
		"A " + Friend(t1, t2) + " is threatened by a tragedy of sickness, legal calamity, or public humiliation, and the only one that seems able to save them is an " + Enemy(t1, t2) + ".",
		"A natural disaster or similar " + Complication(t1, t2) + " strikes a " + Place(t1, t2) + " while the party is present, causing great loss of life and property unless the party is able to immediately respond to the injured and trapped.",
		"A " + Friend(t1, t2) + " with a young business has struck a cache of pretech, valuable minerals, or precious salvage. He needs the party to help him reach the " + Place(t1, t2) + " where the valuables are.",
		"An oppressed segment of society starts a sudden revolt in the " + Place(t1, t2) + " the party is occupying. An " + Enemy(t1, t2) + " simply lumps the party in with the rebels and tries to put the revolt down with force. A " + Friend(t1, t2) + " offers them a way to either help the rebels or clear their names.",
		"A vulnerable " + Friend(t1, t2) + " has been targeted for abduction, and has need of guards. A sudden " + Complication(t1, t2) + " makes guarding them from the " + Enemy(t1, t2) + " seeking their kidnapping much more difficult. If the " + Friend(t1, t2) + " is snatched, they must rescue them from a " + Place(t1, t2) + ".",
		"A mysterious " + Place(t1, t2) + " offers the promise of some precious " + Thing(t1, t2) + ", but access is very dangerous due to wildlife, hostile locals, or a dangerous environment.",
		"An " + Enemy(t1, t2) + " and a " + Friend(t1, t2) + " both have legal claim on a " + Thing(t1, t2) + ", and seek to undermine each other’s case. The " + Enemy(t1, t2) + " is willing to do murder if he thinks he can get away with it.",
		"An " + Enemy(t1, t2) + " seeks the death of his brother, a " + Friend(t1, t2) + ", by arranging the failure of his grav flyer or shuttlecraft in dangerous terrain while the party is coincidentally aboard. The party must survive the environment and bring proof of the crime out with them.",
		"A " + Friend(t1, t2) + " seeks to slip word to a lover, one who is also being courted by the " + Friend(t1, t2) + "’s brother, who is an " + Enemy(t1, t2) + ". A " + Complication(t1, t2) + " threatens to cause death or disgrace to the lover unless they either accept the " + Enemy(t1, t2) + "’s suit or are helped by the party.",
		"An " + Enemy(t1, t2) + " is convinced that one of the party has committed adultery with their flirtatious spouse. He means to lure them to a " + Place(t1, t2) + ", trap them, and have them killed by the dangers there.",
		"An " + Enemy(t1, t2) + " has been driven insane by exotic recreational drugs or excessive psionic torching. He fixes on a PC as being his mortal nemesis, and plots elaborate deaths, attempting to conceal his involvement amid " + Complication(t1, t2) + "s.",
		"A " + Friend(t1, t2) + " has stolen a precious " + Thing(t1, t2) + " from an " + Enemy(t1, t2) + " and fled into a dangerous, inaccessible " + Place(t1, t2) + ". The party must rescue them, and decide what to do with the " + Thing(t1, t2) + " and the outraged " + Enemy(t1, t2) + ".",
		"An " + Enemy(t1, t2) + " has realized that their brother or sister has engaged in a socially unacceptable affair with a " + Friend(t1, t2) + ", and means to kill both of them unless stopped by the party.",
		"A " + Friend(t1, t2) + " has accidentally caused the death of a family member, and wants the party to help him hide the body or fake an accidental death before his family realizes what has happened. A " + Complication(t1, t2) + " suddenly makes the task more difficult.",
		"A " + Friend(t1, t2) + " is a follower of a zealous ideologue who plans to make a violent demonstration of the righteousness of his cause, causing a social " + Complication(t1, t2) + ". The " + Friend(t1, t2) + " will surely be killed in the aftermath if not rescued or protected by the party.",
		"A " + Friend(t1, t2) + "’s sibling is to be " + Place(t1, t2) + "d in a dangerous situation they’ve got no chance of surviving. The " + Friend(t1, t2) + " takes their " + Place(t1, t2) + " at the last moment, and will almost certainly die unless the party aids them.",
		"Suicide bombers detonate an explosive, chemical, or biological weapon in a " + Place(t1, t2) + " occupied by the party where a precious " + Thing(t1, t2) + " is stored The PCs must escape before the " + Place(t1, t2) + " collapses on top of them, navigating throngs of terrified people in the process and saving the " + Thing(t1, t2) + " if possible.",
		"An " + Enemy(t1, t2) + " who controls landing permits, oxygen rations, or some other important resource has a prejudice against one or more of the party members. He demands that they bring him a " + Thing(t1, t2) + " from a dangerous " + Place(t1, t2) + " before he’ll give them the goods.",
		"A " + Friend(t1, t2) + " in a loveless marriage to an " + Enemy(t1, t2) + " seeks escape to be with their beloved, and contacts the party to snatch them from their spouse’s guards at a prearranged " + Place(t1, t2) + ".",
		"A " + Friend(t1, t2) + " seeks to elope with their lover, and contacts the party to help them meet their paramour at a remote, dangerous " + Place(t1, t2) + ". On arrival, they find that the lover is secretly an " + Enemy(t1, t2) + " desirous of their removal and merely lured them to the " + Place(t1, t2) + " to meet their doom.",
		"The party receives or finds a " + Thing(t1, t2) + " which proves the crimes of an " + Enemy(t1, t2) + "- yet a " + Friend(t1, t2) + " was complicit in the crimes, and will be punished as well if the authorities are involved. And the " + Enemy(t1, t2) + " will stop at nothing to get the item back.",
		"A " + Friend(t1, t2) + " needs to get to a " + Place(t1, t2) + " on time in order to complete a business contract, but an " + Enemy(t1, t2) + " means to delay and hinder them until it’s too late, inducing " + Complication(t1, t2) + "s to the trip.",
		"A locked pretech stasis pod has been discovered by a " + Friend(t1, t2) + ", along with directions to the hidden keycode that will open it. The " + Place(t1, t2) + " where the keycode is hidden is now owned by an " + Enemy(t1, t2) + ".",
		"A fierce schism has broken out in the local majority religion, and an " + Enemy(t1, t2) + " is making a play to take control of the local hierarchy. A " + Friend(t1, t2) + " is on the side that will lose badly if the " + Enemy(t1, t2) + " succeeds, and needs a " + Thing(t1, t2) + " to prove the other group’s error.",
		"A former " + Enemy(t1, t2) + " has been given reason to repent his treatment of a " + Friend(t1, t2) + ", and secretly commissions them to help the " + Friend(t1, t2) + " overcome a " + Complication(t1, t2) + ". A different " + Enemy(t1, t2) + " discovers the connection, and tries to paint the PCs as double agents.",
		"An alien or a human with extremely peculiar spiritual beliefs seeks to visit a " + Place(t1, t2) + " for their own reasons. An " + Enemy(t1, t2) + " of their own kind attempts to stop them before they can reach the " + Place(t1, t2) + ", and reveal the " + Thing(t1, t2) + " that was hidden there long ago.",
		"A " + Friend(t1, t2) + "’s sibling is an untrained psychic, and has been secretly using his or her powers to protect the " + Friend(t1, t2) + " from an " + Enemy(t1, t2) + ". The neural damage has finally overwhelmed their sanity, and they’ve now kidnapped the " + Friend(t1, t2) + " to keep them safe. The " + Enemy(t1, t2) + " is taking this opportunity to make sure the " + Friend(t1, t2) + " “dies at the hands of their maddened sibling”.",
		"A " + Friend(t1, t2) + " who is a skilled precognitive has just received a flash of an impending atrocity to be committed by an " + Enemy(t1, t2) + ". He or she needs the party to help them steal the " + Thing(t1, t2) + " that will prove the " + Enemy(t1, t2) + "’s plans while dodging the assassins sent to eliminate the precog.",
		"A " + Friend(t1, t2) + " who is an exotic dancer is sought by an " + Enemy(t1, t2) + " who won’t take no for an answer. The dancer is secretly a Perimeter agent attempting to infiltrate a " + Place(t1, t2) + " to destroy maltech research, and plots to use the party to help get him or her into the facility under the pretext of striking at the " + Enemy(t1, t2) + ".",
		"A young woman on an interplanetary tour needs the hire of local bodyguards. She turns out to be a trained and powerful combat psychic, but touchingly naive about local dangers, causing a social " + Complication(t1, t2) + " that threatens to get the whole group arrested.",
		"A librarian " + Friend(t1, t2) + " has discovered an antique databank with the coordinates of a long-lost pretech cache hidden in a " + Place(t1, t2) + " sacred to a long-vanished religion. The librarian is totally unsuited for danger, but necessary to decipher the obscure religious iconography needed to unlock the cache. The cache is not the anticipated " + Thing(t1, t2) + ", but something more dangerous to the finder.",
		"A fragment of orbital debris clips a shuttle on the way in, and the spaceport is seriously damaged in the crash. The player’s ship or the only vessel capable of getting them off-planet will be destroyed unless the players can organize a response to the dangerous chemical fires and radioactives contaminating the port. A " + Friend(t1, t2) + " is trapped somewhere in the control tower wreckage.",
		"A " + Friend(t1, t2) + " is allied with a reformist religious group that seeks to break the grip of the current, oppressive hierarchy. The current hierarchs have a great deal of political support with the authorities, but the commoners resent them bitterly. The " + Friend(t1, t2) + "  seeks to secure a remote " + Place(t1, t2) + " as a meeting-place for the theological rebels.",
		"A microscopic black hole punctures an orbital station or starship above the world. Its interaction with the station’s artificial grav generators has thrown everything out of whack, and the station’s become a minefield of dangerously high or zero grav zones. It’s tearing itself apart, and it’s going to collapse soon. An " + Enemy(t1, t2) + " seeks to escape aboard the last lifeboat and to Hell with everyone else. Meanwhile, a " + Friend(t1, t2) + " is trying to save his engineer daughter from the radioactive, grav-unstable engine rooms.",
		"The planet has a sealed alien ruin, and an " + Enemy(t1, t2) + "-led cult who worships the vanished builders. They’re convinced that they have the secret to opening and controlling the power inside the ruins, but they’re only half-right. A " + Friend(t1, t2) + " has found evidence that shows that they’ll only devastate the planet if they meddle with the alien power planet. The party has to get inside the ruins and shut down the engines before it’s too late. Little do they realize that a few aliens survive inside, in a stasis field that will be broken by the ruin’s opening.",
		"An " + Enemy(t1, t2) + " and the group are suddenly trapped in a " + Place(t1, t2) + " during an accident or " + Complication(t1, t2) + ". They must work together to escape in time.",
		"A telepathic " + Friend(t1, t2) + " has discovered that an " + Enemy(t1, t2) + " was responsible for a recent atrocity. Telepathic evidence is useless on this world, however, and if she’s discovered to have scanned his mind she’ll be lobotomized as a ’rogue psychic’. A " + Thing(t1, t2) + " might be enough to prove his guilt, if the party can figure out how to get to it without revealing their " + Friend(t1, t2) + "’s meddling.",
		"A " + Friend(t1, t2) + " is responsible for safeguarding a " + Thing(t1, t2) + "yet the " + Thing(t1, t2) + " is suddenly proven to be a fake. The party must find the real object and the " + Enemy(t1, t2) + " who stole it or else their " + Friend(t1, t2) + " will be punished as the thief.",
		"A " + Friend(t1, t2) + " is bitten by a poisonous local animal while in a remote " + Place(t1, t2) + ". The only antidote is back at civilization, yet a " + Complication(t1, t2) + " threatens to delay the group until it is too late.",
		"A lethal plague has started among the residents of the town, but a " + Complication(t1, t2) + " is keeping aid from reaching them. An " + Enemy(t1, t2) + " is taking advantage of the panic to hawk a fake cure at ruinous prices, and a " + Friend(t1, t2) + " is taken in by him. The " + Complication(t1, t2) + " must be overcome before help can reach the town.",
		"A radical political party has started to institute pogroms against “groups hostile to the people”. A " + Friend(t1, t2) + " is among those groups, and needs to get out of town before an " + Enemy(t1, t2) + " uses the riot as cover to settle old scores.",
		"An " + Enemy(t1, t2) + " has sold the party an expensive but worthlessly flawed piece of equipment before lighting out for the back country. He and his plunder are holed up at a remote " + Place(t1, t2) + ".",
		"A concert of offworld music is being held in town, and a " + Friend(t1, t2) + " is slated to be the star performer. Reactionary elements led by an " + Enemy(t1, t2) + " plot to ruin the “corrupting noise” with sabotage that risks getting performers killed. Meanwhile, a crowd of ignorant offworlder fans have landed and are infuriating the locals.",
		"An " + Enemy(t1, t2) + " is wanted on a neighboring world for some heinous act, and a " + Friend(t1, t2) + " turns up as a bounty hunter ready to bring him in alive. This world refuses to extradite him, so the capture and retrieval has to evade local law enforcement.",
		"An unanticipated solar storm blocks communications and grounds the poorly-shielded grav vehicle that brought the group to this remote " + Place(t1, t2) + ". Then people start turning up dead; the storm has awoken a dangerous " + Enemy(t1, t2) + " beast.",
		"A " + Friend(t1, t2) + " has discovered a partially-complete schematic for an ancient pretech refinery unit that produces vast amounts of something precious on this world- water, oxygen, edible compounds, or the like. Several remote " + Place(t1, t2) + "s on the planet are indicated as having the necessary pretech spare parts required to build the device. When finally assembled, embedded self-modification software in the " + Thing(t1, t2) + " modifies itself into a pretech combat bot. The salvage from it remains very valuable.",
		"A " + Complication(t1, t2) + " ensnares the party where they are in an annoying but seemingly ordinary event. In actuality, an " + Enemy(t1, t2) + " is using it as cover to strike at a " + Friend(t1, t2) + " or " + Thing(t1, t2) + " that happens to be where the PCs are.",
		"A " + Friend(t1, t2) + " has a cranky, temperamental artificial heart installed, and the doctor who put it in is the only one who really understands how it works. The heart has recently started to stutter, but the doctor has vanished. An " + Enemy(t1, t2) + " has snatched him to fit his elite assassins with very unsafe combat mods.",
		"A local clinic is doing wonders in providing free health care to the poor. In truth, it’s a front for an offworld eugenics cult, with useful “specimens” kidnapped and shipped offworld while ’cremated remains’ are given to the family. A " + Friend(t1, t2) + " is snatched by them, but the party knows they’d have never consented to cremation as the clinic staff claim.",
		"Space pirates have cut a deal with an isolated backwoods settlement, off loading their plunder to merchants who meet them there. A " + Friend(t1, t2) + " goes home to family after a long absence, but is kidnapped or killed before they can bring back word of the dealings. Meanwhile, the party is entrusted with a valuable " + Thing(t1, t2) + " that must be brought to the " + Friend(t1, t2) + " quickly",
		"A reclusive psychiatrist is offering treatment for violent mentally ill patients at a remote " + Place(t1, t2) + ". His treatments seem to work, calming the subjects and returning them to rationality, though major memory loss is involved and some severe social clumsiness ensues. In actuality, he’s removed large portions of their brains to fit them with remote-control units slaved to an AI in his laboratory. He intends to use them as drones to acquire more “subjects”, and eventual control of the town.",
		"Vital medical supplies against an impending plague have been shipped in from offworld, but the spike drive craft that was due to deliver them misjumped, and has arrived in-system as a lifeless wreck transmitting a blind distress signal. Whoever gets there first can hold the whole planet hostage, and an " + Enemy(t1, t2) + " means to do just that.",
		"A " + Friend(t1, t2) + " has spent a substantial portion of their wealth on an ultra-modern new domicile, and invites the party to enjoy a weekend there. An " + Enemy(t1, t2) + " has hacked the house’s computer system to trap the inhabitants inside and use the automated fittings to kill them.",
		"A mud slide, hurricane, earthquake, or other form of disaster strikes a remote settlement. The party is the closest group of responders, and must rescue the locals while dealing with the unearthed, malfunctioning pretech " + Thing(t1, t2) + " that threatens to cause an even greater calamity if not safely defused.",
		"A " + Friend(t1, t2) + " has found a lost pretech installation, and needs help to loot it. By planetary law, the contents belong to the government.",
		"An " + Enemy(t1, t2) + " mistakes the party for the kind of offworlders who will murder innocents for pay- assuming they aren’t that kind, at least. He’s sloppy with the contact and unwittingly identifies himself, letting the players know that a " + Friend(t1, t2) + " will shortly die unless the " + Enemy(t1, t2) + " is stopped.",
		"A party member is identified as a prophesied savior for an oppressed faith or ethnicity. The believers obstinately refuse to believe any protestations to the contrary, and a cynical " + Enemy(t1, t2) + " in government decides the PC must die simply to prevent the risk of uprising. An equally cynical " + Friend(t1, t2) + " is determined to push the PC forward as a savior, because that’s what’s needed.",
		"Alien beasts escape from a zoo and run wild through the spectators. The panicked owner offers large rewards for recapturing them live, but some of the beasts are quite dangerous.",
		"A trained psychic is accused of going feral by an " + Enemy(t1, t2) + ". The psychic had already suffered severe neural damage before being found for training, so brain scans cannot establish physical signs of madness. The psychic seems unstable, but not violent- at least, on short acquaintance. The psychic offers a psychic PC the secrets of a unique psychic technique if they help him flee.",
		"A " + Thing(t1, t2) + " is the token of rulership on this world, and it’s gone missing. If it’s not found rapidly, the existing ruler will be deposed. Evidence left at a " + Place(t1, t2) + " suggests that an " + Enemy(t1, t2) + " has it, but extralegal means are necessary to investigate fully.",
		"Psychics are vanishing, including a " + Friend(t1, t2) + ". They’re being kidnapped by an ostensibly-rogue government researcher who is using them to research the lost psychic disciplines that helped enable pretech manufacturing, and they are being held at a remote " + Place(t1, t2) + ". The snatcher is a small-time local " + Enemy(t1, t2) + " with unnaturally ample resources.",
		"A " + Friend(t1, t2) + " desperately seeks to hide evidence of some past crime that will ruin his life should it come to light. An " + Enemy(t1, t2) + " holds the " + Thing(t1, t2) + " that proves his involvement, and blackmails him ruthlessly.",
		"A courier mistakes the party for the wrong set of offworlders, and wordlessly deposits a " + Thing(t1, t2) + " with them that implies something awful- med-frozen, child-sized human organs, for example, or a private catalog of gengineered human slaves. The courier’s boss shortly realizes the error, and this " + Enemy(t1, t2) + " tries to silence the PCs while preserving the " + Place(t1, t2) + " where his evil is enacted.",
		"A slowboat system freighter is taken over by " + Enemy(t1, t2) + " separatist terrorists at the same time as the planet’s space defenses are taken offline by internal terrorist attacks. The freighter is aimed straight at the starport, and will crash into it in hours if not stopped.",
		"Alien artifacts on the planet’s surface start beaming signals into the system’s asteroid belt. The signals provoke a social " + Complication(t1, t2) + " in panicked response, and an " + Enemy(t1, t2) + " seeks to use the confusion to take over. The actual effect of the signals might be harmless, or might summon a long-lost alien AI warship to scourge life from the world.",
		"An alien ambassador " + Friend(t1, t2) + " is targeted by xenophobe " + Enemy(t1, t2) + " assassins. Relations are so fragile that if the ambassador even realizes that humans are making a serious effort to kill him, the result may be war.",
		"A new religion is being preached by a " + Friend(t1, t2) + " on this planet. Existing faiths are not amused, and an " + Enemy(t1, t2) + " among the hierarchy is provoking the people to persecute the new believers, hoping for " + Thing(t1, t2) + "s to get out of hand.",
		"An " + Enemy(t1, t2) + " was once the patron of a " + Friend(t1, t2) + " until the latter was betrayed. Now the " + Friend(t1, t2) + " wants revenge, and they think they have the information necessary to get past the " + Enemy(t1, t2) + "’s defenses.",
		"Vital life support or medical equipment has been sabotaged by offworlders or zealots, and must be repaired before time runs out. The only possible source of parts is at a " + Place(t1, t2) + ", and the saboteurs can be expected to be working hard to get there and destroy them, too.",
		"A " + Friend(t1, t2) + " is importing offworld tech that threatens to completely replace the offerings of an " + Enemy(t1, t2) + " businessman. The " + Enemy(t1, t2) + " seeks to sabotage the " + Friend(t1, t2) + "’s stock, and thus ’prove’ its inferiority.",
		"An Exchange diplomat is negotiating for the opening of a branch of the interstellar bank on this world. An " + Enemy(t1, t2) + " among the local banks wants to show the world as being ungovernably unstable, so provokes " + Complication(t1, t2) + "s and riots around the diplomat.",
		"An " + Enemy(t1, t2) + " is infuriated by the uppity presumptio of an ambitious " + Friend(t1, t2) + " of a lower social caste, and tries to pin a local " + Complication(t1, t2) + " on the results of his unnatural rejection of his proper " + Place(t1, t2) + ".",
		"A " + Friend(t1, t2) + " is working for an offworld corporation to open a manufactory, and is ignoring local traditions that privilege certain social or ethnic groups, giving jobs to the most qualified workers instead. An angry " + Enemy(t1, t2) + " seeks to sabotage the factory.",
		"An offworld musician who was revered as little less than a god on his homeworld requires bodyguards. He immediately acquires Enemies on this world with his riotous ways, and his guards must keep him from getting arrested if they are to be paid.",
		"Atmospheric disturbances, dust storms, or other particulate clouds suddenly blow into town, leaving the settlement blind. An " + Enemy(t1, t2) + " commits a murder during the darkness, and attempts to frame the players as convenient scapegoats.",
		"An " + Enemy(t1, t2) + " spikes the oxygen supply of an orbital station or unbreathable-atmosphere hab dome with hallucinogens as cover for a theft. Most victims are merely confused and disoriented, but some become violent in their delusions. By chance, the party’s air supply was not contaminated.",
		"By coincidence, one of the party members is wearing clothing indicative of membership in a violent political group, and thus the party is treated in " + Friend(t1, t2) + "ly fashion by a local " + Enemy(t1, t2) + " for no obvious reason. The " + Enemy(t1, t2) + " assumes that the party will go along with some vicious crime without complaint, and the group isn’t informed of what’s in the offing until they’re in deep.",
		"A local ruler wishes outworlders to advise him of the quality of his execrable poetry- and is the sort to react very poorly to anything less that evidently sincere and fulsome praise. Failure to amuse the ruler results in the party being dumped in a dangerous " + Place(t1, t2) + " to “experience truly poetic solitude”.",
		"A " + Friend(t1, t2) + " among the locals is unreasonably convinced that offworlder tech can repair anything, and has blithely promised a powerful local " + Enemy(t1, t2) + " that the party can easily fix a damaged pretech " + Thing(t1, t2) + ". The " + Enemy(t1, t2) + " has invested in many expensive spare parts, but the truly necessary pieces are kept in a still-dangerous pretech installation in a remote " + Place(t1, t2) + ".",
		"The party’s offworld comm gear picks up a chance transmission from the local government and automatically descrambles the primitive encryption key. The document is proof that an " + Enemy(t1, t2) + " in government intends to commit an atrocity against a local village with a group of “deniable” renegades in order to steal a " + Thing(t1, t2) + " kept in the village.",
		"A " + Friend(t1, t2) + " belongs to a persecuted faith, ethnicity, or social class, and appeals for the PCs to help a cell of rebels get offworld before the " + Enemy(t1, t2) + " law enforcement finds them.",
		"A part on the party’s ship or the only other transport out has failed, and needs immediate replacement. The only available part is held by an " + Enemy(t1, t2) + ", who will only willingly relinquish it in exchange for a " + Thing(t1, t2) + " held by an innocent " + Friend(t1, t2) + " who will refuse to sell at any price.",
		"Eugenics cultists are making gengineered slaves out of genetic material gathered at a local brothel. Some of the unnaturally tempting slaves are being slipped among the prostitutes as bait to infatuate powerful officials, while others are being sold under the table to less scrupulous elites.",
		"Evidence has been unearthed at a " + Place(t1, t2) + " that substantial portions of the planet are actually owned by members of an oppressed and persecuted group. The local courts have no intention of recognizing the rights, but the codes with the ownership evidence would allow someone to bypass a number of antique pretech defenses around the planetary governor’s palace. A " + Friend(t1, t2) + " wants the codes to pass to his " + Friend(t1, t2) + "s among the group’s rebels.",
		"A crop smut threatens the planet’s agriculture, promising large-scale famine. A " + Friend(t1, t2) + " finds evidence that a secret government research station in the system’s asteroid belt was conducting experiments in disease-resistant crop strains for the planet before the Silence struck and cut off communication with the station. The existing government considers it a wild goose chase, but the party might choose to help. The station has stasis-frozen samples of the crop sufficient to avert the famine, but it also has less pleasant relics….",
		"A grasping " + Enemy(t1, t2) + " in local government seizes the party’s ship for some trifling offense. The " + Enemy(t1, t2) + " wants to end offworld trade, and is trying to scare other traders away. The starship is held within a military cordon, and the " + Enemy(t1, t2) + " is confident that by the time other elements of the government countermand the order, the free traders will have been spooked off.",
		"A seemingly useless trinket purchased by a PC turns out to be the security key to a lost pretech facility. It was sold by accident by a bungling and now-dead minion of a local " + Enemy(t1, t2) + ", who is hot after the party to “reclaim” his property… preferably after the party defeats whatever automatic defenses and bots the facility might still support.",
	}
	return allStories[index]
}
