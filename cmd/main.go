package main

import (
	"database/sql"
	"log"

	"github.com/dunky-star/ecomm-proj/cmd/api"
	"github.com/dunky-star/ecomm-proj/configs"
	"github.com/dunky-star/ecomm-proj/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:   configs.Envs.DBUser,
		Passwd: configs.Envs.DBPassword,
		Addr:   configs.Envs.DBAddress,
		DBName: configs.Envs.DBName,
		Net:    "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil{
		log.Fatal(err)
	}

	initStorage(db)
	
	server := api.NewAPIServer(":9090", db)
    if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("DB: Successfully connected!")
}
