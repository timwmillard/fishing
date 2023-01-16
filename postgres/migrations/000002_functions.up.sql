-- -- next_competitor_no returns the next avaliable competitor for a given event.
-- create or replace function fishing.next_competitor_no(event_id_input bigint)
-- returns int
-- language plpgsql
-- as
-- $$
-- declare next_competitor_no int;
-- begin
--     select competitor_no::int into next_competitor_no
--     from fishing.competitor
--     where event_id = event_id_input
--     order by competitor_no desc
--     limit 1;

--     return next_competitor_no + 1;
    
-- end;
-- $$;

-- alter table fishing.competitor
--     alter column competitor_no
--     default next_competitor_no(event_id);

-- -- current_event returns the current event for the give club.
-- create or replace function fishing.current_event(club_input text)
-- returns int
-- language plpgsql
-- as
-- $$
-- declare current_event_id bigint;
-- begin
--     select current_event into current_event_id
--     from fishing.club
--     where slug = club_input;

--     return current_event_id;
    
-- end;
-- $$;
