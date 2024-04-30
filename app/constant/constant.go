package constant

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	InvalidState
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "INVALID_STATE"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "INVALID_STATE"}[r-1]
}
