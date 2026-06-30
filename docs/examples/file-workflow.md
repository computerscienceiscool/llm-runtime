# File Workflow: Read and Write Files Safely

This example shows reading a file and writing a summary using the I/O container.

## Prerequisites
- Docker available (for containerized read/write)
- Optional: set `commands.write.backup_before_write: true` to keep backups

## Steps

1) Read a file:
```bash
echo "Show README <open README.md>" | ./llm-runtime --root /path/to/your/repo
```

2) Write a summary file:
```bash
echo "<write SUMMARY.md>Quick summary of README.</write>" | ./llm-runtime --root /path/to/your/repo
```

3) Use a config file for write options (optional):
```bash
cat <<'YAML' > llm-runtime.config.yaml
commands:
  write:
    enabled: true
    backup_before_write: true
    allowed_extensions:
      - ".md"
      - ".txt"
YAML
```

Notes:
- The I/O container reads/writes with the repo mounted read-only except for temp writes, providing isolation.
- Keep `max_file_size` and `max_write_size` within defaults for safety.
