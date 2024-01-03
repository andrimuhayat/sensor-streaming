package main

import (
	"sensor-streaming/cmd"
	"sensor-streaming/config"
	"sensor-streaming/internal/platform/cpu"
)

func main() {
	cpu.UtilizeCPU()
	var err error
	err = config.SetConfig("")
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
