package knowledge

// Synapse bertindak sebagai jembatan relasi antar simpul konsep dengan bobot tertentu.
type Synapse struct {
	TargetNode *ConceptNode
	Weight     float64
}
