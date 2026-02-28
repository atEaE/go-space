# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Important Notes
- **Always respond in Japanese**

## Tech Stack
- **Go** 1.25.2
- **Ebiten** v2.9.8 (2D game engine)
- **Donburi** v1.15.7 (ECS framework)
- **golangci-lint** (configured in `.golangci.yml`)

## Architecture
ECS (Entity Component System) architecture.

### Package Structure
- `main.go` — Entry point (window 960x720)
- `internal/game` — Game struct, scene management (Title/Playing), ECS setup
- `internal/component` — ECS component & tag definitions
- `internal/archetype` — Entity factory functions
- `internal/system` — ECS systems (update logic) and renderers
- `internal/layer` — Render layer order
- `internal/event` — Event type definitions
- `internal/config` — Screen size constants

## Code Implementation
When changing or creating code, you must follow the rules below.

### Steps
1. **Always create an implementation plan.**
   - If there is insufficient information when creating the implementation plan, ask the user additional questions to supplement the plan.
   - Repeat questions and answers until sufficient information is obtained.
2. Implement the code based on the implementation plan.
   - Do not work with a single agent; **always work with multiple agents.**
   - The supervising agent should assign specific implementation tasks to the working agents based on the work plan. They should also manage the working agents' progress and provide support as needed.
   - The working agents should complete the assigned implementation tasks.
3. Confirm the operation and repeat modifications until the problem is solved.
   - Run the `go tool golangci-lint run ./...` command and confirm that lint check passes.
   - Run the `go test ./... -cover` command and confirm that all tests pass.
   - Run the `go build -o ./out/go-space` command and confirm that the build succeeds.

### Conventions
- Code comments must be written in Japanese, using the format: `// Name : 説明。`
- Events propagate via Donburi's event bus
- Components defined with `donburi.NewComponentType`
- Entity creation through factory functions in `archetype` package

## Security Guidelines
**Security practices:**
- Never expose or log secrets and keys
- Never commit sensitive information to repository
- Always follow security best practices in code
