
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE INDEX users_id_idx ON users(id);
CREATE INDEX users_username_idx ON users(username);

CREATE TABLE notes(
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    title TEXT NOT NULL,
    contents TEXT NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);

CREATE INDEX notes_id_idx ON notes(id);
CREATE INDEX notes_user_id_idx ON notes(user_id);
