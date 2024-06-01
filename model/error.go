package model

import (
	"net/http"

	errors "github.com/bhumong/dbo-test/pkg/error"
)

var ErrAuthJwt = errors.New(
	"ERR_AUTH_001",
	"unauthorized: invalid jwt",
	"unauthorized: invalid jwt",
	http.StatusUnauthorized,
)

var ErrAuthClaimsNotFound = errors.New(
	"ERR_AUTH_002",
	"CLAIMS_NOT_FOUND",
	"claims not found in context",
	http.StatusInternalServerError,
)
