# Mercury — State

## Phase

MVP complete. All Milestone 1 + 2 commands implemented and working.

## What's built

- Go CLI binary, builds via `nix build`, on PATH via nix profile
- SQLite schema: `messages`, `subscriptions`, `cursors` tables
- DB auto-created at `~/.local/share/mercury/mercury.db`
- Commands: `send`, `read`, `subscribe`, `unsubscribe`, `channels`, `log`
- Cursor tracking for per-agent unread position
- WAL mode + busy_timeout=5000 for concurrent agent access
- Input validation: empty `--as`, `--to`, `--channel` rejected
- Hint when reading with no subscriptions
- 7 core db tests (send/read/subscribe/unsubscribe/cursors/channels)

## Decisions made

- Named channels, created on-the-fly
- CLI tool in Go, single binary
- SQLite at ~/.local/share/mercury/mercury.db
- Role-based agent names (self-chosen, unenforced)
- Messages are opaque text, persist as log
- Nix flake for builds and dev shell
- Maximally flexible — convention evolves in practice, not in code

## Milestone status

### Milestone 1: Core (MVP) — DONE
1. ~~Scaffold Go project with Nix flake~~
2. ~~SQLite schema + migration on first run~~
3. ~~`mercury send --as NAME --to CHANNEL BODY`~~
4. ~~`mercury read --as NAME`~~
5. ~~`mercury subscribe --as NAME --channel CHANNEL`~~
6. ~~`mercury channels`~~

### Milestone 2: Usability — DONE
7. ~~`mercury read --follow`~~ (poll-based)
8. ~~`mercury log`~~
9. ~~`mercury unsubscribe`~~
10. ~~Cursor tracking~~

### Milestone 3: Integration — IN PROGRESS
11. ~~Mercury skill~~ — complete, polished with smoke test feedback, symlinked to ~/.claude/skills/mercury
12. ~~Smoke test~~ — oracle + keeper:mercury + worker:smoke-test exchanged messages on the bus
13. Document conventions that emerged during dogfooding — ongoing

## Current work

Settling. v0.1.0 shipped, foundation hardened. Bus is live and carrying real traffic.

## Blockers

None.

## Next actions

- Continue dogfooding — conventions emerge from use
- Consider `mercury who` if discovery need proves real
- Tag v0.2.0 when next batch of improvements lands
