package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"orders.go/m/internal/infrastructure/config"
	"time"
)

type ExternalDatabase struct {
	pools map[string]*sql.DB
}

func (e *ExternalDatabase) GetConnection(connection string) (*sql.DB, error) {
	pool, exists := e.pools[connection]
	if !exists {
		return nil, fmt.Errorf("connection %s does not exist", connection)
	}

	return pool, nil
}

func (e *ExternalDatabase) CloseAll() {
	for _, pool := range e.pools {
		err := pool.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}
}

func New(config config.Config) *ExternalDatabase {
	pgPools := &ExternalDatabase{
		pools: make(map[string]*sql.DB),
	}

	pgConnections := config.MapKeys("postgres.connections")

	var connectionsCount = len(pgConnections)

	if connectionsCount > 0 {

		log.Info().Msgf("Found %d connections", connectionsCount)

		for _, connectionName := range pgConnections {

			log.Info().Msgf("Trying %s connection", connectionName)

			pgUrl := config.GetString(fmt.Sprintf("postgres.connections.%s.url", connectionName))

			var counter int
			var sleep = time.Duration(1) * time.Second
			var connAttempts = config.GetInt("postgres.conn_attempts")

			for {

				counter++

				log.Info().
					Int("iteration", counter).
					Str("connection", connectionName).
					Str("state", "connecting").
					Str("url", pgUrl).
					Send()

				db, err := connect(pgUrl)

				if err != nil {

					log.Error().
						Err(err).
						Int("iteration", counter).
						Str("connection", connectionName).
						Str("state", "connecting").
						Str("url", pgUrl).
						Send()

				} else {

					log.Info().
						Int("iteration", counter).
						Str("connection", connectionName).
						Str("state", "connected").
						Str("url", pgUrl).
						Send()

					pgPools.pools[connectionName] = db
					break
				}

				if counter > connAttempts {

					if err != nil {
						panic(err)
					}

					break
				}

				time.Sleep(sleep) // Пауза в одну секунду
			}

		}
	}

	return pgPools
}

func connect(pgUrl string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	var pool *pgxpool.Pool

	// Get PG Connections Pool
	pool, err = pgxpool.New(context.Background(), pgUrl)
	if err == nil {
		db = stdlib.OpenDBFromPool(pool)
		err = db.Ping()
	}

	// Return connection
	return db, err
}
