package RegionRepository

import (
	"RegionCLI/Models"
	"RegionCLI/Models/Representatives"
	"encoding/json"
	"os"
	"time"
)

func CreateRepresentative(name, address, phoneNumber string, employeeCount, regionId int) (int, error) {
	r := Representatives.Representative{
		Id:            Models.CountLines(Representatives.RepresentativeFilePath) + 1,
		Name:          name,
		Address:       address,
		PhoneNumber:   phoneNumber,
		EmployeeCount: employeeCount,
		RegionId:      regionId,
		CreatedDate:   time.Now(),
	}
	var err error
	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
