-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS movies (
  id UUID PRIMARY KEY NOT NULL  DEFAULT gen_random_uuid(),

  name VARCHAR(50) UNIQUE NOT NULL,
  director TEXT UNIQUE NOT NULL,
  year INTEGER NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()

);
---- create above / drop below ----

DROP TABLE IF EXISTS movies;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
