package pg

import (
	"context"
	"fmt"
	"url-shortener/internal/env"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	conn *pgx.Conn
}

func (c *Client) Close(ctx context.Context) {
	c.conn.Close(ctx)
}

func (c *Client) Exec(ctx context.Context, sql string, arguments ...any) error {
	_, err := c.conn.Exec(ctx, sql, arguments...)
	return err
}

func (c *Client) QueryRow(ctx context.Context, dest []any, sql string, args ...any) error {
	return c.conn.QueryRow(ctx, sql, args...).Scan(dest...)
}

const connectInfo = "host=%s port=%s user=%s password=%s dbname=%s"

var errorConnect = "failed to connect: %s"
var errorPing = "failed to ping: %s"

func New(ctx context.Context) (Client, error) {
	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			connectInfo,
			env.Get("POSTGRES_HOST", "localhost"),
			env.Get("POSTGRES_PORT", "5432"),
			env.Get("POSTGRES_USER"),
			env.Get("POSTGRES_PASSWORD"),
			env.Get("POSTGRES_DB"),
		),
	)
	if err != nil {
		return Client{}, fmt.Errorf(errorConnect, err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return Client{}, fmt.Errorf(errorPing, err)
	}

	return Client{conn: conn}, nil
}
