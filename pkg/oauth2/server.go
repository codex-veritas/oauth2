package oauth2

import (
	"strconv"
)

// ErrNotImplemented is return by a server when an oauth2 option is not implemented.
type ErrNotImplemented string

// Error displays the error as a string.
func (e ErrNotImplemented) Error() string {
	return "not implemented: " + string(e)
}

// AuthorizeParam contains the required parameters for an authorize request.
type AuthorizeParam struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

// AuthorizeResponse contains the fields return to the client after an authorized request.
type AuthorizeResponse struct {
	State            string `json:"state"`
	AccessToken      string `json:"access_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	Error            int    `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// AsFragment returns the AuthorizeResponse as an URL fragment.
func (r AuthorizeResponse) AsFragment() string {
	str := "access_token=" + r.AccessToken
	str += "&token_type=" + r.TokenType
	str += "&expires_id=" + strconv.Itoa(r.ExpiresIn)
	if r.Error != 0 {
		str = "error=" + strconv.Itoa(r.Error)
		str += "&error_description=" + r.ErrorDescription
	}
	str += "&state=" + r.State
	return str
}

// Server represents an OAuth2 server.
type Server struct{}

// Authorize validate the client authentification and return the URL on which he must be redirected.
func (s *Server) Authorize(param AuthorizeParam) (url string, err error) {
	if param.ResponseType != "token" {
		return "", ErrNotImplemented("support only response_type=token")
	}
	r := AuthorizeResponse{
		State:       param.State,
		AccessToken: "qwertyuiop",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}
	url = param.RedirectURI + "#" + r.AsFragment()
	return
}

// Token transform an authorisation code into an access token.
// Not implemented.
func (s *Server) Token() {
}
