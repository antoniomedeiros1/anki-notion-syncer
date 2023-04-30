package notionclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type NotionAPI struct {
	BaseURL string
	APIKey  string
}

func (n NotionAPI) createRequest(method, path string, body interface{}) (*http.Request, error) {
	url := n.BaseURL + path

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+n.APIKey)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (n NotionAPI) doRequest(req *http.Request, result interface{}) error {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("notionapi: bad status code " + resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (n NotionAPI) QueryDatabase(filter interface{}, database string) ([]map[string]interface{}, error) {
	path := "/v1/databases/" + database + "/query"
	req, err := n.createRequest("POST", path, filter)
	if err != nil {
		return nil, err
	}

	var response struct {
		Results []map[string]interface{} `json:"results"`
	}

	err = n.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

func (n NotionAPI) GetPage(PageID string) (map[string]interface{}, error) {
	path := "/v1/pages/" + PageID
	req, err := n.createRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}

	err = n.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
