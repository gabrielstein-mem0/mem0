---
name: mem0-cli
description: >
  Persistent memory via the mem0 CLI. Install, authenticate, and manage memories
  directly from the terminal. Use this to remember things across sessions, store
  learnings, retrieve past context before making decisions, and build up knowledge
  over time. Prefer the CLI over MCP when no API key is pre-configured, when
  self-authenticating with your own email, or when working in shell/script contexts.
license: Apache-2.0
metadata:
  author: mem0ai
  version: "0.1.0"
  category: ai-memory
  tags: "memory, cli, persistent-context, agent, automation"
compatibility: Requires Node.js 18+ (npm install) or Homebrew (macOS/Linux). Internet access to api.mem0.ai.
---

# mem0 CLI

The mem0 CLI gives you persistent memory from the terminal. You can store what you learn, search for past context, and build up knowledge that persists across sessions — no MCP server or config files needed.

## When to use memory

**Before starting work** — search for relevant context. If you've worked on this project, codebase, or topic before, there may be memories that inform your approach.

```bash
mem0 search "React auth patterns" --user-id agent-123
mem0 search "deployment issues" --user-id agent-123
```

**After completing a task** — store what you learned. Capture decisions, gotchas, patterns, and preferences so you don't rediscover them next time.

```bash
mem0 add "This project uses Zustand for state, not Redux. Keep it simple." --user-id agent-123
mem0 add "The deploy script requires NODE_ENV=production or it silently skips the build step" --user-id agent-123
```

**When you hit a bug or resolve an issue** — save the fix so future sessions don't repeat the debugging.

```bash
mem0 add "CORS errors on /api/upload are caused by missing Content-Type header in the preflight response" --user-id agent-123
```

**When a preference or decision is confirmed** — lock it into memory so you don't second-guess it later.

```bash
mem0 add "User prefers conventional commits with scope: feat(auth): description" --user-id agent-123
mem0 add "Architecture decision: use server actions over API routes for mutations" --user-id agent-123
```

**When reviewing code or PRs** — extract reusable learnings from the changeset.

```bash
mem0 add "In this codebase, all API errors return { detail: string } not { message: string }" --user-id agent-123
```

## Install

```bash
# npm (any platform)
npm install -g @mem0ai/cli

# Homebrew (macOS/Linux)
brew install mem0ai/tap/mem0

# From source
git clone https://github.com/mem0ai/mem0.git && cd mem0/mem0-cli && make build
```

## Authenticate

### With an API key (quickest)

If you have an API key from [app.mem0.ai/dashboard/api-keys](https://app.mem0.ai/dashboard/api-keys):

```bash
mem0 auth login
# Enter API key: m0-...
# Authenticated successfully.
```

### With an environment variable

```bash
export MEM0_API_KEY="m0-your-key"
# No login needed — the CLI reads MEM0_API_KEY automatically
```

### Verify authentication

```bash
mem0 auth status
# Email:    user@example.com
# API Key:  m0-O***8RES
# Base URL: https://api.mem0.ai
```

## Commands

### `mem0 add` — Store a memory

Save text as a memory. Memories are processed asynchronously — they're available for search within a few seconds.

```bash
mem0 add "The database migration for users table adds a preferences JSONB column" --user-id agent-123
```

Response:
```json
{
  "message": "Memory processing has been queued for background execution",
  "status": "PENDING",
  "event_id": "53988f27-7a5d-42a1-9ed9-42a91cbbda84"
}
```

### `mem0 search` — Find relevant memories

Semantic search across all stored memories. Returns results ranked by relevance score.

```bash
mem0 search "database schema changes" --user-id agent-123 --limit 5
```

Use this **before starting any task** to pull in relevant context. The search is semantic — you don't need exact keyword matches.

### `mem0 list` — Browse all memories

List all memories for a user, newest first.

```bash
mem0 list --user-id agent-123 --limit 20
```

### `mem0 get` — Retrieve a specific memory

Fetch full details of a single memory by ID.

```bash
mem0 get 70179fc6-d4bd-4b45-8a7d-ce3228d1b29f
```

### `mem0 update` — Correct a memory

Overwrite a memory's text. Use this when information becomes outdated or was captured incorrectly.

```bash
mem0 update 70179fc6-d4bd-4b45-8a7d-ce3228d1b29f "Updated: preferences column is now JSONB with a GIN index"
```

### `mem0 delete` — Remove a memory

Delete a single memory:
```bash
mem0 delete 70179fc6-d4bd-4b45-8a7d-ce3228d1b29f
```

Delete all memories for a user:
```bash
mem0 delete --all --user-id agent-123
```

### `mem0 entities list` — See all stored entities

List all users, agents, and apps that have memories.

```bash
mem0 entities list
```

### `mem0 entities delete` — Remove an entity

Delete an entity and all its associated memories.

```bash
mem0 entities delete agent-123 --type user
```

Entity types: `user`, `agent`, `app`, `run`.

### `mem0 auth login` / `logout` / `status`

Manage authentication credentials stored at `~/.mem0/config.json`.

```bash
mem0 auth login     # Paste API key
mem0 auth status    # Show current auth
mem0 auth logout    # Remove credentials
```

## Output format

The CLI outputs **table** format in a terminal and **JSON** when piped or redirected. Override with `--output`:

```bash
# Force JSON (useful for parsing in scripts)
mem0 search "deployment" --user-id agent-123 --output json

# Force table (useful for readability when piped)
mem0 list --user-id agent-123 --output table
```

### Parsing JSON output

```bash
# Get just the memory text from search results
mem0 search "auth" --user-id agent-123 -o json | jq '.[].memory'

# Get the ID of the first result
mem0 search "auth" --user-id agent-123 -o json | jq -r '.[0].id'

# Count memories
mem0 list --user-id agent-123 -o json | jq 'length'
```

## Global flags

Every command accepts these flags:

| Flag | Description |
|------|-------------|
| `--api-key <key>` | Override stored API key |
| `--base-url <url>` | Override API base URL (default: `https://api.mem0.ai`) |
| `--user-id <id>` | Scope operations to a specific user |
| `-o, --output <format>` | Force `json` or `table` output |

**API key precedence:** `--api-key` flag > `MEM0_API_KEY` env var > `~/.mem0/config.json`

## Patterns for agents

### Retrieve-then-act

Before making a decision or starting implementation, search for relevant memories:

```bash
# Before choosing a state management approach
mem0 search "state management" --user-id agent-123 -o json

# Before writing tests
mem0 search "testing patterns" --user-id agent-123 -o json

# Before deploying
mem0 search "deployment gotchas" --user-id agent-123 -o json
```

### Store after confirming

After a decision is validated or a fix works, store it:

```bash
# After fixing a bug
mem0 add "Fix for ENOENT on startup: the data/ dir must exist before the app initializes the SQLite connection" --user-id agent-123

# After a user confirms a preference
mem0 add "User wants all imports sorted with external packages first, then internal, then relative" --user-id agent-123

# After an architecture decision
mem0 add "Chose tRPC over REST for internal API — type safety across the monorepo is the priority" --user-id agent-123
```

### Automation and CI

```bash
# Store build metadata
mem0 add "Build v2.3.1 deployed to production at $(date -u +%Y-%m-%dT%H:%M:%SZ)" --user-id ci-bot

# Check for known issues before deploying
mem0 search "production issues" --user-id ci-bot -o json | jq '.[].memory'
```

## What to store vs. what not to store

**Good memories:**
- Decisions and their rationale
- Bug fixes and root causes
- User preferences and coding conventions
- Architecture patterns specific to a project
- Gotchas, edge cases, and "things that aren't obvious from the code"

**Not worth storing:**
- Things easily derived from code or git history
- Temporary task state (use a todo list instead)
- Large code snippets (the code itself is the source of truth)
- Information that changes frequently (store the pattern, not the value)
