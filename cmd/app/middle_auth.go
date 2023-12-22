package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bootdotdev/learntocodefyi/internal/auth"
	"github.com/bootdotdev/learntocodefyi/internal/sqlcdb"
)

// middleAuthHandler is a handler function that requires a user to be logged in
type middleAuthHandler func(http.ResponseWriter, *http.Request, sqlcdb.User)

func (hs *handlerState) middleAuth(handler middleAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwt := r.URL.Query().Get("jwt")
		cookie, err := r.Cookie("jwt")
		if err != nil && jwt == "" {
			if err == http.ErrNoCookie {
				http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
				return
			}
		}

		jwtToUse := ""
		if cookie != nil {
			jwtToUse = cookie.Value
		}
		if jwt != "" {
			jwtToUse = jwt
		}

		claims, err := auth.ValidateJWT(jwtToUse, hs.jwtSecret)
		if err != nil {
			log.Println("Error validating JWT:", err)
			http.Error(w, "Error validating JWT", http.StatusUnauthorized)
			return
		}

		user, err := hs.queries.GetUser(r.Context(), claims.UserID)
		if err != nil {
			log.Println("Error getting user:", err)
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}

		expiration := time.Now().Add(24 * time.Hour)
		newCookie := http.Cookie{
			Name:     "jwt",
			Value:    jwtToUse,
			Expires:  expiration,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &newCookie)

		handler(w, r, user)
	}
}
