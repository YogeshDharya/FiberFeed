package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

//So do we need to declare this as an independent go project ?  No right ? In that case should there be a go.mod file at the root ??
func ConnectPostgres() (*sql.DB, error){
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
dbHost,dbPort,dbUser,dbPassword,dbName)
db,err := sql.Open("postgres",dsn)
if err != nil {
	return nil, err
}
if err := db.Ping(); err != nil {
	return nil, err 
}
log.Println("Connected to PostgreSQL Database Successfully !")
return db,nil
}