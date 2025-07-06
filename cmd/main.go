package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	emailverifier "github.com/addix999/email-verifier"
)

func main() {
	v := emailverifier.NewVerifier()

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, `{"error":"email parameter is required"}`, http.StatusBadRequest)
			return
		}

		ret, err := v.Verify(email)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(ret)
	})

	port := "10000"
	fmt.Println("✅ Server berjalan di port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
