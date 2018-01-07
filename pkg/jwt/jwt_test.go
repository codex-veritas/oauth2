package jwt

import (
	"testing"
)

func TestJwtRoundTrip(t *testing.T) {
	secret := "abcd1234"
	userName := "Sam"
	payload := &Payload{userName}

	token, err := Encode(payload, secret)
	if err != nil {
		t.Fatal("Cannot encode token")
	}

	actualPayload, err := Decode(token, secret)
	if err != nil {
		t.Fatal("Cannot decode token")
	}

	if *actualPayload != *payload {
		t.Fatal("Wrong userName")
	}
}

func TestJwtWrongSecret(t *testing.T) {
	token, err := Encode(&Payload{"sam"}, "secret1")
	if err != nil {
		t.Fatal("Cannot encode token")
	}

	_, err = Decode(token, "secret2")
	if err == nil {
		t.Fatalf("Decoding token should fail: %q", err)
	}

	if err.Error() != "invalid JWT signature" {
		t.Fatal("Wrong error")
	}
}
