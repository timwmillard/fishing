
CREATE SCHEMA IF NOT EXISTS fishing;

CREATE TABLE IF NOT EXISTS fishing.club (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  slug text NOT NULL,
  legal_name text NOT NULL DEFAULT '',
  logo_url text,
  custom_domain text,
  billing_address1 text NOT NULL DEFAULT '',
  billing_address2 text NOT NULL DEFAULT '',
  billing_suburb text NOT NULL DEFAULT '',
  billing_state text NOT NULL DEFAULT '',
  billing_postcode text NOT NULL DEFAULT '',
  settings json NOT NULL DEFAULT '{}',
  current_event bigint DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS fishing.event (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  slug text NOT NULL,
  start_timestamp timestamptz,
  end_timestamp timestamptz,
  location text NOT NULL DEFAULT '',
  settings json NOT NULL DEFAULT '{}',
  club_id bigint NOT NULL REFERENCES fishing.club(id)
);

ALTER TABLE fishing.club ADD FOREIGN KEY (current_event) REFERENCES fishing.event(id);

CREATE TABLE IF NOT EXISTS fishing.team (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  team_no int NOT NULL,
  name text NOT NULL DEFAULT '',
  boat_rego text NOT NULL DEFAULT '',
  event_id bigint NOT NULL REFERENCES fishing.event(id)
);

CREATE TABLE IF NOT EXISTS fishing.competitor  (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  competitor_no text NOT NULL DEFAULT '',
  first_name text NOT NULL DEFAULT '',
  last_name text NOT NULL DEFAULT '',
  email text NOT NULL DEFAULT '',
  address1 text NOT NULL DEFAULT '',
  address2 text NOT NULL DEFAULT '',
  suburb text NOT NULL DEFAULT '',
  state text NOT NULL DEFAULT '',
  postcode text NOT NULL DEFAULT '',
  mobile text NOT NULL DEFAULT '',
  event_id bigint NOT NULL REFERENCES fishing.event(id),
  team_id bigint REFERENCES fishing.team(id),
  user_id bigint,
  UNIQUE(event_id, competitor_no)
);

CREATE TABLE fishing.species (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  common_name text UNIQUE NOT NULL,
  scientific_name text NOT NULL DEFAULT '',
  slug text UNIQUE NOT NULL,
  photo_url text
);

CREATE TABLE fishing.catch (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  competitor_id bigint NOT NULL REFERENCES fishing.competitor(id),
  species_id bigint NOT NULL REFERENCES fishing.species(id),
  size int NOT NULL,
  caught_at timestamptz NOT NULL,
  bait text NOT NULL DEFAULT '',
  location text NOT NULL DEFAULT '',
  latitude numeric,
  longitude numeric,
  photo_url text,
  event_id bigint NOT NULL REFERENCES fishing.event(id)
);



-- CREATE TABLE user (
--   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
--   username text UNIQUE NOT NULL,
--   password text NOT NULL,
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



-- CREATE TABLE club_users (
--   club_id int,
--   user_id int,
--   admin smallint,
--   marshall smallint
-- );
