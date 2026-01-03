package target

import "fmt"

// Error codes returned by the Target API.
const (
	ErrCodeUnauthorized         = "UNAUTHORIZED"
	ErrCodeInvalidRequest       = "INVALID_REQUEST"
	ErrCodeNotFound             = "NOT_FOUND"
	ErrCodeUpstreamError        = "UPSTREAM_ERROR"
	ErrCodeInternalError        = "INTERNAL_ERROR"
	ErrCodeInventoryUnavailable = "INVENTORY_UNAVAILABLE"
)

// APIError represents an error returned by the Target API.
type APIError struct {
	Message    string `json:"error"`
	Status     int    `json:"status,omitempty"`
	Details    string `json:"details,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
	Code       string `json:"code,omitempty"`
	HTTPStatus int    `json:"-"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("target: %s (%s)", e.Message, e.Details)
	}
	if e.Code != "" {
		return fmt.Sprintf("target: [%s] %s", e.Code, e.Message)
	}
	return fmt.Sprintf("target: %s", e.Message)
}

// IsUnauthorized returns true if the error is an authentication error.
func (e *APIError) IsUnauthorized() bool {
	return e.HTTPStatus == 401
}

// IsNotFound returns true if the requested resource was not found.
func (e *APIError) IsNotFound() bool {
	return e.HTTPStatus == 404
}

// IsInvalidRequest returns true if the error is due to an invalid request.
func (e *APIError) IsInvalidRequest() bool {
	return e.HTTPStatus == 400
}

// IsUpstreamError returns true if the error is from an upstream service.
func (e *APIError) IsUpstreamError() bool {
	return e.HTTPStatus == 502
}

// IsInternalError returns true if the error is an internal server error.
func (e *APIError) IsInternalError() bool {
	return e.HTTPStatus == 500
}

// IsInventoryUnavailable returns true if the item is not available for the selected fulfillment method.
func (e *APIError) IsInventoryUnavailable() bool {
	return e.Code == ErrCodeInventoryUnavailable || e.HTTPStatus == 424
}
