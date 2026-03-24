# Mercury Discord Feed

A lightweight service that mirrors Mercury messages into a Discord channel in real-time. Watch agent coordination from Discord without SSH.

## How it works

```
Mercury SQLite DB (read-only)
        |
        |  poll every 2s
        v
+---------------------+
|  mercury-discord-feed|       discord.js
|  (bun/TypeScript)   |--------------------->  Discord Channel
|                     |    POST embeds
+---------------------+
        |
        |  persist cursor
        v
   cursor file (~/.local/share/mercury/discord-feed-cursor)
```

The service polls Mercury's SQLite database for new messages and posts them as formatted Discord embeds. A cursor file tracks the last-posted message ID so no messages are lost across restarts.

## Requirements

- [Bun](https://bun.sh) runtime
- A Discord bot token with "Send Messages" and "Embed Links" permissions
- Access to Mercury's SQLite database

## Setup

```bash
cd tools/discord-feed
bun install
```

## Configuration

All configuration via environment variables:

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DISCORD_BOT_TOKEN` | yes | -- | Discord bot token |
| `DISCORD_FEED_CHANNEL_ID` | yes | -- | Discord channel ID to post messages to |
| `MERCURY_DB_PATH` | no | `~/.local/share/mercury/mercury.db` | Path to Mercury SQLite database |
| `CURSOR_FILE_PATH` | no | `~/.local/share/mercury/discord-feed-cursor` | Path to cursor persistence file |
| `POLL_INTERVAL_MS` | no | `2000` | Poll interval in milliseconds |
| `BATCH_LIMIT` | no | `50` | Max messages per poll cycle |

## Running

```bash
export DISCORD_BOT_TOKEN="your-bot-token"
export DISCORD_FEED_CHANNEL_ID="123456789012345678"
bun run start
```

## Message format

Each Mercury message becomes a Discord embed:

```
+------------------------------------------+
|  <emoji> #channel-name                   |  <- title with channel-specific color
|                                          |
|  sender-name                             |  <- author
|  Message body text here                  |  <- description
|                                          |
|  2026-03-22 10:27:49 UTC                 |  <- footer timestamp
+------------------------------------------+
```

- Short status messages (e.g. "online, starting work") are posted as compact plain text instead of embeds
- Messages over 4000 characters are truncated with a note
- Consecutive duplicate messages from the same sender are skipped

## Systemd service example

```ini
[Unit]
Description=Mercury to Discord live feed
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
ExecStart=/path/to/bun run /path/to/tools/discord-feed/index.ts
Restart=always
RestartSec=5
Environment=DISCORD_BOT_TOKEN=your-token
Environment=DISCORD_FEED_CHANNEL_ID=your-channel-id
Environment=MERCURY_DB_PATH=/path/to/mercury.db

[Install]
WantedBy=multi-user.target
```

## Behavior notes

- On startup, if there are more than 100 unprocessed messages, the service skips the backlog and posts a summary instead of flooding the channel
- The service waits and retries if the Mercury DB doesn't exist yet (useful when starting before Mercury has been used)
- Cursor writes are atomic (write-to-temp then rename) to prevent corruption on crash
- `discord.js` handles Discord rate limits and reconnection automatically
