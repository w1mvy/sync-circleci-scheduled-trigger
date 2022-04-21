package main

import (
	"context"
	"fmt"
	"reflect"
)

func Sync(ctx context.Context, client *Client, config *Config, dryRun bool, forceSync bool) ([]*Item, error) {
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
		fmt.Printf("create schedule: %s\n", createSchedule.Name)
	}
	for _, patchSchedule := range FilterPatch(items, config.Schedules) {
		if !dryRun {
			item, err := client.UpdateSchedule(ctx, patchSchedule.ScheduleId, patchSchedule.Schedule)
			if err != nil {
				return nil, err
			}
			savedItem = append(savedItem, item)
		}
		fmt.Printf("update schedule: %s", patchSchedule.Schedule.Name)
	}
	if forceSync {
		for _, deleteItem := range FilterDelete(items, config.Schedules) {
			if !dryRun {
				message, err := client.DeleteSchedule(ctx, deleteItem.ID)
				if err != nil {
					return nil, err
				}
				fmt.Printf("delete schedule message: %s\n", message)
			}
			fmt.Printf("delete schedule: %s\n", deleteItem.Name)
		}
	}
	return savedItem, nil
}

type PatchSchedule struct {
	ScheduleId string
	Schedule   *Schedule
}

func FilterPatch(items []*Item, schedules []*Schedule) []*PatchSchedule {
	var patchSchedules []*PatchSchedule
	for _, schedule := range schedules {
		for _, item := range items {
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
	if len(items) == 0 {
		return schedules
	}
	for _, schedule := range schedules {
		found := false
		for _, item := range items {
			if IsMatch(item, schedule) {
				found = true
				break
			}
		}
		if !found {
			createSchedules = append(createSchedules, schedule)
		}
	}
	return createSchedules
}

func FilterDelete(items []*Item, schedules []*Schedule) []*Item {
	var deleteItems []*Item
	for _, item := range items {
		found := false
		for _, schedule := range schedules {
			if IsMatch(item, schedule) {
				found = true
				break
			}
		}
		if !found {
			deleteItems = append(deleteItems, item)
		}
	}
	return deleteItems
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
