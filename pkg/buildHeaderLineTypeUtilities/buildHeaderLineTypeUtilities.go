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
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// vxMetadata struct definition
type VxMetadata struct {
	ID      string `json:"id"`
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

// data key definitions
var DataKeyMap = map[string][]string{
	"STAT_CNT":              {"FCST_LEAD"},
	"STAT_CTC":              {"FCST_LEAD"},
	"STAT_CTS":              {"FCST_LEAD"},
	"STAT_FHO":              {"FCST_LEAD"},
	"STAT_ISC":              {"FCST_LEAD"},
	"STAT_MCTC":             {"FCST_LEAD"},
	"STAT_MCTS":             {"FCST_LEAD"},
	"STAT_MPR":              {"FCST_LEAD"},
	"STAT_SEEPS":            {"FCST_LEAD"},
	"STAT_SEEPS_MPR":        {"FCST_LEAD"},
	"STAT_NBRCNT":           {"FCST_LEAD"},
	"STAT_NBRCTC":           {"FCST_LEAD"},
	"STAT_NBRCTS":           {"FCST_LEAD"},
	"STAT_GRAD":             {"FCST_LEAD"},
	"STAT_DMAP":             {"FCST_LEAD"},
	"STAT_ORANK":            {"FCST_LEAD"},
	"STAT_PCT":              {"FCST_LEAD"},
	"STAT_PJC":              {"FCST_LEAD"},
	"STAT_PRC":              {"FCST_LEAD"},
	"STAT_PSTD":             {"FCST_LEAD"},
	"STAT_ECLV":             {"FCST_LEAD"},
	"STAT_ECNT":             {"FCST_LEAD"},
	"STAT_RPS":              {"FCST_LEAD"},
	"STAT_RHIST":            {"FCST_LEAD"},
	"STAT_PHIST":            {"FCST_LEAD"},
	"STAT_RELP":             {"FCST_LEAD"},
	"STAT_SAL1L2":           {"FCST_LEAD"},
	"STAT_SL1L2":            {"FCST_LEAD"},
	"STAT_SSVAR":            {"FCST_LEAD"},
	"STAT_VAL1L2":           {"FCST_LEAD"},
	"STAT_VL1L2":            {"FCST_LEAD"},
	"STAT_VCNT":             {"FCST_LEAD"},
	"STAT_GENMPR":           {"FCST_LEAD"},
	"STAT_SSIDX":            {"FCST_LEAD"},
	"MODE_OBJ":              {"FCST_LEAD", "OBJECT_ID"}, // this one is a datafield key
	"MODE_CTS":              {"FCST_LEAD"},
	"MODE_2D":               {"FCST_LEAD", "OBJECT_ID"},
	"MTD_3D_PAIR_CLUSTER":   {"FCST_LEAD", "OBJECT_ID"},
	"MTD_3D_PAIR_SIMPLE":    {"FCST_LEAD", "OBJECT_ID"},
	"MTD_3D_SINGLE_CLUSTER": {"FCST_LEAD", "OBJECT_ID"},
	"MTD_3D_SINGLE_SIMPLE":  {"FCST_LEAD", "OBJECT_ID"},
	"TCST_TCMPR":            {"FCST_LEAD"},
	"TCST_TCDIAG":           {"FCST_LEAD"},
	"TCST_PROBRIRW":         {"FCST_LEAD"},
}

var DateFieldNames = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}

func dateToEpoch(date string) string {
	// convert the date to an epoch from YYYYMMDD_HHMMSS
	// the date is in the format YYYYMMDD_HHMMSS - time zone defaults to UTC
	theTime, _ := time.Parse(`20060102_150405`, date)
	return strconv.FormatInt(theTime.Unix(), 10)
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
	if strings.HasPrefix(fileName, "mode_") {
		if strings.HasSuffix(fileName, "obj.txt") || strings.HasSuffix(fileName, "merge.txt") {
			separatorField = "OBJECT_ID"
			fileLineType = "MODE_OBJ"
		} else if strings.HasSuffix(fileName, "cts.txt") {
			separatorField = "FIELD"
			fileLineType = "MODE_CTS"
		}
	} else if strings.HasPrefix(fileName, "mtd_") {
		separatorField = "OBJECT_ID"
		if strings.Contains(fileName, "2d") {
			fileLineType = "MTD_2D"
		} else if strings.Contains(fileName, "3d_single_cluster") {
			fileLineType = "MTD_3D_SINGLE_CLUSTER"
		} else if strings.Contains(fileName, "3d_single_simple") {
			fileLineType = "MTD_3D_SINGLE_SIMPLE"
		} else if strings.Contains(fileName, "3d_pair_cluster") {
			fileLineType = "MTD_3D_PAIR_CLUSTER"
		} else if strings.Contains(fileName, "3d_pair_simple") {
			fileLineType = "MTD_3D_PAIR_SIMPLE"
		}
	} else {
		separatorField = "LINE_TYPE"
		// don't know the fileLineType for stat files yet.
	}
	headerStringFields, _ := SplitColumnDefLine(separatorField, headerLine)
	dataStartIndex := len(headerStringFields)
	dataData := allData[dataStartIndex:]
	lineTypeIndex := len(headerStringFields) - 1
	if lineTypeIndex > len(allData) {
		return "", nil, nil, "", desc_index, fmt.Errorf("UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line")
	}
	// now we know the lineType for stat files.
	if separatorField == "LINE_TYPE" {
		fileLineType = "STAT" + "_" + allData[lineTypeIndex]
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
				dataKeyFields = append(dataKeyFields, allData[fIndex])
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
				// blank out the dataKeyFields
				headerData = append(headerData, "")
			}
		}
	}
	dataKey := strings.Join(dataKeyFields, "_")
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
func SplitColumnDefLine(fileType string, fieldStr string) ([]string, []string) {
	var headerString string
	var dataString string
	var parts []string
	switch fileType {
	case "MTD", "MTD_2DSINGLE", "MTD_3DSINGLE", "MTD_3DPAIR":
		parts = strings.Split(fieldStr, " OBS_LEV ")
		if len(parts) > 1 {
			headerString = parts[0] + " OBS_LEV"
			dataString = parts[1]
		} else {
			headerString = parts[0]
			dataString = ""
		}
	case "MODE_OBJ", "MODE_CTS":
		parts = strings.Split(fieldStr, " OBTYPE ")
		if len(parts) > 1 {
			headerString = parts[0] + " OBTYPE"
			dataString = parts[1]
		} else {
			headerString = parts[0]
			dataString = ""
		}
	default:
		// get the header fields from line
		parts = strings.Split(fieldStr, " LINE_TYPE ")
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
