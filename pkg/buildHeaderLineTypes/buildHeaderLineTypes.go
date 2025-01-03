package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
	"parser/pkg/buildHeaderLineTypeUtilities"
	"strings"
)

// The output of this program is a series of structs that can be used to define the header
// and data types in the buildHeaderTypes.go file
func main() {
	// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
	// currently 12.0 to get the header column definitions and then using
	// using the src/tools/core/stat_analysis/parse_stat_line.cc to get the stat field types.
	type HeaderStructName string
	type HeaderStructs map[HeaderStructName]string
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
	getDocIDString := "func GetDocForId(fileLineType string, metaData VxMetadata, headerData []string, dataData []string, dataKey string) (interface{}, error) {\n\tdoc := make(map[string] interface{})\n\tswitch fileLineType {\n"
	addDataElementString := "func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) error {\n\tswitch fileLineType {\n"
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
		headerFields, dataFields := buildHeaderLineTypeUtilities.SplitColumnDefLine(fileType, fieldStr)

		// create the header struct
		docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
		headerStructName := HeaderStructName(fmt.Sprintf("%s_header", docStructName))

		keyFields := buildHeaderLineTypeUtilities.GetKeyDataFieldsForLineType(fileType)
		dataKeyFieldsString := strings.Join(keyFields, ", ")
		dataKeyMapString += fmt.Sprintf("    \"%s\": %*s\"%s\"},\n", docStructName, lineTypePadding, "{", dataKeyFieldsString)

		headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
		fillHeaderString := fmt.Sprintf("func (s *%s) fill_%s_Header(fields []string, doc *map[string]interface{}){\n", docStructName, docStructName)
		fillHeaderString += "// fill the met fields\n"
		getDocIDString += fmt.Sprintf("\tcase \"%s\":\n", docStructName)
		getDocIDString += fmt.Sprintf("\t\telem := %s{}\n\t\tdoc[\"ID\"] = metaData.ID\n\t\tdoc[\"Subset\"] = metaData.Subset\n\t\tdoc[\"Type\"] = metaData.Type\n\t\tdoc[\"SubType\"] = metaData.SubType\n", docStructName)

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
			// get the dataType from the cc file
			dataType := "string"
			name := strings.ToLower(term)
			capName := cases.Title(language.English).String(name)
			headerStructString += fmt.Sprintf("    %-*s %s `json:\"%s\"`\n", padding, capName, dataType, name)
			fillHeaderString += fmt.Sprintf("\t(*doc)[\"%s\"] = fields[%d]\n", name, index)
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
			foundType, err := buildHeaderLineTypeUtilities.FindType(term, rawStatBytes)
			if err == nil {
				dataType = foundType
			}
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

	// We have to flesh this out with the vxMetadata struct
	vxMetaDataStruct := fmt.Sprintf("type VxMetadata struct {\n"+
		"    %-*s string `json:\"id\"`\n"+
		"    %-*s string `json:\"subset\"`\n"+
		"    %-*s string `json:\"type\"`\n"+
		"    %-*s string `json:\"subtype\"`\n"+
		"}\n", 7, "ID", 7, "Subset", 7, "Type", 7, "SubType")

	fmt.Println("package structColumnTypes")
	fmt.Println("")
	fmt.Println("import (\n\t\"strconv\"\n\t\"errors\"\n)")
	fmt.Println("\n/*\nTHIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE")
	fmt.Println("To modify this code - modify the buildHeaderLineTypes.go file and run the buildHeaderLineTypes.go program")
	fmt.Println("cd  <repo path>/metlinetypes/pkg/buildHeaderLineTypes")
	fmt.Println("go run . > /tmp/types.go")
	fmt.Println("cp /tmp/types.go ../structColumnTypes/structColumnTypes.go\n*/")
	fmt.Println("")

	// print the vxMetadata struct
	fmt.Println("//vxMetadata struct definition")
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

	// print the addDataElement functions
	fmt.Println("")
	fmt.Println("//addDataElement functions")
	fmt.Println(addDataElementString)

	// print the dataKeys
	fmt.Println("")
	fmt.Println(`	// data key definitions -  these may need to be changed based on the data and how we want to associate records into a single document
	// All of the records with the same header values will be grouped into a single document
	// The key is used to group the data records into the single document data map and it will be indexed by these values as a key.`)
	fmt.Println(dataKeyMapString)
	fmt.Println("}")

	// print the DateFieldNames
	fmt.Println("")
	fmt.Println("	// DateFieldNames is a list of the date fields that need to be converted to epochs")
	fmt.Println("var DateFieldNames = []string{\"FCST_VALID_BEG\", \"FCST_VALID_END\", \"OBS_VALID_BEG\", \"OBS_VALID_END\"}")
	fmt.Println("")
}
