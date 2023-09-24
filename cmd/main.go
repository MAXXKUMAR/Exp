package main

import (
	"fmt"
	"net/http"

	"github.com/MAXXKUMAR/Exp/handlers"
)

func main() {
	http.HandleFunc("/modes/{areaCode}", handlers.GetModeCountsByArea)
	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
