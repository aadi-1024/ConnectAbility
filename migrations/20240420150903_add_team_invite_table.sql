-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE team_invites (
    Id SERIAL PRIMARY KEY,
    Team_Id INT NOT NULL,
    Invited_Id INT NOT NULL,
    FOREIGN KEY (Team_Id) REFERENCES teams(Id),
    FOREIGN KEY (Invited_Id) REFERENCES users(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE team_invite;
-- +goose StatementEnd
