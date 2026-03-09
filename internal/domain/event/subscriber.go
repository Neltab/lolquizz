package event

type Subscriber interface {
	Subscribe(eventName string, handler func(Event))
}
