package claim

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	jwt.StandardClaims
	ID  int   `json:"id"`
	Exp int64 `json:"exp"`
}

func (c *Claim) GetToken(signingString string, expirationTime time.Time) (string, error) {
	c.StandardClaims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(signingString))
}

func GetFromToken(tokenString, signingString string) (*Claim, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	iID, ok := claim["id"]
	if !ok {
		return nil, errors.New("user id not found")
	}

	id, ok := iID.(float64)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	return &Claim{ID: int(id)}, nil
}
