package handler

import (
	"net/http"
	"strings"

	"github.com/solumD/WBTech_L0/internal/handler/response"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type getOrderByUIDResponse struct {
	response.Response
	model.Order `json:"order,omitempty"`
}

// GetOrderByUID gets order by its' uid
// @Summary GetOrder
// @Tags order
// @Description gets order by its' uid
// @ID get-order
// @Produce json
// @Param uid path string true "order's uid"
// @Success 200 {object} getOrderByUIDResponse
// @Failure 400,500
// @Router /order/{uid} [get]
func (h *Handler) GetOrderByUID(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(middleware.RequestIDKey).(string)

	uid := strings.TrimSpace(chi.URLParam(r, "uid"))
	if len(uid) == 0 {
		logger.Error("order uid's length can't be 0", zap.String("request id", requestID))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error("uid's length can't be 0"))

		return
	}

	logger.Info("got order uid from url", zap.String("request id", requestID), zap.String("uid", uid))

	order, err := h.orderService.GetOrderByUID(r.Context(), uid)
	if err != nil {
		logger.Error("failed to get order by uid", zap.String("request id", requestID), zap.Error(err))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error("failed to get order by uid"))

		return
	}

	logger.Info("request completed", zap.String("request id", requestID))

	render.JSON(w, r, getOrderByUIDResponse{
		response.OK(),
		order,
	})
}
