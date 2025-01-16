package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
	"parser/pkg/buildHeaderLineTypeUtilities"
	"regexp"
	"strings"
)

/*
The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types.
*/

func getDataTypeFromCustomFieldNameMap(name string, metDataTypesLines *map[string]string, customFieldNameMap *map[string]string) string {
	/*
		The getDataTypeFromCustomFieldNameMap function is used to get the data type from the customFieldNameMap.
		if the data type is not found in the metDataTypesLines map. The customFieldNameMap is a map of field names
		(defined by a regular expression) to data types that are not found in the MET user guide files.
	*/
	uName := strings.ToUpper(name)
	dataType := (*metDataTypesLines)[strings.ToUpper(uName)]
	if dataType != "" {
		return dataType
	}
	for regExpression, dataType := range *customFieldNameMap {
		regex, err := regexp.Compile(regExpression)
		if err != nil {
			continue
		}
		if regex.MatchString(uName) {
			if dataType == "UNDEFINED" {
				return "string"
			} else {
				return dataType
			}
		}
	}
	return ""
}

func main() {
	// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
	// currently 12.0 to get the required header column definitions and then using
	// using the https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide
	// to get the stat field types. We create a map of field names to field types. Then we have to re-iterate
	// over the header definitions to create the structs and functions to fill the structs.
	type HeaderStructName string
	type HeaderStructs map[HeaderStructName]string
	type DataStructs map[string]string

	/* The following is a map of regexp field names to data types.
	These are the fields that are not found in the MET user guide files.
	These fields would default to "string" in the generated code so
	this is a mechanism to custom define the data types for these fields.
	The code generator will first try to use a field in the metDataTypesLines map
	and if it is not found it will use the customFieldNameMap map. If the field is not found
	in either map it will default to "string". Some of these remain "UNDEFINED" because
	the data type is not found in the MET user guide files and I don't have a better way to
	determine the data type.
	*/
	customFieldNameMap := map[string]string{
		`\d_DIAG`:        "int",
		`.*_WIND\d*`:     "float64",
		`\d_RAN`:         "int",
		`AMRD`:           "UNDEFINED",
		`.*MEAN_*`:       "float64",
		`F[0-9]*O[0-9]*`: "float64",
		`O[0-9]*O[0-9]*`: "float64",
		`RIRW_BE`:        "int",
		`RIRW_EN`:        "int",
		`RIRW_WIDOW`:     "int",
		`INTENSITY_USER`: "string",
		`ARADP`:          "UNDEFINED",
		`\d_ENS`:         "float64",
		`OENERGY`:        "float64",
		`FENERGY`:        "float64",
		`\d_THRSH`:       "float64",
		`\d_BIN`:         "int",
		`ARRP`:           "UNDEFINED",
		`RPS_COM`:        "UNDEFINED",
		`ASPEED`:         "float64",
		`ADIR`:           "UNDEFINED",
		`BIN_i`:          "int",
		`AGUSTS`:         "UNDEFINED",
		`\d_CAT`:         "float64",
		`ADEPTH`:         "UNDEFINED",
		`\d_PTS`:         "UNDEFINED",
		`EIQR_.*`:        "int",
		`AEYE`:           "UNDEFINED",
	}

	// Using the slower regexp instead of a string match because I don't know if the line will have
	// extra leading spaces or not. These documents might get reformatted and the leading spaces might
	// change. They are less likely to change the order of the fields in the table.
	// The regex to split the line into fields
	linePrefix, _ := regexp.Compile(`^ *- `)
	// The regexp to identify the start of a column header
	lineColumnStart, _ := regexp.Compile(`^\s*\* - Column`)
	// read the header columns file
	headerColumnsFileUrl := "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
	resp, err := http.Get(headerColumnsFileUrl)
	if err != nil {
		fmt.Println("error getting met_header_columns file" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rawColumnsBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading met_header_columns file" + err.Error())
		os.Exit(1)
	}

	met_header_columns_lines := strings.Split(string(rawColumnsBytes), "\n")
	fieldNameMap := map[string]string{}
	// split out all the fields to get a map of required fields
	for _, line := range met_header_columns_lines {
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

	// MET user guide files with data type definitions
	var metDataTypesLines map[string]string = make(map[string]string)
	metUserDocFiles := []string{"ensemble-stat.rst", "grid-stat.rst", "gsi-tools.rst", "mode-td.rst", "mode.rst", "point-stat.rst", "stat-analysis.rst", "tc-gen.rst", "tc-pairs.rst", "wavelet-stat.rst"}
	for _, file := range metUserDocFiles {
		url := fmt.Sprintf("https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/%s", file)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error getting met user guide file" + err.Error())
			os.Exit(1)
		}
		defer resp.Body.Close()
		rawBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error reading met user guide file" + err.Error())
			os.Exit(1)
		}
		docFileLines := strings.Split(string(rawBytes), "\n")
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
							metDataTypesLines[fieldName] = dataType
							fieldNameMap[fieldName] = dataType
						}
					}
				}
			}
		}
	}
	// Uncomment the following to look for missing data types in the MET user guide files.
	// for k, v := range fieldNameMap {
	// 	if v == "UNDEFINED" {
	// 		fmt.Printf("%s %s\n", k, v)
	// 	}
	// }

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
		// create the getDocForId case statement element

		// create the data struct for this line type
		dataStruct := fmt.Sprintf("type %s struct {\n", docStructName)
		// find the maximum length of the data fields for formatting (padding)
		padding = 0
		for _, term := range dataFields {
			if len(term) > padding {
				padding = len(term)
			}
		}

		for index, term := range dataFields {
			term = strings.Replace(term, "(", "", -1)
			term = strings.Replace(term, ")", "", -1)
			term = strings.Replace(term, "[0-9]*", "i", -1)
			name := strings.ToLower(term)
			capName := cases.Title(language.English).String(name)
			dataType := getDataTypeFromCustomFieldNameMap(name, &metDataTypesLines, &customFieldNameMap)

			dataStruct += fmt.Sprintf("    %-*s %-*s `json:\"%s\"`\n", padding, capName, 7, dataType, name)
			switch dataType {
			case "int":
				fillStructureString += fmt.Sprintf("\ts.%s, _ = strconv.Atoi(fields[%d])\n", capName, index)
			case "float64":
				fillStructureString += fmt.Sprintf("\ts.%s, _ = strconv.ParseFloat(fields[%d], 64)\n", capName, index)
			default:
				fillStructureString += fmt.Sprintf("\ts.%s = fields[%d]\n", capName, index)
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
