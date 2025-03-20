package scenarios

import (
	"context"
	"net/http"

	"github.com/Axel791/appkit"
	"github.com/Axel791/auth/internal/services"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/Axel791/auth/internal/usecases/auth/repositories"
)

type LoginScenario struct {
	userRepository      repositories.UserRepository
	hashPasswordService services.HashPasswordService
	tokenService        services.TokenService
}

func NewLoginScenario(
	userRepository repositories.UserRepository,
	hashPasswordService services.HashPasswordService,
	tokenService services.TokenService,
) *LoginScenario {
	return &LoginScenario{
		userRepository:      userRepository,
		hashPasswordService: hashPasswordService,
		tokenService:        tokenService,
	}
}

func (s *LoginScenario) Execute(ctx context.Context, userDTO dto.UserDTO) (dto.TokenDTO, error) {
	user, err := s.userRepository.GetUserByLogin(ctx, userDTO.Login)
	if err != nil {
		return dto.TokenDTO{}, appkit.WrapError(
			http.StatusInternalServerError,
			"error login user",
			err,
		)
	}

	if user.ID == 0 {
		return dto.TokenDTO{}, appkit.NotFoundError("user does not exist")
	}

	hashedPassword := s.hashPasswordService.Hash(user.Password)

	if hashedPassword != userDTO.Password {
		return dto.TokenDTO{}, appkit.BadRequestError("invalid password")
	}

	claims := dto.ClaimsDTO{
		UserID: user.ID,
		Login:  userDTO.Login,
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		return dto.TokenDTO{}, appkit.WrapError(
			http.StatusNotFound,
			"error user login",
			err,
		)
	}
	return dto.TokenDTO{Token: token}, nil
}
