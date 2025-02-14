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

// dummy function to satisfy the function signature of getExternalDocForId
func getMissingExternalDocForId(id string) (map[string]interface{}, error) {
	// fmt.Println("getExternalDocForId called with id:", id)
	// Put your own code here in this method but always return this exact error if the document is not found
	return nil, fmt.Errorf("%s: %s", structColumnTypes.DOC_NOT_FOUND, id)
}

// dummy function to satisfy the function signature of getExternalDocForId
func getExistingExternalDocForId(id string) (map[string]interface{}, error) {
	// fmt.Println("getExternalDocForId called with id:", id)
	fileLineType := "STAT_VAL1L2"
	metaDataMap := map[string]interface{}{"id": "MET:DD:MET:V12.0.0:FCST:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LAND_L0:NEAREST:1:VAL1L2", "subset": "MET", "type": "DD", "subtype": "MET"}
	headerData := []string{"V12.0.0", "FCST", "", "", "1333972800", "1333972800", "000000", "1333971000", "1333974600", "UGRD_VGRD", "m/s", "Z10", "UGRD_VGRD", "", "Z10", "ADPSFC", "LAND_L0", "NEAREST", "1", "", "", "", "", "VAL1L2"}
	dataData := []string{"4114", "0.022881", "-0.055846", "-0.23975", "0.11316", "1.40894", "2.39774", "6.07755", "1.35071", "2.1488", "4114", "12.11241", "65.18733", "6744.28012"}
	dataKey := "120000"
	var _err error

	doc, _err := structColumnTypes.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
	if _err != nil {
		return nil, _err
	}
	if _err != nil {
		return doc.(map[string]interface{}), _err
	}

	// Put your own code here in this method but always return this exact error if the document is not found
	return doc.(map[string]interface{}), nil
}

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

func TestGetMissingExternalDocForId(t *testing.T) {
	headerLine := "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV  OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE"
	dataLine := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LAND_L0 NEAREST     1           NA          NA         NA         NA    VAL1L2    4114    0.022881     -0.055846      -0.23975       0.11316       1.40894     2.39774     6.07755      1.35071    2.1488    4114           12.11241   65.18733  6744.28012"
	fName := "grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_900000L_20241104_180000V.stat"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetExistingExternalDocForId(t *testing.T) {
	headerLine := "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV  OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE"
	dataLine := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LAND_L0 NEAREST     1           NA          NA         NA         NA    VAL1L2    4114    0.022881     -0.055846      -0.23975       0.11316       1.40894     2.39774     6.07755      1.35071    2.1488    4114           12.11241   65.18733  6744.28012"
	fName := "grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_900000L_20241104_180000V.stat"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, &doc, fName, getExistingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestParseVAL1L2(t *testing.T) {
	headerLine := "VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV  OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE"
	dataLine := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LAND_L0 NEAREST     1           NA          NA         NA         NA    VAL1L2    4114    0.022881     -0.055846      -0.23975       0.11316       1.40894     2.39774     6.07755      1.35071    2.1488    4114           12.11241   65.18733  6744.28012"
	dataLine2 := "V12.0.0 FCST  NA   120000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"
	dataLine3 := "V12.0.0 FCST  NA   180000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"
	dataLine4 := "V12.0.0 FCST  this_is_a_long_description_field   180000    20120409_120000 20120409_120000 000000   20120409_113000 20120409_123000 UGRD_VGRD m/s        Z10      UGRD_VGRD NA        Z10      ADPSFC LMV     NEAREST     1           NA          NA         NA         NA    VAL1L2     393   -0.32297       0.32197       -0.79039       0.14006       1.34214     1.86519     3.95307      1.23297    1.78245    393           26.10387   54.98572  4500.31836"

	fName := "grid_stat_ECMWF_TMP_vs_ANLYS_TMP_P1000_anom_360000L_20241031_000000V.stat"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine2, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine3, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine4, &doc, fName, getMissingExternalDocForId)
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

/*
This test tests a data field dataKey.
*/
func TestParseMODE_OBJ(t *testing.T) {
	headerLine := "VERSION MODEL N_VALID GRID_RES DESC FCST_LEAD FCST_VALID FCST_ACCUM OBS_LEAD OBS_VALID OBS_ACCUM FCST_RAD FCST_THR OBS_RAD OBS_THR FCST_VAR FCST_UNITS FCST_LEV OBS_VAR OBS_UNITS OBS_LEV OBTYPE OBJECT_ID OBJECT_CAT CENTROID_X CENTROID_Y CENTROID_LAT CENTROID_LON AXIS_ANG LENGTH WIDTH AREA AREA_THRESH CURVATURE CURVATURE_X CURVATURE_Y COMPLEXITY INTENSITY_10 INTENSITY_25 INTENSITY_50 INTENSITY_75 INTENSITY_90 INTENSITY_95 INTENSITY_SUM CENTROID_DIST BOUNDARY_DIST CONVEX_HULL_DIST ANGLE_DIFF ASPECT_DIFF AREA_RATIO INTERSECTION_AREA UNION_AREA SYMMETRIC_DIFF INTERSECTION_OVER_AREA CURVATURE_RATIO COMPLEXITY_RATIO PERCENTILE_INTENSITY_RATIO INTEREST"
	dataLine := "V11.1.0 HRRR_OPS 656523 3 E_CONUS 080000 20241201_190000 000000 000000 20241201_185837 000000 1 >=35 1 >=35 REFC dB L0 REFC dB R1 MRMS F001 CF000 1179.05714 847.2 46.60621 -86.59767 -42.90968 9.89291 4.18823 20 18 1627.71082 1838.91188 1172.93246 0.32203 36.85 39.375 40.59375 42.5 44.95 45.0625 810.3125 NA NA NA NA NA NA NA NA NA NA NA NA NA NA"
	dataLine2 := "V11.1.0 HRRR_OPS 656523 3 E_CONUS 080000 20241201_190000 000000 000000 20241201_185837 000000 1 >=35 1 >=35 REFC dB L0 REFC dB R1 MRMS F002 CF000 1460.47059 790.02941 43.79098 -76.43034 37.76469 8.59583 4.38697 20 20 1907.78256 1981.5673 1571.67157 0.24528 37.36875 40.04688 41.125 42.20312 43.5875 43.81562 815.5 NA NA NA NA NA NA NA NA NA NA NA NA NA NA"
	dataLine3 := "V11.1.0 HRRR_OPS 656523 3 E_CONUS 080000 20241201_190000 000000 000000 20241201_185837 000000 1 >=35 1 >=35 REFC dB L0 REFC dB R1 MRMS F003 CF000 1414.91176 532 37.24534 -79.91852 25.39649 11.60689 2.80129 17 16 1941.15256 1786.31378 1691.79385 0.38182 36.2625 37 37.25 38.25 38.9 39.15 637.4375 NA NA NA NA NA NA NA NA NA NA NA NA NA NA"

	//fileType := "MODE_OBJ"
	fName := "mode_python_mixed_300000L_20120410_180000V_060000A_cts.txt"
	var doc map[string]interface{}
	doc, err := ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine2, &doc, fName, getMissingExternalDocForId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	doc, err = ParseLine(headerLine, dataLine3, &doc, fName, getMissingExternalDocForId)
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
			ext := filepath.Ext(fName)
			filePathParts := strings.Split(ext, ".")
			fileType := strings.ToUpper(filePathParts[1])
			if fileType == "SWP" {
				// skip the swp files - might be editing a file and don't want to parse the .swp file
				continue
			}
			if fileType == "DS_STORE" {
				// skip the .DS_Store files
				continue
			}
			if fileType == "txt" {
				// .txt filepaths are different e.g.
				// ./textfiles/point_stat_GRIB2_SREF_GDAS_150000L_20120409_120000V_sl1l2.txt
				// ./textfiles/mode_MASK_POLY_300000L_20120410_180000V_060000A_cts.txt
				// ./textfiles/mode_python_120000L_20050807_120000V_120000A_obj.txt
				fileType = strings.Split(filePathParts[0], "_")[len(filePathParts[0])-1]
			}
			doc, err = ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
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
					doc, err = ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
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
