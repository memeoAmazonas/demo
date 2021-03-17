package model

type Response struct {
	Message     string  `json:"message"`
	Payload     *Task   `json:"payload,omitempty"`
	PayloadList []*Task `json:"payload_list,omitempty"`
}
