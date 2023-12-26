package forecast

type Config struct {
	ApiKey    string
	City      string
	Units     UnitsOfMeasurment
	UseCoords bool
	Lon, Lat  float64
}
