package adapters

import (
	"github.lancs.ac.uk/library/auditor/pkg/types"
)

type MultiAdapter struct {
	adapters []AuditAdapter
}

func (a *MultiAdapter) Start() error {
	for _, ad := range a.adapters {
		err := ad.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *MultiAdapter) Shutdown() error {
	for _, ad := range a.adapters {
		err := ad.Shutdown()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *MultiAdapter) LogEvent(e *types.AuditEvent) error {
	for _, ad := range a.adapters {
		err := ad.LogEvent(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *MultiAdapter) UseAdapter(ad AuditAdapter) error {
	a.adapters = append(a.adapters, ad)
	return nil
}