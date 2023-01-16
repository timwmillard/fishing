
create schema if not exists fishing;

-- Club
create table if not exists fishing.club (
  id bigint generated always as identity primary key,
  name text not null,
  slug text not null,
  legal_name text not null default '',
  logo_url text,
  custom_domain text,
  billing_address1 text not null default '',
  billing_address2 text not null default '',
  billing_suburb text not null default '',
  billing_state text not null default '',
  billing_postcode text not null default '',
  settings json not null default '{}',
  current_event bigint default null,
    unique(slug)
);

-- Event
create table if not exists fishing.event (
  id bigint generated always as identity primary key,
  name text not null,
  slug text not null,
  start_timestamp timestamptz,
  end_timestamp timestamptz,
  location text not null default '',
  settings json not null default '{}',
  club_id bigint not null references fishing.club(id),
  unique(id, slug)
);

alter table fishing.club add foreign key (current_event) references fishing.event(id);

-- User
create table if not exists fishing.user (
  id bigint generated always as identity primary key,
  username text unique,
  password text,
  firstname text not null default '',
  lastname text not null default '',
  email text not null default '',
  mobile text not null default '',
  address1 text,
  address2 text,
  suburb text,
  state text,
  postcode text,
  meta json,
  settings json
);

-- Team
create table if not exists fishing.team (
  id bigint generated always as identity primary key,
  team_no text not null default '',
  name text not null default '',
  boat_rego text not null default '',
  event_id bigint not null references fishing.event(id)
);

create unique index team_event_id_team_no_key on fishing.team(event_id, team_no) where team_no <> '';

-- Competitor
create table if not exists fishing.competitor  (
  id bigint generated always as identity primary key,
  competitor_no text not null default '',
  first_name text not null default '',
  last_name text not null default '',
  email text not null default '',
  mobile text not null default '',
  phone text not null default '',
  address1 text not null default '',
  address2 text not null default '',
  suburb text not null default '',
  state text not null default '',
  postcode text not null default '',
  event_id bigint not null references fishing.event(id),
  team_id bigint references fishing.team(id),
  user_id bigint references fishing.user(id)
);

create unique index competitor_event_id_competitor_no_key on fishing.competitor(event_id, competitor_no) where competitor_no <> '';

-- Species
create table fishing.species (
  id bigint generated always as identity primary key,
  name text not null,
  scientific_name text not null default '',
  common_names text[],
  slug text not null,
  photo_url text,
  native boolean,
  unique(name),
  unique(slug)
);

-- Catch
create table fishing.catch (
  id bigint generated always as identity primary key,
  competitor_id bigint not null references fishing.competitor(id),
  species_id bigint not null references fishing.species(id),
  size int not null,
  caught_at timestamptz not null,
  bait text not null default '',
  location text not null default '',
  latitude numeric,
  longitude numeric,
  photo_url text,
  event_id bigint not null references fishing.event(id)
);

-- create type user_role as enum (
--   'Admin',
--   'Marshal'
-- );

-- create table club_user (
--   club_id int,
--   user_id int,
--   roles []roles
-- );
