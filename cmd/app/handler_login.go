package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/bootdotdev/learntocodefyi/internal/auth"
	"github.com/bootdotdev/learntocodefyi/internal/sqlcdb"
	"github.com/google/uuid"
)

func (hs *handlerState) handlerLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	formEmail := r.FormValue("email")
	if formEmail == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := hs.queries.InsertUser(
		r.Context(),
		sqlcdb.InsertUserParams{
			ID:    uuid.New().String(),
			Email: strings.ToLower(formEmail),
		},
	)
	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		user, err = hs.queries.GetUserByEmail(r.Context(), formEmail)
		if err != nil {
			log.Println("Error getting user by email:", err)
			http.Error(w, "error getting user by email", http.StatusInternalServerError)
			return
		}
	}
	if err != nil {
		log.Println("Error upserting user:", err)
		http.Error(w, "error upserting user", http.StatusInternalServerError)
		return
	}

	jwt, err := auth.MakeJWT(
		user.ID,
		hs.jwtSecret,
	)
	if err != nil {
		log.Println("Error making JWT:", err)
		http.Error(w, "error making JWT", http.StatusInternalServerError)
		return
	}

	err = hs.sendgridClient.SendMagicLink(
		formEmail,
		"Survey Taker",
		"https://learntocode.fyi/survey?jwt="+jwt,
	)
	if err != nil {
		log.Println("Error sending magic link:", err)
		http.Error(w, "error sending magic link", http.StatusInternalServerError)
		return
	}
	hs.templates.ExecuteTemplate(w, "magiclinkconfirm.html", nil)
}
