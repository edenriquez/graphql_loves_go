package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/edenriquez/graphql_loves_go/backend"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const defaultPort = "3001"

func main() {
	LoadEnvVars()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// LoadEnvVars should load .env file
func LoadEnvVars() {
	base, _ := filepath.Abs(".")
	envPath := path.Join(base, "development.env")
	godotenv.Load(envPath)
}
