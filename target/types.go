package target

// HealthResponse is returned by the health check endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// PingResponse is returned by the ping endpoint.
type PingResponse struct {
	Message    string `json:"message"`
	QuotaUsed  int    `json:"quota_used,omitempty"`
	QuotaLimit int    `json:"quota_limit,omitempty"`
}

// StoreResponse represents a Target store with distance information.
type StoreResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	PostalCode     string  `json:"postalCode"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	DriveUpEnabled bool    `json:"driveUpEnabled"`
	DistanceMiles  float64 `json:"distanceMiles"`
}

// NearbyStoresRequest is the request for finding nearby stores.
type NearbyStoresRequest struct {
	Lat    float64
	Lng    float64
	Limit  int     // Optional, defaults to 10
	Radius float64 // Optional, defaults to 50.0 miles
}

// ProductVariation represents a product variation (color, size, etc.).
type ProductVariation struct {
	TCIN                 string `json:"tcin"`
	Name                 string `json:"name"`
	Value                string `json:"value"`
	SwatchImageURL       string `json:"swatch_image_url,omitempty"`
	PrimaryImageURL      string `json:"primary_image_url"`
	CurrentPrice         string `json:"current_price"`
	InStock              bool   `json:"in_stock"`
	AvailableForShipping bool   `json:"available_for_shipping"`
	AvailableForPickup   bool   `json:"available_for_pickup"`
}

// ProductResponse represents product information.
type ProductResponse struct {
	TCIN                  string             `json:"tcin"`
	Title                 string             `json:"title"`
	CurrentPrice          string             `json:"current_price"`
	RegularPrice          string             `json:"regular_price,omitempty"`
	OnSale                bool               `json:"on_sale"`
	SavingsAmount         string             `json:"savings_amount,omitempty"`
	SavingsPercent        float64            `json:"savings_percent,omitempty"`
	PrimaryImageURL       string             `json:"primary_image_url"`
	InStock               bool               `json:"in_stock"`
	AvailableForShipping  bool               `json:"available_for_shipping"`
	AvailableForPickup    bool               `json:"available_for_pickup"`
	FreeShippingAvailable bool               `json:"free_shipping_available"`
	RatingAverage         float64            `json:"rating_average,omitempty"`
	RatingCount           int                `json:"rating_count,omitempty"`
	ReviewCount           int                `json:"review_count,omitempty"`
	Variations            []ProductVariation `json:"variations,omitempty"`
}

// GetProductRequest is the request for getting product details.
type GetProductRequest struct {
	TCIN    string
	StoreID string // Optional, defaults to "3229" (Tribeca, NY)
	Proxy   string // Required, format: http://user:pass@host:port
}

// FulfillmentType represents the fulfillment method for cart items.
type FulfillmentType string

const (
	FulfillmentShip        FulfillmentType = "SHIP"
	FulfillmentCurbside    FulfillmentType = "CURBSIDE"
	FulfillmentStorePickup FulfillmentType = "STORE_PICKUP"
)

// AddToCartRequest is the request for adding an item to cart.
type AddToCartRequest struct {
	TCIN            string          `json:"tcin"`
	Quantity        int             `json:"quantity"`
	AccessToken     string          `json:"access_token"`
	Proxy           string          `json:"proxy"`
	FulfillmentType FulfillmentType `json:"fulfillment_type,omitempty"`
	StoreID         string          `json:"store_id,omitempty"`
}

// AddedItemSummary represents the item that was added to cart.
type AddedItemSummary struct {
	TCIN      string  `json:"tcin"`
	Title     string  `json:"title"`
	ImageURL  string  `json:"image_url"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
}

// FulfillmentSummary represents fulfillment details for cart items.
type FulfillmentSummary struct {
	Type          string `json:"type"`
	StoreName     string `json:"store_name,omitempty"`
	EstimatedDate string `json:"estimated_date"`
	PickupHours   int    `json:"pickup_hours,omitempty"`
}

// PricingSummary represents the price breakdown for cart items.
type PricingSummary struct {
	ItemTotal float64 `json:"item_total"`
	Shipping  float64 `json:"shipping"`
	Tax       float64 `json:"tax"`
	Total     float64 `json:"total"`
}

// ReturnPolicySummary represents return policy information.
type ReturnPolicySummary struct {
	Days           int `json:"days"`
	DaysWithCircle int `json:"days_with_circle"`
}

// AddToCartResponse is the response after adding an item to cart.
type AddToCartResponse struct {
	Success          bool                `json:"success"`
	Message          string              `json:"message,omitempty"`
	CartID           string              `json:"cart_id"`
	TotalItemsInCart int                 `json:"total_items_in_cart"`
	ItemAdded        AddedItemSummary    `json:"item_added"`
	Fulfillment      FulfillmentSummary  `json:"fulfillment"`
	Pricing          PricingSummary      `json:"pricing"`
	ReturnPolicy     ReturnPolicySummary `json:"return_policy,omitempty"`
}

// ErrorResponse represents an API error response.
type ErrorResponse struct {
	Error      string `json:"error"`
	Status     int    `json:"status,omitempty"`
	Details    string `json:"details,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
	Code       string `json:"code,omitempty"`
}
