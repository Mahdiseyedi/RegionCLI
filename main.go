package main

import (
	"RegionCLI/FileRepository/RepresentativeRepository"
	"fmt"
)

func main() {
	rt2, _ := RepresentativeRepository.GetRepresentativeByID(0)
	fmt.Println(rt2)
}
