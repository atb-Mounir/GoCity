package goCity

import (
	"encoding/json"
	"go/ast"
)

// City Struct (Model)
type City struct {
	ID 			string `json:"id"`
	Name 		string	`json:"name"`
	People 		string	`json:"people"`
	Description string	`json:"description"`
}