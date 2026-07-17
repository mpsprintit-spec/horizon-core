// Package model contains shared data structures for the thinking engine.
package model

import "time"

// ThinkingRequest represents an incoming request to the thinking engine.
type ThinkingRequest struct {
	// ID is a unique identifier for the request.
	ID string `json:"id"`
	// UserID identifies the user making the request.
	UserID string `json:"user_id"`
	// Input is the raw user input or query.
	Input string `json:"input"`
	// ConversationHistory provides recent conversation context.
	ConversationHistory []ConversationTurn `json:"conversation_history,omitempty"`
	// SessionID for maintaining state across interactions.
	SessionID string `json:"session_id,omitempty"`
	// Timestamp when the request was received.
	Timestamp time.Time `json:"timestamp"`
}

// ConversationTurn represents a single turn in conversation history.
type ConversationTurn struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// Action represents a command or tool call extracted from analysis.
type Action struct {
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
	Target string                 `json:"target,omitempty"`
}

// LearningRequest represents knowledge to be stored after approval.
type LearningRequest struct {
	Content    string                 `json:"content"`
	Source     string                 `json:"source"`
	Confidence float64                `json:"confidence"`
	Tags       []string               `json:"tags,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}
