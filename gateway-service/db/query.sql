-- name: GetAccount :many
select *
from Account;

-- name: GetUser :one
select *
from Account
where id = $1
limit 1;

-- name: GetUserByName :one
select *
from Account
where name = $1
limit 1;

-- name: CreateUser :one
insert into Account (name, role, password)
values ($1, $2, $3)
returning *;

-- name: UpdateUser :execrows
update Account
set name     = coalesce(sqlc.narg(name), name),
    role     = coalesce(sqlc.narg(role), role),
    password = coalesce(sqlc.narg(password), password)
where id = sqlc.arg(id)
returning *;

-- name: DeleteUser :execrows
delete
from Account
where id = $1
returning *;