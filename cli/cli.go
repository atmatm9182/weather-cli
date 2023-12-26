package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atmatm9182/weather-cli/forecast"
)

var notEnoughArgsError = errors.New("Not enough arguments. See 'weather-cli --help' for help")

func ParseConfig() (config *forecast.Config, err error) {
	var help bool
	flag.BoolVar(&help, "help", false, "Prints this message")

	var ak string
	flag.StringVar(&ak, "key", "", "Provide an api key for OpenWeatherMap instead of getting it from config file")

	var u string
	flag.StringVar(&u, "units", "metric", "Units of measurment [metric | imperial | standard(kelvin)]")

	var coords bool
	flag.BoolVar(&coords, "coords", false, "Provide city coordinates in format 'longitude latitude' instead of providing city's name")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		err = notEnoughArgsError
		return
	}

	var units forecast.UnitsOfMeasurment
	units, err = forecast.ParseUnits(u)
	if err != nil {
		return
	}

	if ak == "" {
		var fileConfig *Config
		fileConfig, err = parseSystemConfig()
		if err != nil {
			return
		}
		ak = fileConfig.ApiKey
	}

	config = &forecast.Config{
		ApiKey: ak,
		Units:  units,
	}

	if coords {
		if flag.NArg() != 2 {
			err = fmt.Errorf("Expected to get 2 arguments as city coordinates, but got %d instead", flag.NArg())
			return
		}

		var lon float64
		lon, err = strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			err = fmt.Errorf("Error while parsing longitude: %s", err)
			return
		}

		var lat float64
		lat, err = strconv.ParseFloat(flag.Arg(1), 64)
		if err != nil {
			err = fmt.Errorf("Error while parsing latitude: %s", err)
			return
		}

		config.UseCoords = true
		config.Lon, config.Lat = lon, lat
	} else if flag.NArg() == 1 {
		config.City = flag.Arg(0)
	} else {
		config.City = strings.Join(flag.Args(), " ")
	}

	return
}
