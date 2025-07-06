package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	emailverifier "github.com/addix999/email-verifier"
)

func main() {
	v := emailverifier.NewVerifier()

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "email query param is required", http.StatusBadRequest)
			return
		}

		ret, err := v.Verify(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ret)
	})

	fmt.Println("Server jalan di port 10000...")
	http.ListenAndServe(":10000", nil)
}
