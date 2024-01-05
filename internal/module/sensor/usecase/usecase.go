package usecase

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/url"
	module "sensor-streaming/config"
	"sensor-streaming/internal/module/sensor/dto"
	sensor "sensor-streaming/internal/module/sensor/repository"
	"sensor-streaming/internal/platform/helper"
)

type IUseCase interface {
	StreamingData(module.HTTPRequest) error
}

type UseCase struct {
	SensorRepository sensor.IRepository
	MqttClient       mqtt.Client
}

func (u UseCase) StreamingData(request module.HTTPRequest) error {
	//httpError := httpresponse.HTTPError{}
	var err error
	var sensorGenerateRequest dto.SensorDataGenerateRequest

	config := helper.DecoderConfig(&sensorGenerateRequest)
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	if err = decoder.Decode(request.Body); err != nil {
		log.Println("{SensorDataGenerateRequest}{Decode}{Error} : ", err)
		//helper.ResponseWithError(w, http.StatusInternalServerError, httpresponse.ErrorInternalServerError.Message)
	}

	var queryParamReq dto.SensorQueryParam
	config = helper.DecoderConfig(&queryParamReq)
	decoder, err = mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	if err = decoder.Decode(request.Queries); err != nil {
		log.Println("{SensorQueryParam}{Decode}{Error} : ", err)
		//helper.ResponseWithError(w, http.StatusInternalServerError, httpresponse.ErrorInternalServerError.Message)
	}

	if queryParamReq.Frequency == "" {
		queryParamReq.Frequency = "1"
	}

	for i := 0; i < helper.ExpectedInt(queryParamReq.Frequency); i++ {
		marshal, err := json.Marshal(sensorGenerateRequest)
		if err != nil {
			log.Println("{StreamingData}{Marshal}{Error} : ", err)
		}

		t := u.MqttClient.Publish(helper.STREAMSENSOR, 0, false, marshal)
		log.Println(i, u.MqttClient.IsConnected(), t.Error())
	}

	return nil
}

func (u UseCase) listen(uri *url.URL, topic string) {
	u.MqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func NewUseCase(repository sensor.IRepository, client mqtt.Client) IUseCase {
	return UseCase{
		SensorRepository: repository,
		MqttClient:       client,
	}
}
