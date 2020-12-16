package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	// "os"
	_ "github.com/lib/pq"
)

func connector()(*sql.DB, error)  {
	connStr := "postgres://xaghszxlbdtefr:54b7385b48931476fe8dcd41c7f2d2d45ffaf7dc8532999a7ab4b336a659cfc6@ec2-23-23-242-234.compute-1.amazonaws.com:5432/dbp0gdlsb6fe41"
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error from connection")
		return nil, fmt.Errorf("Error from connection")
	}
	return db, nil
}
