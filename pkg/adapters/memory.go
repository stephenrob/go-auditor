package adapters

import "github.lancs.ac.uk/library/auditor/pkg/types"

type MemoryAdapter struct {
	events []*types.AuditEvent
}

func (m *MemoryAdapter) Start() error {
	return nil
}

func (m *MemoryAdapter) Shutdown() error {
	return nil
}

func (m *MemoryAdapter) LogEvent(e *types.AuditEvent) error {
	m.events = append(m.events, e)
	return nil
}

func (m *MemoryAdapter) GetEvents() []*types.AuditEvent {
	return m.events
}

func NewMemoryAdapter() *MemoryAdapter {
	return &MemoryAdapter{}
}