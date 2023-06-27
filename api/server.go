package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/murasame29/GolangBackendEx/db/sqlc"
)

// Bankingサービスに関するHTTPリクエストを受け取る
type Server struct {
	store  db.Store
	router *gin.Engine
}

// 新しいHTTPサーバを作り、ルーティングを設定する
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currensy", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server
}

// サーバのリスニングを開始する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error:": err.Error()}
}
