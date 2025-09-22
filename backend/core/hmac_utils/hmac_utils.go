package hmac_utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
)

func HashUUID(raw uuid.UUID, secret []byte) string {
	mac := hmac.New(sha256.New, secret)

	mac.Write(raw[:])

	sum := mac.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(sum)
}
