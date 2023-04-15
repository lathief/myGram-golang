package helpers

type BaseResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func NewResponse(status int, data interface{}, Message string) *BaseResponse {
	return &BaseResponse{
		Status:  status,
		Data:    data,
		Message: Message,
	}
}