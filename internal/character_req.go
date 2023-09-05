package rick_morty_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListCharacters(page *string) (CharacterResp, error) {
	endpoint := "/character?page=1"
	fullURL := baseURL + endpoint

	if page != nil {
		fullURL = *page
	}

	cacheData , ok := c.cache.Get(fullURL)
	if ok {
		characterList := CharacterResp{}
		err := json.Unmarshal(cacheData, &characterList)
		if err != nil {
			return CharacterResp{}, err 
		}

		return characterList, nil

	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return CharacterResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CharacterResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return CharacterResp{}, fmt.Errorf("bad status code: %v", err)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CharacterResp{}, err
	}

	characterlist := CharacterResp{}
	err = json.Unmarshal(dat, &characterlist)
	if err != nil {
		return CharacterResp{}, err
	}

	c.cache.Add(fullURL, dat)


	return characterlist, nil
}