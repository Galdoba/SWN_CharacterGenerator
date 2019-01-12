package main

import (
	"fmt"
	"strconv"
)

type NPC struct {
	Atributes [6]int
}

func NewNPC() *NPC {
	npc := NPC{}
	//1 - rollAttributes
	fmt.Println("Rolling Attributes:")

	for i := 0; i < 6; i++ {
		npc.Atributes[i] = rollXdY(3, 6)
	}
	fmt.Println(npc.String())
	opt, _ := TakeOptions("Switch Attribute Score to 54654654564?", "Str", "Dex", "Con", "Int", "Wis", "Cha", "-No need-")
	if opt != 7 {
		npc.Atributes[opt-1] = 14
		fmt.Println(npc.String())
	}

	return &npc
}

func (npc *NPC) String() string {
	str := ""
	str = str + "STR = " + strconv.Itoa(npc.Atributes[0]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[0])) + ")\n"
	str = str + "DEX = " + strconv.Itoa(npc.Atributes[1]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[1])) + ")\n"
	str = str + "CON = " + strconv.Itoa(npc.Atributes[2]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[2])) + ")\n"
	str = str + "INT = " + strconv.Itoa(npc.Atributes[3]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[3])) + ")\n"
	str = str + "WIS = " + strconv.Itoa(npc.Atributes[4]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[4])) + ")\n"
	str = str + "CHA = " + strconv.Itoa(npc.Atributes[5]) + "(" + strconv.Itoa(ModifierOf(npc.Atributes[5])) + ")\n"

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
