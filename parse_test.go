package main

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/net/html"
)

const (
	// Snapshot of Dashboard taken 05-Aug-2022
	example string = "example.220805.html"
)

type Tests map[string]struct {
	// Expected is the set of Regions that are expected to be present in this Service's Dashboard entry
	// bool value is tested
	// If Region is available, then true etc.
	Expected Regions
	// Unexpected is the set of Regions that are not expected to be present in this Service's Dashboard entry
	// bool value is ignored
	// Region isn't expected to be present so its status value is irrelevant
	Unexpected Regions
}

var (
	// Tests represent status of services as of the Dashboard snapshot
	tests = Tests{
		"Access Approval": {
			Expected: Regions{
				Global: true,
			},
			Unexpected: Regions{
				Americas:     false,
				Europe:       false,
				AsiaPacific:  false,
				MultiRegions: false,
			},
		},
		"AppSheet": {
			Expected: Regions{
				Americas: true,
				Europe:   true,
				Global:   true,
			},
			Unexpected: Regions{
				AsiaPacific:  false,
				MultiRegions: false,
			},
		},
		"Assured Workloads": {
			Expected: Regions{
				Americas:     true,
				Europe:       true,
				MultiRegions: false,
			},
			Unexpected: Regions{
				AsiaPacific: false,
				Global:      false,
			},
		},

		"Container Registry": {
			Expected: Regions{
				MultiRegions: true,
				Global:       true,
			},
			Unexpected: Regions{
				Americas:    false,
				Europe:      false,
				AsiaPacific: false,
			},
		},
		"Kubernetes Engine": {
			Expected: Regions{
				Americas:    true,
				Europe:      true,
				AsiaPacific: true,
			},
			Unexpected: Regions{
				MultiRegions: false,
				Global:       false,
			},
		},
	}
)

func TestParse(t *testing.T) {

	b, err := os.ReadFile(example)
	if err != nil {
		t.Fatal(err)
	}

	rdr := bytes.NewReader(b)
	node, err := html.Parse(rdr)
	if err != nil {
		t.Fatal(err)
	}

	services := parse(node)
	for _, s := range services {
		t.Run(s.Name, func(t *testing.T) {
			t.Logf("[%s]", s.Name)
			if test, ok := tests[s.Name]; ok {
				for region, want := range test.Expected {
					got, ok := s.Regions[region]
					if !ok {
						t.Fatalf("[%s] expected %s", s.Name, region)
					}
					// Exists
					if got != want {
						t.Fatalf("[%s] expected %s to be available", s.Name, region)
					}
				}
				for region := range test.Unexpected {
					if _, ok := s.Regions[region]; ok {
						t.Fatalf("[%s] unexpected %s", s.Name, region)
					}
				}
			}
		})
	}
}
