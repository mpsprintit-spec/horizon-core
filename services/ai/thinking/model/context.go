package model

// Context holds all assembled information for thinking.
type Context struct {
	// UserInput is the original request.
	UserInput string `json:"user_input"`
	// ConversationHistory from previous interactions.
	ConversationHistory []ConversationTurn `json:"conversation_history"`
	// KnowledgeResults from long-term memory.
	KnowledgeResults []string `json:"knowledge_results,omitempty"`
	// SearchResults from web or external search.
	SearchResults []string `json:"search_results,omitempty"`
	// DeviceState current state of connected devices.
	DeviceState map[string]interface{} `json:"device_state,omitempty"`
	// SessionContext additional session data.
	SessionContext map[string]interface{} `json:"session_context,omitempty"`
	// Metadata for tracking.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Capability describes available modules in the system.
type Capability struct {
	// HasKnowledge indicates knowledge base availability.
	HasKnowledge bool `json:"has_knowledge"`
	// HasWebSearch indicates search capability.
	HasWebSearch bool `json:"has_web_search"`
	// HasDevices indicates device control capability.
	HasDevices bool `json:"has_devices"`
	// HasLearning indicates learning module availability.
	HasLearning bool `json:"has_learning"`
	// SupportedProviders list of available LLM providers.
	SupportedProviders []string `json:"supported_providers"`
}
