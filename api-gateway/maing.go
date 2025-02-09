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

	// Define service URLs
	catalogServiceURL := mustParseURL("http://localhost:8081")
	userServiceURL := mustParseURL("http://localhost:8082")
	orderServiceURL := mustParseURL("http://localhost:8083")

	// Set up reverse proxies
	catalogServiceProxy := createReverseProxy(catalogServiceURL)
	userServiceProxy := createReverseProxy(userServiceURL)
	orderServiceProxy := createReverseProxy(orderServiceURL)

	// Route definitions
	router.HandleFunc("/products", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/products/{id}", proxyRequestHandler(catalogServiceProxy)).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	router.HandleFunc("/users", proxyRequestHandler(userServiceProxy)).Methods(http.MethodGet)
	router.HandleFunc("/register", proxyRequestHandler(userServiceProxy)).Methods(http.MethodPost)
	router.HandleFunc("/login", proxyRequestHandler(userServiceProxy)).Methods(http.MethodGet)

	router.HandleFunc("/orders", proxyRequestHandler(orderServiceProxy)).Methods(http.MethodGet)

	// Start the API Gateway server
	fmt.Println("API Gateway listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Helper function to parse URLs and handle errors
func mustParseURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}
	return parsedURL
}

// Helper function to create a reverse proxy
func createReverseProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

// Request handler for proxying requests
func proxyRequestHandler(target *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Proxying request to: %s\n", r.URL)
		target.ServeHTTP(w, r)
	}
}
