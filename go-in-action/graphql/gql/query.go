package gql

import "github.com/graphql-go/graphql"

var queryFields = graphql.Fields{} //make(map[string]*graphql.Field, 0)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Query",
	Fields: queryFields,
})

var MutationType =graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",

})

