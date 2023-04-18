package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/poneding/learning-go/practice/graphql/gql"
)

func main() {
	fmt.Println("graphql demo.")
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gql.QueryType,
	})

	if err != nil {
		fmt.Errorf(err.Error())
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
