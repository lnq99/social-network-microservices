create or replace function reaction_update()
    returns trigger as
$$
declare
    o int; n int; postId int; r int[];
begin

    if (old is null) then
        postId := new.postId;
    else
        postId := old.postId;
    end if;

    r := (select reaction from Post where id = postId);
    -- raise notice '%', r;
    o := (select array_position(array ['like','love','haha','wow','sad','angry'], old.typ::text));
    n := (select array_position(array ['like','love','haha','wow','sad','angry'], new.typ::text));

    if (o is null) then
        -- raise notice 'insert';
        r[n] := r[n] + 1;
    elsif (n is null) then
        -- raise notice 'delete';
        r[o] := r[o] - 1;
    else
        -- raise notice 'update';
        r[o] := r[o] - 1;
        r[n] := r[n] + 1;
    end if;

    update Post set reaction = r where id = postId;
    -- raise notice '%', r;
    return new;
end;
$$ language plpgsql;


--------------------------------------------------
create trigger reaction_type_update
    after insert or update or delete
    on Reaction
    for each row
execute function reaction_update();


--------------------------------------------------
create or replace function feed(ids_arr int[], lim int default 10, offs int default 0)
    returns int[]
as
$$
select array(
               select id
               from Post
               where userId = any (ids_arr)
               order by created desc
               limit lim offset offs
           )
$$ language sql;
