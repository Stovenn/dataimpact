package util

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/stovenn/dataimpact/internal/model"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateDataDirIfNotExists() {
	if _, err := os.Stat("data/"); errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data", 0755)
	}
}

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomFloat returns a random float64 between min and max
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail returns a random email
func RandomEmail(n int) string {
	return fmt.Sprintf("%s@email.com", RandomString(n))
}

// RandomUser returns a random User
func RandomUser() *model.User {
	password := RandomString(12)
	isActive := false
	balance := RandomString(12)
	age := model.CustomInt(RandomInt(1, 99))
	name := RandomString(15)
	gender := RandomString(1)
	company := RandomString(15)
	email := RandomEmail(10)
	phone := RandomString(13)
	address := RandomString(40)
	about := RandomString(50)
	registered := RandomString(12)
	latitude := RandomFloat(-100, 100)
	longitude := RandomFloat(-100, 100)
	tags := []string{RandomString(10), RandomString(10), RandomString(10)}
	friends := []model.Friend{
		{ID: RandomInt(0, 100), Name: RandomString(10)},
		{ID: RandomInt(0, 100), Name: RandomString(10)},
	}
	data := RandomString(2000)

	return &model.User{
		ID:         RandomString(30),
		Password:   &password,
		IsActive:   &isActive,
		Balance:    &balance,
		Age:        &age,
		Name:       &name,
		Gender:     &gender,
		Company:    &company,
		Email:      &email,
		Phone:      &phone,
		Address:    &address,
		About:      &about,
		Registered: &registered,
		Latitude:   &latitude,
		Longitude:  &longitude,
		Tags:       &tags,
		Friends:    &friends,
		Data:       &data,
	}
}
