-- name: GetPost :one
SELECT *
FROM Post
WHERE id = $1
LIMIT 1;


-- name: GetPostByUserId :one
select array(select id from Post where userId = $1 order by created desc);


-- name: CreatePost :one
INSERT INTO Post (userId, tags, content, atchType, atchId, atchUrl)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: UpdatePost :one
update Post
set Tags=$2,
    Content=$3,
    AtchType=$4,
    AtchId=$5,
    AtchUrl=$6
where id = $1
RETURNING *;


-- name: DeletePost :one
delete
from Post
where id = $2
  and userId = $1
returning *;


-- name: GetReaction :one
select reaction
from Post
where id = $1;


-- name: GetComments :many
select *
from Comment
where postId = $1;


-- name: CreateComment :one
insert into Comment(userId, postId, parentId, content)
values ($1, $2, $3, $4)
returning *;


-- name: GetReactions :many
select *
from Reaction
where postId = $1;


-- name: GetReactionOfUser :one
select typ
from Reaction
where userId = $1
  and postId = $2
limit 1;


-- name: CreateReaction :one
insert into Reaction(userId, postId, typ)
values ($1, $2, $3)
on conflict (userId, postId) do update set typ = $3
returning *;


-- name: DeleteReaction :one
delete
from Reaction
where userId = $1
  and postId = $2
returning *;


-- name: GetPhoto :one
select *
from Photo
where id = $1
limit 1;


-- name: GetPhotoByUserId :many
select *
from Photo
where UserId = $1;


-- name: CreatePhoto :one
insert into Photo(userId, albumId, url)
values ($1, $2, $3)
returning *;


-- name: GetAlbum :one
select *
from Album
where id = $1
limit 1;


-- name: GetAlbumByUserId :many
select *
from Album
where UserId = $1;


-- name: CreateAlbum :one
insert into Album(id, userId, descr, created)
values ($1, $2, $3, $4)
returning *;


-- name: GetNewsfeed :one
select array(
               select id
               from Post
               where userId = any ($1::int[])
               order by created desc
               limit $2 offset $3
           );

-- select feed($1::int[], $2, $3);
