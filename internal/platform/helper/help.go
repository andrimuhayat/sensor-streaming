package helper

import (
	"github.com/mitchellh/mapstructure"
	"strconv"
)

func DecoderConfig(req interface{}) *mapstructure.DecoderConfig {
	config := &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &req,
		TagName:     "json",
	}
	return config
}

func ExpectedInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}
