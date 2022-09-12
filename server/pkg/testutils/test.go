package testutils

import (
	"testing"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
)

type TestWithDBClient struct {
	Client *ent.Client
}

//SetupTestEntClient creates a new ent client for testing with sqlite backend
// Close the connection after calling the function
func SetupTestEntClient(t *testing.T) *ent.Client {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	return client
}
