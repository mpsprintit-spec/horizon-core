package main

import (
	"fmt"
	"github.com/project-horizon/horizon-core/services/ai/core"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
	"github.com/project-horizon/horizon-core/services/ai/plugin"
)

func main() {
	fmt.Println("==================================================================")
	fmt.Println("         🧠 HORIZON DECOUPLED ARCHITECTURE INITIALIZED            ")
	fmt.Println("==================================================================")

	// Inisialisasi engine utama lewat core layanan internal
	horizon := core.NewHorizonEngine()

	drone := &plugin.DronePlugin{}
	screen := &plugin.ChatbotPlugin{}

	horizon.Execution.RegisterPlugin("Terbang", drone)
	horizon.Execution.RegisterPlugin("LogSystem", screen)

	fmt.Println("--- [PHASE 1: LEARNING & ASIMILASI KNOWLEDGE] ---")
	burung := horizon.Knowledge.Store("Burung")
	fungsi := horizon.Knowledge.Store("Fungsi Utama")
	terbang := horizon.Knowledge.Store("Terbang")
	
	burung.Relations[fungsi] = &knowledge.Synapse{TargetNode: terbang, Weight: 0.95}

	elang := horizon.Knowledge.Store("Burung Elang")
	elang.Parent = burung
	burung.Children["Burung Elang"] = elang

	unta := horizon.Knowledge.Store("Burung Unta")
	unta.Parent = burung
	burung.Children["Burung Unta"] = unta
	horizon.Learning.Assimilate("Burung Unta", "Fungsi Utama", "Terbang", 0.95)

	horizon.Learning.Assimilate("Definisi", "Makna Hakiki", "Menjelaskan Esensi Konsep", 0.99)
	fmt.Println()

	fmt.Println("--- [PHASE 2: ENGINE PULSE & HIGH-SPEED EXECUTION] ---")

	horizon.Pulse(core.TaskPulse{Stimulus: "Burung Elang", Context: "Fungsi Utama", Data: "Koordinat Ketinggian 50m"})
	horizon.Pulse(core.TaskPulse{Stimulus: "Burung Unta", Context: "Fungsi Utama", Data: "Koordinat Ketinggian 50m"})
	horizon.Pulse(core.TaskPulse{Stimulus: "Definisi", Context: "Makna Hakiki", Data: "Konteks Logika"})

	fmt.Println("==================================================================")
	fmt.Println("             🧠 ARCHITECTURE SIMULATION COMPLETED                 ")
	fmt.Println("==================================================================")
}

