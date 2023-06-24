package main

import (
	"RegionCLI/FileRepository/RepresentativeRepository"
	"fmt"
)

func main() {
	rt2, _ := RepresentativeRepository.GetRepresentatives()
	fmt.Println(rt2[1].Name)
}
