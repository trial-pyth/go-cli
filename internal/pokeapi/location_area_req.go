package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err :=http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{} ,err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return  LocationAreaResp{}, fmt.Errorf("bas status code: %v", resp.StatusCode)
	} 

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, nil
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, nil
	}

	return locationAreaResp, err

}