# Traffic Tracker

A Go-based traffic monitoring system that tracks commute times between predefined routes using Google Maps API, stores historical data in SQLite, and sends daily/monthly summaries via Telegram.

## What It Does

- **Tracks Traffic**: Monitors 2 routes (A→B) at 5-minute intervals during specified time frames
- **Stores Data**: Saves all traffic data in a SQLite database for historical analysis
- **Daily Summaries**: Sends end-of-day reports via Telegram with best commute times
- **Monthly Reports**: Provides monthly trends and recommendations for optimal routes

## Technology Stack

- **Language**: Go 1.21+
- **Database**: SQLite (via `modernc.org/sqlite` - pure Go, no CGo)
- **APIs**:
  - Google Maps Distance Matrix API (traffic data)
  - Telegram Bot API (notifications)
- **Scheduler**: GitHub Actions (cron-based workflows)
- **Deployment**: Compiled Go binaries

## Project Structure

```
traffic-tracker/
├── cmd/
│   ├── tracker/          # Main traffic tracking command
│   ├── daily/            # Daily summary command
│   └── monthly/          # Monthly summary command
├── internal/
│   ├── database/         # SQLite operations
│   ├── maps/             # Google Maps API client
│   ├── telegram/         # Telegram bot integration
│   ├── analyzer/         # Data analysis logic
│   └── config/           # Configuration management
├── .github/workflows/
│   ├── track-traffic.yml     # Runs every 5 minutes
│   ├── daily-summary.yml     # Runs at end of day
│   └── monthly-summary.yml   # Runs at end of month
├── go.mod
├── go.sum
├── .env.example
├── README.md
└── TODO.md
```

## Prerequisites

1. **Go 1.21+** - [Install Go](https://go.dev/doc/install)
2. **Google Maps API Key**:
   - Go to [Google Cloud Console](https://console.cloud.google.com/)
   - Create a project and enable "Distance Matrix API"
   - Create API credentials
   - Free tier: 40,000 requests/month
3. **Telegram Bot**:
   - Message [@BotFather](https://t.me/botfather) on Telegram
   - Create a new bot and get the API token
   - Get your Chat ID (message [@userinfobot](https://t.me/userinfobot))

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your credentials
3. Build the project: `go build ./cmd/tracker`
4. Initialize the database: `./tracker --init`
5. Test tracking: `./tracker`

## Configuration

Edit `.env` file with your settings:

```env
# Google Maps API
GOOGLE_MAPS_API_KEY=your_api_key_here

# Routes to track
ROUTE_1_ORIGIN=Times Square, New York, NY
ROUTE_1_DESTINATION=JFK Airport, Queens, NY
ROUTE_2_ORIGIN=Brooklyn Bridge, NY
ROUTE_2_DESTINATION=Central Park, NY

# Telegram Bot
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHAT_ID=your_chat_id

TRACK_START_HOUR=7
TRACK_START_MINUTE=30
TRACK_END_HOUR=9
TRACK_END_MINUTE=0
TRACK_TIMEZONE=Europe/Madrid
```

## Usage

### Manual Tracking

```bash
go run ./cmd/tracker
```

### Daily Summary

```bash
go run ./cmd/daily
```

### Monthly Summary

```bash
go run ./cmd/monthly
```

## GitHub Actions Setup

1. Push this repository to GitHub
2. Add secrets in repository settings:
   - `GOOGLE_MAPS_API_KEY`
   - `TELEGRAM_BOT_TOKEN`
   - `TELEGRAM_CHAT_ID`
3. Workflows will run automatically based on schedule

## Data Storage

All traffic data is stored in `traffic.db` (SQLite database) with the following schema:

- Route information (origin, destination)
- Timestamp of measurement
- Duration in traffic (minutes)
- Traffic conditions
