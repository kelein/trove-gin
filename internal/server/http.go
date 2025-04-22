package server

import (
	gohttp "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/kelein/trove-gin/docs"
	"github.com/kelein/trove-gin/internal/handler"
	"github.com/kelein/trove-gin/internal/middleware"
	"github.com/kelein/trove-gin/pkg/jwt"
	"github.com/kelein/trove-gin/pkg/log"
	"github.com/kelein/trove-gin/pkg/server/http"
	"github.com/kelein/trove-gin/pkg/version"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	s.GET("/", home)
	s.GET("/ping", home)
	s.GET("/version", home)

	s.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(gohttp.StatusFound, "/swagger/index.html")
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		noAuthRouter.POST("/login", userHandler.Login)
		noAuthRouter.POST("/register", userHandler.Register)

		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return s
}

func home(ctx *gin.Context) {
	ctx.JSON(gohttp.StatusOK, gin.H{
		"app":    version.AppName,
		"pid":    os.Getpid(),
		"build":  version.Info(),
		"uptime": version.Uptime,
	})
}
