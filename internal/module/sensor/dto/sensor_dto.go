package dto

type SensorQueryParam struct {
	Frequency string `json:"frequency"`
}

type SensorDataGenerateRequest struct {
	SensorValue float32 `json:"sensor_value"`
	SensorType  string  `json:"sensor_type"`
	ID1         string  `json:"ID1" `
	ID2         int     `json:"ID2" `
}
