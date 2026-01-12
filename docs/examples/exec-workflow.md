# Exec Workflow: Enable Exec and Run Tests

This example shows how to enable exec, whitelist commands, and run tests inside the exec container.

## Prerequisites
- Docker available
- Exec image available (default `ubuntu:22.04`; use `python-go` if you need Go/Python preinstalled)

## Steps

1) Create a config that enables exec and sets a whitelist:
```bash
cat <<'YAML' > llm-runtime.config.yaml
commands:
  exec:
    enabled: true
    container_image: "ubuntu:22.04"
    timeout_seconds: 30
    whitelist:
      - "go test"
      - "make test"
YAML
```

2) Run a test command via exec:
```bash
echo "<exec go test ./...>" | ./llm-runtime --root /path/to/your/repo
```

3) For multiple commands, feed an input file:
```bash
cat <<'EOF' > exec_input.txt
<exec go test ./...>
<exec make test>
EOF

./llm-runtime --root /path/to/your/repo --exec-enabled --exec-whitelist "go test,make test" --input exec_input.txt
```

Notes:
- Keep the whitelist tight for your workflow.
- Swap `container_image` to `python-go` if you need Go/Python tools preinstalled.
