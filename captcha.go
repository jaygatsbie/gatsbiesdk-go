package gatsbie

import "context"

// Health checks the API server health status.
func (c *Client) Health(ctx context.Context) (*HealthResponse, error) {
	var resp HealthResponse
	if err := c.doGet(ctx, "/health", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveDatadome solves a Datadome device check challenge.
func (c *Client) SolveDatadome(ctx context.Context, req *DatadomeRequest) (*SolveResponse[DatadomeSolution], error) {
	internal := datadomeRequestInternal{
		TaskType:     "datadome-device-check",
		Proxy:        req.Proxy,
		TargetURL:    req.TargetURL,
		TargetMethod: req.TargetMethod,
	}

	var resp SolveResponse[DatadomeSolution]
	if err := c.doPost(ctx, "/v1/solve/datadome-device-check", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveRecaptchaV3 solves a reCAPTCHA v3 challenge.
func (c *Client) SolveRecaptchaV3(ctx context.Context, req *RecaptchaV3Request) (*SolveResponse[RecaptchaV3Solution], error) {
	internal := recaptchaV3RequestInternal{
		TaskType:   "recaptchav3",
		Proxy:      req.Proxy,
		TargetURL:  req.TargetURL,
		SiteKey:    req.SiteKey,
		Action:     req.Action,
		Title:      req.Title,
		Enterprise: req.Enterprise,
	}

	var resp SolveResponse[RecaptchaV3Solution]
	if err := c.doPost(ctx, "/v1/solve/recaptchav3", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveAkamai solves an Akamai bot management challenge.
func (c *Client) SolveAkamai(ctx context.Context, req *AkamaiRequest) (*SolveResponse[AkamaiSolution], error) {
	internal := akamaiRequestInternal{
		TaskType:    "akamai",
		Proxy:       req.Proxy,
		TargetURL:   req.TargetURL,
		AkamaiJSURL: req.AkamaiJSURL,
		PageFP:      req.PageFP,
	}

	var resp SolveResponse[AkamaiSolution]
	if err := c.doPost(ctx, "/v1/solve/akamai", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveVercel solves a Vercel bot protection challenge.
func (c *Client) SolveVercel(ctx context.Context, req *VercelRequest) (*SolveResponse[VercelSolution], error) {
	internal := vercelRequestInternal{
		TaskType:  "vercel",
		Proxy:     req.Proxy,
		TargetURL: req.TargetURL,
	}

	var resp SolveResponse[VercelSolution]
	if err := c.doPost(ctx, "/v1/solve/vercel", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveShape solves a Shape antibot challenge.
func (c *Client) SolveShape(ctx context.Context, req *ShapeRequest) (*SolveResponse[ShapeSolution], error) {
	internal := shapeRequestInternal{
		TaskType:   "shape",
		Proxy:      req.Proxy,
		TargetURL:  req.TargetURL,
		TargetAPI:  req.TargetAPI,
		ShapeJSURL: req.ShapeJSURL,
		Title:      req.Title,
		Method:     req.Method,
	}

	var resp SolveResponse[ShapeSolution]
	if err := c.doPost(ctx, "/v1/solve/shape", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveTurnstile solves a Cloudflare Turnstile challenge.
func (c *Client) SolveTurnstile(ctx context.Context, req *TurnstileRequest) (*SolveResponse[TurnstileSolution], error) {
	internal := turnstileRequestInternal{
		TaskType:  "turnstile",
		Proxy:     req.Proxy,
		TargetURL: req.TargetURL,
		SiteKey:   req.SiteKey,
	}

	var resp SolveResponse[TurnstileSolution]
	if err := c.doPost(ctx, "/v1/solve/turnstile", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolvePerimeterX solves a PerimeterX Invisible challenge.
func (c *Client) SolvePerimeterX(ctx context.Context, req *PerimeterXRequest) (*SolveResponse[PerimeterXSolution], error) {
	internal := perimeterXRequestInternal{
		TaskType:        "perimeterx_invisible",
		Proxy:           req.Proxy,
		TargetURL:       req.TargetURL,
		PerimeterXJSURL: req.PerimeterXJSURL,
		PxAppID:         req.PxAppID,
	}

	var resp SolveResponse[PerimeterXSolution]
	if err := c.doPost(ctx, "/v1/solve/perimeterx-invisible", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveCloudflareWAF solves a Cloudflare WAF challenge.
func (c *Client) SolveCloudflareWAF(ctx context.Context, req *CloudflareWAFRequest) (*SolveResponse[CloudflareWAFSolution], error) {
	internal := cloudflareWAFRequestInternal{
		TaskType:     "cloudflare_waf",
		Proxy:        req.Proxy,
		TargetURL:    req.TargetURL,
		TargetMethod: req.TargetMethod,
	}

	var resp SolveResponse[CloudflareWAFSolution]
	if err := c.doPost(ctx, "/v1/solve/cloudflare-waf", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveDatadomeSlider solves a Datadome Slider CAPTCHA challenge.
func (c *Client) SolveDatadomeSlider(ctx context.Context, req *DatadomeSliderRequest) (*SolveResponse[DatadomeSliderSolution], error) {
	internal := datadomeSliderRequestInternal{
		TaskType:     "datadome-slider",
		Proxy:        req.Proxy,
		TargetURL:    req.TargetURL,
		TargetMethod: req.TargetMethod,
	}

	var resp SolveResponse[DatadomeSliderSolution]
	if err := c.doPost(ctx, "/v1/solve/datadome-slider", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
