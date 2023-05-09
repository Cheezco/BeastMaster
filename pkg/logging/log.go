package logging

import "time"

type AppLog struct {
	Text    string    `json:"text"`
	Source  string    `json:"source"`
	Created time.Time `json:"created"`
}
