package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Token      string
}

type Item struct {
	ID          string `json:"id"`
	Timetable   Timetable
	UpdatedAt   time.Time  `json:"updated-at"`
	Name        string     `json:"name"`
	CreatedAt   time.Time  `json:"created-at"`
	ProjectSlug string     `json:"project-slug"`
	Parameters  Parameters `json:"parameters"`
	Actor       Actor      `json:"actor"`
	Description string     `json:"description"`
}

type Parameters map[string]interface{}

type Actor struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type Timetable struct {
	PerHour    int      `json:"per-hour"`
	HourOfDay  []int    `json:"hours-of-day"`
	DaysOfWeek []string `json:"days-of-week"`
}

type GetScheduleResponse struct {
	Items         []*Item `json:"items"`
	NextPageToken string  `json:"next_page_token"`
}

func NewClient() (*Client, error) {
	token := os.Getenv("CIRCLECI_TOKEN")
	if len(token) == 0 {
		return nil, fmt.Errorf("must set env `CIRCLECI_TOKEN`")
	}
	return &Client{
		BaseURL:    &url.URL{Host: "circleci.com", Scheme: "https", Path: "/api/v2/"},
		HTTPClient: http.DefaultClient,
		Token:      token,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, bodyStruct interface{}) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, spath)
	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if bodyStruct != nil {
		body, err := json.Marshal(bodyStruct)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("circle-token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// TODO: support page-token
func (c *Client) GetAllSchedules(ctx context.Context, project string) ([]*Item, error) {
	spath := fmt.Sprintf("project/gh/%s/schedule", project)
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var resp GetScheduleResponse
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func (c *Client) CreateSchedule(ctx context.Context, project string, body interface{}) (*Item, error) {
	spath := fmt.Sprintf("project/gh/%s/schedule", project)
	req, err := c.newRequest(ctx, "POST", spath, body)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var resp Item
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSchedule(ctx context.Context, scheduleId string, body interface{}) (*Item, error) {
	spath := fmt.Sprintf("schedule/%s", scheduleId)
	req, err := c.newRequest(ctx, "PATCH", spath, body)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var resp *Item
	if err := decodeBody(res, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
