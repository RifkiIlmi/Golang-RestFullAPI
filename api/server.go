package api

import (
	"fmt"

	db "github.com/Rifkiilmi/simplebank/db/sqlc"
	"github.com/Rifkiilmi/simplebank/token"
	"github.com/Rifkiilmi/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config util.Config
	store  db.IStore
	token  token.Maker
	router *gin.Engine
}

// Create server and setup routes
func NewServer(config util.Config, store db.IStore) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %v", err)
	}

	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.token))

	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts/:id", server.GetAccount)
	authRoutes.GET("/accounts", server.ListAccount)

	authRoutes.POST("/transfers", server.CreateTransfer)

	server.router = router
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
