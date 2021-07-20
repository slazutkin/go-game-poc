package ws

type MessageType string

const MessageConnected = MessageType("MSG_CONNECTED")
const MessageDisconnected = MessageType("MSG_DICSCONNECTED")
const MessageData = MessageType("MSG_DATA")

type InboundMessage struct {
	ClientID string
	Data     interface{}
}

type OutboundMessage struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

type ClientData struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data,omitempty"`
}

type MessageHandler func(msg *InboundMessage)
