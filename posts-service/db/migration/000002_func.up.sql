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
