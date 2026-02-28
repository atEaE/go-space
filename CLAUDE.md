# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Important Notes
- **Always respond in Japanese**

## Code implementation
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

## Security Guidelines
**Security practices:**
- Never expose or log secrets and keys
- Never commit sensitive information to repository
- Always follow security best practices in code
