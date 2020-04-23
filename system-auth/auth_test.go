package auth

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	clientName := faker.Name()
	secret := faker.Password()

	auth := NewAuth(clientName, secret)

	token, err := auth.CreateToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestVerifyToken(t *testing.T) {
	clientName := faker.Name()
	secret := faker.Password()

	auth := NewAuth(clientName, secret)
	token, _ := auth.CreateToken()

	t.Run("Success", func(t *testing.T) {
		claims, err := auth.VerifyToken(token)
		if success := assert.NoError(t, err); success {
			assert.Equal(t, claims.Name, clientName)
		}
	})

	t.Run("Failure: invalid token", func(t *testing.T) {
		claims, err := auth.VerifyToken(token + "x")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("Failure: expired", func(t *testing.T) {
		// Excluded because it is practically indefinite period.
	})
}
