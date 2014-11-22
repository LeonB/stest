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
	collection, _ = collection.AddActiveFilter()
	collection, _ = collection.Slice(10, 20)
	fmt.Println(collection.Get())
}
