package response

type Error struct {
	Message     string      `json:"message"`
	Description interface{} `json:"error"`
}
