package client

import (
	"github.lancs.ac.uk/library/auditor/pkg/adapters"
	"github.lancs.ac.uk/library/auditor/pkg/types"
)

func (a auditor) SubmitEvent(e *types.AuditEvent) error {
	return a.adapter.LogEvent(e)
}

func (a auditor) Shutdown() error {
	return a.adapter.Shutdown()
}

func NewAuditor(adapter adapters.AuditAdapter) Auditor {
	a := auditor{adapter: adapter}
	err := a.adapter.Start()
	if err != nil {
		return nil
	}
	return &a
}
