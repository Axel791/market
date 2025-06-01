package scenarios

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Axel791/appkit"
	"github.com/Axel791/auth/internal/grpc/v1/pb"

	"github.com/Axel791/auth/internal/domains"
	"github.com/Axel791/auth/internal/services"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/Axel791/auth/internal/usecases/auth/repositories"
)

// RegistrationScenario - структура сценария регистрация
type RegistrationScenario struct {
	userRepository      repositories.UserRepository
	hashPasswordService services.HashPasswordService
	loyaltyClient       pb.LoyaltyServiceClient
}

// NewRegistrationScenario - создание сценария
func NewRegistrationScenario(
	userRepository repositories.UserRepository,
	hashPasswordService services.HashPasswordService,
	loyaltyClient pb.LoyaltyServiceClient,
) *RegistrationScenario {
	return &RegistrationScenario{
		userRepository:      userRepository,
		hashPasswordService: hashPasswordService,
		loyaltyClient:       loyaltyClient,
	}
}

// Execute - Функция выполнения сценария
func (s *RegistrationScenario) Execute(ctx context.Context, userDTO dto.UserDTO) error {
	userDomain := domains.User{
		Login:    userDTO.Login,
		Password: userDTO.Password,
	}

	if err := userDomain.ValidatePassword(); err != nil {
		return appkit.ValidationError(err.Error())
	}

	if err := userDomain.ValidateLogin(); err != nil {
		return appkit.ValidationError(err.Error())
	}

	user, err := s.userRepository.GetUserByLogin(ctx, userDomain.Login)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return appkit.WrapError(
				http.StatusInternalServerError,
				"error getting user by login",
				err,
			)
		}
	}

	if user.ID > 0 {
		return appkit.BadRequestError("user login already exists")
	}

	hashedPassword := s.hashPasswordService.Hash(userDomain.Password)
	userDomain.Password = hashedPassword

	createdUser, err := s.userRepository.CreateUser(ctx, userDomain)
	if err != nil {
		return appkit.WrapError(http.StatusInternalServerError, "error creating user", err)
	}

	resp, err := s.loyaltyClient.CreateLoyaltyBalance(ctx, &pb.CreateLoyaltyBalanceRequest{
		UserId: createdUser.ID,
	})
	if err != nil {
		return fmt.Errorf("loyalty service RPC error: %w", err)
	}
	if !resp.Success {
		return fmt.Errorf("loyalty service error: %s", resp.ErrorMessage)
	}
	return nil
}
