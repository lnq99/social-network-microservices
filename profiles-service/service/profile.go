package service

import (
	"context"
	"fmt"
	"time"

	"app/model"
	"app/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProfileServiceImpl struct {
	repo repository.Repo
}

// Функция созадния сервиса профиля
func NewProfileService(repo repository.Repo) ProfileService {
	return &ProfileServiceImpl{repo}
}

// Функция получения профиля
func (r *ProfileServiceImpl) Get(ctx context.Context, id int) (model.Profile, error) {
	p, err := r.repo.GetProfile(ctx, int32(id))
	if err != nil {
		return model.Profile{}, err
	}

	res := model.Profile{
		Id:        int(p.ID),
		Name:      p.Name,
		Gender:    p.Gender,
		Birthdate: p.Birthdate.Time,
		Email:     p.Email,
		Phone:     p.Phone.Int.String(),
		Intro:     p.Intro.String,
		AvatarS:   p.Avatars.String,
		AvatarL:   p.Avatarl.String,
		Created:   p.Created.Time,
	}
	return res, nil
}

// Функция получения профиля по почте
func (r *ProfileServiceImpl) GetByEmail(ctx context.Context, e string) (model.Profile, error) {
	p, err := r.repo.GetProfileByEmail(ctx, e)
	if err != nil {
		return model.Profile{}, err
	}

	res := model.Profile{
		Id:        int(p.ID),
		Name:      p.Name,
		Gender:    p.Gender,
		Birthdate: p.Birthdate.Time,
		Email:     p.Email,
		Phone:     p.Phone.Int.String(),
		Intro:     p.Intro.String,
		AvatarS:   p.Avatars.String,
		AvatarL:   p.Avatarl.String,
		Created:   p.Created.Time,
	}
	return res, nil
}

// Функция поиска профиля по имени
func (r *ProfileServiceImpl) SearchName(ctx context.Context, id int, s string) (string, error) {
	res, err := r.repo.SearchName(ctx, repository.SearchNameParams{
		U:       int32(id),
		Pattern: s,
	})
	return string(res), err
}

// Функция  регистрации нового пользователя
func (r *ProfileServiceImpl) Register(ctx context.Context, body ProfileBody) (err error) {
	//manager := auth.GetManager()
	//salt, hashed := manager.GetHashSalt(body.Password)
	//p := model.Profile{
	//	Email:     body.Email,
	//	Name:      body.Username,
	//	Salt:      salt,
	//	Hash:      hashed,
	//	Gender:    body.Gender,
	//	Birthdate: body.Birthdate,
	//}
	//_, err = r.repo.Insert(&p)

	birthdate, _ := time.Parse(time.RFC3339, body.Birthdate)

	_, err = r.repo.CreateProfile(ctx, repository.CreateProfileParams{
		Name:      body.Username,
		Gender:    body.Gender,
		Birthdate: pgtype.Date{Time: birthdate, Valid: true},
		Email:     body.Email,
		Phone:     pgtype.Numeric{},
	})
	return
}

// Функция установки аватарки
func (r *ProfileServiceImpl) SetAvatar(ctx context.Context, id int, photoUrl string) error {
	_, err := r.repo.UpdateProfile(ctx, repository.UpdateProfileParams{
		ID:      int32(id),
		Avatars: pgtype.Text{String: photoUrl, Valid: true},
		Avatarl: pgtype.Text{String: photoUrl, Valid: true},
	})
	return err
}

// Функция изменения информации о себе
func (r *ProfileServiceImpl) ChangeIntro(ctx context.Context, id int, intro string) error {
	fmt.Println(intro)
	_, err := r.repo.UpdateProfile(ctx, repository.UpdateProfileParams{
		ID:    int32(id),
		Intro: pgtype.Text{String: intro, Valid: true},
	})
	return err
}
