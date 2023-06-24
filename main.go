package main

import (
	"RegionCLI/Models/Representatives"
	"fmt"
	"time"
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

func main() {
	r1, err := Representatives.CreateRepresentative("Mahdi", "Tehran-Mirza", "091233", 6, 1)
	fmt.Println(r1, err)
}
