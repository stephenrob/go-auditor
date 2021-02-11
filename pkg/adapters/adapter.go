package adapters

import "github.lancs.ac.uk/library/auditor/pkg/types"

type AuditAdapter interface {
	Start() error
	Shutdown() error
	LogEvent(e *types.AuditEvent) error
}
