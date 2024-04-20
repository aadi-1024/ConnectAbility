-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Skills (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(64) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Skills;
-- +goose StatementEnd
