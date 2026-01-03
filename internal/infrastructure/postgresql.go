package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"chat-app/internal/domain"
)

func InitDB(cfg domain.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DB)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error open DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("error ping DB: %v", err)
	}

	fmt.Println("Database initialized successfully!")
	return sqlDB
}
