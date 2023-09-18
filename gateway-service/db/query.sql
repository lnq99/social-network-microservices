-- name: GetAccounts :many
select *
from Account;

-- name: GetAccount :one
select *
from Account
where id = $1
limit 1;

-- name: GetAccountByEmail :one
select *
from Account
where email = $1
limit 1;

-- name: CreateAccount :one
insert into Account (email, password, role)
values ($1, $2, $3)
returning *;

-- email: UpdateAccount :execrows
update Account
set email    = coalesce(sqlc.narg(email), email),
    password = coalesce(sqlc.narg(password), password),
    role     = coalesce(sqlc.narg(role), role)
where id = sqlc.arg(id)
returning *;

-- name: DeleteAccount :execrows
delete
from Account
where id = $1
returning *;