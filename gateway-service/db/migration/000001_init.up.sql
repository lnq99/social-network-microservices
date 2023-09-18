CREATE TABLE Account
(
    "id"         serial PRIMARY KEY,
    "email"      text UNIQUE NOT NULL,
    "password"   text        NOT NULL,
    "role"       text      DEFAULT '',
    "created_at" timestamp DEFAULT (now())
);