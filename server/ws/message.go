package ws

type Message[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}
