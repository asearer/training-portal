training-portal/
├── cmd/                    # Entry points
│   └── server/             # Main HTTP server
│       └── main.go
├── internal/               # Application code (clean architecture)
│   ├── domain/             # Business logic interfaces & models
│   │   ├── user/
│   │   │   └── model.go
│   │   ├── course/
│   │   │   └── model.go
│   ├── usecase/            # Business logic implementations
│   │   ├── user/
│   │   │   └── service.go
│   │   ├── course/
│   │   │   └── service.go
│   ├── interface/          # Adapters
│   │   ├── http/           # HTTP delivery layer
│   │   │   ├── middleware/
│   │   │   ├── handler/
│   │   │   │   ├── user.go
│   │   │   │   └── course.go
│   │   │   └── router.go
│   │   ├── repository/     # Interface adapters (DB)
│   │   │   ├── postgres/
│   │   │   │   ├── user.go
│   │   │   │   └── course.go
├── migrations/             # SQL migration scripts
├── configs/                # Config files (e.g., app.yaml)
├── scripts/                # DevOps scripts (Docker, etc.)
├── docs/                   # API docs, architecture diagrams
├── web/                    # Frontend app (React, etc.)
├── .env                    # Environment variables
├── go.mod
└── go.sum
