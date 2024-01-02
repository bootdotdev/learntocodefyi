package main

import (
	"log"
	"net/http"

	"github.com/bootdotdev/learntocodefyi/internal/sendgridwrap"
	"github.com/bootdotdev/learntocodefyi/internal/sqlcdb"
)

func (hs *handlerState) handlerPostSubscribe(w http.ResponseWriter, r *http.Request, user sqlcdb.User) {
	err := hs.sendgridClient.UpsertContact(user.Email, []string{sendgridwrap.LearnToCodeFYIListID})
	if err != nil {
		log.Println("Error subscribing to list:", err)
		http.Error(w, "Error subscribing to list", http.StatusInternalServerError)
		return
	}

	hs.templates.ExecuteTemplate(w, "subscribed.html", nil)
}
