package api

import (
	db "bankapp/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Create a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs HTTP server on a specific adress
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
