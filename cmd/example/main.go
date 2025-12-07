package main

import (
	"fmt"

	"github.com/sureisfun/random_username"
)

func main() {
	username := random_username.GenerateUsername()
	fmt.Println(username)
}
