CREATE DATABASE go-rest-api;

-- Set cursor to go-rest-api
\c go-rest-api

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(200),
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users (name, email) VALUES ('Alice Smith', 'alice@example.com');
INSERT INTO users (name, email) VALUES ('Bob Johnson', 'bob@example.com');
INSERT INTO users (name, email) VALUES ('Carol Davis', 'carol@example.com');
INSERT INTO users (name, email) VALUES ('David Brown', 'david@example.com');
INSERT INTO users (name, email) VALUES ('Eve White', 'eve@example.com');

INSERT INTO posts (user_id, title, content) VALUES (1, 'First Post', 'Content of the first post');
INSERT INTO posts (user_id, title, content) VALUES (2, 'Second Post', 'Content of the second post');
INSERT INTO posts (user_id, title, content) VALUES (3, 'Third Post', 'Content of the third post');
INSERT INTO posts (user_id, title, content) VALUES (4, 'Fourth Post', 'Content of the fourth post');
INSERT INTO posts (user_id, title, content) VALUES (5, 'Fifth Post', 'Content of the fifth post');
