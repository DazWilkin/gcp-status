package main

import "testing"

var (
	// Name is not relevant to the test
	// Americas: 1
	// AsiaPacific: 2
	// Europe: 3
	// MultiRegions: 4
	// Global: 0 (Expected to be not-present in map)
	want = map[Region]uint16{
		Americas:     1,
		AsiaPacific:  2,
		Europe:       3,
		MultiRegions: 4,
		Global:       0,
	}
	test = Services{
		{
			Name: "",
			Regions: Regions{
				Americas:     false,
				AsiaPacific:  false,
				Europe:       false,
				MultiRegions: false,
			},
		},
		{
			Name: "",
			Regions: Regions{
				AsiaPacific:  false,
				Europe:       false,
				MultiRegions: false,
			},
		},
		{
			Name: "",
			Regions: Regions{
				Europe:       false,
				MultiRegions: false,
			},
		},
		{
			Name: "",
			Regions: Regions{
				MultiRegions: false,
			},
		},
	}
)

func TestByRegion(t *testing.T) {
	count := test.ByRegion()

	for r, want := range want {
		t.Run(r.String(), func(t *testing.T) {
			{
				got, ok := count[r]

				// If the region is not present and expected to be (i.e. want>0) then error
				// If the region is not present and not expected to be (i.e. Global, then want==0) then OK
				if !ok && want != 0 {
					t.Fatalf("expected %s", r.String())
				}
				if got != want {
					t.Fatalf("got %d; want: %d", got, want)
				}
			}
		})
	}
}
