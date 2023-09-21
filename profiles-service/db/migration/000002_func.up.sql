create or replace function friends_json(urid int)
    returns jsonb as
$$
with t as (select id, name, avatarS
           from Profile
           where id in (select user2
                        from Relationship
                        where user1 = urid
                          and typ = 'friend'))
select jsonb_agg(t)
from t;
$$ language sql;


--------------------------------------------------
create or replace function mutual_friends(u1 int, u2 int)
    returns int[]
as
$$
select array(
               select R1.user2 friend
               from Relationship R1
                        join (select user2
                              from Relationship
                              where user1 = u2 and typ = 'friend') as R2 on R1.user2 = R2.user2
               where user1 = u1
                 and typ = 'friend'
           )
$$ language sql;


--------------------------------------------------
create or replace function search_name(u int, pattern text)
    returns jsonb as
$$
with t as (select id, cardinality(mutual_friends(u, id)) as mutual
           from Profile
           where lower(name) like format('%%%s%%', lower(pattern))
             and id not in (select id
                            from Relationship
                            where user2 = u
                              and typ = 'block')),
     rel as ((select id,
                     mutual,
                     case typ
                         when 'request' then 'follow'
                         else typ
                         end
              from t
                       left join Relationship r
                                 on r.user1 = u and r.user2 = id)
             union
             (select id, mutual, typ
              from t
                       left join Relationship r
                                 on r.user1 = id and r.user2 = u))
select jsonb_agg(jsonb_build_object('id', id, 'mutual', mutual, 'type', typ))
from rel;
$$ language sql;


--------------------------------------------------
create or replace function n_mutual_friends(u1 int, u2 int)
    returns bigint
as
$$
select count(*)
from Relationship R1
         join (select user2
               from Relationship
               where user1 = u2 and typ = 'friend') as R2 on R1.user2 = R2.user2
where user1 = u1
  and typ = 'friend'
$$ language sql;
