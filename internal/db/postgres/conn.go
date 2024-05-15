package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wildegor/kaspi-rest/internal/configs"
	"log/slog"
)

// PostgresConnection holds db conn
type PostgresConnection struct {
	DB  *pgxpool.Pool
	cfg *configs.PostgresConfig
}

func NewPostgresConnection(
	cfg *configs.PostgresConfig,
) *PostgresConnection {

	conn := &PostgresConnection{
		cfg: cfg,
	}

	conn.Connect()

	return conn
}

// Connect make connect and ping db
func (p *PostgresConnection) Connect() {
	config, err := pgxpool.ParseConfig(p.cfg.URI())
	if err != nil {
		slog.Error("fail parse config")
		panic(err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	dbpool, err := pgxpool.NewWithConfig(context.TODO(), config)
	if err != nil {
		slog.Error("fail connect to postgres")
		panic(err)
	}

	p.DB = dbpool

	if err := dbpool.Ping(context.TODO()); err != nil {
		p.DB.Close()
	}
}

// Close close connection
func (p *PostgresConnection) Close() {
	if p.DB != nil {
		p.DB.Close()
	}

	slog.Info("connection to postgres closed success")
}
