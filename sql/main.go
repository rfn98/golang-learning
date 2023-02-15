package main

import (
	"fmt"
	// "database/sql"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
	"time"
	"log"
	"context"
)

var (
	DB *sql.DB
	once sync.Once
)

type employee struct {
	fullname 		string
	gender			string
	nik 			string
	phone_number	string
	address 		string
}

type religion struct {
	id 	  int
	name  string
}

func SetQueryContext() (context.Context, context.CancelFunc) {
	d := time.Now().Add(time.Millisecond * time.Duration(350000))
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, d)

	return ctx, cancel
}

func getReligions() {
	ctx, cancel := SetQueryContext()
	defer cancel()
    rows, err := DB.QueryxContext(ctx, "select * from religions")
    if err != nil {
        fmt.Println("ERROR SELECT", err.Error())
        return
    }
    defer rows.Close()

    var result []religion

    for rows.Next() {
        var each = religion{}
        var err = rows.Scan(&each.id, &each.name)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
        fmt.Println(each.name)
    }
}

func init() {
	once.Do(func() {
		var err error
		DB, err = sql.Open("postgres", "host=localhost port=5432 dbname=db_hris sslmode=disable")

		if err != nil {
			log.Println("Error creating connection pool: " + err.Error())
		}

		DB.SetMaxOpenConns(10) // Sane default
		DB.SetMaxIdleConns(5)

		DB.SetConnMaxLifetime(time.Minute * 10)

		//log.Printf("Connected!\n")

		if err = DB.Ping(); err != nil {
			log.Panic(err)
		}

		fmt.Println("CONNECT DB SUCCESS")
	})
}

func main() {
	getReligions()
}