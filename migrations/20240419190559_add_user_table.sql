-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    Id SERIAL PRIMARY KEY,
    Email VARCHAR(40) UNIQUE,
    Phone_No VARCHAR(15) UNIQUE,
    First_name VARCHAR(48) NOT NULL,
    Last_Name VARCHAR(48),
    About VARCHAR(1024),
    Profile_Pic VARCHAR(32),
    Resume_Link VARCHAR(32),
    Linkedin VARCHAR(256),
    Github VARCHAR(128),
    Website VARCHAR(256),
    Location_city VARCHAR(64),
    Location_country VARCHAR(64),
    Location_area VARCHAR(64),
    Location_pin VARCHAR(16)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
