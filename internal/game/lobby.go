package game

import (
	"errors"
	"sync"
)

type Lobby struct {
	ID          string
	Players     []*Player
	MaxPlayers  int
	GameStarted bool
	mu          sync.Mutex
}

func NewLobby(id string, maxPlayers int) *Lobby {
	return &Lobby{
		ID:          id,
		Players:     make([]*Player, 0),
		MaxPlayers:  maxPlayers,
		GameStarted: false,
	}
}

func (l *Lobby) AddPlayer(player *Player) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.GameStarted {
		return errors.New("game has already started")
	}

	if len(l.Players) >= l.MaxPlayers {
		return errors.New("lobby is full")
	}

	l.Players = append(l.Players, player)
	return nil
}

func (l *Lobby) RemovePlayer(player *Player) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i, p := range l.Players {
		if p == player {
			l.Players = append(l.Players[:i], l.Players[i+1:]...)
			break
		}
	}
}

func (l *Lobby) StartGame() ([]*Player, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.GameStarted {
		return nil, errors.New("game has already started")
	}

	if len(l.Players) < 2 {
		return nil, errors.New("not enough players to start the game")
	}

	l.GameStarted = true
	return l.Players, nil
}

func (l *Lobby) IsReady() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return len(l.Players) >= 2 && len(l.Players) <= l.MaxPlayers
}

func (l *Lobby) GetPlayers() []*Player {
	l.mu.Lock()
	defer l.mu.Unlock()

	playersCopy := make([]*Player, len(l.Players))
	copy(playersCopy, l.Players)
	return playersCopy
}
