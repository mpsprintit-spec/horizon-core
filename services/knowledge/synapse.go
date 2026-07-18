package knowledge

// Synapse bertindak sebagai kabel penghubung antar kata yang memiliki beban energi dinamis
type Synapse struct {
	TargetNode *ConceptNode
	Weight     float64
}
