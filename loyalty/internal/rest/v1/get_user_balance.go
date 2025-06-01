package v1

import (
	"net/http"
	"strconv"

	"github.com/Axel791/appkit"
	"github.com/Axel791/loyalty/internal/rest/v1/api"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios"
	log "github.com/sirupsen/logrus"
)

// GetUserBalanceHandler - получение баланса пользователя
type GetUserBalanceHandler struct {
	logger                *log.Logger
	getUserBalanceUseCase scenarios.GetUserBalanceUseCase
}

func NewGetUserBalanceHandler(
	logger *log.Logger,
	getUserBalanceUseCase scenarios.GetUserBalanceUseCase,
) *GetUserBalanceHandler {
	return &GetUserBalanceHandler{
		logger:                logger,
		getUserBalanceUseCase: getUserBalanceUseCase,
	}
}

func (h *GetUserBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Параметр id не указан", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}

	balanceDTO, err := h.getUserBalanceUseCase.Execute(r.Context(), userID)
	if err != nil {
		h.logger.Infof("err get balance: %v", err)
		appkit.WriteErrorJSON(w, err)
		return
	}
	balanceResponse := api.LoyaltyBalance{
		ID:     balanceDTO.ID,
		UserID: balanceDTO.UserID,
		Count:  balanceDTO.Count,
	}
	appkit.WriteJSON(w, http.StatusOK, balanceResponse)
}
