package jwt_utils

import (
	"testing"
	"time"
	"vdm/core/locals"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJwtUtils(t *testing.T) {
	input := locals.AuthedUser{ID: uuid.New(), Email: "test@email.com"}

	dummySecret := []byte("dummySecret")

	jwt, err := GenerateJWT(input, dummySecret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	output, err := ParseJWT(jwt, dummySecret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, input, output)
}

func TestJwtUtils_WrongSecret(t *testing.T) {
	input := locals.AuthedUser{ID: uuid.New(), Email: "test@email.com"}

	jwt, err := GenerateJWT(input, []byte("dummySecret"), time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseJWT(jwt, []byte("wrongSecret"))
	assert.Error(t, err)
}

func TestJwtUtils_Expired(t *testing.T) {
	input := locals.AuthedUser{ID: uuid.New(), Email: "test@email.com"}

	jwt, err := GenerateJWT(input, []byte("dummySecret"), time.Now().Add(-time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseJWT(jwt, []byte("dummySecret"))

	assert.Error(t, err)
}
