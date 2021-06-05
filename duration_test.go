package dateparse

import (
	"testing"
)

func TestCheckWordNumber(t *testing.T) {
	for _, tt := range []struct {
		input  string
		output float64
	}{
		{
			"пол",
			0.5,
		},
		{
			"четверть",
			0.25,
		},
	} {
		result := checkWordNumber(tt.input)
		if result != tt.output {
			t.Errorf("%v != %v", result, tt.output)
		}
	}
}
