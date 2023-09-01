package rick_morty_api

import (
	"net/http"
	"time"
	"github.com/Waffenlord/Rick-and-Morty-REPL/internal/rick_morty_cache"
)

const baseURL = "https://rickandmortyapi.com/api"

type Client struct {
	httpClient http.Client
	cache rickmortycache.Cache
}


func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,

		},
		cache: rickmortycache.NewCache(interval),
	}
}