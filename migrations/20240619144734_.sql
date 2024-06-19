-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content VARCHAR(2000) NOT NULL,
    comments_disabled BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    content VARCHAR(2000) NOT NULL,
    parent_id INT,
    FOREIGN KEY (post_id) REFERENCES posts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
