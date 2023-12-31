package Representatives

import (
	"time"
)

const (
	RepresentativeFilePath = "./Representative.txt"
)

type Representative struct {
	Id            int       `json:"Id"`
	Name          string    `json:"Name"`
	Address       string    `json:"Address"`
	PhoneNumber   string    `json:"PhoneNumber"`
	EmployeeCount int       `json:"EmployeeCount"`
	RegionId      int       `json:"RegionId"`
	CreatedDate   time.Time `json:"CreatedDate"`
}
