package thinking

import (
	"fmt"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

func (t *ThinkingEngine) Reason(stimulusStr, konteksStr string) (string, bool) {
	stimulus := t.Kb.Fetch(stimulusStr)
	konteks := t.Kb.Fetch(konteksStr)

	if stimulus == nil || konteks == nil {
		return "", false
	}

	if stimulus.Inhibitions[konteks] {
		fmt.Printf("🛑 [THINKING] Rem Kognitif Aktif! Memutus sirkuit untuk '%s'.\n", stimulusStr)
		return "", false
	}

	curr := stimulus
	for curr != nil {
		synapse, terhubung := curr.Relations[konteks]
		if terhubung && synapse.Weight >= 0.5 {
			return synapse.TargetNode.Word, true
		}
		curr = curr.Parent
	}

	return "", false
}
