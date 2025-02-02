package pg

import (
	"context"
	"errors"
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

const connectInfo = "host=%s port=%s user=%s password=%s dbname=%s"

var errorConnect = errors.New("failed to connect: %s")
var errorPing = errors.New("failed to ping: %s")

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
		return Client{}, fmt.Errorf(errorConnect.Error(), err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return Client{}, fmt.Errorf(errorPing.Error(), err)
	}

	return Client{conn: conn}, nil
}
