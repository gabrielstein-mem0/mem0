# mem0 CLI

Command-line interface for the [Mem0 Platform](https://mem0.ai). Store, search, and manage memories from your terminal.

## Install

### Homebrew

```bash
brew install mem0ai/tap/mem0
```

### npm

```bash
npm install -g @mem0ai/cli
```

### From source

```bash
git clone https://github.com/mem0ai/mem0.git
cd mem0/mem0-cli
make build
# binary is at ./mem0
```

## Setup

1. Get an API key from [app.mem0.ai/dashboard/api-keys](https://app.mem0.ai/dashboard/api-keys)

2. Authenticate:

```bash
mem0 auth login
# paste your m0-... key when prompted
```

Or set the environment variable:

```bash
export MEM0_API_KEY="m0-your-api-key"
```

3. Verify:

```bash
mem0 auth status
```

## Usage

```bash
# Add a memory
mem0 add "User prefers dark mode" --user-id user-123

# Search memories
mem0 search "dark mode" --user-id user-123

# List all memories
mem0 list --user-id user-123

# Get a specific memory
mem0 get <memory-id>

# Update a memory
mem0 update <memory-id> "User prefers light mode"

# Delete a memory
mem0 delete <memory-id>

# Delete all memories for a user
mem0 delete --all --user-id user-123

# List entities
mem0 entities list

# Delete an entity
mem0 entities delete <entity-id>
```

## Output format

Output is **table** in a terminal and **JSON** when piped. Override with `--output`:

```bash
mem0 list --output json
mem0 search "query" --output table
```

## Global flags

| Flag | Description |
|------|-------------|
| `--api-key` | API key (overrides `MEM0_API_KEY` env and config) |
| `--base-url` | API base URL (default: `https://api.mem0.ai`) |
| `--user-id` | User ID for memory operations |
| `-o, --output` | Output format: `json` or `table` |

API key precedence: `--api-key` flag > `MEM0_API_KEY` env var > `~/.mem0/config.json`

## Development

```bash
make build    # build binary
make lint     # gofmt + go vet + golangci-lint
make test     # run tests
make clean    # remove binary
```

## License

Apache-2.0
