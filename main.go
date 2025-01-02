package main

import (
	"fmt"
	"net/http"

	"log"
)

func main() {
	// Instancia o fileServer para que cada rota retorne o html estatico
	fileServer := http.FileServer(http.Dir("./static"))
	// Define a rota principal do servidor
	http.Handle("/", fileServer)
	// Define a rota para o formulario
	http.HandleFunc("/form", formHandler)
	// Define a rota para exibir o texto de Hello, world!
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// Inicia o servidor na porta 8080, e caso ocorra algum erro o logg o retorna
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Checa se a rota é /hello, caso contrario retorna 404
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// Checa se o metodo é GET, caso contrario retorna 404
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	// Retorna para o response o texto "Hello, world!"
	w.Write([]byte("Hello, world!"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	//Checa se a rota a /form contem algum conteudo com um parse no formulario, caso contrario retorna erro de parse no formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "POST request successful\n")

	// Atribui o valor do campo name do formulario html a variavel name
	name := r.FormValue("name")
	// Atribui o valor do campo address do formulario html a variavel address
	address := r.FormValue("address")

	// Retorna e imprime para o response o nome e o endereço informado no formulario
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Addres: %s\n", address)

}
