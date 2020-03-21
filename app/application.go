package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/clients/cassandra"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/accesstoken"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/http"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/repository/db"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/repository/rest"
)

var (
	router = gin.Default()
)

//StartApplication app init
func StartApplication() {
	session := cassandra.GetSession()
	atService := accesstoken.NewService(db.NewRepository(session), rest.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access/:token_id", atHandler.GetByID)
	router.POST("/oauth/access", atHandler.Create)

	router.Run(":8081")
}
