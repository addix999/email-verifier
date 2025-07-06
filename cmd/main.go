package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	emailverifier "github.com/addix999/email-verifier"
)

func main() {
	// Inisialisasi email verifier
	v := emailverifier.NewVerifier()

	// Endpoint /verify?email=...
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		// Set header agar response selalu JSON
		w.Header().Set("Content-Type", "application/json")

		// Ambil parameter email
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, `{"error":"email parameter is required"}`, http.StatusBadRequest)
			return
		}

		// Verifikasi email
		ret, err := v.Verify(email)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		// Kirim hasil JSON
		json.NewEncoder(w).Encode(ret)
	})

	// Ambil PORT dari environment (untuk Render), fallback ke 10000 jika kosong
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	fmt.Println("✅ Server berjalan di port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
