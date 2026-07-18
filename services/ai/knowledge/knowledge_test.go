package knowledge

import "testing"

func TestTokenRegistryPreventsDuplicateNodes(t *testing.T) {
	kb := NewKnowledgeBase()
	first := kb.Store(" Api ")
	second := kb.Store("api")
	if first.ID != second.ID {
		t.Fatalf("expected one unique node for canonical token, got %d and %d", first.ID, second.ID)
	}
	if len(kb.Registry.Nodes()) != 1 {
		t.Fatalf("expected one node, got %d", len(kb.Registry.Nodes()))
	}
}

func TestConnectStrengthensExistingSynapse(t *testing.T) {
	kb := NewKnowledgeBase()
	api := kb.Store("api")
	panas := kb.Store("panas")
	kb.Connect(api, panas, 0.4, 0.7, false)
	kb.Connect(api, panas, 0.8, 0.9, false)
	if len(api.Synapses) != 1 {
		t.Fatalf("expected one synapse, got %d", len(api.Synapses))
	}
	if api.Synapses[panas.ID].Frequency != 2 {
		t.Fatalf("expected reinforced synapse frequency 2, got %d", api.Synapses[panas.ID].Frequency)
	}
}
