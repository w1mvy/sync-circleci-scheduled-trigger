# sync-circleci-scheduled-trigger

sync-circleci-scheduled-triggeris a command-line tool to manage triggers for circleci's scheduled pipeline

## Usage

```
$ export CIRCLECI_TOKEN=your_circleci_token
$ go get -u github.com/w1mvy/sync-circleci-scheduled-trigger
$ sync-circleci-scheduled-trigger --config=pathto/config.json
```

## Config

| key | description |
| --- | --- |
| project | your repository name |
| schedules | list of [schedule parameter](https://circleci.com/docs/api/v2/#operation/createSchedule) |

```
{
  "project": "w1mvy/sync-circleci-scheduled-trigger",
  "schedules": [
    {
      "name": "example schedule 1",
      "description" : "string",
      "attribution-actor": "system",
      "timetable": {
        "per-hour" : 2,
        "hours-of-day" : [ 17 ],
        "days-of-week" : [ "WED" ]
      },
      "parameters": {
        "branch" : "master"
      }
    },
    {
      "name": "example schedule 2",
      "description" : "string",
      "attribution-actor": "system",
      "timetable": {
        "per-hour" : 1,
        "hours-of-day" : [ 17 ],
        "days-of-week" : [ "MON", "FRI" ]
      },
      "parameters": {
        "branch" : "another-branch"
      }
    }
  ]
}
```
