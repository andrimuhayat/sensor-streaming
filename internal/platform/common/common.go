package common

import (
	"github.com/jmoiron/sqlx"
	"sensor-streaming/internal/module/sensor/repository"
	"sensor-streaming/internal/module/sensor/usecase"
)

type Dependency struct {
	DB *sqlx.DB
	//repository
	SensorRepository repository.IRepository
	//usecase
	SensorUseCase usecase.IUseCase
}
