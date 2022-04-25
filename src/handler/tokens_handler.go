package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madrid-recicla/server/src/repository"
)

type tokensHandler struct {
	repository repository.TokenRepository
}

func NewMapboxTokenHandler() TokenHandler {
	return &tokensHandler{repository.NewMapboxTokenRepository()}
}

func (h *tokensHandler) GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk, err := h.repository.GetToken()
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error while getting MapBox token", err)
			return
		}
		c.String(http.StatusOK, tk)
	}
}
