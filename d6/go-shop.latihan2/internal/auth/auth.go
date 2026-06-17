package auth

type Authenticator interface {
	ValidateToken(token string) bool
}

type MerchantAuth struct {
	tokenKey string
}

func (auth *MerchantAuth) ValidateToken(token string) bool {
	return auth.tokenKey == token
}

func NewAuthenticator(secretKey string) Authenticator {
	return &MerchantAuth{
		tokenKey: secretKey,
	}
}
