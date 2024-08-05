package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
)

type Noun struct {
	Noun    string `json:"n"`
	Article string `json:"a"`
}

// As demo returns at random one of 3 colours matching pair
func demoGetRandomKeyValue() (int, string) {
	word := map[string]int{"red": helper.RED, "blue": helper.BLUE, "green": helper.GREEN}

	keys := make([]string, 0, len(word))
	for key := range word {
		keys = append(keys, key)
	}

	index := rand.Intn(len(keys))

	key := keys[index]
	value := word[key]

	return value, key
}

// This function returns a noun and its matching gender article die der das
func GetNoun() (int, string) { // Note only func that start with CAPITAL LETTER is automatically exported and public

	if helper.DemoMode {
		return demoGetRandomKeyValue()
	}

	noun := fetchRandomNoun()
	value, err := getArticle(noun.Article)
	if err != nil {
		log.Fatalf("Failed to assign article: %v", err)
	}
	key := noun.Noun

	return value, key
}

func getArticle(article string) (int, error) {
	switch article {
	case "M":
		return helper.BLUE, nil
	case "F":
		return helper.RED, nil
	case "N":
		return helper.GREEN, nil
	}

	return 0, fmt.Errorf("unepected input")
}

func fetchRandomNoun() Noun {

	// Todo replace with single DB/API noun fetch
	nounList := fetchListFromFile()
	index := rand.Intn(len(nounList))

	return nounList[index]
}

func fetchListFromFile() []Noun {

	path := filepath.Join("data", "100_article_nouns.json")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var nouns []Noun
	err = json.Unmarshal(byteValue, &nouns)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	return nouns
}
