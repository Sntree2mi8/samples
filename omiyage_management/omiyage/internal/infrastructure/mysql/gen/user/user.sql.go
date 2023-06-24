// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package user

import (
	"context"
	"database/sql"
)

const findUser = `-- name: FindUser :one
SELECT id, name FROM user WHERE id = ?
`

func (q *Queries) FindUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const insertUser = `-- name: InsertUser :execresult
INSERT INTO user (id, name) VALUES (?, ?)
`

type InsertUserParams struct {
	ID   string
	Name string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUser, arg.ID, arg.Name)
}