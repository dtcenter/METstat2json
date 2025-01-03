package buildHeaderLineTypeUtilities

import (
	"fmt"
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
	"MODE_OBJ":       {"FCST_LEAD, FCST_LEV"},
	"MODE_CTS":       {"FCST_LEAD, FCST_LEV"},
	"MTD_2DSINGLE":   {"FCST_LEAD, FCST_LEV"},
	"MTD_3DSINGLE":   {"FCST_LEAD, FCST_LEV"},
	"MTD_3DPAIR":     {"FCST_LEAD, FCST_LEV"},
	"TCST_TCMPR":     {"FCST_LEAD"},
	"TCST_TCDIAG":    {"FCST_LEAD"},
	"TCST_PROBRIRW":  {"FCST_LEAD"},
}

var DateFieldNames = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}

func dateToEpoch(date string) string {
	// convert the date to an epoch from YYYYMMDD_HHMMSS
	// the date is in the format YYYYMMDD_HHMMSS - time zone defaults to UTC
	theTime, _ := time.Parse(`20060102_150405`, date)
	return strconv.FormatInt(theTime.Unix(), 10)
}

/* returns the lineType of the data line, and the headerFields and dataFields of the data line as []string.
   The lineType is the last field in the header section of the header line, which is the first line of the file.
   Since this function has to split the line into a header and a data section those fields are returned as well.
   The dataKeyFields are removed from the headerFields and the headerData. The dataFields are returned as is.
   The headerData is converted to epochs if the field is a dateField.
*/

func GetLineType(headerLine string, dataLine string, fileType string) (string, []string, []string, string, error) {
	headerStringFields, _ := SplitColumnDefLine(fileType, headerLine)
	allData := strings.Fields(dataLine)
	lineTypeIndex := len(headerStringFields) - 1
	if lineTypeIndex > len(allData) {
		return "", nil, nil, "", fmt.Errorf("lineTypeIndex is greater than the length of the data line")
	}
	lineType := allData[lineTypeIndex]
	dataData := strings.Fields(dataLine)[lineTypeIndex+1:]

	// have to remove the dataKeyFields from the headerFields and the headerData (dataData AND dataFields won't change)
	headerData := []string{}
	dataKeyFields := []string{}
	for hIndex, h := range headerStringFields {
		isDataKey := false
		for _, d := range DataKeyMap[fileType+"_"+lineType] {
			if h == d {
				isDataKey = true
				dataKeyFields = append(dataKeyFields, allData[hIndex])
				break
			}
		}
		// skip the dataKeyFields - blank them out
		if !isDataKey {
			// keep these fields
			isDateField := false
			for _, d := range DateFieldNames {
				if h == d {
					isDateField = true
					break
				}
			}
			if isDateField {
				// convert the date to an epoch
				headerData = append(headerData, dateToEpoch(allData[hIndex]))
			} else {
				// keep the field as is
				headerData = append(headerData, allData[hIndex])
			}
		} else {
			// blank out the dataKeyFields
			headerData = append(headerData, "")
		}
	}
	dataKey := strings.Join(dataKeyFields, "_")
	return lineType, headerData, dataData, dataKey, nil
}

/*
Given a line from the ColumnDefinitions file this function will return the header fields and the data fields
for the line. The header fields are the fields that are used to identify the document. The data fields
are the fields that are used to populate the data section of the document. The parts are
delimited by specific fields for each file type.
*/
func SplitColumnDefLine(fileType string, fieldStr string) ([]string, []string) {
	var headerString string
	var dataString string
	var parts []string
	switch fileType {
	case "MODE", "MTD":
		parts = strings.Split(fieldStr, " OBS_LEV ")
		if len(parts) > 1 {
			headerString = parts[0] + " OBS_LEV"
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
