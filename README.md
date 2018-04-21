# go-schedule

Database (or other storage) backed scheduling module (WIP).

[![Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/zoh/go-schedule)
[![GitHub tag](https://img.shields.io/github/tag/zoh/go-schedule.svg)](https://github.com/zoh/go-schedule)
[![Build Status](https://travis-ci.org/zoh/go-schedule.svg?branch=master)](https://travis-ci.org/zoh/go-schedule)

This is designed to allow user or application scheduling of events (one off or a variety of repetitions) that can be maintained in a datastore for consistency / coherency.


## Usage


### Creating the scheduler
```go
import (
    "gopkg.in/zoh/go-schedule.v0"
)

s := scheduler.NewScheduler(storer Storer, startTime time.Time, tickRate time.Duration)
go s.Run()
```

### Adding an event
```go
e, err := s.Schedule(name, description string, when, end time.Time, repeat repeat.Repeat)
```

### Subscribing to events
```go
select {
    case e, ok := s.Out:
    if !ok {
        // Scheduler exited / channel closed
    }
    // Do things with event instance
}

```
---

## Execute event 
```go

type MockEvent struct {
    helpers.DefaultEvent
}

func (de *MockEvent) Execute() error {
    log.Println("OK!")

	return nil
}

type MockStorer struct {
	Events []*MockEvent
}
```



