#!/bin/bash

export OPEN_WEATHER_MAP_API_KEY="$(pass open-weather-map/api-key)"

go run ./*.go
