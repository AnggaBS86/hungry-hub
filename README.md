# HungryHub – Restaurant Menu Management API (Echo + MySQL)

Implements the take-home assignment REST API using Go, Echo, and MySQL.

## Tech stack

- Go + Echo (`github.com/labstack/echo/v4`)
- MySQL 8
- GORM (`gorm.io/gorm`) + MySQL driver (`gorm.io/driver/mysql`)

## Quick start (Docker)

1. Copy env:
   - `cp .env.example .env`
2. Run:
   - `docker compose up --build`
3. API:
   - `http://localhost:8080/healthz`

Seeds can be enabled by setting `RUN_SEED=true` in `.env`.
If you run the binary from a different working directory, set `MIGRATIONS_DIR` to point at the `migrations/` folder.

## Local start (no Docker)

1. Create a MySQL database (default `hungryhub`)
2. Export env (or use your own method):
   - `export $(cat .env | xargs)`
3. Run:
   - `go run ./cmd/api`

## Endpoints

- `POST /restaurants`
- `GET /restaurants`
- `GET /restaurants/:id` (includes menu items)
- `PUT /restaurants/:id`
- `DELETE /restaurants/:id`
- `POST /restaurants/:id/menu_items`
- `GET /restaurants/:id/menu_items?category=drink`
- `PUT /menu_items/:id`
- `DELETE /menu_items/:id`

## Design notes

- Migrations are executed automatically on startup (controlled by `RUN_MIGRATIONS=true`).
- Seed data is optional (`RUN_SEED=true`) and is safe to run multiple times (skips if data exists).
- JSON errors follow a consistent shape: `{ "message": "...", "details": {...} }`.
