package controller

import (
	"app/model"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Obj  interface{}
}

type ErrResponse struct {
	Code int    `json:"-"`
	Msg  string `json:"message"`
}

type Msg struct {
	Msg string `json:"message"`
}

func jsonResponse(c *gin.Context, err error, res Response, errRes ErrResponse) {
	if err != nil {
		if errRes.Msg == "" {
			errRes.Msg = err.Error()
		}
		//logger.Err(err)
		c.JSON(errRes.Code, errRes)
	} else if res.Obj != nil {
		c.JSON(res.Code, res.Obj)
	} else {
		c.Status(res.Code)
	}
}

func statusResponse(c *gin.Context, err error, code int, errCode int) {
	if err != nil {
		//logger.Err(err)
		c.Status(errCode)
	} else {
		c.Status(code)
	}
}

type loginResponse struct {
	Token string          `json:"token"`
	User  ProfileResponse `json:"user"`
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

type GetMutualAndTypeResponse struct {
	T      string `json:"type"`
	Mutual []int  `json:"mutual"`
}

type SearchResponse struct {
	Id     int    `json:"id"`
	Mutual int    `json:"mutual"`
	T      string `json:"type"`
}

type FriendResponse struct {
	Id      int    `json:"id"`
	Name    int    `json:"name"`
	AvatarS string `json:"avatars"`
}

type ProfileResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Created   string `json:"created"`
	Intro     string `json:"intro"`
	AvatarS   string `json:"avatars"`
	AvatarL   string `json:"avatarl"`
	//PostCount  string `json:"postCount"`
	//PhotoCount string `json:"photoCount"`
}

func toProfileResponse(profile model.Profile) ProfileResponse {
	return ProfileResponse{
		Id:        profile.Id,
		Name:      profile.Name,
		Gender:    profile.Gender,
		Birthdate: profile.Birthdate.String(),
		Created:   profile.Created.String(),
		Intro:     profile.Intro,
		AvatarS:   profile.AvatarS,
		AvatarL:   profile.AvatarL,
		//PostCount:  profile.PostCount,
		//PhotoCount: profile.PhotoCount,
	}
}
