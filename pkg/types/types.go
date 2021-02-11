package types

import "github.lancs.ac.uk/library/hutch"

type AuditEvent struct{
	Actor string `json:"actor"`
	Entity string `json:"entity"`
	Description string `json:"description"`
}

func (e *AuditEvent) MarshalData() (map[string]interface{}, error) {
	return hutch.JSONDataMarshal(e)
}