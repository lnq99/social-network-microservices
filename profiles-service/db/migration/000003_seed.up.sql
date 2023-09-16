-- copy Profile from '/tmp/csv/profile.csv' delimiter ',' csv header;
--
--
-- begin;
-- create temp table ttmp on commit drop
-- as
-- select *
-- from Relationship
--                        with no data;
--
-- copy ttmp from '/tmp/csv/relationship.csv' delimiter ',' csv header;
--
-- insert into Relationship
-- select distinct on (user1, user2) *
-- from ttmp;
-- commit;
--
--
-- select setval('profile_id_seq', (select max(id) from Profile));


INSERT INTO Profile (name, gender, birthdate, email, phone, intro, avatarS, avatarL)
VALUES ('Sarah Johnson', 'F', '1992-03-14', 'sarahj@example.com', '+1122334455', 'I love to travel!',
        'https://example.com/avatar1_thumb.jpg', 'https://example.com/avatar1.jpg'),
       ('Alex Rodriguez', 'M', '1988-08-08', 'arod@example.com', NULL, 'I am a software developer',
        'https://example.com/avatar2_thumb.jpg', 'https://example.com/avatar2.jpg'),
       ('Emily Chen', 'F', '1996-01-23', 'emilyc@example.com', '+16505551212', 'I enjoy painting and drawing',
        'https://example.com/avatar3_thumb.jpg', 'https://example.com/avatar3.jpg'),
       ('Michael Davis', 'M', '1999-12-25', 'michaeld@example.com', '+14155552671', 'I like to play basketball',
        'https://example.com/avatar4_thumb.jpg', 'https://example.com/avatar4.jpg'),
       ('Jessica Kim', 'F', '1998-05-01', 'jessicakim@example.com', NULL, NULL, 'https://example.com/avatar5_thumb.jpg',
        'https://example.com/avatar5.jpg');

INSERT INTO Relationship (user1, user2, type, other)
VALUES (1, 3, 'friend', ''),
       (3, 1, 'friend', ''),
       (1, 4, 'friend', 'dating'),
       (4, 1, 'friend', 'dating'),
       (2, 3, 'friend', 'worked together on project Y'),
       (3, 2, 'friend', ''),
       (2, 4, 'request', ''),
       (2, 5, 'friend', ''),
       (5, 2, 'friend', ''),
       (3, 4, 'friend', 'engaged'),
       (4, 3, 'friend', 'engaged'),
       (3, 5, 'friend', ''),
       (5, 3, 'friend', ''),
       (4, 5, 'friend', ''),
       (5, 4, 'friend', '');