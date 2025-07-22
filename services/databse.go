package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"searchsong/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo conexión a BD:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("No se puede conectar a la base de datos:", err)
	}

	log.Println("Conectado a la base de datos")
}

func GetDB() *sql.DB {
	return db
}

func SaveSong(song models.Song) error {
	query := `INSERT INTO songs (name, artist, duration, album, artwork, price, origin) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err := db.QueryRow(query,
		song.Name,
		song.Artist,
		song.Duration,
		song.Album,
		song.Artwork,
		song.Price,
		song.Origin,
	).Scan(&id)

	if err != nil {
		return fmt.Errorf("error guardando canción: %v", err)
	}

	fmt.Printf("Canción guardada con ID: %d\n", id)
	return nil
}

func GetAllSongs() ([]models.Song, error) {
	rows, err := db.Query("SELECT name, artist, duration, album, artwork, price, origin FROM songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var s models.Song
		if err := rows.Scan(&s.Name, &s.Artist, &s.Duration, &s.Album, &s.Artwork, &s.Price, &s.Origin); err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	return songs, nil
}
