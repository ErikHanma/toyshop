package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Маршрутизация для catalog-service
	catalogServiceURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	catalogServiceProxy := httputil.NewSingleHostReverseProxy(catalogServiceURL)
	router.HandleFunc("/products", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodGet)
	router.HandleFunc("/products", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodPut)
	router.HandleFunc("/products/{id}", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodDelete)

	// Маршрутизация для user-service
	userServiceURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
	userServiceProxy := httputil.NewSingleHostReverseProxy(userServiceURL)
	router.HandleFunc("/users", proxyRequestHandler(userServiceProxy)).Methods(http.MethodGet)
	router.HandleFunc("/register", proxyRequestHandler(userServiceProxy)).Methods(http.MethodPost)
	router.HandleFunc("/login", proxyRequestHandler(userServiceProxy)).Methods(http.MethodGet)


	// Маршрутизация для order-service
	orderServiceURL, err := url.Parse("http://localhost:8083")
	if err != nil {
		log.Fatal(err)
	}
	orderServiceProxy := httputil.NewSingleHostReverseProxy(orderServiceURL)
	router.HandleFunc("/orders", proxyRequestHandler(orderServiceProxy)).Methods(http.MethodGet)
	// ... (добавь другие маршруты для создания, получения по ID, обновления и удаления заказов)

	// Запуск сервера API Gateway
	fmt.Println("API Gateway listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Функция-обработчик для проксирования запросов
func proxyRequestHandler(target *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Proxying request to: %s\n", r.URL)
		target.ServeHTTP(w, r)
	}
}

