package repository

import (
	"context"
	"github.com/project-horizon/horizon-core/services/ai/learning/internal/model"
)

type LearningRecordRepository interface {
	SaveLearningRecord(context.Context, model.LearningRecord) error
	GetLearningRecord(context.Context, string) (model.LearningRecord, error)
}
type ExperienceMemoryRepository interface {
	SaveExperienceMemory(context.Context, model.ExperienceMemory) error
	GetExperienceMemory(context.Context, string) (model.ExperienceMemory, error)
}
type ModelVersionRepository interface {
	SaveModelVersion(context.Context, model.ModelVersion) error
	GetModelVersion(context.Context, string) (model.ModelVersion, error)
}
type UserPreferenceRepository interface {
	SaveUserPreference(context.Context, model.UserContext) error
	GetUserPreference(context.Context, string) (model.UserContext, error)
}
