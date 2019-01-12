package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	loop := true
	for loop {
		seed := time.Now().UnixNano()
		rand.Seed(seed)
		fmt.Println("Seed =", seed)
		fmt.Println("")
		tectCase := NewNPC()

		fmt.Println(tectCase.String())
		loop = askYesNo("Try Another ?")
	}
}
