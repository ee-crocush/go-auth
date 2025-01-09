CREATE TABLE IF NOT EXISTS blacklist (
    id SERIAL PRIMARY KEY,
    access_token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_blacklist_access_token ON blacklist (access_token);