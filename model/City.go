package model

//City Struct
type City struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	People string `json:"people"`
	CouncilDay string `json:"council_day"`
}