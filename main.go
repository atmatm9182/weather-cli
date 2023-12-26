package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/atmatm9182/weather-cli/cli"
	"github.com/atmatm9182/weather-cli/forecast"
)

const myApiKey = "21f4560db80f5925f55edd78be72ff29"

func capitalize(s string) string {
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

func main() {
	config, err := cli.ParseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	f, err := forecast.GetForecast(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error while trying to get weather forecast: %s\n", err)
		os.Exit(1)
	}

	descriptions := make([]string, len(f.Weather))
	for i, w := range f.Weather {
		w.Description = capitalize(w.Description)
		descriptions[i] = w.Description
	}

	desc := strings.Join(descriptions, ", ")

	fmt.Printf("Weather for %s, %s:\nCurrent temperature is %.2f%c\nFeels like %.2f\n%s\n",
		f.City,
		f.Sys.CountryCode,
		f.Temperature.Actual,
		config.Units.Letter(),
		f.Temperature.FeelsLike,
		desc,
	)
}
