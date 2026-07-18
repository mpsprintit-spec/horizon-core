package learning

import (
	"testing"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

func TestAssimilateLabelsCausalRelations(t *testing.T) {
	kb := knowledge.NewKnowledgeBase()
	NewLearningUnit(kb).Assimilate("api menyebabkan panas", 0.9)
	api := kb.Fetch("api")
	panas := kb.Fetch("panas")
	if api.Synapses[panas.ID].Kind != knowledge.RelationCause {
		t.Fatalf("expected cause relation, got %s", api.Synapses[panas.ID].Kind)
	}
}

func TestOptimizeWeakensStaleSynapse(t *testing.T) {
	kb := knowledge.NewKnowledgeBase()
	api := kb.Store("api")
	panas := kb.Store("panas")
	kb.Connect(api, panas, 1, 1, false)
	api.Synapses[panas.ID].LastActivation = time.Now().Add(-60 * 24 * time.Hour)
	NewLearningUnit(kb).Optimize(time.Now())
	if api.Synapses[panas.ID].Weight >= 1 {
		t.Fatal("expected stale synapse weight to decay")
	}
}
