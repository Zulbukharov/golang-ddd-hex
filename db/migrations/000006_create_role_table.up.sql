BEGIN;

CREATE TABLE roles (
                       id        SERIAL NOT NULL UNIQUE PRIMARY KEY,
                       role_name VARCHAR not NULL
);

INSERT INTO roles(id, role_name) VALUES
    (1, 'USER'),
    (2, 'AUTHOR'),
    (3, 'ADMIN');

COMMIT;