-- migrate:up
CREATE TABLE IF NOT EXISTS stocks
(
    material_code uuid PRIMARY KEY,
    quantity      int4 NOT NULL DEFAULT 0,
    state         BPCHAR(1),
    updated_at    timestamp(8)  DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (material_code)
        REFERENCES materials (code)
);

-- migrate:down
DROP TABLE IF EXISTS stocks;
