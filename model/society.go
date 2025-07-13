package model

import "time"

type Society struct {
	Id      int
	Name    string
	Address string
	Pin     string
	Created time.Time
}
