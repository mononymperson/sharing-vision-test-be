package main

import (
	"log"
	"net/http"
	"os"

	"sharing-vision-be/internal/article"
	"sharing-vision-be/internal/common"

	"github.com/joho/godotenv"
)

type app struct {
	articleHandler *article.ArticleHandler
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()

	db, err := common.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	defer db.Close()

	articleRepo := article.NewArticleRepository(db)
	articleService := article.NewArticleService(articleRepo)
	articleHandler := article.NewArticleHandler(articleService)

	app := &app{articleHandler: articleHandler}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /article/", app.articleHandler.Create)
	mux.HandleFunc("GET /article/{limit}/{offset}", app.articleHandler.FindAll)
	mux.HandleFunc("GET /article/{id}", app.articleHandler.FindByID)
	mux.HandleFunc("PUT /article/{id}", app.articleHandler.Update)
	mux.HandleFunc("DELETE /article/{id}", app.articleHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, cors(mux)))
}
