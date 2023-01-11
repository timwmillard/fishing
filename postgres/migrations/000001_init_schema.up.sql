
create schema if not exists fishing;

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
  current_event bigint default null
);

create table if not exists fishing.event (
  id bigint generated always as identity primary key,
  name text not null,
  slug text not null,
  start_timestamp timestamptz,
  end_timestamp timestamptz,
  location text not null default '',
  settings json not null default '{}',
  club_id bigint not null references fishing.club(id)
);

alter table fishing.club add foreign key (current_event) references fishing.event(id);

create table if not exists fishing.team (
  id bigint generated always as identity primary key,
  team_no int not null,
  name text not null default '',
  boat_rego text not null default '',
  event_id bigint not null references fishing.event(id)
);

create table if not exists fishing.competitor  (
  id bigint generated always as identity primary key,
  competitor_no text not null default '',
  first_name text not null default '',
  last_name text not null default '',
  email text not null default '',
  address1 text not null default '',
  address2 text not null default '',
  suburb text not null default '',
  state text not null default '',
  postcode text not null default '',
  mobile text not null default '',
  event_id bigint not null references fishing.event(id),
  team_id bigint references fishing.team(id),
  user_id bigint,
  unique(event_id, competitor_no)
);

create table fishing.species (
  id bigint generated always as identity primary key,
  common_name text unique not null,
  scientific_name text not null default '',
  slug text unique not null,
  photo_url text
);

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



-- create table user (
--   id bigint generated always as identity primary key,
--   username text unique not null,
--   password text not null,
--   firstname text,
--   lastname text,
--   email text,
--   mobile text,
--   api_token text,
--   address1 text,
--   address2 text,
--   suburb text,
--   state text,
--   postcode text,
--   stripe_billing_id text,
--   settings json
-- );



-- create table club_users (
--   club_id int,
--   user_id int,
--   admin smallint,
--   marshall smallint
-- );
