package service

import (
	"app/model"
	"app/repository"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostServiceImpl struct {
	repo repository.Repo
}

// Функция создания сервиса публикации
func NewPostService(repo repository.Repo) PostService {
	return &PostServiceImpl{repo}
}

// Функция  получения публикации
func (r *PostServiceImpl) Get(ctx context.Context, postId int) (post model.Post, err error) {
	p, err := r.repo.GetPost(ctx, int32(postId))
	post = model.Post{
		Id:       int(p.ID),
		UserId:   int(p.Userid),
		Created:  p.Created.Time.String(),
		Tags:     p.Tags.String,
		Content:  p.Content.String,
		AtchType: p.Atchtype.String,
		AtchId:   int(p.Atchid.Int32),
		AtchUrl:  p.Atchurl.String,
		Reaction: p.Reaction,
		CmtCount: int(p.Cmtcount.Int32),
	}
	return
}

// Функция получения публикации пользователя
func (r *PostServiceImpl) GetByUserId(ctx context.Context, userId int) ([]int32, error) {
	posts, err := r.repo.GetPostByUserId(ctx, int32(userId))
	return posts.([]int32), err
}

// Функция получения реакции на публикацию
func (r *PostServiceImpl) GetReaction(ctx context.Context, postId int) ([]int32, error) {
	reacts, err := r.repo.GetReaction(ctx, int32(postId))
	return reacts, err
}

// Функция добавления публикации
func (r *PostServiceImpl) Post(ctx context.Context, userId int, body PostBody) (err error) {
	post := model.Post{
		UserId:   userId,
		Tags:     body.Tags,
		Content:  body.Content,
		AtchType: body.AtchType,
		AtchId:   body.AtchId,
		AtchUrl:  body.AtchUrl,
	}
	if post.AtchType == "photo" {
		photoId, err := GetService(r.repo).Photo.UploadPhoto(ctx, model.Photo{
			UserId: post.UserId,
			Url:    post.AtchUrl,
		})

		if err != nil {
			return err
		}
		post.AtchId = photoId
	}
	_, err = r.repo.CreatePost(ctx, repository.CreatePostParams{
		Userid:   int32(userId),
		Tags:     pgtype.Text{body.Tags, true},
		Content:  pgtype.Text{body.Content, true},
		Atchtype: pgtype.Text{body.AtchType, true},
		Atchid:   pgtype.Int4{int32(body.AtchId), true},
		Atchurl:  pgtype.Text{body.AtchUrl, true},
	})
	return
}

// Функция удаления публикации
func (r *PostServiceImpl) Delete(ctx context.Context, userId int, postId int) error {
	_, err := r.repo.DeletePost(ctx, repository.DeletePostParams{
		Userid: int32(userId),
		ID:     int32(postId),
	})
	return err
}
