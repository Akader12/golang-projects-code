package database

import (
	"context"
	"fmt"
	"log"
	

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectB(){
	connStr := "postgres://postgres:biite@localhost:5432/myapp"
	var err error
	DB,err = pgx.Connect(context.Background(),connStr)
	if err != nil {
		log.Fatal("Connection failed")
	}
	fmt.Println("Connected to postgres")

	
}