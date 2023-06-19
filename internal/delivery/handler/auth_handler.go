package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
