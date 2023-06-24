package Representatives

import (
	"RegionCLI/Models"
	"encoding/json"
	"os"
	"time"
)

const (
	representativeFilePath = "./Representative.txt"
)

type representative struct {
	Id            int       `json:"Id"`
	Name          string    `json:"Name"`
	Address       string    `json:"Address"`
	PhoneNumber   string    `json:"PhoneNumber"`
	EmployeeCount int       `json:"EmployeeCount"`
	RegionId      int       `json:"RegionId"`
	CreatedDate   time.Time `json:"CreatedDate"`
}

func CreateRepresentative(name, address, phoneNumber string, employeeCount, regionId int) (int, error) {
	r := representative{
		Id:            Models.CountLines(representativeFilePath) + 1,
		Name:          name,
		Address:       address,
		PhoneNumber:   phoneNumber,
		EmployeeCount: employeeCount,
		RegionId:      regionId,
		CreatedDate:   time.Now(),
	}
	var err error
	f, oErr := os.OpenFile(representativeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	eErr := enc.Encode(r)
	if eErr != nil {
		err = eErr
	}
	return r.Id, err
}
