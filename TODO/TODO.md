# TODO Index

Maintain zero-padded IDs starting at 001 and do not renumber. Keep only this index file in the TODO folder. Reference items by number in commits/PRs. Move completed items to the DONE section (checked) instead of deleting them. List any sibling TODO files below.

## High Priority
- [ ] 014 - Define approach for shell exec command injection protections (decide security modes and enforcement for exec whitelist vs shell flexibility; document outcome).
- [ ] 018 - Speed up test suite (cache modules, reduce Docker-dependent cases, add fast paths/flags).
- [ ] 060 - Set up Docker access for test environment. Add current user to `docker` group (`sudo usermod -aG docker $USER`) or configure rootless Docker so Docker-dependent tests in `cmd/llm-runtime`, `pkg/app`, `pkg/evaluator`, and `pkg/sandbox` can actually run.
## Medium Priority
- [ ] 006 - Document the config persona system when implemented.
- [ ] 015 - Implement MCP integration (Model Context Protocol) for standardized LLM tool integration; document usage.
- [ ] 016 - Add CLI project detection (`llm-runtime --auto`) to suggest configs based on repo type.

## Low Priority
- [ ] 009 - Add architecture diagrams as images in documentation.
- [ ] 010 - Implement streaming output for large command results.
- [ ] 017 - Add additional commands: `<git status>`, `<git diff>`, `<tree>`, `<grep pattern>` for richer repo introspection.
- [ ] 036 - Add scanner timeout / context support. `Scanner.Scan()` blocks indefinitely on `ReadString('\n')` with no way to cancel.
- [ ] 037 - Scanner processes input byte-by-byte (`scanner.go:108`), which may split multi-byte UTF-8 characters. Consider rune-based iteration for correctness with non-ASCII content.
- [ ] 084 - Remove unused `ContainerID` field from `ExecutionResult` struct (`pkg/scanner/types.go`). Field is set but never read by any caller.
- [ ] 085 - Remove unused `StartPos`, `EndPos`, `Original` fields from `Command` struct (`pkg/scanner/types.go`). Fields are never set or read.
- [ ] 086 - Remove unused `SearchResult.Relevance` field and `GetRelevanceLabel` function (`pkg/search/results.go`). Relevance is never set; callers use `Score` instead.
- [ ] 087 - Fix `filepath` variable shadowing imported `path/filepath` package in `database.go:112`. Rename local variable to avoid shadow.
- [ ] 088 - Fix wrong I/O timeout default in README (line 291) and docs (`configuration.md:44,277,417`). Docs say `60s`; actual default is `30s` (`DefaultIOTimeout` in `constants.go:13`).
- [ ] 089 - Fix wrong container pool defaults in README (lines 27-30, 614-623) and `container-pooling.md`. Docs say `size: 5`, `startup_containers: 2`, `health_check_interval: 60s`; actual defaults are `10`, `3`, `30s` (`constants.go:31-35`).
- [ ] 090 - Fix wrong default exec whitelist in README (lines 296-328). Lists 20+ commands; actual default is `["go test", "go build", "npm test", "make"]` (`defaults.go:48`).
- [ ] 091 - Clarify `python-go` image references across README and 14 doc files. No Dockerfile or build target exists for this image. It is an undefined external image presented as if part of the project.
- [ ] 092 - Fix broken references in README: `docs/installation.md` should be `docs/installation-guide.md` (line 754); `./security_test.sh` and `./write_demo.sh` should be `scripts/` (lines 553, 569); `examples/` directory does not exist (line 876); `internal/core/` and `docs/.index/` are empty dirs in project structure (lines 88-93); duplicate numbering "5." in Contributing (lines 863-864).
- [ ] 093 - Remove stale config fields from docs. `chunk_size`/`chunk_overlap` in `quick-reference.md:147-148` and `llm-runtime-overview.md:147-148` (removed in 075). `ollama_timeout` in `troubleshooting.md:307`, `quick-reference.md:146`, `llm-runtime-overview.md:146` (never existed in code).
- [ ] 094 - Delete `docs/llm-runtime-overview.md`. It is a duplicate of `docs/quick-reference.md`. Keep `quick-reference.md`.
- [ ] 095 - Fix exec mount read-only claim in README (line 279). Single-use exec containers are read-only, but pooled containers are read-write. README does not distinguish.
- [ ] 096 - Fix stale package paths in `CLAUDE.md` (lines 133-139). Shows `internal/app`, `internal/cli`, `internal/config`, `internal/search`, `internal/session` but all live under `pkg/`. Also stale refs at lines 202, 348-349 (`internal/config/defaults.go` should be `pkg/config/defaults.go`).
- [ ] 097 - Fix wrong subcommand syntax across README and 10+ doc files. Docs show `./llm-runtime --reindex` but CLI uses subcommands: `./llm-runtime reindex`, `search-validate`, `search-status`, `search-cleanup`, `search-update`, `check-ollama`.
- [ ] 098 - Fix stale default whitelist in `docs/SYSTEM_PROMPT.md` (lines 56-62). Lists 20+ commands across Go/Node/Python/Rust/system; actual default is `["go test", "go build", "npm test", "make"]`.
- [ ] 099 - Update `docs/index.md` after deleting `llm-runtime-overview.md` (094). Remove references and redirect to `quick-reference.md`.
- [ ] 100 - Clean up `docs/TODO.md`. Container image validation (line 51) marked pending but was completed (007). Several items duplicate main `TODO/TODO.md`.

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
- [x] 001 - Closed as stale. `temp_tests/` directory no longer exists.
- [x] 002 - Closed as stale. `PythonPath` field was removed; no references remain in codebase.
- [x] 003 - Closed as stale. All 22 scanner tests pass; mid-line matching was already resolved.
- [x] 028 - Add `Err()` method to Scanner (bufio.Scanner pattern). Non-EOF read errors are now stored and retrievable after `Scan()` returns nil. Caller in `app.go` updated to check and log.
- [x] 047 - Fix race condition in pool `Return()`. Hold write lock across closed check and channel send so `Close()` cannot close the channel between them. Removed dead `ctx.Done()` select case.
- [x] 033 - Surface scanner buffer overflow errors. `checkBufferLimit()` failures in `StateWriteBody`/`StateExecBody` now set `lastErr` with a `BUFFER_OVERFLOW` message instead of silently discarding.
- [x] 038 - Add concurrent audit log test. 10 goroutines x 20 writes with `-race` flag verifies no interleaving or corruption.
- [x] 051 - Remove misleading `ExecNetworkEnabled` config flag. Network is always disabled (`NetworkMode: "none"`) by design. Removed field from config struct, CLI flag, and config file. Documented in config docs.
- [x] 004 - Closed as stale. App test failures are Docker permission errors, not nil pointer issues.
- [x] 053 - Add I/O timeout context to `ExecuteOpen` and `ExecuteWrite`. Both now use `context.WithTimeout(ctx, cfg.IOTimeout)` instead of `context.Background()`.
- [x] 054 - Fix `--io-timeout` Cobra flag default from `"60s"` to `"30s"` in `cli/root.go` to match `DefaultIOTimeout` (30s).
- [x] 055 - Check `ContainerRemove` error in `pool.go`. Cleanup failure now surfaces in the returned error message.
- [x] 056 - Remove unused `pool` parameter from `ExecuteExec` and `ExecuteSearch`. Updated signatures, callers in `executor.go`, and all test call sites.
- [x] 057 - Remove 7 unused constants from `constants.go`: `DefaultContainerCPUs`, `AuditLogMaxSize`, `AuditLogMaxBackups`, `AuditLogMaxAge`, `MaxSessionsPerUser`, `MaxBackups`, `BackupExtension`.
- [x] 058 - Hoist compiled regexes in `errors.go` to package-level `var` block. `sanitizePaths` and `sanitizeUserInfo` no longer recompile on every call.
- [x] 059 - Add `t.Skip` guards to all Docker-dependent tests. Added `dockerAvailable()` helpers and skip checks to 40 tests across 7 files. `go test ./...` now passes clean without Docker.
- [x] 061 - Remove dead `content []byte` variable in `evaluator/open.go`. Eliminated pointless `string->[]byte->string` round-trip; use `contentStr` directly.
- [x] 062 - Remove unused error return from `FormatContent` and dead error check in `ExecuteWrite` (`evaluator/write.go`). Changed signature to return `string` only. Removed 10-line dead error block. Updated all test callers.
- [x] 063 - Fix import ordering in `evaluator/write.go`. Moved `"context"` to correct alphabetical position; consolidated third-party import group.
- [x] 064 - Check `hijackedResp.CloseWrite()` error in `sandbox/container.go`. Returns error instead of silently dropping it.
- [x] 065 - Remove unused `repoRoot` parameter from `executeInExistingContainer` (`sandbox/container.go`). Updated caller in `ExecuteInPooledContainer`.
- [x] 066 - Consolidate duplicate `parseMemoryLimitIO` into `parseMemoryLimit` (`sandbox/io_container.go`). Removed duplicate function, updated callers and tests.
- [x] 067 - Remove unused `StateExecute` constant from `scanner/scanner.go`. Removed from enum, `String()` method, and test.
- [x] 068 - Remove unused `showPrompts` field from Scanner struct (`scanner/scanner.go`). Removed field and parameter from `NewScanner`. Updated caller in `app.go` and all test call sites.
- [x] 069 - Log `destroyContainer` error in `healthCheckLoop` (`sandbox/pool.go`). Error now logged to stderr instead of silently dropped.
- [x] 070 - Add `maxLogPayloadSize` (10MB) guard in `demuxLogs` and `readDockerLogs` (`sandbox/container.go`, `sandbox/io_container.go`). Rejects corrupted Docker stream headers before allocation.
- [x] 071 - Fix nil pointer dereference in `ValidateIndex` (`search/indexing.go`). Added `else if err != nil` guard after `os.IsNotExist` check so non-existent errors (e.g. permission denied) don't nil-dereference `info.ModTime()`.
- [x] 072 - Add missing `rows.Err()` check after iteration loop in `ValidateIndex` (`search/indexing.go`). Database iteration errors no longer silently lost.
- [x] 073 - Remove unused `storedHash` variable and `content_hash` from SQL query in `ValidateIndex` (`search/indexing.go`). Hash validation was never implemented despite the comment.
- [x] 074 - Remove unused `Session.CommandsRun` field (`session/session.go`). Command counting lives in `Executor.commandsRun`. Removed field and associated tests.
- [x] 075 - Remove unused `SearchConfig.ChunkSize` field. Removed from `search/config.go`, `config/defaults.go` (default, viper, loader), `config/defaults_test.go`, and `DefaultEmbeddingDims` constant.
- [x] 076 - Remove unused `SearchConfig.EmbeddingDimensions` field. Removed from `search/config.go`, `config/defaults.go` (default, viper, loader), `config/defaults_test.go`. Dimension remains hardcoded as `const embeddingDimensions = 768` in `similarity.go`.
- [x] 077 - Check `session.Close()` error in `app.Close()` (`app/app.go`). Now collects both session and pool close errors using `errors.Join`.
- [x] 078 - Remove stale `.backup` files tracked in git (`pkg/config/defaults.go.backup`, `defaults_test.go.backup`, `types.go.backup`).
- [x] 079 - Fix pool `healthCheckLoop` deadlock on shutdown (`sandbox/pool.go`). Added `done chan struct{}` to `ContainerPool`. `Close()` closes the channel before `ticker.Stop()`. `healthCheckLoop` select now has `case <-p.done: return` so the goroutine exits promptly.
- [x] 080 - Remove unused `verbose` parameter from `PullDockerImage` (`sandbox/client.go`). Removed parameter, fixed stale comment, updated callers in `pool.go`, `exec.go`, and all test call sites. Removed `TestPullDockerImage_VerboseMode` test.
- [x] 081 - Remove duplicate `checkOllamaAvailability` from `cli/commands.go`. Had `runCheckOllama` call `search.CheckOllamaSetup` instead. Fixed `CheckOllamaSetup` to use `http.StatusOK` instead of hardcoded `200`. Removed `net/http` import from CLI. Removed redundant CLI tests.
- [x] 082 - Remove unused `DefaultSessionTimeout` constant from `config/constants.go`. Never referenced anywhere.
- [x] 083 - Check `io.ReadAll` error in `search/embedding.go`. Now returns a fallback error message if body read fails instead of silently using empty body.

## Other TODO Files
- docs/TODO.md
