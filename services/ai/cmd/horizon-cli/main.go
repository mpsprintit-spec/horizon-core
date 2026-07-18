package main

import (
	"fmt"

	"github.com/project-horizon/horizon-core/services/ai/core"
	"github.com/project-horizon/horizon-core/services/ai/plugin"
)

func main() {
	fmt.Println("==================================================================")
	fmt.Println("         🧠 HORIZON NEURAL SEMANTIC MEMORY INITIALIZED            ")
	fmt.Println("==================================================================")

	horizon := core.NewHorizonEngine()
	horizon.Execution.RegisterPlugin("terbang", &plugin.DronePlugin{})
	horizon.Execution.RegisterPlugin("logsystem", &plugin.ChatbotPlugin{})

	fmt.Println("--- [PHASE 1: LEARNING WITHOUT SENTENCE STORAGE] ---")
	horizon.Learning.Assimilate("burung terbang", 0.95)
	horizon.Learning.Assimilate("burung elang terbang", 0.95)
	horizon.Learning.Assimilate("burung unta", 0.70)
	horizon.Learning.Inhibit("unta", "terbang", 0.95)
	fmt.Println()

	fmt.Println("--- [PHASE 2: ACTIVATION-BASED THINKING] ---")
	horizon.Pulse(core.TaskPulse{Stimulus: "burung elang", Context: "terbang", Data: "Koordinat Ketinggian 50m"})
	horizon.Pulse(core.TaskPulse{Stimulus: "burung unta", Context: "terbang", Data: "Koordinat Ketinggian 50m"})

	fmt.Println("==================================================================")
	fmt.Println("             🧠 ARCHITECTURE SIMULATION COMPLETED                 ")
	fmt.Println("==================================================================")
}
