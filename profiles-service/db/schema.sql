-- create type gender_t as enum('M','F');
-- create type relation_t as enum('friend','block','request', 'follow');

create table Profile
(
    id        serial primary key,
    name      text    not null,
    gender    char(1) not null,
    birthdate date,
    email     text    not null unique,
    phone     decimal(13) unique,
    intro     text        default '',
    avatarS   text        default '',
    avatarL   text        default '',
    created   timestamptz default now()
);

create table Relationship
(
    user1   int not null references Profile (id) on delete cascade,
    user2   int not null references Profile (id) on delete cascade,
    type    text        default '',
    other   text        default '',
    created timestamptz default now(),
    primary key (user1, user2)
);