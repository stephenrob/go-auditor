package events

import "time"

type AuditEvent struct {
	ID int
	Action string
	Actor string
	ActorID string
	CatalogService string
	CategoryType string
	Timestamp time.Time
	ServerID string
	Note string
	Metadata map[string]interface{}
}
