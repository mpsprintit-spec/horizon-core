package handler

import (
	"context"
	"github.com/project-horizon/horizon-core/services/ai/learning/internal/event"
	"github.com/project-horizon/horizon-core/services/ai/learning/internal/model"
)

type EventHandler interface {
	HandleLearningEvent(context.Context, event.Event) error
}
type FeedbackHandler interface {
	ApplyFeedback(context.Context, model.LearningRecord) error
}
type ModelUpdateHandler interface {
	RequestModelImprovement(context.Context, model.ModelVersion, []model.LearningRecord) error
}
