# Horizon Thinking Engine

## Purpose
The Thinking Engine is the central intelligence module of Horizon Core. It orchestrates all cognitive processes, decision making, and workflow execution. The LLM serves as an internal tool rather than the primary brain.

## Architecture
- **Clean Architecture**: Separation of concerns with clear boundaries.
- **SOLID Principles**: Single responsibility, open/closed, etc.
- **Interface-Driven**: All major components are abstracted via interfaces.
- **Dependency Injection**: Constructor-based injection for testability and flexibility.

## Responsibilities
- Context building from multiple sources.
- Capability assessment.
- Prompt construction and LLM orchestration.
- Response analysis and action extraction.
- Knowledge integration and learning decisions.
- Error-resilient workflow execution.

## Workflow
1. Receive user request
2. Build comprehensive context
3. Check available capabilities
4. Determine need for knowledge/search
5. Generate prompt
6. Invoke LLM via adapter
7. Analyze response
8. Handle learning if needed
9. Return structured response

## Directory Structure
```
services/thinking/
├── README.md
├── engine.go          # Main orchestrator
├── interface.go       # Public interfaces
├── workflow.go        # Execution flow control
├── context.go         # Context assembly
├── prompt.go          # Prompt generation
├── analyzer.go        # Response parsing
├── capability.go      # Capability checks
├── llm/               # LLM adapter layer
├── model/             # Shared data models
└── internal/          # Private helpers
```

## Dependencies
- Go 1.21+
- External providers configurable via config
- Context from Horizon Core modules (Knowledge, WebSearch, etc.)

## Examples
```go
engine, err := thinking.NewEngine(config, deps)
resp, err := engine.Think(ctx, req)
```

This module is production-ready and follows idiomatic Go standards.