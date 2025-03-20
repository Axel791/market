package scenarios

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Axel791/appkit"
	"github.com/Axel791/auth/internal/domains"
	"github.com/Axel791/auth/internal/usecases/auth/dto"

	"github.com/Axel791/auth/internal/services/mock"
	mockRepo "github.com/Axel791/auth/internal/usecases/auth/repositories/mock"
)

func TestLoginScenario_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockHashSvc := mock.NewMockHashPasswordService(ctrl)
	mockTokenSvc := mock.NewMockTokenService(ctrl)

	scenario := NewLoginScenario(mockUserRepo, mockHashSvc, mockTokenSvc)

	ctx := context.Background()
	userDTO := dto.UserDTO{
		Login:    "some_login",
		Password: "user_input_password",
	}

	dbUser := domains.User{
		ID:       10,
		Login:    "some_login",
		Password: "hashed_from_DB",
	}

	mockUserRepo.
		EXPECT().
		GetUserByLogin(ctx, "some_login").
		Return(dbUser, nil)

	mockHashSvc.
		EXPECT().
		Hash("hashed_from_DB").
		Return("user_input_password")

	mockTokenSvc.
		EXPECT().
		GenerateToken(dto.ClaimsDTO{
			UserID: 10,
			Login:  "some_login",
		}).
		Return("some_jwt_token", nil)

	tokenDTO, err := scenario.Execute(ctx, userDTO)

	assert.NoError(t, err)
	assert.Equal(t, "some_jwt_token", tokenDTO.Token)
}

func TestLoginScenario_Execute_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockHashSvc := mock.NewMockHashPasswordService(ctrl)
	mockTokenSvc := mock.NewMockTokenService(ctrl)

	scenario := NewLoginScenario(mockUserRepo, mockHashSvc, mockTokenSvc)

	ctx := context.Background()
	userDTO := dto.UserDTO{
		Login:    "some_login",
		Password: "123",
	}

	mockUserRepo.
		EXPECT().
		GetUserByLogin(ctx, "some_login").
		Return(domains.User{}, errors.New("db connection error"))

	tokenDTO, err := scenario.Execute(ctx, userDTO)

	assert.Error(t, err)

	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusInternalServerError, appErr.Code)
		assert.Contains(t, appErr.Error(), "error internal")
	} else {
		t.Errorf("expected error of type *appkit.AppError, got %T", err)
	}
	assert.Empty(t, tokenDTO.Token)
}

func TestLoginScenario_Execute_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockHashSvc := mock.NewMockHashPasswordService(ctrl)
	mockTokenSvc := mock.NewMockTokenService(ctrl)

	scenario := NewLoginScenario(mockUserRepo, mockHashSvc, mockTokenSvc)

	ctx := context.Background()
	userDTO := dto.UserDTO{
		Login:    "unknown_user",
		Password: "any_password",
	}

	mockUserRepo.
		EXPECT().
		GetUserByLogin(ctx, "unknown_user").
		Return(domains.User{ID: 0}, nil)

	tokenDTO, err := scenario.Execute(ctx, userDTO)
	assert.Error(t, err)

	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusNotFound, appErr.Code)
		assert.Contains(t, appErr.Error(), "user does not exist")
	} else {
		t.Errorf("expected error of type *appkit.AppError, got %T", err)
	}
	assert.Empty(t, tokenDTO.Token)
}

func TestLoginScenario_Execute_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockHashSvc := mock.NewMockHashPasswordService(ctrl)
	mockTokenSvc := mock.NewMockTokenService(ctrl)

	scenario := NewLoginScenario(mockUserRepo, mockHashSvc, mockTokenSvc)

	ctx := context.Background()
	userDTO := dto.UserDTO{
		Login:    "some_login",
		Password: "entered_password",
	}

	dbUser := domains.User{
		ID:       10,
		Login:    "some_login",
		Password: "hashed_from_DB",
	}

	mockUserRepo.
		EXPECT().
		GetUserByLogin(ctx, "some_login").
		Return(dbUser, nil)

	mockHashSvc.
		EXPECT().
		Hash("hashed_from_DB").
		Return("some_other_hash")

	tokenDTO, err := scenario.Execute(ctx, userDTO)

	assert.Error(t, err)
	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusBadRequest, appErr.Code)
		assert.Contains(t, appErr.Error(), "error bad request login")
	} else {
		t.Errorf("expected error of type *appkit.AppError, got %T", err)
	}
	assert.Empty(t, tokenDTO.Token)
}

func TestLoginScenario_Execute_TokenGenerationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockHashSvc := mock.NewMockHashPasswordService(ctrl)
	mockTokenSvc := mock.NewMockTokenService(ctrl)

	scenario := NewLoginScenario(mockUserRepo, mockHashSvc, mockTokenSvc)

	ctx := context.Background()
	userDTO := dto.UserDTO{
		Login:    "some_login",
		Password: "ok_password",
	}

	dbUser := domains.User{
		ID:       10,
		Login:    "some_login",
		Password: "hashed_from_DB",
	}

	mockUserRepo.
		EXPECT().
		GetUserByLogin(ctx, "some_login").
		Return(dbUser, nil)

	mockHashSvc.
		EXPECT().
		Hash("hashed_from_DB").
		Return("ok_password") // Пароли совпадают

	// Генерация токена упадёт с ошибкой
	mockTokenSvc.
		EXPECT().
		GenerateToken(gomock.Any()).
		Return("", errors.New("token service error"))

	tokenDTO, err := scenario.Execute(ctx, userDTO)
	assert.Error(t, err)
	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusInternalServerError, appErr.Code)
		assert.Contains(t, appErr.Error(), "internal error")
	} else {
		t.Errorf("expected error of type *appkit.AppError, got %T", err)
	}
	assert.Empty(t, tokenDTO.Token)
}
