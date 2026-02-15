# TODO Index

Maintain zero-padded IDs prefixed with the developer's initial (e.g., J001 for JJ, S001 for Steve). Do not renumber. Keep only this index file in the TODO folder. Reference items by number in commits/PRs. Move completed items to the DONE section (checked) instead of deleting them. List any sibling TODO files below.

## High Priority
- [ ] J014 -Define approach for shell exec command injection protections (decide security modes and enforcement for exec whitelist vs shell flexibility; document outcome).
- [ ] J018 -Speed up test suite (cache modules, reduce Docker-dependent cases, add fast paths/flags).
- [ ] J060 -Set up Docker access for test environment. Add current user to `docker` group (`sudo usermod -aG docker $USER`) or configure rootless Docker so Docker-dependent tests in `cmd/llm-runtime`, `pkg/app`, `pkg/evaluator`, and `pkg/sandbox` can actually run.
## Medium Priority
- [ ] J006 -Document the config persona system when implemented.
- [ ] J015 -Implement MCP integration (Model Context Protocol) for standardized LLM tool integration; document usage.
- [ ] J016 -Add CLI project detection (`llm-runtime --auto`) to suggest configs based on repo type.

## Low Priority
- [ ] J009 -Add architecture diagrams as images in documentation.
- [ ] J010 -Implement streaming output for large command results.
- [ ] J017 -Add additional commands: `<git status>`, `<git diff>`, `<tree>`, `<grep pattern>` for richer repo introspection.
- [ ] J036 -Add scanner timeout / context support. `Scanner.Scan()` blocks indefinitely on `ReadString('\n')` with no way to cancel.
- [ ] J037 -Scanner processes input byte-by-byte (`scanner.go:108`), which may split multi-byte UTF-8 characters. Consider rune-based iteration for correctness with non-ASCII content.
- [ ] J084 -Remove unused `ContainerID` field from `ExecutionResult` struct (`pkg/scanner/types.go`). Field is set but never read by any caller.
- [ ] J085 -Remove unused `StartPos`, `EndPos`, `Original` fields from `Command` struct (`pkg/scanner/types.go`). Fields are never set or read.
- [ ] J086 -Remove unused `SearchResult.Relevance` field and `GetRelevanceLabel` function (`pkg/search/results.go`). Relevance is never set; callers use `Score` instead.
- [ ] J087 -Fix `filepath` variable shadowing imported `path/filepath` package in `database.go:112`. Rename local variable to avoid shadow.
- [ ] J088 -Fix wrong I/O timeout default in README (line 291) and docs (`configuration.md:44,277,417`). Docs say `60s`; actual default is `30s` (`DefaultIOTimeout` in `constants.go:13`).
- [ ] J089 -Fix wrong container pool defaults in README (lines 27-30, 614-623) and `container-pooling.md`. Docs say `size: 5`, `startup_containers: 2`, `health_check_interval: 60s`; actual defaults are `10`, `3`, `30s` (`constants.go:31-35`).
- [ ] J090 -Fix wrong default exec whitelist in README (lines 296-328). Lists 20+ commands; actual default is `["go test", "go build", "npm test", "make"]` (`defaults.go:48`).
- [ ] J091 -Clarify `python-go` image references across README and 14 doc files. No Dockerfile or build target exists for this image. It is an undefined external image presented as if part of the project.
- [ ] J092 -Fix broken references in README: `docs/installation.md` should be `docs/installation-guide.md` (line 754); `./security_test.sh` and `./write_demo.sh` should be `scripts/` (lines 553, 569); `examples/` directory does not exist (line 876); `internal/core/` and `docs/.index/` are empty dirs in project structure (lines 88-93); duplicate numbering "5." in Contributing (lines 863-864).
- [ ] J093 -Remove stale config fields from docs. `chunk_size`/`chunk_overlap` in `quick-reference.md:147-148` and `llm-runtime-overview.md:147-148` (removed in J075). `ollama_timeout` in `troubleshooting.md:307`, `quick-reference.md:146`, `llm-runtime-overview.md:146` (never existed in code).
- [ ] J094 -Delete `docs/llm-runtime-overview.md`. It is a duplicate of `docs/quick-reference.md`. Keep `quick-reference.md`.
- [ ] J095 -Fix exec mount read-only claim in README (line 279). Single-use exec containers are read-only, but pooled containers are read-write. README does not distinguish.
- [ ] J096 -Fix stale package paths in `CLAUDE.md` (lines 133-139). Shows `internal/app`, `internal/cli`, `internal/config`, `internal/search`, `internal/session` but all live under `pkg/`. Also stale refs at lines 202, 348-349 (`internal/config/defaults.go` should be `pkg/config/defaults.go`).
- [ ] J097 -Fix wrong subcommand syntax across README and 10+ doc files. Docs show `./llm-runtime --reindex` but CLI uses subcommands: `./llm-runtime reindex`, `search-validate`, `search-status`, `search-cleanup`, `search-update`, `check-ollama`.
- [ ] J098 -Fix stale default whitelist in `docs/SYSTEM_PROMPT.md` (lines 56-62). Lists 20+ commands across Go/Node/Python/Rust/system; actual default is `["go test", "go build", "npm test", "make"]`.
- [ ] J099 -Update `docs/index.md` after deleting `llm-runtime-overview.md` (J094). Remove references and redirect to `quick-reference.md`.
- [ ] J100 -Clean up `docs/TODO.md`. Container image validation (line 51) marked pending but was completed (J007). Several items duplicate main `TODO/TODO.md`.

## DONE
- [x] J024 -Fix container pool not assigned to App struct in bootstrap.go (pool created but never stored; leaked containers on shutdown).
- [x] J025 -Fix audit log file descriptor never closed in session.go (added `auditFile` field, `Close()` method, and wired into `App.Close()`).
- [x] J026 -Fix missing `rows.Err()` check after iteration in `search/engine.go`.
- [x] J027 -Fix unchecked `binary.Write`/`binary.Read` errors in `search/similarity.go` (`serializeEmbedding` now returns error).
- [x] J029 -Fix `removeFileInfo` errors silently dropped in `search/indexing.go`.
- [x] J030 -Fix hardcoded audit log path in session.go (now reads from `config.AuditLogPath`).
- [x] J032 -Add directory detection in `evaluator/open.go`. `ExecuteOpen` now checks `IsDir()` and returns `IS_DIRECTORY` error. Unskipped two tests.
- [x] J034 -Remove dead code: `fullConfig` struct, `setFullConfigDefaults()`, and their tests from `config/defaults.go`, `config/types.go`, `config/defaults_test.go`.
- [x] J035 -Remove leftover debug comment in `cli/config.go`.
- [x] J039 -Update `Dockerfile.io` base image from `golang:1.22.2-alpine` to `alpine:3.21` (Go toolchain not needed for I/O container).
- [x] J040 -Consolidate double `init()` in `cli/root.go` into a single function. Viper defaults and config file setup now run before Cobra flag registration.
- [x] J041 -Fix `InitializeSearchIndex` never triggering: `commands.go:117` compared `int64` to string `"0"` (always false). Changed to `.(int64) == 0`.
- [x] J042 -Replace fake content hash in `indexing.go:185`. Changed `fmt.Sprintf("%x", content)` to `fmt.Sprintf("%x", sha256.Sum256(content))`.
- [x] J049 -Fix wrong build path in 5 scripts (`demo.sh`, `exec_demo.sh`, `write_demo.sh`, `example_usage.sh`, `security_test.sh`). Changed `go build -o llm-runtime main.go` to `go build -o llm-runtime ./cmd/llm-runtime`.
- [x] J052 -Add missing Viper default for `io-timeout` in `config/defaults.go`. Without it, `viper.GetString("io-timeout")` returns empty string when no flag or config file value is set, causing `time.ParseDuration` to fail.
- [x] J012 -Align exec defaults in docs/README with code (docs claimed exec always enabled and default image `python-go`; defaults disable exec and use `ubuntu:22.04`).
- [x] J013 -Align search docs with code defaults (docs called for `nomic-embed-text` as default; defaults use `all-MiniLM-L6-v2`).
- [x] J011 -Enforce `commands.exec.enabled` flag (Config lacked enable field and `ExecuteExec` ran regardless of config; added plumbing + guard + tests).
- [x] J019 -Document exec container image validation rules and expected errors in docs.
- [x] J020 -Add an example workflow to docs/examples/ (e.g., read/write and run tests).
- [x] J021 -Add a stub `--auto` detection flag that reports planned project-type suggestions.
- [x] J022 -Improve `make help` output to list targets clearly and make `make` show the help menu.
- [x] J023 -Document `make test-fast` usage in README/docs.
- [x] J007 -Decide and document the container image validation approach (whitelist vs digest pinning vs patterns).
- [x] J043 -Escape pipe delimiters in audit log fields (`sandbox/audit.go`). Added escape function to replace `|` with `\|` in all variable fields. Updated tests to verify escaping.
- [x] J045 -Validate memory limit parsing in `sandbox/container.go` and `sandbox/io_container.go`. Both `parseMemoryLimit()` and `parseMemoryLimitIO()` now return `(int64, error)` and reject invalid formats. Updated all callers and tests.
- [x] J046 -Validate path length against `MaxPathLength` constant in `sandbox/path.go`. Added check at top of `ValidatePath()`. Added tests.
- [x] J031 -Add binary file detection in `evaluator/open.go`. Sniff first 512 bytes with `http.DetectContentType`; reject non-text content types with `BINARY_FILE` error. Unskipped test.
- [x] J044 -Add timeout to container cleanup in `pool.Close()`. Changed `context.Background()` to 30-second timeout so pool shutdown doesn't hang on unresponsive Docker.
- [x] J048 -Add timeout to health check loop context in `pool.go`. Each health check tick now uses a 10-second timeout context instead of unbounded `context.Background()`.
- [x] J050 -Remove broken `--io-containerized` flag from Makefile `test-io-container` target. The flag was never implemented in the CLI.
- [x] J001 -Closed as stale. `temp_tests/` directory no longer exists.
- [x] J002 -Closed as stale. `PythonPath` field was removed; no references remain in codebase.
- [x] J003 -Closed as stale. All 22 scanner tests pass; mid-line matching was already resolved.
- [x] J028 -Add `Err()` method to Scanner (bufio.Scanner pattern). Non-EOF read errors are now stored and retrievable after `Scan()` returns nil. Caller in `app.go` updated to check and log.
- [x] J047 -Fix race condition in pool `Return()`. Hold write lock across closed check and channel send so `Close()` cannot close the channel between them. Removed dead `ctx.Done()` select case.
- [x] J033 -Surface scanner buffer overflow errors. `checkBufferLimit()` failures in `StateWriteBody`/`StateExecBody` now set `lastErr` with a `BUFFER_OVERFLOW` message instead of silently discarding.
- [x] J038 -Add concurrent audit log test. 10 goroutines x 20 writes with `-race` flag verifies no interleaving or corruption.
- [x] J051 -Remove misleading `ExecNetworkEnabled` config flag. Network is always disabled (`NetworkMode: "none"`) by design. Removed field from config struct, CLI flag, and config file. Documented in config docs.
- [x] J004 -Closed as stale. App test failures are Docker permission errors, not nil pointer issues.
- [x] J053 -Add I/O timeout context to `ExecuteOpen` and `ExecuteWrite`. Both now use `context.WithTimeout(ctx, cfg.IOTimeout)` instead of `context.Background()`.
- [x] J054 -Fix `--io-timeout` Cobra flag default from `"60s"` to `"30s"` in `cli/root.go` to match `DefaultIOTimeout` (30s).
- [x] J055 -Check `ContainerRemove` error in `pool.go`. Cleanup failure now surfaces in the returned error message.
- [x] J056 -Remove unused `pool` parameter from `ExecuteExec` and `ExecuteSearch`. Updated signatures, callers in `executor.go`, and all test call sites.
- [x] J057 -Remove 7 unused constants from `constants.go`: `DefaultContainerCPUs`, `AuditLogMaxSize`, `AuditLogMaxBackups`, `AuditLogMaxAge`, `MaxSessionsPerUser`, `MaxBackups`, `BackupExtension`.
- [x] J058 -Hoist compiled regexes in `errors.go` to package-level `var` block. `sanitizePaths` and `sanitizeUserInfo` no longer recompile on every call.
- [x] J059 -Add `t.Skip` guards to all Docker-dependent tests. Added `dockerAvailable()` helpers and skip checks to 40 tests across 7 files. `go test ./...` now passes clean without Docker.
- [x] J061 -Remove dead `content []byte` variable in `evaluator/open.go`. Eliminated pointless `string->[]byte->string` round-trip; use `contentStr` directly.
- [x] J062 -Remove unused error return from `FormatContent` and dead error check in `ExecuteWrite` (`evaluator/write.go`). Changed signature to return `string` only. Removed 10-line dead error block. Updated all test callers.
- [x] J063 -Fix import ordering in `evaluator/write.go`. Moved `"context"` to correct alphabetical position; consolidated third-party import group.
- [x] J064 -Check `hijackedResp.CloseWrite()` error in `sandbox/container.go`. Returns error instead of silently dropping it.
- [x] J065 -Remove unused `repoRoot` parameter from `executeInExistingContainer` (`sandbox/container.go`). Updated caller in `ExecuteInPooledContainer`.
- [x] J066 -Consolidate duplicate `parseMemoryLimitIO` into `parseMemoryLimit` (`sandbox/io_container.go`). Removed duplicate function, updated callers and tests.
- [x] J067 -Remove unused `StateExecute` constant from `scanner/scanner.go`. Removed from enum, `String()` method, and test.
- [x] J068 -Remove unused `showPrompts` field from Scanner struct (`scanner/scanner.go`). Removed field and parameter from `NewScanner`. Updated caller in `app.go` and all test call sites.
- [x] J069 -Log `destroyContainer` error in `healthCheckLoop` (`sandbox/pool.go`). Error now logged to stderr instead of silently dropped.
- [x] J070 -Add `maxLogPayloadSize` (10MB) guard in `demuxLogs` and `readDockerLogs` (`sandbox/container.go`, `sandbox/io_container.go`). Rejects corrupted Docker stream headers before allocation.
- [x] J071 -Fix nil pointer dereference in `ValidateIndex` (`search/indexing.go`). Added `else if err != nil` guard after `os.IsNotExist` check so non-existent errors (e.g. permission denied) don't nil-dereference `info.ModTime()`.
- [x] J072 -Add missing `rows.Err()` check after iteration loop in `ValidateIndex` (`search/indexing.go`). Database iteration errors no longer silently lost.
- [x] J073 -Remove unused `storedHash` variable and `content_hash` from SQL query in `ValidateIndex` (`search/indexing.go`). Hash validation was never implemented despite the comment.
- [x] J074 -Remove unused `Session.CommandsRun` field (`session/session.go`). Command counting lives in `Executor.commandsRun`. Removed field and associated tests.
- [x] J075 -Remove unused `SearchConfig.ChunkSize` field. Removed from `search/config.go`, `config/defaults.go` (default, viper, loader), `config/defaults_test.go`, and `DefaultEmbeddingDims` constant.
- [x] J076 -Remove unused `SearchConfig.EmbeddingDimensions` field. Removed from `search/config.go`, `config/defaults.go` (default, viper, loader), `config/defaults_test.go`. Dimension remains hardcoded as `const embeddingDimensions = 768` in `similarity.go`.
- [x] J077 -Check `session.Close()` error in `app.Close()` (`app/app.go`). Now collects both session and pool close errors using `errors.Join`.
- [x] J078 -Remove stale `.backup` files tracked in git (`pkg/config/defaults.go.backup`, `defaults_test.go.backup`, `types.go.backup`).
- [x] J079 -Fix pool `healthCheckLoop` deadlock on shutdown (`sandbox/pool.go`). Added `done chan struct{}` to `ContainerPool`. `Close()` closes the channel before `ticker.Stop()`. `healthCheckLoop` select now has `case <-p.done: return` so the goroutine exits promptly.
- [x] J080 -Remove unused `verbose` parameter from `PullDockerImage` (`sandbox/client.go`). Removed parameter, fixed stale comment, updated callers in `pool.go`, `exec.go`, and all test call sites. Removed `TestPullDockerImage_VerboseMode` test.
- [x] J081 -Remove duplicate `checkOllamaAvailability` from `cli/commands.go`. Had `runCheckOllama` call `search.CheckOllamaSetup` instead. Fixed `CheckOllamaSetup` to use `http.StatusOK` instead of hardcoded `200`. Removed `net/http` import from CLI. Removed redundant CLI tests.
- [x] J082 -Remove unused `DefaultSessionTimeout` constant from `config/constants.go`. Never referenced anywhere.
- [x] J083 -Check `io.ReadAll` error in `search/embedding.go`. Now returns a fallback error message if body read fails instead of silently using empty body.

## Other TODO Files
- docs/TODO.md
