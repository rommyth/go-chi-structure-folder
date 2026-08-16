package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Database struct {
	db *pgxpool.Pool
}

func NewDatabase(viper *viper.Viper) (*pgxpool.Pool, error) {
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	sslmode := viper.GetString("database.sslmode")
	database := viper.GetString("database.name")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt32("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	// dsn
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode,
	)

	// Parse Config
	parseConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database: %w \n", err)
	}

	// Setup delay
	parseConfig.MaxConns = maxConnection
	parseConfig.MaxConnIdleTime = time.Duration(idleConnection)
	parseConfig.MaxConnLifetime = time.Second * time.Duration(maxLifeTimeConnection)

	pool, err := pgxpool.NewWithConfig(context.Background(), parseConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w \n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to pool: %w \n", err)
	}

	return pool, nil
}
