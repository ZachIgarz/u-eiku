package alarm

import "time"

//Alarm  is the struct of a alarm
type Alarm struct {
	Name string
	Date time.Time
	Days []string
}
