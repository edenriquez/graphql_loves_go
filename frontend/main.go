package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
)

type query struct {
	Todos []struct {
		Text graphql.String
		Done graphql.Boolean
		User struct {
			Name graphql.String
		}
	}
}

/*
	query findTodos {
		todos {
			text
			done
			user {
				name
			}
		}
	}
*/
func main() {
	client := graphql.NewClient("http://localhost:8080/query", nil)
	result := &query{}
	err := client.Query(context.Background(), result, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.Todos[0])
}
