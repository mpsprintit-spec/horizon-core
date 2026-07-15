package grpc

import (
	"context"
	"github.com/project-horizon/horizon-core/services/ai/learning/internal/model"
)

type RecordExperienceHandler interface {
	RecordExperience(context.Context, model.LearningRecord) error
}
type ApplyFeedbackHandler interface {
	ApplyFeedback(context.Context, model.LearningRecord) error
}
type GetLearningStateHandler interface {
	GetLearningState(context.Context, string) (model.ExperienceMemory, error)
}
type ListModelsHandler interface {
	ListModels(context.Context) ([]model.ModelVersion, error)
}
