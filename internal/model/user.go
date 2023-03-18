package model

type Friend struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type User struct {
	ID       string `json:"id" bson:"_id"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
	Balance  string `json:"balance"`
	//TODO: handle age field
	Age        string   `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []Friend `json:"friends"`
	Data       string   `json:"data"`
}
