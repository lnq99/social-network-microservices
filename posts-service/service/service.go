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
	Post     PostService
	Comment  CommentService
	Reaction ReactionService
	Photo    PhotoService
	Feed     FeedService
}

func GetService(repo repository.Repo) *Service {
	once.Do(func() {
		service = &Service{
			Post:     NewPostService(repo),
			Comment:  NewCommentService(repo),
			Reaction: NewReactionService(repo),
			Photo:    NewPhotoService(repo),
			Feed:     NewFeedService(repo),
		}
	})
	return service
}

type PostService interface {
	Get(ctx context.Context, postId int) (model.Post, error)
	GetReaction(ctx context.Context, postId int) ([]int32, error)
	GetByUserId(ctx context.Context, userId int) ([]int32, error)
	Post(ctx context.Context, userId int, body PostBody) error
	Delete(ctx context.Context, userId int, postId int) error
}

type CommentService interface {
	GetTree(ctx context.Context, postId int) (string, error)
	Add(ctx context.Context, userId int, body CommentBody) error
}

type ReactionService interface {
	Get(ctx context.Context, postId int) ([]model.Reaction, error)
	GetByUserPost(ctx context.Context, userId, postId int) (string, error)
	UpdateReaction(ctx context.Context, userId, postId int, t string) error
}

type PhotoService interface {
	GetAlbumByUserId(ctx context.Context, userId int) ([]model.Album, error)
	GetAlbumId(ctx context.Context, userId int, album string) (int, error)

	GetPhoto(ctx context.Context, id int) (model.Photo, error)
	GetPhotoByUserId(ctx context.Context, userId int) ([]model.Photo, error)

	UploadPhotoToAlbum(ctx context.Context, p model.Photo, album string) (int, error)
	UploadPhoto(ctx context.Context, p model.Photo) (int, error)
	SetAvatar(ctx context.Context, p model.Photo) error
}

type FeedService interface {
	GetNewsfeed(ctx context.Context, ids_arr []int32, limit, offset int32) ([]int32, error)
}

type CommentBody struct {
	PostId   int    `json:"postId"`
	ParentId int    `json:"parentId"`
	Content  string `json:"content"`
}

type PostBody struct {
	Tags     string `json:"tags"`
	Content  string `json:"content"`
	AtchType string `json:"atchType"`
	AtchId   int    `json:"atchId,omitempty"`
	AtchUrl  string `json:"atchUrl"`
}
