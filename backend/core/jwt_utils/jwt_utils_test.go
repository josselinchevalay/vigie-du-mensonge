package jwt_utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJwtUtils(t *testing.T) {
	inputID := uuid.New()

	dummySecret := []byte("dummySecret")

	jwt, err := GenerateJWT(inputID, dummySecret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	parsedID, err := ParseJWT(jwt, dummySecret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, inputID, parsedID)
}

func TestJwtUtils_WrongSecret(t *testing.T) {
	id := uuid.New()

	jwt, err := GenerateJWT(id, []byte("dummySecret"), time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseJWT(jwt, []byte("wrongSecret"))
	assert.Error(t, err)
}

func TestJwtUtils_Expired(t *testing.T) {
	id := uuid.New()

	jwt, err := GenerateJWT(id, []byte("dummySecret"), time.Now().Add(-time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseJWT(jwt, []byte("dummySecret"))

	assert.Error(t, err)
}
