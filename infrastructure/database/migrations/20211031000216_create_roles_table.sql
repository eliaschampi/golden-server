-- migrate:up
CREATE TABLE IF NOT EXISTS roles (
    code UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL,
    description VARCHAR(255)
);
-- migrate:down
DROP TABLE IF EXISTS roles;