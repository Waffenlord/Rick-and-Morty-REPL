package rick_morty_api

import "time"

type CharactersResp struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  *string `json:"next"`
		Prev  *string    `json:"prev"`
	} `json:"info"`
	Results []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Species string `json:"species"`
		Type    string `json:"type"`
		Gender  string `json:"gender"`
		Origin  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"origin"`
		Location struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		Image   string   `json:"image"`
		Episode []string `json:"episode"`
		URL     string   `json:"url"`
		Created time.Time `json:"created"`
	} `json:"results"`
}


type SingleCharacter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string    `json:"image"`
	Episode []string  `json:"episode"`
	URL     string    `json:"url"`
	Created time.Time `json:"created"`
}


type SavedChar struct {
	ID 			int `json:"id"`
	Name    	string `json:"name"`
	Status  	string `json:"status"`
	Species 	string `json:"species"`
	Type    	string `json:"type"`
	Gender  	string `json:"gender"`
	Origin  	string `json:"origin"`
	Location 	string `json:"location"`
	Image 		string `json:"image"`
}