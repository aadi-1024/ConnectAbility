-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Save_List (
    User_Id INT,
    Saved_Id INT,
    FOREIGN KEY (User_Id) REFERENCES Users(Id),
    FOREIGN KEY (Saved_Id) REFERENCES Users(Id),
    PRIMARY KEY (User_Id, Saved_Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Save_List;
-- +goose StatementEnd
