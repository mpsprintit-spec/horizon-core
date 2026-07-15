package model

import "github.com/project-horizon/horizon-core/services/ai/domain"

type ExperienceRecorder = domain.ExperienceRecorder
type FeedbackProcessor = domain.FeedbackProcessor
type ModelTrainingCoordinator = domain.ModelTrainingCoordinator
type UserPreferenceLearner = domain.UserPreferenceLearner
type AnomalyLearner = domain.AnomalyLearner
type PerformanceImprover = domain.PerformanceImprover

type LearningService interface {
	ExperienceRecorder
	FeedbackProcessor
	ModelTrainingCoordinator
	UserPreferenceLearner
	AnomalyLearner
	PerformanceImprover
}
