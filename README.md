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

## API docs

Base URL: `http://localhost:8080`  
Content-Type: `application/json`

### Health

- `GET /healthz`

Response:
```json
{ "status": "ok" }
```

### Errors

Common error shape:
```json
{ "message": "validation failed", "details": { "field": "reason" } }
```

### Restaurants

#### Create restaurant

- `POST /restaurants`

Body:
```json
{
  "name": "Bangkok Bistro",
  "address": "123 Sukhumvit Rd, Bangkok",
  "phone": "+66-2-111-1111",
  "opening_hours": "Mon-Sun 10:00-22:00"
}
```

Responses:
- `201` with created restaurant
- `400` if `name`/`address` missing

#### List restaurants

- `GET /restaurants`

Response: `200` array of restaurants (without menu items).

#### Get restaurant detail (includes menu items)

- `GET /restaurants/:id`

Response: `200` restaurant with `menu_items`.

#### Update restaurant

- `PUT /restaurants/:id`

Body (partial update; omitted fields stay unchanged):
```json
{ "phone": "+66-2-999-9999" }
```

Responses:
- `200` updated restaurant
- `404` if restaurant not found

#### Delete restaurant

- `DELETE /restaurants/:id`

Responses:
- `200` empty body
- `404` if restaurant not found

### Menu items

#### Add menu item to restaurant

- `POST /restaurants/:id/menu_items`

Body (`price` is a decimal string):
```json
{
  "name": "Pad Thai",
  "description": "Classic stir-fried noodles",
  "price": "180.00",
  "category": "main",
  "is_available": true
}
```

Responses:
- `201` created menu item
- `400` if `name`/`price` missing or `price` invalid
- `404` if restaurant not found

#### List menu items for restaurant

- `GET /restaurants/:id/menu_items`

Query params:
- `category` (optional): `GET /restaurants/:id/menu_items?category=drink`

Responses:
- `200` array of menu items
- `404` if restaurant not found

#### Update menu item

- `PUT /menu_items/:id`

Body (partial update):
```json
{ "is_available": false }
```

Responses:
- `200` updated menu item
- `404` if menu item not found

#### Delete menu item

- `DELETE /menu_items/:id`

Responses:
- `200` empty body
- `404` if menu item not found

## Design notes

- Migrations are executed automatically on startup (controlled by `RUN_MIGRATIONS=true`).
- Seed data is optional (`RUN_SEED=true`) and is safe to run multiple times (skips if data exists).
- JSON errors follow a consistent shape: `{ "message": "...", "details": {...} }`.
