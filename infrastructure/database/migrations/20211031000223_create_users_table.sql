-- migrate:up
CREATE TABLE IF NOT EXISTS users (
    dni CHAR(8) PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    rol_code UUID NOT NULL,
    gender CHAR(1),
    image VARCHAR(100),
    address VARCHAR(100),
    phone VARCHAR(40),
    email VARCHAR(50),
    password VARCHAR(100),
    created_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    FOREIGN KEY (rol_code) REFERENCES roles (code)
);
-- migrate:down
DROP TABLE IF EXISTS users;