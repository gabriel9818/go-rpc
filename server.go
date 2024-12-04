//go:build server
// +build server

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Argumentos struct {
	A int `json:"A"`
	B int `json:"B"`
}

type Respuesta struct {
	Resultado int `json:"resultado"`
}

func sumarHandler(w http.ResponseWriter, r *http.Request) {

	// Agregar encabezados CORS para permitir solicitudes desde cualquier origen
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir todos los orígenes
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		// Si es una solicitud preflight (OPTIONS), solo respondemos con los encabezados
		w.WriteHeader(http.StatusOK)
		return
	}
	// Asegúrate de que la solicitud es POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, usa POST", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica el cuerpo de la solicitud JSON
	var args Argumentos
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		http.Error(w, "Error al leer los argumentos: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Depuración: Imprime los valores recibidos
	fmt.Printf("Recibido: A = %d, B = %d\n", args.A, args.B)

	// Realiza la suma
	resultado := Respuesta{
		Resultado: args.A + args.B,
	}

	// Devuelve el resultado como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}

func main() {
	http.HandleFunc("/sumar", sumarHandler)

	fmt.Println("Servidor JSON-RPC escuchando en http://localhost:8080/sumar...")
	http.ListenAndServe(":8080", nil)
}
