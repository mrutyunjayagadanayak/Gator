// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: delete.sql

package database

import (
	"context"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
`

func (q *Queries) DeleteFeedFollow(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users
`

func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsers)
	return err
}

const deletefeeds = `-- name: Deletefeeds :exec
DELETE FROM feeds
`

func (q *Queries) Deletefeeds(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deletefeeds)
	return err
}
