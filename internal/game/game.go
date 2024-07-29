package game

import (
	"errors"
	"strings"
	"time"
)

type Game struct {
	Lobby       *Lobby
	Players     []*Player
	CurrentWord string
	Round       int
	TimeLimit   time.Duration
	WordBank    *WordBank
	Timer       *Timer
	IsActive    bool
	CurrentTurn int
}

func NewGame(lobby *Lobby, initialTimeLimit time.Duration) (*Game, error) {
	players := lobby.GetPlayers()
	if len(players) < 2 {
		return nil, errors.New("not enough players to start the game")
	}

	return &Game{
		Lobby:       lobby,
		Players:     players,
		Round:       0,
		TimeLimit:   initialTimeLimit,
		WordBank:    NewWordBank(),
		Timer:       NewTimer(initialTimeLimit),
		IsActive:    false,
		CurrentTurn: 0,
	}, nil
}

func (g *Game) Start() error {
	if g.IsActive {
		return errors.New("game is already in progress")
	}
	g.IsActive = true
	return g.nextRound()
}

func (g *Game) nextRound() error {
	g.Round++
	if g.Round > 10 {
		return errors.New("game over: max rounds reached")
	}

	g.CurrentWord = g.WordBank.GetWord(g.Round)
	if g.CurrentWord == "" {
		return errors.New("failed to get word for the round")
	}

	g.TimeLimit = g.TimeLimit - time.Second // Decrease time limit
	if g.TimeLimit < 5*time.Second {
		return errors.New("game over: time limit too short")
	}

	g.Timer.Reset(g.TimeLimit)
	g.CurrentTurn = 0
	return nil
}

func (g *Game) SubmitSentence(playerIndex int, sentence string) error {
	if !g.IsActive {
		return errors.New("game has not started")
	}
	if g.CurrentTurn != playerIndex {
		return errors.New("not this player's turn")
	}
	if g.Timer.IsExpired() {
		return errors.New("time's up")
	}
	if !g.containsCurrentWord(sentence) {
		return errors.New("sentence does not contain the current word")
	}

	g.Players[playerIndex].IncrementScore()
	g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)

	if g.CurrentTurn == 0 {
		return g.nextRound()
	}

	g.Timer.Reset(g.TimeLimit)
	return nil
}

func (g *Game) containsCurrentWord(sentence string) bool {
	return strings.Contains(strings.ToLower(sentence), strings.ToLower(g.CurrentWord))
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.Players[g.CurrentTurn]
}

func (g *Game) IsGameOver() bool {
	return !g.IsActive || g.Round > 10 || g.TimeLimit < 5*time.Second
}

func (g *Game) GetWinner() *Player {
	if len(g.Players) == 0 {
		return nil
	}

	winner := g.Players[0]
	for _, player := range g.Players[1:] {
		if player.Score > winner.Score {
			winner = player
		}
	}
	return winner
}

func (g *Game) GetGameState() GameState {
	return GameState{
		Round:         g.Round,
		CurrentWord:   g.CurrentWord,
		CurrentTurn:   g.CurrentTurn,
		TimeRemaining: g.Timer.RemainingTime(),
		Players:       g.Players,
		IsActive:      g.IsActive,
	}
}

type GameState struct {
	Round         int
	CurrentWord   string
	CurrentTurn   int
	TimeRemaining time.Duration
	Players       []*Player
	IsActive      bool
}
