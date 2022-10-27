package events

import (
	"github.com/qwertyqq2/event-listener/pkg/mongodb/event"
)

type DepositCreate struct {
	Type  string `bson:"type"`
	From  string `bson:"owner"`
	Value int    `bson:"value"`
}

func NewDeposit(t, from string, value int) event.Event {
	return &DepositCreate{
		Type:  t,
		From:  from,
		Value: value,
	}
}
