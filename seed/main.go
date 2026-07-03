package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"sharing-vision-be/internal/common"

	"github.com/joho/godotenv"
)

var categories = []string{"Technology", "Lifestyle", "Programming", "Design", "Business", "Science", "Health", "Travel"}
var statuses = []string{"publish", "draft", "thrash"}

func randomTitle(n int) string {
	words := []string{
		"Getting Started", "Advanced Guide", "Deep Dive", "Complete Tutorial",
		"Beginner Friendly", "Expert Tips", "Best Practices", "Modern Approach",
		"The Ultimate", "A Comprehensive", "Practical", "Essential",
	}
	return fmt.Sprintf("%s Guide to Building Scalable Applications Part %d", words[rand.Intn(len(words))], n)
}

func randomContent() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
}

func main() {
	godotenv.Load()

	db, err := common.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	defer db.Close()

	for i := 1; i <= 100; i++ {
		now := time.Now().Add(-time.Duration(100-i) * time.Hour)
		_, err := db.Exec(
			`INSERT INTO posts (title, content, category, status, created_date, updated_date) VALUES (?, ?, ?, ?, ?, ?)`,
			randomTitle(i),
			randomContent(),
			categories[rand.Intn(len(categories))],
			statuses[rand.Intn(len(statuses))],
			now,
			now,
		)
		if err != nil {
			log.Fatal("failed to insert seed data: ", err)
		}
	}

	fmt.Println("Successfully inserted 100 articles!")
}
