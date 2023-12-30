package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zakhaev26/mail_service/function"
)

func main() {
	r := mux.NewRouter()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env")
		return
	}

	r.HandleFunc("/send-mail", MailRequestHandler).Methods("POST")
	fmt.Println("Server live at :6969")
	http.ListenAndServe(":6969", r)
}

type MailingParams struct {
	Subject     string   `json:"subject,omitempty"`
	Content     string   `json:"content,omitempty"`
	To          []string `json:"to,omitempty"`
	Cc          []string `json:"cc,omitempty"`
	Bcc         []string `json:"bcc,omitempty"`
	AttachFiles []string `json:"attachFiles,omitempty"`
}

func MailRequestHandler(w http.ResponseWriter, r *http.Request) {
	var paramInstance MailingParams

	senderName := os.Getenv("EMAIL_SENDER_NAME")
	senderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	senderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")

	sender := function.NewGmailSender(senderName, senderAddress, senderPassword)

	err := json.NewDecoder(r.Body).Decode(&paramInstance)

	if err != nil {
		jsonPayload := `{"err": "` + err.Error() + `"}`
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonPayload)
		return
	}

	fmt.Println(paramInstance)

	err = sender.SendEmail(paramInstance.Subject, paramInstance.Content, paramInstance.To, nil, nil, nil)

	if err != nil {
		jsonPayload := `{"err": "` + err.Error() + `"}`
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonPayload)
		return
	}

	fmt.Println("Sucess")
	w.Write([]byte("E-mail sent Successfully."))
}
