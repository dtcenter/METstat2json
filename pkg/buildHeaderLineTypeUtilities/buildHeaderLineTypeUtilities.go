package buildHeaderLineTypeUtilities

/*
This package contains the utilities for building the header and line type for the data files. The data files are
MET output files that contain a header section and a data section. The header section contains the header fields that
are used to identify the document. The data section contains the data fields that are used to populate the document.
This package is separate from the metLineTypeDefinitions package because the metLineTypeDefinitions package is automatically
generated from the buildHeaderLineTypes.go program and there is a desire to avoid a circular dependency.
This package defines a VxMetaData struct that is used to store the metadata for the mapped documents.
The metadata is used to uniquely identify each document and is used to merge documents with the same metadata.

This package also defines the DataKeyMap that is used to determine the key data fields for a given line type.
The key data fields are used to merge documents with the same header field values excluding the key data fields.
It also defines header keys that are not allowed to be in the header.
Other utilities exist to convert the date fields to epochs, to get the line type of the data line, to get the key
data fields for a given line type, and to find the data type of a given field in addition to some other utility functions.
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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

var (
	MetHeaderColumnsFileUrl_v12_0 = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
	MetHeaderColumnsFileUrl_v11_1 = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V11.1.txt"
	MetHeaderColumnsFileUrl_v11_0 = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V11.0.txt"
	MetHeaderColumnsFileUrl_v10_1 = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V10.1.txt"
	MetHeaderColumnsFileUrl_v10_0 = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V10.0.txt"
	MetHeaderColumnsFileUrl       = MetHeaderColumnsFileUrl_v12_0
)

var MetUserDocFiles = []string{
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/ensemble-stat.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/grid-stat.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/gsi-tools.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/mode-td.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/mode.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/point-stat.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/stat-analysis.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/tc-gen.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/tc-pairs.rst",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/wavelet-stat.rst",
}

var MetSrcFiles = []string{
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/mode_line.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_grid/unstructured_grid.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_info_base.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_pair_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/track_pair_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/core/stat_analysis/parse_stat_line.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/tc_utils/tc_stat/tc_stat_job.cc",
}

// vxMetadata struct definition
type VxMetadata struct {
	ID      string `json:"id"`
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

type DataKeyEntry struct {
	DataKey        []string
	HeaderDisallow []string
}

// data key definitions
var DataKeyMap = map[string]DataKeyEntry{
	"STAT_CNT":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_CTC":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_CTS":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_FHO":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_ISC":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_MCTC":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_MCTS":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_MPR":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SEEPS":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SEEPS_MPR": {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_NBRCNT":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_NBRCTC":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_NBRCTS":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_GRAD":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_DMAP":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_ORANK":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_PCT":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_PJC":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_PRC":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_PSTD":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_ECLV":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_ECNT":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_RPS":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_RHIST":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_PHIST":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_RELP":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SAL1L2":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SL1L2":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SSVAR":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_VAL1L2":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_VL1L2":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_VCNT":      {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_GENMPR":    {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"STAT_SSIDX":     {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"MODE_OBJ":       {DataKey: []string{"FCST_LEAD", "OBJECT_ID"}, HeaderDisallow: nil},
	"MODE_CTS":       {DataKey: []string{"FCST_LEAD"}, HeaderDisallow: nil},
	"MTD_2DSINGLE":   {DataKey: []string{"OBJECT_ID", "TIME_INDEX"}, HeaderDisallow: nil},
	"MTD_3DSINGLE":   {DataKey: []string{"OBJECT_ID"}, HeaderDisallow: nil},
	"MTD_3DPAIR":     {DataKey: []string{"OBJECT_ID"}, HeaderDisallow: nil},
	"TCST_TCMPR":     {DataKey: []string{"LEAD"}, HeaderDisallow: []string{"INIT"}},
	"TCST_TCDIAG":    {DataKey: []string{"LEAD"}, HeaderDisallow: []string{"INIT"}},
	"TCST_PROBRIRW":  {DataKey: []string{"LEAD"}, HeaderDisallow: []string{"INIT"}},
}

// these fields will be converted to an int in the header section
var IntFieldNames = []string{"FCST_LEAD", "OBS_LEAD", "LEAD"}

// these fields will be converted to an epoch int in the header section
var DateFieldNames = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END", "INIT", "VALID"}

func dateToEpoch(date string) string {
	// convert the date to an epoch from YYYYMMDD_HHMMSS
	// the date is in the format YYYYMMDD_HHMMSS - time zone defaults to UTC
	theTime, _ := time.Parse(`20060102_150405`, date)
	return strconv.FormatInt(theTime.Unix(), 10)
}

func getLeadFromInitValid(headerFields []string, data []string, dataFieldIndex int) string {
	if len(headerFields) != len(data) {
		return "MISSING"
	}
	if len(data) < 3 {
		return "MISSING"
	}
	if (headerFields[dataFieldIndex] != "LEAD" &&
		headerFields[dataFieldIndex] != "FCST_LEAD") ||
		headerFields[dataFieldIndex-1] != "INIT" ||
		headerFields[dataFieldIndex+1] != "VALID" {
		return "MISSING"
	}
	initTime, err := time.Parse("20060102_150405", data[dataFieldIndex-1])
	if err != nil {
		return "MISSING"
	}
	validTime, err := time.Parse("20060102_150405", data[dataFieldIndex+1])
	if err != nil {
		return "MISSING"
	}
	lead := validTime.Sub(initTime)
	intVal := int64(lead.Hours())
	// padd for minutes and seconds (4 chars on the end)
	zeroVal := false
	if intVal < 0 {
		zeroVal = true
		intVal = intVal * -1
	}
	var strVal string
	if intVal < 10 {
		strVal = fmt.Sprintf("%02d", intVal) // result will be like "01"
	} else {
		strVal = fmt.Sprintf("%d", intVal)
	}
	if zeroVal {
		strVal = "-" + strVal
	}
	return strVal + "0000" // add the minutes and seconds
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
		The DataKey is the concatenation of the DataKeyFields. The desc_index is the index of the DESC field in the headerFields.
		The DataKeyFields are used to merge documents. A DataKey field may be derived from either header fields or data fields.
		If the DataKey is derived from header fields they will be removed from the headerFields and the headerData.
		If the DataKey is derived from data fields they will not be removed from the data element.
		Any headerData is converted to epochs if the field is a dateField.

		NOTE: Often a stat file will not have a header line that contains data fields, i.e. it will end with the LINE_TYPE.
		This is when the data file has different types of stat data in the same file.
		In this case the dataFields cannot be determined and if the DataKeyMap specifies a DataKeyField that is not in the
		headerFields then the DataKey will be "" and the entire header will be used to generate the id.
*/
func GetLineType(headerLine string, dataLine string, fileName string, version string) (string, []string, []string, string, int, error) {
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
		columnDefHeaderFields, err := getLineTypeFromColumnDefsFile(headerLine, version)
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
	// have to remove the DataKeyFields from the headerFields and the headerData (dataData AND dataFields won't change)
	headerData := []string{}
	DataKeyFields := []string{}
	// look for DataKey fields in allData - remove the DataKeyFields from the headerData if the key is in the header
	// or if the key is in the disallowed header fields
	for fIndex, field := range allHeaderFields {
		isDataKey := false
		for _, dk := range DataKeyMap[fileLineType].DataKey {
			if field == dk {
				isDataKey = true
				if (dk == "LEAD" || dk == "FCST_LEAD") && allData[fIndex] == "NA" {
					// This is a special case.
					// if the datakey is LEAD or FCST_LEAD and the lead is NA then we
					// have to get the lead from the init and valid fields, if possible,
					// and use that as the lead. If it isn't possible to get the lead
					// from the init and valid fields then we will set the lead to "MISSING"
					// and the data section will have one entry with the key set to "MISSING"
					if allData[fIndex] == "NA" {
						allData[fIndex] = getLeadFromInitValid(allHeaderFields, allData, fIndex)
					}
				}
				if allData[fIndex] != "NA" {
					// the DataKey is the value that belongs to the associated field
					// not the field itself. Datakeys deliniate data sections
					// based on their values
					DataKeyFields = append(DataKeyFields, allData[fIndex])
				}
				break
			}
		}
		// handle disallowed header fields
		for _, disDk := range DataKeyMap[fileLineType].HeaderDisallow {
			if field == disDk {
				isDataKey = true
				break
			}
		}
		isHeaderField := fIndex <= lineTypeIndex
		// iterate through the header fields and
		// if the field is a header field and a DataKeyField then blank it out
		// convert header DATE fields that are not DataKeys to epochs
		// Note: DateFieldNames also get a type of int in getDataType function
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
				// blank out the DataKeyFields - they will be squeezed out by the join below
				headerData = append(headerData, "")
			}
		}
	}
	DataKey := strings.Join(DataKeyFields, "_")
	if DataKey == "" {
		// if the DataKey is empty this is an error
		return "", nil, nil, "", desc_index, fmt.Errorf("UNPARSABLE_LINE: DataKey is empty")
	}
	// return the lineType, the headerData, the dataData, the DataKey, and the desc_index
	return fileLineType, headerData, dataData, DataKey, desc_index, nil
}

/*
Given a line from the ColumnDefinitions file this function will return the header fields and the data fields
for the line. The header fields are the fields that are used to identify the document. The data fields
are the fields that are used to populate the data section of the document. The parts are
delimited by specific fields for each file type.

NOTE: This is different than the DataKeyFields. The DataKeyFields are used to merge documents. This
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

// GetHeaderValue returns the value of a header field in the headerData slice.
// If the field is not found, it returns -1.
func GetHeaderValue(headerFields []string, headerData []string, field string) (string, error) {
	for i, h := range headerFields {
		if h == field {
			// if the field is a date field then convert it to an epoch
			if slices.Contains(DateFieldNames, field) {
				return dateToEpoch(headerData[i]), nil
			}
			return headerData[i], nil
		}
	}
	return "", nil
}

func GetId(dataSetName string, tmpHeaderData []string, metaData *VxMetadata) (VxMetadata, error) {
	idElems := []string{metaData.Subset, metaData.Type, metaData.SubType, dataSetName}
	idElems = append(idElems, tmpHeaderData...)
	id := strings.Join(idElems, ":")
	if len(id) > 250 {
		return VxMetadata{}, fmt.Errorf("Calculated ID is too long: %d - id:\"%s\"", len(id), id)
	}
	metaData.ID = id
	return *metaData, nil
}

func getLineTypeFromColumnDefsFile(headerLine string, version string) (HeaderFields, error) {
	// this function will read the column definitions file and return the line type for the header line
	// if it is found in the column definitions file
	// If the columnsDefinition file is not present then the file will be downloaded from the
	// metLineTypeDefinitions.MetHeaderColumnsFileUrl
	switch version {
	case "v12_0":
		MetHeaderColumnsFileUrl = MetHeaderColumnsFileUrl_v12_0
	case "v11_1":
		MetHeaderColumnsFileUrl = MetHeaderColumnsFileUrl_v11_1
	case "v11_0":
		MetHeaderColumnsFileUrl = MetHeaderColumnsFileUrl_v11_0
	case "v10_1":
		MetHeaderColumnsFileUrl = MetHeaderColumnsFileUrl_v10_1
	case "v10_0":
		MetHeaderColumnsFileUrl = MetHeaderColumnsFileUrl_v10_0
	default:
		log.Printf("version not found, %s - using v12.0", version)
		return HeaderFields{}, fmt.Errorf("Unsupported version %s", version)
	}

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
	// metLineTypeDefinitions.MetHeaderColumnsFileUrl.
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
