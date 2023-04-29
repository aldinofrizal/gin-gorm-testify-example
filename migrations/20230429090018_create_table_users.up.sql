CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email varchar NOT NULL UNIQUE,
  name varchar NOT NULL,
  password varchar NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
)