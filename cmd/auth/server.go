package main

import (
	"strconv"
)

type errNotImplemented string

func (e errNotImplemented) Error() string {
	return "not implemented: " + string(e)
}

type authorizeParam struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (r authResponse) AsFragment() string {
	str := "access_token=" + r.AccessToken
	str += "&token_type=" + r.TokenType
	str += "&expires_id=" + strconv.Itoa(r.ExpiresIn)
	return str
}

type server struct{}

func (s *server) Authorize(param authorizeParam) (url string, err error) {
	if param.ResponseType != "token" {
		return "", errNotImplemented("support only response_type=token")
	}
	r := authResponse{
		AccessToken: "qwertyuiop",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}
	url = param.RedirectURI + "#" + r.AsFragment()
	return
}

func (s *server) Token() {
}
