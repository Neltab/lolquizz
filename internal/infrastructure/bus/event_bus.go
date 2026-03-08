package bus

import (
	"lolquizz/internal/domain/event"
	"sync"
)

type EventBus struct {
	handlers map[string][]func(event.Event)
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]func(event.Event)),
	}
}

func (b *EventBus) Subscribe(eventName string, handler func(event.Event)) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.handlers[eventName] == nil {
		b.handlers[eventName] = make([]func(event.Event), 0)
	}
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Publish(event event.Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, handler := range b.handlers[event.EventName()] {
		go handler(event)
	}
}
