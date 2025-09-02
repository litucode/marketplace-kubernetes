package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "Marketplace Kubernetes Client: GET recibido")
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
