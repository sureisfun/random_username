package random_username

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateUsername() string {
	// Random adjective and noun
	adj := Adjectives[rand.Intn(len(Adjectives))]
	noun := Nouns[rand.Intn(len(Nouns))]
	number := rand.Intn(9999) + 1

	// Capitalize first letter of adjective and noun
	adj = strings.ToUpper(adj[:1]) + adj[1:]
	noun = strings.ToUpper(noun[:1]) + noun[1:]

	return fmt.Sprintf("%s%s-%d", adj, noun, number)
}
