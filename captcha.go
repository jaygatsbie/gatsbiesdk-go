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

// SolveShape solves a Shape antibot challenge (v1).
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

// SolveShapeV2 solves a Shape antibot challenge using the v2 API with TLS fingerprinting.
func (c *Client) SolveShapeV2(ctx context.Context, req *ShapeV2Request) (*SolveResponse[ShapeV2Solution], error) {
	// Build metadata map for the API
	metadata := map[string]interface{}{
		"proxy": req.Proxy,
	}
	if req.Pkey != "" {
		metadata["pkey"] = req.Pkey
	}
	if req.ScriptURL != "" {
		metadata["script_url"] = req.ScriptURL
	}
	if len(req.Request) > 0 {
		metadata["request"] = req.Request
	}
	if req.Country != "" {
		metadata["country"] = req.Country
	}
	if req.Timeout > 0 {
		metadata["timeout"] = req.Timeout
	}

	internal := map[string]interface{}{
		"url":      req.URL,
		"metadata": metadata,
	}

	var resp SolveResponse[ShapeV2Solution]
	if err := c.doPost(ctx, "/v1/solve/shape-v2", internal, &resp); err != nil {
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

// SolveCaptchaFox solves a CaptchaFox challenge.
func (c *Client) SolveCaptchaFox(ctx context.Context, req *CaptchaFoxRequest) (*SolveResponse[CaptchaFoxSolution], error) {
	internal := captchaFoxRequestInternal{
		TaskType:  "captchafox",
		Proxy:     req.Proxy,
		TargetURL: req.TargetURL,
		SiteKey:   req.SiteKey,
	}

	var resp SolveResponse[CaptchaFoxSolution]
	if err := c.doPost(ctx, "/v1/solve/captchafox", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveCastle solves a Castle challenge.
func (c *Client) SolveCastle(ctx context.Context, req *CastleRequest) (*SolveResponse[CastleSolution], error) {
	internal := castleRequestInternal{
		TaskType:   "castle",
		Proxy:      req.Proxy,
		TargetURL:  req.TargetURL,
		ConfigJSON: req.ConfigJSON,
	}

	var resp SolveResponse[CastleSolution]
	if err := c.doPost(ctx, "/v1/solve/castle", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveReese84 solves an Incapsula Reese84 challenge.
func (c *Client) SolveReese84(ctx context.Context, req *Reese84Request) (*SolveResponse[Reese84Solution], error) {
	internal := reese84RequestInternal{
		TaskType:     "reese84",
		Proxy:        req.Proxy,
		Reese84JsUrl: req.Reese84JsUrl,
	}

	var resp SolveResponse[Reese84Solution]
	if err := c.doPost(ctx, "/v1/solve/reese84", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveForter solves a Forter challenge.
func (c *Client) SolveForter(ctx context.Context, req *ForterRequest) (*SolveResponse[ForterSolution], error) {
	internal := forterRequestInternal{
		TaskType:    "forter",
		Proxy:       req.Proxy,
		TargetURL:   req.TargetURL,
		ForterJsUrl: req.ForterJsUrl,
		SiteID:      req.SiteID,
	}

	var resp SolveResponse[ForterSolution]
	if err := c.doPost(ctx, "/v1/solve/forter", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveFuncaptcha solves a Funcaptcha (Arkose Labs) challenge.
func (c *Client) SolveFuncaptcha(ctx context.Context, req *FuncaptchaRequest) (*SolveResponse[FuncaptchaSolution], error) {
	internal := funcaptchaRequestInternal{
		TaskType:      "funcaptcha",
		Proxy:         req.Proxy,
		TargetURL:     req.TargetURL,
		CustomApiHost: req.CustomApiHost,
		PublicKey:     req.PublicKey,
	}

	var resp SolveResponse[FuncaptchaSolution]
	if err := c.doPost(ctx, "/v1/solve/funcaptcha", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SolveSBSD solves an Akamai SBSD challenge.
func (c *Client) SolveSBSD(ctx context.Context, req *SBSDRequest) (*SolveResponse[SBSDSolution], error) {
	internal := sbsdRequestInternal{
		TaskType:     "sbsd",
		Proxy:        req.Proxy,
		TargetURL:    req.TargetURL,
		TargetMethod: req.TargetMethod,
	}

	var resp SolveResponse[SBSDSolution]
	if err := c.doPost(ctx, "/v1/solve/sbsd", internal, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
