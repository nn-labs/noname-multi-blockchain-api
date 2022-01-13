package errors

type Status string

const (
	statusInternalError   Status = "internal_error"
	statusNotFoundError   Status = "not_found_error"
	statusBadRequestError Status = "bad_request_error"
)
