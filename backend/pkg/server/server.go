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

type PartialWorld struct {
	Entities                 map[string]*ecs.Entity `json:"entities"`
	OwnedPropertiesIndex     map[int][]int          `json:"owned_properties_index"`     // ownerID -> propertyIDs
	GroupPropertiesIndex     map[int][]int          `json:"group_properties_index"`     // groupID -> propertyIDs
	GroupUpgradedPercentages map[int]float64        `json:"group_upgraded_percentages"` // groupID -> upgradedPercentage
	GroupUpgradedCounts      map[int]int            `json:"group_upgraded_counts"`      // groupID -> number of properties with >=1 upgrade
	Players                  []*ecs.Entity          `json:"players"`
}

func sendPartialWorld(w http.ResponseWriter, world *ecs.World) {
	partial := PartialWorld{
		Entities:                 world.Entities,
		OwnedPropertiesIndex:     world.OwnedPropertiesIndex,
		GroupPropertiesIndex:     world.GroupPropertiesIndex,
		GroupUpgradedPercentages: world.GroupUpgradedPercentages,
		GroupUpgradedCounts:      world.GroupUpgradedCounts,
		Players:                  world.Players,
	}
	json.NewEncoder(w).Encode(partial)
}

func StartServer(world *ecs.World) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		actions.HandleAction(world, w, r)
	})

	mux.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received GET request for /state")
		mu.Lock()
		defer mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		sendPartialWorld(w, world)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		log.Println("Server started at http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
