package rick_morty_api

import (
	"net/http"
	"time"
	"github.com/Waffenlord/Rick-and-Morty-REPL/internal/rick_morty_cache"
)

const baseURL = "https://rickandmortyapi.com/api"

// Struct representing the current user
type Client struct {
	httpClient http.Client
	cache rickmortycache.Cache
	SavedChars SavedCharacters
}

type SavedCharacters struct {
	Characters map[int]SavedChar `json:"characters"`
}

// Create an instance of Client
func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,

		},
		cache: rickmortycache.NewCache(interval),
		SavedChars: SavedCharacters{
			Characters: make(map[int]SavedChar),
		},
	}
}