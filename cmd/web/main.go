package main

import (
	"flag"
	"github.com/jackc/pgx"
	"log"
	"module1/pkg/postgre"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgre.SnippetModel
}


func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	config := pgx.ConnConfig{
		Host: "localhost",
		Port: 5432,
		Database: "GoDB",
		User: "postgres",
		Password: "beks300900",
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 10,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	conn, err := pgx.NewConnPool(poolConfig)

	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgre.SnippetModel{Conn: conn},
	}

	server := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
