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
	"github.com/Axel791/auth/internal/services/mock"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	repoMock "github.com/Axel791/auth/internal/usecases/auth/repositories/mock"
)

func TestValidateScenario_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repoMock.NewMockUserRepository(ctrl)
	mockTokenService := mock.NewMockTokenService(ctrl)

	scenario := NewValidateScenario(mockUserRepo, mockTokenService)

	ctx := context.Background()
	token := "valid_token"

	// Подготовка
	mockTokenService.
		EXPECT().
		ValidateToken(token).
		Return(dto.ClaimsDTO{UserID: 10, Login: "john_doe"}, nil)

	mockUserRepo.
		EXPECT().
		GetUserById(ctx, int64(10)).
		Return(domains.User{ID: 10, Login: "john_doe", Password: "somehash"}, nil)

	err := scenario.Execute(ctx, token)
	assert.NoError(t, err)
}

func TestValidateScenario_Execute_InvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repoMock.NewMockUserRepository(ctrl)
	mockTokenService := mock.NewMockTokenService(ctrl)

	scenario := NewValidateScenario(mockUserRepo, mockTokenService)

	ctx := context.Background()
	token := "invalid_or_expired_token"

	mockTokenService.
		EXPECT().
		ValidateToken(token).
		Return(dto.ClaimsDTO{}, errors.New("token expired"))

	err := scenario.Execute(ctx, token)
	assert.Error(t, err)

	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusUnauthorized, appErr.Code)
		assert.Contains(t, appErr.Error(), "invalid token")
	} else {
		t.Errorf("expected *appkit.AppError, got %T", err)
	}
}

func TestValidateScenario_Execute_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repoMock.NewMockUserRepository(ctrl)
	mockTokenService := mock.NewMockTokenService(ctrl)

	scenario := NewValidateScenario(mockUserRepo, mockTokenService)

	ctx := context.Background()
	token := "valid_token"

	mockTokenService.
		EXPECT().
		ValidateToken(token).
		Return(dto.ClaimsDTO{UserID: 999}, nil)

	mockUserRepo.
		EXPECT().
		GetUserById(ctx, int64(999)).
		Return(domains.User{}, errors.New("db is down"))

	err := scenario.Execute(ctx, token)
	assert.Error(t, err)

	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusInternalServerError, appErr.Code)
		assert.Contains(t, appErr.Error(), "error getting user")
	} else {
		t.Errorf("expected *appkit.AppError, got %T", err)
	}
}

func TestValidateScenario_Execute_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repoMock.NewMockUserRepository(ctrl)
	mockTokenService := mock.NewMockTokenService(ctrl)

	scenario := NewValidateScenario(mockUserRepo, mockTokenService)

	ctx := context.Background()
	token := "valid_token"

	mockTokenService.
		EXPECT().
		ValidateToken(token).
		Return(dto.ClaimsDTO{UserID: 1000}, nil)

	mockUserRepo.
		EXPECT().
		GetUserById(ctx, int64(1000)).
		Return(domains.User{ID: 0}, nil)

	err := scenario.Execute(ctx, token)
	assert.Error(t, err)

	var appErr *appkit.AppError
	if errors.As(err, &appErr) {
		assert.Equal(t, http.StatusNotFound, appErr.Code)
		assert.Contains(t, appErr.Error(), "user not found")
	} else {
		t.Errorf("expected *appkit.AppError, got %T", err)
	}
}
