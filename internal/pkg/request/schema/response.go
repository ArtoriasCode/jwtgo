package schema

import (
	"time"
)

type Cookie struct {
	Name     string
	Value    string
	Duration time.Duration
}
