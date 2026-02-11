# TODO Index

Maintain zero-padded IDs starting at 001 and do not renumber. Keep only this index file in the TODO folder. Reference items by number in commits/PRs. Move completed items to the DONE section (checked) instead of deleting them. List any sibling TODO files below.

## High Priority
- [ ] 001 - Fix test files in `temp_tests/` for new import paths.
- [ ] 002 - Update tests referencing the removed `PythonPath` field.
- [ ] 003 - Fix parser tests expecting mid-line command matching.
- [ ] 004 - Resolve app test nil pointer issues.
- [ ] 014 - Define approach for shell exec command injection protections (decide security modes and enforcement for exec whitelist vs shell flexibility; document outcome).
- [ ] 018 - Speed up test suite (cache modules, reduce Docker-dependent cases, add fast paths/flags).
- [ ] 028 - Fix scanner silently ignoring non-EOF read errors. `scanner.go:99-105` treats all errors the same as EOF when there is remaining line data -- disk errors, permission errors, etc. are swallowed.

## Medium Priority
- [ ] 006 - Document the config persona system when implemented.
- [ ] 015 - Implement MCP integration (Model Context Protocol) for standardized LLM tool integration; document usage.
- [ ] 016 - Add CLI project detection (`llm-runtime --auto`) to suggest configs based on repo type.
- [ ] 033 - Scanner buffer overflow silently aborts commands. When `checkBufferLimit()` fails in `StateWriteBody`/`StateExecBody`, the command is discarded and the scanner moves on with no error surfaced to the caller.

## Low Priority
- [ ] 009 - Add architecture diagrams as images in documentation.
- [ ] 010 - Implement streaming output for large command results.
- [ ] 017 - Add additional commands: `<git status>`, `<git diff>`, `<tree>`, `<grep pattern>` for richer repo introspection.
- [ ] 036 - Add scanner timeout / context support. `Scanner.Scan()` blocks indefinitely on `ReadString('\n')` with no way to cancel.
- [ ] 037 - Scanner processes input byte-by-byte (`scanner.go:108`), which may split multi-byte UTF-8 characters. Consider rune-based iteration for correctness with non-ASCII content.
- [ ] 038 - Add concurrent audit log tests. `session.LogAudit` has no synchronization; multiple goroutines writing to the same logger can interleave entries.
- [ ] 047 - Fix race condition in container pool `Return()` method (`sandbox/pool.go:176-228`). Check-then-act on `p.closed` without holding the lock; pool can close between check and container return.
- [ ] 051 - Wire `ExecNetworkEnabled` config flag to container creation or document that network is always disabled. Currently all container code hardcodes `NetworkMode: "none"`. Config default corrected to `false` in `llm-runtime.config.yaml`.

## DONE
- [x] 024 - Fix container pool not assigned to App struct in bootstrap.go (pool created but never stored; leaked containers on shutdown).
- [x] 025 - Fix audit log file descriptor never closed in session.go (added `auditFile` field, `Close()` method, and wired into `App.Close()`).
- [x] 026 - Fix missing `rows.Err()` check after iteration in `search/engine.go`.
- [x] 027 - Fix unchecked `binary.Write`/`binary.Read` errors in `search/similarity.go` (`serializeEmbedding` now returns error).
- [x] 029 - Fix `removeFileInfo` errors silently dropped in `search/indexing.go`.
- [x] 030 - Fix hardcoded audit log path in session.go (now reads from `config.AuditLogPath`).
- [x] 032 - Add directory detection in `evaluator/open.go`. `ExecuteOpen` now checks `IsDir()` and returns `IS_DIRECTORY` error. Unskipped two tests.
- [x] 034 - Remove dead code: `fullConfig` struct, `setFullConfigDefaults()`, and their tests from `config/defaults.go`, `config/types.go`, `config/defaults_test.go`.
- [x] 035 - Remove leftover debug comment in `cli/config.go`.
- [x] 039 - Update `Dockerfile.io` base image from `golang:1.22.2-alpine` to `alpine:3.21` (Go toolchain not needed for I/O container).
- [x] 040 - Consolidate double `init()` in `cli/root.go` into a single function. Viper defaults and config file setup now run before Cobra flag registration.
- [x] 041 - Fix `InitializeSearchIndex` never triggering: `commands.go:117` compared `int64` to string `"0"` (always false). Changed to `.(int64) == 0`.
- [x] 042 - Replace fake content hash in `indexing.go:185`. Changed `fmt.Sprintf("%x", content)` to `fmt.Sprintf("%x", sha256.Sum256(content))`.
- [x] 049 - Fix wrong build path in 5 scripts (`demo.sh`, `exec_demo.sh`, `write_demo.sh`, `example_usage.sh`, `security_test.sh`). Changed `go build -o llm-runtime main.go` to `go build -o llm-runtime ./cmd/llm-runtime`.
- [x] 052 - Add missing Viper default for `io-timeout` in `config/defaults.go`. Without it, `viper.GetString("io-timeout")` returns empty string when no flag or config file value is set, causing `time.ParseDuration` to fail.
- [x] 012 - Align exec defaults in docs/README with code (docs claimed exec always enabled and default image `python-go`; defaults disable exec and use `ubuntu:22.04`).
- [x] 013 - Align search docs with code defaults (docs called for `nomic-embed-text` as default; defaults use `all-MiniLM-L6-v2`).
- [x] 011 - Enforce `commands.exec.enabled` flag (Config lacked enable field and `ExecuteExec` ran regardless of config; added plumbing + guard + tests).
- [x] 019 - Document exec container image validation rules and expected errors in docs.
- [x] 020 - Add an example workflow to docs/examples/ (e.g., read/write and run tests).
- [x] 021 - Add a stub `--auto` detection flag that reports planned project-type suggestions.
- [x] 022 - Improve `make help` output to list targets clearly and make `make` show the help menu.
- [x] 023 - Document `make test-fast` usage in README/docs.
- [x] 007 - Decide and document the container image validation approach (whitelist vs digest pinning vs patterns).
- [x] 043 - Escape pipe delimiters in audit log fields (`sandbox/audit.go`). Added escape function to replace `|` with `\|` in all variable fields. Updated tests to verify escaping.
- [x] 045 - Validate memory limit parsing in `sandbox/container.go` and `sandbox/io_container.go`. Both `parseMemoryLimit()` and `parseMemoryLimitIO()` now return `(int64, error)` and reject invalid formats. Updated all callers and tests.
- [x] 046 - Validate path length against `MaxPathLength` constant in `sandbox/path.go`. Added check at top of `ValidatePath()`. Added tests.
- [x] 031 - Add binary file detection in `evaluator/open.go`. Sniff first 512 bytes with `http.DetectContentType`; reject non-text content types with `BINARY_FILE` error. Unskipped test.
- [x] 044 - Add timeout to container cleanup in `pool.Close()`. Changed `context.Background()` to 30-second timeout so pool shutdown doesn't hang on unresponsive Docker.
- [x] 048 - Add timeout to health check loop context in `pool.go`. Each health check tick now uses a 10-second timeout context instead of unbounded `context.Background()`.
- [x] 050 - Remove broken `--io-containerized` flag from Makefile `test-io-container` target. The flag was never implemented in the CLI.

## Other TODO Files
- docs/TODO.md
