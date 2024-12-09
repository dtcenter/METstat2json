package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"net/http"
	"regexp"
)

func findType(name string, data []byte) (string, error){
	r, _ := regexp.Compile("ato.*get_item.*" + name)
	res := r.FindSubmatch(data)
	if res == nil {
		return "", fmt.Errorf("could not find type for %s", name)
	}
	dataType := strings.Split(string(res[0]),"(")[0]
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
	type DataStructName string
	type DataStructs map[DataStructName]string

	// read the header columns file
	headerColumnsFileUrl := "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
	resp, err := http.Get(headerColumnsFileUrl)
	if err != nil {
		fmt.Println("error getting columns file" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rawBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading columns file" + err.Error())
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
		fmt.Println("error reading parse_stat_line file" + err.Error())
		os.Exit(1)
	}

	// Use a map to keep track of unique headerStructs.
	var dataStructs = make(DataStructs)
	var headerStructs = make(HeaderStructs)

	file_lines := strings.Split(string(rawBytes), "\n")
	for _, line := range file_lines {
		// get the prefix from the line
		parts := strings.Split(line, ": VERSION")
		if len(parts) < 2 {
			if line != "" {
				fmt.Println("error parsing line: " + "line:'" + line + "'")
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
		version := strings.Replace(strings.TrimSpace(parts[0]), ".", "_", -1)
		fileType := strings.TrimSpace(parts[1])
		lineType := strings.TrimSpace(parts[2])
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
		headerStructName := HeaderStructName(fmt.Sprintf("%s_header_%s", fileType, version))
		headerStructString := fmt.Sprintf("type %s struct {\n", headerStructName)
		for _, term := range headerFields {
			// get the dataType from the cc file
			dataType := "string"
			lowerName := strings.ToLower(term)
			name := strings.ToUpper(lowerName[:1]) + lowerName[1:]
			headerStructString += fmt.Sprintf("    %s %s `json:\"%s\"`\n", name, dataType, term)
		}
		headerStructString += "}\n"
		headerStructString += "\n"
		// add the header struct to the map for printing later
		headerStructs[headerStructName] = headerStructString
		// create the data struct for this line type
		dataStructName := DataStructName(fmt.Sprintf("%s_%s_%s", fileType, lineType, version))
		dataStruct := fmt.Sprintf("%s struct {\n    vxMetaData\n    %s\n", dataStructName, headerStructName)
		for _, term := range dataFields {
			lowerName := strings.ToLower(term)
			name := strings.ToUpper(lowerName[:1]) + lowerName[1:]
			dataType := "string"
			foundType, err := findType(term, rawStatBytes)
			if err == nil {
				dataType = foundType
			}
			dataStruct += fmt.Sprintf("    %s %s `json:\"%s\"`\n", name, dataType, term)
		}
		dataStruct += "}\n"
		dataStruct += "\n"
		//fmt.Println(dataStruct)
		dataStructs[dataStructName] = dataStruct
	}

	// We have to flesh this out with the vxMetaData struct
	vxMetaDataStruct := "type vxMetadata struct {" +
	    "    ID string `json:\"ID\"`\n" +
		"    subset string `json:\"MET\"`\n" +
		"    version string `json:\"version\"`\n" +
		"    type string `json:\"type\"`\n" +
		"}\n"

	// print the vxMetaData struct
	fmt.Println(vxMetaDataStruct)

	// print the header structs
	for _, headerStruct := range headerStructs {
		fmt.Println(headerStruct)
	}
	// print the data structs
	for _, dataStruct := range dataStructs {
		fmt.Println(dataStruct)
	}
}
