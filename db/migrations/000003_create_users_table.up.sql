CREATE TABLE users (
  id      SERIAL NOT NULL UNIQUE PRIMARY KEY,
  username VARCHAR not NULL UNIQUE,
  password VARCHAR not null
);
