package gowes

func Hello() string {
	return "Hello World"
}

type SocketPayload struct {
	From        string      `json:"from"`
	To          []Client    `json:"to"`
	MessageType string      `json:"message_type"`
	MessageData interface{} `json:"message_data"`
}

type Client struct {
	ClientID string `json:"client_id"`
}

// standard message data
type Message struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

const (
	// TypeMessage is a message type
	TypeGetNotification  = "get_notification"
	TypeSendNotification = "send_notification"
)
