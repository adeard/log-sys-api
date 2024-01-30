package main

import (
	"log"
	"log-sys-api/config"
	"log-sys-api/docs"
	"log-sys-api/modules/application"
	"log-sys-api/modules/logging"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Host localhost:6969
// @title API SWAGGER FOR LOGGING SYSTEM API SERVICE
// @version 1.0.0
// @description LOGGING SYSTEM API SERVICE
// @termsOfService http://swagger.io/terms/

// @contact.name ADE ARDIAN
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath

func main() {
	db := config.Connect()

	docs.SwaggerInfo.BasePath = "/log-sys"

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.GET("log-sys/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":         "Logging System API Service",
			"documentation": "/swagger/index.html",
		})
	})

	router.GET("log-sys/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("log-sys/api/v1")

	logging.NewLoggingHandler(v1, logging.LoggingRegistry(db))
	application.NewApplicationHandler(v1, application.ApplicationRegistry(db))

	// router.Run(":86")

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "87"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	// Menampilkan log koneksi sukses
	log.Println("App Service run in port:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// Menampilkan log ketika koneksi gagal
		log.Fatal("Connection Fail -> port "+port+":", err)
	}
}
