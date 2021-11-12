-- migrate:up
CREATE TABLE IF NOT EXISTS materials
(
    code          uuid PRIMARY KEY default uuid_generate_v4(),
    name          VARCHAR(150) not null,
    description   VARCHAR(300),
    category_code uuid         NOT NULL,
    created_at    timestamp(8)     DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (category_code)
        REFERENCES categories (code)
);

-- migrate:down
DROP TABLE IF EXISTS materials;
