-- migrate:up
CREATE TABLE IF NOT EXISTS cashes
(
    code        uuid PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_dni    BPCHAR(8)      NOT NULL,
    amount      NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    state       bool           NOT NULL default true,
    description VARCHAR(300),
    created_at  timestamp(6)            DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6)            DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (user_dni) REFERENCES users (dni)
);

-- migrate:down
DROP TABLE IF EXISTS cashes;
