package buildHeaderLineTypeUtilities

import (
	"testing"
)

func TestGetLineType(t *testing.T) {
	tests := []struct {
		headerLine    string
		dataLine      string
		fileType      string
		wantType      string
		wantHeader    []string
		wantData      []string
		wantKey       string
		wantDescIndex int
		wantErr       bool
	}{
		{
			headerLine:    "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE",
			dataLine:      "V12.0.0 ECMWF NA   060000    20241031_000000 20241031_000000 000000   20241031_000000 20241031_000000 TMP       K          P1000    TMP       K         P1000   ANLYS  FULL    NEAREST     1           NA          NA         NA         NA    SAL1L2    1038240   0.84332      0.45998         6.19154        7.38574        7.9421     1.02924",
			fileType:      "STAT",
			wantType:      "SAL1L2",
			wantHeader:    []string{"V12.0.0", "ECMWF", "NA", "", "1730332800", "1730332800", "000000", "1730332800", "1730332800", "TMP", "K", "P1000", "TMP", "K", "P1000", "ANLYS", "FULL", "NEAREST", "1", "NA", "NA", "NA", "NA", "SAL1L2"},
			wantData:      []string{"1038240", "0.84332", "0.45998", "6.19154", "7.38574", "7.9421", "1.02924"},
			wantKey:       "060000",
			wantDescIndex: 2,
			wantErr:       false,
		},
		{
			headerLine:    "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE",
			dataLine:      "V12.0.0 ECMWF NA   060000    20241031_000000 20241031_000000 000000   20241031_000000 20241031_000000 TMP       K          P1000    TMP       K         P1000   ANLYS  NAO     NEAREST     1           NA          NA         NA         NA    SAL1L2      18444  -0.56214     -0.71048        12.15072       12.51109       12.28397    0.54932",
			fileType:      "STAT",
			wantType:      "SAL1L2",
			wantHeader:    []string{"V12.0.0", "ECMWF", "NA", "", "1730332800", "1730332800", "000000", "1730332800", "1730332800", "TMP", "K", "P1000", "TMP", "K", "P1000", "ANLYS", "NAO", "NEAREST", "1", "NA", "NA", "NA", "NA", "SAL1L2"},
			wantData:      []string{"18444", "-0.56214", "-0.71048", "12.15072", "12.51109", "12.28397", "0.54932"},
			wantKey:       "060000",
			wantDescIndex: 2,
			wantErr:       false,
		},
		{
			headerLine:    "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE",
			dataLine:      "V12.0.0 ECMWF NA   060000    20241031_000000 20241031_000000 000000   20241031_000000 20241031_000000 TMP       K          P850     TMP       K         P850    ANLYS  FULL    NEAREST     1           NA          NA         NA         NA    SAL1L2    1038240   0.82265      0.63798        11.1637        11.68235       12.24054    0.8205",
			fileType:      "STAT",
			wantType:      "SAL1L2",
			wantHeader:    []string{"V12.0.0", "ECMWF", "NA", "", "1730332800", "1730332800", "000000", "1730332800", "1730332800", "TMP", "K", "P850", "TMP", "K", "P850", "ANLYS", "FULL", "NEAREST", "1", "NA", "NA", "NA", "NA", "SAL1L2"},
			wantData:      []string{"1038240", "0.82265", "0.63798", "11.1637", "11.68235", "12.24054", "0.8205"},
			wantKey:       "060000",
			wantDescIndex: 2,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.headerLine, func(t *testing.T) {
			gotType, gotHeader, gotData, gotKey, gotDescIndex, err := GetLineType(tt.headerLine, tt.dataLine, tt.fileType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLineType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotType != tt.wantType {
				t.Errorf("GetLineType() gotType = %v, want %v", gotType, tt.wantType)
			}
			if !equal(gotHeader, tt.wantHeader) {
				t.Errorf("GetLineType() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !equal(gotData, tt.wantData) {
				t.Errorf("GetLineType() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetLineType() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotDescIndex != tt.wantDescIndex {
				t.Errorf("GetLineType() gotKey = %v, want %v", gotDescIndex, tt.wantDescIndex)
			}
		})
	}
}

func TestSplitColumnDefLine(t *testing.T) {
	tests := []struct {
		fileType   string
		fieldStr   string
		wantHeader []string
		wantData   []string
	}{
		{
			fileType:   "MODE",
			fieldStr:   "FCST_LEAD FCST_VALID_BEG FCST_VALID_END OBS_LEV LINE_TYPE",
			wantHeader: []string{"FCST_LEAD", "FCST_VALID_BEG", "FCST_VALID_END", "OBS_LEV"},
			wantData:   []string{"LINE_TYPE"},
		},
		{
			fileType:   "MTD",
			fieldStr:   "FCST_LEAD FCST_VALID_BEG FCST_VALID_END OBS_LEV LINE_TYPE",
			wantHeader: []string{"FCST_LEAD", "FCST_VALID_BEG", "FCST_VALID_END", "OBS_LEV"},
			wantData:   []string{"LINE_TYPE"},
		},
		{
			fileType:   "STAT",
			fieldStr:   "FCST_LEAD FCST_VALID_BEG FCST_VALID_END LINE_TYPE DATA_FIELD1 DATA_FIELD2",
			wantHeader: []string{"FCST_LEAD", "FCST_VALID_BEG", "FCST_VALID_END", "LINE_TYPE"},
			wantData:   []string{"DATA_FIELD1", "DATA_FIELD2"},
		},
		{
			fileType:   "STAT",
			fieldStr:   "FCST_LEAD FCST_VALID_BEG FCST_VALID_END LINE_TYPE",
			wantHeader: []string{"FCST_LEAD", "FCST_VALID_BEG", "FCST_VALID_END", "LINE_TYPE"},
			wantData:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.fileType, func(t *testing.T) {
			gotHeader, gotData := SplitColumnDefLine(tt.fileType, tt.fieldStr)
			if !equal(gotHeader, tt.wantHeader) {
				t.Errorf("SplitColumnDefLine() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !equal(gotData, tt.wantData) {
				t.Errorf("SplitColumnDefLine() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestGetKeyDataFieldsForLineType(t *testing.T) {
	tests := []struct {
		lineType string
		want     []string
	}{
		{
			lineType: "MODE",
			want:     []string{"FCST_LEAD", "FCST_LEV"},
		},
		{
			lineType: "MTD",
			want:     []string{"FCST_LEAD", "FCST_LEV"},
		},
		{
			lineType: "STAT",
			want:     []string{"FCST_LEAD"},
		},
		{
			lineType: "unknown",
			want:     []string{"FCST_LEAD"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.lineType, func(t *testing.T) {
			got := GetKeyDataFieldsForLineType(tt.lineType)
			if !equal(got, tt.want) {
				t.Errorf("GetKeyDataFieldsForLineType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindType(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    string
		wantErr bool
	}{
		{
			name:    "item1",
			data:    []byte("atoi(get_item) item1"),
			want:    "int",
			wantErr: false,
		},
		{
			name:    "item2",
			data:    []byte("atof(get_item) item2"),
			want:    "float64",
			wantErr: false,
		},
		{
			name:    "item3",
			data:    []byte("atos(get_item) item3"),
			want:    "string",
			wantErr: false,
		},
		{
			name:    "item4",
			data:    []byte("get_item item4"),
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindType(tt.name, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
