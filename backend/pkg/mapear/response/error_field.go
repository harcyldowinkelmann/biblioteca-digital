package response

type ErrorResponseField struct {
	Message     string            `json:"message"`
	Description map[string]string `json:"errors"`
}
