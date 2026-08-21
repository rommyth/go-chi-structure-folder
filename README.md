# Go Chi Structure Folder

An example Restaurant Management API that demonstrates a modular folder structure using [Chi](https://github.com/go-chi/chi). The currently implemented feature is food creation through the `food` module's `CreateFood` handler. Other modules and endpoints remain examples or placeholders unless noted below.

## Framework and libraries

- **HTTP router:** [go-chi/chi](https://github.com/go-chi/chi) v5
- **Database:** PostgreSQL through `pgx/v5` and `pgxpool`
- **Configuration:** Viper, loaded from `config.yml`
- **Authentication:** JWT with `go-chi/jwtauth`
- **Validation:** `go-playground/validator`
- **Logging:** Go's standard `log/slog`

## Folder structure

```text
.
├── cmd/
│   └── api/                     # Application entry point and root router setup
│       ├── main.go              # Creates shared dependencies and starts the server
│       ├── app.go               # Wires module dependencies together
│       └── routes.go            # Registers global, guest, and protected routes
├── internal/
│   ├── config/                  # Viper, database, JWT, logger, and validator setup
│   ├── middleware/              # Reusable HTTP middleware (JWT verification)
│   └── modules/                 # Feature modules
│       ├── auth/                # Authentication DTOs and claims
│       ├── food/                # Food module; CreateFood is implemented
│       │   ├── dto.go           # CreateFood request DTO
│       │   ├── model.go         # Food domain/database model
│       │   ├── repository.go    # Food database access
│       │   ├── service.go       # Menu check and food creation logic
│       │   ├── handler.go       # HTTP handlers
│       │   └── routes.go        # /foods routes
│       ├── health/              # Health-check handler
│       └── user/                # Example feature module
│           ├── dto.go           # Request/response data-transfer objects
│           ├── model.go         # Domain/database models
│           ├── repository.go    # Database access interface and implementation
│           ├── service.go       # Business-logic layer
│           ├── handler.go       # HTTP handler layer
│           └── routes.go        # Module-specific route registration
├── migrations/                  # Reserved for database migrations
├── pkg/
│   └── response/                # Shared JSON success and error response helpers
├── config.yml                   # Server, database, and JWT configuration
└── go.mod
```

`internal` keeps application-specific packages private to this project. Put code intended to be reused by multiple modules in `pkg`.

## Request flow

```text
HTTP request
  -> cmd/api/routes.go (Chi router and global middleware)
  -> guest or JWT-protected route group
  -> module routes.go
  -> Handler
  -> Service
  -> Repository
  -> PostgreSQL
  -> pkg/response JSON response
```

At startup, `cmd/api/main.go` loads configuration, initializes the logger, validator, PostgreSQL connection pool, and JWT signer. `cmd/api/app.go` creates each module's repository, service, and handler, then passes the handlers to the root router.

## The `food` module: CreateFood

The active endpoint is `POST /api/foods/`. It is registered in the JWT-protected route group, so every request must include a valid bearer token:

```http
Authorization: Bearer <token-jwt>
Content-Type: application/json
```

Before creating food, ensure that `menu_id` already exists in the `menus` table. The service checks the menu through `menu.Repository`, and the `foods.menu_id` database foreign key provides an additional safeguard.

### Request

```json
{
  "menu_id": 1,
  "name": "Nasi Goreng",
  "price": 25000,
  "image": "https://example.com/images/nasi-goreng.jpg"
}
```

| Field | Type | Required | Rule |
| --- | --- | --- | --- |
| `menu_id` | integer | Yes | Existing menu ID |
| `name` | string | Yes | 2 to 100 characters |
| `price` | number | Yes | Application validation accepts `0` or greater; database requires a value greater than `0` |
| `image` | string or `null` | No | Optional image URL/path |

Example using `curl`:

```bash
curl -X POST http://localhost:3000/api/foods/ \
  -H "Authorization: Bearer <token-jwt>" \
  -H "Content-Type: application/json" \
  -d '{"menu_id":1,"name":"Nasi Goreng","price":25000,"image":"https://example.com/images/nasi-goreng.jpg"}'
```

On success, the API returns `201 Created`:

```json
{
  "status": true,
  "message": "success create food",
  "data": {
    "id": 1,
    "name": "Nasi Goreng",
    "price": 25000,
    "image": "https://example.com/images/nasi-goreng.jpg",
    "created_at": "2026-08-21T00:00:00Z",
    "updated_at": "2026-08-21T00:00:00Z",
    "menu_id": 1
  }
}
```

Malformed JSON or a failed request validation returns `400 Bad Request`. Errors while checking the menu or persisting the food currently return `500 Internal Server Error`.

### CreateFood request flow

```text
Client
  │ POST /api/foods/ + Bearer JWT
  ▼
cmd/api/routes.go
  │ Global middleware: RequestID, Logger, Recoverer
  │ JWT Verify and Authenticator middleware
  ▼
internal/modules/food/routes.go → Handler.CreateFood
  │ Decode JSON and validate CreateFoodRequest
  ▼
food.Service.Create
  │ Verify menu_id with menu.Repository.GetByID
  │ Build Food model
  ▼
food.Repository.Create
  │ INSERT INTO foods (...) RETURNING ...
  ▼
PostgreSQL → pkg/response.Success → 201 JSON response → Client
```

At startup, `cmd/api/app.go` wires the dependencies in this order: `food.NewRepository` receives the PostgreSQL pool, `food.NewService` receives the food and menu repositories, and `food.NewHandler` receives the service, logger, and validator.

### Module dependency pattern

```text
routes -> handler -> service -> repository -> database
```

This separation keeps HTTP concerns, business rules, and data access independent.

To add another module, follow the same pattern:

1. Create `internal/modules/<module-name>`.
2. Add model and DTO types as needed.
3. Define a repository interface and implementation.
4. Add a service that depends on the repository interface.
5. Add a handler that depends on the service.
6. Register the module routes in its `routes.go`.
7. Wire the module in `cmd/api/app.go` and register it from `cmd/api/routes.go` in the appropriate route group.

## Running the project

1. Install Go, run a PostgreSQL instance, and install the `migrate` CLI when using the provided migration targets.
2. Update the `database` and `jwt` values in `config.yml` for your environment.
3. Create the configured database (by default, `restaurant_management`).
4. Apply the migrations, including `foods` and its dependency tables:

   ```bash
   make migrate-up
   ```

   The current `Makefile` migration command uses local PostgreSQL credentials `postgres:postgres`; update it when your database configuration differs.

5. From the repository root, start the API:

   ```bash
   go run ./cmd/api
   ```

The server listens on `http://localhost:3000` by default.

## Current routes

| Method | Path | Access | Purpose |
| --- | --- | --- | --- |
| `GET` | `/api/health` | Guest | Health check |
| `GET` | `/api/guest` | Guest | Guest-route example |
| `GET` | `/api/protected` | JWT | Protected-route example |
| `GET` | `/api/users/test` | JWT | User module example; returns a JSON response |
| `POST` | `/api/foods/` | JWT | Creates a food record; implemented |
| `GET` | `/api/foods/` | JWT | Registered, handler is still a placeholder |
| `GET` | `/api/foods/{id}` | JWT | Registered, handler is still a placeholder |
| `PATCH` | `/api/foods/{id}` | JWT | Registered, handler is still a placeholder |
| `/api/menus/*` | JWT | Registered, menu handlers are still placeholders |

All protected routes require a valid JWT in the `Authorization` header:

```text
Authorization: Bearer <token>
```

The current middleware expects the token to contain string `user_id`, `email`, and `role` claims when they are read through `middleware.GetClaims`.
