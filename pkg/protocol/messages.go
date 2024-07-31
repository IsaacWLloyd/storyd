package protocol

import (
	"time"
)

type MessageType string

const (
	SubmitSentence MessageType = "SUBMIT_SENTENCE"
	RequestGameState MessageType = "REQUEST_GAME_STATE"
	GameState MessageType = "GAME_STATE"
	PlayerJoined MessageType = "PLAYER_JOINED"
	// Add more message types as needed
)

type Message struct {
	Type    MessageType     `json:"type"`
	Content json.RawMessage `json:"content"`
}

type SubmitSentenceContent struct {
	Sentence string `json:"sentence"`
}

type GameStateContent struct {
	Round         int           `json:"round"`
	CurrentWord   string        `json:"currentWord"`
	TimeRemaining time.Duration `json:"timeRemaining"`
	Players       []PlayerState `json:"players"`
}

type PlayerState struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type PlayerJoinedContent struct {
	PlayerName string `json:"playerName"`
}

// Add more message content structures as needed
