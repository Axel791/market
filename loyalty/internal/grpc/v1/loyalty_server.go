package v1

import (
	"context"
	"fmt"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"

	"github.com/Axel791/loyalty/internal/grpc/v1/pb"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoyaltyServer struct {
	pb.UnimplementedLoyaltyServiceServer

	createLoyaltyBalanceUseCase scenarios.CreateUserBalanceUseCase
	inputLoyaltyBalanceUseCase  scenarios.InputUserBalanceUseCase
}

func NewLoyaltyServer(
	createLoyaltyBalanceUseCase scenarios.CreateUserBalanceUseCase,
	inputLoyaltyBalanceUseCase scenarios.InputUserBalanceUseCase,
) *LoyaltyServer {
	return &LoyaltyServer{
		createLoyaltyBalanceUseCase: createLoyaltyBalanceUseCase,
		inputLoyaltyBalanceUseCase:  inputLoyaltyBalanceUseCase,
	}
}

// CreateLoyaltyBalance - Реализация метода CreateLoyaltyBalance из .proto
func (s *LoyaltyServer) CreateLoyaltyBalance(
	ctx context.Context,
	req *pb.CreateLoyaltyBalanceRequest,
) (*pb.CreateLoyaltyBalanceResponse, error) {
	err := s.createLoyaltyBalanceUseCase.Execute(ctx, req.GetUserId())
	if err != nil {
		return &pb.CreateLoyaltyBalanceResponse{
				Success:      false,
				ErrorMessage: fmt.Sprintf("failed to create loyalty balance: %v", err),
			},
			status.Errorf(codes.Internal, "failed to create loyalty balance: %v", err)
	}

	return &pb.CreateLoyaltyBalanceResponse{
		Success: true,
	}, nil
}

func (s *LoyaltyServer) ConclusionLoyaltyBalance(
	ctx context.Context,
	req *pb.ConcludeUserBalanceRequest,
) (*pb.ConcludeUserBalanceResponse, error) {
	var loyaltyBalance dto.LoyaltyBalance

	loyaltyBalance = dto.LoyaltyBalance{
		UserID: req.UserId,
		Count:  req.Count,
	}

	err := s.inputLoyaltyBalanceUseCase.Execute(ctx, req.OrderId, loyaltyBalance)
	if err != nil {
		return &pb.ConcludeUserBalanceResponse{
				Success:      false,
				ErrorMessage: fmt.Sprintf("failed to input loyalty balance: %v", err),
			},
			status.Errorf(codes.Internal, "failed to input loyalty balance: %v", err)
	}
	return &pb.ConcludeUserBalanceResponse{
		Success: true,
	}, nil
}
