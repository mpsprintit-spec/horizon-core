package knowledge

// ConceptNode melambangkan satu simpul unik representasi makna di memori saraf.
type ConceptNode struct {
	Word        string
	Parent      *ConceptNode
	Children    map[string]*ConceptNode
	Relations   map[*ConceptNode]*Synapse
	Inhibitions map[*ConceptNode]bool
}
