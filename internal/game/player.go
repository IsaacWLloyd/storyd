package game

type Player struct {
	Name  string
	Score int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Score: 0,
	}
}

func (p *Player) IncrementScore() {
	p.Score++
}
