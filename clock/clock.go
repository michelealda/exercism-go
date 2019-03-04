package clock

import "fmt"

//Clock define the clock type
type Clock struct {
	hour, minute int
}

//New construct a new clock
func New(hour, minute int) Clock {
	normH := (24 + hour%24) % 24
	totM := (1440 + normH*60 + minute%1440) % 1440
	return Clock{totM / 60, totM % 60}
}

//String print a clock in a readable way
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

//Add minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

//Subtract minutes to a clock
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
