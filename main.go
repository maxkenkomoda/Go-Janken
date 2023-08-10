package main

import (
	"fmt"
	"math/rand"
	"time"
)

var hs = []string{"rock", "scissors", "paper"}

func main() {
	game()
}

func game() {
	rand.Seed(time.Now().UnixNano())
	fh := hs[rand.Intn(3)]
	sh := hs[rand.Intn(3)]

	fmt.Println("First player:", fh, ",", "Second player:", sh)

	if fh == sh {
		fmt.Println("Draw")
		fmt.Println("Play again")
		game()
	}

	if fh == "rock" && sh == "scissors" {
		fmt.Println("First player win, game set")
	}

	if fh == "rock" && sh == "paper" {
		fmt.Println("Second player win, game set")
	}
}
