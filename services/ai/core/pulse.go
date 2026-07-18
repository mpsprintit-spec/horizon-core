package core

import (
	"fmt"
	"time"
)

func (h *HorizonEngine) Pulse(task TaskPulse) {
	fmt.Printf("⚡ [ENGINE] Memproses impuls saraf: '%s' dengan konteks '%s'...\n", task.Stimulus, task.Context)
	time.Sleep(50 * time.Millisecond)

	decision, success := h.Thinking.Reason(task.Stimulus, task.Context)
	if !success {
		fmt.Println("📭 [ENGINE] Arus kognisi meredup. Tidak ada tindakan aman yang diambil.\n")
		return
	}

	fmt.Printf("💡 [ENGINE] Hasil penalaran didapatkan: [%s]. Mengirim ke Kubu Execution...\n", decision)

	h.Execution.Dispatch(decision, task.Data)
	fmt.Println()
}
