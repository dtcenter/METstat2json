package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"parser/pkg/buildHeaderLineTypeUtilities"
	"regexp"
	"strings"
)

/*
NOTES:
See https://github.com/dtcenter/MET/blob/main_v12.0/data/table_files/met_header_columns_V12.0.txt#L24.
also see https://met.readthedocs.io/en/latest/Users_Guide/index.html and look at the data type definitions
for the various MET file types. The data type definitions are in the "Column Definitions" section of the
various MET file type documentation. For example see https://met.readthedocs.io/en/latest/Users_Guide/mode.html#mode-output
for the MODE file type output format.

Delineating header from data sections in the header records:
The field that delineates the header section from the data section in most line types is the LINE_TYPE field
The exceptions are MODE and MTD.
MODE object header records use the "OBJECT_ID" field and MODE cts headers use the "FIELD" field.
MTD headers use the "OBJECT_ID" field.

Recognizing data types:
The way that you can tell if a data file is a mode file is that
MODE object files begin with mode_ and end with _obj.txt and the CTS files begin with mode_
and end with _cts.txt. However, Grid-Stat and Point-Stat also write output files with the
_cts.txt suffix, but they begin with grid_stat_ and point_stat_ , respectively, and they have a LINE_TYPE field.

MTD files start with mtd_ and end with either _2d.txt, _3d_pair_cluster.txt, _3d_pair_simple.txt, _3d_single_cluster.txt, or _3d_single_simple.txt
Other files have a line type field that is used to determine the data types of the fields in the data section of the document.
There can be multiples of these line types in a single file.
*/
type Pattern struct {
	match       *regexp.Regexp
	dType       string
	structField string
	structType  string
}

var patterns = make(map[string]Pattern)

var metSrcFiles = []string{
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/mode_line.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/stat_job.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_grid/unstructured_grid.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_info_base.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_pair_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/track_pair_info.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/core/stat_analysis/parse_stat_line.cc",
	"https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/tc_utils/tc_stat/tc_stat_job.cc",
}

var metUserDocFiles = []string{
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

var metHeaderColumnsFileUrl = buildHeaderLineTypeUtilities.MetHeaderColumnsFileUrl

/*
The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types.
There are several repeating fields that are handled in a special way. These fields are handled in the getRepeatingSequenceStructureString
function. See the notes in that function for more information.
*/

func main() {
	type HeaderStructs map[string]string
	type DataStructs map[string]string

	// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
	// currently 12.0 to get the required header column definitions and then using
	// using the https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide
	// to get the stat field types. We create a map of field names to field types. Then we have to re-iterate
	// over the header definitions to create the structs and functions to fill the structs.
	// read the header columns file
	met_header_columns_lines, fieldNameMap := getColumnLinesAndMapForUrl(metHeaderColumnsFileUrl)
	metDataTypesForLines := make(map[string]string)
	metDataTypesForLines = fillMetDataMapFromSrcFiles(metDataTypesForLines, fieldNameMap)
	metDataTypesForLines = fillMetDataMapFromUserGuide(metDataTypesForLines, fieldNameMap)
	// Use a map to keep track of unique headerStructs and dataStructs.
	dataStructs := make(DataStructs)
	headerStructs := make(HeaderStructs)
	// fill...Funcs are the same as header and data Structs but for the fill functions
	fillHeaderFuncs := make(HeaderStructs)
	fillDataFuncs := make(DataStructs)
	// create the getDocFoID function
	docIDString := "func GetDocForId(fileLineType string, metaDataMap map[string]interface{}, headerData []string, dataData []string, dataKey string) (map[string]interface{}, error) {\n\tdoc := make(map[string] interface{})\n\t" +
		"// add the metadata to the doc\n\tfor key, value := range metaDataMap {\n\t\tdoc[key] = value\n\t}\n\tswitch fileLineType {\n"
	addDataElementString := "func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) (map[string]interface{}, error) {\n\tswitch fileLineType {\n"
	// iterate through every line in the met_header_columns file to create the getDocId case and the structs and functions for each met header column line
	var docStructName, headerStructName, headerStructString, fillHeaderString string
	for _, line := range met_header_columns_lines {
		// get the prefix from the line
		fieldStr, fileType, lineType, err := getFileLineType(line)
		if err != nil {
			continue // skip this line - it didn't parse properly
		}
		fileLineType := fileType + "_" + lineType
		// split the line into header and data fields
		headerFields, dataFields := buildHeaderLineTypeUtilities.SplitColumnDefLine(fileLineType, fieldStr)
		// create the header struct string and the fillHeader function string
		docStructName, headerStructName, headerStructString, fillHeaderString, docIDString, addDataElementString = getHeaderStructureString(fileType, lineType, docIDString, addDataElementString, headerFields)
		// add the header struct string to the map for printing later
		headerStructs[headerStructName] = headerStructString
		// add the fillHeader function string to the map for printing later
		fillHeaderFuncs[headerStructName] = fillHeaderString
		// create dataStructure and fillStructure function strings for the data struct
		fillStructureString, dataStruct := getFillStructureString(docStructName, dataFields, metDataTypesForLines, fileType, lineType)
		dataStructs[docStructName] = dataStruct
		fillDataFuncs[docStructName] = fillStructureString
	}
	// end the switch statements in the getDocForId and addDataElementStrings now that the line loop is over
	docIDString += "\tdefault:\n\t\treturn nil, errors.New(\"GetDocForId: Unknown file_line type:\" + fileLineType)\n\t}\n\treturn doc, nil\n}\n"
	addDataElementString += "\tdefault:\n\t\treturn nil, errors.New(\"AddDataElement: Unknown file_line type:\" + fileLineType)\n\t}\n\treturn *doc, nil\n}\n"

	// print the package - header structs, fillHeader functions, data structs, fillStructure functions, getDocForId functions, addDataElement functions
	fmt.Println("package structColumnTypes")
	fmt.Println("")
	fmt.Println("import (\n\t\"strconv\"\n\t\"errors\"\n\t\"fmt\"\n)")
	fmt.Println("\n/*\nTHIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE")
	fmt.Println("To modify this code - modify the buildHeaderLineTypes.go file and run the buildHeaderLineTypes.go program")
	fmt.Println("cd  <repo path>/metlinetypes/pkg/buildHeaderLineTypes")
	fmt.Println("go run . > /tmp/types.go")
	fmt.Println("cp /tmp/types.go ../structColumnTypes/structColumnTypes.go\n*/")
	fmt.Println("")

	fmt.Println("const DOC_NOT_FOUND = \"Document not found:\"")
	// print the header structs
	fmt.Println("")
	fmt.Println("//Header struct definitions")
	for _, headerStruct := range headerStructs {
		// print the struct
		fmt.Println(headerStruct)
	}
	// print the fillHeader functions
	fmt.Println("")
	fmt.Println("//fillHeader functions")
	for _, fillHeaderFunc := range fillHeaderFuncs {
		// print the function
		fmt.Println(fillHeaderFunc)
	}

	// print the data structs
	fmt.Println("")
	fmt.Println("//line data struct definitions")
	for _, dataStruct := range dataStructs {
		fmt.Println(dataStruct)
	}

	// print the fillStructure functions
	fmt.Println("")
	fmt.Println("//fillStructure functions")
	for _, fillStructureFunc := range fillDataFuncs {
		// print the function
		fmt.Println(fillStructureFunc)
	}

	// print the getDocForId functions
	fmt.Println("")
	fmt.Println("//getDocForId functions")
	fmt.Println(docIDString)

	// print the addDataElement functions
	fmt.Println("")
	fmt.Println("//addDataElement functions")
	fmt.Println(addDataElementString)

	// print the DateFieldNames
	fmt.Println("")
	fmt.Println("	// DateFieldNames is a list of the date fields that need to be converted to epochs")
	fmt.Println("var DateFieldNames = []string{\"FCST_VALID_BEG\", \"FCST_VALID_END\", \"OBS_VALID_BEG\", \"OBS_VALID_END\"}")
	fmt.Println("var MetHeaderColumnsFileUrl = \"" + metHeaderColumnsFileUrl + "\"")
	fmt.Println("")
}

// private functions
func getFileLineType(line string) (string, string, string, error) {
	parts := strings.Split(line, ": VERSION")
	if len(parts) < 2 {
		if line != "" {
			fmt.Println("error parsing line: met_header_columns" + "line:'" + line + "'")
		}
		return "", "", "", fmt.Errorf("error parsing line: met_header_columns line:'%s'", line)
	}
	prefix := parts[0]
	fieldStr := "VERSION " + parts[1]
	// get the version from the line
	parts = strings.Split(prefix, " : ")
	if len(parts) < 3 {
		fmt.Println("error parsing line: " + line)
		return "", "", "", fmt.Errorf("error parsing line: met_header_columns line:'%s'", line)
	}
	fileType := strings.ToUpper(strings.TrimSpace(parts[1]))
	lineType := strings.ToUpper(strings.TrimSpace(parts[2]))
	return fieldStr, fileType, lineType, nil
}

func getHeaderStructureString(fileType string, lineType string, getDocIDString string, addDataElementString string, headerFields []string) (string, string, string, string, string, string) {
	docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
	headerStructName := fmt.Sprintf("%s_header", docStructName)

	keyFields := buildHeaderLineTypeUtilities.GetKeyDataFieldsForLineType(fileType)
	headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
	fillHeaderString := fmt.Sprintf("func (s *%s) fill_%s_Header(fields []string, doc *map[string]interface{}){\n\tdataLen := len(fields)\n\ti := -1\n", docStructName, docStructName)
	fillHeaderString += "\t// fill the met fields leaving out \"\" and NA values\n"
	getDocIDString += fmt.Sprintf("\tcase \"%s\":\n", docStructName)
	getDocIDString += fmt.Sprintf("\t\telem := %s{}\n", docStructName)

	getDocIDString += fmt.Sprintf("\t\telem.fill_%s_Header(headerData, &doc)\n", docStructName)
	getDocIDString += fmt.Sprintf("\t\telem.fill_%s(dataData)\n\t\tif exists := (doc)[\"data\"]; exists == nil {\n", docStructName)
	getDocIDString += fmt.Sprintf("\t\t\t(doc)[\"data\"] = make(map[string]%s)\n\t\t}\n\t\tif val, ok := (doc)[\"data\"].(map[string]%s); ok {\n\t\t\tval[dataKey] = elem\n\t\t\t(doc)[\"data\"] = val\n\t\t}\n", docStructName, docStructName)
	addDataElementString += fmt.Sprintf("\tcase \"%s\":\n", docStructName)
	addDataElementString += fmt.Sprintf("\t\telem := %s{}\n", docStructName)
	addDataElementString += fmt.Sprintf("\t\telem.fill_%s(dataData)\n", docStructName)
	addDataElementString += fmt.Sprintf("\t\tif val, ok := (*doc)[\"data\"].(map[string]%s); ok {\n", docStructName)
	addDataElementString += "\t\t\tval[dataKey] = elem\n\t\t\t(*doc)[\"data\"] = val\n\t\t}\n"

	// find the maximum length of the header fields for formatting (padding)
	padding := 0
	for _, term := range headerFields {
		if len(term) > padding {
			padding = len(term)
		}
	}
	for _, term := range headerFields {
		// skip the dataKey fields
		isDataKey := false
		for _, d := range keyFields {
			if term == d {
				isDataKey = true
				break
			}
		}
		// skip the dataKeyFields
		if isDataKey {
			continue
		}
		// change regex type terms
		term = strings.Replace(term, "(", "", -1)
		term = strings.Replace(term, ")", "", -1)
		term = strings.Replace(term, "[0-9]*", "i", -1)
		// name := cases.Title(language.English).String(term)
		name := strings.ToUpper(term)
		dataType := "string"
		jsonName := strings.ToLower(name)
		// headerStructString += fmt.Sprintf("    %-*s %s `json:\"%s\"`\n", padding, capName, dataType, name)
		headerStructString += fmt.Sprintf("    %-*s %s `json:\"%s\"`\n", padding, name, dataType, jsonName)
		fillHeaderString += fmt.Sprintf("\ti++; if i <= dataLen && fields[i] != \"\"  && fields[i] != \"NA\" {\n\t\t(*doc)[\"%s\"] = fields[i]\n\t}\n", name)
	}
	headerStructString += "}\n"
	fillHeaderString += "}\n"
	return docStructName, headerStructName, headerStructString, fillHeaderString, getDocIDString, addDataElementString
}

func getFillStructureString(docStructName string, dataFields []string, metDataTypesForLines map[string]string, fileType string, lineType string) (string, string) {
	// returns fillStructureString and the dataStruct
	fillStructureString := fmt.Sprintf("func (s *%s) fill_%s(fields []string) {\n\tdataLen := len(fields) - 1\n\ti := -1\n",
		docStructName, docStructName)
	// create the data struct for this line type
	dataStruct := fmt.Sprintf("type %s struct {\n", docStructName)
	// find the maximum length of the data fields for formatting (padding)
	padding := 0
	for _, term := range dataFields {
		if len(term) > padding {
			padding = len(term)
		}
	}
	padding2 := len("map[string]interface{}")
	//	for index, term := range dataFields {
	for index := 0; index < len(dataFields); index++ {
		term := dataFields[index]
		cleanTerm, dataType := getDataType(term, &metDataTypesForLines)
		jsonTerm := strings.ToLower(cleanTerm)

		dataStruct += fmt.Sprintf("    %-*s %-*s `json:\"%s\"`\n", padding, cleanTerm, padding2, dataType, jsonTerm)
		var numFields int
		var err error
		var repeatFillStructureString string
		fillStructureString += "\ti++; if i <= dataLen {"
		switch dataType {
		case "int":
			fillStructureString += fmt.Sprintf("s.%s, _ = strconv.Atoi(fields[%d])", cleanTerm, index)
		case "float64":
			fillStructureString += fmt.Sprintf("s.%s, _ = strconv.ParseFloat(fields[%d], 64)", cleanTerm, index)
		case "map[string]interface{}":
			// this is a map which means that there are a sequence of fields that are repeated
			numFields, repeatFillStructureString, err = getRepeatingSequenceStructureString(term, cleanTerm, fileType, lineType, index)
			if err != nil {
				fmt.Println("error in getRepeatingSequenceStructureString: ", err)
			}
			fillStructureString += repeatFillStructureString
			index += numFields
		default:
			fillStructureString += fmt.Sprintf("s.%s = fields[%d]", cleanTerm, index)
		}
		fillStructureString += "}\n"
	}
	fillStructureString += "}\n"

	dataStruct += "}\n"
	return fillStructureString, dataStruct
}

func getRepeatingSequenceStructureString(term string, cleanTerm string, fileType string, lineType string, index int) (numFields int, structureString string, err error) {
	/*

		The function definition is already in the fillStructureString in the caller so here we just need to return the
		part of the fillStructure string that does the logic of filling the repeating struct fields which are of
		differing lengths depending on the data file. This string will get appended to the caller's string. We also
		have to return the newIndex which is the number of fields that get processed in this part of the function,
		because there may be single fields remaining in the data line after the repeating fields are swallowed up, and
		the caller needs to know what the index of those fields is.

		There are several possible repeating fields depending on line and file type.
			Repeated sequences:
			Some line types can have repeated sequences of fields.
			These fields must be represented in the data sections of documents as embedded maps.
			These are identified by the N_ prefix in the field name and the field name being surrounded by parenthesis.
			The key for the map is patterned after the field name with the N_ prefix removed.

			(N_CAT) for MCTC files, the repeated sequence is F[0-9]*_O[0-9]* that will be contained in a map[string]int.

			(N_THRESH) for PCT which is a sequence of THRESH_n OY_n ON_n key/values that will be contained in a map[string]interface{}
			where the keys are THRESH_n, OY_n, or ON_n with the n being the threshold number.

			(N_THRESH) for PJC which is a sequence of THRESH_n, OY_TP_n, ON_TP_n, CALIBRATION_n, REFINEMENT_n, LIKELIHOOD_n, BASER_n key/values
			that will be contained in a map[string]interface{}
			where the keys are THRESH_n, OY_TP_n, ON_TP_n, CALIBRATION_n, REFINEMENT_n, LIKELIHOOD_n, and BASER_n.

			(N_THRESH) for PRC which is a sequence of THRESH_n PODY_n POFD_n key/values that will be contained in a map[string]interface{}
			where the keys are THRESH_n PODY_n POFD_n.

			(N_THRESH) for PSTD which is a map[string]interface{}
			where the keys are all THRESH_n

			(N_THRESH) for PROBRIRW which is a sequence of THRESH_n PROB_n that will be contained in a map[string]interface{}
			where the keys are THRESH_n and PROB_n.

			(N_PTS) for ECLV files which is a sequence of CL_n VALUE_n key/values that will be contained in a map[string]interface{}
			where the keys are CL_n and VALUE_n e.g. CL_1, VALUE_1, CL_2, VALUE_2, etc.

			(N_RANK) for RHIST files, the repeated sequence is RANK_n that will be contained in a map[string]int
			where the keys are RANK_n e.g. RANK_1, RANK_2, etc.

			(N_BIN) for PHIST files, the repeated sequence is BIN_n that will be contained in a map[string]int
			where the keys are BIN_n e.g. BIN_1, BIN_2, etc.

			(N_ENS) for ORANK files, the repeated sequence is ENS_n that will be contained in a map[string]interface{}
			where the keys are ENS_n e.g. ENS_1, ENS_2, etc.

			(N_ENS) for RELP files, the repeated sequence is RELP_n that will be contained in a map[string]interface{}
			where the keys are RELP_N e.g. RELP_1, RELP_2, etc.

			(N_DIAG) for TCDIAG files, the repeated sequence is DIAG_n VALUE_n that will be contained in a map[string]interface{}
			where the keys are DIAG_n and VALUE_n e.g. DIAG_1, VALUE_1, DIAG_2, VALUE_2 etc.
	*/
	fileLineType := fileType + "_" + lineType
	switch term {
	case "(N_CAT)":
		/*  MCTC files have a sequence of Fn_On key/values in an n dimensional array of ints.
		MCTC records - I find these in different file types, e.g. grid_stat_APCP as well as grid_stat...mctc.txt files.
		The 25th field is the (NCAT) i.e.start of the repeating sequence and the number of dimensions in a n dimensional array.
		e.g. if NCAT is 4 then there will be 16 fields in this order 4x4...
		N_CAT F1_O1 F1_O2 F1_O3 F1_O4 F2_O1 F2_O2 F2_O3 F2_O4 F3_O1 F3_O2 F3_O3 F3_O4 F4_O1 F4_O2 F4_O3 F4_O4
		Cannot depend on there always being a header line that tells us the number of dimensions in the array.
		The dimensions AND the order must be inferred from the NCAT and the knowledge that they go in sorted
		order 1st dimension then second dimension. As far as I know these are always ints
		*/
		return getNCATStructureString(cleanTerm, index)
	case "(N_THRESH)": // PCT, PJC, PRC, PSTD, PROBRIRW files
		switch lineType {
		case "PCT":
			// THRESH PCT which is a sequence of THRESH_n OY_n ON_n which are float64 values
			return getFillStructureSequenceString([]string{"THRESH_", "OY_", "ON_"}, cleanTerm, "float64", index)
		case "PJC":
			// PJC files have a sequence of THRESH_n OY_TP_n ON_TP_n CALIBRATION_n REFINEMENT_n LIKELIHOOD_n BASER_n which are float64 values
			return getFillStructureSequenceString([]string{"THRESH_", "OY_TP_", "ON_TP_", "CALIBRATION_", "REFINEMENT", "LIKELIHOOD_", "BASER_"}, cleanTerm, "float64", index)
		case "PRC":
			// PRC files have a sequence of THRESH_n PODY_n POFD_n which are float64 values
			return getFillStructureSequenceString([]string{"THRESH_", "PODY_", "POFD_"}, cleanTerm, "float64", index)
		case "PSTD":
			// PSTD files have a sequence of THRESH_n which are float64 values
			return getFillStructureSequenceString([]string{"THRESH_"}, cleanTerm, "float64", index)
		case "PROBRIRW":
			// PROBRIRW files have a sequence of THRESH_n PROB_n which are int values
			return getFillStructureSequenceString([]string{"THRESH_", "PROB_"}, cleanTerm, "int", index)
		}

	case "(N_PTS)": // ECLV files (N_PTS) for ECLV files which is a sequence of CL_n VALUE_n which are float64 values
		return getFillStructureSequenceString([]string{"CL_", "VALUE_"}, cleanTerm, "float64", index)
	case "(N_RANK)": // RHIST files (N_RANK) for RHIST files, the repeated sequence is RANK_n which are int values
		return getFillStructureSequenceString([]string{"RANK_"}, cleanTerm, "int", index)
	case "(N_BIN)": // PHIST files (N_BIN) for PHIST files, the repeated sequence is BIN_n which are int values
		return getFillStructureSequenceString([]string{"BIN_"}, cleanTerm, "int", index)
	case "(N_ENS)": // ORANK, RELP files (N_ENS) for ORANK files, the repeated sequence is ENS_n
		if fileLineType == "STAT_ORANK" {
			// (N_ENS) for ORANK files, the repeated sequence is ENS_n which are ints (can have NA values)
			return getFillStructureSequenceString([]string{"ENS_"}, cleanTerm, "int", index)
		}
		if fileLineType == "STAT_RELP" {
			// (N_ENS) for RELP files, the repeated sequence is RELP_n which are float64 values
			return getFillStructureSequenceString([]string{"RELP_"}, cleanTerm, "float64", index)
		}
	case "(N_DIAG)": // TCDIAG files (no sample data for this type)
		return getFillStructureSequenceString([]string{"DIAG_", "VALUE_"}, cleanTerm, "string", index)
	}
	return index, "", nil
}

func getFillStructureSequenceString(keyPrefixes []string, cleanTerm string, elemType string, index int) (numFields int, structureString string, err error) {
	var convStr string
	// sometimes we need to do a nilCheck (if it is a conversion to an int) - otherwise we will get a panic
	keyPrefixesStr := `"` + strings.Join(keyPrefixes, `","`) + `"`
	if elemType == "float64" {
		convStr = `value, err = strconv.ParseFloat(fields[index],64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}`
	} else if elemType == "int" {
		convStr = `value, err = strconv.Atoi(fields[index])
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}`
	} else {
		convStr = "value = fields[index]"
	}
	str := `    // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value interface{}
	count, err := strconv.Atoi(fields[%d])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{%s}
	s.%s = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := %d; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%%s_%%d",keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
				value = "NA"
			} else {
				%s
			}
			s.%s[key] = value
		}` + "\n\t}\n"

	str = fmt.Sprintf(str, index, keyPrefixesStr, cleanTerm, index+1, convStr, cleanTerm)
	return len(keyPrefixes), str, nil
}

func getNCATStructureString(cleanTerm string, index int) (numFields int, structureString string, err error) {
	// these values seem to always be ints (or "NA")
	str := `    // these values seem to always be ints (or "NA")
	var value interface{}
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	s.%s = make(map[string]interface{})
	for i1 := %d; i1 <= count; i1++ {
		for i2 := 1; i2 <= count; i2++ {
			// generate the particular key for the map i.e. F1_O1, F1_O2, F1_O3, F1_O4, F2_O1, F2_O2, F2_O3, F2_O4, etc.
			key := fmt.Sprintf("F%%d_O%%d", i1, i2)
			index := (i1-1)*count + i2
			if index >= len(fields) {
				value = "NA"
			} else {
				value, err = strconv.Atoi(fields[index])
			}
			if err != nil {
				value = "NA"
			}
			s.%s[key] = value
		}` + "\n\t}\n"

	str = fmt.Sprintf(str, cleanTerm, index, cleanTerm)
	return 1, str, nil
}

func fillMetDataMapFromSrcFiles(metDataTypesForLines map[string]string, fieldNameMap map[string]string) map[string]string {
	// use a map (atoLines) as a set to avoid duplicate lines
	atoLines := make(map[string]bool)
	// iterate through the metSrcFiles to get the data types for the fields
	matchConvertLine := regexp.MustCompile(`= ato[fi]\(l.get_item\(`)
	for _, url := range metSrcFiles {
		lines := getLinesForUrl(url)
		// iterate through the lines to find the data types
		for _, line := range lines {
			parts := matchConvertLine.Split(line, -1)
			if len(parts) > 1 {
				atoLines[line] = true
			}
		}
	}
	// iterate the fieldNames to see if we can find data types in the atoLines from the src code files
	for token := range fieldNameMap {
		// iterate through the atoLines to get any data types
		for line := range atoLines {
			if matched, err := regexp.Match(token, []byte(line)); err == nil && matched {
				parts := strings.Split(line, " =")
				if len(parts) > 1 {
					atPart := strings.TrimSpace(strings.Split(parts[1], "(")[0])
					var dataType string
					switch atPart {
					case "atoi":
						dataType = "int"
					case "atof":
						dataType = "float64"
					default:
						dataType = "string"
					}
					fieldNameMap[token] = dataType
					metDataTypesForLines[token] = dataType
				}
			}
		}
	}
	return metDataTypesForLines
}

func fillMetDataMapFromUserGuide(metDataTypesForLines, fieldNameMap map[string]string) map[string]string {
	// MET user guide files with data type definitions
	// Using the slower regexp instead of a string match because I don't know if the line will have
	// extra leading spaces or not. These documents might get reformatted and the leading spaces might
	// change. They are less likely to change the order of the fields in the table.
	// The regex to split the line into fields
	linePrefix := regexp.MustCompile(`^ *- `)
	// The regexp to identify the start of a column header
	lineColumnStart := regexp.MustCompile(`^\s*\* - Column`)
	for _, url := range metUserDocFiles {
		docFileLines := getLinesForUrl(url)
		var parts []string
		for i := 0; i < len(docFileLines)-1; i++ {
			line := docFileLines[i]
			if line == "" {
				continue
			}
			/*
					look for the lines that have data type table entries - they all start with "  * - " and
					have Column in the first line. If a line starts with "  * - " and does not have Column in it
					then it is the end of the data type table.
					* - Column Number       - start of table (this is the header)
					- Header Column Name
					- Description
					- Data Type
				* - 1                       - 1st entry of the table (this is the column number)
					- VERSION               - field name
					- Version number        - field description
					- String                - data type
			*/
			// skip the lines that are not the start of the data type table i.e. don't start with '* - Column Number'
			if strings.Contains(line, "* - Column") {
				if !strings.Contains(docFileLines[i+1], "Name") ||
					!strings.Contains(docFileLines[i+2], "Description") ||
					!strings.Contains(docFileLines[i+3], "Data Type") {
					continue // this is not the start of a table we are interested in
				}
				// process the table
				for ; i < len(docFileLines)-1; i++ {
					line = docFileLines[i]
					if line == "" {
						break // this table is ended, go to the next one.
					}
					// is this a Column Header Start line (header of table)?
					if lineColumnStart.MatchString(line) {
						// skip to the end of the table header
						i = i + 4
					}
					// skip the first line - it is the column number
					i++
					line = docFileLines[i]
					// remove possible embedded html - ugh!
					line = strings.Replace(line, ":raw-html:`<br />`", "", -1)
					line = strings.Replace(line, `\`, "", -1)
					if line == "    -" {
						break // this table is formatted poorly (empty fieldName), go to the next one.
					}
					parts = linePrefix.Split(line, -1)
					fieldName := strings.Replace(parts[1], " ", "", -1) // remove extra spaces
					parts := strings.Split(fieldName, "/")              // remove any front slashes
					if len(parts) > 1 {
						fieldName = parts[1]
					}
					// It seems that the actual fieldNames that are specified in the doc as abc_i are labeled as abc_[0-9]* in the code
					fieldName = strings.Replace(fieldName, `_i`, `_[0-9]*`, -1) // replace _i with _[0-9]
					// skip the description line
					i = i + 2
					line = docFileLines[i]
					// skip any extra line (for some reason these lines can have notes and things that take up multiple lines) that do not start with "* - "
					if !linePrefix.MatchString(line) {
						i = i + 1
						line = docFileLines[i]
					}
					parts = linePrefix.Split(line, -1)
					dataType := strings.TrimSpace(parts[1]) // get the data type
					// any unknown data types will default to string
					switch dataType {
					case "Integer":
						dataType = "int"
					case "Double":
						dataType = "float64"
					case "String":
						dataType = "string"
					default:
						dataType = "string"
					}
					fields := []string{}
					// INTENSITY fields seem to be sort of one-off fields that have partial subfields i.e. "INTENSITY_10, _25, _50, _75, _90"
					// These must be handled differently to get the actual subfields.
					if strings.Contains(fieldName, "INTENSITY") ||
						strings.Contains(fieldName, "CURVATURE") ||
						strings.Contains(fieldName, "CENTROID") {
						parts = strings.Split(fieldName, ",")
						fields = append(fields, parts[0])
						if len(parts) > 1 {
							pre := strings.Split(parts[0], "_")[0]
							for _, p := range parts[1:] {
								subFieldName := pre + strings.TrimSpace(p)
								fields = append(fields, subFieldName)
							}
						} else {
							fields = append(fields, fieldName)
						}
					} else {
						// some fields have multiple subfields i.e. "FCST_LEAD, FCST_LEV"
						// These also must be handled differently to get the actual subfields.
						majorFields := strings.Split(fieldName, ",")
						if len(majorFields) > 1 {
							for _, majorField := range majorFields {
								// remove any  remaining spaces - some fields have nonsense spaces in them
								subfield := strings.Replace(majorField, " ", "", -1)
								fields = append(fields, subfield)
							}
						} else {
							fields = append(fields, fieldName)
						}
					}
					// now fields should have all the field names including subfields
					for _, fieldName := range fields {
						if fieldName != "" {
							/*
								NOTE: many of the fieldNameMap entries remain "UNDEFINED" because the
								data type is not found in the MET user guide files. The data type
								will default to "string" in the generated code. This is the best that
								can be done without a more comprehensive data type definition. I am
								leaving the fieldNameMap entries as a way to track down
								the missing data types, but it is not used in the code generation.
							*/
							metDataTypesForLines[fieldName] = dataType
							fieldNameMap[fieldName] = dataType
						}
					}
				}
			}
		}
	}
	// Fill undefined's that I have found in the MET user guide files in text (not column tables), or in data files themselves.
	metDataTypesForLines["RIRW_WINDOW"] = "int"
	metDataTypesForLines["F[0-9]*_O[0-9]*"] = "string"
	metDataTypesForLines["INTENSITY_USER"] = "float64"
	metDataTypesForLines["INTENSITY_USER_MIN"] = "float64"
	metDataTypesForLines["INTENSITY_USER_MAX"] = "float64"
	metDataTypesForLines["RPS_COMP"] = "float64"
	metDataTypesForLines["ARADP"] = "string"
	metDataTypesForLines["AMRD"] = "int"
	metDataTypesForLines["AGUSTS"] = "int"
	metDataTypesForLines["ADIR"] = "int"
	metDataTypesForLines["AEYE"] = "int"
	metDataTypesForLines["EIQR_BCL"] = "float64"
	metDataTypesForLines["EIQR_BCU"] = "float64"
	metDataTypesForLines["ASPEED"] = "int"
	metDataTypesForLines["ARRP"] = "int"
	metDataTypesForLines["ADEPTH"] = "int"

	fieldNameMap["RIRW_WINDOW"] = "int"
	fieldNameMap["F[0-9]*_O[0-9]*"] = "string"
	fieldNameMap["INTENSITY_USER"] = "float64"
	fieldNameMap["INTENSITY_USER_MIN"] = "float64"
	fieldNameMap["INTENSITY_USER_MAX"] = "float64"
	fieldNameMap["RPS_COMP"] = "float64"
	fieldNameMap["ARADP"] = "string"
	fieldNameMap["AMRD"] = "int"
	fieldNameMap["AGUSTS"] = "int"
	fieldNameMap["ADIR"] = "int"
	fieldNameMap["AEYE"] = "int"
	fieldNameMap["EIQR_BCL"] = "float64"
	fieldNameMap["EIQR_BCU"] = "float64"
	fieldNameMap["ASPEED"] = "int"
	fieldNameMap["ARRP"] = "int"
	fieldNameMap["ADEPTH"] = "int"

	// Uncomment the following to look for missing data types in the MET user guide files.
	var found bool
	var undefineds []string = []string{}
	for k, v := range fieldNameMap {
		if v == "UNDEFINED" {
			found = false
			for _, v1 := range *getPatterns() {
				if v1.match.MatchString(strings.ToUpper(k)) {
					found = true
					break
				}
			}
			if !found {
				undefineds = append(undefineds, k)
			}
		}
	}
	if len(undefineds) > 0 {
		fmt.Printf("\n/*undefined: %v*/\n", undefineds)
	}

	return metDataTypesForLines
}

func getPatterns() *map[string]Pattern {
	// These are Regular expression patterns that are used to find the data types for the fields that are not simple,
	// linear or not found in the user guide files.
	// PROBRIRW example
	// VERSION AMODEL BMODEL DESC STORM_ID BASIN CYCLONE STORM_NAME INIT            LEAD   VALID           INIT_MASK VALID_MASK LINE_TYPE ALAT ALON  BLAT BLON  INITIALS TK_ERR    X_ERR      Y_ERR ADLAND     BDLAND     RIRW_BEG RIRW_END RIRW_WINDOW AWIND_END BWIND_BEG BWIND_END BDELTA BDELTA_MAX BLEVEL_BEG BLEVEL_END N_THRESH THRESH_1 PROB_1 THRESH_2 PROB_2 THRESH_3 PROB_3 THRESH_4 PROB_4 THRESH_5 PROB_5
	// V12.0.0 GPMI   BEST   NA   AL012015 AL    01      ANA        20150508_120000 240000 20150509_120000 NA        NA         PROBRIRW  31.6 -77.7 32.5 -77.8       NA  54.23894    5.08551   -54  135.63956   80.31028        0       24          24        44        40        50     10         10         TS         TS        5      -30      0      -10      0        0    100       10      0       30      0
	// Only define and Compile these things once
	if len(patterns) == 0 {
		// these `\(N_[A-Z]+\)` are repeated sequences of data fields - a map of data fields
		// that indicate repeated sequences of data fields in the data sections.
		// These do not have direct corresponding data elements, but they have
		// corresponding sequences of data elements in the data sections.
		// Those elements are identified in the actual data section headers by
		// things like THRESH_[0-9]* PROB_[0-9]* for PROBRIRW. There are associated data
		// elements in the data sections that are of different types and they are
		// defined by tokens like THRESH_1 PROB_1 THRESH_2 PROB_2 THRESH_3 PROB_3
		// Since these are essentially key value pairs of varying length (per individual data file)
		// they are not directly represented in the data structures as a map.
		// The data structures are created dynamically based on the data files and the types
		// of the values in the map are specific to each pattern e.g. bin_n values are ints
		// but cl_n values are float64s.
		// structField is for repeating fields only. These patterns are surrounded by "()" i.e. (N_CAT).
		// The structure generator needs to know what it is that is repeating. Most likely a map of some sort.

		// repeating patterns
		patterns["(n_cat)"] = Pattern{match: regexp.MustCompile(`(N_CAT)`), dType: "int", structField: "CAT", structType: "map[string]interface{}"}
		patterns["(n_thresh)"] = Pattern{match: regexp.MustCompile(`(N_THRESH)`), dType: "int", structField: "THRESH", structType: "map[string]interface{}"}
		patterns["(n_pts)"] = Pattern{match: regexp.MustCompile(`(N_PTS)`), dType: "int", structField: "PTS", structType: "map[string]interface{}"}
		patterns["(n_ens)"] = Pattern{match: regexp.MustCompile(`(N_ENS)`), dType: "int", structField: "ENS", structType: "map[string]interface{}"}
		patterns["(n_rank)"] = Pattern{match: regexp.MustCompile(`(N_RANK)`), dType: "int", structField: "RANK", structType: "map[string]interface{}"}
		patterns["(n_bin)"] = Pattern{match: regexp.MustCompile(`(N_BIN)`), dType: "int", structField: "BIN", structType: "map[string]interface{}"}
		patterns["(n_diag)"] = Pattern{match: regexp.MustCompile(`(N_DIAG)`), dType: "int", structField: "DIAG", structType: "map[string]interface{}"}
		// single patterns
		patterns["baser_n"] = Pattern{match: regexp.MustCompile("BASER_[0-9]*"), dType: "float64", structField: "BASER_I", structType: "map[string]interface{}"}
		patterns["bin_n"] = Pattern{match: regexp.MustCompile("BIN_[0-9]*"), dType: "int", structField: "BIN_I", structType: "int"}
		patterns["calibration_n"] = Pattern{match: regexp.MustCompile("CALIBRATION_[0-9]*"), dType: "float64", structField: "CALIBRATION_I", structType: "float64"}
		patterns["cl_n"] = Pattern{match: regexp.MustCompile("CL_[0-9]*"), dType: "float64", structField: "CL_I", structType: "float64"}
		patterns["diag_n"] = Pattern{match: regexp.MustCompile("DIAG_[0-9]*"), dType: "float64", structField: "DIAG_I", structType: "float64"}
		patterns["ens_n"] = Pattern{match: regexp.MustCompile("ENS_[0-9]*"), dType: "int", structField: "ENS_I", structType: "int"}
		patterns["fi_oi"] = Pattern{match: regexp.MustCompile("F[0-9]*_O[0-9]*"), dType: "string", structField: "FI_OI", structType: "string"}
		patterns["azfi_azoi"] = Pattern{match: regexp.MustCompile("[A-Z]F[0-9]*_[A-Z]O[0-9]*"), dType: "string", structField: "AZFI_AZOI", structType: "string"}
		patterns["likelihood_n"] = Pattern{match: regexp.MustCompile("LIKELIHOOD_[0-9]*"), dType: "float64", structField: "LIKELIHOOD_I", structType: "float64"}
		patterns["aal_wind_n"] = Pattern{match: regexp.MustCompile("AAL_WIND_[0-9]*"), dType: "float64", structField: "AAL_WIND_I", structType: "float64"}
		patterns["ase_wind_n"] = Pattern{match: regexp.MustCompile("ASE_WIND_[0-9]*"), dType: "float64", structField: "ASE_WIND_I", structType: "float64"}
		patterns["asw_wind_n"] = Pattern{match: regexp.MustCompile("ASW_WIND_[0-9]*"), dType: "float64", structField: "ASW_WIND_I", structType: "float64"}
		patterns["ane_wind_n"] = Pattern{match: regexp.MustCompile("ANE_WIND_[0-9]*"), dType: "float64", structField: "ANE_WIND_I", structType: "float64"}
		patterns["anw_wind_n"] = Pattern{match: regexp.MustCompile("ANW_WIND_[0-9]*"), dType: "float64", structField: "ANW_WIND_I", structType: "float64"}
		patterns["on_tp_n"] = Pattern{match: regexp.MustCompile("ON_TP_[0-9]*"), dType: "float64", structField: "ON_TP_I", structType: "float64"}
		patterns["on_n"] = Pattern{match: regexp.MustCompile("ON_[0-9]*"), dType: "float64", structField: "ON_I", structType: "float64"}
		patterns["oy_tp_n"] = Pattern{match: regexp.MustCompile("OY_TP_[0-9]*"), dType: "float64", structField: "OY_TP_I", structType: "float64"}
		patterns["oy_n"] = Pattern{match: regexp.MustCompile("OY_[0-9]*"), dType: "float64", structField: "OY_I", structType: "float64"}
		patterns["pody_n"] = Pattern{match: regexp.MustCompile("PODY_[0-9]*"), dType: "float64", structField: "PODY_I", structType: "float64"}
		patterns["pofd_n"] = Pattern{match: regexp.MustCompile("POFD_[0-9]*"), dType: "float64", structField: "POFD_I", structType: "float64"}
		patterns["prob_n"] = Pattern{match: regexp.MustCompile("PROB_[0-9]*"), dType: "float64", structField: "PROB_I", structType: "float64"}
		patterns["rank_n"] = Pattern{match: regexp.MustCompile("RANK_[0-9]*"), dType: "int", structField: "RANK_I", structType: "int"}
		patterns["refinement_n"] = Pattern{match: regexp.MustCompile("REFINEMENT_[0-9]*"), dType: "float64", structField: "REFINEMENT_I", structType: "float64"}
		patterns["relp_n"] = Pattern{match: regexp.MustCompile("RELP_[0-9]*"), dType: "float64", structField: "RELP_I", structType: "float64"}
		patterns["thresh_n"] = Pattern{match: regexp.MustCompile("THRESH_[0-9]*"), dType: "int", structField: "THRESH_I", structType: "int"}
		patterns["value_n"] = Pattern{match: regexp.MustCompile("VALUE_[0-9]*"), dType: "int", structField: "VALUE_I", structType: "int"}
	}
	return &patterns
}

func getColumnLinesAndMapForUrl(fileUrl string) ([]string, map[string]string) {
	fieldNameMap := make(map[string]string)
	lines := getLinesForUrl(fileUrl)
	// split out all the fields to get a map of required fields
	for _, line := range lines {
		// get the prefix from the line
		parts := strings.Split(line, ": VERSION")
		if len(parts) < 2 {
			if line != "" {
				fmt.Println("error parsing line: met_header_columns" + "line:'" + line + "'")
			}
			continue
		}
		// get all the required fields for this line
		fields := strings.Fields(parts[1])
		// use a map as a SET to get unique field names
		// this makes an assumption the field names are unique or that
		// the field names have the same data type all the different structs i.e. columnDef lines.
		for _, field := range fields {
			name := strings.ToLower(field)
			// fieldNameMap[cases.Title(language.English).String(name)] = "UNDEFINED"
			fieldNameMap[strings.ToUpper(name)] = "UNDEFINED"
		}
	}
	return lines, fieldNameMap
}

func getLinesForUrl(fileUrl string) []string {
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("error getting met_header_columns file" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rawColumnsBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
		os.Exit(1)
	}
	lines := strings.Split(string(rawColumnsBytes), "\n")
	return lines
}

func getDataType(name string, metDataTypesLines *map[string]string) (key string, dType string) {
	/*
		The getDataType function is used to get the data type for a field name. The field name is used to look up the data type in the metDataTypesLines map.
		If the data type is not found in the map then the data type is set to "string". The field name is returned as the key and the data type is returned as the dType.
		There are some fields that are arrays of values and these are handled differently. The field name is modified to reflect the array of values and the data type is set to the array type.
		Some interpretations of the patterns depend upon the line type.
	*/
	//uName := cases.Title(language.English).String(name)
	uName := strings.ToUpper(name)
	// is it exactly the same as a pattern? This is the case for a structure repeating field.
	for _, v := range *getPatterns() {
		if v.match.String() == uName {
			return v.structField, v.structType
		}
	}
	// does it match a pattern? This is the case for a structure simple field.
	for _, v := range *getPatterns() {
		if v.match.MatchString(uName) {
			return uName, v.dType
		}
	}

	// it wasn't one of the patterns so try to look it up in the metDataTypesLines
	dataType := (*metDataTypesLines)[uName]
	// it could still be undefined - if it is then set it to string
	if dataType == "" {
		dataType = "string"
	}
	return uName, dataType
}
