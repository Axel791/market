package v1

import (
	"encoding/json"

	"github.com/Axel791/appkit"

	userAPI "github.com/Axel791/auth/internal/rest/v1/api"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	authScenarios "github.com/Axel791/auth/internal/usecases/auth/scenarios"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type LoginHandler struct {
	logger       *log.Logger
	loginUseCase authScenarios.Login
}

func NewLoginHandler(
	logger *log.Logger,
	loginUseCase authScenarios.Login,
) *LoginHandler {
	return &LoginHandler{
		logger:       logger,
		loginUseCase: loginUseCase,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input userAPI.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Infof("err decode body: %v", err)
		appkit.WriteErrorJSON(w, appkit.BadRequestError("invalid request body"))
		return
	}

	loginDTO := dto.UserDTO{
		Login:    input.Login,
		Password: input.Password,
	}

	token, err := h.loginUseCase.Execute(r.Context(), loginDTO)
	if err != nil {
		h.logger.Infof("err login: %v", err)
		appkit.WriteErrorJSON(w, err)
		return
	}
	appkit.WriteJSON(w, http.StatusOK, userAPI.Token{Token: token.Token})
}
