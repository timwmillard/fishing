
-- CREATE TABLE competitions (
--   id uuid NOT NULL PRIMARY KEY,
--   organisation_id int,
--   short_name varchar(63),
--   name varchar(255),
--   logo_url varchar(255),
--   custom_domain varchar(255),
--   current_event int,
--   settings json
-- );

-- CREATE TABLE events (
--   id uuid NOT NULL PRIMARY KEY,
--   competition_id int,
--   slug varchar(50),
--   name varchar(255),
--   start_date date,
--   end_date date,
--   location varchar(255),
--   status int,
--   settings json
-- );

-- CREATE TABLE team (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   team_no int,
--   name varchar(255),
--   boat_rego varchar(20)
-- );

CREATE TABLE IF NOT EXISTS competitors  (
  id uuid NOT NULL PRIMARY KEY,
  -- event_id uuid NOT NULL,
  competitor_no varchar(255) NOT NULL DEFAULT '',
  firstname varchar(255) NOT NULL DEFAULT '',
  lastname varchar(255) NOT NULL DEFAULT '',
  email varchar(255) NOT NULL DEFAULT '',
  address1 varchar(255) NOT NULL DEFAULT '',
  address2 varchar(255) NOT NULL DEFAULT '',
  suburb varchar(255) NOT NULL DEFAULT '',
  state varchar(255) NOT NULL DEFAULT '',
  postcode varchar(20) NOT NULL DEFAULT '',
  phone varchar(20) NOT NULL DEFAULT '',
  mobile varchar(20) NOT NULL DEFAULT ''
  -- paid smallint DEFAULT 0,
  -- registered smallint DEFAULT 0,
  -- checkin smallint DEFAULT 0,
  -- ticket int DEFAULT 0,
  -- team_id int,
  -- user_id int
);

-- CREATE TABLE catches (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   competitor_id int,
--   species_id int,
--   size int,
--   caught_at datetime,
--   bait varchar(255) NOT NULL DEFAULT '',
--   location varchar(255) NOT NULL DEFAULT '',
--   latitude double NOT NULL DEFAULT 0.0,
--   longitude double NOT NULL DEFAULT 0.0,
--   marshall varchar(255) NOT NULL DEFAULT '',
--   marshall_id int,
--   status int DEFAULT 0
-- );

-- CREATE TABLE species (
--   id uuid NOT NULL PRIMARY KEY,
--   slug varchar(50) UNIQUE NOT NULL,
--   common_name varchar(255) UNIQUE NOT NULL,
--   scientific_name varchar(255) NOT NULL DEFAULT '',
--   photo_url varchar(255) NOT NULL DEFAULT ''
-- );

-- CREATE TABLE users (
--   id uuid NOT NULL PRIMARY KEY,
--   username varchar(255) UNIQUE NOT NULL,
--   password varchar(255) NOT NULL,
--   firstname varchar(255),
--   lastname varchar(255),
--   email varchar(255),
--   mobile varchar(255),
--   api_token varchar(255),
--   address1 varchar(255),
--   address2 varchar(255),
--   suburb varchar(255),
--   state varchar(255),
--   postcode varchar(255),
--   stripe_billing_id varchar(255),
--   settings json
-- );

-- CREATE TABLE tickets (
--   id uuid NOT NULL PRIMARY KEY,
--   event_id int,
--   name varchar(255),
--   start_competitor_no int,
--   next_competitor_no int,
--   price int,
--   stripe_product_id int,
--   max_no_competitors int
-- );

-- CREATE TABLE clubs (
--   id uuid NOT NULL PRIMARY KEY,
--   name varchar(255),
--   billing_address1 varchar(255),
--   billing_address2 varchar(255),
--   billing_suburb varchar(255),
--   billing_state varchar(255),
--   billing_postcode varchar(255),
--   stripe_billing_id varchar(255),
--   owner int,
--   settings json
-- );

-- CREATE TABLE organisation_users (
--   organisation_id int,
--   user_id int,
--   admin smallint,
--   marshall smallint
-- );

-- ALTER TABLE competitions ADD FOREIGN KEY (organisation_id) REFERENCES organisations (id);

-- ALTER TABLE competitions ADD FOREIGN KEY (current_event) REFERENCES events (id);

-- ALTER TABLE events ADD FOREIGN KEY (competition_id) REFERENCES competitions (id);

-- ALTER TABLE team ADD FOREIGN KEY (event_id) REFERENCES events (id);

-- ALTER TABLE competitors ADD FOREIGN KEY (event_id) REFERENCES events (id);

-- /* ALTER TABLE competitors ADD FOREIGN KEY (ticket) REFERENCES tickets (id); */

-- ALTER TABLE competitors ADD FOREIGN KEY (team_id) REFERENCES team (id);

-- ALTER TABLE competitors ADD FOREIGN KEY (user_id) REFERENCES users (id);

-- ALTER TABLE catches ADD FOREIGN KEY (event_id) REFERENCES events (id);

-- ALTER TABLE catches ADD FOREIGN KEY (competitor_id) REFERENCES competitors (id);

-- ALTER TABLE catches ADD FOREIGN KEY (species_id) REFERENCES species (id);

-- ALTER TABLE catches ADD FOREIGN KEY (marshall_id) REFERENCES users (id);

-- ALTER TABLE tickets ADD FOREIGN KEY (event_id) REFERENCES events (id);

-- ALTER TABLE organisations ADD FOREIGN KEY (owner) REFERENCES users (id);

-- ALTER TABLE organisation_users ADD FOREIGN KEY (organisation_id) REFERENCES organisations (id);

-- ALTER TABLE organisation_users ADD FOREIGN KEY (user_id) REFERENCES users (id);
