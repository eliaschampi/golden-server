-- migrate:up
CREATE TABLE IF NOT EXISTS categories
(
    code        uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(120) NOT NULL,
    description VARCHAR(300)
);

-- migrate:down
DROP TABLE IF EXISTS categories;
