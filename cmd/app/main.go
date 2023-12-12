package main

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/libsql/libsql-client-go/libsql"
)

func main() {
	godotenv.Load()

	// https://github.com/libsql/libsql-client-go/#open-a-connection-to-sqld
	// libsql://[your-database].turso.io?authToken=[your-auth-token]
	sqlURL := os.Getenv("SQLURL")
	if sqlURL == "" {
		log.Fatal("No SQLURL set in .env")
	}
	db, err := sql.Open("libsql", sqlURL)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	err = ensureDB(db)
	if err != nil {
		log.Fatal("error ensuring database: ", err)
	}
	//queries := sqlcdb.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	templates := template.Must(template.New("").ParseGlob("pages/*.html"))
	templates = template.Must(templates.ParseGlob("partials/*.html"))

	fileServer := http.FileServer(http.Dir("./public"))

	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	r.HandleFunc("/devreload", handlerDevReload)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})
	r.Get("/survey", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "survey.html", nil)
	})
	r.Get("/results", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "results.html", nil)
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "about.html", nil)
	})

	r.Get("/partials/questions/{num}", func(w http.ResponseWriter, r *http.Request) {
		numStr := chi.URLParam(r, "num")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		questions := getQuestions()
		if num < 0 || num >= len(questions) {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}

		switch questions[num].QuestionType {
		case SingleSelectQuestionType:
			templates.ExecuteTemplate(
				w,
				"singleselect.html",
				questions[num].SingleSelectQuestion,
			)
			return
		case MultiSelectQuestionType:
			templates.ExecuteTemplate(
				w,
				"multiselect.html",
				questions[num].MultiSelectQuestion,
			)
			return
		}
		http.Error(w, "Question type not found", http.StatusNotFound)
	})

	fmt.Println("Server running on: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

//go:embed sqlc/schema/*.sql
var embedMigrations embed.FS

func ensureDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		return fmt.Errorf("error setting goose dialect: %w", err)
	}
	if err := goose.Up(db, "sqlc/schema"); err != nil {
		return fmt.Errorf("error running goose migrations: %w", err)
	}
	fmt.Println("Database tables ensured!")
	return nil
}
