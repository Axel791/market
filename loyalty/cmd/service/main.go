package main

import (
	"net"
	"net/http"

	"github.com/Axel791/loyalty/internal/config"
	"github.com/Axel791/loyalty/internal/db"
	v1 "github.com/Axel791/loyalty/internal/grpc/v1"
	"github.com/Axel791/loyalty/internal/grpc/v1/pb"
	apiV1Handler "github.com/Axel791/loyalty/internal/rest/v1"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	log.Infof("databse_dsn: %s", cfg.DatabaseDSN)
	dbConn, err := db.ConnectDB(cfg.DatabaseDSN, cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer func() {
		if dbConn != nil {
			_ = dbConn.Close()
		}
	}()

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Logger)

	grpcServer := grpc.NewServer()

	// Repositories
	loyaltyBalanceRepository := repositories.NewSqlLoyaltyRepository(dbConn)
	loyaltyHistoryRepository := repositories.NewSqlLoyaltyHistoryRepository(dbConn)
	unitOfWork := repositories.NewUnitOfWork(dbConn)

	// UseCases
	createUserBalance := scenarios.NewCreateLoyaltyBalance(loyaltyBalanceRepository)
	conclusionUserBalance := scenarios.NewConclusionUserBalance(
		loyaltyBalanceRepository,
		loyaltyHistoryRepository,
		unitOfWork,
	)
	inputUserBalanceUseCase := scenarios.NewInputUserBalance(
		loyaltyBalanceRepository,
		loyaltyHistoryRepository,
		unitOfWork,
	)
	getUserBalance := scenarios.NewGetUserBalance(loyaltyBalanceRepository)

	// gRPC
	loyaltyServer := v1.NewLoyaltyServer(createUserBalance, inputUserBalanceUseCase)

	pb.RegisterLoyaltyServiceServer(grpcServer, loyaltyServer)

	lis, err := net.Listen(cfg.GrpcNetwork, cfg.GrpcAddress)
	if err != nil {
		log.Fatalf("failed to listen on :50051: %v", err)
	}
	log.Println("Starting Loyalty gRPC server on", cfg.GrpcAddress)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
	// REST
	router.Route("/api/v1", func(r chi.Router) {
		r.Method(
			http.MethodPost,
			"/loyalty/balance",
			apiV1Handler.NewGetUserBalanceHandler(log, getUserBalance),
		)
		r.Method(
			http.MethodPost,
			"/loyalty/balance/conclusion",
			apiV1Handler.NewConclusionUserBalanceHandler(log, conclusionUserBalance),
		)
	})

	log.Infof("server started on %s", cfg.Address)
	err = http.ListenAndServe(cfg.Address, router)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
