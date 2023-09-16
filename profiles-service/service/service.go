package service

import (
	"context"
	"sync"

	"app/model"
	"app/repository"
)

var (
	once    sync.Once
	service *Service
)

type Service struct {
	Profile      ProfileService
	Relationship RelationshipService
}

func GetService(repo repository.Repo) *Service {
	once.Do(func() {
		service = &Service{
			Profile:      NewProfileService(repo),
			Relationship: NewRelationshipService(repo),
		}
	})
	return service
}

type ProfileService interface {
	Get(ctx context.Context, id int) (model.Profile, error)
	GetByEmail(ctx context.Context, e string) (model.Profile, error)
	SearchName(ctx context.Context, id int, s string) (string, error)
	Register(ctx context.Context, p ProfileBody) error
	ChangeIntro(ctx context.Context, id int, intro string) error
	SetAvatar(ctx context.Context, id int, photoUrl string) error
}

type RelationshipService interface {
	Get(ctx context.Context, id int) ([]model.Relationship, error)
	Friends(ctx context.Context, id int) ([]model.Relationship, error)
	Requests(ctx context.Context, id int) ([]model.Relationship, error)
	FriendsDetail(ctx context.Context, id int) (string, error)
	MutualFriends(ctx context.Context, u1, u2 int) ([]int, error)
	GetRelationshipWith(ctx context.Context, u1, u2 int) string
	ChangeType(ctx context.Context, u1, u2 int, t string) error
}

type IntroBody struct {
	Intro string `json:"intro"`
}

type ProfileBody struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}
