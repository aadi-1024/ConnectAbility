-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Teams (
    Id SERIAL PRIMARY KEY,
    Owner_Id INT NOT NULL,
    Name VARCHAR(64) NOT NULL,
    CREATION_TIME TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Teams;
-- +goose StatementEnd
