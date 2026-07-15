# Learning Engine

Learning Engine is the HC-002 scaffold for asynchronous experience memory, feedback, model-improvement coordination, user preference learning, anomaly learning, and performance improvement.

It reuses shared contracts from `services/ai/domain` and does not implement AI models, training frameworks, or synchronous control-path logic.

## HTTP

- `POST /ai/learning/experience`
- `POST /ai/learning/feedback`
- `GET /ai/learning/state/{missionId}`
- `GET /ai/learning/models`

## gRPC

- `LearningEngineService.RecordExperience`
- `LearningEngineService.ApplyFeedback`
- `LearningEngineService.GetLearningState`
