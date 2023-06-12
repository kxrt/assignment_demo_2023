package main

import (
	"database/sql"
	"fmt"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitializeDatabase() (*sql.DB, error) {
	// Get the environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func PushMessage(m *rpc.Message) error {
	db, err := InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	chat := m.GetChat()
	sendTime := m.GetSendTime()
	sender := m.GetSender()
	text := m.GetText()

	query := `INSERT INTO messages (id, message, sender, created_at) VALUES ($1, $2, $3, $4)`

	// Execute the INSERT query.
	_, err = db.Exec(query, chat, text, sender, sendTime)
	if err != nil {
		return err
	}

	return nil

}
