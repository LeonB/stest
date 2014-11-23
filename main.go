package main

import (
	"fmt"
	"github.com/leonb/stest/user"
)

func main() {
	fmt.Println("Test")
	var repo user.Repository
	repo = user.MysqlRepository{}

	collection, _ := repo.All()
	collection.AddActiveFilter()
	collection.Slice(10, 20)

	users, _ := collection.Get()
	fmt.Println(users)
	fmt.Println(len(users))
}
