package db

import (
	"context"
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
)

const databaseDriver = "postgres"

// instance holds a single instance of the database
var instance *ReaderWriter

var ready atomic.Bool

// once is used to ensure that there is only a single instance of the database
var once sync.Once

// Store holds the database connection and other stuff.
type Store struct {
	db    func() *sql.DB
	dbs   *ReaderWriter
	ready *atomic.Bool
}

// NewDbConnectionFromSettings sets up a db connection from the settings, only once
func NewDbConnectionFromSettings(ctx context.Context, settings *Settings, withSearchPath bool) Store {
	once.Do(func() {
		instance = NewDbConnection(
			ctx,
			&ready,
			ConnectOptions{
				Retries:            5,
				RetryDelay:         time.Second * 10,
				ConnectTimeout:     time.Minute * 5,
				DSN:                settings.BuildConnectionString(withSearchPath),
				MaxOpenConnections: settings.MaxOpenConnections,
				MaxIdleConnections: settings.MaxIdleConnections,
				ConnMaxLifetime:    time.Minute * 5,
				DriverName:         databaseDriver,
			},
			ConnectOptions{
				Retries:            5,
				RetryDelay:         time.Second * 10,
				ConnectTimeout:     time.Minute * 5,
				DSN:                settings.BuildConnectionString(true),
				MaxOpenConnections: settings.MaxOpenConnections,
				MaxIdleConnections: settings.MaxIdleConnections,
				ConnMaxLifetime:    time.Minute * 5,
				DriverName:         databaseDriver,
			},
		)
	})

	return Store{db: instance.GetWriterConn, dbs: instance, ready: &ready}
}

// IsReady returns if db is ready to connect to
func (store *Store) IsReady() bool {
	return store.ready.Load()
}

// DBS returns the reader and writer databases to connect to
func (store *Store) DBS() *ReaderWriter {
	return store.dbs
}

// NewDbConnectionForTest use this for tests as we have multiple sessions in parallel and don't want the synced one
func NewDbConnectionForTest(ctx context.Context, settings *Settings, withSearchPath bool) Store {
	localReady := atomic.Bool{}
	dbConnection := NewDbConnection(
		ctx,
		&localReady,
		ConnectOptions{
			Retries:            5,
			RetryDelay:         time.Second * 10,
			ConnectTimeout:     time.Minute * 5,
			DSN:                settings.BuildConnectionString(withSearchPath),
			MaxOpenConnections: settings.MaxOpenConnections,
			MaxIdleConnections: settings.MaxIdleConnections,
			ConnMaxLifetime:    time.Minute * 5,
			DriverName:         databaseDriver,
		},
		ConnectOptions{
			Retries:            5,
			RetryDelay:         time.Second * 10,
			ConnectTimeout:     time.Minute * 5,
			DSN:                settings.BuildConnectionString(true),
			MaxOpenConnections: settings.MaxOpenConnections,
			MaxIdleConnections: settings.MaxIdleConnections,
			ConnMaxLifetime:    time.Minute * 5,
			DriverName:         databaseDriver,
		},
	)
	return Store{db: dbConnection.GetWriterConn, dbs: dbConnection, ready: &localReady}
}

// WaitForDB waits 30 seconds for db to become available, and Fatal panic if can't connect.
func (store *Store) WaitForDB(logger zerolog.Logger) {
	totalTime := 0
	for !store.IsReady() {
		if totalTime > 30 {
			logger.Fatal().Msg("could not connect to postgres after 30 seconds")
		}
		time.Sleep(time.Second)
		totalTime++
	}
}
