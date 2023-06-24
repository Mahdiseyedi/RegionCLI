package main

import (
	"RegionCLI/FileRepository/RepresentativeRepository"
	"fmt"
)

func main() {
	err := RepresentativeRepository.EditRepresentative(1, "Ali", "", "0911123", 32, 5)
	fmt.Println(err)
}
