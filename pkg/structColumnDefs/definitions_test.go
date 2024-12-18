package structColumnDefs

import (
	"encoding/json"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"parser/pkg/structColumnTypes"
)

func ReadJsonFromFile(filename string) ([]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var parsedDoc []interface{}
	err = decoder.Decode(&parsedDoc)
	if err != nil {
		return nil, err
	}
	return parsedDoc, nil
}

func TestParseIt(t *testing.T) {
	headerLine := "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV  OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE"
	dataLine := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LAND_L0 NEAREST     1           NA          NA         NA         NA    VAL1L2    4114    0.022881     -0.055846      -0.23975       0.11316       1.40894     2.39774     6.07755      1.35071    2.1488    4114           12.11241   65.18733  6744.28012"
	dataLine2 := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"
	dataLine3 := "V12.0.0 FCST  NA   180000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"

	fileType := "STAT"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, fileType, doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine2, fileType, doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine3, fileType, doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if doc == nil {
		t.Fatalf("Expected parsed document, got nil")
	}
	err = WriteJsonToFile(doc, "test_output.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// read the file back in
	parsedDoc, err := ReadJsonFromFile("test_output.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.NotNil(t, parsedDoc)
	assert.Equal(t, 2, len(parsedDoc), "expected 2 but got %d", len(parsedDoc)) // two top level elements
	doc0 := doc["V12.0.0:FCST:NA:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:NA:Z10:ADPSFC:LAND_L0:NEAREST:1:NA:NA:NA:NA:VAL1L2"].(map[string]interface{})
	doc0Data := doc0["data"].(map[string]structColumnTypes.STAT_VAL1L2)
	doc0Data120000 := doc0Data["120000"]
	doc0DataTotal := doc0Data120000.Total
	parsedDoc0 := parsedDoc[0].(map[string]interface{})
	parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	// Don't know why the parsedDoc (what we read back in) came back as a float and not an int
	parsedDoc0DataTotal := int(parsedDoc0Data["120000"].(map[string]interface{})["total"].(float64))
	assert.Equal(t, doc0DataTotal, parsedDoc0DataTotal, "expected doc and parsedDoc to have equal values")
	// Add more assertions based on the expected structure of parsedDoc
}
