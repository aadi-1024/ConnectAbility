-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Team_Members (
    Team_Id INT,
    Member_Id INT,
    Joining_Date TIMESTAMP NOT NULL,
    Leaving_Date TIMESTAMP,
    PRIMARY KEY (Team_Id, Member_Id),
    FOREIGN KEY (Team_Id) REFERENCES Teams(Id),
    FOREIGN KEY (Member_Id) REFERENCES Users(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Team_Members;
-- +goose StatementEnd
