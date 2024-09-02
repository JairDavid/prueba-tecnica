package response

type Common struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
