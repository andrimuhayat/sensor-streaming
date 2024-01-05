package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	docs "sensor-streaming/docs" // Import the docs
	"sensor-streaming/internal/module/sensor"
	"sensor-streaming/internal/platform/app"
	module "sensor-streaming/internal/platform/common"
	"sensor-streaming/internal/platform/httpengine"
	"sensor-streaming/internal/platform/httpengine/echoserver"
	"sensor-streaming/internal/platform/messagebroker/mqqt"
	internalMdw "sensor-streaming/internal/platform/middleware"
	"sync"
)

func RunConsumer(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

// Server httpengine service
type Server struct {
	DB        *sqlx.DB
	Router    httpengine.Router
	AppRouter *echo.Echo
}

// NewServer httpengine initialization
func NewServer() (*Server, error) {
	var (
		err        error
		appConfig  app.Config
		server     = new(Server)
		echoRouter = echoserver.NewEchoRouter()
	)

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}

	server.Router = echoRouter
	server.Run(appConfig)

	return server, nil
}

func (s *Server) Run(config app.Config) {
	s.AppRouter = s.Router.GetRouter()
	s.AppRouter.Use(internalMdw.PanicException)
	s.AppRouter.Use(middleware.RequestID())

	docs.SwaggerInfo.Title = "Sensor Streaming(A) API"
	docs.SwaggerInfo.Description = "This is a swagger for Sensor Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8081"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	s.AppRouter.GET("/swagger/*", echoSwagger.WrapHandler)
	s.initModuleDependency(&config)

	s.Router.SERVE(config.App.Port)
}

func (s *Server) initModuleDependency(appConfig *app.Config) module.Dependency {
	var (
		dependency module.Dependency
	)
	dependency.MqttClient = mqqt.Connect("pub", appConfig.Mqtt)
	sensor.StartService(dependency, s.AppRouter)
	return dependency
}
