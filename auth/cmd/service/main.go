package main

import (
	"net/http"

	"github.com/Axel791/auth/internal/config"
	"github.com/Axel791/auth/internal/db"
	"github.com/Axel791/auth/internal/grpc/v1/pb"
	apiV1Handlers "github.com/Axel791/auth/internal/rest/v1"
	"github.com/Axel791/auth/internal/services"
	"github.com/Axel791/auth/internal/usecases/auth/repositories"
	"github.com/Axel791/auth/internal/usecases/auth/scenarios"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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

	// Репозитории
	userRepository := repositories.NewUserRepository(dbConn)

	// Сервисы
	tokenService := services.NewTokenService(cfg.SecretKey)
	passwordService := services.NewHashPasswordService(cfg.PasswordSecret)

	// gRPC
	conn, err := grpc.NewClient(
		cfg.GrpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create new client: %v", err)
	}

	defer conn.Close()

	// gRPC clients
	loyaltyClient := pb.NewLoyaltyServiceClient(conn)

	// UseCases
	validationUseCase := scenarios.NewValidateScenario(userRepository, tokenService)
	loginUseCase := scenarios.NewLoginScenario(userRepository, passwordService, tokenService)
	registrationUseCase := scenarios.NewRegistrationScenario(userRepository, passwordService, loyaltyClient)

	// Routers V1
	router.Route("/public/api/v1", func(r chi.Router) {
		r.Method(
			http.MethodPost,
			"/users/registration",
			apiV1Handlers.NewRegistrationHandler(registrationUseCase, log),
		)
		r.Method(
			http.MethodPost,
			"/users/login",
			apiV1Handlers.NewLoginHandler(log, loginUseCase),
		)
		r.Method(
			http.MethodPost,
			"/users/validate",
			apiV1Handlers.NewValidationHandler(log, validationUseCase),
		)
	})

	log.Infof("server started on %s", cfg.Address)
	err = http.ListenAndServe(cfg.Address, router)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
