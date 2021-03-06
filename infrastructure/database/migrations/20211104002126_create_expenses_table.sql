-- migrate:up
CREATE TABLE IF NOT EXISTS expenses (
    code UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    description VARCHAR(350) NOT NULL,
    user_dni CHAR(8) NOT NULL,
    cash_code UUID NOT NULL,
    total NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    created_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (user_dni) REFERENCES users (dni),
    FOREIGN KEY (cash_code) REFERENCES cashes (code)
);
-- migrate:down
DROP TABLE IF EXISTS expenses;