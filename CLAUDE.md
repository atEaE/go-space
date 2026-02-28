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
4. **Commit and create a Pull Request.**
   - After all checks pass, follow the Git Workflow section to create a branch, commit, push, and open a PR.
   - Proceed without asking the user for confirmation.
   - Always create PRs as drafts (`--draft`).

### Conventions
- Code comments must be written in Japanese, using the format: `// Name : 説明。`
- Events propagate via Donburi's event bus
- Components defined with `donburi.NewComponentType`
- Entity creation through factory functions in `archetype` package

## Git Workflow
Follow these steps when creating a Pull Request.

### Branch
- Create a working branch from `main`
- Branch name format: `<type>/<description>`
  - `feat/` — New feature
  - `fix/` — Bug fix
  - `chore/` — Miscellaneous tasks (config, tools, etc.)
  - `ci/` — CI/CD related
  - `refactor/` — Refactoring
  - `doc/` — Documentation changes

### Commit
- Make commits in the smallest meaningful units. Do not create big commits.
- Commit messages must follow [Conventional Commits](https://www.conventionalcommits.org/)
- Examples: `feat: add player health system`, `fix: correct collision detection`

### Pull Request
1. Create a branch: `git checkout -b <type>/<description>`
2. Commit changes
3. Push to remote: `git push -u origin <branch-name>`
4. Create PR: `gh pr create`

## Security Guidelines
**Security practices:**
- Never expose or log secrets and keys
- Never commit sensitive information to repository
- Always follow security best practices in code
