package scheduler

import (
	"time"

	"github.com/zoh/go-schedule/repeat"
)

// Storer defines the backing storage implementation required by the scheduler
type Storer interface {
	AddEvent(name, description string, when, end, next time.Time, string repeat.Repeat) (Event, error)
	GetEvent(id string) (Event, error)
	UpdateEvent(event Event) (Event, error)
	GetEventsFiltered(start, end time.Time, completed bool) ([]Event, error)
}
