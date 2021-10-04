package dao

import (
	"context"
	"database/sql"
	"https/github.com/atb-Mounir/GoCity/database"
	"https/github.com/atb-Mounir/GoCity/model"
	"log"
	"strconv"
	"time"

	_ "https/github.com/atb-Mounir/GoCity/database"

	_ "https/github.com/atb-Mounir/GoCity/model"
)

	// Get database connection
	var db, _ = database.GetDb()

	// Get Single City
func GetCity(id string)  (model.City, error){
	var city model.City
	// Query for a valu based on a single row
	if err := db.QueryRow("SELECT Id , name , description, people, council_day FROM city WHERE id=?", id).Scan(&city.Id, &city.Name, &city.Description, &city.People, &city.CouncilDay); err != nil {
		if err == sql.ErrNoRows{
			return model.City{}, err
		}
		return model.City{}, err
	}
	return city, nil

}