package main

import (
	"Go-Janken/src/player"
	"fmt"
	"math/rand"
	"time"
)

var hs = []player.HandType{"rock", "scissors", "paper"}

func main() {
	rand.Seed(time.Now().UnixNano())
	fh := hs[rand.Intn(3)]
	sh := hs[rand.Intn(3)]

	taro := player.Main("human", false, "Taro", fh)
	bot := player.Main("computer", false, "ThinkPad", sh)
	player.Greeting(taro)
	player.Greeting(bot)
	fmt.Println("start game")
	game(taro, bot)
	fmt.Println("end game")
}

func changeHand(fp *player.Player, sp *player.Player) {
	rand.Seed(time.Now().UnixNano())
	fh := hs[rand.Intn(3)]
	sh := hs[rand.Intn(3)]

	fp.ChangeHand(fh)
	sp.ChangeHand(sh)

}

func game(fp *player.Player, sp *player.Player) {

	for {
		if fp.Hand == sp.Hand {
			fmt.Println(fp.Name, ":", fp.Hand, ",", sp.Name, ":", sp.Hand)
			fmt.Println("Draw")
			fmt.Println("Play again")
			changeHand(fp, sp)
		}

		if fp.Hand == "rock" && sp.Hand == "scissors" || fp.Hand == "scissors" && sp.Hand == "paper" || fp.Hand == "paper" && sp.Hand == "rock" {
			fmt.Println(fp.Name, ":", fp.Hand, ",", sp.Name, ":", sp.Hand)
			fmt.Println("1")
			fmt.Println(fp.Name, " win, game set")
			break
		}

		if fp.Hand == "rock" && sp.Hand == "paper" || fp.Hand == "scissors" && sp.Hand == "rock" || fp.Hand == "paper" && sp.Hand == "scissors" {
			fmt.Println(fp.Name, ":", fp.Hand, ",", sp.Name, ":", sp.Hand)
			fmt.Println("2")
			fmt.Println(sp.Name, " win, game set")
			break
		}
	}
}
