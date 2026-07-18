package language

import (
	"fmt"
	"strings"

	"github.com/project-horizon/horizon-core/services/ai/thinking"
)

type Engine struct{}

func NewEngine() *Engine { return &Engine{} }

func (e *Engine) Generate(thought thinking.Thought) string {
	if len(thought.Concepts) == 0 || thought.Confidence < 0.25 {
		return "Informasi internal belum cukup untuk menjawab dengan yakin."
	}
	if thought.NeedsWebSearch {
		return fmt.Sprintf("Pemahaman sementara mengarah ke %s dengan confidence %.2f, tetapi Horizon membutuhkan persepsi tambahan untuk validasi.", strings.Join(thought.Concepts, ", "), thought.Confidence)
	}
	return fmt.Sprintf("Pemahaman yang muncul mengarah ke %s dengan confidence %.2f.", strings.Join(thought.Concepts, ", "), thought.Confidence)
}
