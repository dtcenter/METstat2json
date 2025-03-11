package buildHeaderLineTypeUtilities

/*
This package contains the utilities for building the header and line type for the data files. The data files are
MET output files that contain a header section and a data section. The header section contains the header fields that
are used to identify the document. The data section contains the data fields that are used to populate the document.
This package is separate from the structColumnTypes package because the structColumnTypes package is automatically
generated from the buildHeaderLineTypes.go program and there is a desire to avoid a circular dependency.
This package defines a VxMetaData struct that is used to store the metadata for the mapped documents.
The metadata is used to uniquely identify each document and is used to merge documents with the same metadata.

This package also defines the DataKeyMap that is used to determine the key data fields for a given line type.
The key data fields are used to merge documents with the same header field values excluding the key data fields.
Other utilities exist to convert the date fields to epochs, to get the line type of the data line, to get the key
data fields for a given line type, and to find the data type of a given field in addition to some other utility functions.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HeaderFields struct {
	Header         string
	Version        string
	LineType       string
	FileType       string
	SeparatorField string
}

var MetHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"

// vxMetadata struct definition
type VxMetadata struct {
	ID      string `json:"id"`
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

// data key definitions
var DataKeyMap = map[string][]string{
	"STAT_CNT":       {"FCST_LEAD"},
	"STAT_CTC":       {"FCST_LEAD"},
	"STAT_CTS":       {"FCST_LEAD"},
	"STAT_FHO":       {"FCST_LEAD"},
	"STAT_ISC":       {"FCST_LEAD"},
	"STAT_MCTC":      {"FCST_LEAD"},
	"STAT_MCTS":      {"FCST_LEAD"},
	"STAT_MPR":       {"FCST_LEAD"},
	"STAT_SEEPS":     {"FCST_LEAD"},
	"STAT_SEEPS_MPR": {"FCST_LEAD"},
	"STAT_NBRCNT":    {"FCST_LEAD"},
	"STAT_NBRCTC":    {"FCST_LEAD"},
	"STAT_NBRCTS":    {"FCST_LEAD"},
	"STAT_GRAD":      {"FCST_LEAD"},
	"STAT_DMAP":      {"FCST_LEAD"},
	"STAT_ORANK":     {"FCST_LEAD"},
	"STAT_PCT":       {"FCST_LEAD"},
	"STAT_PJC":       {"FCST_LEAD"},
	"STAT_PRC":       {"FCST_LEAD"},
	"STAT_PSTD":      {"FCST_LEAD"},
	"STAT_ECLV":      {"FCST_LEAD"},
	"STAT_ECNT":      {"FCST_LEAD"},
	"STAT_RPS":       {"FCST_LEAD"},
	"STAT_RHIST":     {"FCST_LEAD"},
	"STAT_PHIST":     {"FCST_LEAD"},
	"STAT_RELP":      {"FCST_LEAD"},
	"STAT_SAL1L2":    {"FCST_LEAD"},
	"STAT_SL1L2":     {"FCST_LEAD"},
	"STAT_SSVAR":     {"FCST_LEAD"},
	"STAT_VAL1L2":    {"FCST_LEAD"},
	"STAT_VL1L2":     {"FCST_LEAD"},
	"STAT_VCNT":      {"FCST_LEAD"},
	"STAT_GENMPR":    {"FCST_LEAD"},
	"STAT_SSIDX":     {"FCST_LEAD"},
	"MODE_OBJ":       {"FCST_LEAD", "OBJECT_ID"},
	"MODE_CTS":       {"FCST_LEAD"},
	"MTD_2DSINGLE":   {"OBJECT_ID", "TIME_INDEX"},
	"MTD_3DSINGLE":   {"OBJECT_ID"},
	"MTD_3DPAIR":     {"OBJECT_ID"},
	"TCST_TCMPR":     {"LEAD"},
	"TCST_TCDIAG":    {"LEAD"},
	"TCST_PROBRIRW":  {"LEAD"},
}

var DateFieldNames = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}

func dateToEpoch(date string) string {
	// convert the date to an epoch from YYYYMMDD_HHMMSS
	// the date is in the format YYYYMMDD_HHMMSS - time zone defaults to UTC
	theTime, _ := time.Parse(`20060102_150405`, date)
	return strconv.FormatInt(theTime.Unix(), 10)
}

func getLeadFromInitValid(data []string, dataFieldIndex int) string {
	initTime, _ := time.Parse("20060102_150405", data[dataFieldIndex-1])
	validTime, _ := time.Parse("20060102_150405", data[dataFieldIndex+1])
	lead := validTime.Sub(initTime)
	return strconv.FormatInt(int64(lead.Hours()), 10)
}

/*
	    Returns the lineType of the data line, and the headerFields and dataFields of the data line as []string.
			There are some assumptions here for mode and mtd files. The assumptions are that we can derive the fileType
			from the fileName using this logic...
			If the base fileName is a mode file (i.e. starts with "mode_") and ends in either
			obj.txt or merge.txt then it can be treated as an obj type file and the line
	        separator is OBJECT_ID. If the base fileName is a mode file and ends in cts.txt
			then it can be treated as a cts type file and the line separator is FIELD.
			If the base fileName is an mtd file (i.e. starts with "mtd_") then it can be treated
			as an mtd type file and the line separator is also OBJECT_ID.
			example fileNames...
			// ./textfiles/point_stat_GRIB2_SREF_GDAS_150000L_20120409_120000V_sl1l2.txt
			// ./textfiles/mode_MASK_POLY_300000L_20120410_180000V_060000A_cts.txt
			// ./textfiles/mode_python_120000L_20050807_120000V_120000A_obj.txt
			// ./textfiles/mtd_PYTHON_20050807_120000V_2d.txt
			// ./textfiles/mtd_PYTHON_20050807_120000V_3d_pair_cluster.txt
			// ./textfiles/mtd_PYTHON_20050807_120000V_3d_pair_simple.txt
			// ./textfiles/mtd_PYTHON_20050807_120000V_3d_single_cluster.txt
			// ./textfiles/mtd_PYTHON_20050807_120000V_3d_single_simple.txt

		For stat files the header and data section are delineated by "LINE_TYPE" which is the last field of the header section of the header line.
		Since this function has to split the line into a header and a data section those fields are returned as well.
		The dataKey is the concatenation of the dataKeyFields. The desc_index is the index of the DESC field in the headerFields.
		The dataKeyFields are used to merge documents. A dataKey field may be derived from either header fields or data fields.
		If the dataKey is derived from header fields they will be removed from the headerFields and the headerData.
		If the dataKey is derived from data fields they will not be removed from the data element.
		Any headerData is converted to epochs if the field is a dateField.

		NOTE: Often a stat file will not have a header line that contains data fields, i.e. it will end with the LINE_TYPE.
		This is when the data file has different types of stat data in the same file.
		In this case the dataFields cannot be determined and if the dataKeyMap specifies a dataKeyField that is not in the
		headerFields then the dataKey will be "" and the entire header will be used to generate the id.
*/
func GetLineType(headerLine string, dataLine string, fileName string) (string, []string, []string, string, int, error) {
	// make sure we have the basename here
	fileName = filepath.Base(fileName)
	desc_index := -1
	allHeaderFields := strings.Fields(headerLine)
	// get the data fields from the data line
	allData := strings.Fields(dataLine)

	// different fileTypes have different separator fields
	separatorField := "LINE_TYPE"
	// lineType depends on the fileName
	// If the base fileName is a mode file (i.e. starts with "mode_") and ends in either
	// obj.txt or merge.txt then it can be treated as an obj type file and the line
	// separator is OBJECT_ID. If the base fileName is a mode file and ends in cts.txt
	// then it can be treated as a cts type file and the line separator is FIELD.
	// If the base fileName is an mtd file (i.e. starts with "mtd_") then it can be treated
	// as an mtd type file and the line separator is also OBJECT_ID.
	var fileLineType string
	if strings.Contains(fileName, "stat") {
		separatorField = "LINE_TYPE"
	} else if strings.Contains(fileName, "mode") {
		if strings.Contains(fileName, "obj") || strings.Contains(fileName, "merge") {
			separatorField = "OBJECT_ID"
			fileLineType = "MODE_OBJ"
		} else if strings.Contains(fileName, "cts") {
			separatorField = "FIELD"
			fileLineType = "MODE_CTS"
		}
	} else if strings.Contains(fileName, "mtd") {
		separatorField = "OBJECT_ID"
		if strings.Contains(fileName, "2d") {
			fileLineType = "MTD_2DSINGLE"
		} else if strings.Contains(fileName, "3d_single_cluster") {
			fileLineType = "MTD_3DSINGLE"
		} else if strings.Contains(fileName, "3d_single_simple") {
			fileLineType = "MTD_3DSINGLE"
		} else if strings.Contains(fileName, "3d_pair_cluster") {
			fileLineType = "MTD_3DPAIR"
		} else if strings.Contains(fileName, "3d_pair_simple") {
			fileLineType = "MTD_3DPAIR"
		}
	} else {
		// check for differently named files?
		// there appear to be some files that aren't named like the others
		// Check the header line to see if it matches a line in the column defs file
		// if it does then we can use the line type from the column defs file
		columnDefHeaderFields, err := getLineTypeFromColumnDefsFile(headerLine)
		if err == nil {
			fileLineType = columnDefHeaderFields.FileType + "_" + columnDefHeaderFields.LineType
			separatorField = columnDefHeaderFields.SeparatorField
		}
	}
	headerStringFields, _ := SplitColumnDefLine(fileLineType, headerLine)
	dataStartIndex := len(headerStringFields)
	dataData := allData[dataStartIndex:]
	lineTypeIndex := len(headerStringFields) - 1
	if lineTypeIndex > len(allData) {
		return "", nil, nil, "", desc_index, fmt.Errorf("UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line")
	}
	// now we know the lineType for  files.
	if separatorField == "LINE_TYPE" {
		if strings.Contains(fileName, "tcst") {
			fileLineType = "TCST" + "_" + allData[lineTypeIndex]
		} else {
			fileLineType = "STAT" + "_" + allData[lineTypeIndex]
		}
	}
	// get the desc_index
	for i, h := range headerStringFields {
		if strings.ToUpper(h) == "DESC" {
			desc_index = i
			break
		}
	}
	// have to remove the dataKeyFields from the headerFields and the headerData (dataData AND dataFields won't change)
	headerData := []string{}
	dataKeyFields := []string{}
	// look for dataKey fields in allData - remove the dataKeyFields from the headerData if the key is in the header
	for fIndex, field := range allHeaderFields {
		isDataKey := false
		for _, dk := range DataKeyMap[fileLineType] {
			if field == dk {
				isDataKey = true
				if (dk == "LEAD" || dk == "FCST_LEAD") && allData[fIndex] == "NA" {
					// This is a special case.
					// if the datakey is LEAD or FCST_LEAD and the lead is NA then we
					// have to get the lead from the init and valid fields
					if allData[fIndex] == "NA" {
						allData[fIndex] = getLeadFromInitValid(allData, fIndex)
					}
				}
				if allData[fIndex] != "NA" {
					// the dataKey is the value that belongs to the associated field
					// not the field itself. Datakeys deliniate data sections
					// based on their values
					dataKeyFields = append(dataKeyFields, allData[fIndex])
				}
				break
			}
		}

		isHeaderField := fIndex <= lineTypeIndex
		// iterate through the header fields and
		// if the field is a header field and a dataKeyField then blank it out
		// convert header DATE fields that are not dataKeys to epochs
		if isHeaderField {
			if !isDataKey {
				// keep these header fields but convert them
				isDateField := false
				for _, d := range DateFieldNames {
					if field == d {
						isDateField = true
						break
					}
				}
				if isDateField {
					// convert the date to an epoch
					headerData = append(headerData, dateToEpoch(allData[fIndex]))
				} else {
					// keep the field as is
					headerData = append(headerData, allData[fIndex])
				}
			} else {
				// blank out the dataKeyFields - they will be squeezed out by the join below
				headerData = append(headerData, "")
			}
		}
	}
	dataKey := strings.Join(dataKeyFields, "_")
	if dataKey == "" {
		// if the dataKey is empty this is an error
		return "", nil, nil, "", desc_index, fmt.Errorf("UNPARSABLE_LINE: dataKey is empty")
	}
	// return the lineType, the headerData, the dataData, the dataKey, and the desc_index
	return fileLineType, headerData, dataData, dataKey, desc_index, nil
}

/*
Given a line from the ColumnDefinitions file this function will return the header fields and the data fields
for the line. The header fields are the fields that are used to identify the document. The data fields
are the fields that are used to populate the data section of the document. The parts are
delimited by specific fields for each file type.

NOTE: This is different than the dataKeyFields. The dataKeyFields are used to merge documents. This
function is used to split the header and data fields from the line.
*/
func SplitColumnDefLine(fileLineType string, headerLine string) ([]string, []string) {
	var headerString string
	var dataString string
	var parts []string
	switch {
	case fileLineType == "MTD",
		fileLineType == "MTD_2DSINGLE",
		fileLineType == "MTD_3DSINGLE",
		fileLineType == "MTD_3DPAIR":
		parts = strings.Split(headerLine, " OBS_LEV ")
		if len(parts) > 1 {
			headerString = parts[0] + " OBS_LEV"
			dataString = parts[1]
		} else {
			headerString = parts[0]
			dataString = ""
		}
	case fileLineType == "MODE_OBJ",
		fileLineType == "MODE_CTS":
		parts = strings.Split(headerLine, " OBTYPE ")
		if len(parts) > 1 {
			headerString = parts[0] + " OBTYPE"
			dataString = parts[1]
		} else {
			headerString = parts[0]
			dataString = ""
		}
	default:
		// get the header fields from line
		parts = strings.Split(headerLine, " LINE_TYPE ")
		if len(parts) > 1 {
			headerString = parts[0] + " LINE_TYPE"
			dataString = parts[1]
		} else {
			headerString = parts[0]
			dataString = ""
		}
	}
	// squeeze the spaces out of the header string
	headerFields := strings.Fields(headerString)
	// get the data fields from the line
	dataFields := strings.Fields(dataString)
	return headerFields, dataFields
}

/*
This routine is used to get the key data fields for a given line type. These fields are used to
essentially merge documents. All the documents with the same header field values excluding the key data fields
are merged into a single document. The key data fields are used to index the data section of the document.
*/
func GetKeyDataFieldsForLineType(lineType string) []string {
	lineTypeUpper := strings.ToUpper(lineType)
	switch lineTypeUpper {
	case "MODE":
		// I have no idea if this is the correct key for MODE!
		return []string{"FCST_LEAD", "FCST_LEV"}
	case "MTD":
		// I have no idea if this is the correct key for MTD!
		return []string{"FCST_LEAD", "FCST_LEV"}
	default:
		return []string{"FCST_LEAD"}
	}
}

func FindType(name string, data []byte) (string, error) {
	r, _ := regexp.Compile("ato.*get_item.*" + name)
	res := r.FindSubmatch(data)
	if res == nil {
		return "", fmt.Errorf("could not find type for %s", name)
	}
	dataType := strings.Split(string(res[0]), "(")[0]
	switch dataType {
	case "atoi":
		dataType = "int"
	case "atof":
		dataType = "float64"
	default:
		dataType = "string"
	}
	return dataType, nil
}

func GetId(tmpHeaderData []string, metaData *VxMetadata) VxMetadata {
	idElems := []string{metaData.Subset, metaData.Type, metaData.SubType}
	idElems = append(idElems, tmpHeaderData...)
	id := strings.Join(idElems, ":")
	metaData.ID = id
	return *metaData
}

func getLineTypeFromColumnDefsFile(headerLine string) (HeaderFields, error) {
	// this function will read the column definitions file and return the line type for the header line
	// if it is found in the column definitions file
	// If the columnsDefinition file is not present then the file will be downloaded from the
	// structColumnTypes.MetHeaderColumnsFileUrl
	var columnDefHeaderFields HeaderFields
	var trimmedColumnDefsLine string
	var lines []string
	wd, _ := os.Getwd()
	if strings.Contains(headerLine, " CYCLONE ") {
		// these are TCST files
		columnDefHeaderFields.FileType = "TCST"
		columnDefHeaderFields.SeparatorField = "LINE_TYPE"
		columnDefHeaderFields.Header = headerLine
		if strings.Contains(headerLine, " WATCH_WARN ") {
			columnDefHeaderFields.LineType = "TCMPR"
		} else if strings.Contains(headerLine, " DIAG_SOURCE ") {
			columnDefHeaderFields.LineType = "TCDIAG"
		} else if strings.Contains(headerLine, " RIRW_WINDOW ") {
			columnDefHeaderFields.LineType = "PROBRIRW"
		}
		return columnDefHeaderFields, nil
	} else if strings.Contains(headerLine, " SPACE_CENTROID_DIST ") {
		// these are MTD 3dpair files
		columnDefHeaderFields.Header = headerLine
		columnDefHeaderFields.SeparatorField = "OBJECT_ID"
		columnDefHeaderFields.FileType = "MTD"
		columnDefHeaderFields.LineType = "3DPAIR"
		return columnDefHeaderFields, nil
	} else if strings.Contains(headerLine, " TIME_INDEX ") {
		// MTD 2dsingle files
		columnDefHeaderFields.Header = headerLine
		columnDefHeaderFields.SeparatorField = "OBJECT_ID"
		columnDefHeaderFields.FileType = "MTD"
		columnDefHeaderFields.LineType = "2DSINGLE"
		return columnDefHeaderFields, nil
	} else if strings.Contains(headerLine, " CDIST_TRAVELLED ") {
		// MTD 3dsingle files
		columnDefHeaderFields.Header = headerLine
		columnDefHeaderFields.SeparatorField = "OBJECT_ID"
		columnDefHeaderFields.FileType = "MTD"
		columnDefHeaderFields.LineType = "3DSINGLE"
		return columnDefHeaderFields, nil
	} else if strings.Contains(headerLine, " GRID_RES ") {
		// these are MODE files
		columnDefHeaderFields.Header = headerLine
		columnDefHeaderFields.FileType = "MODE"
		if strings.Contains(headerLine, " OBJECT_ID ") {
			columnDefHeaderFields.SeparatorField = "OBJECT_ID"
			columnDefHeaderFields.LineType = "OBJ"
		} else if strings.Contains(headerLine, " FIELD ") {
			columnDefHeaderFields.SeparatorField = "FIELD"
			columnDefHeaderFields.LineType = "CTS"
		}
		return columnDefHeaderFields, nil
	}
	// If I made it to here I couldn't parse it from the headerline easily so I have to read the column definitions file
	// and find the line type for the header line by looking line by line in the column definitions file.
	// The column definitions file is a text file that contains the line type for each header line.
	// If the column definitions file is not present then the file will be downloaded from the
	// structColumnTypes.MetHeaderColumnsFileUrl.
	columnDefsFilePath := wd + "/" + "./column_defs.txt"
	_, err := os.Stat(columnDefsFilePath)
	if os.IsNotExist(err) {
		resp, err := http.Get(MetHeaderColumnsFileUrl)
		if err != nil {
			return HeaderFields{}, err
		}
		defer resp.Body.Close()
		out, err := os.Create(columnDefsFilePath)
		if err != nil {
			return HeaderFields{}, err
		}
		defer out.Close()
		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("error getting met_header_columns file: " + err.Error())
			return HeaderFields{}, err
		}
	}
	// read the column definitions file
	rawColumnsBytes, err := os.ReadFile(columnDefsFilePath)
	if err != nil {
		return HeaderFields{}, err
	}
	// Either the ColumnsDefinitions file is present or we just downloaded it
	// Now we need to find the header line in the file
	lines = strings.Split(string(rawColumnsBytes), "\n")
	// find the trimmed header line in the column definitions file.
	trimmedHeaderLine := strings.Join(strings.Fields(headerLine), " ")
	found := false
	for _, line := range lines {
		if line == "" {
			continue
		}
		// ignore the first 3 fields of the columndefs line for comparison
		// squeeze extra white space out of the line and remove the first 3 prefix fields
		trimmedColumnDefsLine = strings.Split(strings.Join(strings.Fields(line), " "), ": ")[3]
		if strings.HasPrefix(trimmedHeaderLine, trimmedColumnDefsLine) {
			columnDefHeaderFields.Header = line
			lineFields := strings.Split(columnDefHeaderFields.Header, ":")
			columnDefHeaderFields.FileType = strings.TrimSpace(lineFields[1])
			columnDefHeaderFields.LineType = strings.TrimSpace(lineFields[2])
			switch columnDefHeaderFields.FileType {
			case "STAT":
				columnDefHeaderFields.SeparatorField = "LINE_TYPE"
			case "MODE":
				if strings.Contains(columnDefHeaderFields.LineType, "OBJ") {
					columnDefHeaderFields.SeparatorField = "OBJECT_ID"
				} else if strings.Contains(columnDefHeaderFields.LineType, "CTS") {
					columnDefHeaderFields.SeparatorField = "FIELD"
				}
			case "MTD":
				columnDefHeaderFields.SeparatorField = "OBJECT_ID"
			case "TCST":
				columnDefHeaderFields.SeparatorField = "LINE_TYPE"
			}
			found = true
			break
		}
	}
	if found {
		return columnDefHeaderFields, nil
	} else {
		// we didn't find the separator so we can't process this line
		return HeaderFields{}, fmt.Errorf("separator not found in column definitions file")
	}
}
