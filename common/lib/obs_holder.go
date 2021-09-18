package lib

import (
	"context"

	"github.com/IlyaFloppy/observable"
)

type ObsHolder struct {
	events   *observable.Object[[]IEvent]
	eventsCh <-chan []IEvent
}

func NewObsHolder() *ObsHolder {
	r := &ObsHolder{
		events: observable.New[[]IEvent](nil),
	}

	r.eventsCh = r.events.Subscribe(context.Background(), false)

	return r
}

func (h *ObsHolder) Process() {
	for {
		select {
		default:
			return
		case events, ok := <-h.eventsCh:
			if !ok {
				panic("events channel is closed")
			}
			for _, e := range events {
				e.Apply(h)
			}
		}
	}
}
