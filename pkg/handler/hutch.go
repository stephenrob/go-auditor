package handler

import (
	"context"
	"github.lancs.ac.uk/library/auditor/pkg/events"
	"github.lancs.ac.uk/library/auditor/pkg/types"
	"github.lancs.ac.uk/library/hutch"
	"log"
	"os"
)

type Handler interface {
	HandleEvent(event types.AuditEvent) error
}

type HutchHandler struct {
	Client hutch.Client
	logger *log.Logger
	ex     hutch.Exchange
	q hutch.Queue
	repo events.Repository
}

func (h *HutchHandler) HandleEvent(event types.AuditEvent) error {
	ev := events.AuditEvent(event)

	err := h.repo.Create(&ev)
	if err != nil {
		return err
	}

	h.logger.Printf("Stored audit event - %s by %s\n", ev.Action, ev.Actor)

	return nil
}

func NewHutchHandler(c hutch.Client, name string, repo events.Repository) (*HutchHandler, error) {

	ex, err := c.NewTopicExchange(name)
	if err != nil {
		return nil, err
	}

	h := &HutchHandler{
		Client: c,
		repo: repo,
		ex: ex,
		logger: log.New(os.Stdout, "HANDLER: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return h, nil
}

func (h *HutchHandler) HandleMessage(r hutch.RawMessage) {
	e := types.AuditEvent{}
	err := hutch.JSONDataUnmarshal(r.Data(), &e)
	if err != nil {
		return
	}
	err = h.HandleEvent(e)
	if err != nil {
		return
	}
}

func (h *HutchHandler) Subscribe(queue string, ctx context.Context) error {
	q, err := h.Client.NewQueue(queue)
	if err != nil {
		return err
	}
	h.q = q
	err = q.Bind(h.ex, "auditor.event")
	if err != nil {
		return err
	}
	err = q.Subscribe(ctx, h.HandleMessage)
	if err != nil {
		return err
	}
	return nil
}

func (h *HutchHandler) Close() error {
	h.q.WG().Wait()
	err := h.Client.Close()
	if err != nil {
		return err
	}
	return nil
}