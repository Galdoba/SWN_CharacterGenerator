package main

import (
	"strconv"
)

const (
	STR = iota
	DEX
	CON
	INT
	WIS
	CHA
)

type NPC struct {
	Atributes  [6]int
	Background string
	Class      string
	HP         int
	AtkBon     int
	Skills     map[string]int
	Foci       []string
}

func NewNPC() *NPC {
	npc := NPC{}
	npc.Skills = make(map[string]int)
	npc.presetClassAndStyle(roll1dX(8, 0))
	npc.Background = randomBackground()
	npc.defaultSkills()
	npc.defaultFoci()

	return &npc
}

func (npc *NPC) Strength() int {
	return npc.Atributes[0]
}

func (npc *NPC) Dexterity() int {
	return npc.Atributes[1]
}

func (npc *NPC) Constitution() int {
	return npc.Atributes[2]
}

func (npc *NPC) Intelligence() int {
	return npc.Atributes[3]
}

func (npc *NPC) Wisdom() int {
	return npc.Atributes[4]
}

func (npc *NPC) Charisma() int {
	return npc.Atributes[5]
}

func (npc *NPC) SetAttribute(i int, newScore int) {
	npc.Atributes[i] = newScore
}

func (npc *NPC) String() string {
	str := ""
	str = str + "Class: " + npc.Class + "\n"
	str = str + "Background: " + npc.Background + "\n"
	str = str + "STR = " + int2Str(npc.Strength()) + " (" + strconv.Itoa(ModifierOf(npc.Strength())) + ")\n"
	str = str + "DEX = " + strconv.Itoa(npc.Atributes[1]) + " (" + strconv.Itoa(ModifierOf(npc.Atributes[1])) + ")\n"
	str = str + "CON = " + strconv.Itoa(npc.Atributes[2]) + " (" + strconv.Itoa(ModifierOf(npc.Atributes[2])) + ")\n"
	str = str + "INT = " + strconv.Itoa(npc.Atributes[3]) + " (" + strconv.Itoa(ModifierOf(npc.Atributes[3])) + ")\n"
	str = str + "WIS = " + strconv.Itoa(npc.Atributes[4]) + " (" + strconv.Itoa(ModifierOf(npc.Atributes[4])) + ")\n"
	str = str + "CHA = " + strconv.Itoa(npc.Atributes[5]) + " (" + strconv.Itoa(ModifierOf(npc.Atributes[5])) + ")\n"
	for key, val := range npc.Skills {
		str = str + "Skill " + key + "-" + strconv.Itoa(val) + "\n"
	}
	for i := range npc.Foci {
		str = str + "Aply effects from Focus: " + npc.Foci[i] + "\n"
	}
	str = str + "Add bonus Skill: " + randomSkill() + "\n"
	str = str + "Character have " + strconv.Itoa(rollXdY(2, 6)*100) + " credits \n"

	return str
}

func ModifierOf(score int) int {
	if score < 4 {
		return -2
	}
	if score < 8 {
		return -1
	}
	if score < 14 {
		return 0
	}
	if score < 18 {
		return 1
	}
	return 2
}

func pickBackground(i int) string {
	back := "Custom"
	switch i {
	case 1:
		back = "Barbarian"
	case 2:
		back = "Clergy"
	case 3:
		back = "Courtesan"
	case 4:
		back = "Criminal"
	case 5:
		back = "Dilettante"
	case 6:
		back = "Entertainer"
	case 7:
		back = "Merchant"
	case 8:
		back = "Noble"
	case 9:
		back = "Official"
	case 10:
		back = "Peasant"
	case 11:
		back = "Physician"
	case 12:
		back = "Pilot"
	case 13:
		back = "Politician"
	case 14:
		back = "Scholar"
	case 15:
		back = "Soldier"
	case 16:
		back = "Spacer"
	case 17:
		back = "Technician"
	case 18:
		back = "Thug"
	case 19:
		back = "Vagabond"
	case 20:
		back = "Worker"

	}
	return back
}

func randomBackground() string {
	i := roll1dX(20, 0)
	return pickBackground(i)
}

func pickSkill(i int) string {
	class := ""
	switch i {
	case 1:
		class = "Administer"
	case 2:
		class = "Connect"
	case 3:
		class = "Exert"
	case 4:
		class = "Fix"
	case 5:
		class = "Heal"
	case 6:
		class = "Know"
	case 7:
		class = "Lead"
	case 8:
		class = "Notice"
	case 9:
		class = "Perform"
	case 10:
		class = "Pilot"
	case 11:
		class = "Program"
	case 12:
		class = "Punch"
	case 13:
		class = "Shoot"
	case 14:
		class = "Sneak"
	case 15:
		class = "Stab"
	case 16:
		class = "Survive"
	case 17:
		class = "Talk"
	case 18:
		class = "Trade"
	case 19:
		class = "Work"
	case 20:
		class = "Change a level-0 skill you have to level-1"

	}
	return class
}

func randomSkill() string {
	i := roll1dX(20, 0)
	return pickSkill(i)
}

func (npc *NPC) SetHP(newHp int) {
	if newHp < 1 {
		newHp = 1
	}
	npc.HP = newHp
}

func (npc *NPC) SetAtkBon(newAB int) {
	npc.AtkBon = newAB
}

func (npc *NPC) presetClassAndStyle(i int) {
	switch i {
	case 1:
		npc.Class = "Expert (Smart)"
		npc.SetAttribute(0, 10)
		npc.SetAttribute(1, 12)
		npc.SetAttribute(2, 11)
		npc.SetAttribute(3, 14)
		npc.SetAttribute(4, 7)
		npc.SetAttribute(5, 9)
		npc.SetHP(roll1dX(6, 0))
		npc.SetAtkBon(0)
	case 2:
		npc.Class = "Expert (Smooth)"
		npc.SetAttribute(0, 7)
		npc.SetAttribute(1, 9)
		npc.SetAttribute(2, 10)
		npc.SetAttribute(3, 12)
		npc.SetAttribute(4, 11)
		npc.SetAttribute(5, 14)
		npc.SetHP(roll1dX(6, 0))
		npc.SetAtkBon(0)
	case 3:
		npc.Class = "Expert (Nimble)"
		npc.SetAttribute(0, 10)
		npc.SetAttribute(1, 14)
		npc.SetAttribute(2, 12)
		npc.SetAttribute(3, 11)
		npc.SetAttribute(4, 9)
		npc.SetAttribute(5, 7)
		npc.SetHP(roll1dX(6, 0))
		npc.SetAtkBon(0)
	case 4:
		npc.Class = "Warrior (Melee)"
		npc.SetAttribute(0, 14)
		npc.SetAttribute(1, 9)
		npc.SetAttribute(2, 12)
		npc.SetAttribute(3, 7)
		npc.SetAttribute(4, 10)
		npc.SetAttribute(5, 11)
		npc.SetHP(roll1dX(6, 2))
		npc.SetAtkBon(1)
	case 5:
		npc.Class = "Warrior (Ranged)"
		npc.SetAttribute(0, 9)
		npc.SetAttribute(1, 14)
		npc.SetAttribute(2, 12)
		npc.SetAttribute(3, 10)
		npc.SetAttribute(4, 7)
		npc.SetAttribute(5, 11)
		npc.SetHP(roll1dX(6, 2))
		npc.SetAtkBon(1)
	case 6:
		npc.Class = "Warrior (Leader)"
		npc.SetAttribute(0, 7)
		npc.SetAttribute(1, 10)
		npc.SetAttribute(2, 9)
		npc.SetAttribute(3, 11)
		npc.SetAttribute(4, 12)
		npc.SetAttribute(5, 14)
		npc.SetHP(roll1dX(6, 2))
		npc.SetAtkBon(1)
	case 7:
		npc.Class = "Psychic (Seer)"
		npc.SetAttribute(0, 9)
		npc.SetAttribute(1, 11)
		npc.SetAttribute(2, 12)
		npc.SetAttribute(3, 10)
		npc.SetAttribute(4, 14)
		npc.SetAttribute(5, 7)
		npc.SetHP(roll1dX(6, 0))
		npc.SetAtkBon(0)
		npc.pickPsyhicSkill(roll1dX(6, 0))
		npc.pickPsyhicSkill(roll1dX(6, 0))
	case 8:
		npc.Class = "Psychic (Adept)"
		npc.SetAttribute(0, 12)
		npc.SetAttribute(1, 10)
		npc.SetAttribute(2, 14)
		npc.SetAttribute(3, 9)
		npc.SetAttribute(4, 11)
		npc.SetAttribute(5, 7)
		npc.SetHP(roll1dX(6, 1))
		npc.SetAtkBon(0)
		npc.pickPsyhicSkill(roll1dX(6, 0))
		npc.pickPsyhicSkill(roll1dX(6, 0))
	}
}

func (npc *NPC) defaultSkills() {
	switch npc.Background {
	case "Barbarian":
		npc.Skills["Survive"] = 0
		npc.Skills["Notice"] = 0
		npc.Skills["Any Combat"] = 0

	case "Clergy":
		npc.Skills["Talk"] = 0
		npc.Skills["Perform"] = 0
		npc.Skills["Know"] = 0

	case "Courtesan":
		npc.Skills["Perform"] = 0
		npc.Skills["Notice"] = 0
		npc.Skills["Connect"] = 0
	case "Criminal":
		npc.Skills["Connect"] = 0
		npc.Skills["Sneak"] = 0
		npc.Skills["Talk"] = 0
	case "Dilettante":
		npc.Skills["Connect"] = 0
		npc.Skills["Know"] = 0
		npc.Skills["Talk"] = 0
	case "Entertainer":
		npc.Skills["Perform"] = 0
		npc.Skills["Talk"] = 0
		npc.Skills["Connect"] = 0
	case "Merchant":
		npc.Skills["Trade"] = 0
		npc.Skills["Talk"] = 0
		npc.Skills["Connect"] = 0
	case "Noble":
		npc.Skills["Lead"] = 0
		npc.Skills["Connect"] = 0
		npc.Skills["Administer"] = 0
	case "Official":
		npc.Skills["Administer"] = 0
		npc.Skills["Talk"] = 0
		npc.Skills["Connect"] = 0
	case "Peasant":
		npc.Skills["Exert"] = 0
		npc.Skills["Sneak"] = 0
		npc.Skills["Survive"] = 0
	case "Physician":
		npc.Skills["Heal"] = 0
		npc.Skills["Know"] = 0
		npc.Skills["Notice"] = 0
	case "Pilot":
		npc.Skills["Pilot"] = 0
		npc.Skills["Fix"] = 0
		npc.Skills["Shoot or Trade"] = 0
	case "Politician":
		npc.Skills["Talk"] = 0
		npc.Skills["Lead"] = 0
		npc.Skills["Connect"] = 0
	case "Scholar":
		npc.Skills["Know"] = 0
		npc.Skills["Administer"] = 0
		npc.Skills["Connect"] = 0
	case "Soldier":
		npc.Skills["Any Combat"] = 0
		npc.Skills["Exert"] = 0
		npc.Skills["Survive"] = 0
	case "Spacer":
		npc.Skills["Fix"] = 0
		npc.Skills["Pilot"] = 0
		npc.Skills["Program"] = 0
	case "Technician":
		npc.increaseSkill("Fix")
		npc.Skills["Notice"] = 0
		npc.Skills["Exert"] = 0
	case "Thug":
		npc.Skills["Any Combat"] = 0
		npc.Skills["Talk"] = 0
		npc.Skills["Connect"] = 0
	case "Vagabond":
		npc.Skills["Notice"] = 0
		npc.Skills["Sneak"] = 0
		npc.Skills["Survive"] = 0
	case "Worker":
		npc.Skills["Connect"] = 0
		npc.Skills["Exert"] = 0
		npc.Skills["Work"] = 0
	}

}

func (npc *NPC) increaseSkill(skill string) {
	if val, ok := npc.Skills[skill]; ok {
		if val > 3 {
			return
		}
		npc.Skills[skill] = val + 1
	} else {
		npc.Skills[skill] = 0
	}
}

func (npc *NPC) pickPsyhicSkill(i int) {
	switch i {
	case 1:
		npc.increaseSkill("Biopsionics")
	case 2:
		npc.increaseSkill("Metapsionics")
	case 3:
		npc.increaseSkill("Precognition")
	case 4:
		npc.increaseSkill("Telekinesis")
	case 5:
		npc.increaseSkill("Telepathy")
	case 6:
		npc.increaseSkill("Teleportation")
	}
}

func (npc *NPC) defaultFoci() {
	if npc.Class == "Expert (Smart)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Specialist/Fix")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 2:
			npc.Foci = append(npc.Foci, "Hacker")
			npc.Foci = append(npc.Foci, "Tinker")
		case 3:
			npc.Foci = append(npc.Foci, "Specialist/Know")
			npc.Foci = append(npc.Foci, "Healer")
		case 4:
			npc.Foci = append(npc.Foci, "Specialist/Fix")
			npc.Foci = append(npc.Foci, "Tinker")
		case 5:
			npc.Foci = append(npc.Foci, "Healer")
			npc.Foci = append(npc.Foci, "Ironhide")
		case 6:
			npc.Foci = append(npc.Foci, "Specialist/Fix")
			npc.Foci = append(npc.Foci, "Hacker")
		}
	}
	if npc.Class == "Expert (Smooth)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Diplomat")
			npc.Foci = append(npc.Foci, "Connected")
		case 2:
			npc.Foci = append(npc.Foci, "Specialist/Talk")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 3:
			npc.Foci = append(npc.Foci, "Diplomat")
			npc.Foci = append(npc.Foci, "Alert")
		case 4:
			npc.Foci = append(npc.Foci, "Specialist/Lead")
			npc.Foci = append(npc.Foci, "Authority")
		case 5:
			npc.Foci = append(npc.Foci, "Healer")
			npc.Foci = append(npc.Foci, "Specialist/Talk")
		case 6:
			npc.Foci = append(npc.Foci, "Specialist/Notice")
			npc.Foci = append(npc.Foci, "Specialist/Talk")
		}
	}
	if npc.Class == "Expert (Nimble)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Specialist/Pilot")
			npc.Foci = append(npc.Foci, "Starfarer")
		case 2:
			npc.Foci = append(npc.Foci, "Healer")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 3:
			npc.Foci = append(npc.Foci, "Tinker")
			npc.Foci = append(npc.Foci, "Gunslinger")
		case 4:
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
			npc.Foci = append(npc.Foci, "Assassin")
		case 5:
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
			npc.Foci = append(npc.Foci, "Specialist/Exert")
		case 6:
			npc.Foci = append(npc.Foci, "Specialist/Entertain")
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
		}
	}
	if npc.Class == "Warrior (Melee)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Savage Fray")
			npc.Foci = append(npc.Foci, "Shocking Assault")
		case 2:
			npc.Foci = append(npc.Foci, "Assassin")
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
		case 3:
			npc.Foci = append(npc.Foci, "Armsman")
			npc.Foci = append(npc.Foci, "Close Combatant")
		case 4:
			npc.Foci = append(npc.Foci, "Close Combatant")
			npc.Foci = append(npc.Foci, "Savage Fray")
		case 5:
			npc.Foci = append(npc.Foci, "Ironhide")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 6:
			npc.Foci = append(npc.Foci, "Unarmed Combatant")
			npc.Foci = append(npc.Foci, "Close Combatant")
		}
	}
	if npc.Class == "Warrior (Ranged)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Gunslinger")
			npc.Foci = append(npc.Foci, "Close Combatant")
		case 2:
			npc.Foci = append(npc.Foci, "Sniper")
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
		case 3:
			npc.Foci = append(npc.Foci, "Assassin")
			npc.Foci = append(npc.Foci, "Specialist/Sneak")
		case 4:
			npc.Foci = append(npc.Foci, "Ironhide")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 5:
			npc.Foci = append(npc.Foci, "Gunslinger")
			npc.Foci = append(npc.Foci, "Tinker")
		case 6:
			npc.Foci = append(npc.Foci, "Close Combatant")
			npc.Foci = append(npc.Foci, "Alert")
		}
	}
	if npc.Class == "Warrior (Leader)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Gunslinger")
			npc.Foci = append(npc.Foci, "Authority")
		case 2:
			npc.Foci = append(npc.Foci, "Ironhide")
			npc.Foci = append(npc.Foci, "Connected")
		case 3:
			npc.Foci = append(npc.Foci, "Armsman")
			npc.Foci = append(npc.Foci, "Specialist/Lead")
		case 4:
			npc.Foci = append(npc.Foci, "Gunslinger")
			npc.Foci = append(npc.Foci, "Specialist/Talk")
		case 5:
			npc.Foci = append(npc.Foci, "Assassin")
			npc.Foci = append(npc.Foci, "Die Hard")
		case 6:
			npc.Foci = append(npc.Foci, "Close Combatant")
			npc.Foci = append(npc.Foci, "Henchkeeper")
		}
	}
	if npc.Class == "Psychic (Seer)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Alert")
		case 2:
			npc.Foci = append(npc.Foci, "Healer")
		case 3:
			npc.Foci = append(npc.Foci, "Specialist/Notice")
		case 4:
			npc.Foci = append(npc.Foci, "Psychic Training")
		case 5:
			npc.Foci = append(npc.Foci, "Savage Fray")
		case 6:
			npc.Foci = append(npc.Foci, "Hacker")
		}
	}
	if npc.Class == "Psychic (Adept)" {
		i := roll1dX(6, 0)
		switch i {
		case 1:
			npc.Foci = append(npc.Foci, "Armsman")
		case 2:
			npc.Foci = append(npc.Foci, "Ironhide")
		case 3:
			npc.Foci = append(npc.Foci, "Die Hard")
		case 4:
			npc.Foci = append(npc.Foci, "Psychic Training")
		case 5:
			npc.Foci = append(npc.Foci, "Healer")
		case 6:
			npc.Foci = append(npc.Foci, "Unarmed Combatant")
		}
	}
}

/*
fw
fw
fw
fe
fe
fe
fp
fp
fp
wp
we
ep


Diplomat
Connected
Specialist/Talk
Die Hard
Diplomat
Alert
Specialist/Lead
Authority
Healer
Specialist/Talk
Specialist/Notice
Specialist/Talk
*/
