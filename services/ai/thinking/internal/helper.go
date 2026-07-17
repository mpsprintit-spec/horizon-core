// Package internal contains private utilities for the thinking package.
package internal

import "strings"

// CleanPrompt removes unnecessary whitespace and normalizes prompt.
func CleanPrompt(prompt string) string {
	return strings.TrimSpace(prompt)
}

// ExtractJSON attempts to extract JSON from LLM response if structured.
func ExtractJSON(resp string) (string, bool) {
	// Simple implementation - in production use regex or parser.
	start := strings.Index(resp, "{")
	if start == -1 {
		return "", false
	}
	// Basic extraction logic
	return resp[start:], true
}

// ValidateConfidence ensures confidence score is in valid range.
func ValidateConfidence(c float64) float64 {
	if c < 0 {
		return 0
	}
	if c > 1 {
		return 1
	}
	return c
}
