package main

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bootdotdev/learntocodefyi/internal/sendgridwrap"
	"github.com/bootdotdev/learntocodefyi/internal/sqlcdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/libsql/libsql-client-go/libsql"
)

type handlerState struct {
	queries        *sqlcdb.Queries
	templates      *template.Template
	sendgridClient sendgridwrap.Client
	jwtSecret      string
}

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
	queries := sqlcdb.New(db)
	templates := template.Must(template.New("").ParseGlob("pages/*.html"))
	templates = template.Must(templates.ParseGlob("partials/*.html"))
	sendgridClient := sendgridwrap.NewClient("", os.Getenv("PLATFORM"))
	hs := handlerState{
		queries:        queries,
		templates:      templates,
		sendgridClient: sendgridClient,
		jwtSecret:      os.Getenv("JWT_SECRET"),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))
	r.HandleFunc("/devreload", handlerDevReload)

	r.Post("/login", hs.handlerLogin)
	r.Get("/survey", hs.middleAuth(hs.handlerGetSurvey))
	r.Post("/survey/{question_id}", hs.middleAuth(hs.handlerPostSurvey))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})
	r.Get("/results", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "results.html", nil)
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "about.html", nil)
	})

	fmt.Println("Server running on: http://localhost:8080")

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
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
