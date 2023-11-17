package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) cardNew(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "card_modal.html.tmpl", nil)
}
