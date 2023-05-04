package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
)

// DB Global db connection
var DB *sqlx.DB

// InitGlobalDatabaseConnection sets the global database connection
func InitGlobalDatabaseConnection() {
	fmt.Println("init scripts connection ....")
	db, err := sqlx.Open("mysql", viper.GetString("mysql.url"))

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("db stats:", db.Stats())

	DB = db
}
