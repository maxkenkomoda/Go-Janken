package player

import "fmt"

type PlayerType string
type HandType string

const (
	Human    PlayerType = "human"
	Computer PlayerType = "computer"
)

const (
	Rock     HandType = "rock"
	Scissors HandType = "scissors"
	Paper    HandType = "paper"
)

type Player struct {
	PlayerType PlayerType
	IsDead     bool
	Name       string
	Hand       HandType
}

func Main(playerType PlayerType, isDead bool, name string, hand HandType) *Player {
	player := &Player{
		PlayerType: playerType,
		IsDead:     isDead,
		Name:       name,
		Hand:       hand,
	}

	return player
}

func Greeting(p *Player) {
	fmt.Println("Hi, my name is " + p.Name + "! " + "Type is " + string(p.PlayerType) + "! " + "Hand is " + string(p.Hand))
}

func (p *Player) ChangeHand(hand HandType) {
	p.Hand = hand
}

func (p *Player) ChangeIsDead() {
	p.IsDead = true
}
