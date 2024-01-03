package sensor

import (
	"github.com/labstack/echo/v4"
	"sensor-streaming/internal/module/sensor/handler"
	"sensor-streaming/internal/module/sensor/repository"
	"sensor-streaming/internal/module/sensor/usecase"
	module "sensor-streaming/internal/platform/common"
)

func StartService(dependency module.Dependency, router *echo.Echo) {
	//init repo
	dependency.SensorRepository = repository.NewRepository(dependency.DB)
	//init usecase
	dependency.SensorUseCase = usecase.NewUseCase(dependency.SensorRepository)
	// define handler
	sensorHandler := handler.NewHandler(dependency.SensorUseCase)
	//init route
	versionRoute := router.Group("/api")

	handler.NewSensorRoute(sensorHandler, versionRoute)
}
