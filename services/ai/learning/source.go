package learning

import "time"

type SourceEvidence struct {
	Source         string
	EvidenceCount  int
	Confidence     float64
	Reputation     float64
	Usage          int
	LastValidation time.Time
}
