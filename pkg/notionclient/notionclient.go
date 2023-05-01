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

	var err error
	var req *http.Request
	var jsonBody []byte

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+n.APIKey)
	req.Header.Set("Notion-Version", "2022-06-28")
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

func (n NotionAPI) GetBlock(BlockID string) (map[string]any, error) {
	path := "/v1/blocks/" + BlockID
	req, err := n.createRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var res map[string]any

	err = n.doRequest(req, &res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (n NotionAPI) GetBlockChildren(BlockID string) (map[string]any, error) {
	path := "/v1/blocks/" + BlockID + "/children?page_size=100"
	req, err := n.createRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var res map[string]any

	err = n.doRequest(req, &res)
	if err != nil {
		return nil, err
	}

	// pagination
	// if res["nextCursor"] != nil {
	// 	n.GetBlockChildren(BlockID, res["nextCursor"])
	// }

	return res, nil

}
