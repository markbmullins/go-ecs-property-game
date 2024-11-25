// pkg/server/server.go

package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/markbmullins/city-developer/pkg/actions"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/rs/cors"
)

var mu sync.Mutex

func StartServer(world *ecs.World) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		actions.HandleAction(world, w, r)
	})

	mux.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(world) // Include game state
	})

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow your frontend's origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true, // If you need to allow cookies or authentication
	})

	handler := c.Handler(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Server started at http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
