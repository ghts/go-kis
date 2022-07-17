package kis

import (
	"context"
	"net/http"
)

type Oauth2Service service

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	AppKey    string `json:"appkey"`
	AppSecret string `json:"appsecret"`
}

func (s *Oauth2Service) TokenP(ctx context.Context) (*TokenResponse, *Response, error) {
	u := "oauth2/tokenP"

	tokenRequest := &TokenRequest{
		AppSecret: s.client.AppSecret,
		AppKey:    s.client.AppKey,
		GrantType: "client_credentials",
	}

	req, err := s.client.NewRequest(http.MethodPost, u, tokenRequest)
	if err != nil {
		return nil, nil, err
	}

	tokenResponse := new(TokenResponse)
	resp, err := s.client.Do(ctx, req, tokenResponse)
	if err != nil {
		return nil, resp, err
	}

	return tokenResponse, resp, nil
}
