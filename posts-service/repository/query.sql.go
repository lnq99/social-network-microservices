// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAlbum = `-- name: CreateAlbum :one
insert into Album(id, userId, descr, created)
values ($1, $2, $3, $4)
returning id, userid, descr, created
`

type CreateAlbumParams struct {
	ID      int32              `json:"id"`
	Userid  int32              `json:"userid"`
	Descr   pgtype.Text        `json:"descr"`
	Created pgtype.Timestamptz `json:"created"`
}

func (q *Queries) CreateAlbum(ctx context.Context, arg CreateAlbumParams) (Album, error) {
	row := q.db.QueryRow(ctx, createAlbum,
		arg.ID,
		arg.Userid,
		arg.Descr,
		arg.Created,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Descr,
		&i.Created,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
insert into Comment(userId, postId, parentId, content)
values ($1, $2, $3, $4)
returning id, userid, postid, parentid, content, created
`

type CreateCommentParams struct {
	Userid   int32       `json:"userid"`
	Postid   int32       `json:"postid"`
	Parentid pgtype.Int4 `json:"parentid"`
	Content  pgtype.Text `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.Userid,
		arg.Postid,
		arg.Parentid,
		arg.Content,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Postid,
		&i.Parentid,
		&i.Content,
		&i.Created,
	)
	return i, err
}

const createPhoto = `-- name: CreatePhoto :one
insert into Photo(userId, albumId, url)
values ($1, $2, $3)
returning id, userid, albumid, url, created
`

type CreatePhotoParams struct {
	Userid  int32       `json:"userid"`
	Albumid int32       `json:"albumid"`
	Url     pgtype.Text `json:"url"`
}

func (q *Queries) CreatePhoto(ctx context.Context, arg CreatePhotoParams) (Photo, error) {
	row := q.db.QueryRow(ctx, createPhoto, arg.Userid, arg.Albumid, arg.Url)
	var i Photo
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Albumid,
		&i.Url,
		&i.Created,
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO Post (userId, tags, content, atchType, atchId, atchUrl)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, userid, created, tags, content, atchtype, atchid, atchurl, reaction, cmtcount
`

type CreatePostParams struct {
	Userid   int32       `json:"userid"`
	Tags     pgtype.Text `json:"tags"`
	Content  pgtype.Text `json:"content"`
	Atchtype pgtype.Text `json:"atchtype"`
	Atchid   pgtype.Int4 `json:"atchid"`
	Atchurl  pgtype.Text `json:"atchurl"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.Userid,
		arg.Tags,
		arg.Content,
		arg.Atchtype,
		arg.Atchid,
		arg.Atchurl,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Created,
		&i.Tags,
		&i.Content,
		&i.Atchtype,
		&i.Atchid,
		&i.Atchurl,
		&i.Reaction,
		&i.Cmtcount,
	)
	return i, err
}

const createReaction = `-- name: CreateReaction :one
insert into Reaction(userId, postId, typ)
values ($1, $2, $3)
on conflict (userId, postId) do update set typ = $3
returning userid, postid, typ
`

type CreateReactionParams struct {
	Userid int32       `json:"userid"`
	Postid int32       `json:"postid"`
	Typ    pgtype.Text `json:"typ"`
}

func (q *Queries) CreateReaction(ctx context.Context, arg CreateReactionParams) (Reaction, error) {
	row := q.db.QueryRow(ctx, createReaction, arg.Userid, arg.Postid, arg.Typ)
	var i Reaction
	err := row.Scan(&i.Userid, &i.Postid, &i.Typ)
	return i, err
}

const deletePost = `-- name: DeletePost :one
delete
from Post
where id = $2
  and userId = $1
returning id, userid, created, tags, content, atchtype, atchid, atchurl, reaction, cmtcount
`

type DeletePostParams struct {
	Userid int32 `json:"userid"`
	ID     int32 `json:"id"`
}

func (q *Queries) DeletePost(ctx context.Context, arg DeletePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, deletePost, arg.Userid, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Created,
		&i.Tags,
		&i.Content,
		&i.Atchtype,
		&i.Atchid,
		&i.Atchurl,
		&i.Reaction,
		&i.Cmtcount,
	)
	return i, err
}

const deleteReaction = `-- name: DeleteReaction :one
delete
from Reaction
where userId = $1
  and postId = $2
returning userid, postid, typ
`

type DeleteReactionParams struct {
	Userid int32 `json:"userid"`
	Postid int32 `json:"postid"`
}

func (q *Queries) DeleteReaction(ctx context.Context, arg DeleteReactionParams) (Reaction, error) {
	row := q.db.QueryRow(ctx, deleteReaction, arg.Userid, arg.Postid)
	var i Reaction
	err := row.Scan(&i.Userid, &i.Postid, &i.Typ)
	return i, err
}

const getAlbum = `-- name: GetAlbum :one
select id, userid, descr, created
from Album
where id = $1
limit 1
`

func (q *Queries) GetAlbum(ctx context.Context, id int32) (Album, error) {
	row := q.db.QueryRow(ctx, getAlbum, id)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Descr,
		&i.Created,
	)
	return i, err
}

const getAlbumByUserId = `-- name: GetAlbumByUserId :many
select id, userid, descr, created
from Album
where UserId = $1
`

func (q *Queries) GetAlbumByUserId(ctx context.Context, userid int32) ([]Album, error) {
	rows, err := q.db.Query(ctx, getAlbumByUserId, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Album{}
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Descr,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getComments = `-- name: GetComments :many
select id, userid, postid, parentid, content, created
from Comment
where postId = $1
`

func (q *Queries) GetComments(ctx context.Context, postid int32) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getComments, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Postid,
			&i.Parentid,
			&i.Content,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNewsfeed = `-- name: GetNewsfeed :one
select array(
               select id
               from Post
               where userId = any ($1::int[])
               order by created desc
               limit $2 offset $3
           )
`

type GetNewsfeedParams struct {
	Column1 []int32 `json:"column_1"`
	Limit   int32   `json:"limit"`
	Offset  int32   `json:"offset"`
}

func (q *Queries) GetNewsfeed(ctx context.Context, arg GetNewsfeedParams) (interface{}, error) {
	row := q.db.QueryRow(ctx, getNewsfeed, arg.Column1, arg.Limit, arg.Offset)
	var exists interface{}
	err := row.Scan(&exists)
	return exists, err
}

const getPhoto = `-- name: GetPhoto :one
select id, userid, albumid, url, created
from Photo
where id = $1
limit 1
`

func (q *Queries) GetPhoto(ctx context.Context, id int32) (Photo, error) {
	row := q.db.QueryRow(ctx, getPhoto, id)
	var i Photo
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Albumid,
		&i.Url,
		&i.Created,
	)
	return i, err
}

const getPhotoByUserId = `-- name: GetPhotoByUserId :many
select id, userid, albumid, url, created
from Photo
where UserId = $1
`

func (q *Queries) GetPhotoByUserId(ctx context.Context, userid int32) ([]Photo, error) {
	rows, err := q.db.Query(ctx, getPhotoByUserId, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Photo{}
	for rows.Next() {
		var i Photo
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Albumid,
			&i.Url,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPost = `-- name: GetPost :one
SELECT id, userid, created, tags, content, atchtype, atchid, atchurl, reaction, cmtcount
FROM Post
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Created,
		&i.Tags,
		&i.Content,
		&i.Atchtype,
		&i.Atchid,
		&i.Atchurl,
		&i.Reaction,
		&i.Cmtcount,
	)
	return i, err
}

const getPostByUserId = `-- name: GetPostByUserId :one
select array(select id from Post where userId = $1 order by created desc)
`

func (q *Queries) GetPostByUserId(ctx context.Context, userid int32) (interface{}, error) {
	row := q.db.QueryRow(ctx, getPostByUserId, userid)
	var exists interface{}
	err := row.Scan(&exists)
	return exists, err
}

const getReaction = `-- name: GetReaction :one
select reaction
from Post
where id = $1
`

func (q *Queries) GetReaction(ctx context.Context, id int32) ([]int32, error) {
	row := q.db.QueryRow(ctx, getReaction, id)
	var reaction []int32
	err := row.Scan(&reaction)
	return reaction, err
}

const getReactionOfUser = `-- name: GetReactionOfUser :one
select typ
from Reaction
where userId = $1
  and postId = $2
limit 1
`

type GetReactionOfUserParams struct {
	Userid int32 `json:"userid"`
	Postid int32 `json:"postid"`
}

func (q *Queries) GetReactionOfUser(ctx context.Context, arg GetReactionOfUserParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getReactionOfUser, arg.Userid, arg.Postid)
	var typ pgtype.Text
	err := row.Scan(&typ)
	return typ, err
}

const getReactions = `-- name: GetReactions :many
select userid, postid, typ
from Reaction
where postId = $1
`

func (q *Queries) GetReactions(ctx context.Context, postid int32) ([]Reaction, error) {
	rows, err := q.db.Query(ctx, getReactions, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reaction{}
	for rows.Next() {
		var i Reaction
		if err := rows.Scan(&i.Userid, &i.Postid, &i.Typ); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
update Post
set Tags=$2,
    Content=$3,
    AtchType=$4,
    AtchId=$5,
    AtchUrl=$6
where id = $1
RETURNING id, userid, created, tags, content, atchtype, atchid, atchurl, reaction, cmtcount
`

type UpdatePostParams struct {
	ID       int32       `json:"id"`
	Tags     pgtype.Text `json:"tags"`
	Content  pgtype.Text `json:"content"`
	Atchtype pgtype.Text `json:"atchtype"`
	Atchid   pgtype.Int4 `json:"atchid"`
	Atchurl  pgtype.Text `json:"atchurl"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.ID,
		arg.Tags,
		arg.Content,
		arg.Atchtype,
		arg.Atchid,
		arg.Atchurl,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Created,
		&i.Tags,
		&i.Content,
		&i.Atchtype,
		&i.Atchid,
		&i.Atchurl,
		&i.Reaction,
		&i.Cmtcount,
	)
	return i, err
}
