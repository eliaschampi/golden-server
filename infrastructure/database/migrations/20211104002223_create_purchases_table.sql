-- migrate:up
CREATE TABLE IF NOT EXISTS purchases (
    expense_code UUID PRIMARY KEY,
    contact_code UUID NOT NULL,
    voucher_num VARCHAR(20),
    FOREIGN KEY (expense_code) REFERENCES expenses (code),
    FOREIGN KEY (contact_code) REFERENCES contacts (code)
);
-- migrate:down
DROP TABLE IF EXISTS purchases;