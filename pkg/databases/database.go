package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/Samestora/WebOverflow/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB;
}
func Connect() (*Database, error) {
	config := configs.New();

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost,
		config.DBUser,
		config.DBPass,
		config.DBName,
		config.DBPort,
	)

	Db, err := sqlx.Connect("postgres", dsn);

	if err != nil {
		log.Fatal("Failed to connect to DB!\n", err);
		os.Exit(2)
	}

	return &Database{Db}, nil;
}
