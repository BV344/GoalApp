.PHONY: run-backend run-frontend test migrate lint

run-backend:
	cd backend && go run ./cmd/api

run-frontend:
	cd frontend && pnpm dev

test:
	cd backend && go test ./... && cd ../frontend && pnpm test

migrate:
	cd backend && go run ./cmd/migrate up

lint:
	cd backend && go vet ./... && cd ../frontend && pnpm lint
