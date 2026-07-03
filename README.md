# Sharing Vision BE

## Prerequisites

- Docker & Docker Compose

## Quick Start

```bash
# 1. Copy env and adjust if needed
cp .env.example .env

# 2. Start all services
docker compose up
```

Server runs on `http://localhost:8080`. MySQL auto-starts and migration runs automatically on app boot.

## Seed Data (optional)

Insert 100 dummy articles to test pagination:

```bash
go run seed/main.go
```

## API Endpoints

| Method | Path                     | Description          |
|--------|--------------------------|----------------------|
| POST   | `/article/`              | Create article       |
| GET    | `/article/{limit}/{offset}` | List with paging  |
| GET    | `/article/{id}`          | Get by ID            |
| PUT    | `/article/{id}`          | Update by ID         |
| DELETE | `/article/{id}`          | Delete by ID         |
