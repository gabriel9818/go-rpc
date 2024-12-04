package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Argumentos struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Respuesta struct {
	Resultado int `json:"resultado"`
}

func main() {
	// Definir los números que se van a sumar
	args := Argumentos{
		A: 10,
		B: 20,
	}

	// Convertir los datos en formato JSON
	data, err := json.Marshal(args)
	if err != nil {
		log.Fatalf("Error al convertir los datos a JSON: %v", err)
	}

	// Realizar la solicitud POST al servidor RPC
	resp, err := http.Post("http://localhost:8080/sumar", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error al conectar con el servidor: %v", err)
	}
	defer resp.Body.Close()

	// Decodificar la respuesta del servidor
	var respuesta Respuesta
	err = json.NewDecoder(resp.Body).Decode(&respuesta)
	if err != nil {
		log.Fatalf("Error al leer la respuesta del servidor: %v", err)
	}

	// Mostrar el resultado
	fmt.Printf("Resultado de la suma: %d\n", respuesta.Resultado)
}
