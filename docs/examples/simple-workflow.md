# Simple Workflow: Read, Write, and Run Tests

This example shows a minimal end-to-end session: read a file, write a summary, and run tests. Adjust commands for your project.

## Prerequisites
- Docker available
- Optional: exec container image (default `ubuntu:22.04`; set to `python-go` if you need Go/Python)

## Steps

1) Read a file:
```
echo "Show README <open README.md>" | ./llm-runtime --root /path/to/your/repo
```

2) Write a new file:
```
echo "Create summary <write SUMMARY.md>This is a quick summary.</write>" | ./llm-runtime --root /path/to/your/repo
```

3) Run tests (enable exec and whitelist):
```
./llm-runtime --root /path/to/your/repo \
  --exec-enabled \
  --exec-whitelist "go test,make test" \
  --exec-image ubuntu:22.04 \
  --input <(echo "<exec go test ./...>")
```

Notes:
- Use `--exec-image python-go` if you need Go/Python preinstalled.
- Keep the exec whitelist tight for your workflow.
