package main

import (
	"fmt"
	"log"
	"main/src/common/aws"
	"main/src/common/db"
	"main/src/graph/generated"
	"main/src/graph/resolver"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	if err := aws.InitAws(); err != nil {
		fmt.Println(err)
		return
	}
	if err := db.InitMongo(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("tag test3")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
