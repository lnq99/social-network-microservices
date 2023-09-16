package model

import (
	"time"
)

type Profile struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Intro     string    `json:"intro"`
	AvatarS   string    `json:"avatars"`
	AvatarL   string    `json:"avatarl"`
	Created   time.Time `json:"created"`
	//PostCount  string  `json:"postCount"`
	//PhotoCount string  `json:"photoCount"`
}

type Relationship struct {
	User1   int       `json:"user1"`
	User2   int       `json:"user2"`
	Type    string    `json:"type"`
	Other   string    `json:"other"`
	Created time.Time `json:"created"`
}

type ShortInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AvatarS string `json:"avatars"`
}
