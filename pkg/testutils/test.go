package testutils

import (
	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
	"testing"
)

type TestWithDBClient struct {
	Client *ent.Client
}

//SetupTestEntClient creates a new ent client for testing with sqlite backend
// Close the connection after calling the function
func SetupTestEntClient(t *testing.T) *ent.Client {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	return client
}
