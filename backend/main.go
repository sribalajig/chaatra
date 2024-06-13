package main

import (
	"chaatra/persistence"
	"chaatra/service"
	"log"
	"net/http"
	"os"

	h "chaatra/http"

	"github.com/rs/cors"
)

func main() {
	// initialize elastic search
	persistence.InitEs()

	var err error
	h.Dictionary, err = service.ParseApteDictionary(`dictionary.xml`)
	if err != nil {
		log.Println(`error parsing the dicrionary : `, err.Error())

		os.Exit(1)
	}

	h.Trie = service.BuildTrie(h.Dictionary)

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
	err = http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatalf("Shutting down server : %s", err.Error())
	} else {
		log.Println("Shutting down server")
	}
}
