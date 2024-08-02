package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	repr "github.com/v3nooom/st3llar/supplier/internal/api/representation"
	svcOAuth "github.com/v3nooom/st3llar/supplier/internal/service/oauth"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var body repr.Login
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("---> Request body:")
	fmt.Printf("%#v\n", body)

	// Extract Query Parameters
	q1 := r.URL.Query().Get("k1")
	q2 := r.URL.Query().Get("k2")

	fmt.Printf("---> Query param: %s\n", q1)
	fmt.Printf("---> Query param: %s\n", q2)

	svcOAuth.Logout(body)

	// TODO: response handler
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "login"}`))
}
