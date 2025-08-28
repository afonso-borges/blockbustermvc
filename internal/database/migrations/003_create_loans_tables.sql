-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS loans (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),

  movie_id UUID NOT NULL,
  user_id UUID NOT NULL,
  borrowed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  returned_at TIMESTAMPTZ,
  status VARCHAR(20) NOT NULL DEFAULT 'active',

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT fk_loans_movie_id FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
  CONSTRAINT fk_loans_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS loans;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
