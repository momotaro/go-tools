package auth

import (
	"math"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims :
type Claims struct {
	Name string
}

// Auth :
type Auth interface {
	CreateToken() (string, error)
	VerifyToken(string) (*Claims, error)
}

// NewAuth :
func NewAuth(name string, secret string) Auth {
	parser := new(jwt.Parser)
	return &auth{name, secret, parser}
}

type auth struct {
	name      string
	secret    string
	jwtParser *jwt.Parser
}

func (a *auth) buildSecret() []byte {
	return []byte(a.secret)
}

// CreateToken :
func (a *auth) CreateToken() (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": a.name,
		"exp": time.Now().Unix() + int64(math.MaxInt32*10),
	})

	token, err := claims.SignedString(a.buildSecret())
	if err != nil {
		return "", err
	}

	return token, nil
}

// VerifyToken :
func (a *auth) VerifyToken(token string) (*Claims, error) {
	t, err := a.jwtParser.ParseWithClaims(
		token,
		jwt.MapClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return a.buildSecret(), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if err := t.Claims.Valid(); err != nil {
		return nil, err
	}

	claims := &Claims{
		Name: t.Claims.(jwt.MapClaims)["sub"].(string),
	}

	return claims, nil
}
