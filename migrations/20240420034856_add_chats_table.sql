-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Chats (
    Id SERIAL PRIMARY KEY,
    Member1 INT NOT NULL,
    Member2 INT NOT NULL,
    FOREIGN KEY (Member1) REFERENCES Users(Id),
    FOREIGN KEY (Member2) REFERENCES Users(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Chats;
-- +goose StatementEnd
