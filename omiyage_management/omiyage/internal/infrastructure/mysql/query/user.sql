-- name: FindUser :one
SELECT * FROM user WHERE id = ?;

-- name: InsertUser :execresult
INSERT INTO user (id, name) VALUES (?, ?);
