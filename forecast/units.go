package forecast

import (
	"fmt"
	"log"
)

type UnitsOfMeasurment uint8

const (
	UnitsStandard UnitsOfMeasurment = iota
	UnitsMetric
	UnitsImperial
)

func (u UnitsOfMeasurment) String() (s string) {
	switch u {
	case UnitsStandard:
		s = "standard"
	case UnitsImperial:
		s = "imperial"
	case UnitsMetric:
		s = "metric"
	default:
		log.Fatal("Tried to convert invalid unit of measurment to string")
	}

	return
}

func (u UnitsOfMeasurment) Letter() (r rune) {
	switch u {
	case UnitsMetric:
		r = 'C'
	case UnitsStandard:
		r = 'K'
	case UnitsImperial:
		r = 'F'
	default:
		log.Fatal("Tried to get a letter for an invalid unit of measurement")
	}

	return
}

func ParseUnits(s string) (units UnitsOfMeasurment, err error) {
	switch s {
	case "metric":
		units = UnitsMetric
	case "imperial":
		units = UnitsImperial
	case "standard", "kelvin":
		units = UnitsStandard
	default:
		err = fmt.Errorf("\"%s\" is not a valid unit of measurement", s)
	}

	return
}
