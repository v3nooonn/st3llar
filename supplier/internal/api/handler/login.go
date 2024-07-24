package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	repr "github.com/v3nooom/st3llar/supplier/internal/api/representation"
	oauthSvc "github.com/v3nooom/st3llar/supplier/internal/service/oauth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Extract JSON Body
	id := r.Context().Value("id").(string)
	fmt.Println("id")
	fmt.Println(id)

	var body repr.Login
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("body")
	fmt.Printf("%#v\n", body)

	// Extract Query Parameters
	q1 := r.URL.Query().Get("k1")
	q2 := r.URL.Query().Get("k2")

	fmt.Println("q1")
	fmt.Println(q1)
	fmt.Println("q2")
	fmt.Println(q2)

	oauthSvc.Login(body)

	// TODO: response handler
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "login"}`))
}
