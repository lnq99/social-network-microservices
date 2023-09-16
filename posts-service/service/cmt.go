package service

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"

	"app/model"
	"app/repository"
)

type CommentServiceImpl struct {
	repo repository.Repo
}

// Функция сервиса комментария
func NewCommentService(repo repository.Repo) CommentService {
	return &CommentServiceImpl{repo}
}

// Функция получения дерева комментариев
func (r *CommentServiceImpl) GetTree(ctx context.Context, postId int) (res string, err error) {
	var c []model.Comment
	cmts, err := r.repo.GetComments(ctx, int32(postId))
	for _, cmt := range cmts {
		c = append(c, model.Comment{
			Id:       int(cmt.ID),
			UserId:   int(cmt.Userid),
			PostId:   int(cmt.Postid),
			ParentId: int(cmt.Parentid.Int32),
			Content:  cmt.Content.String,
			Created:  cmt.Created.Time.String(),
		})
	}
	res = r.BuildCmtTree(c)
	return
}

// Функция создания создания дерева комментарией
func (r *CommentServiceImpl) BuildCmtTree(cmts []model.Comment) (tree string) {
	m := make(map[int]*model.Comment)

	for i := range cmts {
		m[cmts[i].Id] = &cmts[i]
	}

	for i, n := range cmts {
		if m[n.ParentId] != nil {
			m[n.ParentId].Children = append(m[n.ParentId].Children, &cmts[i])
		}
	}

	out := []*model.Comment{}
	for _, v := range m {
		if v.ParentId == 0 {
			out = append(out, v)
		}
	}

	bytes, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	tree = string(bytes)

	return
}

// Функция добавления комментария
func (r *CommentServiceImpl) Add(ctx context.Context, userId int, body CommentBody) (err error) {
	_, err = r.repo.CreateComment(ctx, repository.CreateCommentParams{
		Userid:   int32(userId),
		Postid:   int32(body.PostId),
		Parentid: pgtype.Int4{int32(body.ParentId), true},
		Content:  pgtype.Text{body.Content, true},
	})
	return
}
