package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("ping")
	tectCase := NewNPC()

	fmt.Println(tectCase.String())
}
