package knowledge

import "sync"

// KnowledgeBase mengelola repositori penyimpanan ingatan jangka panjang di RAM.
type KnowledgeBase struct {
	Mu   sync.RWMutex
	Data map[string]*ConceptNode
}

func NewKnowledgeBase() *KnowledgeBase {
	return &KnowledgeBase{Data: make(map[string]*ConceptNode)}
}

func (k *KnowledgeBase) Fetch(word string) *ConceptNode {
	k.Mu.RLock()
	defer k.Mu.RUnlock()
	return k.Data[word]
}

func (k *KnowledgeBase) Store(word string) *ConceptNode {
	k.Mu.Lock()
	defer k.Mu.Unlock()
	if node, exists := k.Data[word]; exists {
		return node
	}
	newNode := &ConceptNode{
		Word:        word,
		Children:    make(map[string]*ConceptNode),
		Relations:   make(map[*ConceptNode]*Synapse),
		Inhibitions: make(map[*ConceptNode]bool),
	}
	k.Data[word] = newNode
	return newNode
}
