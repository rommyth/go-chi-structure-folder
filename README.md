# Go Chi Structure Folder

An example Restaurant Management API that demonstrates a modular folder structure using [Chi](https://github.com/go-chi/chi). The `user` module is the reference implementation for adding a feature to the project.

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
├── Makefile                     # Terminal short way command
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

## The `user` module example

The `internal/modules/user` package shows the intended dependency direction:

```text
routes -> handler -> service -> repository -> database
```

`NewRepository` receives the shared `pgxpool.Pool`; `NewService` receives the repository interface; and `NewHandler` receives the service. `user.RegisterRoutes` attaches user endpoints to the router supplied by the application. This separation keeps HTTP concerns, business rules, and data access independent.

To add another module, follow the same pattern:

1. Create `internal/modules/<module-name>`.
2. Add model and DTO types as needed.
3. Define a repository interface and implementation.
4. Add a service that depends on the repository interface.
5. Add a handler that depends on the service.
6. Register the module routes in its `routes.go`.
7. Wire the module in `cmd/api/app.go` and register it from `cmd/api/routes.go` in the appropriate route group.

## Running the project

1. Install Go and run a PostgreSQL instance.
2. Update the `database` and `jwt` values in `config.yml` for your environment.
3. Create the configured database (by default, `restaurant_management`).
4. From the repository root, start the API:

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

All protected routes require a valid JWT in the `Authorization` header:

```text
Authorization: Bearer <token>
```

The current middleware expects the token to contain string `user_id`, `email`, and `role` claims when they are read through `middleware.GetClaims`.
