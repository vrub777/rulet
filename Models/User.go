package Models

import "time"

type User struct {
	Id               int
	Name             string
	Email            string
	DateRegistration time.Time
	DateActivation   time.Time
	IsLocked         bool
	IsActivate       bool
	CountTryAuth     int
	Phone            string
	Description      string
	PassworsHash     string
	Rating           float32
	Error
}
