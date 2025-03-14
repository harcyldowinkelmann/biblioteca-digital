package response

type SuccessBody struct {
	Message string      `json:"message"`
	Body    interface{} `Body:"message"`
}
