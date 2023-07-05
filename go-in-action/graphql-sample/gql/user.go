package gql

import (
	"fmt"
	"strconv"

	"graphql-sample/gql/data"
	"graphql-sample/gql/types"

	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(types.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user, ok := p.Source.(types.User)
				if ok {
					//if u, ok := p.Source.(User); ok {
					return user.Name, nil
				} else {
					fmt.Println(user)
				}
				return nil, nil
			},
		},
		"age": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(types.User); ok {
					return user.Age, nil
				}
				return nil, nil
			},
		},
		"sex": &graphql.Field{
			Type: enumTypeSexType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(types.User); ok {
					return user.Sex, nil
				}
				return nil, nil
			},
		},
	},
})

var enumTypeSexType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Sex",
	Values: graphql.EnumValueConfigMap{
		"Male": &graphql.EnumValueConfig{
			Value: "Male",
		},
		"Female": &graphql.EnumValueConfig{
			Value: "Female",
		},
		"UnKnow": &graphql.EnumValueConfig{
			Value: "UnKnow",
		},
	},
})

func init() {
	queryFields["user"] = &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "User ip",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId, err := strconv.Atoi(p.Args["id"].(string))
			if err != nil {
				return nil, err
			}
			return data.GetUserByID(userId), nil
		},
	}

	UserType.AddFieldConfig("mates", &graphql.Field{
		Type: graphql.NewList(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(types.User); ok {
				return data.GetUserMates(user.ID), nil
			}
			return nil, nil
		},
	})
}
