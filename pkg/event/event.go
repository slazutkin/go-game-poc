package event

const (
	EventConnected    = "EVT_CONNECTED"
	EventDisconnected = "EVT_DISCONNECTED"
	EventData         = "EVT_DATA"
)

type Event struct {
	ClientID string
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}
