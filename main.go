package main

import (
	"RegionCLI/Models"
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

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r1, err := Models.CreateRegion("Shiraz")
	fmt.Println(r1, err)
}
