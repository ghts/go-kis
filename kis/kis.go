package kis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func AuthOptionsFromEnv() AuthOptions {
	return AuthOptions{
		AppKey:    os.Getenv("KIS_APP_KEY"),
		AppSecret: os.Getenv("KIS_APP_SECRET"),
		AccountNo: os.Getenv("KIS_ACCOUNT_NO"),
	}
}

const defaultBaseURL = "https://openapi.koreainvestment.com:9443"

type AuthOptions struct {
	AppKey    string
	AppSecret string
	AccountNo string
}

type Client struct {
	client    *http.Client
	UserAgent string
	BaseURL   *url.URL

	AppKey    string
	AppSecret string
	AccountNo string

	AccessToken string
	// TODO: check expire before req
	// AccessTokenExpires time.Time

	// TODO: rate limiter
	// rateMu sync.Mutex
	// rateLimits [categories]Rate

	Debug bool

	common                  service // Reuse a single struct instead of allocating one for each service on the heap.
	DomesticStockQuotations *DomesticStockQuotationsService
	Oauth2                  *Oauth2Service
}

type service struct {
	client *Client
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	req = req.WithContext(ctx)

	// TODO: rate limiter
	// rateLimitCategory := category(req.URL.Path)
	//
	// if bypass := ctx.Value(bypassRateLimitCheck); bypass == nil {
	// 	if err := c.checkRateLimitBeforeDo(req, rateLimitCategory); err != nil {
	// 		return &Response{
	// 			Response: err.Response,
	// 			Rate:     err.Rate,
	// 		}, err
	// 	}
	// }

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	response := newResponse(resp)

	// TODO: rate limiter
	// // Don't update the rate limits if this was a cached response.
	// // X-From-Cache is set by https://github.com/gregjones/httpcache
	// if response.Header.Get("X-From-Cache") == "" {
	// 	c.rateMu.Lock()
	// 	c.rateLimits[rateLimitCategory] = response.Rate
	// 	c.rateMu.Unlock()
	// }

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}
	return response, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func NewClient(httpClient *http.Client, options AuthOptions) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	// uploadURL, _ := url.Parse(uploadBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		AppKey:    options.AppKey,
		AppSecret: options.AppSecret,
		AccountNo: options.AccountNo,
	}
	c.common.client = c
	c.DomesticStockQuotations = (*DomesticStockQuotationsService)(&c.common)
	c.Oauth2 = (*Oauth2Service)(&c.common)
	return c
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	// response.populatePageValues()
	// response.Rate = parseRate(r)
	// response.TokenExpiration = parseTokenExpiration(r)
	return response
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	switch {
	// TODO: Advanced error handling?
	default:
		return errorResponse
	}
}

type ErrorResponse struct {
	Response         *http.Response
	ErrorDescription string `json:"error_description"`
	ErrorCode        string `json:"error_code"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.ErrorCode, r.ErrorDescription)
}
