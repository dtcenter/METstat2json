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
)

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
		headerString = parts[0] + " OBS_LEV"
		// get the data fields from line
		dataString = parts[1]
	default:
		// get the header fields from line
		parts = strings.Split(fieldStr, " LINE_TYPE ")
		headerString = parts[0] + " LINE_TYPE"
		// get the data fields from line
		dataString = parts[1]
	}
	// squeeze the spaces out of the header string
	headerFields := strings.Fields(headerString)
	// get the data fields from line
	dataFields := strings.Fields(dataString)
	return headerFields, dataFields
}

/*
This routine is used to get the key data fields for a given line type. These fields are used to
essentially merge documents. All the documents with the same header field values excluding the key data fields
are merged into a single document. The key data fields are used to index the data section of the document.
*/
func getKeyDataFieldsForLineType(lineType string) []string {
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

func findType(name string, data []byte) (string, error) {
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

// The output of this program is a series of structs that can be used to define the header
// and data types in the buildHeaderTypes.go file
func main() {
	// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
	// currently 12.0 to get the header column definitions and then using
	// using the src/tools/core/stat_analysis/parse_stat_line.cc to get the stat field types.
	type HeaderStructName string
	type HeaderStructs map[HeaderStructName]string
	type DocStructs map[string]string
	type DataStructs map[string]string

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

	// read the parse_stat_line file
	parse_stat_line_cc := "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/core/stat_analysis/parse_stat_line.cc"
	resp, err = http.Get(parse_stat_line_cc)
	if err != nil {
		fmt.Println("error getting parse_stat_line file" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rawStatBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading parse_stat_line file " + err.Error())
		os.Exit(1)
	}

	// Use a map to keep track of unique headerStructs.
	dataStructs := make(DataStructs)
	headerStructs := make(HeaderStructs)

	// fillHeaderFuncs is the same as headerStructs but for the fillHeader function
	fillHeaderFuncs := make(HeaderStructs)
	fillDataFuncs := make(DataStructs)

	dataKeyMapString := "var DataKeyMap = map[string][]string{\n"
	// create the getDocFoID function
	getDocIDString := fmt.Sprintf("func GetDocForId(fileLineType string, metaData VxMetadata, headerData []string, dataData []string, dataKey string) (interface{}, error) {\n\tswitch fileLineType {\n")

	file_lines := strings.Split(string(rawColumnsBytes), "\n")
	for _, line := range file_lines {
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
		lineTypePadding := 10 - len(lineType)
		headerFields, dataFields := SplitColumnDefLine(fileType, fieldStr)

		// create the header struct
		docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
		headerStructName := HeaderStructName(fmt.Sprintf("%s_header", docStructName))

		keyFields := getKeyDataFieldsForLineType(fileType)
		dataKeyFieldsString := strings.Join(keyFields, ", ")
		dataKeyMapString += fmt.Sprintf("    \"%s\": %*s\"%s\"},\n", docStructName, lineTypePadding, "{", dataKeyFieldsString)

		headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
		fillHeaderString := fmt.Sprintf("func (s *%s) fill_%s_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {\n", docStructName, docStructName)
		fillHeaderString += "// fill the met fields\n" +
			"(*doc)[\"ID\"] = metaData.ID\n" +
			"(*doc)[\"Subset\"] = metaData.Subset\n" +
			"(*doc)[\"Type\"] = metaData.Type\n" +
			"(*doc)[\"SubType\"] = metaData.SubType\n"
		getDocIDString += fmt.Sprintf("\tcase \"%s\":\n", docStructName)
		getDocIDString += fmt.Sprintf("\t{\n")
		getDocIDString += fmt.Sprintf("\t\telem := %s{}\n", docStructName)
		getDocIDString += fmt.Sprintf("\t\telem.fill_%s_Header(data)\n", docStructName)
		getDocIDString += fmt.Sprintf("\t\tif exists := (*doc)[\"data\"]; exists == nil {\n")
		getDocIDString += fmt.Sprintf("\t\t\t(*doc)[\"data\"] = make(map[string]%s)\n", docStructName)
		getDocIDString += fmt.Sprintf("\t\t}\n")
		getDocIDString += fmt.Sprintf("\t\tif val, ok := (*doc)[\"data\"].(map[string]%s); ok {\n", docStructName)
		getDocIDString += fmt.Sprintf("\t\t\tval[dataKey] = elem\n")
		getDocIDString += fmt.Sprintf("\t\t\t(*doc)[\"data\"] = val\n")
		getDocIDString += fmt.Sprintf("\t\t}\n}\n")

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
			// get the dataType from the cc file
			dataType := "string"
			name := strings.ToLower(term)
			capName := cases.Title(language.English).String(name)
			headerStructString += fmt.Sprintf("    %-*s %s `json:\"%s\"`\n", padding, capName, dataType, name)
			fillHeaderString += fmt.Sprintf("(*doc)[\"%s\"] = fields[%d]\n", name, index)
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
			dataType := "string"
			foundType, err := findType(term, rawStatBytes)
			if err == nil {
				dataType = foundType
			}
			dataStruct += fmt.Sprintf("    %-*s %-*s `json:\"%s\"`\n", padding, capName, 7, dataType, name)
			switch dataType {
			case "int":
				fillStructureString += fmt.Sprintf("s.%s, _ = strconv.Atoi(fields[%d])\n", capName, index)
			case "float64":
				fillStructureString += fmt.Sprintf("s.%s, _ = strconv.ParseFloat(fields[%d], 64)\n", capName, index)
			default:
				fillStructureString += fmt.Sprintf("s.%s = fields[%d]\n", capName, index)
			}
		}
		dataStruct += "}\n"
		fillStructureString += "}\n"
		dataStructs[docStructName] = dataStruct
		fillDataFuncs[docStructName] = fillStructureString
	}
	getDocIDString += fmt.Sprintf("\tdefault:\n")
	getDocIDString += fmt.Sprintf("return errors.New(\"Unknonwn file_line type:\" + fileLineType)\n")
	getDocIDString += fmt.Sprintf("\t}\n")
	getDocIDString += fmt.Sprintf("\t\treturn nil")
	getDocIDString += fmt.Sprintf("}\n")

	// We have to flesh this out with the vxMetadata struct
	vxMetaDataStruct := fmt.Sprintf("type VxMetadata struct {\n"+
		"    %-*s string `json:\"id\"`\n"+
		"    %-*s string `json:\"subset\"`\n"+
		"    %-*s string `json:\"type\"`\n"+
		"    %-*s string `json:\"subtype\"`\n"+
		"}\n", 7, "ID", 7, "Subset", 7, "Type", 7, "SubType")

	fmt.Println("package structColumnTypes")
	fmt.Println("")
	fmt.Println("import\n(\n\"strconv\"\n\"errors\"\n)")
	fmt.Println("//vxMetadata struct definition")
	// print the vxMetadata struct
	fmt.Println(vxMetaDataStruct)

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

	// print the dataKeys
	fmt.Println("")
	fmt.Println("//data key definitions")
	fmt.Println(dataKeyMapString)
	fmt.Println("}")
}
