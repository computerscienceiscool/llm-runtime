# TODO Index

Maintain zero-padded IDs starting at 001 and do not renumber. Keep only this index file in the TODO folder. Reference items by number in commits/PRs. Move completed items to the DONE section (checked) instead of deleting them. List any sibling TODO files below.

## High Priority
- [ ] 001 - Fix test files in `temp_tests/` for new import paths.
- [ ] 002 - Update tests referencing the removed `PythonPath` field.
- [ ] 003 - Fix parser tests expecting mid-line command matching.
- [ ] 004 - Resolve app test nil pointer issues.
- [ ] 011 - Enforce `commands.exec.enabled` flag (Config lacks an exec-enabled field and `ExecuteExec` runs regardless of config, so exec cannot be disabled; add config plumbing + guard + tests).
- [ ] 014 - Define approach for shell exec command injection protections (decide security modes and enforcement for exec whitelist vs shell flexibility; document outcome).

## Medium Priority
- [ ] 005 - Add troubleshooting for common Ollama issues.
- [ ] 006 - Document the config persona system when implemented.
- [ ] 007 - Decide and document the container image validation approach (whitelist vs digest pinning vs patterns).
- [ ] 012 - Align exec defaults in docs/README with code (docs claim exec always enabled and default image `python-go`, but `pkg/config/defaults.go` disables exec by default and uses `ubuntu:22.04`; update docs or defaults to match).
- [ ] 013 - Align search docs with code defaults (docs call for Ollama `nomic-embed-text` as default model, while `pkg/config/defaults.go` sets `EmbeddingModel: all-MiniLM-L6-v2`; decide canonical default and update docs/config).
- [ ] 015 - Implement MCP integration (Model Context Protocol) for standardized LLM tool integration; document usage.
- [ ] 016 - Add CLI project detection (`llm-runtime --auto`) to suggest configs based on repo type.

## Low Priority
- [ ] 008 - Add example workflows in `docs/examples/`.
- [ ] 009 - Add architecture diagrams as images in documentation.
- [ ] 010 - Implement streaming output for large command results.
- [ ] 017 - Add additional commands: `<git status>`, `<git diff>`, `<tree>`, `<grep pattern>` for richer repo introspection.

## DONE
- [ ] (move completed items here and mark with [x])

## Other TODO Files
- docs/TODO.md
