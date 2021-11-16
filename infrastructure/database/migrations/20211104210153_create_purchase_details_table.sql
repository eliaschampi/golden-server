-- migrate:up
CREATE TABLE IF NOT EXISTS purchase_details (
    purchase_code UUID NOT NULL,
    material_code UUID NOT NULL,
    quantity int4 NOT NULL DEFAULT 0,
    price NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    FOREIGN KEY (purchase_code) REFERENCES purchases (expense_code),
    FOREIGN KEY (material_code) REFERENCES materials (code),
    PRIMARY KEY (purchase_code, material_code)
);
-- migrate:down
DROP TABLE IF EXISTS purchase_details;