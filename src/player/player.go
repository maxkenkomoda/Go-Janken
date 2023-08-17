package player

import "fmt"

type PlayerType string

const (
	Human    PlayerType = "human"
	Computer PlayerType = "computer"
)

type Player struct {
	playerType PlayerType
	isDead     bool
	name       string
}

func Main() {
	plyer := &Player{
		playerType: Human,
		isDead:     false,
		name:       "Player 1",
	}
	fmt.Println(plyer.CreateGreetingText())
}

func (p *Player) CreateGreetingText() string {
	return "Hello, " + p.name + "!"
}
