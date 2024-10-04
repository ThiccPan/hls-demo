package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main2() {
	router := http.NewServeMux()
	router.HandleFunc("GET /files", func(w http.ResponseWriter, r *http.Request) {
		res := map[string]any{
			"data": "hi there",
		}
		resJson, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
