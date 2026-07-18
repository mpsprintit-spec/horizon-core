package knowledge

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
	"time"
)

// TokenRegistry is the only authority allowed to create token nodes.
type TokenRegistry struct {
	mu      sync.RWMutex
	nextID  NodeID
	byID    map[NodeID]*ConceptNode
	byToken map[string]NodeID
}

func NewTokenRegistry() *TokenRegistry {
	return &TokenRegistry{nextID: 1, byID: make(map[NodeID]*ConceptNode), byToken: make(map[string]NodeID)}
}

func canonicalToken(token string) string { return strings.ToLower(strings.TrimSpace(token)) }

// GetOrCreate returns the single canonical node for token, creating it only if absent.
func (r *TokenRegistry) GetOrCreate(token string) (*ConceptNode, bool, error) {
	canonical := canonicalToken(token)
	if canonical == "" {
		return nil, false, errors.New("token is empty")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if id, ok := r.byToken[canonical]; ok {
		n := r.byID[id]
		n.Frequency++
		n.LastActivation = time.Now().UTC()
		return n, false, nil
	}
	n := newConceptNode(r.nextID, canonical)
	r.nextID++
	n.Frequency = 1
	r.byToken[canonical] = n.ID
	r.byID[n.ID] = n
	return n, true, nil
}

func (r *TokenRegistry) Get(token string) *ConceptNode {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.byID[r.byToken[canonicalToken(token)]]
}

func (r *TokenRegistry) GetByID(id NodeID) *ConceptNode {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.byID[id]
}

func (r *TokenRegistry) Nodes() []*ConceptNode {
	r.mu.RLock()
	defer r.mu.RUnlock()
	nodes := make([]*ConceptNode, 0, len(r.byID))
	for _, n := range r.byID {
		nodes = append(nodes, n)
	}
	return nodes
}

// KnowledgeBase is Horizon's neural semantic memory graph.
type KnowledgeBase struct{ Registry *TokenRegistry }

func NewKnowledgeBase() *KnowledgeBase                   { return &KnowledgeBase{Registry: NewTokenRegistry()} }
func (k *KnowledgeBase) Fetch(token string) *ConceptNode { return k.Registry.Get(token) }
func (k *KnowledgeBase) Store(token string) *ConceptNode {
	n, _, _ := k.Registry.GetOrCreate(token)
	return n
}

func (k *KnowledgeBase) Connect(source, target *ConceptNode, weight, confidence float64, inhibitory bool) {
	k.ConnectKind(source, target, RelationAssociation, weight, confidence, inhibitory)
}

func (k *KnowledgeBase) ConnectKind(source, target *ConceptNode, kind RelationKind, weight, confidence float64, inhibitory bool) {
	if source == nil || target == nil {
		return
	}
	now := time.Now().UTC()
	s, ok := source.Synapses[target.ID]
	if !ok {
		s = &Synapse{TargetID: target.ID, Kind: kind, Confidence: confidence, Inhibitory: inhibitory}
		source.Synapses[target.ID] = s
	}
	if s.Kind == "" {
		s.Kind = kind
	}
	s.Weight = clamp01((s.Weight*float64(s.Frequency) + weight) / float64(s.Frequency+1))
	s.Confidence = clamp01((s.Confidence + confidence) / 2)
	s.Frequency++
	s.LastActivation = now
}

// Optimize weakens stale or rarely reinforced paths without deleting nodes.
func (k *KnowledgeBase) Optimize(now time.Time, staleAfter time.Duration, decay float64) {
	if staleAfter <= 0 || decay <= 0 {
		return
	}
	for _, node := range k.Registry.Nodes() {
		for _, synapse := range node.Synapses {
			if synapse.LastActivation.IsZero() || now.Sub(synapse.LastActivation) > staleAfter {
				synapse.Weight = clamp01(synapse.Weight * (1 - decay))
			}
		}
	}
}

func (k *KnowledgeBase) Save(path string) error {
	b, e := json.MarshalIndent(k.Registry.Nodes(), "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, b, 0644)
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
