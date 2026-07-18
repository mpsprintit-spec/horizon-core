package language

import (
	"fmt"
	"strings"

	"github.com/project-horizon/horizon-core/services/ai/thinking"
)

type Engine struct{}

func NewEngine() *Engine { return &Engine{} }

func (e *Engine) Generate(thought thinking.Thought) string {
	if len(thought.Concepts) == 0 {
		return "Saya belum memiliki aktivasi yang cukup untuk menjawab dengan yakin."
	}
	return fmt.Sprintf("Pemahaman yang muncul mengarah ke %s dengan confidence %.2f.", strings.Join(thought.Concepts, ", "), thought.Confidence)
}
