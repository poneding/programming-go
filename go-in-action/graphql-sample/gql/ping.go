package gql

import "github.com/graphql-go/graphql"

func init() {
	queryFields["ping"] = &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pong(), nil
		},
	}
}

func pong() string{
	return "ping"
}
