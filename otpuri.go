package frothy

import (
	"fmt"
	"net/url"
	"strings"
)

type OTPType string

const (
	OTPTypeTOTP = "totp"
	OTPTypeHOTP = "hotp"
)

type OTPSecret struct {
	Name    string
	Secret  string
	Issuer  string
	Type    OTPType
	Counter string
}

func ParseOTPSecretFromURI(uri string) (*OTPSecret, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "otpauth" {
		return nil, fmt.Errorf("invalid scheme: %s", u.Scheme)
	}

	query := u.Query()
	secret := query.Get("secret")
	if secret == "" {
		return nil, fmt.Errorf("secret not found")
	}

	var counter string
	switch u.Host {
	case OTPTypeTOTP:
		// no special cases needed
	case OTPTypeHOTP:
		counter = query.Get("counter")
		if counter == "" {
			return nil, fmt.Errorf("htop counter not found but required")
		}
	default:
		return nil, fmt.Errorf("unknown otp type: %s", u.Host)
	}

	return &OTPSecret{
		Name:    strings.TrimPrefix(u.Path, "/"),
		Secret:  secret,
		Issuer:  query.Get("issuer"),
		Counter: counter,
		Type:    OTPType(u.Host),
	}, nil
}
