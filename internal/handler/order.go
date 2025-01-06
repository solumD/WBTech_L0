package handler

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/solumD/WBTech_L0/internal/logger"
	"github.com/solumD/WBTech_L0/internal/model"
	"github.com/solumD/WBTech_L0/internal/response"
	"go.uber.org/zap"
)

type getOrderByUIDResponse struct {
	response.Response
	model.Order `json:"order"`
}

// GetOrderByUID gets order by its' uid
func (h *Handler) GetOrderByUID(w http.ResponseWriter, r *http.Request) {
	uid := strings.TrimSpace(chi.URLParam(r, "uid"))

	if len(uid) == 0 {
		logger.Error("order uid's length can't be 0")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error("uid's length can't be 0"))

		return
	}

	logger.Info("got order uid from url", zap.String("uid", uid))

	order, err := h.orderService.GetOrderByUID(r.Context(), uid)
	if err != nil {
		logger.Error("failed to get order by uid", zap.Error(err))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error("failed to get order by uid"))

		return
	}

	render.JSON(w, r, getOrderByUIDResponse{
		response.OK(),
		order,
	})
}
