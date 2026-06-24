package integration_test

import "testing"

// These tests require a live Postgres instance (set DATABASE_URL env var).
// Run with: go test ./tests/integration/...

func TestGoalCRUD(t *testing.T) {
	t.Skip("integration tests require a running database — remove this line to enable")
}
