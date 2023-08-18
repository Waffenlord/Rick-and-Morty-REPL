package rick_morty_api

import (
	"net/http"
	"time"
)

const baseURL = "https://rickandmortyapi.com/api"

type Client struct {
	httpClient http.Client
}


func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}