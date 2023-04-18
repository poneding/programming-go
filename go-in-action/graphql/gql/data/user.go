package data

import "github.com/poneding/learning-go/practice/graphql/gql/types"

var users = map[int]types.User{
	1: {
		ID:   1,
		Name: "Pone",
		Age:  20,
		Sex:  "Male",
	},
	2: {
		ID:   2,
		Name: "Bob",
		Age:  20,
		Sex:  "Male",
	},
	3: {
		ID:   3,
		Name: "Jack",
		Age:  20,
		Sex:  "Male",
	},
	4: {
		ID:   4,
		Name: "John",
		Age:  20,
		Sex:  "Male",
	},
	5: {
		ID:   5,
		Name: "Alice",
		Age:  20,
		Sex:  "Female",
	},
}

func GetUserByID(id int) types.User {
	return users[id]
}

func GetUserMates(id int) []types.User {
	temp := make([]types.User, 0)
	for i := 1; i <= len(users); i++ {
		if i == id {
			continue
		}
		temp = append(temp, users[i])
	}
	return temp
}
