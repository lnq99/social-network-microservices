package service

import (
	"app/model"
	"app/repository"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

type PhotoServiceImpl struct {
	repo repository.Repo
}

// Функция создания сервиса фото
func NewPhotoService(repo repository.Repo) PhotoService {
	return &PhotoServiceImpl{repo}
}

// Функция получения альбома фото пользователя
func (r *PhotoServiceImpl) GetAlbumByUserId(ctx context.Context, userId int) (res []model.Album, err error) {
	albums, err := r.repo.GetAlbumByUserId(ctx, int32(userId))
	for _, album := range albums {
		res = append(res, model.Album{
			Id:      int(album.ID),
			UserId:  int(album.Userid),
			Descr:   album.Descr.String,
			Created: album.Created.Time.String(),
		})
	}
	return
}

// Функция получения альбома
func (r *PhotoServiceImpl) GetAlbumId(ctx context.Context, userId int, album string) (albumId int, err error) {
	albums, err := r.repo.GetAlbumByUserId(ctx, int32(userId))
	if err == nil {
		for _, a := range albums {
			if a.Descr.String == album {
				albumId = int(a.ID)
				return
			}
		}
	}
	return
}

// Функция получения фото
func (r *PhotoServiceImpl) GetPhoto(ctx context.Context, id int) (model.Photo, error) {
	p, err := r.repo.GetPhoto(ctx, int32(id))
	photo := model.Photo{
		Id:      int(p.ID),
		UserId:  int(p.Userid),
		AlbumId: int(p.Albumid),
		Url:     p.Url.String,
		Created: p.Created.Time.String(),
	}
	return photo, err
}

// Функция получения фото пользователя
func (r *PhotoServiceImpl) GetPhotoByUserId(ctx context.Context, userId int) (res []model.Photo, err error) {
	photos, err := r.repo.GetPhotoByUserId(ctx, int32(userId))
	for _, p := range photos {
		res = append(res, model.Photo{
			Id:      int(p.ID),
			UserId:  int(p.Userid),
			AlbumId: int(p.Albumid),
			Url:     p.Url.String,
			Created: p.Created.Time.String(),
		})
	}
	return
}

// Функция загрузки фото в альбом
func (r *PhotoServiceImpl) UploadPhotoToAlbum(ctx context.Context, p model.Photo, album string) (photoId int, err error) {
	p.AlbumId, err = r.GetAlbumId(ctx, p.UserId, album)
	fmt.Println(p.UserId, album, p.AlbumId, err)
	if err != nil {
		return -1, err
	}
	photo, err := r.repo.CreatePhoto(ctx, repository.CreatePhotoParams{
		Userid:  int32(p.UserId),
		Albumid: int32(p.AlbumId),
		Url:     pgtype.Text{p.Url, true},
	})
	return int(photo.ID), err
}

// Функция-оберка загрузки фото
func (r *PhotoServiceImpl) UploadPhoto(ctx context.Context, p model.Photo) (photoId int, err error) {
	return r.UploadPhotoToAlbum(ctx, p, "Upload")
}

// Функция-обертка загрузки новой аватарки
func (r *PhotoServiceImpl) SetAvatar(ctx context.Context, p model.Photo) (err error) {
	_, err = r.UploadPhotoToAlbum(ctx, p, "Avatar")
	return
}
