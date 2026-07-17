package model

import "time"

// ThinkingResponse is the structured output from the thinking engine.
type ThinkingResponse struct {
	// ID matches the request ID.
	ID string `json:"id"`
	// Output is the final response to the user.
	Output string `json:"output"`
	// Confidence indicates how certain the engine is about the response.
	Confidence float64 `json:"confidence"`
	// Actions lists any tool/device actions to be executed.
	Actions []Action `json:"actions,omitempty"`
	// LearningSuggestions contains potential knowledge items for review.
	LearningSuggestions []LearningRequest `json:"learning_suggestions,omitempty"`
	// Metadata for additional context or debugging.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Timestamp of response generation.
	Timestamp time.Time `json:"timestamp"`
}

// Analysis represents the parsed result from LLM output.
type Analysis struct {
	Intent        string            `json:"intent"`
	Confidence    float64           `json:"confidence"`
	Actions       []Action          `json:"actions,omitempty"`
	LearningItems []LearningRequest `json:"learning_items,omitempty"`
	ResponseText  string            `json:"response_text"`
	ToolRequests  []ToolRequest     `json:"tool_requests,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// ToolRequest represents a request to use external tools.
type ToolRequest struct {
	ToolType string                 `json:"tool_type"`
	Params   map[string]interface{} `json:"params"`
}
