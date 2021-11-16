-- migrate:up
CREATE TABLE IF NOT EXISTS stocks (
    material_code UUID PRIMARY KEY,
    quantity int4 NOT NULL DEFAULT 0,
    state CHAR(1),
    updated_at timestamp(8) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (material_code) REFERENCES materials (code)
);
-- migrate:down
DROP TABLE IF EXISTS stocks;