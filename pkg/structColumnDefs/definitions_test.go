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
	"parser/pkg/structColumnTypes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	// Put your own code here in this method but always return this exact error if the document is not found
	return doc.(map[string]interface{}), nil
}

func ReadJsonFromGzipFile(filename string) (map[string]interface{}, error) {
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
	var result []interface{}
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	parsedDoc := make(map[string]interface{})
	for _, v := range result {
		elem := v.(map[string]interface{})
		parsedDoc[elem["id"].(string)] = elem
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
	/* two of these data lines (dataLine and dataLine2) are the same. Only one of them should show up in the doc. The dataLine3 and dataLine4 only differ by their desc.*/
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
	doc0DataTotal := doc0Data120000.TOTAL
	doc2DataTotal := doc2Data180000.TOTAL

	parsedDoc0 := parsedDoc["MET:DD:MET:V12.0.0:FCST:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LAND_L0:NEAREST:1:VAL1L2"].(map[string]interface{})
	parsedDoc2 := parsedDoc["MET:DD:MET:V12.0.0:FCST:this_is_a_:1333972800:1333972800:000000:1333971000:1333974600:UGRD_VGRD:m/s:Z10:UGRD_VGRD:Z10:ADPSFC:LMV:NEAREST:1:VAL1L2"].(map[string]interface{})
	parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	parsedDoc2Data := parsedDoc2["data"].(map[string]interface{})
	parsedDoc0Data120000 := parsedDoc0Data["120000"].(map[string]interface{})
	parsedDoc2Data180000 := parsedDoc2Data["180000"].(map[string]interface{})
	parsedDoc0DataTotal := int(parsedDoc0Data120000["total"].(float64))
	parsedDoc2DataTotal := int(parsedDoc2Data180000["total"].(float64))

	assert.Equal(t, doc0DataTotal, parsedDoc0DataTotal, "expected doc and parsedDoc to have equal values")
	assert.Equal(t, doc2DataTotal, parsedDoc2DataTotal, "expected doc and parsedDoc to have equal values")
	// Add more assertions based on the expected structure of parsedDoc
}

/*
This test tests a data field dataKey.
*/
func TestParseMODE_OBJ(t *testing.T) {
	headerLine := "VERSION MODEL N_VALID GRID_RES DESC FCST_LEAD FCST_VALID      FCST_ACCUM OBS_LEAD OBS_VALID       OBS_ACCUM FCST_RAD FCST_THR OBS_RAD OBS_THR FCST_VAR FCST_UNITS FCST_LEV OBS_VAR OBS_UNITS OBS_LEV OBTYPE FIELD  TOTAL FY_OY FY_ON FN_OY FN_ON BASER   FMEAN    ACC     FBIAS  PODY       PODN    POFD     FAR     CSI        GSS       HK        HSS       ODDS      LODDS   ORSS     EDS      SEDS     EDI      SEDI     BAGSS"
	dataLine := "V12.0.0 FCST  26026   9        NA   300000    20120410_180000 060000     120000   20050807_120000 120000    2        >=5.0    2       >=5.0   APCP_06  kg/m^2     A6       OBS     None      Surface STAGE4    RAW 26026    47  1356  5898 18725 0.22843 0.053908 0.72128 0.236  0.0079058  0.93247 0.067527 0.9665  0.0064375  -0.039178 -0.059621 -0.08155  0.11004   -2.2069 -0.80173 -0.53249 -0.3039  -0.28465 -0.28988 -0.11236"
	dataLine2 := "V12.0.0 FCST  26026   9        this_is_a_long_description   300000    20120410_180000 060000     120000   20050807_120000 120000    2        >=5.0    2       >=5.0   APCP_06  kg/m^2     A6       OBS     None      Surface STAGE4 OBJECT 26026     4  1315  6322 18385 0.24306 0.05068  0.70656 0.2085 0.00063231 0.93325 0.066751 0.99697 0.00052349 -0.043249 -0.066119 -0.090409 0.0088459 -4.7278 -0.98246 -0.67783 -0.49927 -0.46256 -0.46613 -0.13686"

	// fileType := "MODE_OBJ"
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
	// write the doc to a file
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
	assert.Equal(t, 2, len(parsedDoc), "expected 3 but got %d", len(parsedDoc)) // two top level elements
	doc0 := doc["MET:DD:MET:V12.0.0:FCST:26026:9:20120410_180000:060000:120000:20050807_120000:120000:2:>=5.0:2:>=5.0:APCP_06:kg/m^2:A6:OBS:None:Surface:STAGE4"].(map[string]interface{})
	doc2 := doc["MET:DD:MET:V12.0.0:FCST:26026:9:this_is_a_:20120410_180000:060000:120000:20050807_120000:120000:2:>=5.0:2:>=5.0:APCP_06:kg/m^2:A6:OBS:None:Surface:STAGE4"].(map[string]interface{})
	doc0Data := doc0["data"].(map[string]structColumnTypes.MODE_CTS)
	doc2Data := doc2["data"].(map[string]structColumnTypes.MODE_CTS)
	doc0Data120000 := doc0Data["300000"]
	doc2Data180000 := doc2Data["300000"]
	doc0DataTotal := doc0Data120000.TOTAL
	doc2DataTotal := doc2Data180000.TOTAL

	parsedDoc0 := parsedDoc["MET:DD:MET:V12.0.0:FCST:26026:9:20120410_180000:060000:120000:20050807_120000:120000:2:>=5.0:2:>=5.0:APCP_06:kg/m^2:A6:OBS:None:Surface:STAGE4"].(map[string]interface{})
	parsedDoc2 := parsedDoc["MET:DD:MET:V12.0.0:FCST:26026:9:this_is_a_:20120410_180000:060000:120000:20050807_120000:120000:2:>=5.0:2:>=5.0:APCP_06:kg/m^2:A6:OBS:None:Surface:STAGE4"].(map[string]interface{})
	parsedDoc0Data := parsedDoc0["data"].(map[string]interface{})
	parsedDoc2Data := parsedDoc2["data"].(map[string]interface{})
	parsedDoc0Data120000 := parsedDoc0Data["300000"].(map[string]interface{})
	parsedDoc2Data180000 := parsedDoc2Data["300000"].(map[string]interface{})
	parsedDoc0DataTotal := int(parsedDoc0Data120000["total"].(float64))
	parsedDoc2DataTotal := int(parsedDoc2Data180000["total"].(float64))

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
			if line == 0 || lines[line] == "" {
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
	// add other test assertions here
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
				if strings.Contains(fileInfo.Name(), ".swp") {
					// skip the swp files - might be editing a file and don't want to parse the .swp file
					return nil
				}
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
					if line == 0 || lines[line] == "" {
						continue
					}
					dataLine := lines[line]
					doc, err = ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
					if err != nil {
						if strings.Contains(err.Error(), "UNPARSABLE_LINE") {
							// skip the line
							// I know that the file ./G2G_v12/20241101-18z/grid_stat/grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
							// has a truncated line at the end. I don't want to fail the test because of it.
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
	// add other test assertions here
}
