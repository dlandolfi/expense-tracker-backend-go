package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	addr := flag.String("Addr", ":8080", "network port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)

	fmt.Println("Server starting. Listening on port", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
