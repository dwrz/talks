package weather

type response struct {
	Temperatures temperatures `json:"main"`
}

type temperatures struct {
	Temp float64 `json:"temp"`
}
