package hmac_utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestHmacUtils(t *testing.T) {
	secret := []byte("secret")

	u := uuid.New()

	h1 := HashUUID(u, secret)
	h2 := HashUUID(u, secret)

	if h1 != h2 {
		t.Fatalf("Hashes do not match: %s != %s", h1, h2)
	}

	h3 := HashUUID(uuid.New(), []byte("wrong-secret"))
	if h1 == h3 {
		t.Fatalf("Hashes match with wrong secret: %s == %s", h1, h3)
	}
}
