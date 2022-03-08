package main

import (
	"context"
	"reflect"
)

func Sync(ctx context.Context, client *Client, config *Config, dryRun bool) ([]*Item, error) {
	items, err := client.GetAllSchedules(ctx, config.Project)
	if err != nil {
		return nil, err
	}
	var savedItem []*Item
	for _, createSchedule := range FilterCreate(items, config.Schedules) {
		if !dryRun {
			item, err := client.CreateSchedule(ctx, config.Project, createSchedule)
			if err != nil {
				return nil, err
			}
			savedItem = append(savedItem, item)
		}
		client.Logger.Printf("create schedule: %s\n", createSchedule.Name)
	}
	for _, patchSchedule := range FilterPatch(items, config.Schedules) {
		if !dryRun {
			item, err := client.UpdateSchedule(ctx, patchSchedule.ScheduleId, patchSchedule)
			if err != nil {
				return nil, err
			}
			savedItem = append(savedItem, item)
		}
		client.Logger.Printf("update schedule: %s\n", patchSchedule.Schedule.Name)
	}
	return savedItem, nil
}

type PatchSchedule struct {
	ScheduleId string
	Schedule   *Schedule
}

func FilterPatch(items []*Item, schedules []*Schedule) []*PatchSchedule {
	var patchSchedules []*PatchSchedule
	for _, item := range items {
		for _, schedule := range schedules {
			if IsMatch(item, schedule) && DiffExist(item, schedule) {
				patchSchedules = append(patchSchedules, &PatchSchedule{ScheduleId: item.ID, Schedule: schedule})
				break
			}
		}
	}
	return patchSchedules
}

func FilterCreate(items []*Item, schedules []*Schedule) []*Schedule {
	var createSchedules []*Schedule
	for _, item := range items {
		for _, schedule := range schedules {
			if !IsMatch(item, schedule) {
				createSchedules = append(createSchedules, schedule)
				break
			}
		}
	}
	return createSchedules
}

// always diff exist when schedule.AttributionActor not system
func DiffExist(item *Item, schedule *Schedule) bool {
	if item.Description != schedule.Description || !reflect.DeepEqual(item.Timetable, schedule.Timetable) || !reflect.DeepEqual(item.Parameters, schedule.Parameters) || !(item.Actor.Login == "system-actor" && schedule.AttributionActor == "system") {
		return true
	}
	return false
}

func IsMatch(item *Item, schedule *Schedule) bool {
	if item.Name == schedule.Name {
		return true
	}
	return false
}
