package random_username

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed adjectives.txt
var adjectivesData string

//go:embed nouns.txt
var nounsData string

var (
	Adjectives []string
	Nouns      []string
)

func init() {
	Adjectives = parseWords(adjectivesData)
	Nouns = parseWords(nounsData)
}

// parseWords splits embedded text data into a slice of words
func parseWords(data string) []string {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	words := make([]string, 0, len(lines))
	for _, line := range lines {
		if word := strings.TrimSpace(line); word != "" {
			words = append(words, word)
		}
	}
	return words
}

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
