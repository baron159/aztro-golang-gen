package msg

import "time"

// Message represents a single message in a chat conversation.
type Message struct {
	Sender    string      `json:"sender" yaml:"sender"`       // The sender of the message.
	Content   string      `json:"content" yaml:"content"`     // The content of the message.
	Timestamp time.Time   `json:"timestamp" yaml:"timestamp"` // The timestamp of the message.
	Sent      bool        `json:"sent" yaml:"sent"`           // Indicates if the message was sent by the user.
	Party     interface{} `json:"party" yaml:"party"`
}

func NewMessage(content string, sent bool) *Message {
	return &Message{
		Sender:    "User",
		Content:   content,
		Timestamp: time.Now(),
		Sent:      sent,
	}
}
