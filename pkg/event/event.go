package event

const (
	EventConnected    = "EVT_CONNECTED"
	EventDisconnected = "EVT_DISCONNECTED"
	EventData         = "EVT_DATA"
)

type Event struct {
	Type     string
	ClientID string
	Data     interface{}
}
