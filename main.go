package main

import (
	"fmt"
	"net/http"
	"taskflow/handlers"
)

func main() {
	mux := http.NewServeMux()

	TaskHandler := handlers.NewTaskHandler()
	TaskHandler.RegisterRoutes(mux)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TaskFlow API is running")
	})

	fmt.Println("Server running on http://localhost8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("server error:", err)

	}

}
