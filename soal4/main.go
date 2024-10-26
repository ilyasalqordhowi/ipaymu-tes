package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxRequests = 100
	timeFrame   = 1 * time.Hour
)

type RateLimiter struct {
	db *sql.DB
}

func NewRateLimiter(connStr string) (*RateLimiter, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &RateLimiter{db: db}, nil
}

func (rl *RateLimiter) AllowRequest(userID string) (bool, error) {
	now := time.Now()

	
	_, err := rl.db.Exec("DELETE FROM requests WHERE user_id = $1 AND timestamp < $2", userID, now.Add(-timeFrame))
	if err != nil {
		return false, err
	}

	var count int
	err = rl.db.QueryRow("SELECT COUNT(*) FROM requests WHERE user_id = $1 AND timestamp >= $2", userID, now.Add(-timeFrame)).Scan(&count)
	if err != nil {
		return false, err
	}

	if count < maxRequests {
	
		_, err = rl.db.Exec("INSERT INTO requests (user_id, timestamp) VALUES ($1, $2)", userID, now)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func main() {
	connStr := "postgresql://postgres:1@localhost:5432/soal4?sslmode=disable"
	rateLimiter, err := NewRateLimiter(connStr)
	if err != nil {
		log.Fatal(err)
	}

	userID := "user123"
	allowed, err := rateLimiter.AllowRequest(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if allowed {
		fmt.Println("Request allowed")
	} else {
		fmt.Println("Rate limit exceeded")
	}
}


// run in terminal 
//docker run --name soal4 -e POSTGRES_PASSWORD=1 -p 5432:5432 -d postgres
//make migrate:reset
//go run main.go