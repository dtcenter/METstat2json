package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getKeyDataFieldsForLineType(lineType string) []string {
	lineTypeUpper := strings.ToUpper(lineType)
	switch lineTypeUpper {
	case "MODE":
		// I have no idea if this is the correct key for MODE!
		return []string{"Fcst_lead", "Fcst_lev"}
	default:
		return []string{"Fcst_lead"}
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
	docStructs := make(DocStructs)

	var dataKeyMapString = "var dataKeyMap = map[string][]string{\n"

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
		postfix := "VERSION " + parts[1]
		// get the version from the line
		parts = strings.Split(prefix, " : ")
		if len(parts) < 3 {
			fmt.Println("error parsing line: " + line)
			continue
		}
		fileType := strings.ToUpper(strings.TrimSpace(parts[1]))
		lineType := strings.ToUpper(strings.TrimSpace(parts[2]))
		// MODE files are apparently parsed differently than other files by separating the header and data sections with OBTYPE
		var HeaderString string
		if fileType == "MODE" {
			parts = strings.Split(postfix, " OBTYPE ")
			HeaderString = parts[0] + " OBTYPE"
		} else {
			// get the header fields from line
			parts = strings.Split(postfix, " LINE_TYPE ")
			HeaderString = parts[0] + " LINE_TYPE"
		}
		// squeeze the spaces out of the header string
		headerFields := strings.Fields(HeaderString)
		dataString := parts[1]
		// get the data fields from line
		dataFields := strings.Fields(dataString)
		// squeeze the spaces out of the data string
		// create the header struct
		docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
		headerStructName := HeaderStructName(fmt.Sprintf("%s_header", docStructName))

		keyFields := getKeyDataFieldsForLineType(fileType)
		dataKeyFieldsString := strings.Join(keyFields, ", ")
		dataKeyMapString += fmt.Sprintf("    \"%s\":  {\"%s\"},\n", docStructName, dataKeyFieldsString)

		headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
		for _, term := range headerFields {
			// change regex type terms
			term = strings.Replace(term, "(", "", -1)
			term = strings.Replace(term, ")", "", -1)
			term = strings.Replace(term, "[0-9]*", "i", -1)
			// get the dataType from the cc file
			dataType := "string"
			name := strings.ToLower(term)
			capName := cases.Title(language.English).String(name)
			headerStructString += fmt.Sprintf("    %s %s `json:\"%s\"`\n", capName, dataType, name)
		}
		headerStructString += "}\n"
		// add the header struct to the map for printing later
		headerStructs[headerStructName] = headerStructString
		// add the docStruct to the map for printing later
		docStructString := fmt.Sprintf("type %s_doc struct {\n    VxMetadata      `json:\"vx_metadata\"`\n    %s_header `json:\"met_header\"`\n    Data map[string]%s `json:\"data\"`\n}", docStructName, docStructName, docStructName)
		docStructs[docStructName] = docStructString
		// create the key type for this data struct
		// create the data struct for this line type
		dataStruct := fmt.Sprintf("type %s struct {\n", docStructName)
		for _, term := range dataFields {
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
			dataStruct += fmt.Sprintf("    %s %s `json:\"%s\"`\n", capName, dataType, name)
		}
		dataStruct += "}\n"
		dataStructs[docStructName] = dataStruct
	}

	// We have to flesh this out with the vxMetadata struct
	vxMetaDataStruct := "type VxMetadata struct {\n" +
		"    ID string `json:\"id\"`\n" +
		"    Subset string `json:\"met\"`\n" +
		"    Version string `json:\"version\"`\n" +
		"    Type string `json:\"type\"`\n" +
		"    SubType string `json:\"subtype\"`\n" +
		"}\n"

	fmt.Println("package structColumnTypes")
	fmt.Println("")
	fmt.Println("import \"strconv\"")
	fmt.Println("")

	fmt.Println(`type ColumnDef struct {
	    Name string
	    Doc interface{}
	}`)

	fmt.Println("")
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

	// print the data structs
	fmt.Println("")
	fmt.Println("//line data struct definitions")
	for _, dataStruct := range dataStructs {
		fmt.Println(dataStruct)
	}

	// print the doc structs
	fmt.Println("")
	fmt.Println("//doc struct definitions")
	for _, docStruct := range docStructs {
		fmt.Println(docStruct)
	}

	// print the dataKeys
	fmt.Println("")
	fmt.Println("//data key definitions")
	fmt.Println(dataKeyMapString)
	fmt.Println("}")

	// print the parserMap
	fmt.Println("")
	fmt.Println("//ParserMap declaration")

	fmt.Println("var ParserMap = map[string]ColumnDef{")
	for _, dataStruct := range dataStructs {
		fmt.Println("    \"" + strings.Split(dataStruct, " ")[1] + "\": {")
		fmt.Println("		Name: \"" + strings.Split(dataStruct, " ")[1] + "\",")
		fmt.Println("		Doc:  " + strings.Split(dataStruct, " ")[1] + "_doc{},")
		fmt.Println("	},")
	}
	fmt.Println("}")
}
