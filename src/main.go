package main

import (
	"Go-Janken/src/player"
	"fmt"
	"math/rand"
	"time"
)

var hs = []string{"rock", "scissors", "paper"}

func main() {
	player.Main()
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

	if fh == "rock" && sh == "scissors" || fh == "scissors" && sh == "paper" || fh == "paper" && sh == "rock" {
		fmt.Println("First player win, game set")
	}

	if fh == "rock" && sh == "paper" || fh == "scissors" && sh == "rock" || fh == "paper" && sh == "scissors" {
		fmt.Println("Second player win, game set")
	}
}
