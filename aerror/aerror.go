package aerror

import (
	"encoding/json"
	"fmt"
)

type Code string

const (
	CodeUnknown               Code = "unknown"
	CodeUnauthorized          Code = "unauthorized"
	CodeForbidden             Code = "forbidden"
	CodeNotFound              Code = "not_found"
	CodeAlreadyExists         Code = "already_exists"
	CodeConflict              Code = "conflict"
	CodeGone                  Code = "gone"
	CodeInvalid               Code = "invalid"
	CodeServerTimeout         Code = "server_timeout"
	CodeTimeout               Code = "timeout"
	CodeTooManyRequests       Code = "too_many_requests"
	CodeBadRequest            Code = "bad_request"
	CodeMethodNotAllowed      Code = "method_not_allowed"
	CodeNotAcceptable         Code = "not_acceptable"
	CodeRequestEntityTooLarge Code = "request_entity_too_large"
	CodeUnsupportedMediaType  Code = "unsupported_media_type"
	CodeInternalError         Code = "internal_error"
	CodeExpired               Code = "expired"
	CodeServiceUnavailable    Code = "service_unavailable"
)

func ErrUnknown() *Error {
	return New(CodeUnknown).WithHttpStatusCode(500)
}

func ErrUnauthorized() *Error {
	return New(CodeUnauthorized).WithHttpStatusCode(401)
}

func ErrForbidden() *Error {
	return New(CodeForbidden).WithHttpStatusCode(403)
}

func ErrNotFound() *Error {
	return New(CodeNotFound).WithHttpStatusCode(404)
}

func ErrAlreadyExists() *Error {
	return New(CodeAlreadyExists).WithHttpStatusCode(409)
}

func ErrConflict() *Error {
	return New(CodeConflict).WithHttpStatusCode(409)
}

func ErrGone() *Error {
	return New(CodeGone).WithHttpStatusCode(410)
}

func ErrInvalid() *Error {
	return New(CodeInvalid).WithHttpStatusCode(422)
}

func ErrServerTimeout() *Error {
	return New(CodeServerTimeout).WithHttpStatusCode(500)
}

func ErrTimeout() *Error {
	return New(CodeTimeout).WithHttpStatusCode(504)
}

func ErrTooManyRequests() *Error {
	return New(CodeTooManyRequests).WithHttpStatusCode(429)
}

func ErrBadRequest() *Error {
	return New(CodeBadRequest).WithHttpStatusCode(400)
}

func ErrMethodNotAllowed() *Error {
	return New(CodeMethodNotAllowed).WithHttpStatusCode(405)
}

func ErrNotAcceptable() *Error {
	return New(CodeNotAcceptable).WithHttpStatusCode(406)
}

func ErrRequestEntityTooLarge() *Error {
	return New(CodeRequestEntityTooLarge).WithHttpStatusCode(413)
}

func ErrUnsupportedMediaType() *Error {
	return New(CodeUnsupportedMediaType).WithHttpStatusCode(415)
}

func ErrInternalError() *Error {
	return New(CodeInternalError).WithHttpStatusCode(500)
}

func ErrExpired() *Error {
	return New(CodeExpired).WithHttpStatusCode(410)
}

func ErrServiceUnavailable() *Error {
	return New(CodeServiceUnavailable).WithHttpStatusCode(503)
}

type Details map[string]interface{}

type errorPayload struct {
	Code    Code    `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Details Details `json:"details,omitempty"`
}

type Error struct {
	Err            errorPayload `json:"error,omitempty"`
	HTTPStatusCode int          `json:"-"`
}

func New(code Code) *Error {
	return &Error{
		Err: errorPayload{
			Code: code,
		},
	}
}

func NewWithAll(code Code, message string, details Details, httpStatusCode int) *Error {
	return &Error{
		Err: errorPayload{
			Code:    code,
			Message: message,
			Details: details,
		},
		HTTPStatusCode: httpStatusCode,
	}
}

func UnmarshallJSON(data []byte) (*Error, bool) {
	ae := &Error{}
	if err := json.Unmarshal(data, ae); err != nil {
		return nil, false
	}

	if len(ae.Err.Code) == 0 {
		return nil, false
	}

	return ae, true
}

func (ae *Error) Error() string {
	return ae.Err.Message
}

func (ae *Error) WithMessage(message string) *Error {
	ae.Err.Message = message

	return ae
}

func (ae *Error) WithMessagef(format string, a ...interface{}) *Error {
	ae.Err.Message = fmt.Sprintf(format, a...)

	return ae
}

func (ae *Error) WithHttpStatusCode(statusCode int) *Error {
	ae.HTTPStatusCode = statusCode

	return ae
}

func (ae *Error) WithDetails(details Details) *Error {
	ae.Err.Details = details

	return ae
}

func (ae *Error) WithError(err error) *Error {
	if e, ok := err.(*Error); ok {
		*ae = *e
	} else {
		ae.Err.Message = err.Error()
	}

	return ae
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.Err.Code == CodeNotFound {
		return true
	}

	return false
}
