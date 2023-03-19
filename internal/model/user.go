package model

import (
	"encoding/json"
	"strconv"
)

type CustomInt int

type Friend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type User struct {
	ID         string     `json:"id" bson:"_id"`
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

type UserResponse struct {
	ID         string    `json:"id"`
	IsActive   bool      `json:"isActive"`
	Balance    string    `json:"balance"`
	Age        CustomInt `json:"age"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Company    string    `json:"company"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	About      string    `json:"about"`
	Registered string    `json:"registered"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Tags       []string  `json:"tags"`
	Friends    []Friend  `json:"friends"`
	Data       string    `json:"data"`
}
type LoginCredentials struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func ToResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:         u.ID,
		IsActive:   *u.IsActive,
		Balance:    *u.Balance,
		Age:        *u.Age,
		Name:       *u.Name,
		Gender:     *u.Gender,
		Company:    *u.Company,
		Email:      *u.Email,
		Phone:      *u.Phone,
		Address:    *u.Address,
		About:      *u.About,
		Registered: *u.Registered,
		Latitude:   *u.Latitude,
		Longitude:  *u.Longitude,
		Tags:       *u.Tags,
		Friends:    *u.Friends,
		Data:       *u.Data,
	}
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
