package perception

import (
	"strings"
	"time"
)

type PerceptionKind string

const (
	PerceptionUserInput PerceptionKind = "user_input"
	PerceptionWebSearch PerceptionKind = "web_search"
)

type PerceptionSignal struct {
	Kind       PerceptionKind
	Source     string
	Tokens     []string
	Confidence float64
	ObservedAt time.Time
}

type PerceptionLayer interface {
	Perceive(input string) ([]PerceptionSignal, error)
}

type UserInputPerception struct{}

func (UserInputPerception) Perceive(input string) ([]PerceptionSignal, error) {
	return []PerceptionSignal{{Kind: PerceptionUserInput, Source: "user", Tokens: strings.Fields(strings.ToLower(strings.TrimSpace(input))), Confidence: 1, ObservedAt: time.Now().UTC()}}, nil
}
