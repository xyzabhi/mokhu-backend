go-api/
├── cmd/
│   └── main.go                  → Application entry point
├── internal/
│   ├── api/                     → Handlers, route definitions
│   ├── config/                  → App configuration (env, setup)
│   ├── db/                      → Database connection + sqlc code
│   ├── models/                  → Structs & validation rules
│   ├── repository/              → Query logic (sqlc wrapper)
│   └── service/                 → Business logic
├── pkg/                         → Utilities (logger, auth, etc.)
├── go.mod
└── .env
