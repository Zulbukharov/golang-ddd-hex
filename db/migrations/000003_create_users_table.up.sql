CREATE TABLE users (
  id      serial NOT NULL,
  username VARCHAR not NULL UNIQUE,
  password VARCHAR not null
);
