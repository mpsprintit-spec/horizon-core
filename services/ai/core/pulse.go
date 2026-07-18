package core

import (
	"fmt"
	"strings"
	"time"
)

func (h *HorizonEngine) Pulse(task TaskPulse) {
	prompt := strings.TrimSpace(task.Stimulus + " " + task.Context)
	fmt.Printf("⚡ [ENGINE] Mengaktifkan pulsa saraf: %q...\n", prompt)
	time.Sleep(50 * time.Millisecond)

	thought, success := h.Thinking.Think(prompt, strings.Fields(strings.ToLower(task.Context)))
	if !success {
		fmt.Println("📭 [ENGINE] Aktivasi belum mencapai ambang konvergensi.")
		return
	}

	decision := thought.Concepts[0]
	fmt.Printf("💡 [ENGINE] %s\n", h.Language.Generate(thought))
	h.Execution.Dispatch(decision, task.Data)
	fmt.Println()
}
