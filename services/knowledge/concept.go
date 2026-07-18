package knowledge

// ConceptNode adalah representasi SATU sel kata/makna unik di memori.
type ConceptNode struct {
	Word        string
	Parent      *ConceptNode
	Children    map[string]*ConceptNode
	Relations   map[*ConceptNode]*Synapse
	Inhibitions map[*ConceptNode]bool
}
