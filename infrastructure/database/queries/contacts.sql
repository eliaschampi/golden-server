-- name: GetContacts :many
SELECT * from contacts;

-- name: CreateContact :one
INSERT INTO contacts (name, type, email, phone, address, description)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;