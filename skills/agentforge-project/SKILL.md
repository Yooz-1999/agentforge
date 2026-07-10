---
name: agentforge-project
description: AgentForge backend project memory and workflow guide. Use when working in this repository, especially for local startup, go-zero service wiring, MySQL/Redis configuration, implemented API status, project conventions, or when deciding whether new conversation learnings should be recorded for future Codex sessions.
---

# AgentForge Project

## Required First Step

Read `references/project-memory.md` before making project changes or explaining local startup.

That file contains confirmed project facts, current local development choices, known unfinished areas, and recurring pitfalls.

## How To Use This Skill

Use this skill as the project memory for AgentForge backend work.

Before coding:

1. Check the current request against `references/project-memory.md`
2. Prefer confirmed project facts over assumptions
3. If the user reports a startup or config issue, compare it with the known local workflow first
4. Keep README and docs in Chinese unless the user explicitly asks otherwise

After coding:

1. Run focused validation for the changed area
2. Run `go test ./...` for backend-wide changes
3. Run `go vet ./...` when changing shared logic, startup code, config, or generated-facing behavior
4. Update `references/project-memory.md` if the conversation produced durable knowledge

## When To Update Memory

Update `references/project-memory.md` when a fact will likely matter in a future session.

Good examples:

1. A confirmed local environment detail
2. A chosen architecture rule
3. A recurring startup command
4. A known pitfall and its fix
5. A user preference about project conventions
6. The real implementation status of major APIs

Do not store:

1. Passwords, API keys, tokens, or private secrets
2. One-off test data
3. Guesses that were not confirmed
4. Large logs or command output
5. Temporary debugging notes that will not matter later

When adding memory, write short entries and clearly label them as confirmed facts, decisions, or open items.
