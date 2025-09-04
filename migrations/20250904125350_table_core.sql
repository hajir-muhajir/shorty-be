-- +goose Up
-- +goose StatementBegin
-- users table
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- links table
CREATE TABLE IF NOT EXISTS links(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    original_url TEXT NOT NULL,
    alias VARCHAR(64) NOT NULL UNIQUE,
    password_hash TEXT NULL,
    expires_at TIMESTAMPTZ NULL,
    max_clicks INT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    click_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_links_user_created ON links(user_id, created_at);
CREATE INDEX IF NOT EXISTS idx_links_active_exp ON links(is_active, expires_at);

-- clicks table
CREATE TABLE IF NOT EXISTS clicks(
    id BIGSERIAL PRIMARY KEY,
    link_id UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,
    ts TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ip_hash TEXT NOT NULL,
    country TEXT NULL,
    city TEXT NULL,
    referrer TEXT NULL,
    ua TEXT NULL,
    device TEXT NULL,
    os TEXT NULL,
    browser TEXT NULL
);

CREATE INDEX IF NOT EXISTS idx_clicks_links_ts ON clicks(link_id, ts);
CREATE INDEX IF NOT EXISTS idx_clicks_links_ip ON clicks(link_id, ip_hash);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clicks;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
