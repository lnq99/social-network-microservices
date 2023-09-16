package service

import (
	"app/model"
	"app/repository"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

type RelationshipServiceImpl struct {
	repo repository.Repo
}

// Функция создания сервиса связи пользователей
func NewRelationshipService(repo repository.Repo) RelationshipService {
	return &RelationshipServiceImpl{repo}
}

// Функция получения связей
func (r *RelationshipServiceImpl) Get(ctx context.Context, id int) (rels []model.Relationship, err error) {
	//return r.repo.Select(id)
	relationships, err := r.repo.GetRelationshipsTo(ctx, int32(id))
	for _, rel := range relationships {
		rels = append(rels, model.Relationship{
			User1:   int(rel.User1),
			User2:   int(rel.User2),
			Type:    rel.Type.String,
			Other:   rel.Other.String,
			Created: rel.Created.Time,
		})
	}
	return
}

// Функция получения "Друзей"
func (r *RelationshipServiceImpl) Friends(ctx context.Context, id int) (rels []model.Relationship, err error) {
	relationships, err := r.repo.GetFriendRelationships(ctx, int32(id))
	for _, rel := range relationships {
		rels = append(rels, model.Relationship{
			User1:   int(rel.User1),
			User2:   int(rel.User2),
			Type:    rel.Type.String,
			Other:   rel.Other.String,
			Created: rel.Created.Time,
		})
	}
	return
}

// Функция получения запросов на "Дружбу"
func (r *RelationshipServiceImpl) Requests(ctx context.Context, id int) (rels []model.Relationship, err error) {
	relationships, err := r.repo.GetRequestRelationships(ctx, int32(id))
	for _, rel := range relationships {
		rels = append(rels, model.Relationship{
			User1:   int(rel.User1),
			User2:   int(rel.User2),
			Type:    rel.Type.String,
			Other:   rel.Other.String,
			Created: rel.Created.Time,
		})
	}
	return
}

// Функция получения информации о "Друзьях"
func (r *RelationshipServiceImpl) FriendsDetail(ctx context.Context, id int) (string, error) {
	//return r.repo.FriendsDetail(id)

	res, err := r.repo.FriendsDetail(ctx, int32(id))
	return string(res), err
}

func (r *RelationshipServiceImpl) MutualFriends(ctx context.Context, u1, u2 int) ([]int, error) {
	//return r.repo.MutualFriends(u1, u2)

	mf, err := r.repo.MutualFriends(ctx, repository.MutualFriendsParams{
		U1: int32(u1),
		U2: int32(u2),
	})

	res := make([]int, len(mf))
	for i, v := range mf {
		res[i] = int(v)
	}
	return res, err
}

// Функция изменения типа связи между пользователями
func (r *RelationshipServiceImpl) ChangeType(ctx context.Context, user1, user2 int, t string) error {
	var rel repository.Relationship

	u1 := int32(user1)
	u2 := int32(user2)

	rel, _ = r.repo.GetRelationship(ctx, repository.GetRelationshipParams{
		User1: u1,
		User2: u2,
	})
	t12 := rel.Type.String

	rel, _ = r.repo.GetRelationship(ctx, repository.GetRelationshipParams{
		User1: u2,
		User2: u1,
	})
	t21 := rel.Type.String

	//t12 := r.repo.SelectRelationshipWith(u1, u2)
	//t21 := r.repo.SelectRelationshipWith(u2, u1)

	fmt.Println(u1, u2, t)
	fmt.Println(t12, t21)

	// TODO: err checking
	switch t {
	case "accept":
		if t21 == "request" {
			r.repo.UpdateRelationship(ctx, repository.UpdateRelationshipParams{User1: u1, User2: u2, Type: pgtype.Text{String: "friend"}})
			r.repo.UpdateRelationship(ctx, repository.UpdateRelationshipParams{User1: u2, User2: u1, Type: pgtype.Text{String: "friend"}})
		}
	case "delete":
		if t21 == "request" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u2, u1})
		}
	case "unfollow":
		if t12 == "request" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u1, u2})
		}
	case "request":
		if t12 != "friend" && t21 != "block" {
			r.repo.UpdateRelationship(ctx, repository.UpdateRelationshipParams{User1: u1, User2: u2, Type: pgtype.Text{String: "request"}})
		}
	case "unfriend":
		if t12 == "friend" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u1, u2})
		}
		if t21 == "friend" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u2, u1})
		}
	case "block":
		r.repo.UpdateRelationship(ctx, repository.UpdateRelationshipParams{User1: u1, User2: u2, Type: pgtype.Text{String: "block"}})
		if t21 != "block" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u2, u1})
		}
	case "unblock":
		if t12 == "block" {
			r.repo.DeleteRelationship(ctx, repository.DeleteRelationshipParams{u1, u2})
		}
	default:
		return fmt.Errorf("unknown type of relationship command")
	}

	//r.repo.UpdateRelationship(ctx, repository.UpdateRelationshipParams{
	//	Type:  pgtype.Text{},
	//	Other: pgtype.Text{},
	//	User1: 0,
	//	User2: 0,
	//})

	return nil
}

func (r *RelationshipServiceImpl) GetRelationshipWith(ctx context.Context, u1, u2 int) string {
	rel, err := r.repo.GetRelationship(ctx, repository.GetRelationshipParams{
		User1: int32(u1),
		User2: int32(u2),
	})

	if err != nil {
		return ""
	}

	t := rel.Type.String
	if t == "request" {
		t = "follow"
	}
	return t
}
