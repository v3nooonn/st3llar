package constant

import (
	"fmt"
	"strconv"
)

type ConfigConst string

const (
	EnvPrefix ConfigConst = "st3llar"

	// ConfigFilePath ConfigConst = "./internal/config/"

	ConfigFileName ConfigConst = ".st3llar"
	ConfigFileType ConfigConst = ".yaml"
)

func (cc ConfigConst) ValStr() string {
	return string(cc)
}

func (cc ConfigConst) ValInt() int {
	intVal, err := strconv.Atoi(cc.ValStr())
	if err != nil {
		fmt.Printf("Error converting <%s> to int: %s\n", cc.ValStr(), err.Error())
		return 0
	}

	return intVal
}

func (cc ConfigConst) ValFloat32() float32 {
	floatValue, err := strconv.ParseFloat(cc.ValStr(), 32)
	if err != nil {
		fmt.Printf("Error converting <%s> to float32: %s\n", cc.ValStr(), err.Error())
		return float32(0.0)
	}

	return float32(floatValue)
}

func (cc ConfigConst) ValFloat64() float64 {
	floatValue, err := strconv.ParseFloat(cc.ValStr(), 64)
	if err != nil {
		fmt.Printf("Error converting <%s> to float64: %s\n", cc.ValStr(), err.Error())
		return 0.0
	}

	return floatValue
}
