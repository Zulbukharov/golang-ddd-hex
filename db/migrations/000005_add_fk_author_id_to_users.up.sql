ALTER TABLE posts ADD CONSTRAINT fk_author_id_to_users FOREIGN KEY (author_id) REFERENCES users(id);
