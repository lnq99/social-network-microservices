copy Post from '/tmp/csv/post.csv' delimiter ',' csv header;
copy Comment from '/tmp/csv/comment.csv' delimiter ',' csv header;
copy Reaction from '/tmp/csv/reaction.csv' delimiter ',' csv header;

copy Album from '/tmp/csv/album.csv' delimiter ',' csv header;
copy Photo from '/tmp/csv/photo.csv' delimiter ',' csv header;

select setval('photo_id_seq', (select max(id) from Photo));
select setval('album_id_seq', (select max(id) from Album));

select setval('post_id_seq', (select max(id) from Post));
select setval('comment_id_seq', (select max(id) from Comment));


-- insert into post (id, userid, created, tags, content, atchtype, atchid, atchurl, reaction, cmtcount)
-- values ();
-- insert into comment (id, userid, postid, parentid, content, created)
-- values ();
-- insert into reaction (userid, postid, type)
-- values ();