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

//go:embed pages/*.html
var pagesFS embed.FS

//go:embed partials/*.html
var partialsFS embed.FS

//go:embed public/*
var publicFS embed.FS

//go:embed sqlc/schema/*.sql
var embedMigrations embed.FS

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
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	if sendgridAPIKey == "" {
		log.Fatal("No SENDGRID_API_KEY set in .env")
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

	templates, err := template.New("").ParseFS(pagesFS, "pages/*.html")
	if err != nil {
		log.Fatal("error parsing templates: ", err)
	}
	templates, err = templates.ParseFS(partialsFS, "partials/*.html")
	if err != nil {
		log.Fatal("error parsing templates: ", err)
	}

	sendgridClient := sendgridwrap.NewClient(sendgridAPIKey, os.Getenv("PLATFORM"))
	hs := handlerState{
		queries:        queries,
		templates:      templates,
		sendgridClient: sendgridClient,
		jwtSecret:      os.Getenv("JWT_SECRET"),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.FS(publicFS))
	r.Handle("/public/*", fileServer)
	r.HandleFunc("/devreload", handlerDevReload)

	r.Post("/login", hs.handlerLogin)
	r.Get("/survey", hs.middleAuth(hs.handlerGetSurvey))
	r.Post("/survey/{question_id}", hs.middleAuth(hs.handlerPostSurvey))
	r.Post("/subscribe", hs.middleAuth(hs.handlerPostSubscribe))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})
	r.Get("/results", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "results.html", nil)
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "about.html", nil)
	})
	r.Get("/thanks", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "thanks.html", nil)
	})

	fmt.Println("Server running on: http://localhost:8080")

	addr := ":8080"
	if os.Getenv("PLATFORM") == "PROD" {
		addr = "0.0.0.0:8080"
	}

	server := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

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
