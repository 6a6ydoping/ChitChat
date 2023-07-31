package api

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Ok struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
