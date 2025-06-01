package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Axel791/appkit"
	"github.com/Axel791/loyalty/internal/rest/v1/api"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios"
	log "github.com/sirupsen/logrus"
)

// ConclusionUserBalanceHandler - структура хэндлера по выводу баланса
type ConclusionUserBalanceHandler struct {
	logger                   *log.Logger
	conclusionBalanceUseCase scenarios.ConclusionUserBalanceUseCase
}

func NewConclusionUserBalanceHandler(
	logger *log.Logger,
	conclusionBalanceUseCase scenarios.ConclusionUserBalanceUseCase,
) *ConclusionUserBalanceHandler {
	return &ConclusionUserBalanceHandler{
		logger:                   logger,
		conclusionBalanceUseCase: conclusionBalanceUseCase,
	}
}

func (h *ConclusionUserBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userBalanceInput api.LoyaltyBalance
	if err := json.NewDecoder(r.Body).Decode(&userBalanceInput); err != nil {
		h.logger.Infof("err decode body: %v", err)
		appkit.WriteErrorJSON(w, appkit.BadRequestError("invalid request body"))
		return
	}

	userBalanceDTO := dto.LoyaltyBalance{
		UserID: userBalanceInput.UserID,
		Count:  userBalanceInput.Count,
	}

	err := h.conclusionBalanceUseCase.Execute(r.Context(), userBalanceDTO)
	if err != nil {
		h.logger.Infof("err conlusion balance: %v", err)
		appkit.WriteErrorJSON(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
