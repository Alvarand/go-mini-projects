-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS short_url (
    id int NOT NULL,
    short_url text PRIMARY KEY,
    url text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS short_url;
-- +goose StatementEnd
