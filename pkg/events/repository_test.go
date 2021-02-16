package events

import (
	"github.com/go-pg/pg/v10"
	"os"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {

	dbs := os.Getenv("DATABASE_CONNECTION_STRING")

	opt, err := pg.ParseURL(dbs)

	if err != nil {
		t.Fatalf("Failed to connect to Postgres DB: %s", err)
	}

	db := pg.Connect(opt)

	t.Run("CreateEvent", func(t *testing.T) {

		r := &DBRepository{DB: db}

		e := &AuditEvent{
			Action:         "user.reload",
			Actor:          "username",
			ActorID:        "123456",
			CatalogService: "Service/Accounts",
			CategoryType:   "User Management",
			Timestamp:      time.Now(),
			ServerID:      "pod_abcd_12345",
			Note:           "Manual reload of user account",
			Metadata: map[string]interface{}{"hello": "world", "service": "123"},
		}

		err = r.Create(e)

		if err != nil {
			t.Fatalf("Failed to create audit event: %s", err)
		}

	})

	t.Run("GetEvent", func(t *testing.T) {

		r := &DBRepository{DB: db}

		e, err := r.GetByID(1)

		if err != nil {
			t.Fatalf("Failed to get audit event: %s", err)
		}

		if e.Actor != "username" {
			t.Errorf("Actor - GOT %s WANT %s", e.Actor, "username")
		}

	})
}
