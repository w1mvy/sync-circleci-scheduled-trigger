package main

import (
	"testing"
)

func TestFilterCreate_noexist(t *testing.T) {
	items := []*Item{
		{
			Name: "test1",
		},
		{
			Name: "test2",
		},
		{
			Name: "test3",
		},
	}
	schedules := []*Schedule{
		{
			Name: "test1",
		},
		{
			Name: "test2",
		},
	}

	create := FilterCreate(items, schedules)
	if len(create) != 0 {
		t.Errorf("FilterCreate expected to return empty array: returns %v", create)
	}
}

func TestFilterCreate_exist(t *testing.T) {
	items := []*Item{
		{
			Name: "test1",
		},
		{
			Name: "test3",
		},
	}
	schedules := []*Schedule{
		{
			Name: "test1",
		},
		{
			Name: "test2",
		},
		{
			Name: "test2",
		},
	}

	create := FilterCreate(items, schedules)
	if create[0].Name != "test2" {
		t.Errorf("FilterCreate expected to return test2 : returns %v", create[0].Name)
	}
}

func TestFilterPatch_noexist(t *testing.T) {
	items := []*Item{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			Name:        "test3",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
	}
	schedules := []*Schedule{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test3",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
	}

	patch := FilterPatch(items, schedules)
	if len(patch) != 0 {
		t.Errorf("FilterPatch expected to return empty array: returns %v", patch)
	}
}

func TestFilterPatch_exists(t *testing.T) {
	items := []*Item{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "c",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			Name:        "test3",
			Description: "nodiff",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
	}
	schedules := []*Schedule{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "w1mvy",
		},
		{
			Name:        "test3",
			Description: "nodiff",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test4",
			Description: "new schedule",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
	}

	patch := FilterPatch(items, schedules)
	if len(patch) != 2 {
		t.Errorf("FilterPatch expected to return two elements: returns %v", patch)
	}
	if patch[0].Schedule.Name != "test1" {
		t.Errorf("FilterPatch expected to return two elements: returns %v", patch)
	}
	if patch[1].Schedule.Name != "test2" {
		t.Errorf("FilterPatch expected to return two elements: returns %v", patch)
	}
}

func TestFilterDelete_noexist(t *testing.T) {
	items := []*Item{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
	}
	schedules := []*Schedule{
		{
			Name:        "test1",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test3",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
	}

	deleteItems := FilterDelete(items, schedules)
	if len(deleteItems) != 0 {
		t.Errorf("FilterDelete expected to return empty array: returns %v", deleteItems)
	}
}

func TestFilterDelete_exists(t *testing.T) {
	items := []*Item{
		{
			ID:          "not-exist-id-from-schedules",
			Name:        "not-exist-id-from-schedules",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			ID:          "test2",
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			ID:          "test3-diff-from-schedule",
			Name:        "test3-diff-from-schedule",
			Description: "nodiff",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
		{
			ID:          "test4",
			Name:        "test4",
			Description: "schedule",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			Actor: Actor{
				Login: "system-actor",
				Name:  "Scheduled",
			},
		},
	}
	schedules := []*Schedule{
		{
			Name:        "test2",
			Description: "desc",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "w1mvy",
		},
		{
			Name:        "test3",
			Description: "nodiff",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
		{
			Name:        "test4",
			Description: "schedule",
			Parameters: Parameters{
				"testA": "a",
				"testB": "b",
			},
			AttributionActor: "system",
		},
	}

	deleteItems := FilterDelete(items, schedules)
	if len(deleteItems) != 2 {
		t.Errorf("FilterDelete expected to return two elements: returns %v", deleteItems)
	}
	if deleteItems[0].ID != "not-exist-id-from-schedules" {
		t.Errorf("FilterDelete expected to return two elements: returns %v", deleteItems)
	}
	if deleteItems[1].ID != "test3-diff-from-schedule" {
		t.Errorf("FilterDelete expected to return two elements: returns %v", deleteItems)
	}
}

func TestIsMatch_notmatch(t *testing.T) {
	item := &Item{Name: "test"}
	attr := &Schedule{Name: "notmatch"}
	if IsMatch(item, attr) {
		t.Errorf("IsMatch must be return false. item: %v, attr: %v", item, attr)
	}
}

func TestIsMatch_match(t *testing.T) {
	item := &Item{Name: "test"}
	attr := &Schedule{Name: "test"}
	if !IsMatch(item, attr) {
		t.Errorf("IsMatch must be return true. item: %v, attr: %v", item, attr)
	}
}

func TestDiffExist_nodiff(t *testing.T) {
	item := &Item{
		Name:        "test1",
		Description: "desc",
		Parameters: Parameters{
			"testA": true,
			"testB": "b",
		},
		Actor: Actor{
			Login: "system-actor",
			Name:  "Scheduled",
		},
	}

	schedule := &Schedule{
		Name:        "test1",
		Description: "desc",
		Parameters: Parameters{
			"testA": true,
			"testB": "b",
		},
		AttributionActor: "system",
	}
	if DiffExist(item, schedule) {
		t.Errorf("DiffExist expected to be false. item %v, schedule: %v", item, schedule)
	}
}

func TestDiffExist_diff(t *testing.T) {
	item := &Item{
		Description: "desc",
		Parameters: Parameters{
			"testA": "a",
			"testB": "b",
		},
		Actor: Actor{
			ID:   "system-actor",
			Name: "Scheduled",
		},
	}
	schedule := &Schedule{
		Description: "desc",
		Parameters: Parameters{
			"testA": "a",
			"testC": "c",
		},
		AttributionActor: "current",
	}
	if !DiffExist(item, schedule) {
		t.Errorf("DiffExist expected to be true. item %v, schedule: %v", item, schedule)
	}
}
