package constant

import (
	"fmt"
	"strconv"
)

type Configuration string

const (
	Environment  Configuration = "DEVELOPMENT"
	EnvPrefix    Configuration = "ST3LLAR"
	Organization Configuration = "57B"
)

const (
	ConfigFileName Configuration = ".st3llar"
	ConfigFileType Configuration = "yaml"
)

func (cc Configuration) ValStr() string {
	return string(cc)
}

func (cc Configuration) ValInt() int {
	intVal, err := strconv.Atoi(cc.ValStr())
	if err != nil {
		fmt.Printf("Error converting <%s> to int: %s\n", cc.ValStr(), err.Error())
		return 0
	}

	return intVal
}

func (cc Configuration) ValFloat32() float32 {
	floatValue, err := strconv.ParseFloat(cc.ValStr(), 32)
	if err != nil {
		fmt.Printf("Error converting <%s> to float32: %s\n", cc.ValStr(), err.Error())
		return float32(0.0)
	}

	return float32(floatValue)
}

func (cc Configuration) ValFloat64() float64 {
	floatValue, err := strconv.ParseFloat(cc.ValStr(), 64)
	if err != nil {
		fmt.Printf("Error converting <%s> to float64: %s\n", cc.ValStr(), err.Error())
		return 0.0
	}

	return floatValue
}
