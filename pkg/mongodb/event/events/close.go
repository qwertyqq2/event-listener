package events

import (
	"github.com/qwertyqq2/event-listener/pkg/mongodb/event"
)

type DepositClose struct {
	Type string `bson:"type"`
	From string `bson:"owner"`
}

func NewClose(t, from string) event.Event {
	return &DepositClose{Type: t, From: from}
}
