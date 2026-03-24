package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Memory struct {
	ID         string         `json:"id"`
	Memory     string         `json:"memory"`
	UserID     string         `json:"user_id"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  string         `json:"updated_at"`
	Metadata   map[string]any `json:"metadata,omitempty"`
	Categories []string       `json:"categories,omitempty"`
	Score      float64        `json:"score,omitempty"`
}

type AddResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	EventID string `json:"event_id"`
}

type Entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func (c *Client) AddMemory(text, userID string) (*AddResult, error) {
	body := map[string]any{
		"messages": []map[string]string{
			{"role": "user", "content": text},
		},
	}
	if userID != "" {
		body["user_id"] = userID
	}

	resp, err := c.Do("POST", "/v1/memories/", body)
	if err != nil {
		return nil, err
	}

	var results []AddResult
	if err := json.Unmarshal(resp, &results); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no result returned from API")
	}
	return &results[0], nil
}

func (c *Client) SearchMemories(query string, userID string, limit int) ([]Memory, error) {
	body := map[string]any{
		"query": query,
		"limit": limit,
	}
	if userID != "" {
		body["user_id"] = userID
	}

	resp, err := c.Do("POST", "/v1/memories/search/", body)
	if err != nil {
		return nil, err
	}

	var memories []Memory
	if err := json.Unmarshal(resp, &memories); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return memories, nil
}

func (c *Client) ListMemories(userID string, limit, page int) ([]Memory, error) {
	params := url.Values{}
	if userID != "" {
		params.Set("user_id", userID)
	}
	if limit > 0 {
		params.Set("page_size", strconv.Itoa(limit))
	}
	if page > 1 {
		params.Set("page", strconv.Itoa(page))
	}

	path := "/v1/memories/"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	resp, err := c.Do("GET", path, nil)
	if err != nil {
		return nil, err
	}

	// API returns a plain array without page param, paginated object with it.
	var memories []Memory
	if err := json.Unmarshal(resp, &memories); err == nil {
		return memories, nil
	}

	var paginated struct {
		Results []Memory `json:"results"`
	}
	if err := json.Unmarshal(resp, &paginated); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return paginated.Results, nil
}

func (c *Client) GetMemory(id string) (*Memory, error) {
	resp, err := c.Do("GET", "/v1/memories/"+id+"/", nil)
	if err != nil {
		return nil, err
	}

	var mem Memory
	if err := json.Unmarshal(resp, &mem); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &mem, nil
}

func (c *Client) UpdateMemory(id, text string) (*Memory, error) {
	body := map[string]string{"text": text}
	resp, err := c.Do("PUT", "/v1/memories/"+id+"/", body)
	if err != nil {
		return nil, err
	}

	var mem Memory
	if err := json.Unmarshal(resp, &mem); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &mem, nil
}

func (c *Client) DeleteMemory(id string) error {
	_, err := c.Do("DELETE", "/v1/memories/"+id+"/", nil)
	return err
}

func (c *Client) DeleteAllMemories(userID string) error {
	body := map[string]string{"user_id": userID}
	_, err := c.Do("DELETE", "/v1/memories/", body)
	return err
}

func (c *Client) ListEntities() ([]Entity, error) {
	resp, err := c.Do("GET", "/v1/entities/", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Results []Entity `json:"results"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return result.Results, nil
}

func (c *Client) DeleteEntity(entityType, id string) error {
	_, err := c.Do("DELETE", fmt.Sprintf("/v1/entities/%s/%s/", entityType, id), nil)
	return err
}
