CREATE TABLE IF NOT EXISTS profiles (
    id int PRIMARY KEY,
    fname VARCHAR(255),
    lname VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255),
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users(id)
);