package protocol

type MessageType string

const (
	Error         MessageType = "error"
	QuoteRequest              = "quote-request"
	QuoteResponse             = "quote-response"
)

type Message struct {
	Type MessageType `json:"type"`
	Body string      `json:"body"`
}
