package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		Host:     getEnv("BLK_DATABASE_HOST", "localhost"),
		Port:     getEnv("BLK_DATABASE_PORT", "5580"),
		User:     getEnv("BLK_DATABASE_USER", "postgres"),
		Password: getEnv("BLK_DATABASE_PASSWORD", "123456789"),
		DBName:   getEnv("BLK_DATABASE_NAME", "blockbuster"),
		SSLMode:  getEnv("BLK_DATABASE_SSL_MODE", "disable"),
	}
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(config *Config) (*Database, error) {
	pool, err := pgxpool.New(context.Background(), config.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		log.Println("Database connection closed")
	}
}
