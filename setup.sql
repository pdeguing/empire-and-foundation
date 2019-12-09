create table users (
  id serial primary key,
  uuid varchar(64) not null unique,
  name varchar(255),
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp not null
);
create table planets (
  id serial primary key,
  uuid varchar(64) not null unique,
  metal_stock bigint not null,
  metal_mine integer not null,
  user_id integer references users(id),
  created_at timestamp not null,
  last_metal_update timestamp not null
);
CREATE TABLE sessions (
  token TEXT PRIMARY KEY,
  data BYTEA NOT NULL,
  expiry TIMESTAMPTZ NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);