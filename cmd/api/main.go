package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type application struct{}

func main() {
	addr := flag.String("Addr", ":8080", "network port")
	flag.Parse()

	app := &application{}

	fmt.Println("Server starting. Listening on port", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
