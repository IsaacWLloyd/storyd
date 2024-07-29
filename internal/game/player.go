package game

// Player represents a player in the game
type Player struct {
	Name  string
	Score int
}

// NewPlayer creates and initializes a new Player instance
func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Score: 0,
	}
}

// IncrementScore increases the player's score by 1
func (p *Player) IncrementScore() {
	p.Score++
}
