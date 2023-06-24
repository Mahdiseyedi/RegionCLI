package main

import (
	"RegionCLI/FileRepository/RegionRepository"
	"fmt"
)

func main() {
	n, err := RegionRepository.DeleteRegion(5)
	fmt.Println(n, err)
}
