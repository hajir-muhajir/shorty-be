-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS health_cheks(
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS health_checks;
-- +goose StatementEnd
