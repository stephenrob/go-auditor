package types

import (
	"github.lancs.ac.uk/library/hutch"
	"time"
)

type AuditEvent struct{
	ID int `json:"id"`
	Action string `json:"action"`
	Actor string `json:"actor"`
	ActorID string `json:"actor_id"`
	CatalogService string `json:"catalog_service"`
	CategoryType string `json:"category_type"`
	Timestamp time.Time `json:"timestamp"`
	ServerID string `json:"server_id"`
	Note string `json:"note"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (e *AuditEvent) MarshalData() (map[string]interface{}, error) {
	return hutch.JSONDataMarshal(e)
}