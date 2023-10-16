package schedule

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	vsys    string
	dg      string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 test recurring weekly", "", "", version.Number{10, 0, 0, ""}, Entry{
			Name:         "one",
			WeeklySunday: []string{""},
		}},
		{"v1 test non-recurring", "", "", version.Number{10, 0, 0, ""}, Entry{
			Name:         "one",
			NonRecurring: []string{""},
		}},
	}
}
