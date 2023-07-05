package main

import (
	"fmt"
	"log"
	"net/http"

	"graphql-sample/gql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	fmt.Println("graphql demo.")
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gql.QueryType,
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		//GraphiQL: true,
		Playground: true,
	})
	http.Handle("/graphql", h)

	err = http.ListenAndServe(":5011", nil)
	if err != nil {
		fmt.Printf("???? %s\n", err.Error())
	}
}
