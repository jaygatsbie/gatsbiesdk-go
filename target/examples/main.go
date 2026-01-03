package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jaygatsbie/gatsbiesdk-go/target"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("GATSBIE_API_KEY")
	if apiKey == "" {
		log.Fatal("GATSBIE_API_KEY environment variable is required")
	}

	// Create client
	client := target.NewClient(apiKey)
	ctx := context.Background()

	// Health check
	fmt.Println("=== Health Check ===")
	health, err := client.Health(ctx)
	if err != nil {
		log.Printf("Health check failed: %v", err)
	} else {
		fmt.Printf("Status: %s\n", health.Status)
	}

	// Ping (get quota info)
	fmt.Println("\n=== Ping ===")
	ping, err := client.Ping(ctx)
	if err != nil {
		log.Printf("Ping failed: %v", err)
	} else {
		fmt.Printf("Message: %s\n", ping.Message)
		if ping.QuotaUsed > 0 {
			fmt.Printf("Quota: %d / %d\n", ping.QuotaUsed, ping.QuotaLimit)
		}
	}

	// Find nearby stores
	fmt.Println("\n=== Nearby Stores ===")
	stores, err := client.GetNearbyStores(ctx, target.NearbyStoresRequest{
		Lat:   40.7147,
		Lng:   -74.0112,
		Limit: 5,
	})
	if err != nil {
		log.Printf("Failed to get nearby stores: %v", err)
	} else {
		for _, store := range stores {
			fmt.Printf("- %s (%s, %s) - %.2f miles\n",
				store.Name, store.City, store.State, store.DistanceMiles)
		}
	}

	// Get product details (requires proxy)
	proxy := os.Getenv("PROXY_URL")
	if proxy != "" {
		fmt.Println("\n=== Product Details ===")
		product, err := client.GetProduct(ctx, target.GetProductRequest{
			TCIN:  "86777236",
			Proxy: proxy,
		})
		if err != nil {
			log.Printf("Failed to get product: %v", err)
		} else {
			fmt.Printf("Title: %s\n", product.Title)
			fmt.Printf("Price: %s\n", product.CurrentPrice)
			fmt.Printf("In Stock: %v\n", product.InStock)
			if product.OnSale {
				fmt.Printf("On Sale! Save %s (%.0f%% off)\n",
					product.SavingsAmount, product.SavingsPercent)
			}
			if len(product.Variations) > 0 {
				fmt.Printf("Variations: %d\n", len(product.Variations))
				for _, v := range product.Variations {
					fmt.Printf("  - %s: %s (%s)\n", v.Name, v.Value, v.CurrentPrice)
				}
			}
		}
	}

	// Add to cart (requires proxy and access token)
	accessToken := os.Getenv("TARGET_ACCESS_TOKEN")
	if proxy != "" && accessToken != "" {
		fmt.Println("\n=== Add to Cart ===")
		cartResp, err := client.AddToCart(ctx, target.AddToCartRequest{
			TCIN:        "94716087",
			Quantity:    1,
			AccessToken: accessToken,
			Proxy:       proxy,
		})
		if err != nil {
			log.Printf("Failed to add to cart: %v", err)
		} else {
			fmt.Printf("Success: %v\n", cartResp.Success)
			fmt.Printf("Cart ID: %s\n", cartResp.CartID)
			fmt.Printf("Items in Cart: %d\n", cartResp.TotalItemsInCart)
			fmt.Printf("Item: %s\n", cartResp.ItemAdded.Title)
			fmt.Printf("Total: $%.2f\n", cartResp.Pricing.Total)
		}
	}
}
