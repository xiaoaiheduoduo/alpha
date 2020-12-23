package aerror

import (
	"encoding/json"
	"fmt"
	"strings"
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

func ErrUnknown(messages ...string) *Error {
	return New(CodeUnknown).WithHttpStatusCode(500).WithMessages(messages...)
}

func ErrUnauthorized(messages ...string) *Error {
	return New(CodeUnauthorized).WithHttpStatusCode(401).WithMessages(messages...)
}

func ErrForbidden(messages ...string) *Error {
	return New(CodeForbidden).WithHttpStatusCode(403).WithMessages(messages...)
}

func ErrNotFound(messages ...string) *Error {
	return New(CodeNotFound).WithHttpStatusCode(404).WithMessages(messages...)
}

func ErrAlreadyExists(messages ...string) *Error {
	return New(CodeAlreadyExists).WithHttpStatusCode(409).WithMessages(messages...)
}

func ErrConflict(messages ...string) *Error {
	return New(CodeConflict).WithHttpStatusCode(409).WithMessages(messages...)
}

func ErrGone(messages ...string) *Error {
	return New(CodeGone).WithHttpStatusCode(410).WithMessages(messages...)
}

func ErrInvalid(messages ...string) *Error {
	return New(CodeInvalid).WithHttpStatusCode(422).WithMessages(messages...)
}

func ErrServerTimeout(messages ...string) *Error {
	return New(CodeServerTimeout).WithHttpStatusCode(500).WithMessages(messages...)
}

func ErrTimeout(messages ...string) *Error {
	return New(CodeTimeout).WithHttpStatusCode(504).WithMessages(messages...)
}

func ErrTooManyRequests(messages ...string) *Error {
	return New(CodeTooManyRequests).WithHttpStatusCode(429).WithMessages(messages...)
}

func ErrBadRequest(messages ...string) *Error {
	return New(CodeBadRequest).WithHttpStatusCode(400).WithMessages(messages...)
}

func ErrMethodNotAllowed(messages ...string) *Error {
	return New(CodeMethodNotAllowed).WithHttpStatusCode(405).WithMessages(messages...)
}

func ErrNotAcceptable(messages ...string) *Error {
	return New(CodeNotAcceptable).WithHttpStatusCode(406).WithMessages(messages...)
}

func ErrRequestEntityTooLarge(messages ...string) *Error {
	return New(CodeRequestEntityTooLarge).WithHttpStatusCode(413).WithMessages(messages...)
}

func ErrUnsupportedMediaType(messages ...string) *Error {
	return New(CodeUnsupportedMediaType).WithHttpStatusCode(415).WithMessages(messages...)
}

func ErrInternalError(messages ...string) *Error {
	return New(CodeInternalError).WithHttpStatusCode(500).WithMessages(messages...)
}

func ErrExpired(messages ...string) *Error {
	return New(CodeExpired).WithHttpStatusCode(410).WithMessages(messages...)
}

func ErrServiceUnavailable(messages ...string) *Error {
	return New(CodeServiceUnavailable).WithHttpStatusCode(503).WithMessages(messages...)
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
func (ae *Error) WithMessages(messages ...string) *Error {
	if len(messages) > 0 {
		ae.Err.Message = strings.Join(messages, ";")
	}

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
