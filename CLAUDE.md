# GoalApp

## Stack
- Backend: Go 1.25 / Gin / PostgreSQL
- Frontend: Vue 3 / TypeScript / Pinia / Tailwind CSS / Vite (pnpm)

## Commands

### Backend
```
make run-backend   # start API server (port 8080)
make test          # run all tests
make migrate       # run DB migrations
make lint          # run linters
```

### Frontend
```
pnpm dev           # start dev server
pnpm test:unit     # run Vitest unit tests
pnpm lint          # run ESLint + oxlint
pnpm build         # production build → frontend/dist/
```

## Architecture

**Backend** — clean arch under `backend/internal/`:
- `handlers/` → HTTP layer (Gin)
- `service/` → business logic
- `repository/` → DB queries (Postgres)
- `domain/` → structs (Goal, Schedule)
- `middleware/` → JWT auth (`auth.go`), CORS (`cors.go`), logger

**Frontend** — `frontend/src/`:
- `views/` → page components
- `stores/` → Pinia state (authStore, goalStore, scheduleStore)
- `services/` → API calls
- `composables/` → reusable logic
- `router/index.ts` → route definitions

## Key Constraints
- JWT auth is enforced in `backend/internal/middleware/auth.go` — all protected routes pass through it
- DB connection via `DATABASE_URL` env var; copy `.env.example` → `.env` to configure
- Frontend API base set via `VITE_API_BASE_URL` env var (default `http://localhost:8080`)
- Entry point: `backend/cmd/api/main.go`
