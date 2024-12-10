package structColumnDefs

import "testing"

func TestParseIt(t *testing.T) {
	tests := []struct {
		lineType string
		line     string
		expected string
	}{
		{
			lineType: "SL1L2",
			line:     "V12.0.0 FCST  NA   120000    20050807_120000 20050807_120000 000000   20120409_113000 20120409_123000 FCST     None       Surface  TMP     NA        Z2      ADPSFC FULL    NEAREST     1           NA          NA         NA         NA    SL1L2      4511 59.12348 279.18203 16572.91166 14380.0603  77998.04015 220.05855",
			expected: "parsed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.lineType, func(t *testing.T) {
			result, err := ParseIt(tt.lineType, tt.line)
			if err != nil {
				t.Errorf("ParseIt() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("ParseIt() = %v, want %v", result, tt.expected)
			}
		})
	}
}
