package model

import (
	"encoding/json"
	"strconv"
)

type CustomInt int

type Friend struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type User struct {
	ID         string    `json:"id" bson:"_id"`
	Password   *string    `json:"password"`
	IsActive   *bool      `json:"isActive"`
	Balance    *string    `json:"balance"`
	Age        *CustomInt `json:"age"`
	Name       *string    `json:"name"`
	Gender     *string    `json:"gender"`
	Company    *string    `json:"company"`
	Email      *string    `json:"email"`
	Phone      *string    `json:"phone"`
	Address    *string    `json:"address"`
	About      *string    `json:"about"`
	Registered *string    `json:"registered"`
	Latitude   *float64   `json:"latitude"`
	Longitude  *float64   `json:"longitude"`
	Tags       *[]string  `json:"tags"`
	Friends    *[]Friend  `json:"friends"`
	Data       *string    `json:"data"`
}

func (ci *CustomInt) UnmarshalJSON(b []byte) error {
	// Check if the value is a string
	if b[0] == '"' {
		// Parse the string value
		var str string
		err := json.Unmarshal(b, &str)
		if err != nil {
			return err
		}
		// Convert the string to an integer
		i, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*ci = CustomInt(i)
	} else {
		// Parse the integer value
		var i int
		err := json.Unmarshal(b, &i)
		if err != nil {
			return err
		}
		*ci = CustomInt(i)
	}
	return nil
}
