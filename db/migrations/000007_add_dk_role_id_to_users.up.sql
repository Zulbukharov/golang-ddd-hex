BEGIN;

ALTER TABLE users ADD COLUMN role_id INTEGER;
ALTER TABLE users ADD CONSTRAINT fk_role_id_to_users FOREIGN KEY (role_id) REFERENCES roles(id);

COMMIT;