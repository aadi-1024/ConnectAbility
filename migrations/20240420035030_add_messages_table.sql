-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Messages (
    Id SERIAL PRIMARY KEY,
    Chat_Id INT NOT NULL,
    Content VARCHAR(256) NOT NULL,
    Sent_By INT NOT NULL,
    Timestamp TIMESTAMP NOT NULL,
    FOREIGN KEY (Chat_Id) REFERENCES Chats(Id),
    FOREIGN KEY (Sent_By) REFERENCES Users(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Messages;
SELECT 'down SQL query';
-- +goose StatementEnd
