package app

import (
	"github.com/Kungfucoding23/bookstore_oauth_api/src/clients/cassandra"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/http"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/repository/db"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	session := cassandra.GetSession()
	session.Close()
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run()
}
