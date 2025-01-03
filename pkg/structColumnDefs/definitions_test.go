package structColumnDefs

/*
For profiling -
brew install --cask approf
cd .../pkg/structColumnDefs
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

*/
import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"parser/pkg/structColumnTypes"
)

func ReadJsonFromGzipFile(filename string) ([]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Print(err)
	}
	defer gz.Close()
	decoder := json.NewDecoder(gz)
	var parsedDoc []interface{}
	err = decoder.Decode(&parsedDoc)
	if err != nil {
		return nil, err
	}
	return parsedDoc, nil
}

func TestParseVAL1L2(t *testing.T) {
	headerLine := "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV  OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE"
	dataLine := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LAND_L0 NEAREST     1           NA          NA         NA         NA    VAL1L2    4114    0.022881     -0.055846      -0.23975       0.11316       1.40894     2.39774     6.07755      1.35071    2.1488    4114           12.11241   65.18733  6744.28012"
	dataLine2 := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"
	dataLine3 := "V12.0.0 FCST  NA   180000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"
	dataLine4 := "V12.0.0 FCST  this_is_a_long_description_field   180000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"

	fileType := "STAT"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, fileType, &doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine2, fileType, &doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine3, fileType, &doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine4, fileType, &doc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if doc == nil {
		t.Fatalf("Expected parsed document, got nil")
	}
	err = WriteJsonToGzipFile(doc, "/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// read the file back in
	parsedDoc, err := ReadJsonFromGzipFile("/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.NotNil(t, parsedDoc)
	assert.Equal(t, 3, len(parsedDoc), "expected 3 but got %d", len(parsedDoc)) // two top level elements
	doc0 := doc["MET:DD:MET:V12.0.0:FCST:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LAND_L0:NEAREST:1:VAL1L2"].(map[string]interface{})
	doc2 := doc["MET:DD:MET:V12.0.0:FCST:this_is_a_:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LMV:NEAREST:1:VAL1L2"].(map[string]interface{})
	doc0Data := doc0["data"].(map[string]structColumnTypes.STAT_VAL1L2)
	doc2Data := doc2["data"].(map[string]structColumnTypes.STAT_VAL1L2)
	doc0Data120000 := doc0Data["120000"]
	doc2Data180000 := doc2Data["180000"]
	doc0DataTotal := doc0Data120000.Total
	doc2DataTotal := doc2Data180000.Total
	parsedDoc0 := parsedDoc[0].(map[string]interface{})
	parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	parsedDoc2 := parsedDoc[2].(map[string]interface{})
	parsedDoc2Data := parsedDoc2["data"].(map[string]interface{})
	// Don't know why the parsedDoc (what we read back in) came back as a float and not an int
	parsedDoc0DataTotal := int(parsedDoc0Data["120000"].(map[string]interface{})["total"].(float64))
	parsedDoc2DataTotal := int(parsedDoc2Data["180000"].(map[string]interface{})["total"].(float64))
	assert.Equal(t, doc0DataTotal, parsedDoc0DataTotal, "expected doc and parsedDoc to have equal values")
	assert.Equal(t, doc2DataTotal, parsedDoc2DataTotal, "expected doc and parsedDoc to have equal values")
	// Add more assertions based on the expected structure of parsedDoc
}

func TestParseRegressionSuite(t *testing.T) {
	var doc map[string]interface{}
	var err error
	path, err := os.Getwd()
	if err != nil {
		t.Fatal("error getting working directory:", err)
	}
	directory := path + "/../../test_data/statfiles/" // The statfiles directory

	files, err := os.Open(directory) // open the directory to read files in the directory
	if err != nil {
		t.Fatal("error opening directory:", err) // print error if directory is not opened
	}
	defer files.Close()                 // close the directory opened
	fileInfos, err := files.Readdir(-1) // read the files from the directory
	if err != nil {
		t.Fatal("error reading directory:", err) // if directory is not read properly print error message
	}
	for _, fileInfos := range fileInfos {
		file, err := os.Open(directory + fileInfos.Name()) // open the file
		if err != nil {
			t.Fatal("error opening file", err)
		}
		defer file.Close()
		fName := fileInfos.Name()
		// fmt.Println("Parsing file:", fName)
		ext := filepath.Ext(fName)
		fileType := strings.ToUpper(strings.Split(ext, ".")[1])
		if fileType == "SWP" {
			// skip the swp files - might be editing a file and don't want to parse the .swp file
			continue
		}
		rawData, err := io.ReadAll(file)
		if err != nil {
			t.Fatal("error reading file", err)
		}
		lines := strings.Split(string(rawData), "\n")
		headerLine := lines[0]
		for line := range lines {
			if line == 0 {
				continue
			}
			dataLine := lines[line]
			doc, err = ParseLine(headerLine, dataLine, fileType, &doc)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		}
	}
	if doc == nil {
		t.Fatalf("Expected parsed document, got nil")
	}
	err = WriteJsonToGzipFile(doc, "/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// read the file back in
	parsedDoc, err := ReadJsonFromGzipFile("/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.NotNil(t, parsedDoc)
	// assert.Equal(t, 2, len(parsedDoc), "expected 2 but got %d", len(parsedDoc)) // two top level elements
	// doc0 := doc["V12.0.0:FCST:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LAND_L0:NEAREST:1:VAL1L2"].(map[string]interface{})
	// doc0Data := doc0["data"].(map[string]structColumnTypes.STAT_VAL1L2)
	// doc0Data120000 := doc0Data["120000"]
	// doc0DataTotal := doc0Data120000.Total
	// parsedDoc0 := parsedDoc[0].(map[string]interface{})
	// parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	// // Don't know why the parsedDoc (what we read back in) came back as a float and not an int
	// parsedDoc0DataTotal := int(parsedDoc0Data["120000"].(map[string]interface{})["total"].(float64))
	// assert.Equal(t, doc0DataTotal, parsedDoc0DataTotal, "expected doc and parsedDoc to have equal values")
	// Add more assertions based on the expected structure of parsedDoc
}

func TestParseG2G_v12_Suite(t *testing.T) {
	var doc map[string]interface{}
	var err error
	var path string

	path, err = os.Getwd()
	if err != nil {
		t.Fatal("error getting working directory:", err)
	}
	directory := path + "/../../test_data/G2G_v12" // The G2G_v12 top level directory

	err = filepath.Walk(directory,
		func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				// this is a directory
				fmt.Println(path, fileInfo.Size())
			} else {
				// this is a file so process it
				file, err := os.Open(path) // open the file
				if err != nil {
					t.Fatal("error opening file", err)
				}
				defer file.Close()
				fName := fileInfo.Name()
				// uncomment the following for debugging
				// fmt.Println("Parsing file:", fName)
				ext := filepath.Ext(fName)
				fileType := strings.ToUpper(strings.Split(ext, ".")[1])
				if fileType == "SWP" {
					// skip the swp files - might be editing a file and don't want to parse the .swp file
					return nil
				}
				rawData, err := io.ReadAll(file)
				if err != nil {
					t.Fatal("error reading file", err)
				}
				lines := strings.Split(string(rawData), "\n")
				headerLine := lines[0]
				for line := range lines {
					if line == 0 {
						continue
					}
					dataLine := lines[line]
					doc, err = ParseLine(headerLine, dataLine, fileType, &doc)
					if err != nil {
						if strings.Contains(err.Error(), "UNPARSABLE_LINE") {
							// skip the line
							fmt.Printf("Skipping line: %s because it isn't parsable from file %s\n", dataLine, fName)
						} else {
							t.Fatalf("Expected no error, got %v", err)
						}
					}
				}
			}
			return nil
		})
	if err != nil {
		t.Fatal(err)
	}

	if doc == nil {
		t.Fatalf("Expected parsed document, got nil")
	}
	err = WriteJsonToGzipFile(doc, "/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// read the file back in
	parsedDoc, err := ReadJsonFromGzipFile("/tmp/test_output.json.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.NotNil(t, parsedDoc)
	// assert.Equal(t, 2, len(parsedDoc), "expected 2 but got %d", len(parsedDoc)) // two top level elements
	// doc0 := doc["V12.0.0:FCST:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LAND_L0:NEAREST:1:VAL1L2"].(map[string]interface{})
	// doc0Data := doc0["data"].(map[string]structColumnTypes.STAT_VAL1L2)
	// doc0Data120000 := doc0Data["120000"]
	// doc0DataTotal := doc0Data120000.Total
	// parsedDoc0 := parsedDoc[0].(map[string]interface{})
	// parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	// // Don't know why the parsedDoc (what we read back in) came back as a float and not an int
	// parsedDoc0DataTotal := int(parsedDoc0Data["120000"].(map[string]interface{})["total"].(float64))
	// assert.Equal(t, doc0DataTotal, parsedDoc0DataTotal, "expected doc and parsedDoc to have equal values")
	// Add more assertions based on the expected structure of parsedDoc
}
