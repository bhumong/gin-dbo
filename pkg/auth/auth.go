package auth

import (
	"time"

	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type AuthJwt struct {
	secret           string
	expiredDurration time.Duration
	signingMethod    *jwt.SigningMethodHMAC
	tokenOptions     []jwt.TokenOption
}

func New(secret string,
	expiredDurration time.Duration,
	signingMethod *jwt.SigningMethodHMAC,
	tokenOptions []jwt.TokenOption) *AuthJwt {
	return &AuthJwt{
		secret:           secret,
		expiredDurration: expiredDurration,
		signingMethod:    signingMethod,
		tokenOptions:     tokenOptions,
	}
}

func (a *AuthJwt) Validate(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("method not jwt.SigningMethodHMAC")
		} else if method != a.signingMethod {
			return nil, errors.New("method not same a.signingMethod")
		}
		return ([]byte(a.secret)), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt.Parse")
	}

	if !token.Valid {
		return nil, errors.New("!token.Valid")
	}
	return token.Claims, nil
}

func (a *AuthJwt) NewWithClaims(claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(a.signingMethod, claims, a.tokenOptions...)
	if jwtToken == nil {
		return "", errors.New("jwt.NewWithClaims is nil")
	}
	stringToken, err := jwtToken.SignedString([]byte(a.secret))
	if err != nil {
		return "", errors.Wrap(err, "jwtToken.SigningString return an error")
	}
	return stringToken, nil
}
