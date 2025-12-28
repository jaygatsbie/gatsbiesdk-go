package gatsbie

// HealthResponse is returned by the health check endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// SolveResponse is the generic response for successful captcha solves.
type SolveResponse[T any] struct {
	Success   bool    `json:"success"`
	TaskID    string  `json:"taskId"`
	Service   string  `json:"service"`
	Solution  T       `json:"solution"`
	Cost      float64 `json:"cost"`
	SolveTime float64 `json:"solveTime"`
}

// DatadomeRequest is the request for solving Datadome device check challenges.
type DatadomeRequest struct {
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

// DatadomeSolution is returned when solving Datadome challenges.
type DatadomeSolution struct {
	Datadome  string `json:"datadome"`
	UserAgent string `json:"ua"`
}

// RecaptchaV3Request is the request for solving reCAPTCHA v3 challenges.
type RecaptchaV3Request struct {
	Proxy      string `json:"proxy"`
	TargetURL  string `json:"target_url"`
	SiteKey    string `json:"site_key"`
	Action     string `json:"action,omitempty"`
	Title      string `json:"title,omitempty"`
	Enterprise bool   `json:"enterprise,omitempty"`
}

// RecaptchaV3Solution is returned when solving reCAPTCHA v3 challenges.
type RecaptchaV3Solution struct {
	Token     string `json:"token"`
	UserAgent string `json:"ua"`
}

// AkamaiRequest is the request for solving Akamai challenges.
type AkamaiRequest struct {
	Proxy       string `json:"proxy"`
	TargetURL   string `json:"target_url"`
	AkamaiJSURL string `json:"akamai_js_url"`
	PageFP      string `json:"page_fp,omitempty"`
}

// AkamaiCookies contains the cookies returned by Akamai.
type AkamaiCookies struct {
	Country   string `json:"Country,omitempty"`
	UsrLocale string `json:"UsrLocale,omitempty"`
	Abck      string `json:"_abck"`
	BmSz      string `json:"bm_sz"`
}

// AkamaiSolution is returned when solving Akamai challenges.
type AkamaiSolution struct {
	CookiesDict AkamaiCookies `json:"cookies_dict"`
	UserAgent   string        `json:"ua"`
}

// VercelRequest is the request for solving Vercel challenges.
type VercelRequest struct {
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
}

// VercelSolution is returned when solving Vercel challenges.
type VercelSolution struct {
	Vcrcs     string `json:"_vcrcs"`
	UserAgent string `json:"ua"`
}

// ShapeRequest is the request for solving Shape challenges.
type ShapeRequest struct {
	Proxy      string `json:"proxy"`
	TargetURL  string `json:"target_url"`
	TargetAPI  string `json:"target_api"`
	ShapeJSURL string `json:"shape_js_url"`
	Title      string `json:"title"`
	Method     string `json:"method"`
}

// ShapeSolution is returned when solving Shape challenges.
// Note: Shape uses dynamic header names that vary by site.
// Use the Headers map to access all solution headers.
type ShapeSolution map[string]string

// TurnstileRequest is the request for solving Cloudflare Turnstile challenges.
type TurnstileRequest struct {
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
	SiteKey   string `json:"site_key"`
}

// TurnstileSolution is returned when solving Cloudflare Turnstile challenges.
type TurnstileSolution struct {
	Token     string `json:"token"`
	UserAgent string `json:"ua"`
}

// PerimeterXRequest is the request for solving PerimeterX Invisible challenges.
type PerimeterXRequest struct {
	Proxy          string `json:"proxy"`
	TargetURL      string `json:"target_url"`
	PerimeterXJSURL string `json:"perimeterx_js_url"`
	PxAppID        string `json:"pxAppId"`
}

// PerimeterXCookies contains the PerimeterX cookies needed for requests.
type PerimeterXCookies struct {
	Px3   string `json:"_px3"`
	Pxde  string `json:"_pxde"`
	Pxvid string `json:"_pxvid"`
	Pxcts string `json:"pxcts"`
}

// PerimeterXSolution is returned when solving PerimeterX challenges.
type PerimeterXSolution struct {
	Cookies   PerimeterXCookies `json:"perimeterx_cookies"`
	UserAgent string            `json:"ua"`
}

// CloudflareWAFRequest is the request for solving Cloudflare WAF challenges.
type CloudflareWAFRequest struct {
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

// CloudflareWAFCookies contains the cookies returned by Cloudflare WAF.
type CloudflareWAFCookies struct {
	CfClearance string `json:"cf_clearance"`
}

// CloudflareWAFSolution is returned when solving Cloudflare WAF challenges.
type CloudflareWAFSolution struct {
	Cookies   CloudflareWAFCookies `json:"cookies"`
	UserAgent string               `json:"ua"`
}

// DatadomeSliderRequest is the request for solving Datadome Slider challenges.
type DatadomeSliderRequest struct {
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

// DatadomeSliderSolution is returned when solving Datadome Slider challenges.
type DatadomeSliderSolution struct {
	Datadome  string `json:"datadome"`
	UserAgent string `json:"ua"`
}

// CaptchaFoxRequest is the request for solving CaptchaFox challenges.
type CaptchaFoxRequest struct {
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
	SiteKey   string `json:"site_key"`
}

// CaptchaFoxCookies contains the cookies returned by CaptchaFox.
type CaptchaFoxCookies struct {
	BmS  string `json:"bm_s"`
	BmSc string `json:"bm_sc"`
}

// CaptchaFoxSolution is returned when solving CaptchaFox challenges.
type CaptchaFoxSolution struct {
	Cookie    CaptchaFoxCookies `json:"cookie"`
	UserAgent string            `json:"ua"`
}

// CastleConfigJSON contains the Castle configuration parameters.
type CastleConfigJSON struct {
	AvoidCookies bool   `json:"avoidCookies,omitempty"`
	PK           string `json:"pk"`
	WUrl         string `json:"wUrl"`
	SwUrl        string `json:"swUrl"`
}

// CastleRequest is the request for solving Castle challenges.
type CastleRequest struct {
	Proxy      string           `json:"proxy"`
	TargetURL  string           `json:"target_url"`
	ConfigJSON CastleConfigJSON `json:"config_json"`
}

// CastleSolution is returned when solving Castle challenges.
type CastleSolution struct {
	Token     string `json:"token"`
	UserAgent string `json:"ua"`
}

// Reese84Request is the request for solving Incapsula Reese84 challenges.
type Reese84Request struct {
	Proxy        string `json:"proxy"`
	Reese84JsUrl string `json:"reese84_js_url"`
}

// Reese84Solution is returned when solving Reese84 challenges.
type Reese84Solution struct {
	Reese84   string `json:"reese84"`
	UserAgent string `json:"user_agent"`
}

// ForterRequest is the request for solving Forter challenges.
type ForterRequest struct {
	Proxy       string `json:"proxy"`
	TargetURL   string `json:"target_url"`
	ForterJsUrl string `json:"forter_js_url"`
	SiteID      string `json:"site_id"`
}

// ForterSolution is returned when solving Forter challenges.
type ForterSolution struct {
	Token     string `json:"token"`
	UserAgent string `json:"ua"`
}

// FuncaptchaRequest is the request for solving Funcaptcha (Arkose Labs) challenges.
type FuncaptchaRequest struct {
	Proxy         string `json:"proxy"`
	TargetURL     string `json:"target_url"`
	CustomApiHost string `json:"custom_api_host"`
	PublicKey     string `json:"public_key"`
}

// FuncaptchaSolution is returned when solving Funcaptcha challenges.
type FuncaptchaSolution struct {
	Token     string `json:"token"`
	UserAgent string `json:"ua"`
}

// SBSDRequest is the request for solving Akamai SBSD challenges.
type SBSDRequest struct {
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

// SBSDSolution is returned when solving SBSD challenges.
type SBSDSolution struct {
	BmS       string `json:"bm_s"`
	BmSc      string `json:"bm_sc"`
	UserAgent string `json:"ua"`
}

// internal request wrappers that include task_type
type datadomeRequestInternal struct {
	TaskType     string `json:"task_type"`
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

type recaptchaV3RequestInternal struct {
	TaskType   string `json:"task_type"`
	Proxy      string `json:"proxy"`
	TargetURL  string `json:"target_url"`
	SiteKey    string `json:"site_key"`
	Action     string `json:"action,omitempty"`
	Title      string `json:"title,omitempty"`
	Enterprise bool   `json:"enterprise,omitempty"`
}

type akamaiRequestInternal struct {
	TaskType    string `json:"task_type"`
	Proxy       string `json:"proxy"`
	TargetURL   string `json:"target_url"`
	AkamaiJSURL string `json:"akamai_js_url"`
	PageFP      string `json:"page_fp,omitempty"`
}

type vercelRequestInternal struct {
	TaskType  string `json:"task_type"`
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
}

type shapeRequestInternal struct {
	TaskType   string `json:"task_type"`
	Proxy      string `json:"proxy"`
	TargetURL  string `json:"target_url"`
	TargetAPI  string `json:"target_api"`
	ShapeJSURL string `json:"shape_js_url"`
	Title      string `json:"title"`
	Method     string `json:"method"`
}

type turnstileRequestInternal struct {
	TaskType  string `json:"task_type"`
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
	SiteKey   string `json:"site_key"`
}

type perimeterXRequestInternal struct {
	TaskType        string `json:"task_type"`
	Proxy           string `json:"proxy"`
	TargetURL       string `json:"target_url"`
	PerimeterXJSURL string `json:"perimeterx_js_url"`
	PxAppID         string `json:"pxAppId"`
}

type cloudflareWAFRequestInternal struct {
	TaskType     string `json:"task_type"`
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

type datadomeSliderRequestInternal struct {
	TaskType     string `json:"task_type"`
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}

type captchaFoxRequestInternal struct {
	TaskType  string `json:"task_type"`
	Proxy     string `json:"proxy"`
	TargetURL string `json:"target_url"`
	SiteKey   string `json:"site_key"`
}

type castleRequestInternal struct {
	TaskType   string           `json:"task_type"`
	Proxy      string           `json:"proxy"`
	TargetURL  string           `json:"target_url"`
	ConfigJSON CastleConfigJSON `json:"config_json"`
}

type reese84RequestInternal struct {
	TaskType     string `json:"task_type"`
	Proxy        string `json:"proxy"`
	Reese84JsUrl string `json:"reese84_js_url"`
}

type forterRequestInternal struct {
	TaskType    string `json:"task_type"`
	Proxy       string `json:"proxy"`
	TargetURL   string `json:"target_url"`
	ForterJsUrl string `json:"forter_js_url"`
	SiteID      string `json:"site_id"`
}

type funcaptchaRequestInternal struct {
	TaskType      string `json:"task_type"`
	Proxy         string `json:"proxy"`
	TargetURL     string `json:"target_url"`
	CustomApiHost string `json:"custom_api_host"`
	PublicKey     string `json:"public_key"`
}

type sbsdRequestInternal struct {
	TaskType     string `json:"task_type"`
	Proxy        string `json:"proxy"`
	TargetURL    string `json:"target_url"`
	TargetMethod string `json:"target_method"`
}
