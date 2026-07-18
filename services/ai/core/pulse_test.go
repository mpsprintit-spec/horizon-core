package core

import "testing"

type recordingPlugin struct{ triggered []string }

func (p *recordingPlugin) Name() string { return "recording" }
func (p *recordingPlugin) Trigger(data string) {
	p.triggered = append(p.triggered, data)
}

func TestPulseKeepsContextSelectableAsAction(t *testing.T) {
	horizon := NewHorizonEngine()
	drone := &recordingPlugin{}
	fallback := &recordingPlugin{}
	horizon.Execution.RegisterPlugin("terbang", drone)
	horizon.Execution.RegisterPlugin("LogSystem", fallback)
	horizon.Learning.Assimilate("burung terbang", 0.95)
	horizon.Learning.Assimilate("burung elang terbang", 0.95)

	horizon.Pulse(TaskPulse{Stimulus: "burung elang", Context: "terbang", Data: "Koordinat Ketinggian 50m"})

	if len(drone.triggered) != 1 || drone.triggered[0] != "Koordinat Ketinggian 50m" {
		t.Fatalf("expected terbang plugin to receive pulse data, got %#v", drone.triggered)
	}
	if len(fallback.triggered) != 0 {
		t.Fatalf("expected no fallback dispatch, got %#v", fallback.triggered)
	}
}
