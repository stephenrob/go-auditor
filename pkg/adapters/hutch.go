package adapters

import (
	"fmt"
	"github.com/google/uuid"
	"github.lancs.ac.uk/library/auditor/pkg/types"
	"github.lancs.ac.uk/library/hutch"
)

type HutchAdapter struct {
	Client hutch.Client
	ex     hutch.Exchange
}

func NewHutchAdapter(c hutch.Client) *HutchAdapter {
	return &HutchAdapter{
		Client: c,
	}
}

func (h *HutchAdapter) Start() error {
	ex, err := h.Client.NewTopicExchange("auditor")
	if err != nil {
		return err
	}
	h.ex = ex
	return nil
}

func (h *HutchAdapter) Shutdown() error {
	err := h.Client.Close()
	if err != nil {
		return err
	}
	return nil
}

func (h *HutchAdapter) LogEvent(e *types.AuditEvent) error {
	msg := hutch.Message{
		Meta: hutch.NewMessageMeta(uuid.NewString(), "auditor.event", "0.1.0"),
		Data: e,
	}
	err := h.ex.Publish(hutch.DeliverableMessage(msg), "auditor.event")
	if err != nil {
		return err
	}
	fmt.Printf("Published log event\n")
	return nil
}