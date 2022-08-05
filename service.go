package main

// Region is Google Cloud region
type Region uint8

const (
	Americas     Region = 0
	Europe       Region = 1
	AsiaPacific  Region = 2
	MultiRegions Region = 3
	Global       Region = 4
)

// String converts a Region to a string
func (r Region) String() string {
	switch r {
	case Americas:
		return "Americas"
	case Europe:
		return "Europe"
	case AsiaPacific:
		return "Asia Pacific"
	case MultiRegions:
		return "Multi-regions"
	case Global:
		return "Global"
	default:
		return "Unknown"
	}

}

type Regions map[Region]bool

// Service describes a Google Cloud Platform service status
type Service struct {
	Name    string
	Regions Regions
}

// NewService is a function that returns a new Service
func NewService(name string) Service {
	return Service{
		Name:    name,
		Regions: map[Region]bool{},
	}
}
