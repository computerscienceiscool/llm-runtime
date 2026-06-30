# Search Workflow: Enable and Run Semantic Search

This example shows how to enable search, index the repo, and run a query.

## Prerequisites
- Docker available
- Ollama installed (optional but required for `nomic-embed-text`)

## Steps

1) Enable search in config (set `commands.search.enabled: true`, model and URL as needed):
```bash
cat <<'YAML' > llm-runtime.config.yaml
commands:
  search:
    enabled: true
    embedding_model: "nomic-embed-text"
    ollama_url: "http://localhost:11434"  # if using Ollama
YAML
```

2) Build (optional if already built):
```bash
make build
```

3) Reindex the repo:
```bash
./llm-runtime --reindex
```

4) Run a search query:
```bash
echo "Find authentication code <search auth login>" | ./llm-runtime --root /path/to/your/repo
```

Notes:
- Use `embedding_model: "nomic-embed-text"` and ensure `ollama pull nomic-embed-text` if you prefer Ollama embeddings.
- Keep `max_file_size` reasonable to avoid large-file indexing overhead.
