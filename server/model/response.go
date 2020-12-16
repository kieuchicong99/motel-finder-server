package model

type InsertResponse struct {
	Message string      `json:"Message"`
	Error   string      `json:"Error"`
	Data    interface{} `json:"Data"`
}

type UpdateResponse struct {
	Message string      `json:"Message"`
	Error   string      `json:"Error"`
	Data    interface{} `json:"Data"`
}

type GetOneResponse struct {
	Message string      `json:"Message"`
	Error   string      `json:"Error"`
	Data    interface{} `json:"Data"`
}

type GetManyResponse struct {
	Message string        `json:"Message"`
	Error   string        `json:"Error"`
	Data    []interface{} `json:"Data"`
}
