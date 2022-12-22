
CREATE SCHEMA IF NOT EXISTS fishing;

-- CREATE TABLE competition (
--   id uuid NOT NULL PRIMARY KEY,
--   organisation_id int,
--   short_name text,
--   name text,
--   logo_url text,
--   custom_domain text,
--   current_event int,
--   settings json
-- );

-- CREATE TABLE event (
--   id uuid NOT NULL PRIMARY KEY,
--   competition_id int,
--   slug varchar(50),
--   name text,
--   start_date date,
--   end_date date,
--   location text,
--   status int,
--   settings json
-- );

-- CREATE TABLE team (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   team_no int,
--   name text,
--   boat_rego varchar(20)
-- );

CREATE TABLE IF NOT EXISTS fishing.competitor  (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  -- event_id uuid NOT NULL,
  competitor_no text NOT NULL DEFAULT '',
  first_name text NOT NULL DEFAULT '',
  last_name text NOT NULL DEFAULT '',
  email text NOT NULL DEFAULT '',
  address1 text NOT NULL DEFAULT '',
  address2 text NOT NULL DEFAULT '',
  suburb text NOT NULL DEFAULT '',
  state text NOT NULL DEFAULT '',
  postcode text NOT NULL DEFAULT '',
  mobile text NOT NULL DEFAULT ''
  -- team_id int,
  -- user_id int
);

-- CREATE TABLE catch (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   competitor_id int,
--   species_id int,
--   size int,
--   caught_at datetime,
--   bait text NOT NULL DEFAULT '',
--   location text NOT NULL DEFAULT '',
--   latitude double NOT NULL DEFAULT 0.0,
--   longitude double NOT NULL DEFAULT 0.0,
--   marshall text NOT NULL DEFAULT '',
--   marshall_id int,
--   status int DEFAULT 0
-- );

-- CREATE TABLE species (
--   id uuid NOT NULL PRIMARY KEY,
--   slug varchar(50) UNIQUE NOT NULL,
--   common_name text UNIQUE NOT NULL,
--   scientific_name text NOT NULL DEFAULT '',
--   photo_url text NOT NULL DEFAULT ''
-- );

-- CREATE TABLE user (
--   id uuid NOT NULL PRIMARY KEY,
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

-- CREATE TABLE ticket (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   name text,
--   start_competitor_no int,
--   next_competitor_no int,
--   price int,
--   stripe_product_id int,
--   max_no_competitors int
-- );

-- CREATE TABLE club (
--   id uuid NOT NULL PRIMARY KEY,
--   name text,
--   billing_address1 text,
--   billing_address2 text,
--   billing_suburb text,
--   billing_state text,
--   billing_postcode text,
--   stripe_billing_id text,
--   owner int,
--   settings json
-- );

-- CREATE TABLE club_users (
--   club_id int,
--   user_id int,
--   admin smallint,
--   marshall smallint
-- );

-- ALTER TABLE competition ADD FOREIGN KEY (club_id) REFERENCES clubs (id);

-- ALTER TABLE competition ADD FOREIGN KEY (current_event) REFERENCES event (id);

-- ALTER TABLE event ADD FOREIGN KEY (competition_id) REFERENCES competition (id);

-- ALTER TABLE team ADD FOREIGN KEY (event_id) REFERENCES event (id);

-- ALTER TABLE competitor ADD FOREIGN KEY (event_id) REFERENCES event (id);

-- /* ALTER TABLE competitor ADD FOREIGN KEY (ticket) REFERENCES tickets (id); */

-- ALTER TABLE competitor ADD FOREIGN KEY (team_id) REFERENCES team (id);

-- ALTER TABLE competitor ADD FOREIGN KEY (user_id) REFERENCES users (id);

-- ALTER TABLE catche ADD FOREIGN KEY (event_id) REFERENCES event (id);

-- ALTER TABLE catche ADD FOREIGN KEY (competitor_id) REFERENCES competitor (id);

-- ALTER TABLE catche ADD FOREIGN KEY (species_id) REFERENCES species (id);

-- ALTER TABLE catche ADD FOREIGN KEY (marshall_id) REFERENCES user (id);

-- ALTER TABLE tickets ADD FOREIGN KEY (event_id) REFERENCES event (id);

-- ALTER TABLE club ADD FOREIGN KEY (owner) REFERENCES user (id);

-- ALTER TABLE club_user ADD FOREIGN KEY (club_id) REFERENCES clubs (id);

-- ALTER TABLE club_user ADD FOREIGN KEY (user_id) REFERENCES user (id);
