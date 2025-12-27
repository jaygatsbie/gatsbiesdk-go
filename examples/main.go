package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gatsbie "github.com/jaygatsbie/sdk-go"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("GATSBIE_API_KEY")
	if apiKey == "" {
		log.Fatal("GATSBIE_API_KEY environment variable is required")
	}

	// Create client with default options
	client := gatsbie.NewClient(apiKey)

	// Or with custom options
	// client := gatsbie.NewClient(apiKey,
	// 	gatsbie.WithTimeout(60*time.Second),
	// 	gatsbie.WithBaseURL("https://custom.api.url"),
	// )

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Check API health
	health, err := client.Health(ctx)
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Printf("API Status: %s\n\n", health.Status)

	// Example: Solve Turnstile challenge
	solveTurnstileExample(ctx, client)

	// Example: Solve Datadome challenge
	// solveDatadomeExample(ctx, client)

	// Example: Solve reCAPTCHA v3 challenge
	// solveRecaptchaV3Example(ctx, client)

	// Example: Solve Akamai challenge
	// solveAkamaiExample(ctx, client)
}

func solveTurnstileExample(ctx context.Context, client *gatsbie.Client) {
	fmt.Println("Solving Turnstile challenge...")

	resp, err := client.SolveTurnstile(ctx, &gatsbie.TurnstileRequest{
		Proxy:     "http://user:pass@proxy.example.com:8080",
		TargetURL: "https://example.com/protected-page",
		SiteKey:   "0x4AAAAAAABS7TtLxsNa7Z2e",
	})

	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Success! Task ID: %s\n", resp.TaskID)
	fmt.Printf("Token: %s...\n", truncate(resp.Solution.Token, 50))
	fmt.Printf("User-Agent: %s\n", resp.Solution.UserAgent)
	fmt.Printf("Cost: %.4f credits\n", resp.Cost)
	fmt.Printf("Solve Time: %.2f ms\n\n", resp.SolveTime)
}

func solveDatadomeExample(ctx context.Context, client *gatsbie.Client) {
	fmt.Println("Solving Datadome device check...")

	resp, err := client.SolveDatadome(ctx, &gatsbie.DatadomeRequest{
		Proxy:        "http://user:pass@proxy.example.com:8080",
		TargetURL:    "https://www.cma-cgm.com/",
		TargetMethod: "GET",
	})

	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Success! Task ID: %s\n", resp.TaskID)
	fmt.Printf("Datadome Cookie: %s...\n", truncate(resp.Solution.Datadome, 50))
	fmt.Printf("User-Agent: %s\n", resp.Solution.UserAgent)
	fmt.Printf("Cost: %.4f credits\n", resp.Cost)
	fmt.Printf("Solve Time: %.2f ms\n\n", resp.SolveTime)
}

func solveRecaptchaV3Example(ctx context.Context, client *gatsbie.Client) {
	fmt.Println("Solving reCAPTCHA v3...")

	resp, err := client.SolveRecaptchaV3(ctx, &gatsbie.RecaptchaV3Request{
		Proxy:      "http://user:pass@proxy.example.com:8080",
		TargetURL:  "https://2captcha.com/demo/recaptcha-v3",
		SiteKey:    "6Lcyqq8oAAAAAJE7eVJ3aZp_hnJcI6LgGdYD8lge",
		Action:     "demo_action",
		Title:      "Google reCAPTCHA V3 demo: Sample Form with Google reCAPTCHA V3",
		Enterprise: false,
	})

	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Success! Task ID: %s\n", resp.TaskID)
	fmt.Printf("Token: %s...\n", truncate(resp.Solution.Token, 50))
	fmt.Printf("User-Agent: %s\n", resp.Solution.UserAgent)
	fmt.Printf("Cost: %.4f credits\n", resp.Cost)
	fmt.Printf("Solve Time: %.2f ms\n\n", resp.SolveTime)
}

func solveAkamaiExample(ctx context.Context, client *gatsbie.Client) {
	fmt.Println("Solving Akamai challenge...")

	resp, err := client.SolveAkamai(ctx, &gatsbie.AkamaiRequest{
		Proxy:       "http://user:pass@proxy.example.com:8080",
		TargetURL:   "https://shop.lululemon.com/",
		AkamaiJSURL: "https://shop.lululemon.com/WGlx/lc_w/w/vez/w0HNXw/EmubktLXh3Npr6Nab5/TXUGYQ/Lh9aC2xK/H34",
	})

	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("Success! Task ID: %s\n", resp.TaskID)
	fmt.Printf("_abck: %s...\n", truncate(resp.Solution.Abck, 50))
	fmt.Printf("bm_sz: %s...\n", truncate(resp.Solution.BmSz, 50))
	fmt.Printf("User-Agent: %s\n", resp.Solution.UserAgent)
	fmt.Printf("Cost: %.4f credits\n", resp.Cost)
	fmt.Printf("Solve Time: %.2f ms\n\n", resp.SolveTime)
}

func handleError(err error) {
	var apiErr *gatsbie.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API Error [%s]: %s\n", apiErr.Code, apiErr.Message)
		if apiErr.Details != "" {
			fmt.Printf("Details: %s\n", apiErr.Details)
		}

		// Handle specific error types
		switch {
		case apiErr.IsAuthError():
			fmt.Println("Check your API key")
		case apiErr.IsInsufficientCredits():
			fmt.Println("Please add more credits to your account")
		case apiErr.IsSolveFailed():
			fmt.Println("The captcha could not be solved, try again")
		}
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
