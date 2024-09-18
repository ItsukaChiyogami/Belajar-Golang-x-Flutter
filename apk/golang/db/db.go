package db

import (
	"database/sql"
	"fmt"
	"golang/config" // Pastikan path sesuai dengan project Anda
	"log"

	_ "github.com/go-sql-driver/mysql" // Pastikan driver MySQL diimpor
)

var db *sql.DB
var err error

func Init() {
	// Mengambil konfigurasi dari file config
	conf := config.GetConfig()

	// Membuat connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)

	// Tambahkan log untuk memeriksa connection string
	fmt.Println("Connection String:", connectionString)

	// Membuka koneksi ke database
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}

	// Melakukan ping untuk memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping Error: %v", err)
	}

	fmt.Println("Database connection established successfully.")
}

// Fungsi untuk mengembalikan koneksi database
func CreateCon() *sql.DB {
	return db
}
