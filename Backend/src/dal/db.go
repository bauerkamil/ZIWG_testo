package dal

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var loadError = godotenv.Load("../.env")
var dsn = "host=" + os.Getenv("DB_HOST") +
	" user=" + os.Getenv("DB_USER") +
	" password=" + os.Getenv("DB_PASSWORD") +
	" dbname=" + os.Getenv("DB_NAME") +
	" port=" + os.Getenv("DB_PORT")
var DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "system.", // table name prefix, table for `User` would be `t_users`
		SingularTable: true,      // use singular table name, table for `User` would be `user` with this option enabled
	}})

func init() {
	if loadError != nil {
		panic("Error loading .env file")
	}
	if err != nil {
		panic("failed to connect database")
	}
}
