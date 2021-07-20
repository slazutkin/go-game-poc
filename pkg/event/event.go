package event

type Event struct {
	Type     Type
	ClientID string
	Data     interface{}
}

type Handler interface {
	HandleEvent(event *Event)
}

type Type string

const EventConnected = Type("EVENT_CONNECTED")
const EventDisconnected = Type("EVENT_DISCONNECTED")
const EventData = Type("EVENT_DATA")
