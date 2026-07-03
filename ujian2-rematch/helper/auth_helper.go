package helper

import (
	"time"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Cost 12 adalah standar industri yang aman dan efisien
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(
	userId uint,
	employeeId *uint,
	email string,
	role model.UserRole,
) (string, error) {
	claims := dto.JWTClaims{
		UserID:     userId,
		EmployeeID: employeeId,
		Email:      email,
		Role:       role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(dto.JWTSecretKey)
}
