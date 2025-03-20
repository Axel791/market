package services

import "github.com/Axel791/auth/internal/usecases/auth/dto"

type HashPasswordService interface {
	Hash(string) string
}

type TokenService interface {
	GenerateToken(claimsDTO dto.ClaimsDTO) (string, error)
	ValidateToken(tokenStr string) (dto.ClaimsDTO, error)
}
