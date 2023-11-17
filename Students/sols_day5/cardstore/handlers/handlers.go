package handlers

import (
	"local/hw/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", server.cardsList)
	router.GET("/card-modal", server.cardNew)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run(":3000")
}
