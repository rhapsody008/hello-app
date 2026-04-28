package main
 
import (
	"fmt"
	"log"
	"net/http"
	"os"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "App is up and running!")
 
	clusterName := os.Getenv("CLUSTER_NAME")
	if clusterName == "" {
		clusterName = "unknown"
	}
	fmt.Fprintf(w, "Hello from %s!\n VERSION 2\n", clusterName)
}
 
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}
 
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)
	mux.HandleFunc("/healthz", healthz)
 
	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}
 
	log.Printf("App is up and running! Listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
