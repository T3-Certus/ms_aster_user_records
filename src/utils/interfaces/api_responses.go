package utils_interfaces

type DefaultResponse struct {
	HttpStatus    uint16 `json:"httpStatus"`
	ServerMessage string `json:"serverMessage"`
	Error         bool   `json:"error"`
	Success       bool   `json:"success"`
}

type GenericSuccessResponse struct {
	DefaultResponse
	Payload     interface{} `json:"responseBody,omitempty"`
	MoreDetails string      `json:"moreDetails"`
}

type GenericErrorResponse struct {
	DefaultResponse
	ErrorMessage string `json:"errorMessage"`
}
