CREATE TABLE IF NOT EXISTS user (
    id INT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    email varchar(100),
    password varchar(255)
)
