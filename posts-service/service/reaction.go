package service

import (
	"app/model"
	"app/repository"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

type ReactionServiceImpl struct {
	repo repository.Repo
}

// Функция создания сервиса реакции
func NewReactionService(repo repository.Repo) ReactionService {
	return &ReactionServiceImpl{repo}
}

// Функция  получения реакции
func (r *ReactionServiceImpl) Get(ctx context.Context, postId int) (res []model.Reaction, err error) {
	reacts, err := r.repo.GetReactions(ctx, int32(postId))
	for _, react := range reacts {
		res = append(res, model.Reaction{
			UserId: int(react.Userid),
			PostId: int(react.Postid),
			T:      react.Typ.String,
		})
	}
	return
}

// Функция получения реакции на публикацию пользователя
func (r *ReactionServiceImpl) GetByUserPost(ctx context.Context, userId, postId int) (string, error) {
	react, err := r.repo.GetReactionOfUser(ctx, repository.GetReactionOfUserParams{
		Userid: int32(userId),
		Postid: int32(postId),
	})
	return react.String, err
}

// Функция обновления реакции на публикацию
func (r *ReactionServiceImpl) UpdateReaction(ctx context.Context, userId, postId int, t string) error {
	var err error
	if t == "del" {
		_, err = r.repo.DeleteReaction(ctx, repository.DeleteReactionParams{
			Userid: int32(userId),
			Postid: int32(postId),
		})
	} else {
		_, err = r.repo.CreateReaction(ctx, repository.CreateReactionParams{
			Userid: int32(userId),
			Postid: int32(postId),
			Typ:    pgtype.Text{t, true},
		})
	}
	return err
}
