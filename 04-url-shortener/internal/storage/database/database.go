package database

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"url-shortener/internal/generator"
)

const sqlGetMaxID = "SELECT MAX(id) FROM short_url"
const sqlInsertURL = "INSERT INTO short_url (id, short_url, url) VALUES($1, $2, $3)"
const sqlGetURL = `SELECT url
FROM short_url
WHERE short_url = $1
`

var errorUnexistsShortURL = errors.New("url does not exist")

type Database struct {
	maxID    int
	dbClient IPGClient
	mx       sync.Mutex
}

type IPGClient interface {
	Exec(ctx context.Context, sql string, arguments ...any) error
	QueryRow(ctx context.Context, dest []any, sql string, args ...any) error
}

func New(ctx context.Context, dbClient IPGClient) (*Database, error) {
	db := &Database{
		dbClient: dbClient,
		mx:       sync.Mutex{},
	}
	err := db.initMaxID(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) initMaxID(ctx context.Context) error {
	err := db.dbClient.QueryRow(ctx, []any{&db.maxID}, sqlGetMaxID)
	if err != nil {
		return fmt.Errorf("failed to get maxID: %s", err)
	}
	db.maxID++
	return nil
}

func (db *Database) SaveURL(ctx context.Context, url string) (string, error) {
	db.mx.Lock()
	maxID := db.maxID
	db.maxID++
	db.mx.Unlock()

	shortURL := generator.GenerateURL(maxID)
	err := db.dbClient.Exec(
		ctx,
		sqlInsertURL,
		maxID,
		shortURL,
		url,
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert url '%s': %s", url, err)
	}

	return shortURL, nil
}

func (db *Database) GetURL(ctx context.Context, shortURL string) (string, error) {
	var url string
	err := db.dbClient.QueryRow(ctx, []any{&url}, sqlGetURL, shortURL)
	if url == "" || err != nil {
		return "", errorUnexistsShortURL
	}

	return url, nil
}
