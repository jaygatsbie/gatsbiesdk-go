package gatsbie

import "fmt"

// Error codes returned by the Gatsbie API.
const (
	ErrCodeAuthFailed           = "AUTH_FAILED"
	ErrCodeInsufficientCredits  = "INSUFFICIENT_CREDITS"
	ErrCodeInvalidRequest       = "INVALID_REQUEST"
	ErrCodeUpstreamError        = "UPSTREAM_ERROR"
	ErrCodeSolveFailed          = "SOLVE_FAILED"
	ErrCodeInternalError        = "INTERNAL_ERROR"
)

// APIError represents an error returned by the Gatsbie API.
type APIError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
	Timestamp  int64  `json:"timestamp"`
	HTTPStatus int    `json:"-"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("gatsbie: %s: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("gatsbie: %s: %s", e.Code, e.Message)
}

// IsAuthError returns true if the error is an authentication error.
func (e *APIError) IsAuthError() bool {
	return e.Code == ErrCodeAuthFailed
}

// IsInsufficientCredits returns true if the error is due to insufficient credits.
func (e *APIError) IsInsufficientCredits() bool {
	return e.Code == ErrCodeInsufficientCredits
}

// IsInvalidRequest returns true if the error is due to an invalid request.
func (e *APIError) IsInvalidRequest() bool {
	return e.Code == ErrCodeInvalidRequest
}

// IsUpstreamError returns true if the error is from an upstream service.
func (e *APIError) IsUpstreamError() bool {
	return e.Code == ErrCodeUpstreamError
}

// IsSolveFailed returns true if the captcha solving failed.
func (e *APIError) IsSolveFailed() bool {
	return e.Code == ErrCodeSolveFailed
}

// IsInternalError returns true if the error is an internal server error.
func (e *APIError) IsInternalError() bool {
	return e.Code == ErrCodeInternalError
}

// errorResponse is the structure returned by the API on error.
type errorResponse struct {
	Success bool      `json:"success"`
	TaskID  string    `json:"taskId,omitempty"`
	Error   *APIError `json:"error"`
}
