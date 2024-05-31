CREATE TABLE  users (
    id INT PRIMARY KEY,
    email VARCHAR(255),
    password varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);