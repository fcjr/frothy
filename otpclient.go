package frothy

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

type TOTP struct {
	Code      string
	ExpiresAt time.Time
}

// NewTOTP Generates a TOTP from the given secret string
// currently only supports the most common type of TOTP:
// Six Digit, SHA1 based, 30 second intervals
//
// follows algo from https://tools.ietf.org/html/rfc4226
func NewTOTP(secret string) (*TOTP, error) {

	hashFunc := sha1.New  // TODO support mutliple hash functions
	intervalSeconds := 30 // TODO support multiple intervals
	codeLen := 6          // TODO support multiple code lengths

	now := time.Now() // TODO support timestamps?
	key := uint64(float64(now.Unix()) / float64(intervalSeconds))

	// transform invalid but fixable secrets
	// 1) trim secret
	secret = strings.TrimSpace(secret)

	// 2) ensure capitalization
	secret = strings.ToUpper(secret)

	// 3) ensure proper padding
	secretBytes, err := base32.StdEncoding.WithPadding('=').DecodeString(secret)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 8)
	mac := hmac.New(hashFunc, secretBytes)
	binary.BigEndian.PutUint64(buf, key)

	mac.Write(buf)
	sum := mac.Sum(nil)

	// https://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[19] & 0xf
	binCode := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	code := int32(binCode % int64(math.Pow10(codeLen)))

	return &TOTP{
		Code:      fmt.Sprintf(fmt.Sprintf("%0%%dd", codeLen), code), // pad beginning with zeros
		ExpiresAt: now.Truncate(time.Second * time.Duration(intervalSeconds)).Add(time.Second * time.Duration(intervalSeconds)),
	}, nil
}
