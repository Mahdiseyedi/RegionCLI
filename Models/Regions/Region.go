package Regions

const (
	RegionFilePath = "./Regions.txt"
)

type Region struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}
