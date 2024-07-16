package ws

type Message[T any] struct {
	Type string `json:"type"` // 訊息類型
	Data T      `json:"data"`
}
