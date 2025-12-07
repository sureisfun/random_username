package random_username

import (
	"fmt"
	"math/rand"
)

func GenerateUsername() string {
	// Random adjective and noun
	adj := data.Adjectives[rand.Intn(len(data.Adjectives))]
	noun := data.Nouns[rand.Intn(len(data.Nouns))]
	number := rand.Intn(9999) + 1

	return fmt.Sprintf("%s%s%d", adj, noun, number)
}
