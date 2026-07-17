package thinking

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/project-horizon/horizon-core/services/ai/thinking/internal"
	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// responseAnalyzer implements ResponseAnalyzer.
type responseAnalyzer struct{}

// NewResponseAnalyzer creates a new analyzer.
func NewResponseAnalyzer() ResponseAnalyzer {
	return &responseAnalyzer{}
}

// Analyze parses the LLM response into structured Analysis.
func (a *responseAnalyzer) Analyze(ctx context.Context, llmResp string) (*model.Analysis, error) {
	if llmResp == "" {
		return nil, fmt.Errorf("empty LLM response")
	}

	// Try to extract JSON
	jsonStr, ok := internal.ExtractJSON(llmResp)
	if !ok {
		// Fallback to simple text response
		return &model.Analysis{
			Intent:       "general_response",
			Confidence:   0.7,
			ResponseText: llmResp,
		}, nil
	}

	var analysis model.Analysis
	if err := json.Unmarshal([]byte(jsonStr), &analysis); err != nil {
		// Partial parse fallback
		return &model.Analysis{
			Intent:       "parse_failed",
			Confidence:   0.5,
			ResponseText: llmResp,
		}, fmt.Errorf("json parse error: %w", err)
	}

	analysis.Confidence = internal.ValidateConfidence(analysis.Confidence)
	return &analysis, nil
}
