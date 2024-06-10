package main

import (
	"chaatra/core/parser"
	"chaatra/core/trans"
	"chaatra/persistence"
	"log"
	"net/http"

	h "chaatra/http"

	"github.com/rs/cors"
)

var d parser.Dictionary

func main() {
	// initialize elastic search
	persistence.InitEs()

	trans.T = &trans.Trie{
		Root: &trans.Node{
			Letter: &trans.Letter{
				Devanagari: '*',
			},
			Children: make(map[rune]*trans.Node),
		},
	}

	if d = parser.Parse(trans.T); d != nil {
		parser.D = d
		persistence.IndexEntries(d)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/search", h.SearchHandler)
	mux.HandleFunc("/complete", h.AutoCompleteHandler)
	mux.HandleFunc("/dhatus", h.SearchDhatuHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allowing only http://localhost:3000
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Assuming you might need Authorization
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(mux)

	log.Println("Starting server on port : 8081")
	err := http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatalf("Shutting down server : %s", err.Error())
	} else {
		log.Println("Shutting down server")
	}
}
