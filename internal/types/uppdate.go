package types

type Update struct {
	Type string `json:"type"`
	Data any    `json:"data,omitempty"`
}
