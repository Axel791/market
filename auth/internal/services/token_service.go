package services

import (
	"fmt"
	"time"

	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/golang-jwt/jwt/v4"
)

// TokenServiceHandler - структура сервиса работы с токеном
type TokenServiceHandler struct {
	secretKey string
}

// NewTokenService - конструктор сервиса работы с токеном
func NewTokenService(secretKey string) *TokenServiceHandler {
	return &TokenServiceHandler{secretKey: secretKey}
}

// GenerateToken - генерация токена
func (s *TokenServiceHandler) GenerateToken(claimsDTO dto.ClaimsDTO) (string, error) {
	claims := jwt.MapClaims{
		"userID": claimsDTO.UserID,
		"login":  claimsDTO.Login,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	return signedToken, nil
}

// ValidateToken - валидация токена
func (s *TokenServiceHandler) ValidateToken(tokenStr string) (dto.ClaimsDTO, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return dto.ClaimsDTO{}, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return dto.ClaimsDTO{}, fmt.Errorf("token has expired")
			}
		}

		userIDFloat, ok := claims["userID"].(float64)
		if !ok {
			return dto.ClaimsDTO{}, fmt.Errorf("invalid userID in token")
		}

		login, ok := claims["login"].(string)
		if !ok {
			return dto.ClaimsDTO{}, fmt.Errorf("invalid login in token")
		}

		return dto.ClaimsDTO{
			UserID: int64(userIDFloat),
			Login:  login,
		}, nil
	}

	return dto.ClaimsDTO{}, fmt.Errorf("invalid token")
}
