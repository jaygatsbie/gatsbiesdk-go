package target

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://target.gatsbie.io"
	defaultTimeout = 120 * time.Second
)

// Client is the Target API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// Option is a functional option for configuring the Client.
type Option func(*Client)

// WithBaseURL sets a custom base URL for the API.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new Target API client.
// The apiKey should be a valid Gatsbie API key starting with "gats_".
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// do performs an HTTP request with authentication.
func (c *Client) do(ctx context.Context, method, path string, body any, result any) error {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("target: failed to marshal request: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("target: failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("target: request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("target: failed to read response: %w", err)
	}

	// Check for error responses
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return &APIError{
				Message:    fmt.Sprintf("unexpected error response: %s", string(respBody)),
				HTTPStatus: resp.StatusCode,
			}
		}
		apiErr.HTTPStatus = resp.StatusCode
		return &apiErr
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("target: failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// doGet performs a GET request.
func (c *Client) doGet(ctx context.Context, path string, result any) error {
	return c.do(ctx, http.MethodGet, path, nil, result)
}

// doPost performs a POST request.
func (c *Client) doPost(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPost, path, body, result)
}

// Health checks the API server health status.
func (c *Client) Health(ctx context.Context) (*HealthResponse, error) {
	var result HealthResponse
	if err := c.doGet(ctx, "/health", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ping checks API connectivity and returns quota information.
func (c *Client) Ping(ctx context.Context) (*PingResponse, error) {
	var result PingResponse
	if err := c.doGet(ctx, "/api/v1/ping", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetNearbyStores returns a list of Target stores near the specified coordinates.
func (c *Client) GetNearbyStores(ctx context.Context, req NearbyStoresRequest) ([]StoreResponse, error) {
	params := url.Values{}
	params.Set("lat", strconv.FormatFloat(req.Lat, 'f', -1, 64))
	params.Set("lng", strconv.FormatFloat(req.Lng, 'f', -1, 64))

	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}
	if req.Radius > 0 {
		params.Set("radius", strconv.FormatFloat(req.Radius, 'f', -1, 64))
	}

	path := "/api/v1/stores/nearby?" + params.Encode()

	var result []StoreResponse
	if err := c.doGet(ctx, path, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetProduct returns detailed product information for a specific Target product.
func (c *Client) GetProduct(ctx context.Context, req GetProductRequest) (*ProductResponse, error) {
	if req.TCIN == "" {
		return nil, &APIError{Message: "tcin is required", HTTPStatus: 400}
	}
	if req.Proxy == "" {
		return nil, &APIError{Message: "proxy is required", HTTPStatus: 400}
	}

	params := url.Values{}
	params.Set("proxy", req.Proxy)
	if req.StoreID != "" {
		params.Set("store_id", req.StoreID)
	}

	path := fmt.Sprintf("/api/v1/products/%s?%s", req.TCIN, params.Encode())

	var result ProductResponse
	if err := c.doGet(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddToCart adds an item to the Target shopping cart.
func (c *Client) AddToCart(ctx context.Context, req AddToCartRequest) (*AddToCartResponse, error) {
	if req.TCIN == "" {
		return nil, &APIError{Message: "tcin is required", HTTPStatus: 400}
	}
	if req.Quantity < 1 {
		return nil, &APIError{Message: "quantity must be at least 1", HTTPStatus: 400}
	}
	if req.AccessToken == "" {
		return nil, &APIError{Message: "access_token is required", HTTPStatus: 400}
	}
	if req.Proxy == "" {
		return nil, &APIError{Message: "proxy is required", HTTPStatus: 400}
	}

	// Validate fulfillment type requires store_id
	if (req.FulfillmentType == FulfillmentCurbside || req.FulfillmentType == FulfillmentStorePickup) && req.StoreID == "" {
		return nil, &APIError{Message: "store_id is required when fulfillment_type is CURBSIDE or STORE_PICKUP", HTTPStatus: 400}
	}

	body := map[string]any{
		"tcin":         req.TCIN,
		"quantity":     req.Quantity,
		"access_token": req.AccessToken,
		"proxy":        req.Proxy,
	}
	if req.FulfillmentType != "" {
		body["fulfillment_type"] = string(req.FulfillmentType)
	}
	if req.StoreID != "" {
		body["store_id"] = req.StoreID
	}

	var result AddToCartResponse
	if err := c.doPost(ctx, "/api/v1/cart/items", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
