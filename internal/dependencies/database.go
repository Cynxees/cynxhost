package dependencies

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

type DatabaseClient struct {
	Db *sql.DB
}

func NewDatabaseClient(config *Config) (*DatabaseClient, error) {

	dataSourceName := config.Database.MySQL.Username + ":" + config.Database.MySQL.Password + "@tcp(" + config.Database.MySQL.Host + ":" + strconv.Itoa(config.Database.MySQL.Port) + ")/" + config.Database.MySQL.Database + "?parseTime=true"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseClient{Db: db}, nil
}

func (client *DatabaseClient) Close() error {
	return client.Db.Close()
}

func (client *DatabaseClient) RunMigrations(migrationsPath string) {
	// Set the dialect to MySQL
	if err := goose.SetDialect("mysql"); err != nil {
    log.Fatalf("Failed to set dialect: %v", err)
  }

	// Run migrations
	if err := goose.Up(client.Db, migrationsPath); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}