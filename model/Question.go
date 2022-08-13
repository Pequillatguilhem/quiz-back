package model

type Question struct {
	Id       *int   `json:"id"`
	Name     string `json:"name"`
	Response string `json:"response"`
	SerieId  int    `json:"serieId"`
}
