`weather-cli` is a simple and easy to use CLI app to quickly and conviniently check the weather using [OpenWeatherMap](https://openweathermap.org/) 

# Quick start

Just run:

```shell
go install github.com/atmatm9182/weather-cli
weather-cli -help
```

Or if you want to build from source:

```shell
git clone https://github.com/atmatm9182/weather-cli
cd weather-cli
go build 
./weather-cli -help
```

# Configuration
The configuration is done via a configuration file written in [json](https://en.wikipedia.org/wiki/JSON) and has the following parameters:
- `ApiKey: string` - your api key that you can get from [OpenWeatherMap](https://openweathermap.org/) website after registering

## Path to config
Linux: `$XDG_CONFIG_HOME/weather-cli/config.json` (or `$HOME/.config/config.json` if `$XDG_CONFIG_HOME` is not set)  
Windows: `%APPDATA%\Weather CLI\config.json`