package entities

import "time"

type Company struct {
	name         string
	uid          string
	display_name string
	external_id  string
	create_time  time.Time
	update_time  time.Time
}
