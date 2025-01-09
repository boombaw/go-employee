package constants

type DefaultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors,omitempty"`
}

type PaginationData struct {
	Page        uint `json:"page"`
	TotalPages  uint `json:"totalPages"`
	TotalItems  uint `json:"totalItems"`
	Limit       uint `json:"limit"`
	HasNext     bool `json:"hasNext"`
	HasPrevious bool `json:"hasPrevious"`
}

type PaginationResponseData struct {
	Results        interface{} `json:"results"`
	PaginationData `json:"pagination"`
}

const (
	MESSAGE_SUCCESS                = "Success"
	MESSAGE_STILL_PROCESS          = "Transaction is being process"
	MESSAGE_FAILED                 = "Something went wrong"
	MESSAGE_INVALID_REQUEST_FORMAT = "Invalid Request Format"
	MESSAGE_UNAUTHORIZED           = "Unauthorized"
	MESSAGE_FORBIDDEN              = "Forbidden"
	MESSAGE_CONFLICT               = "Conflict"
)
