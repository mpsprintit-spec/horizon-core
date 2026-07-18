package learning

import (
	"strings"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

// Assimilate learns a transient phrase by tokenizing it, reusing registry nodes,
// and strengthening co-activation synapses. The sentence itself is discarded.
func (l *LearningUnit) Assimilate(text string, weight float64) {
	tokens := tokenize(text)
	if len(tokens) == 0 {
		return
	}
	var nodes []*knowledge.ConceptNode
	for _, token := range tokens {
		nodes = append(nodes, l.Kb.Store(token))
	}
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if i != j {
				l.Kb.ConnectKind(nodes[i], nodes[j], relationKind(tokens, i, j), weight, 0.8, false)
			}
		}
	}
	l.buildShortcuts(nodes, weight)
}

func (l *LearningUnit) Inhibit(a, b string, weight float64) {
	l.Kb.ConnectKind(l.Kb.Store(a), l.Kb.Store(b), knowledge.RelationAssociation, weight, 0.8, true)
}

func (l *LearningUnit) Optimize(now time.Time) {
	l.Kb.Optimize(now, 30*24*time.Hour, 0.08)
}

func (l *LearningUnit) buildShortcuts(nodes []*knowledge.ConceptNode, weight float64) {
	for i := 0; i+2 < len(nodes); i++ {
		l.Kb.ConnectKind(nodes[i], nodes[i+2], knowledge.RelationAssociation, weight*0.55, 0.55, false)
	}
}

func tokenize(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

func relationKind(tokens []string, i, j int) knowledge.RelationKind {
	between := tokens[min(i, j)+1 : max(i, j)]
	for _, token := range between {
		switch token {
		case "karena", "menyebabkan", "sebab":
			return knowledge.RelationCause
		case "untuk", "fungsi", "berguna":
			return knowledge.RelationFunction
		case "bagian", "memiliki", "punya":
			return knowledge.RelationPartWhole
		case "di", "dari", "lokasi":
			return knowledge.RelationLocation
		case "sebelum", "sesudah", "ketika", "saat":
			return knowledge.RelationTime
		}
	}
	return knowledge.RelationAssociation
}
