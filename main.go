// Academic project to implement a consistent traffic redirection server using the consistent hashing implementation.
// Sets up a REST web server to accept server registration requests (and general service discovery mechanics) as well as redirect traffic to appropriate servers.

package main

import (
	"log"
	"net/http"

	"trafficDistributor/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.Root)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
