package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// sample response from : https://circleci.com/docs/api/v2/#tag/Schedule

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func TestGetAllSchedules(t *testing.T) {
	ctx := context.Background()
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	defer server.Close()
	url, err := url.Parse(server.URL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse : %s", server.URL))
	}
	client = &Client{
		BaseURL:    url,
		HTTPClient: http.DefaultClient,
	}
	mux.HandleFunc("/project/gh/CircleCI-Public/api-preview-docs/schedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, `
		{
		  "items": [
		    {
		      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		      "timetable": {
		        "per-hour": 0,
		        "hours-of-day": [
		          0
		        ],
		        "days-of-week": [
		          "TUE"
		        ]
		      },
		      "updated-at": "2019-08-24T14:15:22Z",
		      "name": "string",
		      "created-at": "2019-08-24T14:15:22Z",
		      "project-slug": "gh/CircleCI-Public/api-preview-docs",
		      "parameters": {
		        "deploy_prod": true,
		        "branch": "feature/design-new-api"
		      },
		      "actor": {
		        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		        "login": "string",
		        "name": "string"
		      },
		      "description": "string"
		    }
		  ],
		  "next_page_token": "string"
		}
		`)
	})
	resp, err := client.GetAllSchedules(ctx, "CircleCI-Public/api-preview-docs")
	if err != nil {
		t.Fatalf("Client.GetAllSchedules returns error %v", err)
	}

	if resp[0].ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("expected id to be %s, but was %s", "497f6eca-6276-4993-bfeb-53cbbbba6f08", resp[0].ID)
	}
}

func TestCreateSchedule(t *testing.T) {
	ctx := context.Background()
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	defer server.Close()
	url, err := url.Parse(server.URL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse : %s", server.URL))
	}
	client = &Client{
		BaseURL:    url,
		HTTPClient: http.DefaultClient,
	}
	mux.HandleFunc("/project/gh/CircleCI-Public/api-preview-docs/schedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, `
		{
		  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		  "timetable": {
		    "per-hour": 0,
		    "hours-of-day": [
		      0
		    ],
		    "days-of-week": [
		      "TUE"
		    ]
		  },
		  "updated-at": "2019-08-24T14:15:22Z",
		  "name": "string",
		  "created-at": "2019-08-24T14:15:22Z",
		  "project-slug": "gh/CircleCI-Public/api-preview-docs",
		  "parameters": {
		    "deploy_prod": true,
		    "branch": "feature/design-new-api"
		  },
		  "actor": {
		    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		    "login": "string",
		    "name": "string"
		  },
		  "description": "string"
		}
		`)
	})

	project := "CircleCI-Public/api-preview-docs"
	attr := `
	{
	  "description": "string",
	  "name": "string",
	  "timetable": {
	    "per-hour": 0,
	    "hours-of-day": [
	      0
	    ],
	    "days-of-week": [
	      "TUE"
	    ]
	  },
	  "attribution-actor": "current",
	  "parameters": {
	    "deploy_prod": true,
	    "branch": "feature/design-new-api"
	  }
	}
	`
	resp, err := client.CreateSchedule(ctx, project, attr)
	if err != nil {
		t.Fatalf("Client.CreateSchedule returns error %v", err)
	}

	if resp.Name != "string" {
		t.Errorf("expected id to be %s, but was %s", "string", resp.Name)
	}
}

func TestUpdateSchedule(t *testing.T) {
	ctx := context.Background()
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	defer server.Close()
	url, err := url.Parse(server.URL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse : %s", server.URL))
	}
	scheduleId := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	client = &Client{
		BaseURL:    url,
		HTTPClient: http.DefaultClient,
	}
	mux.HandleFunc(fmt.Sprintf("/schedule/%s", scheduleId), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, `
		{
		  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		  "timetable": {
		    "per-hour": 0,
		    "hours-of-day": [
		      0
		    ],
		    "days-of-week": [
		      "TUE"
		    ]
		  },
		  "updated-at": "2019-08-24T14:15:22Z",
		  "name": "string",
		  "created-at": "2019-08-24T14:15:22Z",
		  "project-slug": "gh/CircleCI-Public/api-preview-docs",
		  "parameters": {
		    "deploy_prod": true,
		    "branch": "feature/design-new-api"
		  },
		  "actor": {
		    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		    "login": "string",
		    "name": "string"
		  },
		  "description": "string"
		}
		`)
	})

	attr := `
	{
	  "description": "string",
	  "name": "string",
	  "timetable": {
	    "per-hour": 0,
	    "hours-of-day": [
	      0
	    ],
	    "days-of-week": [
	      "TUE"
	    ]
	  },
	  "attribution-actor": "current",
	  "parameters": {
	    "deploy_prod": true,
	    "branch": "feature/design-new-api"
	  }
	}
	`
	resp, err := client.UpdateSchedule(ctx, scheduleId, attr)
	if err != nil {
		t.Fatalf("Client.GetAllSchedules returns error %v", err)
	}

	if resp.ID != scheduleId {
		t.Errorf("expected id to be %s, but was %s", "497f6eca-6276-4993-bfeb-53cbbbba6f08", resp.ID)
	}
}
