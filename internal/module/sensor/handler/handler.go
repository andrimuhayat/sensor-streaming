package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sensor-streaming/config"
	"sensor-streaming/internal/module/sensor/usecase"
	"sensor-streaming/internal/platform/httpengine/httpresponse"
)

type Handler struct {
	UseCase usecase.IUseCase
}

// @Summary Generate stream data
// @Description stream data sensor
// @Accept  json
// @Produce  json
// @Param frequency query string true "example: 2"
// @Param data body dto.SensorDataGenerateRequest true "Stream data request"
// @Router /api/streaming/sensor-generate [post]
func (h Handler) GenerateStream(c echo.Context) error {
	request, err := config.MappingRequest(c)
	if err != nil {
		return httpresponse.ResponseWithError(c, http.StatusInternalServerError, err.Error())
	}

	err = h.UseCase.StreamingData(request)
	if err != nil {
		log.Println("{Sensor}{GenerateStream}{Error} : ", err)
		return httpresponse.ResponseWithError(c, http.StatusBadRequest, httpresponse.ErrorBadRequest.Message)
	}

	return httpresponse.ResponseWithJSON(c, http.StatusOK, httpresponse.ResponseSuccess(http.StatusOK, "success", nil))
}

func (h Handler) Health(c echo.Context) error {
	return httpresponse.ResponseWithJSON(c, http.StatusOK, "ok")
}

func NewHandler(useCase usecase.IUseCase) Handler {
	return Handler{
		UseCase: useCase,
	}
}
