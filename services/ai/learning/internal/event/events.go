package event

import "github.com/project-horizon/horizon-core/services/ai/domain"

type Name = domain.EventName
type Event = domain.DomainEvent

const (
	DecisionCreated             = domain.EventDecisionCreated
	BehaviorExecuted            = domain.EventBehaviorExecuted
	DroneCommandExecuted        = domain.EventDroneCommandExecuted
	HUDRendered                 = domain.EventHUDRendered
	MissionCompleted            = domain.EventMissionCompleted
	UserFeedbackReceived        = domain.EventUserFeedbackReceived
	FailureLogged               = domain.EventFailureLogged
	EnvironmentalChangeDetected = domain.EventEnvironmentalChangeDetected
	ExperienceRecorded          = domain.EventExperienceRecorded
	LearningUpdated             = domain.EventLearningUpdated
	AnomalyLearned              = domain.EventAnomalyLearned
	ModelImprovementRequested   = domain.EventModelImprovementRequested
	UserPreferenceLearned       = domain.EventUserPreferenceLearned
)
