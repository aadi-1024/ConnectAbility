-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE HasSkill (
    User_id INT,
    Skill_id INT,
    PRIMARY KEY (User_id, Skill_id),
    FOREIGN KEY (User_id) REFERENCES Users(Id),
    FOREIGN KEY (Skill_id) REFERENCES Skills(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE HasSkill;
-- +goose StatementEnd
