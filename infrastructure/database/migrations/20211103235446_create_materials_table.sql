-- migrate:up
CREATE TABLE IF NOT EXISTS materials (
    code UUID PRIMARY KEY default uuid_generate_v4(),
    name VARCHAR(150) not null,
    description VARCHAR(300),
    category_code UUID NOT NULL,
    created_at timestamp(6) DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (category_code) REFERENCES categories (code)
);
-- migrate:down
DROP TABLE IF EXISTS materials;