CREATE TABLE Account
(
    "id"         serial PRIMARY KEY,
    "email"      text UNIQUE NOT NULL,
    "role"       text      DEFAULT '',
    "password"   text        NOT NULL,
    "created_at" timestamp DEFAULT (now())
);