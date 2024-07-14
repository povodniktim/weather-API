package models

type Response struct {
	Status int         `json:"status"`
	Error  bool        `json:"error"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}
