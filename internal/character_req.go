package rick_morty_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListCharacters(page *string) (CharactersResp, error) {
	endpoint := "/character?page=1"
	fullURL := baseURL + endpoint

	if page != nil {
		fullURL = *page
	}

	cacheData , ok := c.cache.Get(fullURL)
	if ok {
		characterList := CharactersResp{}
		err := json.Unmarshal(cacheData, &characterList)
		if err != nil {
			return CharactersResp{}, err 
		}

		return characterList, nil

	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return CharactersResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CharactersResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return CharactersResp{}, fmt.Errorf("bad status code: %v", err)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CharactersResp{}, err
	}

	characterlist := CharactersResp{}
	err = json.Unmarshal(dat, &characterlist)
	if err != nil {
		return CharactersResp{}, err
	}

	c.cache.Add(fullURL, dat)


	return characterlist, nil
}

func (c *Client) GetSingleCharacter(id string) (SingleCharacter, error) {
	endpoint := fmt.Sprintf("/character/%s", id)
	fulURL := baseURL + endpoint

	dat, ok := c.cache.Get(fulURL)
	if ok {
		fmt.Println("Cache hit!")
		character := SingleCharacter{}
		err := json.Unmarshal(dat, &character)
		if err != nil {
			return SingleCharacter{}, err
		}

		return character, nil
	}

	fmt.Println("Cache miss!")
	req, err := http.NewRequest("GET", fulURL, nil)
	if err != nil {
		return SingleCharacter{}, err 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SingleCharacter{}, err
	}

	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return SingleCharacter{}, err
	}

	character := SingleCharacter{}
	err = json.Unmarshal(dat, &character)
	if err != nil {
		return SingleCharacter{}, err
	}

	c.cache.Add(fulURL, dat)

	return character, nil
}