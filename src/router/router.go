package router

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/madrid-recicla/server/src/handler"
	cors "github.com/rs/cors/wrapper/gin"
)

var clothesHandler = handler.NewClothesHandler()
var mapboxHandler = handler.NewMapboxTokenHandler()

func InitRouter() {
	log.Println("Initializing HTTP router")
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Use(cors.New(cors.Options{
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
	}))

	configureInternalApi(router)
	configureExternalApi(router)

	router.Run(":" + os.Getenv("PORT"))
}

func configureInternalApi(r *gin.Engine) {
	in := r.Group("int/v1/")
	con := in.Group("containers")
	clo := con.Group("clothes")
	{
		clo.POST("/load", clothesHandler.Load())
	}
}

func configureExternalApi(r *gin.Engine) {
	ex := r.Group("v1/")
	// Tokens
	tok := ex.Group("token")
	{
		tok.GET("mapbox", mapboxHandler.GetToken())
	}
	// Containers
	con := ex.Group("containers")
	clo := con.Group("clothes")
	{
		clo.GET("", clothesHandler.ListAll())
	}
}
