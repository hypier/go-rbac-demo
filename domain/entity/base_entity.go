package entity

import "time"

type BaseEntity struct {
	CreateTime time.Time
	UpdateTime time.Time
	Version    int
}
