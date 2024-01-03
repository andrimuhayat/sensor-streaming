package usecase

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	module "sensor-streaming/config"
	"sensor-streaming/internal/module/sensor/dto"
	sensor "sensor-streaming/internal/module/sensor/repository"
	"sensor-streaming/internal/platform/util"
)

type IUseCase interface {
	StreamingData(module.HTTPRequest) error
}

type UseCase struct {
	SensorRepository sensor.IRepository
}

func (u UseCase) StreamingData(request module.HTTPRequest) error {
	//httpError := httpresponse.HTTPError{}
	var err error
	var sensorGenerateRequest dto.SensorDataGenerateRequest

	config := util.DecoderConfig(&sensorGenerateRequest)
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	if err = decoder.Decode(request.Body); err != nil {
		log.Println("{SensorDataGenerateRequest}{Decode}{Error} : ", err)
		//helper.ResponseWithError(w, http.StatusInternalServerError, httpresponse.ErrorInternalServerError.Message)
	}

	var queryParamReq dto.SensorQueryParam
	config = util.DecoderConfig(&queryParamReq)
	decoder, err = mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	if err = decoder.Decode(request.Queries); err != nil {
		log.Println("{SensorQueryParam}{Decode}{Error} : ", err)
		//helper.ResponseWithError(w, http.StatusInternalServerError, httpresponse.ErrorInternalServerError.Message)
	}
	for i := 0; i < util.ExpectedInt(queryParamReq.Frequency); i++ {
		fmt.Println("index:", i)
	}
	return nil
}

func NewUseCase(repository sensor.IRepository) IUseCase {
	return UseCase{
		SensorRepository: repository,
	}
}
