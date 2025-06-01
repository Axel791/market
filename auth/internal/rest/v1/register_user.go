package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Axel791/appkit"

	userAPI "github.com/Axel791/auth/internal/rest/v1/api"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	authScenarios "github.com/Axel791/auth/internal/usecases/auth/scenarios"
	log "github.com/sirupsen/logrus"
)

type RegistrationHandler struct {
	logger              *log.Logger
	registrationUseCase authScenarios.Registration
}

func NewRegistrationHandler(
	registrationUseCase authScenarios.Registration,
	logger *log.Logger,
) *RegistrationHandler {
	return &RegistrationHandler{
		registrationUseCase: registrationUseCase,
		logger:              logger,
	}
}

func (h *RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input userAPI.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Infof("err decode body: %v", err)
		appkit.WriteErrorJSON(w, appkit.BadRequestError("invalid request body"))
		return
	}
	userDTO := dto.UserDTO{
		Login:    input.Login,
		Password: input.Password,
	}

	err := h.registrationUseCase.Execute(r.Context(), userDTO)
	if err != nil {
		h.logger.Infof("err login: %v", err)
		appkit.WriteErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
