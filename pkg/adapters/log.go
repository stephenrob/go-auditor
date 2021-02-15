package adapters

import (
	"github.lancs.ac.uk/library/auditor/pkg/types"
	"log"
)

type LogAdapter struct {
	logger *log.Logger
}

func NewLogAdapter(logger *log.Logger) *LogAdapter {
	return &LogAdapter{
		logger: logger,
	}
}

// Start is a NO-OP function for the LogAdapter
func (l *LogAdapter) Start() error {
	return nil
}

// Shutdown is a NO-OP function for the LogAdapter
func (l *LogAdapter) Shutdown() error {
	return nil
}

// LogEvent outputs the audit event to the logger
func (l *LogAdapter) LogEvent(e *types.AuditEvent) error {
	l.logger.Printf("actor:%s,entity:%s,description:\"%s\"", e.Actor, e.Entity, e.Description)
	return nil
}

