-- name: GetProfile :one
SELECT *
FROM Profile
WHERE id = $1
LIMIT 1;


-- name: GetProfileByEmail :one
SELECT *
FROM Profile
WHERE email = $1
LIMIT 1;


-- name: CreateProfile :one
INSERT INTO Profile (name, gender, birthdate, email, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;


-- name: UpdateProfile :one
UPDATE Profile
SET name      = COALESCE(sqlc.narg(name), name),
    gender    = COALESCE(sqlc.narg(gender), gender),
    birthdate = COALESCE(sqlc.narg(birthdate), birthdate),
    phone     = COALESCE(sqlc.narg(phone), phone),
    intro     = COALESCE(sqlc.narg(intro), intro),
    avatarS   = COALESCE(sqlc.narg(avatarS), avatarS),
    avatarL   = COALESCE(sqlc.narg(avatarL), avatarL)
WHERE id = sqlc.arg(id)
RETURNING *;


-- name: SearchName :one
SELECT search_name($1, $2);


-- name: CreateRelationship :one
INSERT INTO Relationship (user1, user2, typ, other)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetRelationship :one
SELECT *
FROM Relationship
WHERE user1 = $1
  AND user2 = $2
LIMIT 1;


-- name: GetRelationshipsFrom :many
SELECT *
FROM Relationship
WHERE user1 = $1;


-- name: GetRelationshipsTo :many
SELECT *
FROM Relationship
WHERE user2 = $1;


-- name: GetFriendRelationships :many
SELECT *
FROM Relationship
WHERE user1 = $1
  and typ = 'friend';


-- name: GetRequestRelationships :many
SELECT *
FROM Relationship
WHERE user1 = $1
  and typ = 'request';


-- name: UpdateRelationship :one
insert into Relationship(user1, user2, typ, other)
values ($1, $2, sqlc.narg(typ), sqlc.narg(other))
on conflict (user1, user2) do update set typ   = COALESCE(sqlc.narg(typ), Relationship.typ),
                                         other = COALESCE(sqlc.narg(other), Relationship.other)
returning *;

-- UPDATE Relationship
-- SET typ  = COALESCE(sqlc.narg(typ), typ),
--     other = COALESCE(sqlc.narg(other), other)
-- WHERE user1 = sqlc.arg(user1)
--   AND user2 = sqlc.arg(user2)
-- RETURNING *;


-- name: DeleteRelationship :one
DELETE
FROM Relationship
WHERE user1 = $1
  AND user2 = $2
RETURNING *;


-- name: FriendsDetail :one
select friends_json($1);


-- name: MutualFriends :one
select mutual_friends($1, $2)::int[];
