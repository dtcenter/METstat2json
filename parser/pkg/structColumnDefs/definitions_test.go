package structColumnDefs

import (
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name      string
		columnDef ColumnDef
		line      string
		expected  interface{}
	}{
		{
			name: "Test VL1L2Doc",
			columnDef: ColumnDef{
				Name: "VL1L2",
				doc:  VL1L2Doc{},
			},
			line: "V11.1.0 GFS   NA   120000    20240203_120000 20240203_120000 000000   20240203_120000 20240203_120000 TMP       K          P850     TMP       K         P850    ANLYS  FULL    NEAREST     1           NA          NA         NA         NA    SAL1L2    10512   1.3349     1.33074      20.76681     20.78507      21.03788    0.35771"

			expected: VL1L2Doc{
				HeaderFields: statHeaderL1{
					FcstLev:      "1000",
					FcstUnits:    "hPa",
					FcstValidBeg: 20230101,
					FcstValidEnd: 20230102,
					FcstVar:      "TMP",
					ID:           "1234",
					InterpMthd:   "BILIN",
					InterpPnts:   10,
					LineType:     "SL1L2",
					Model:        "GFS",
					ObsLev:       "1000",
					ObsUnits:     "hPa",
					ObsValidBeg:  20230101,
					ObsValidEnd:  20230102,
					ObsVar:       "TMP",
					Obtype:       "1234",
					Version:      "1.0",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pb := &parserBuilder{
				columnDef: tt.columnDef,
			}
			parser := pb.Build(tt.line)
			doc := reflect.ValueOf(parser).Elem().FieldByName("Columns").FieldByName("doc").Interface()
			if !reflect.DeepEqual(doc, tt.expected) {
				t.Errorf("Build() = %v, want %v", doc, tt.expected)
			}
		})
	}
}
