package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type header struct {
	Type string `json:"typ"`
	Algo string `json:"alg"`
}

type Payload struct {
	UserName string `json:"user_name"`
	//
	// TODO Extend
}

// Encode takes a userName and secret and returns a full JWT
func Encode(payload *Payload, secret string) (string, error) {
	header := header{
		Algo: "HS256",
		Type: "JWT",
	}

	jsonHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	encodedHeader := base64.StdEncoding.EncodeToString(jsonHeader)
	encodedPayload := base64.StdEncoding.EncodeToString(jsonPayload)

	signature := sign(encodedHeader, encodedPayload, secret)

	return encodedHeader + "." + encodedPayload + "." + signature, nil
}

func sign(encodedHeader, encodedPayload, secret string) string {
	message := encodedHeader + "." + encodedPayload
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	macSum := mac.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(macSum)
	return signature
}

// Decode takes a token and a checks it has been signed by a given secret
// It everything matches, it returns the UserName contained inside the token
func Decode(token, secret string) (*Payload, error) {
	parts := strings.Split(token, ".")
	encodedHeader := parts[0]
	encodedPayload := parts[1]
	signature := parts[2]

	expectedSignature := sign(encodedHeader, encodedPayload, secret)
	if expectedSignature != signature {
		return nil, errors.New("invalid JWT signature")
	}

	// Ignoring the header
	jsonPayload, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		return nil, errors.New("payload is not valid base64")
	}

	var payload Payload
	json.Unmarshal([]byte(jsonPayload), &payload)

	if payload.UserName == "" {
		return nil, errors.New("missing UserName in JWT")
	}

	return &payload, nil
}
