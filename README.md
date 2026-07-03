# Sharing Vision BE

## Prerequisites

- Go 1.22+
- Docker & Docker Compose (for MySQL)

## Quick Start

```bash
# 1. Copy env and adjust if needed
cp .env.example .env

# 2. Start MySQL
docker compose up mysql -d

# 3. Run migration
#    Table is auto-created on first MySQL startup via migration/init.sql.
#    If MySQL was already running, run manually:
mysql -u root -p sharing_vision < migration/init.sql

# 4. Start the server
go mod tidy
go run main.go
```

## With Docker Compose (MySQL + App)

```bash
docker compose up
```

Server runs on `http://localhost:8080`.

## API Endpoints

| Method | Path                     | Description          |
|--------|--------------------------|----------------------|
| POST   | `/article/`              | Create article       |
| GET    | `/article/{limit}/{offset}` | List with paging  |
| GET    | `/article/{id}`          | Get by ID            |
| PUT    | `/article/{id}`          | Update by ID         |
| DELETE | `/article/{id}`          | Delete by ID         |
