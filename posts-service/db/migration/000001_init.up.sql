create table Post
(
    id       serial primary key,
    userId   int not null,
    created  timestamptz default now(),
    tags     text        default '',
    content  text,
    atchType text        default 'none',
    atchId   int         default 0,
    atchUrl  text        default '',
    reaction int[6]      default '{0,0,0,0,0,0}',
    cmtCount int         default 0
);

create table Comment
(
    id       bigserial primary key,
    userId   int not null,
    postId   int not null references Post (id) on delete cascade,
    parentId int,
    content  text,
    created  timestamptz default now()
);

create table Reaction
(
    userId int not null,
    postId int not null references Post (id) on delete cascade,
    typ    text default 'like',
    primary key (userId, postId)
);

create table Album
(
    id      serial primary key,
    userId  int not null,
    descr   text        default '',
    created timestamptz default now()
);

create table Photo
(
    id      serial primary key,
    userId  int not null,
    albumId int not null references Album (id) on delete cascade,
    url     text,
    created timestamptz default now()
);
