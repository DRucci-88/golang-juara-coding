package auth

import "strings"

// Authenticator endefinisikan interface kontrak otorisasi (Exported)
type Authenticator interface {
	ValidateToken(token string) bool
}

// merchantAuth adalah struct konkret (Private)
type merchantAuth struct {
	secretKey string
}

// NewAuthenticator bergungsi sebagai Constructor Function
func NewAuthenticator(key string) Authenticator {
	return &merchantAuth{
		secretKey: strings.TrimSpace(key),
	}
}

// ValidateToken mengimplementasikan interface Authenticator
func (ma *merchantAuth) ValidateToken(token string) bool {
	return ma.secretKey == strings.TrimSpace(token)
}
