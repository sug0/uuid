package uuid

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	table := []struct {
		input string
		valid bool
	}{
		{"d21484ef-9159-4d45-bebc-9e9439f832b1", true},
		{"810d3661-9f7c-471a-ac09-150eeb9e19f5", true},
		{"92BAC1D0-3079-4B96-A64D-03F6CA5B1185", true},
		{"1f508e72-3876-4086-9dae-8dfce4614c8b", true},
		{"aae4d00d-adc3-489e-9e84-0c0848352b20", true},
		{"aae4d00d-adc3-489e-9e84-0c0848352b2", false},
		{"aae4d00d-adc3-489e-9e84-0c0848352b20\n", false},
		{"aae4d00d0adc3-489e-9e84-0c0848352b20", false},
		{"aae4d00d-adc3f489e-9e84-0c0848352b20", false},
		{"aae4d00d-adc3-489ea9e84-0c0848352b20", false},
		{"aae4d00d-adc3-489e-9e84b0c0848352b20", false},
		{"aae4d0xd-adc3-489e-9e84-0c0848352b20", false},
		{"aae4d00d-adc3-489e-9e84-0c0848352-20", false},
	}
	for _, row := range table {
		u, err := Parse(row.input)
		if err != nil {
			if row.valid {
				t.Errorf("Parse(%#v), got error, want success", row.input)
			}
		} else {
			if !row.valid {
				t.Errorf("Parse(%#v), got success, want error", row.input)
			} else {
				s := u.String()
				if !strings.EqualFold(row.input, s) {
					t.Errorf("Parse(%#v), got %#v, want %#v",
						row.input, s, strings.ToLower(row.input))
				}
			}
		}
	}
}
