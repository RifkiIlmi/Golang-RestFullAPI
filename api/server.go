package api

import (
	db "github.com/Rifkiilmi/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.IStore
	router *gin.Engine
}

// Create server and setup routes
func NewServer(store db.IStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)

	router.POST("/transfers", server.CreateTransfer)

	server.router = router

	return server
}

// Runs server with specific address
func (server *Server) Run(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
