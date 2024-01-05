package common

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

	MqttClient mqtt.Client
}
