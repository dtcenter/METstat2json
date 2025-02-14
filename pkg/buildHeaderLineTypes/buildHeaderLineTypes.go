package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"parser/pkg/buildHeaderLineTypeUtilities"
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

(N_ENS) for RELP files, the repeated sequence is N_ENS that will be contained in a map[string]interface{}
where the keys are RELP_N e.g. RELP_1, RELP_2, etc.

(N_DIAG) for TCDIAG files, the repeated sequence is DIAG_n that will be contained in a map[string]interface{}
where the keys are DIAG_n and VALUE_n e.g. DIAG_1, VALUE_1, DIAG_2, VALUE_2 etc.
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

var metHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"

/*
The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types.
*/

func main() {
	type HeaderStructName string
	type HeaderStructs map[HeaderStructName]string
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

	// Use a map to keep track of unique headerStructs.
	dataStructs := make(DataStructs)
	headerStructs := make(HeaderStructs)

	// fillHeaderFuncs is the same as headerStructs but for the fillHeader function
	fillHeaderFuncs := make(HeaderStructs)
	fillDataFuncs := make(DataStructs)

	// create the getDocFoID function
	getDocIDString := "func GetDocForId(fileLineType string, metaDataMap map[string]interface{}, headerData []string, dataData []string, dataKey string) (interface{}, error) {\n\tdoc := make(map[string] interface{})\n\t" +
		"// add the metadata to the doc\n\tfor key, value := range metaDataMap {\n\t\tdoc[key] = value\n\t}\n\tswitch fileLineType {\n"

	addDataElementString := "func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) error {\n\tswitch fileLineType {\n"

	// iterate through every line in the met_header_columns file to create the structs and functions
	// for each line type. The line type is the combination of the fileType and the lineType.
	// The lineType is the last part of the prefix in the line.
	// The fileType is the second part of the prefix in the line.
	// We will create a header struct and a data struct for each line type.
	// This is very much assuming that the column definitions are consistent across the different
	// file types and line types and that the column definitions are complete and correct.
	for _, line := range met_header_columns_lines {
		// get the prefix from the line
		parts := strings.Split(line, ": VERSION")
		if len(parts) < 2 {
			if line != "" {
				fmt.Println("error parsing line: met_header_columns" + "line:'" + line + "'")
			}
			continue
		}
		prefix := parts[0]
		fieldStr := "VERSION " + parts[1]
		// get the version from the line
		parts = strings.Split(prefix, " : ")
		if len(parts) < 3 {
			fmt.Println("error parsing line: " + line)
			continue
		}
		fileType := strings.ToUpper(strings.TrimSpace(parts[1]))
		lineType := strings.ToUpper(strings.TrimSpace(parts[2]))
		headerFields, dataFields := buildHeaderLineTypeUtilities.SplitColumnDefLine(fileType, fieldStr)

		// create the header struct
		docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
		headerStructName := HeaderStructName(fmt.Sprintf("%s_header", docStructName))

		keyFields := buildHeaderLineTypeUtilities.GetKeyDataFieldsForLineType(fileType)
		headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
		fillHeaderString := fmt.Sprintf("func (s *%s) fill_%s_Header(fields []string, doc *map[string]interface{}){\n", docStructName, docStructName)
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
		for index, term := range headerFields {
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
			name := strings.ToUpper(term)
			dataType := "string"
			capName := cases.Title(language.English).String(name)
			headerStructString += fmt.Sprintf("    %-*s %s `json:\"%s\"`\n", padding, capName, dataType, name)
			fillHeaderString += fmt.Sprintf("\tif fields[%d] != \"\"  && fields[0] != \"NA\" {\n\t\t(*doc)[\"%s\"] = fields[%d]\n\t}\n", index, name, index)
		}
		headerStructString += "}\n"
		fillHeaderString += "}\n"
		// add the header struct to the map for printing later
		headerStructs[headerStructName] = headerStructString
		fillHeaderFuncs[headerStructName] = fillHeaderString
		// create a fillStructure function for the data struct
		fillStructureString := fmt.Sprintf("func (s *%s) fill_%s(fields []string) {\n", docStructName, docStructName)
		// create the data struct for this line type
		dataStruct := fmt.Sprintf("type %s struct {\n", docStructName)
		// find the maximum length of the data fields for formatting (padding)
		padding = 0
		for _, term := range dataFields {
			if len(term) > padding {
				padding = len(term)
			}
		}

		padding2 := len("map[string]interface{}")
		for index, term := range dataFields {
			cleanTerm, dataType := getDataType(term, &metDataTypesForLines)
			dataStruct += fmt.Sprintf("    %-*s %-*s `json:\"%s\"`\n", padding, cleanTerm, padding2, dataType, cleanTerm)
			switch dataType {
			case "int":
				fillStructureString += fmt.Sprintf("\ts.%s, _ = strconv.Atoi(fields[%d])\n", cleanTerm, index)
			case "float64":
				fillStructureString += fmt.Sprintf("\ts.%s, _ = strconv.ParseFloat(fields[%d], 64)\n", cleanTerm, index)
			case "map[string]int":
				fillStructureString += fmt.Sprintf("\ts.%s = make(map[string]int)\n", cleanTerm)
			case "map[string]float64":
				fillStructureString += fmt.Sprintf("\ts.%s = make(map[string]float64)\n", cleanTerm)
			case "map[string]interface{}":
				fillStructureString += fmt.Sprintf("\ts.%s = make(map[string]interface{})\n", cleanTerm)
			default:
				fillStructureString += fmt.Sprintf("\ts.%s = fields[%d]\n", cleanTerm, index)
			}
		}
		dataStruct += "}\n"
		fillStructureString += "}\n"
		dataStructs[docStructName] = dataStruct
		fillDataFuncs[docStructName] = fillStructureString
	}
	getDocIDString += "\tdefault:\n\t\treturn nil, errors.New(\"Unknown file_line type:\" + fileLineType)\n\t}\n\treturn doc, nil\n}\n"
	addDataElementString += "\tdefault:\n\t\treturn errors.New(\"Unknown file_line type:\" + fileLineType)\n\t}\n\treturn nil\n}\n"

	fmt.Println("package structColumnTypes")
	fmt.Println("")
	fmt.Println("import (\n\t\"strconv\"\n\t\"errors\"\n)")
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
	fmt.Println(getDocIDString)

	// print the addDataElement functions
	fmt.Println("")
	fmt.Println("//addDataElement functions")
	fmt.Println(addDataElementString)

	// print the DateFieldNames
	fmt.Println("")
	fmt.Println("	// DateFieldNames is a list of the date fields that need to be converted to epochs")
	fmt.Println("var DateFieldNames = []string{\"FCST_VALID_BEG\", \"FCST_VALID_END\", \"OBS_VALID_BEG\", \"OBS_VALID_END\"}")
	fmt.Println("")
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
					dataType := "string"
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
				if v1.match.MatchString(k) {
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
		patterns["(n_cat)"] = Pattern{match: regexp.MustCompile(`(N_CAT)`), dType: "int", structField: "CAT", structType: "map[string]int"}
		patterns["(n_thresh)"] = Pattern{match: regexp.MustCompile(`(N_THRESH)`), dType: "int", structField: "THRESH", structType: "map[string]interface{}"}
		patterns["(n_pts)"] = Pattern{match: regexp.MustCompile(`(N_PTS)`), dType: "int", structField: "PTS", structType: "map[string]interface{}"}
		patterns["(n_ens)"] = Pattern{match: regexp.MustCompile(`(N_ENS)`), dType: "int", structField: "ENS", structType: "map[string]interface{}"}
		patterns["(n_rank)"] = Pattern{match: regexp.MustCompile(`(N_RANK)`), dType: "int", structField: "RANK", structType: "map[string]int"}
		patterns["(n_bin)"] = Pattern{match: regexp.MustCompile(`(N_BIN)`), dType: "int", structField: "BIN", structType: "map[string]int"}
		patterns["(n_diag)"] = Pattern{match: regexp.MustCompile(`(N_DIAG)`), dType: "int", structField: "DIAG", structType: "map[string]interface{}"}
		// single patterns
		patterns["baser_n"] = Pattern{match: regexp.MustCompile("BASER_[0-9]*"), dType: "float64", structField: "BASER_I", structType: "map[string]float64"}
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
	var fieldNameMap = make(map[string]string)
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
			fieldNameMap[field] = "UNDEFINED"
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
			return name, v.dType
		}
	}

	// it wasn't one of the patterns so try to look it up in the metDataTypesLines
	dataType := (*metDataTypesLines)[strings.ToUpper(uName)]
	// it could still be undefined - if it is then set it to string
	if dataType == "" {
		dataType = "string"
	}
	return name, dataType
}
