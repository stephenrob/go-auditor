package client

import (
	"github.lancs.ac.uk/library/auditor/pkg/adapters"
	"github.lancs.ac.uk/library/auditor/pkg/types"
)

type Auditor interface {
	SubmitEvent(e *types.AuditEvent) error
	Shutdown() error
}

type auditor struct {
	adapter adapters.AuditAdapter
}